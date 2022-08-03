package vo

import (
	"time"
)

type User struct {
	ID string `json:"id"`
	Subscriptions []Subscription `json:"subscriptions"` 
	ReadingList []Article `json:"readingList"`
	LastUpdated string `json:"lastUpdated"`
}

func (u *User) RefreshReadingList() {

	for _, s := range(u.Subscriptions) {
		for _, a := range(s.Articles) {
			if true { // TODO: If not in reading list
				u.ReadingList = append(u.ReadingList, a)
			}
		}
	}

	u.LastUpdated = time.Now().String()

}
