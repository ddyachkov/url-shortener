package handlers

import (
	"io"
	"net/http"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/storage"
)

type URLHandler struct {
	shortener app.URLShortener //shortener logic
}

// NewURLHandler returns a new URLHandler object that implements the Handler interface.
func NewURLHandler() URLHandler {
	storage := storage.NewURLStorage()
	shortener := app.NewURLShortener(&storage)
	return URLHandler{
		shortener: shortener,
	}
}

// ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.
func (handler URLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		url, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusBadRequest, err.Error())
			return
		}
		uri, err := handler.shortener.ReturnURI(string(url))
		if err != nil {
			writeResponce(w, http.StatusBadRequest, err.Error())
			return
		}
		writeResponce(w, http.StatusCreated, "http://localhost:8080/"+uri)
	case http.MethodGet:
		uri := r.URL.Path[1:]
		url, err := handler.shortener.GetFullURL(uri)
		if err != nil {
			writeResponce(w, http.StatusBadRequest, err.Error())
			return
		}
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

// writeResponce writes http code to header and text to body.
func writeResponce(w http.ResponseWriter, code int, text string) {
	w.WriteHeader(code)
	w.Write([]byte(text))
}
