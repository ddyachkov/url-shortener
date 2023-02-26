package storage

import (
	"github.com/jackc/pgx"
)

type URLDBStorage struct {
	conn *pgx.Conn
}

func NewURLDBStorage(c *pgx.Conn) URLDBStorage {
	return URLDBStorage{
		conn: c,
	}
}

func (s URLDBStorage) WriteData(url string, userID int) (dataID int, err error) {
	err = s.conn.QueryRow("INSERT INTO public.url_data (url, user_id) VALUES ($1, $2) RETURNING id", url, userID).Scan(&dataID)
	if err != nil {
		return 0, err
	}

	return dataID, nil
}

func (s URLDBStorage) GetData(dataID int) (url string, err error) {
	err = s.conn.QueryRow("SELECT ud.url FROM public.url_data ud WHERE ud.id = $1", dataID).Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s URLDBStorage) CheckUser(userID int) (exists bool, err error) {
	err = s.conn.QueryRow("SELECT EXISTS (SELECT 1 FROM public.user u WHERE u.id = $1)", userID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s URLDBStorage) MakeNewUser() (userID int, err error) {
	err = s.conn.QueryRow("INSERT INTO public.user DEFAULT VALUES RETURNING id").Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s URLDBStorage) GetUserURL(userID int) (urlData []URLData, err error) {
	rows, err := s.conn.Query("SELECT ud.id, ud.url FROM public.url_data ud WHERE ud.user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ud URLData
		err = rows.Scan(&ud.ID, &ud.URL)
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

func (s URLDBStorage) Prepare() (err error) {
	_, err = s.conn.Exec("CREATE TABLE IF NOT EXISTS public.user (id SERIAL PRIMARY KEY)")
	if err != nil {
		return err
	}

	_, err = s.conn.Exec("CREATE TABLE IF NOT EXISTS public.url_data (id SERIAL PRIMARY KEY, url text NOT NULL, user_id integer REFERENCES public.user (id))")
	if err != nil {
		return err
	}

	_, err = s.conn.Exec("CREATE INDEX IF NOT EXISTS idx_ud_url on public.url_data(url)")
	if err != nil {
		return err
	}

	_, err = s.conn.Exec("CREATE INDEX IF NOT EXISTS idx_ud_user_id on public.url_data(user_id)")
	if err != nil {
		return err
	}

	return nil
}
