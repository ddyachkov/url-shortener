package app

import (
	"fmt"
	"os"
	"testing"

	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cfg config.ServerConfig = config.ServerConfig{FileStoragePath: "/tmp/data.txt", SecretKey: "thisisthirtytwobytelongsecretkey"}

func createFile(t *testing.T, fileStoragePath string, content string) {
	file, err := os.Create(fileStoragePath)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = os.Remove(file.Name())
	})
	t.Cleanup(func() {
		_ = file.Close()
	})
	_, err = file.Write([]byte(content))
	require.NoError(t, err)
}

func Test_makeURI(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantURI string
	}{
		{
			name:    "Positive_ValidURI",
			args:    args{id: 1},
			wantURI: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURI := makeURI(tt.args.id)
			assert.Equal(t, gotURI, tt.wantURI)
		})
	}
}

func Test_makeID(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name   string
		args   args
		wantID int
	}{
		{
			name:   "Positive_ValidID",
			args:   args{uri: "b"},
			wantID: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID := makeID(tt.args.uri)
			assert.Equal(t, gotID, tt.wantID)
		})
	}
}

func TestURLShortener_ReturnURI(t *testing.T) {
	storage := storage.NewURLStorage()
	shortener := NewURLShortener(&storage)
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantURI string
		wantErr bool
	}{
		{
			name:    "Positive_NewURL",
			args:    args{url: "https://www.google.ru"},
			wantURI: "b",
			wantErr: false,
		},
		{
			name:    "Positive_SameURL",
			args:    args{url: "https://www.google.ru"},
			wantURI: "b",
			wantErr: false,
		},
		{
			name:    "Negative_InvalidURL",
			args:    args{url: "www.google.ru"},
			wantURI: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURI, err := shortener.ReturnURI(tt.args.url, 1)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLShortener.ReturnURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotURI, tt.wantURI)
		})
	}
}

func TestURLShortener_GetFullURL(t *testing.T) {
	storage := storage.NewURLStorage()
	shortener := NewURLShortener(&storage)
	url := "https://www.google.ru"
	gotURI, err := shortener.ReturnURI(url, 1)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		wantURL string
		wantErr bool
	}{
		{
			name:    "Positive_URLFound",
			args:    args{uri: gotURI},
			wantURL: url,
			wantErr: false,
		},
		{
			name:    "Negative_URLNotFound",
			args:    args{uri: "a"},
			wantURL: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, err := shortener.GetFullURL(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLShortener.GetFullURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotURL, tt.wantURL)
		})
	}
}

func TestURLShortener_GetFullURL_WithFileStorage(t *testing.T) {
	url := "https://www.google.ru"
	dataID := 1
	userID := 1
	createFile(t, cfg.FileStoragePath, fmt.Sprint(dataID, " ", url, " ", userID))
	storage := storage.NewURLFileStorage(&cfg)
	storage.LoadData()
	shortener := NewURLShortener(&storage)

	type args struct {
		uri string
	}
	tests := []struct {
		name    string
		args    args
		wantURL string
		wantErr bool
	}{
		{
			name:    "Positive_URLFound",
			args:    args{uri: makeURI(dataID)},
			wantURL: url,
			wantErr: false,
		},
		{
			name:    "Negative_URLNotFound",
			args:    args{uri: "a"},
			wantURL: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, err := shortener.GetFullURL(tt.args.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLShortener.GetFullURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotURL, tt.wantURL)
		})
	}
}
