package server

import (
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
)

func (s *Server) Get(c string, id string) (*gocb.GetResult, error) {
	col := s.DB.Collection(c)
	result, err := col.Get(id, nil)
	if err != nil {
		return nil,  err
	}
	return result, nil
}

func(s *Server) Upsert(c string, id string, val interface{}) error {
		col := s.DB.Collection(c)
		_, err := col.Upsert(id, val, nil)
		if err != nil {
			return err
		}
		return nil
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

