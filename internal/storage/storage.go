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
	s.saveData(id, url)

	return id, nil
}

func (s URLStorage) GetData(id int) (url string, err error) {
	url, ok := s.urls[id]
	if !ok {
		return "", errors.New("URL not found")
	}
	return url, nil
}

func (s *URLStorage) LoadData() {
	if _, err := os.Stat(config.FileStoragePath); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(config.FileStoragePath), os.ModePerm)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("Created", filepath.Dir(config.FileStoragePath), "directory")
	}

	file, err := os.Open(config.FileStoragePath)
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

func (s URLStorage) saveData(id int, url string) {
	file, err := os.OpenFile(config.FileStoragePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%d %s\n", id, url)
}
