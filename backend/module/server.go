package module

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/knightspore/rss-reader-app/backend/vo"
)

type Server struct {
	DB     *sql.DB
	Router *mux.Router
	Users []vo.User // TODO: For TUI Replication Purposes
}

func (s *Server) Routes() {
	s.Router.HandleFunc("/api/user/create", s.handleUserCreate)
	s.Router.HandleFunc("/api/subscription/create", s.handleSubscriptionCreate)
	s.Router.HandleFunc("/api/readinglist/get", s.handleGetReadingList)
}

func (s *Server) Start() {

	s.Routes()
	s.Router.Headers("Content-Type", "application/json")

	srv := &http.Server{
		Handler:      s.Router,
		Addr:         "127.0.0.1:1337",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()

}

// Handler Functions
// For now, these just handle the direct logic required for the route
// These should be moved elsewhere in the module package
// and these routes should only handle HTTP

type UserEvent struct {
	ID string `json:"id"`
}

type SubscriptionEvent struct {
	ID string `json:"id"`
	Title string `json:"title"`
	URL   string `json:"url"`
	UserID string `json:"userId"` 
}

func (s *Server) handleUserCreate(w http.ResponseWriter, r *http.Request) {

	var e UserEvent
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
	}

	user := vo.NewUser(e.ID)

	s.Users = append(s.Users, user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (s *Server) handleSubscriptionCreate(w http.ResponseWriter, r *http.Request) {

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

	for _, u := range s.Users {
		if u.ID == e.UserID {
			user = u
		}
	}

	user.AddSubscription(sub)
	user.RefreshReadingList()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ReadingList)

}

func (s *Server) handleGetReadingList(w http.ResponseWriter, r *http.Request) {

	var e UserEvent
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		fmt.Println(err)
	}
	
	var user vo.User
	for _, u := range s.Users {
		if u.ID == e.ID {
			user = u
		}
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.ReadingList)

}