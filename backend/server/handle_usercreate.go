package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func (s *Server) HandleUserCreate() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := vo.NewUserEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		user := vo.NewUser(e.ID)
		err = s.Upsert("users", user.ID, user)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(http.StatusOK)

	}
}