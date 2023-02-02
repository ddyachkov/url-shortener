package app

import (
	"errors"
	"regexp"
	"strings"
)

type Storage interface {
	WriteData(string) (int, error)
	GetData(int) (string, error)
}

type URLShortener struct {
	storage Storage // URL storage
}

// NewURLShortener() returns a new URLShortener object.
func NewURLShortener(st Storage) URLShortener {
	return URLShortener{
		storage: st,
	}
}

// ReturnURI returns URI for received URL. If URL is invalid then returns error.
func (shortener *URLShortener) ReturnURI(url string) (uri string, err error) {
	_, err = ParseRequestURI(url)
	if err != nil {
		return "", err
	}
	id, err := shortener.storage.WriteData(url)
	if err != nil {
		return "", err
	}
	return makeURI(id), nil
}

// GetFullURL returns full URL for received URI. IF URL is not found then returns error.
func (shortener *URLShortener) GetFullURL(uri string) (url string, err error) {
	id := makeID(uri)
	url, err = shortener.storage.GetData(id)
	if err != nil {
		return "", err
	}
	return url, nil
}

// makeURI returns URI from data ID
func makeURI(id int) (uri string) {
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	shortURL := make([]byte, 0)
	for id > 0 {
		shortURL = append(shortURL, characters[id%len(characters)])
		id = id / len(characters)
	}
	l := len(shortURL)
	for i := 0; i < l/2; i++ {
		shortURL[i], shortURL[l-1-i] = shortURL[l-1-i], shortURL[i]
	}
	return string(shortURL)
}

// makeID returns data ID from URI
func makeID(uri string) (id int) {
	for i := 0; i < len(uri); i++ {
		if 'a' <= uri[i] && uri[i] <= 'z' {
			id = id*62 + int(uri[i]) - 'a'
		}
		if 'A' <= uri[i] && uri[i] <= 'Z' {
			id = id*62 + int(uri[i]) - 'A' + 26
		}
		if '0' <= uri[i] && uri[i] <= '9' {
			id = id*62 + int(uri[i]) - '0' + 52
		}
	}
	return id
}
