package main

import (
	"fmt"

	"github.com/knightspore/rss-reader-app/backend/module"
)

func main() {

	s := module.NewServer()

	fmt.Println("> Server starting on :1337")

	s.Start()

}
