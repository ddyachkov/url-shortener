package storage

import "errors"

type URLStorage struct {
	urls   map[int]string // url storage
	ids    map[string]int // id storage
	lastID int            // last used ID
}

// NewURLStorage() returns a new URLStorage object that implements the Storage interface.
func NewURLStorage() URLStorage {
	return URLStorage{
		urls:   make(map[int]string),
		ids:    make(map[string]int),
		lastID: 0,
	}
}

// WriteData writes data into storage and returns a new ID. If data is already in storage, then return its ID.
func (storage *URLStorage) WriteData(url string) (id int, err error) {
	id, ok := storage.ids[url]
	if ok {
		return id
	}
	storage.lastID += 1
	id = storage.lastID
	storage.urls[id] = url
	storage.ids[url] = id

	return id, nil
}

// GetData returns data from storage. If data is not found then returns error.
func (storage URLStorage) GetData(id int) (url string, err error) {
	url, ok := storage.urls[id]
	if !ok {
		return "", errors.New("URL not found")
	}
	return url, nil
}
