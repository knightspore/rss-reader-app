package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func (s *Server) HandleSubscriptionGet() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		events, err := NewSubscriptionEvents(r.Body)
		if err != nil {
			log.Panic(err)
		}

		var subs []vo.Subscription

		for _, e := range events {

			log.Printf("Get Subscription: %q", e.ID)

			var sub vo.Subscription

			res, err := s.Get("subscriptions", e.ID)
			if err != nil {
				log.Panic(err)
			}

			err = res.Content(&sub)
			if err != nil {
				log.Panic(err)
			}

			subs = append(subs, sub)

		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(subs)

	}
}