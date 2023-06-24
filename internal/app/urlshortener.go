// Package app implements url shortener logic
package app

import (
	"context"
	"net/url"

	"github.com/ddyachkov/url-shortener/internal/storage"
)

// URLShortener is a middleware between handler and storage
type URLShortener struct {
	storage storage.URLStorage
}

// NewURLShortener returns a new URLShortener object.
func NewURLShortener(st storage.URLStorage) *URLShortener {
	return &URLShortener{
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

// ReturnBatchURI returns batch of URIs for received batch of URLs. If any URL is invalid then returns error.
func (sh *URLShortener) ReturnBatchURI(ctx context.Context, batchURL []string, userID int) (batchURI []string, err error) {
	for i := range batchURL {
		_, err = url.ParseRequestURI(batchURL[i])
		if err != nil {
			return nil, err
		}
	}

	batchID, err := sh.storage.WriteBatchData(ctx, batchURL, userID)
	if err != nil {
		return nil, err
	}

	batchURI = make([]string, 0)
	for i := range batchID {
		batchURI = append(batchURI, makeURI(batchID[i]))
	}

	return batchURI, nil
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

// GetUser receives searchID and checks if user is already exists. Returns same id for existing user or new id for new user.
func (sh URLShortener) GetUser(ctx context.Context, searchID int) (foundID int, err error) {
	return sh.storage.CheckUser(ctx, searchID)
}

// GetURLByUser returns all URLs shortened by user.
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

// DeleteUserURL deletes batch of URLs shortened by user.
func (sh URLShortener) DeleteUserURL(ctx context.Context, uriList []string, userID int) {
	idList := make([]int, len(uriList))
	for i, uri := range uriList {
		idList[i] = makeID(uri)
	}
	sh.storage.DeleteBatchData(ctx, idList, userID)
}

// GetStats returns total count of short URLs and users
func (sh URLShortener) GetStats(ctx context.Context) (cURLs, cUsers int, err error) {
	return sh.storage.GetStats(ctx)
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
