package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func (s *Server) HandleSubscriptionCreate() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(http.StatusOK)

		events, err := NewSubscriptionEvents(r.Body)
		if err != nil {
			log.Panic(err)
		}

		var toAdd []string

		for _, e := range events {

			log.Printf("Create Subscription: %q (User: %q)", e.ID, e.UserID)

			sub, arts, err := vo.NewSubscription(e.URL, e.Title)
			if err != nil {
				log.Panic(err)
			}

			err = s.Upsert("subscriptions", sub.ID, sub)
			if err != nil {
				log.Panic(err)
			}

			for _, a := range arts {
				err = s.Upsert("articles", a.ID, a)
				if err != nil {
					log.Panic(err)
				}
			}

			toAdd = append(toAdd, sub.ID)

		}

		var u vo.User

		res, err := s.Get("users", events[0].UserID)
		if err != nil {
			log.Panic(err)
		}

		err = res.Content(&u)
		if err != nil {
			log.Panic(err)
		}

		u.Subscriptions = append(u.Subscriptions, toAdd...)

		err = s.Upsert("users", u.ID, u)
		if err != nil {
			log.Panic(err)
		}


	}
}
