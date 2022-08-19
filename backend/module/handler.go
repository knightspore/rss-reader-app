package module

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

type UserEvent struct {
	ID string `json:"id"`
}

type SubscriptionEvent struct {
	ID string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	UserID string `json:"userId"` 
}

func (s *Server) UserCreate(w http.ResponseWriter, r *http.Request) {

	var e UserEvent
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
	}

	user := vo.NewUser(e.ID)

	col := s.DB.Collection("users")
	_, err = col.Upsert(user.ID, user, nil)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (s *Server) SubscriptionCreate(w http.ResponseWriter, r *http.Request) {

	var e SubscriptionEvent
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
	}

	sub, err := vo.NewSubscription(e.URL, e.Title)
	if err != nil {
		fmt.Println(err)
	}

	var user vo.User

	col := s.DB.Collection("users")
	result, err := col.Get(e.UserID, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = result.Content(&user)
	if err != nil {
		fmt.Println(err)
	}

	user.AddSubscription(sub)
	user.RefreshReadingList()

	_, err = col.Upsert(e.UserID, vo.User{
		ID: user.ID,
		Subscriptions: user.Subscriptions,
		ReadingList: user.ReadingList,
	}, nil)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Subscriptions)

}

func (s *Server) ReadingListGet(w http.ResponseWriter, r *http.Request) {

	var e UserEvent
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
	}
	
	var user vo.User

	col := s.DB.Collection("users")
	result, err := col.Get(e.ID, nil)
	if err != nil {
		fmt.Println(err)
	}

	err = result.Content(&user)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ReadingList)

}