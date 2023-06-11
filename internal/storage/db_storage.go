package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// URLDBStorage stores pointer to pg connection pool object.
type URLDBStorage struct {
	db *pgxpool.Pool
}

// NewURLDBStorage returns a new URLDBStorage object.
func NewURLDBStorage(dbpool *pgxpool.Pool, ctx context.Context) (storage *URLDBStorage, err error) {
	storage = &URLDBStorage{
		db: dbpool,
	}

	err = storage.Prepare(ctx)
	if err != nil {
		return storage, err
	}

	return storage, nil
}

// WriteData writes URL data to storage and returns new data id.
func (s URLDBStorage) WriteData(ctx context.Context, url string, userID int) (dataID int, err error) {
	err = s.db.QueryRow(ctx, "INSERT INTO public.url_data (url, user_id) VALUES ($1, $2) RETURNING id", url, userID).Scan(&dataID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			err = s.db.QueryRow(ctx, "SELECT id FROM public.url_data ud WHERE ud.url = $1", url).Scan(&dataID)
			if err != nil {
				return 0, err
			}
			return dataID, ErrWriteDataConflict
		}
		return 0, err
	}

	return dataID, nil
}

// WriteData writes batch of URLs data to storage and returns new data ids.
func (s URLDBStorage) WriteBatchData(ctx context.Context, batchURL []string, userID int) (batchID []int, err error) {
	query := `INSERT INTO public.url_data (url, user_id) VALUES (@url, @userID) RETURNING id`

	batch := &pgx.Batch{}
	for i := range batchURL {
		args := pgx.NamedArgs{
			"url":    batchURL[i],
			"userID": userID,
		}
		batch.Queue(query, args)
	}

	results := s.db.SendBatch(ctx, batch)
	defer results.Close()

	batchID = make([]int, 0)
	for i := 0; i < len(batchURL); i++ {
		var dataID int
		err := results.QueryRow().Scan(&dataID)
		if err != nil {
			return nil, err
		}
		batchID = append(batchID, dataID)
	}

	return batchID, nil
}

// GetData returns URL data by data id.
func (s URLDBStorage) GetData(ctx context.Context, dataID int) (url string, err error) {
	var deleted bool
	err = s.db.QueryRow(ctx, "SELECT ud.url, ud.deleted FROM public.url_data ud WHERE ud.id = $1", dataID).Scan(&url, &deleted)
	if err != nil {
		return "", ErrURLNotFound
	}

	if deleted {
		return "", ErrURLIsDeleted
	}

	return url, nil
}

// CheckUser receives searchID and checks if user is already exists. Returns same id for existing user or new id for new user.
func (s URLDBStorage) CheckUser(ctx context.Context, searchID int) (foundID int, err error) {
	var exists bool
	err = s.db.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM public.user u WHERE u.id = $1)", searchID).Scan(&exists)
	if err != nil {
		return searchID, err
	}

	if exists {
		return searchID, nil
	}

	return s.makeNewUser(ctx)
}

// GetUserURL returns batch of URL data by user id.
func (s URLDBStorage) GetUserURL(ctx context.Context, userID int) (urlData []URLData, err error) {
	rows, err := s.db.Query(ctx, "SELECT ud.id, ud.url FROM public.url_data ud WHERE ud.user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ud URLData
		err = rows.Scan(&ud.ID, &ud.OriginalURL)
		if err != nil {
			return nil, err
		}

		urlData = append(urlData, ud)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return urlData, nil
}

// DeleteBatchData deletes batch of URL data by user id.
func (s URLDBStorage) DeleteBatchData(ctx context.Context, batchID []int, userID int) {
	query := `UPDATE public.url_data SET deleted = true WHERE id = @id and user_id = @userID`

	batch := &pgx.Batch{}
	for i := range batchID {
		args := pgx.NamedArgs{
			"id":     batchID[i],
			"userID": userID,
		}
		batch.Queue(query, args)
	}

	results := s.db.SendBatch(ctx, batch)
	results.Close()
}

// GetStats returns total count of short URLs and users
func (s URLDBStorage) GetStats(ctx context.Context) (cURLs, cUsers int, err error) {
	err = s.db.QueryRow(ctx, "SELECT count(ud.id) FROM public.url_data ud WHERE NOT ud.deleted").Scan(&cURLs)
	if err != nil {
		return 0, 0, err
	}

	err = s.db.QueryRow(ctx, "SELECT count(u.id) FROM public.user u").Scan(&cUsers)
	if err != nil {
		return 0, 0, err
	}

	return cURLs, cUsers, nil
}

// Prepare prepares db storage to work with.
func (s URLDBStorage) Prepare(ctx context.Context) (err error) {
	_, err = s.db.Exec(ctx, "CREATE TABLE IF NOT EXISTS public.user (id SERIAL PRIMARY KEY)")
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, "CREATE TABLE IF NOT EXISTS public.url_data (id SERIAL PRIMARY KEY, url TEXT UNIQUE NOT NULL, user_id INTEGER REFERENCES public.user (id) NOT NULL, deleted BOOLEAN DEFAULT FALSE)")
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, "CREATE INDEX IF NOT EXISTS idx_ud_user_id ON public.url_data(user_id)")
	if err != nil {
		return err
	}

	return nil
}

func (s URLDBStorage) makeNewUser(ctx context.Context) (userID int, err error) {
	err = s.db.QueryRow(ctx, "INSERT INTO public.user DEFAULT VALUES RETURNING id").Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
