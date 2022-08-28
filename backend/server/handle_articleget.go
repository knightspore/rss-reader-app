package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func (s *Server) HandleArticleGet() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewArticleEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("Get Article: %q (User: %q)", e.ID, e.UserID)

		var a vo.Article

		res, err := s.Get("articles", e.ID)
		if err != nil {
			fmt.Println(err)
		}

		err = res.Content(&a)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(a)

	}
}