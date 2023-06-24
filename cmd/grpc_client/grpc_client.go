package main

import (
	"context"
	"log"

	"github.com/ddyachkov/url-shortener/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewShortenerClient(conn)
	TestShortener(c)
}

// TestShortener is a function to test gRPC server
func TestShortener(c proto.ShortenerClient) {
	md := metadata.New(map[string]string{"userID": "1"})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	req := &proto.CreateShortURLRequest{Url: "https://www.google.ru"}
	resp, err := c.CreateShortURL(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.ShortUrl, resp.Error)
}
