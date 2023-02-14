package handler

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type handler struct {
	service *app.URLShortener
}

func NewURLHandler(shortener *app.URLShortener) http.Handler {
	router := chi.NewRouter()

	h := handler{
		service: shortener,
	}

	router.Use(middleware.Compress(5))
	router.Post("/", h.ReturnTextShortURL)
	router.Post("/api/shorten", h.ReturnJSONShortURL)

	router.Get("/{URI}", h.RedirectToFullURL)

	return router
}

func (h handler) ReturnTextShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := returnRequestBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uri, err := h.service.ReturnURI(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("ReturnTextShortURL:", string(body), "->", config.BaseURL+"/"+uri)
	writeResponse(w, []byte(config.BaseURL+"/"+uri), http.StatusCreated)
}

func (h handler) ReturnJSONShortURL(w http.ResponseWriter, r *http.Request) {
	body, err := returnRequestBody(r)
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

	responceBody := struct {
		Result string `json:"result"`
	}{Result: config.BaseURL + "/" + uri}
	responce, err := json.Marshal(responceBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("ReturnJSONShortURL:", requestBody.URL, "->", config.BaseURL+"/"+uri)
	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, responce, http.StatusCreated)
}

func (h handler) RedirectToFullURL(w http.ResponseWriter, r *http.Request) {
	uri := chi.URLParam(r, "URI")
	url, err := h.service.GetFullURL(uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Println("RedirectToFullURL:", config.BaseURL+"/"+uri, "->", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func returnRequestBody(r *http.Request) (body []byte, err error) {
	var reader io.Reader

	if r.Header.Get(`Content-Encoding`) == `gzip` {
		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			return nil, err
		}
		reader = gz
		defer gz.Close()
	} else {
		reader = r.Body
	}

	body, err = io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func writeResponse(w http.ResponseWriter, text []byte, code int) {
	w.WriteHeader(code)
	w.Write(text)
}
