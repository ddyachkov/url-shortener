package handlers

import (
	"io"
	"net/http"

	"github.com/ddyachkov/url-shortener/internal/app"
)

type URLPostHandler struct {
	shortener app.URLShortener //shortener logic
}

type URLGetHandler struct {
	shortener app.URLShortener //shortener logic
}

// NewURLPostHandler returns a new URLPostHandler object that implements the Handler interface.
func NewURLPostHandler(storage app.Storage) URLPostHandler {
	shortener := app.NewURLShortener(storage)
	return URLPostHandler{
		shortener: shortener,
	}
}

// NewURLGetHandler returns a new URLGetHandler object that implements the Handler interface.
func NewURLGetHandler(storage app.Storage) URLGetHandler {
	shortener := app.NewURLShortener(storage)
	return URLGetHandler{
		shortener: shortener,
	}
}

// ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.
func (handler URLPostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uri, err := handler.shortener.ReturnURI(string(url))
	if err != nil {
		writeResponce(w, http.StatusBadRequest, err.Error())
		return
	}
	writeResponce(w, http.StatusCreated, "http://localhost:8080/"+uri)
}

// ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.
func (handler URLGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Path[1:]
	url, err := handler.shortener.GetFullURL(uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// writeResponce writes http code to header and text to body.
func writeResponce(w http.ResponseWriter, code int, text string) {
	w.WriteHeader(code)
	w.Write([]byte(text))
}
