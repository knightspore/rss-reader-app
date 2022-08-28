package main

import (
	"log"

	"github.com/knightspore/rss-reader-app/backend/server"
)

func main() {
	s := server.NewServer()
	log.Println("> Server starting on 127.0.0.1:1337")
	s.Start()
}
