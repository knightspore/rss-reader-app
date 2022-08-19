package main

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/knightspore/rss-reader-app/backend/module"
	"github.com/knightspore/rss-reader-app/backend/vo"
	_ "github.com/mattn/go-sqlite3"
)

const file string = "/tmp/test.db"

func main() {
	
	var Users []vo.User

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := &module.Server{
		DB:     db,
		Router: mux.NewRouter(),
		Users: Users,
	}

	fmt.Println("> Server starting on :1337")

	s.Start()

}
