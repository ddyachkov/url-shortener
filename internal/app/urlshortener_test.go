package app

import (
	"context"
	"math/rand"
	"testing"

	"github.com/ddyachkov/url-shortener/internal/random"
	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestURLShortener_ReturnURI(t *testing.T) {
	storage := storage.NewURLMemStorage()
	service := NewURLShortener(storage)
	url := random.URL().String()
	tests := []struct {
		name    string
		url     string
		wantURI bool
		wantErr bool
	}{
		{
			name:    "Positive_NewURL",
			url:     url,
			wantURI: true,
			wantErr: false,
		},
		{
			name:    "Positive_SameURL",
			url:     url,
			wantURI: true,
			wantErr: true,
		},
		{
			name:    "Negative_InvalidURL",
			url:     random.ASCIIString(5, 15),
			wantURI: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURI, err := service.ReturnURI(context.Background(), tt.url, 1)
			assert.Equal(t, len(gotURI) > 0, tt.wantURI)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestURLShortener_ReturnBatchURI(t *testing.T) {
	storage := storage.NewURLMemStorage()
	service := NewURLShortener(storage)
	correctURLs := make([]string, 0)
	for i := 0; i < 2; i++ {
		correctURLs = append(correctURLs, random.URL().String())
	}
	tests := []struct {
		name    string
		urls    []string
		wantURI bool
		wantErr bool
	}{
		{
			name:    "Positive_NewURLs",
			urls:    correctURLs,
			wantURI: true,
			wantErr: false,
		},
		{
			name:    "Negative_InvalidURLs",
			urls:    []string{random.ASCIIString(5, 15)},
			wantURI: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURI, err := service.ReturnBatchURI(context.Background(), tt.urls, 1)
			assert.Equal(t, len(gotURI) > 0, tt.wantURI)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestURLShortener_GetFullURL(t *testing.T) {
	storage := storage.NewURLMemStorage()
	service := NewURLShortener(storage)

	url := random.URL().String()
	uri, err := service.ReturnURI(context.Background(), url, 1)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		uri     string
		url     string
		wantErr bool
	}{
		{
			name:    "Positive_URLFound",
			uri:     uri,
			url:     url,
			wantErr: false,
		},
		{
			name:    "Negative_URLNotFound",
			uri:     random.ASCIIString(5, 15),
			url:     "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, err := service.GetFullURL(context.Background(), tt.uri)
			assert.Equal(t, gotURL, tt.url)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestURLShortener_GetUser(t *testing.T) {
	storage := storage.NewURLMemStorage()
	service := NewURLShortener(storage)

	userID := rand.Intn(100)
	tests := []struct {
		name    string
		newUser bool
	}{
		{
			name:    "Positive_NewUser",
			newUser: true,
		},
		{
			name:    "Positive_SameUser",
			newUser: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, _ := service.GetUser(context.Background(), userID)
			assert.Equal(t, gotID != userID, tt.newUser)
			if tt.newUser {
				userID = gotID
			}
		})
	}
}

func TestURLShortener_GetURLByUser(t *testing.T) {
	storage := storage.NewURLMemStorage()
	service := NewURLShortener(storage)

	userID, err := service.GetUser(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}

	urls := make([]string, 0)
	for i := 0; i < 2; i++ {
		urls = append(urls, random.URL().String())
	}
	_, err = service.ReturnBatchURI(context.Background(), urls, userID)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		userID   int
		wantData bool
	}{
		{
			name:     "Positive_FoundURLs",
			userID:   userID,
			wantData: true,
		},
		{
			name:     "Negative_InvalidURLs",
			userID:   userID + 1,
			wantData: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, _ := service.GetURLByUser(context.Background(), tt.userID)
			assert.Equal(t, len(gotData) > 0, tt.wantData)
		})
	}
}

func Test_makeURI(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		wantURI string
	}{
		{
			name:    "Positive_ValidURI",
			id:      1,
			wantURI: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURI := makeURI(tt.id)
			assert.Equal(t, gotURI, tt.wantURI)
		})
	}
}

func Test_makeID(t *testing.T) {
	tests := []struct {
		name   string
		uri    string
		wantID int
	}{
		{
			name:   "Positive_ValidID",
			uri:    "b",
			wantID: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID := makeID(tt.uri)
			assert.Equal(t, gotID, tt.wantID)
		})
	}
}
