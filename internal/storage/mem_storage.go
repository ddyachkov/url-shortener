package storage

import (
	"errors"
)

type URLMemStorage struct {
	urls       map[int]string
	ids        map[string]int
	users      map[int][]URLData
	lastDataID int
	lastUserID int
}

func NewURLMemStorage() URLMemStorage {
	return URLMemStorage{
		urls:       make(map[int]string),
		ids:        make(map[string]int),
		users:      make(map[int][]URLData),
		lastDataID: 0,
		lastUserID: 0,
	}
}

func (s *URLMemStorage) WriteData(url string, userID int) (dataID int, err error) {
	dataID, ok := s.ids[url]
	if ok {
		return dataID, nil
	}
	s.lastDataID += 1
	dataID = s.lastDataID
	s.urls[dataID] = url
	s.ids[url] = dataID
	s.users[userID] = append(s.users[userID], URLData{ID: dataID, OriginalURL: url})

	return dataID, nil
}

func (s URLMemStorage) GetData(dataID int) (url string, err error) {
	url, ok := s.urls[dataID]
	if !ok {
		return "", errors.New("URL not found")
	}
	return url, nil
}

func (s *URLMemStorage) MakeNewUser() (userID int, err error) {
	s.lastUserID += 1
	return s.lastUserID, nil
}

func (s URLMemStorage) CheckUser(userID int) (exists bool, err error) {
	_, exists = s.users[userID]
	return exists, nil
}

func (s URLMemStorage) GetUserURL(userID int) (urlData []URLData, err error) {
	return s.users[userID], nil
}
