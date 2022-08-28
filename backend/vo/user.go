package vo

import (
	"time"
)

type User struct {
	ID            string            `json:"id"`
	Subscriptions []string `json:"subscriptions"`
	ReadingList   []string      `json:"readingList"`
	LastUpdated   string            `json:"lastUpdated"`
}

func NewUser(id string) User {
	return User{
		ID:            id,
		Subscriptions: []string{},
		ReadingList:   []string{},
		LastUpdated:   time.Now().String(),
	}
}

// func (u *User) RefreshReadingList() {

// 	var unsorted []Article

// 	for _, s := range u.Subscriptions {
// 		for _, a := range s.Articles {
// 			if !a.IsRead {
// 				unsorted = append(unsorted, a.id)
// 			}
// 		}
// 	}

// 	// Remove Duplicates
// 	combined := append(unsorted, u.ReadingList...)
// 	var unique []Article
// 	in := make(map[string]bool)
// 	for _, a := range combined {
// 		if !in[a.ID] {
// 			in[a.ID] = true
// 			unique = append(unique, a)
// 		}
// 	}

// 	if len(unique) > 1 {
// 		// Order by PubDate (Unix Timestamp)
// 		sort.Slice(unique, func(i, j int) bool {
// 			return unsorted[i].PubEpoch > unsorted[j].PubEpoch
// 		})
// 	}

// 	u.ReadingList = unique

// 	u.LastUpdated = time.Now().String()

// }