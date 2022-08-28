package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func (s *Server) HandleUserGet() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewUserEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("Get User: %q", e.ID)

		var u vo.User

		res, err := s.Get("users", e.ID)
		if err != nil {
			fmt.Println(err)
		}

		err = res.Content(&u)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(u)

	}
}