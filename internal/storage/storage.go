package storage

import (
	"errors"
)

type URLStorage struct {
	urls       map[int]string
	ids        map[string]int
	users      map[int][]int
	lastDataID int
	lastUserID int
}

func NewURLStorage() URLStorage {
	return URLStorage{
		urls:       make(map[int]string),
		ids:        make(map[string]int),
		users:      make(map[int][]int),
		lastDataID: 0,
		lastUserID: 0,
	}
}

func (s *URLStorage) WriteData(url string, userID int) (dataID int, err error) {
	dataID, ok := s.ids[url]
	if ok {
		return dataID, nil
	}
	s.lastDataID += 1
	dataID = s.lastDataID
	s.urls[dataID] = url
	s.ids[url] = dataID
	s.users[userID] = append(s.users[userID], dataID)

	return dataID, nil
}

func (s URLStorage) GetData(dataID int) (url string, err error) {
	url, ok := s.urls[dataID]
	if !ok {
		return "", errors.New("URL not found")
	}
	return url, nil
}

func (s *URLStorage) MakeNewUser() (userID int, err error) {
	s.lastUserID += 1
	return s.lastUserID, nil
}

func (s URLStorage) GetUserData(userID int) (dataIDs []int, err error) {
	return s.users[userID], nil
}
