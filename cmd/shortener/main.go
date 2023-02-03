package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ddyachkov/url-shortener/internal/handlers"
	"github.com/ddyachkov/url-shortener/internal/storage"
	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	storage := storage.NewURLStorage()
	postHandler := handlers.NewURLPostHandler(&storage)
	getHandler := handlers.NewURLGetHandler(&storage)
	router.Method(http.MethodPost, "/", postHandler)
	router.Method(http.MethodGet, "/{URI}", getHandler)
	router.Get("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is shutting down..."))
		cancel()
	})
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	server.Shutdown(ctx)
	log.Printf("Finished")
}
