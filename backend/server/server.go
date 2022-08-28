package server

import (
	"log"
	"net/http"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/gorilla/mux"
)

type Server struct {
	DB     *gocb.Bucket
	Router *mux.Router
}

type Handler func(http.ResponseWriter, *http.Request)

func NewServer() *Server {
	s := &Server{}
	s.Router = mux.NewRouter()
	err := s.ConnectDatabase()
	if err != nil {
		log.Panic(err)
	}
	return s
}

func (s *Server) Start() {

	s.Routes()

	s.Router.Use(Middleware)

	srv := &http.Server{
		Handler:      s.Router,
		Addr:         "127.0.0.1:1337",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()

}