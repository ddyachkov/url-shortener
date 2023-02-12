package main

import (
	"flag"

	"github.com/ddyachkov/url-shortener/internal/server"
)

func main() {
	flag.Parse()
	server.Run()
}
