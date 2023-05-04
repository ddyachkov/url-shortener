package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ddyachkov/url-shortener/internal/cookie"
)

// GetEncryptedUserID gets user id from encrypted cookie and passes it further to handler.
func (h handler) GetEncryptedUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userID int
		cookieValue, err := cookie.GetEncryptedValue(r, "user_id", []byte(h.config.SecretKey))
		if err != nil {
			userID = 0
		}

		userID, err = strconv.Atoi(cookieValue)
		if err != nil {
			userID = 0
		}

		ctx := context.WithValue(r.Context(), contextUserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// SetEncryptedUserID sets encrypted cookie with user id to return to client.
func (h handler) SetEncryptedUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(contextUserIDKey).(int)
		gotUserID, err := h.service.GetUser(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if userID != gotUserID {
			r = r.WithContext(context.WithValue(r.Context(), contextUserIDKey, gotUserID))
			err = cookie.WriteEncryptedValue(w, "user_id", strconv.Itoa(gotUserID), []byte(h.config.SecretKey))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
