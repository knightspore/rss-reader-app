package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/knightspore/rss-reader-app/backend/module"
)

func main() {
	
	s := &module.Server{
		Router: mux.NewRouter(),
	}

	err := s.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("> Server starting on :1337")

	s.Start()

}
