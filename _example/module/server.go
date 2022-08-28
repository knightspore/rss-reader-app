package module

import (
	"fmt"
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
	s.Router.HandleFunc("/api/user/create", s.HandleUserCreate())
	s.Router.HandleFunc("/api/subscription/create", s.HandleSubscriptionCreate())
	s.Router.HandleFunc("/api/subscription/get", s.HandleSubscriptionGet())
	s.Router.HandleFunc("/api/readinglist/get", s.HandleReadingListGet())
	s.Router.HandleFunc("/api/article/read", s.HandleArticleRead())
}

func NewServer() *Server {

	s := &Server{}
	s.Router = mux.NewRouter()

	err := s.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
	}

	return s

}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)

		defer func(startedAt time.Time) {
			log.Println(r.RequestURI, time.Since(startedAt))
		}(time.Now())

		next.ServeHTTP(w, r)
	})
}

func (s *Server) Start() {

	port := "1337"

	s.Routes()

	srv := &http.Server{
		Handler:      s.Router,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()

}

type DB struct {
	Bucket *gocb.Bucket
}

func (d *DB) ConnectDatabase() error {

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

	DB.Bucket = bucket

	return err

}
