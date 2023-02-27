package app

import (
	"context"
	"net/url"

	"github.com/ddyachkov/url-shortener/internal/storage"
)

type URLShortener struct {
	storage storage.URLStorage // URL storage
}

// NewURLShortener() returns a new URLShortener object.
func NewURLShortener(st storage.URLStorage) URLShortener {
	return URLShortener{
		storage: st,
	}
}

// ReturnURI returns URI for received URL. If URL is invalid then returns error.
func (sh *URLShortener) ReturnURI(ctx context.Context, fullURL string, userID int) (uri string, err error) {
	_, err = url.ParseRequestURI(fullURL)
	if err != nil {
		return "", err
	}

	id, err := sh.storage.WriteData(ctx, fullURL, userID)
	if id == 0 {
		return "", err
	}

	return makeURI(id), err
}

func (sh *URLShortener) ReturnBatchURI(ctx context.Context, batchData []storage.URLData, userID int) (err error) {
	for i := range batchData {
		_, err = url.ParseRequestURI(batchData[i].OriginalURL)
		if err != nil {
			return err
		}
	}

	err = sh.storage.WriteBatchData(ctx, batchData, userID)
	if err != nil {
		return err
	}

	for i := range batchData {
		batchData[i].URI = makeURI(batchData[i].ID)
	}

	return nil
}

// GetFullURL returns full URL for received URI. IF URL is not found then returns error.
func (sh *URLShortener) GetFullURL(ctx context.Context, uri string) (fullURL string, err error) {
	id := makeID(uri)
	fullURL, err = sh.storage.GetData(ctx, id)
	if err != nil {
		return "", err
	}
	return fullURL, nil
}

func (sh URLShortener) GetUser(ctx context.Context, userID *int) (new bool, err error) {
	exists, err := sh.storage.CheckUser(ctx, *userID)
	if err != nil {
		return false, err
	}

	if exists {
		return false, nil
	}

	*userID, err = sh.storage.MakeNewUser(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (sh URLShortener) GetURLByUser(ctx context.Context, userID int) (urlData []storage.URLData, err error) {
	urlData, err = sh.storage.GetUserURL(ctx, userID)
	if err != nil {
		return nil, err
	}

	for i := range urlData {
		urlData[i].URI = makeURI(urlData[i].ID)
	}

	return urlData, nil
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
