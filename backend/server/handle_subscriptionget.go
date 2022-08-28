package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func (s *Server) HandleSubscriptionGet() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewSubscriptionEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("Get Subscription: %q", e.ID)

		var sub vo.Subscription

		res, err := s.Get("subscriptions", e.ID)
		if err != nil {
			fmt.Println(err)
		}

		err = res.Content(&sub)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(sub)

	}
}