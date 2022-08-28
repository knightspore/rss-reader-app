package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func (s *Server) HandleSubscriptionCreate() Handler {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewSubscriptionEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("Create Subscription: %q (User: %q)", e.ID, e.UserID)

		sub, arts, err := vo.NewSubscription(e.URL, e.Title)
		if err != nil {
			fmt.Println(err)
		}

		err = s.Upsert("subscriptions", e.ID, sub)
		if err != nil {
			fmt.Println(err)
		}

		for _, a := range(arts) {
			err = s.Upsert("articles", a.ID, a)
			if err != nil {
				fmt.Println(err)
			}
		}

		var u vo.User

		res, err := s.Get("users", e.UserID)
		if err != nil {
			fmt.Println(err)
		}

		err = res.Content(&u)
		if err != nil {
			fmt.Println(err)
		}

		u.Subscriptions = append(u.Subscriptions, sub.ID)

		err = s.Upsert("users", u.ID, u)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(http.StatusOK)

	}
}
