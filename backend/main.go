package main

import (
	"fmt"

	"github.com/knightspore/rss-reader-app/backend/server"
)

func main() {

	s := server.NewServer()

	fmt.Println("> Server starting on :1337")

	s.Start()

}
