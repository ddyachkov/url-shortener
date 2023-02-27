package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ddyachkov/url-shortener/internal/config"
)

type URLFileStorage struct {
	config     *config.ServerConfig
	urls       map[int]string
	ids        map[string]int
	users      map[int][]URLData
	lastDataID int
	lastUserID int
}

func NewURLFileStorage(cfg *config.ServerConfig) URLFileStorage {
	return URLFileStorage{
		config:     cfg,
		urls:       make(map[int]string),
		ids:        make(map[string]int),
		users:      make(map[int][]URLData),
		lastDataID: 0,
		lastUserID: 0,
	}
}

func (s *URLFileStorage) WriteData(ctx context.Context, url string, userID int) (dataID int, err error) {
	dataID, ok := s.ids[url]
	if ok {
		return dataID, errors.New("Conflict")
	}
	s.lastDataID += 1
	dataID = s.lastDataID
	s.urls[dataID] = url
	s.ids[url] = dataID
	s.users[userID] = append(s.users[userID], URLData{ID: dataID, OriginalURL: url})
	s.saveData(dataID, url, userID)

	return dataID, nil
}

func (s *URLFileStorage) WriteBatchData(ctx context.Context, batchData []URLData, userID int) (err error) {
	for i := range batchData {
		s.lastDataID += 1
		batchData[i].ID = s.lastDataID
		s.urls[batchData[i].ID] = batchData[i].OriginalURL
		s.ids[batchData[i].OriginalURL] = batchData[i].ID
		s.users[userID] = append(s.users[userID], batchData[i])
		s.saveData(batchData[i].ID, batchData[i].OriginalURL, userID)
	}

	return nil
}

func (s URLFileStorage) GetData(ctx context.Context, dataID int) (url string, err error) {
	url, ok := s.urls[dataID]
	if !ok {
		return "", errors.New("URL not found")
	}
	return url, nil
}

func (s URLFileStorage) CheckUser(ctx context.Context, userID int) (exists bool, err error) {
	_, exists = s.users[userID]
	return exists, nil
}

func (s *URLFileStorage) MakeNewUser(ctx context.Context) (userID int, err error) {
	s.lastUserID += 1
	return s.lastUserID, nil
}

func (s URLFileStorage) GetUserURL(ctx context.Context, userID int) (urlData []URLData, err error) {
	return s.users[userID], nil
}

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
