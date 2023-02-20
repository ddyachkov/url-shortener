package storage

import (
	"errors"
)

type URLStorage struct {
	urls   map[int]string
	ids    map[string]int
	lastID int
}

func NewURLStorage() URLStorage {
	return URLStorage{
		urls:   make(map[int]string),
		ids:    make(map[string]int),
		lastID: 0,
	}
}

func (s *URLStorage) WriteData(url string) (id int, err error) {
	id, ok := s.ids[url]
	if ok {
		return id, nil
	}
	s.lastID += 1
	id = s.lastID
	s.urls[id] = url
	s.ids[url] = id

	return id, nil
}

func (s URLStorage) GetData(id int) (url string, err error) {
	url, ok := s.urls[id]
	if !ok {
		return "", errors.New("URL not found")
	}
	return url, nil
}
