// Package handler provides endpoints for url shortener server
package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"

	_ "net/http/pprof"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/ddyachkov/url-shortener/internal/middleware"
	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

type contextKey string

// String implements Stringer interface for contextKey type.
func (c contextKey) String() string {
	return string(c)
}

const contextUserIDKey contextKey = "user_id"

type handler struct {
	service *app.URLShortener
	config  *config.ServerConfig
	db      *pgxpool.Pool
}

// NewURLHandler returns http.Handler with url shortener endpoints.
func NewURLHandler(shortener *app.URLShortener, cfg *config.ServerConfig, dbpool *pgxpool.Pool) http.Handler {
	router := chi.NewRouter()

	h := handler{
		service: shortener,
		config:  cfg,
		db:      dbpool,
	}

	router.Group(func(router chi.Router) {
		router.Use(middleware.Decompress, middleware.Compress, h.GetEncryptedUserID)
		router.Get("/api/user/urls", h.GetUserURL)
		router.Delete("/api/user/urls", h.DeleteUserURL)

		router.Group(func(router chi.Router) {
			router.Use(h.SetEncryptedUserID)
			router.Post("/", h.ReturnTextShortURL)
			router.Post("/api/shorten", h.ReturnJSONShortURL)
			router.Post("/api/shorten/batch", h.ReturnBatchJSONShortURL)
		})

		router.Group(func(router chi.Router) {
			router.Use(h.CheckSubnet)
			router.Get("/api/internal/stats", h.GetStats)
		})
	})

	router.Get("/{URI}", h.RedirectToFullURL)
	router.Get("/ping", h.PingDatabase)

	router.Route("/debug/pprof", func(router chi.Router) {
		router.Handle("/", http.DefaultServeMux)
		router.Handle("/{cmd}", http.DefaultServeMux)
	})

	return router
}

// ReturnTextShortURL returns response with short URL for request with URL in simple text body.
func (h handler) ReturnTextShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.Context().Value(contextUserIDKey).(int)

	httpStatus := http.StatusCreated
	uri, err := h.service.ReturnURI(r.Context(), string(body), userID)
	if err != nil {
		if errors.Is(err, storage.ErrWriteDataConflict) {
			httpStatus = http.StatusConflict
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	shortURL, err := url.JoinPath(h.config.BaseURL, uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("ReturnTextShortURL:", string(body), "->", shortURL)
	writeResponse(w, []byte(shortURL), httpStatus)
}

// ReturnJSONShortURL returns response with short URL for request with URL in JSON body.
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

	userID := r.Context().Value(contextUserIDKey).(int)

	httpStatus := http.StatusCreated
	uri, err := h.service.ReturnURI(r.Context(), requestBody.URL, userID)
	if err != nil {
		if errors.Is(err, storage.ErrWriteDataConflict) {
			httpStatus = http.StatusConflict
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
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

	w.Header().Set("Content-Type", "application/json")
	log.Println("ReturnJSONShortURL:", requestBody.URL, "->", shortURL)
	writeResponse(w, responce, httpStatus)
}

// ReturnJSONShortURL returns response with batch of short URLs for request with URLs in JSON body.
func (h handler) ReturnBatchJSONShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	batchData := []storage.URLData{}
	if err = json.Unmarshal(body, &batchData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(contextUserIDKey).(int)

	batchURL := make([]string, 0)
	for i := range batchData {
		batchURL = append(batchURL, batchData[i].OriginalURL)
	}
	batchURI, err := h.service.ReturnBatchURI(r.Context(), batchURL, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range batchURI {
		shortURL, err := url.JoinPath(h.config.BaseURL, batchURI[i])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		batchData[i].ShortURL = shortURL
		batchData[i].OriginalURL = ""
	}

	responce, err := json.Marshal(batchData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, responce, http.StatusCreated)
}

// RedirectToFullURL returns redirect from short URL to full URL.
func (h handler) RedirectToFullURL(w http.ResponseWriter, r *http.Request) {
	uri := chi.URLParam(r, "URI")
	fullURL, err := h.service.GetFullURL(r.Context(), uri)
	if err != nil {
		switch err {
		case storage.ErrURLNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case storage.ErrURLIsDeleted:
			http.Error(w, err.Error(), http.StatusGone)
		}
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

// GetUserURL returns batch of URLs shortened by user.
func (h handler) GetUserURL(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(contextUserIDKey).(int)
	if userID == 0 {
		http.Error(w, "[]", http.StatusNoContent)
	}

	responceBody, err := h.service.GetURLByUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(responceBody) == 0 {
		http.Error(w, "[]", http.StatusNoContent)
		return
	}

	for i := range responceBody {
		shortURL, err := url.JoinPath(h.config.BaseURL, responceBody[i].URI)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responceBody[i].ShortURL = shortURL
	}

	responce, err := json.Marshal(responceBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("GetUserURL:", userID)
	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, responce, http.StatusOK)
}

// DeleteUserURL deletes batch of URLs shortened by user.
func (h handler) DeleteUserURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var uriList []string
	err = json.Unmarshal(body, &uriList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.Context().Value(contextUserIDKey).(int)

	go func() {
		h.service.DeleteUserURL(context.Background(), uriList, userID)
	}()
	writeResponse(w, []byte("Accepted"), http.StatusAccepted)
}

// GetStats returns total count of short URLs and users
func (h handler) GetStats(w http.ResponseWriter, r *http.Request) {
	cURLs, cUsers, err := h.service.GetStats(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responceBody := struct {
		URLs  int `json:"urls"`
		Users int `json:"users"`
	}{URLs: cURLs, Users: cUsers}
	responce, err := json.Marshal(responceBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	log.Println("GetStats: URLs -", responceBody.URLs, ", Users -", responceBody.Users)
	writeResponse(w, responce, http.StatusOK)
}

// PingDatabase provides "ping DB" functionality.
func (h handler) PingDatabase(w http.ResponseWriter, r *http.Request) {
	if h.db == nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err := h.db.Ping(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, []byte(""), http.StatusOK)
}

func writeResponse(w http.ResponseWriter, text []byte, code int) {
	w.WriteHeader(code)
	w.Write(text)
}
