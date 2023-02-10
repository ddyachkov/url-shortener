package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestURLHandler_ServeHTTP(t *testing.T) {
	storage := storage.NewURLStorage()
	service := app.NewURLShortener(&storage)
	handler := NewURLHandler(&service)
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
			header: header{
				contentType: "text/plain",
			},
			body: "https://www.google.ru",
			want: want{
				code: http.StatusCreated,
				text: "http://localhost:8080/b",
			},
		},
		{
			name:   "Negative_POST_Text_Code415",
			method: http.MethodPost,
			path:   "/",
			header: header{
				contentType: "application/json",
			},
			body: "https://www.google.ru",
			want: want{
				code: http.StatusUnsupportedMediaType,
				text: "Unsupported Media Type\n",
			},
		},
		{
			name:   "Negative_POST_Text_Code400",
			method: http.MethodPost,
			path:   "/",
			header: header{
				contentType: "text/plain",
			},
			body: "www.google.ru",
			want: want{
				code: http.StatusBadRequest,
				text: "parse \"www.google.ru\": invalid URI for request\n",
			},
		},
		{
			name:   "Positive_POST_JSON_Code201",
			method: http.MethodPost,
			path:   "/api/shorten",
			header: header{
				contentType: "application/json",
			},
			body: "{\"URL\":\"https://www.google.ru\"}",
			want: want{
				code: http.StatusCreated,
				text: "{\"result\":\"http://localhost:8080/b\"}",
			},
		},
		{
			name:   "Negative_POST_JSON_Code415",
			method: http.MethodPost,
			path:   "/api/shorten",
			header: header{
				contentType: "text/plain",
			},
			body: "{\"URL\":\"https://www.google.ru\"}",
			want: want{
				code: http.StatusUnsupportedMediaType,
				text: "Unsupported Media Type\n",
			},
		},
		{
			name:   "Negative_POST_JSON_Code400",
			method: http.MethodPost,
			path:   "/api/shorten",
			header: header{
				contentType: "application/json",
			},
			body: "{\"URL\":\"www.google.ru\"}",
			want: want{
				code: http.StatusBadRequest,
				text: "parse \"www.google.ru\": invalid URI for request\n",
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
					location: "https://www.google.ru",
				},
			},
		},
		{
			name:   "Negative_GET_Code400",
			method: http.MethodGet,
			path:   "/c",
			want: want{
				code: http.StatusNotFound,
				text: "URL not found\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			bodyReader := strings.NewReader(tt.body)
			r := httptest.NewRequest(tt.method, tt.path, bodyReader)
			r.Header.Set("Content-Type", tt.header.contentType)
			handler.ServeHTTP(w, r)
			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)
			assert.Equal(t, tt.want.header.location, res.Header.Get("Location"))

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			assert.Equal(t, tt.want.text, string(resBody))
		})
	}
}
