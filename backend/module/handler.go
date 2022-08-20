package module

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func (s *Server) UserCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewUserEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		user := vo.NewUser(e.ID)

		err = s.UserUpdate(user)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(user)

	}
}


func (s *Server) SubscriptionCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewSubscriptionEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		sub, err := vo.NewSubscription(e.URL, e.Title)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(http.StatusOK)

		user, err := s.UserGet(e.UserID)
		if err != nil {
			fmt.Println(err)
		}

		user.AddSubscription(sub)
		user.RefreshReadingList()

		s.UserUpdate(user)

	}
}

func (s *Server) SubscriptionGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewUserEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		user, err := s.UserGet(e.ID)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(user.Subscriptions)

	}
}

func (s *Server) ReadingListGet() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		
		e, err := NewUserEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		user, err := s.UserGet(e.ID)
		if err != nil {
			fmt.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(user.ReadingList)

	}
}

func (s *Server) ArticleRead() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		e, err := NewArticleEvent(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		user, err := s.UserGet(e.UserID)
		if err != nil {
			fmt.Println(err)
		}

		article, err := user.GetArticle(e.URL)
		if err != nil {
			fmt.Println(err)
		}

		md, err := article.Read()
		if err != nil {
			fmt.Println(err)
		}

		article.Content = md

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(article.Content)

		article.IsRead = true
		user.RefreshReadingList()
		s.UserUpdate(user)

	}

}
