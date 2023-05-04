package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ddyachkov/url-shortener/internal/config"
)

// URLFileStorage stores file storage data.
type URLFileStorage struct {
	config     *config.ServerConfig
	urls       map[int]string
	ids        map[string]int
	users      map[int][]URLData
	lastDataID int
	lastUserID int
}

// NewURLFileStorage returns a new URLFileStorage object.
func NewURLFileStorage(cfg *config.ServerConfig) (storage *URLFileStorage) {
	storage = &URLFileStorage{
		config:     cfg,
		urls:       make(map[int]string),
		ids:        make(map[string]int),
		users:      make(map[int][]URLData),
		lastDataID: 0,
		lastUserID: 0,
	}
	storage.LoadData()

	return storage
}

// WriteData writes URL data to storage and returns new data id.
func (s *URLFileStorage) WriteData(ctx context.Context, url string, userID int) (dataID int, err error) {
	dataID, ok := s.ids[url]
	if ok {
		return dataID, ErrWriteDataConflict
	}
	s.lastDataID += 1
	dataID = s.lastDataID
	s.urls[dataID] = url
	s.ids[url] = dataID
	s.users[userID] = append(s.users[userID], URLData{ID: dataID, OriginalURL: url})
	s.saveData(dataID, url, userID)

	return dataID, nil
}

// WriteData writes batch of URLs data to storage and returns new data ids.
func (s *URLFileStorage) WriteBatchData(ctx context.Context, batchURL []string, userID int) (batchID []int, err error) {
	batchID = make([]int, 0)
	for i := range batchURL {
		s.lastDataID += 1
		batchID = append(batchID, s.lastDataID)
		s.urls[s.lastDataID] = batchURL[i]
		s.ids[batchURL[i]] = s.lastDataID
		s.users[userID] = append(s.users[userID], URLData{ID: s.lastDataID, OriginalURL: batchURL[i]})
		s.saveData(s.lastDataID, batchURL[i], userID)
	}

	return batchID, nil
}

// GetData returns URL data by data id.
func (s URLFileStorage) GetData(ctx context.Context, dataID int) (url string, err error) {
	url, ok := s.urls[dataID]
	if !ok {
		return "", ErrURLNotFound
	}
	if url == "" {
		return "", ErrURLIsDeleted
	}
	return url, nil
}

// CheckUser receives searchID and checks if user is already exists. Returns same id for existing user or new id for new user.
func (s URLFileStorage) CheckUser(ctx context.Context, searchID int) (foundID int, err error) {
	if _, exists := s.users[searchID]; exists {
		return searchID, nil
	}
	return s.makeNewUser(ctx)
}

// GetUserURL returns batch of URL data by user id.
func (s URLFileStorage) GetUserURL(ctx context.Context, userID int) (urlData []URLData, err error) {
	return s.users[userID], nil
}

// DeleteBatchData deletes batch of URL data by user id.
func (s URLFileStorage) DeleteBatchData(ctx context.Context, batchID []int, userID int) {
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

// LoadData loads data from file on disk.
func (s *URLFileStorage) LoadData() {
	if _, err := os.Stat(s.config.FileStoragePath); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(s.config.FileStoragePath), os.ModePerm)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}

	file, err := os.Open(s.config.FileStoragePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()

	var dataID int
	var url string
	var userID int
	for {
		_, err := fmt.Fscanf(file, "%d %s %d\n", &dataID, &url, &userID)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err.Error())
			return
		}
		s.urls[dataID] = url
		s.ids[url] = dataID
		s.users[userID] = append(s.users[userID], URLData{ID: dataID, OriginalURL: url})
		if dataID > s.lastDataID {
			s.lastDataID = dataID
		}
		if userID > s.lastUserID {
			s.lastUserID = userID
		}
	}
}

func (s URLFileStorage) saveData(dataID int, url string, userID int) {
	file, err := os.OpenFile(s.config.FileStoragePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%d %s %d\n", dataID, url, userID)
}

func (s *URLFileStorage) makeNewUser(ctx context.Context) (userID int, err error) {
	s.lastUserID += 1
	return s.lastUserID, nil
}
