package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cfg = config.ServerConfig{
	ServerAddress:   "localhost:8080",
	BaseURL:         "http://localhost:8080",
	FileStoragePath: "/tmp/data.txt",
	SecretKey:       "thisisthirtytwobytelongsecretkey",
}

func TestURLHandler_ServeHTTP(t *testing.T) {
	storage := storage.NewURLMemStorage()
	service := app.NewURLShortener(storage)
	handler := NewURLHandler(service, &cfg, nil)
	type header struct {
		contentType string
		location    string
	}
	type want struct {
		code   int
		text   string
		header header
	}
	tests := []struct {
		name   string
		method string
		path   string
		header header
		body   string
		want   want
	}{
		{
			name:   "Positive_POST_Text_Code201",
			method: http.MethodPost,
			path:   "/",
			body:   "https://www.google.ru",
			want: want{
				code: http.StatusCreated,
				text: cfg.BaseURL + "/b",
			},
		},
		{
			name:   "Negative_POST_Text_Code409",
			method: http.MethodPost,
			path:   "/",
			body:   "https://www.google.ru",
			want: want{
				code: http.StatusConflict,
				text: cfg.BaseURL + "/b",
			},
		},
		{
			name:   "Negative_POST_Text_Code400",
			method: http.MethodPost,
			path:   "/",
			body:   "www.google.ru",
			want: want{
				code: http.StatusBadRequest,
				text: "parse \"www.google.ru\": invalid URI for request\n",
				header: header{
					contentType: "text/plain; charset=utf-8",
				},
			},
		},
		{
			name:   "Positive_POST_JSON_Code201",
			method: http.MethodPost,
			path:   "/api/shorten",
			body:   "{\"URL\":\"https://www.google.com\"}",
			want: want{
				code: http.StatusCreated,
				text: "{\"result\":\"" + cfg.BaseURL + "/c\"}",
				header: header{
					contentType: "application/json",
				},
			},
		},
		{
			name:   "Negative_POST_JSON_Code409",
			method: http.MethodPost,
			path:   "/api/shorten",
			body:   "{\"URL\":\"https://www.google.ru\"}",
			want: want{
				code: http.StatusConflict,
				text: "{\"result\":\"" + cfg.BaseURL + "/b\"}",
				header: header{
					contentType: "application/json",
				},
			},
		},
		{
			name:   "Negative_POST_JSON_Code400",
			method: http.MethodPost,
			path:   "/api/shorten",
			body:   "{\"URL\":\"www.google.ru\"}",
			want: want{
				code: http.StatusBadRequest,
				text: "parse \"www.google.ru\": invalid URI for request\n",
				header: header{
					contentType: "text/plain; charset=utf-8",
				},
			},
		},
		{
			name:   "Positive_GET_Code307",
			method: http.MethodGet,
			path:   "/b",
			want: want{
				code: http.StatusTemporaryRedirect,
				text: "<a href=\"https://www.google.ru\">Temporary Redirect</a>.\n\n",
				header: header{
					contentType: "text/html; charset=utf-8",
					location:    "https://www.google.ru",
				},
			},
		},
		{
			name:   "Negative_GET_Code400",
			method: http.MethodGet,
			path:   "/d",
			want: want{
				code: http.StatusNotFound,
				text: "URL not found\n",
				header: header{
					contentType: "text/plain; charset=utf-8",
				},
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
			assert.Equal(t, tt.want.header.contentType, res.Header.Get("Content-Type"))
			assert.Equal(t, tt.want.header.location, res.Header.Get("Location"))

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Equal(t, tt.want.text, string(resBody))
		})
	}
}

func TestURLHandler_ServeHTTP_WithFileStorage(t *testing.T) {
	storage := storage.NewURLFileStorage(&cfg)

	t.Cleanup(func() {
		_ = os.Remove(cfg.FileStoragePath)
	})
	service := app.NewURLShortener(storage)
	handler := NewURLHandler(service, &cfg, nil)
	type header struct {
		contentType string
		location    string
	}
	type want struct {
		code   int
		text   string
		header header
	}
	tests := []struct {
		name   string
		method string
		path   string
		header header
		body   string
		want   want
	}{
		{
			name:   "Positive_POST_Text_Code201",
			method: http.MethodPost,
			path:   "/",
			body:   "https://www.google.ru",
			want: want{
				code: http.StatusCreated,
				text: cfg.BaseURL + "/b",
			},
		},
		{
			name:   "Negative_POST_Text_Code409",
			method: http.MethodPost,
			path:   "/",
			body:   "https://www.google.ru",
			want: want{
				code: http.StatusConflict,
				text: cfg.BaseURL + "/b",
			},
		},
		{
			name:   "Negative_POST_Text_Code400",
			method: http.MethodPost,
			path:   "/",
			body:   "www.google.ru",
			want: want{
				code: http.StatusBadRequest,
				text: "parse \"www.google.ru\": invalid URI for request\n",
				header: header{
					contentType: "text/plain; charset=utf-8",
				},
			},
		},
		{
			name:   "Positive_POST_JSON_Code201",
			method: http.MethodPost,
			path:   "/api/shorten",
			body:   "{\"URL\":\"https://www.google.com\"}",
			want: want{
				code: http.StatusCreated,
				text: "{\"result\":\"" + cfg.BaseURL + "/c\"}",
				header: header{
					contentType: "application/json",
				},
			},
		},
		{
			name:   "Negative_POST_JSON_Code409",
			method: http.MethodPost,
			path:   "/api/shorten",
			body:   "{\"URL\":\"https://www.google.ru\"}",
			want: want{
				code: http.StatusConflict,
				text: "{\"result\":\"" + cfg.BaseURL + "/b\"}",
				header: header{
					contentType: "application/json",
				},
			},
		},
		{
			name:   "Negative_POST_JSON_Code400",
			method: http.MethodPost,
			path:   "/api/shorten",
			body:   "{\"URL\":\"www.google.ru\"}",
			want: want{
				code: http.StatusBadRequest,
				text: "parse \"www.google.ru\": invalid URI for request\n",
				header: header{
					contentType: "text/plain; charset=utf-8",
				},
			},
		},
		{
			name:   "Positive_GET_Code307",
			method: http.MethodGet,
			path:   "/b",
			want: want{
				code: http.StatusTemporaryRedirect,
				text: "<a href=\"https://www.google.ru\">Temporary Redirect</a>.\n\n",
				header: header{
					contentType: "text/html; charset=utf-8",
					location:    "https://www.google.ru",
				},
			},
		},
		{
			name:   "Negative_GET_Code400",
			method: http.MethodGet,
			path:   "/d",
			want: want{
				code: http.StatusNotFound,
				text: "URL not found\n",
				header: header{
					contentType: "text/plain; charset=utf-8",
				},
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
			assert.Equal(t, tt.want.header.contentType, res.Header.Get("Content-Type"))
			assert.Equal(t, tt.want.header.location, res.Header.Get("Location"))

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Equal(t, tt.want.text, string(resBody))
		})
	}
}
