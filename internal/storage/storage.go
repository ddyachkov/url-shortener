package storage

import (
	"context"
	"errors"
)

var ErrURLNotFound = errors.New("URL not found")
var ErrWriteDataConflict = errors.New("write data conflict")
var ErrURLIsDeleted = errors.New("URL is deleted")

type URLStorage interface {
	WriteData(context.Context, string, int) (int, error)
	WriteBatchData(context.Context, []string, int) ([]int, error)
	GetData(context.Context, int) (string, error)
	CheckUser(context.Context, int) (int, error)
	MakeNewUser(context.Context) (int, error)
	GetUserURL(context.Context, int) ([]URLData, error)
	DeleteBatchData(context.Context, []int, int)
}

type URLData struct {
	ID          int    `json:"-"`
	URI         string `json:"-"`
	OriginalURL string `json:"original_url,omitempty"`
	ShortURL    string `json:"short_url,omitempty"`
	CorrID      string `json:"correlation_id,omitempty"`
}
