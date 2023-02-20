package storage

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ddyachkov/url-shortener/internal/config"
)

type URLFileStorage struct {
	urls   map[int]string
	ids    map[string]int
	lastID int
	config *config.ServerConfig
}

func NewURLFileStorage(cfg *config.ServerConfig) URLFileStorage {
	return URLFileStorage{
		urls:   make(map[int]string),
		ids:    make(map[string]int),
		lastID: 0,
		config: cfg,
	}
}

func (s *URLFileStorage) WriteData(url string) (id int, err error) {
	id, ok := s.ids[url]
	if ok {
		return id, nil
	}
	s.lastID += 1
	id = s.lastID
	s.urls[id] = url
	s.ids[url] = id
	s.saveData(id, url)

	return id, nil
}

func (s URLFileStorage) GetData(id int) (url string, err error) {
	url, ok := s.urls[id]
	if !ok {
		return "", errors.New("URL not found")
	}
	return url, nil
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

	var id int
	var url string
	for {
		_, err := fmt.Fscanf(file, "%d %s\n", &id, &url)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err.Error())
			return
		}
		s.urls[id] = url
		s.ids[url] = id
	}
	s.lastID = id
}

func (s URLFileStorage) saveData(id int, url string) {
	file, err := os.OpenFile(s.config.FileStoragePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%d %s\n", id, url)
}
