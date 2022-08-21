package module

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/gorilla/mux"
	"github.com/knightspore/rss-reader-app/backend/vo"
)

type Server struct {
	DB     *gocb.Bucket
	Router *mux.Router
}

func (s *Server) Routes() {
	// Users
	s.Router.HandleFunc("/api/user/create", s.UserCreate())
	// Subscriptions
	s.Router.HandleFunc("/api/subscription/create", s.SubscriptionCreate())
	s.Router.HandleFunc("/api/subscription/get", s.SubscriptionGet())
	// Reading List
	s.Router.HandleFunc("/api/readinglist/get", s.ReadingListGet())
	s.Router.HandleFunc("/api/article/read", s.ArticleRead())
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

func (s *Server) UserGet(id string) (vo.User, error) {
	var user vo.User

	col := s.DB.Collection("users")
	result, err := col.Get(id, nil)
	if err != nil {
		return user, err
	}

	err = result.Content(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

func (s *Server) UserUpdate(user vo.User) error {
	col := s.DB.Collection("users")

	_, err := col.Upsert(user.ID, user, nil)
	if err != nil {
		return err
	}

	return nil
}
