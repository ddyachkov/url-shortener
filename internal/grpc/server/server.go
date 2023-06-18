package server

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ddyachkov/url-shortener/internal/app"
	"github.com/ddyachkov/url-shortener/internal/config"
	"github.com/ddyachkov/url-shortener/internal/grpc/proto"
	"github.com/ddyachkov/url-shortener/internal/storage"
	"google.golang.org/grpc/metadata"
)

// ShortenerServer is a struct for gRPC server
type ShortenerServer struct {
	proto.UnimplementedShortenerServer
	Service *app.URLShortener
	Config  *config.ServerConfig
}

// CreateShortURL returns response with short URL for request with URL.
func (s *ShortenerServer) CreateShortURL(ctx context.Context, in *proto.CreateShortURLRequest) (*proto.CreateShortURLResponce, error) {
	var responce proto.CreateShortURLResponce
	var userID int

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("userID")
		if len(values) > 0 {
			userID, _ = strconv.Atoi(values[0])
		}
	}

	uri, err := s.Service.ReturnURI(ctx, in.Url, userID)
	if err != nil {
		if errors.Is(err, storage.ErrWriteDataConflict) {
			responce.Error = fmt.Sprintf("URL %s уже сокращен", in.Url)
		} else {
			responce.Error = fmt.Sprint(err)
			return &responce, nil
		}
	}

	shortURL, err := url.JoinPath(s.Config.BaseURL, uri)
	if err != nil {
		responce.Error = fmt.Sprint(err)
		return &responce, nil
	}

	responce.ShortUrl = shortURL
	return &responce, nil
}
