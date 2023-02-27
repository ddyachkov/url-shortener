package storage

import (
	"context"
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

func (s *URLMemStorage) WriteData(ctx context.Context, url string, userID int) (dataID int, err error) {
	dataID, ok := s.ids[url]
	if ok {
		return dataID, errors.New("Conflict")
	}
	s.lastDataID += 1
	dataID = s.lastDataID
	s.urls[dataID] = url
	s.ids[url] = dataID
	s.users[userID] = append(s.users[userID], URLData{ID: dataID, OriginalURL: url})

	return dataID, nil
}

func (s *URLMemStorage) WriteBatchData(ctx context.Context, batchData []URLData, userID int) (err error) {
	for i := range batchData {
		s.lastDataID += 1
		batchData[i].ID = s.lastDataID
		s.urls[batchData[i].ID] = batchData[i].OriginalURL
		s.ids[batchData[i].OriginalURL] = batchData[i].ID
		s.users[userID] = append(s.users[userID], batchData[i])
	}

	return nil
}

func (s URLMemStorage) GetData(ctx context.Context, dataID int) (url string, err error) {
	url, ok := s.urls[dataID]
	if !ok {
		return "", errors.New("URL not found")
	}
	return url, nil
}

func (s *URLMemStorage) MakeNewUser(ctx context.Context) (userID int, err error) {
	s.lastUserID += 1
	return s.lastUserID, nil
}

func (s URLMemStorage) CheckUser(ctx context.Context, userID int) (exists bool, err error) {
	_, exists = s.users[userID]
	return exists, nil
}

func (s URLMemStorage) GetUserURL(ctx context.Context, userID int) (urlData []URLData, err error) {
	return s.users[userID], nil
}
