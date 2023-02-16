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
		name    string
		args    args
		wantID  int
		wantErr bool
	}{
		{
			name:    "Positive_NewData",
			args:    args{url: "https://www.google.ru"},
			wantID:  1,
			wantErr: false,
		},
		{
			name:    "Positive_SameData",
			args:    args{url: "https://www.google.ru"},
			wantID:  1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := storage.WriteData(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLStorage.GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantID, gotID)
		})
	}
}

func TestURLStorage_GetData(t *testing.T) {
	storage := NewURLStorage()
	url := "https://www.google.ru"
	gotID, err := storage.WriteData(url)
	if err != nil {
		t.Fatal(err)
	}
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
			assert.Equal(t, tt.wantURL, gotURL)
		})
	}
}
