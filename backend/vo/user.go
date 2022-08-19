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

func (u *User) RefreshReadingList() {

	var unsorted []Article

	for _, s := range u.Subscriptions {
		for _, a := range s.Articles {
			if true { // TODO: If not in reading list
				unsorted = append(unsorted, a)
			}
		}
	}

	// Order by PubDate (Unix Timestamp)
	sort.Slice(unsorted, func(i, j int) bool {
		return unsorted[i].PubEpoch > unsorted[j].PubEpoch
	})

	u.ReadingList = unsorted

	u.LastUpdated = time.Now().String()

}
