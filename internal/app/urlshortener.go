package app

import (
	"net/url"
)

type Storage interface {
	WriteData(string, int) (int, error)
	GetData(int) (string, error)
	MakeNewUser() (int, error)
	GetUserData(int) ([]int, error)
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
func (sh *URLShortener) ReturnURI(fullURL string, userID int) (uri string, err error) {
	_, err = url.ParseRequestURI(fullURL)
	if err != nil {
		return "", err
	}
	id, err := sh.storage.WriteData(fullURL, userID)
	if err != nil {
		return "", err
	}
	return makeURI(id), nil
}

// GetFullURL returns full URL for received URI. IF URL is not found then returns error.
func (sh *URLShortener) GetFullURL(uri string) (fullURL string, err error) {
	id := makeID(uri)
	fullURL, err = sh.storage.GetData(id)
	if err != nil {
		return "", err
	}
	return fullURL, nil
}

func (sh URLShortener) GetNewUser() (userID int, err error) {
	userID, err = sh.storage.MakeNewUser()
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (sh URLShortener) GetURLByUser(userID int) (urls map[string]string, err error) {
	urls = make(map[string]string)
	dataIDs, err := sh.storage.GetUserData(userID)
	if err != nil {
		return nil, err
	}

	for _, dataID := range dataIDs {
		uri := makeURI(dataID)
		url, err := sh.storage.GetData(dataID)
		if err != nil {
			return nil, err
		}
		urls[uri] = url
	}

	return urls, nil
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
