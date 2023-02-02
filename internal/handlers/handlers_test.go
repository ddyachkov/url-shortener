package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLHandler_ServeHTTP(t *testing.T) {
	handler := NewURLHandler()
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
			body:   "htts://www.google.ru",
			want: want{
				code: http.StatusBadRequest,
				text: "URL is invalid",
			},
		},
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
