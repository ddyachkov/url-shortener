package storage

import (
	"context"
)

type URLMemStorage struct {
	urls       map[int]string
	ids        map[string]int
	users      map[int][]URLData
	lastDataID int
	lastUserID int
}

func NewURLMemStorage() *URLMemStorage {
	return &URLMemStorage{
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
		return dataID, ErrWriteDataConflict
	}
	s.lastDataID += 1
	dataID = s.lastDataID
	s.urls[dataID] = url
	s.ids[url] = dataID
	s.users[userID] = append(s.users[userID], URLData{ID: dataID, OriginalURL: url})

	return dataID, nil
}

func (s *URLMemStorage) WriteBatchData(ctx context.Context, batchURL []string, userID int) (batchID []int, err error) {
	batchID = make([]int, 0)
	for i := range batchURL {
		s.lastDataID += 1
		batchID = append(batchID, s.lastDataID)
		s.urls[s.lastDataID] = batchURL[i]
		s.ids[batchURL[i]] = s.lastDataID
		s.users[userID] = append(s.users[userID], URLData{ID: s.lastDataID, OriginalURL: batchURL[i]})
	}

	return batchID, nil
}

func (s URLMemStorage) GetData(ctx context.Context, dataID int) (url string, err error) {
	url, ok := s.urls[dataID]
	if !ok {
		return "", ErrURLNotFound
	}
	if url == "" {
		return "", ErrURLIsDeleted
	}
	return url, nil
}

func (s *URLMemStorage) MakeNewUser(ctx context.Context) (userID int, err error) {
	s.lastUserID += 1
	return s.lastUserID, nil
}

func (s URLMemStorage) CheckUser(ctx context.Context, searchID int) (foundID int, err error) {
	if _, exists := s.users[searchID]; exists {
		return searchID, nil
	}
	return s.MakeNewUser(ctx)
}

func (s URLMemStorage) GetUserURL(ctx context.Context, userID int) (urlData []URLData, err error) {
	return s.users[userID], nil
}

func (s URLMemStorage) DeleteBatchData(ctx context.Context, batchID []int, userID int) {
	for _, id := range batchID {
		urlData := s.users[userID]
		found := false
		for _, data := range urlData {
			if data.ID == id {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		url := s.urls[id]
		s.urls[id] = ""
		delete(s.ids, url)
	}
}
