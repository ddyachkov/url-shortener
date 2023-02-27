package storage

import "context"

type URLStorage interface {
	WriteData(context.Context, string, int) (int, error)
	WriteBatchData(context.Context, []URLData, int) error
	GetData(context.Context, int) (string, error)
	CheckUser(context.Context, int) (bool, error)
	MakeNewUser(context.Context) (int, error)
	GetUserURL(context.Context, int) ([]URLData, error)
}

type URLData struct {
	ID          int    `json:"-"`
	URI         string `json:"-"`
	OriginalURL string `json:"original_url,omitempty"`
	ShortURL    string `json:"short_url,omitempty"`
	CorrID      string `json:"correlation_id,omitempty"`
}
