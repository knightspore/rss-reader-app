package module

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

func (s *Server) Routes() {
	// Users
	s.Router.HandleFunc("/api/user/create", s.UserCreate).Methods(http.MethodGet, http.MethodPost)
	// Subscriptions
	s.Router.HandleFunc("/api/subscription/create", s.SubscriptionCreate)
	s.Router.HandleFunc("/api/subscription/get", s.SubscriptionGet)
	// Reading List
	s.Router.HandleFunc("/api/readinglist/get", s.ReadingListGet)
}

func (s *Server) Start() {

	s.Routes()
	s.Router.Headers("Content-Type", "application/json")

	srv := &http.Server{
		Handler:      s.Router,
		Addr:         "127.0.0.1:1337",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()

}

func (s *Server) ConnectDatabase() error {	

	endpoint := "cb.63kdb50zq6bwphve.cloud.couchbase.com"
	bucketName := "rss-reader"
	username := "go-app"
	password := "Barcka1332!"

	cluster, err := gocb.Connect("couchbases://"+endpoint, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		return err
	}

	bucket := cluster.Bucket(bucketName)

	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}

	s.DB = bucket

	return err

}