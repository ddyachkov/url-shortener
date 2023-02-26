package storage

type URLStorage interface {
	WriteData(string, int) (int, error)
	GetData(int) (string, error)
	CheckUser(int) (bool, error)
	MakeNewUser() (int, error)
	GetUserURL(int) ([]URLData, error)
}

type URLData struct {
	ID  int
	URI string
	URL string
}
