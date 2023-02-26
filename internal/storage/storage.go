package storage

type URLStorage interface {
	WriteData(string, int) (int, error)
	GetData(int) (string, error)
	CheckUser(int) (bool, error)
	MakeNewUser() (int, error)
	GetUserURL(int) ([]URLData, error)
}

type URLData struct {
	ID          int    `json:"-"`
	URI         string `json:"-"`
	OriginalURL string `json:"original_url,omitempty"`
	ShortURL    string `json:"short_url,omitempty"`
	CorrID      string `json:"correlation_id,omitempty"`
}
