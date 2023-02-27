package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLDBStorage struct {
	db *pgxpool.Pool
}

func NewURLDBStorage(dbpool *pgxpool.Pool) URLDBStorage {
	return URLDBStorage{
		db: dbpool,
	}
}

func (s URLDBStorage) WriteData(ctx context.Context, url string, userID int) (dataID int, err error) {
	err = s.db.QueryRow(ctx, "INSERT INTO public.url_data (url, user_id) VALUES ($1, $2) RETURNING id", url, userID).Scan(&dataID)
	if err != nil {
		/* Не работает, хз почему
		var pgErr pgx.PgError
		if errors.As(err, &pgErr) && pgErr.code == pgerrcode.UniqueViolation{
			err = s.db.QueryRow(ctx, "SELECT id FROM public.url_data ud WHERE ud.url = $1", url).Scan(&dataID)
			if err != nil {
				return 0, err
			}
			return dataID, nil
		}
		*/
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"url_data_url_key\" (SQLSTATE 23505)" {
			err = s.db.QueryRow(ctx, "SELECT id FROM public.url_data ud WHERE ud.url = $1", url).Scan(&dataID)
			if err != nil {
				return 0, err
			}
			return dataID, errors.New("Conflict")
		}
		return 0, err
	}

	return dataID, nil
}

func (s URLDBStorage) WriteBatchData(ctx context.Context, batchData []URLData, userID int) (err error) {
	query := `INSERT INTO public.url_data (url, user_id) VALUES (@url, @userID) RETURNING id`

	batch := &pgx.Batch{}
	for i := range batchData {
		args := pgx.NamedArgs{
			"url":    batchData[i].OriginalURL,
			"userID": userID,
		}
		batch.Queue(query, args)
	}

	results := s.db.SendBatch(ctx, batch)
	defer results.Close()

	for i := range batchData {
		err := results.QueryRow().Scan(batchData[i].ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s URLDBStorage) GetData(ctx context.Context, dataID int) (url string, err error) {
	err = s.db.QueryRow(ctx, "SELECT ud.url FROM public.url_data ud WHERE ud.id = $1", dataID).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s URLDBStorage) CheckUser(ctx context.Context, userID int) (exists bool, err error) {
	err = s.db.QueryRow(ctx, "SELECT EXISTS (SELECT 1 FROM public.user u WHERE u.id = $1)", userID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s URLDBStorage) MakeNewUser(ctx context.Context) (userID int, err error) {
	err = s.db.QueryRow(ctx, "INSERT INTO public.user DEFAULT VALUES RETURNING id").Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

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

func (s URLDBStorage) Prepare(ctx context.Context) (err error) {
	_, err = s.db.Exec(ctx, "CREATE TABLE IF NOT EXISTS public.user (id SERIAL PRIMARY KEY)")
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, "CREATE TABLE IF NOT EXISTS public.url_data (id SERIAL PRIMARY KEY, url text UNIQUE NOT NULL, user_id integer REFERENCES public.user (id) NOT NULL)")
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, "CREATE INDEX IF NOT EXISTS idx_ud_user_id on public.url_data(user_id)")
	if err != nil {
		return err
	}

	return nil
}
