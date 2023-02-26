package handler

import (
	"context"
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
	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx"
)

type handler struct {
	service *app.URLShortener
	config  *config.ServerConfig
	conn    *pgx.Conn
}

func NewURLHandler(shortener *app.URLShortener, cfg *config.ServerConfig, c *pgx.Conn) http.Handler {
	router := chi.NewRouter()

	h := handler{
		service: shortener,
		config:  cfg,
		conn:    c,
	}

	router.Use(middleware.Decompress)
	router.Use(middleware.Compress)

	router.Post("/", h.ReturnTextShortURL)
	router.Post("/api/shorten", h.ReturnJSONShortURL)
	router.Post("/api/shorten/batch", h.ReturnBatchJSONShortURL)

	router.Get("/{URI}", h.RedirectToFullURL)
	router.Get("/api/user/urls", h.GetUserURL)
	router.Get("/ping", h.PingDatabase)

	return router
}

func (h handler) ReturnTextShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := getUserID(r, []byte(h.config.SecretKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newUser, err := h.service.GetUser(&userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if newUser {
		err = cookie.WriteEncryptedValue(w, "user_id", strconv.Itoa(userID), []byte(h.config.SecretKey))
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

	userID, err := getUserID(r, []byte(h.config.SecretKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newUser, err := h.service.GetUser(&userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if newUser {
		err = cookie.WriteEncryptedValue(w, "user_id", strconv.Itoa(userID), []byte(h.config.SecretKey))
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

	userID, err := getUserID(r, []byte(h.config.SecretKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newUser, err := h.service.GetUser(&userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if newUser {
		err = cookie.WriteEncryptedValue(w, "user_id", strconv.Itoa(userID), []byte(h.config.SecretKey))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err = h.service.ReturnBatchURI(batchData, userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range batchData {
		shortURL, err := url.JoinPath(h.config.BaseURL, batchData[i].URI)
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
	userID, err := getUserID(r, []byte(h.config.SecretKey))
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "[]", http.StatusNoContent)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	responceBody, err := h.service.GetURLByUser(userID)
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

func (h handler) PingDatabase(w http.ResponseWriter, r *http.Request) {
	if h.conn == nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err := h.conn.Ping(context.Background())
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

func getUserID(r *http.Request, secretKey []byte) (userID int, err error) {
	cookieValue, err := cookie.GetEncryptedValue(r, "user_id", []byte(secretKey))
	if err != nil {
		return 0, nil
	}

	userID, err = strconv.Atoi(cookieValue)
	if err != nil {
		return 0, nil
	}

	return userID, nil
}
