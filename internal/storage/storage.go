// Package storage provides functionality to store and receive URL and user data
package storage

import (
	"context"
	"errors"
)

// Errors
var (
	ErrURLNotFound       = errors.New("URL not found")
	ErrWriteDataConflict = errors.New("write data conflict")
	ErrURLIsDeleted      = errors.New("URL is deleted")
)

// URLStorage describes interface to store and receive URL and user data
type URLStorage interface {
	WriteData(context.Context, string, int) (int, error)
	WriteBatchData(context.Context, []string, int) ([]int, error)
	GetData(context.Context, int) (string, error)
	CheckUser(context.Context, int) (int, error)
	GetUserURL(context.Context, int) ([]URLData, error)
	DeleteBatchData(context.Context, []int, int)
	GetStats(context.Context) (int, int, error)
}

// URLStorage saves URL data
type URLData struct {
	ID          int    `json:"-"`
	URI         string `json:"-"`
	OriginalURL string `json:"original_url,omitempty"`
	ShortURL    string `json:"short_url,omitempty"`
	CorrID      string `json:"correlation_id,omitempty"`
}
