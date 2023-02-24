package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/ddyachkov/url-shortener/internal/cookie"
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
	router.Get("/api/user/urls", h.GetUserURL)

	return router
}

func (h handler) ReturnTextShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var userID int
	cookieValue, err := cookie.GetEncryptedValue(r, "user_id", []byte(h.config.SecretKey))
	if err != nil {
		userID, err = h.service.GetNewUser()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = cookie.WriteEncryptedValue(w, "user_id", strconv.Itoa(userID), []byte(h.config.SecretKey))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		userID, err = strconv.Atoi(cookieValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	uri, err := h.service.ReturnURI(string(body), userID)
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

	var userID int
	cookieValue, err := cookie.GetEncryptedValue(r, "user_id", []byte(h.config.SecretKey))
	if err != nil {
		userID, err = h.service.GetNewUser()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = cookie.WriteEncryptedValue(w, "user_id", strconv.Itoa(userID), []byte(h.config.SecretKey))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		userID, err = strconv.Atoi(cookieValue)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	uri, err := h.service.ReturnURI(requestBody.URL, userID)
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

	w.Header().Set("Content-Type", "application/json")
	log.Println("ReturnJSONShortURL:", requestBody.URL, "->", shortURL)
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

func (h handler) GetUserURL(w http.ResponseWriter, r *http.Request) {
	type userURL struct {
		ShortURL    string `json:"short_url"`
		OriginalURL string `json:"original_url"`
	}

	cookieValue, err := cookie.GetEncryptedValue(r, "user_id", []byte(h.config.SecretKey))
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "[]", http.StatusNoContent)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	userID, err := strconv.Atoi(cookieValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	urls, err := h.service.GetURLByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(urls) == 0 {
		http.Error(w, "[]", http.StatusNoContent)
		return
	}
	responceBody := make([]userURL, 0)
	for uri, originalURL := range urls {
		shortURL, err := url.JoinPath(h.config.BaseURL, uri)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responceBody = append(responceBody, userURL{ShortURL: shortURL, OriginalURL: originalURL})
	}

	responce, err := json.Marshal(responceBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("GetUserURL:", cookieValue)
	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, responce, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, text []byte, code int) {
	w.WriteHeader(code)
	w.Write(text)
}
