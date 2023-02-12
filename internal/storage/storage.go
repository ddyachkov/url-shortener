package storage

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ddyachkov/url-shortener/internal/config"
)

type URLStorage struct {
	urls   map[int]string
	ids    map[string]int
	lastID int
	config config.StorageConfig
}

func NewURLStorage() URLStorage {
	return URLStorage{
		urls:   make(map[int]string),
		ids:    make(map[string]int),
		lastID: 0,
		config: config.GetStorageConfig(),
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

func (s *URLStorage) LoadData() {
	file, err := os.Open(s.config.StoragePath + "/data.txt")
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

func (s URLStorage) SaveData() {
	if _, err := os.Stat(s.config.StoragePath); os.IsNotExist(err) {
		err := os.Mkdir(s.config.StoragePath, os.ModePerm)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
	file, err := os.OpenFile(s.config.StoragePath+"/data.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()

	for id, url := range s.urls {
		fmt.Fprintf(file, "%d %s\n", id, url)
	}
}
