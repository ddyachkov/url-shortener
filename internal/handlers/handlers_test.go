package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestURLPostHandler_ServeHTTP(t *testing.T) {
	storage := storage.NewURLStorage()
	handler := NewURLPostHandler(&storage)
	type want struct {
		code           int
		text           string
		headerLocation string
	}
	tests := []struct {
		name   string
		method string
		path   string
		body   string
		want   want
	}{
		{
			name:   "Positive_POST_Code201",
			method: http.MethodPost,
			path:   "/",
			body:   "https://www.google.ru",
			want: want{
				code: http.StatusCreated,
				text: "http://localhost:8080/b",
			},
		},
		{
			name:   "Negative_POST_Code400",
			method: http.MethodPost,
			path:   "/",
			body:   "www.google.ru",
			want: want{
				code: http.StatusBadRequest,
				text: "parse \"www.google.ru\": invalid URI for request",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			bodyReader := strings.NewReader(tt.body)
			r := httptest.NewRequest(tt.method, tt.path, bodyReader)
			handler.ServeHTTP(w, r)
			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)
			assert.Equal(t, tt.want.headerLocation, res.Header.Get("Location"))

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want.text, string(resBody))
		})
	}
}

func TestURLGetHandler_ServeHTTP(t *testing.T) {
	storage := storage.NewURLStorage()
	t.Run("Positive_POST_Code201", func(t *testing.T) {
		w := httptest.NewRecorder()
		bodyReader := strings.NewReader("https://www.google.ru")
		r := httptest.NewRequest(http.MethodPost, "/", bodyReader)
		postHandler := NewURLPostHandler(&storage)
		postHandler.ServeHTTP(w, r)
		res := w.Result()

		assert.Equal(t, http.StatusCreated, res.StatusCode)

		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "http://localhost:8080/b", string(resBody))
	})
	handler := NewURLGetHandler(&storage)
	type want struct {
		code           int
		text           string
		headerLocation string
	}
	tests := []struct {
		name   string
		method string
		path   string
		body   string
		want   want
	}{
		{
			name:   "Positive_GET_Code307",
			method: http.MethodGet,
			path:   "/b",
			want: want{
				code:           http.StatusTemporaryRedirect,
				text:           "<a href=\"https://www.google.ru\">Temporary Redirect</a>.\n\n",
				headerLocation: "https://www.google.ru",
			},
		},
		{
			name:   "Negative_GET_Code400",
			method: http.MethodGet,
			path:   "/c",
			want: want{
				code: http.StatusBadRequest,
				text: "URL not found\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			bodyReader := strings.NewReader(tt.body)
			r := httptest.NewRequest(tt.method, tt.path, bodyReader)
			handler.ServeHTTP(w, r)
			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)
			assert.Equal(t, tt.want.headerLocation, res.Header.Get("Location"))

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.want.text, string(resBody))
		})
	}
}
