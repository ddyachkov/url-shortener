package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/ddyachkov/url-shortener/internal/middleware"
	"github.com/go-chi/chi"
)

type handler struct {
	service *app.URLShortener
	config  *config.ServerConfig
}

func NewURLHandler(shortener *app.URLShortener, cfg *config.ServerConfig) http.Handler {
	router := chi.NewRouter()

	h := handler{
		service: shortener,
		config:  cfg,
	}

	router.Use(middleware.Decompress)
	router.Use(middleware.Compress)

	router.Post("/", h.ReturnTextShortURL)
	router.Post("/api/shorten", h.ReturnJSONShortURL)

	router.Get("/{URI}", h.RedirectToFullURL)

	return router
}

func (h handler) ReturnTextShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uri, err := h.service.ReturnURI(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := url.JoinPath(h.config.BaseURL, uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("ReturnTextShortURL:", string(body), "->", shortURL)
	writeResponse(w, []byte(shortURL), http.StatusCreated)
}

func (h handler) ReturnJSONShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	requestBody := struct {
		URL string `json:"url"`
	}{}
	if err = json.Unmarshal(body, &requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uri, err := h.service.ReturnURI(requestBody.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := url.JoinPath(h.config.BaseURL, uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responceBody := struct {
		Result string `json:"result"`
	}{Result: shortURL}
	responce, err := json.Marshal(responceBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("ReturnJSONShortURL:", requestBody.URL, "->", shortURL)
	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, responce, http.StatusCreated)
}

func (h handler) RedirectToFullURL(w http.ResponseWriter, r *http.Request) {
	uri := chi.URLParam(r, "URI")
	fullURL, err := h.service.GetFullURL(uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	shortURL, err := url.JoinPath(h.config.BaseURL, uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("RedirectToFullURL:", shortURL, "->", fullURL)
	http.Redirect(w, r, fullURL, http.StatusTemporaryRedirect)
}

func writeResponse(w http.ResponseWriter, text []byte, code int) {
	w.WriteHeader(code)
	w.Write(text)
}
