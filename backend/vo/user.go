package vo

import (
	"sort"
	"time"
)

type User struct {
	ID            string         `json:"id"`
	Subscriptions []Subscription `json:"subscriptions"`
	ReadingList   []Article      `json:"readingList"`
	LastUpdated   string         `json:"lastUpdated"`
}

func NewUser(id string) User {
	return User{
		ID:            id,
		Subscriptions: []Subscription{},
		ReadingList:   []Article{},
		LastUpdated:   time.Now().String(),
	}
}

func (u *User) RefreshReadingList() {

	var unsorted []Article

	for _, s := range u.Subscriptions {
		for _, a := range s.Articles {
			if !a.IsRead {
				unsorted = append(unsorted, a)
			}
		}
	}

	// Remove Duplicates
	combined := append(unsorted, u.ReadingList...)
	var unique []Article
	in := make(map[string]bool)
	for _, a := range combined {
		if !in[a.ID] {
			in[a.ID] = true
			unique = append(unique, a)
		}
	}

	if len(unique) > 1 {
		// Order by PubDate (Unix Timestamp)
		sort.Slice(unique, func(i, j int) bool {
			return unsorted[i].PubEpoch > unsorted[j].PubEpoch
		})
	}

	u.ReadingList = unique

	u.LastUpdated = time.Now().String()

}

func (u *User) AddSubscription(s Subscription) {

	u.Subscriptions = append(u.Subscriptions, s)
	in := make(map[string]bool)
	var unique []Subscription
	for _, item := range u.Subscriptions {
		if _, ok := in[item.ID]; !ok {
			in[item.ID] = true
			unique = append(unique, item)
		}
	}
	u.Subscriptions = unique

}

func (u *User) GetArticle(url string) (Article, error) {

	var a *Article

	for _, s := range u.Subscriptions {
		for _, f := range s.Articles {
			if f.URL == url {
				a = &f
			}
		}
	}

	return *a, nil

}
