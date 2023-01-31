package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLStorage_WriteData(t *testing.T) {
	storage := NewURLStorage()
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		args   args
		wantID int
	}{
		{
			name:   "Positive_NewData",
			args:   args{url: "https://www.google.ru"},
			wantID: 1,
		},
		{
			name:   "Positive_SameData",
			args:   args{url: "https://www.google.ru"},
			wantID: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID := storage.WriteData(tt.args.url)
			assert.Equal(t, gotID, tt.wantID)
		})
	}
}

func TestURLStorage_GetData(t *testing.T) {
	storage := NewURLStorage()
	url := "https://www.google.ru"
	gotID := storage.WriteData(url)
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantURL string
		wantErr bool
	}{
		{
			name:    "Positive_FoundData",
			args:    args{id: gotID},
			wantURL: url,
			wantErr: false,
		},
		{
			name:    "Negative_NotFoundData",
			args:    args{id: 2},
			wantURL: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, err := storage.GetData(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLStorage.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotURL, tt.wantURL)
		})
	}
}
