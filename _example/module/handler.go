package module

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func (s *Server) HandleUserCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewUserEvent(r.Body)
		if err != nil {
			log.Panic(err)
		}

		user := vo.NewUser(e.ID)

		err = s.UserUpsert(user)
		if err != nil {
			log.Panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(http.StatusOK)

	}
}

func (s *Server) HandleSubscriptionCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewSubscriptionEvent(r.Body)
		if err != nil {
			log.Panic(err)
		}

		err = s.SubscriptionCreate(e.URL, e.Title)
		if err != nil {
			log.Panic(err)
		}

		user, err := s.UserGet(e.UserID)
		if err != nil {
			log.Panic(err)
		}
		user.Subscriptions = append(user.Subscriptions, e.IDs...)
		err = s.UserUpsert(user)
		if err != nil {
			log.Panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(http.StatusOK)

	}
}

func (s *Server) HandleSubscriptionGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewSubscriptionEvent(r.Body)
		if err != nil {
			log.Panic(err)
		}

		sub, err := s.SubscriptionsGet(e.IDs)
		if err != nil {
			log.Panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(sub)

	}
}

func (s *Server) HandleReadingListGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewUserEvent(r.Body)
		if err != nil {
			log.Panic(err)
		}

		user, err := s.UserGet(e.ID)
		if err != nil {
			log.Panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(user.ReadingList)

	}
}

func (s *Server) HandleArticleRead() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewArticleEvent(r.Body)
		if err != nil {
			log.Panic(err)
		}

		_, err = s.UserGet(e.UserID)
		if err != nil {
			log.Panic(err)
		}

		// GetArticle Not Currently Working
		// article, err := user.GetArticle(e.URL)
		// if err != nil {
		// 	log.Panic(err)
		// }

		// Workaround
		article := vo.Article{
			URL: e.URL,
		}

		md, err := article.Read()
		if err != nil {
			log.Panic(err)
		}

		article.Content = md

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(article.Content)

		article.IsRead = true

	}
}
