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
	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
)

type handler struct {
	service *app.URLShortener
	config  *config.ServerConfig
	db      *pgxpool.Pool
}

func NewURLHandler(shortener *app.URLShortener, cfg *config.ServerConfig, dbpool *pgxpool.Pool) http.Handler {
	router := chi.NewRouter()

	h := handler{
		service: shortener,
		config:  cfg,
		db:      dbpool,
	}

	router.Use(middleware.Decompress, middleware.Compress)

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
	gotUserID, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if userID != gotUserID {
		err = cookie.WriteEncryptedValue(w, "user_id", strconv.Itoa(gotUserID), []byte(h.config.SecretKey))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

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
	gotUserID, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if userID != gotUserID {
		err = cookie.WriteEncryptedValue(w, "user_id", strconv.Itoa(gotUserID), []byte(h.config.SecretKey))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

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
	gotUserID, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if userID != gotUserID {
		err = cookie.WriteEncryptedValue(w, "user_id", strconv.Itoa(gotUserID), []byte(h.config.SecretKey))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

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

func (h handler) RedirectToFullURL(w http.ResponseWriter, r *http.Request) {
	uri := chi.URLParam(r, "URI")
	fullURL, err := h.service.GetFullURL(r.Context(), uri)
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
