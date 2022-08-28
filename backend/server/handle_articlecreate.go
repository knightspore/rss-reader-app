package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *Server) HandleArticleCreate() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewArticleEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("Create Article: %q - %q (User: %q)", e.URL, e.ParentID, e.UserID)
		log.Printf(">> Not Working Currently")

		// TODO: Add NewArticle Code and Complete
	
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(http.StatusOK)

	}
}