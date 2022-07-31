package vo

type User struct {
	ID string `json:"id"`
	Subscriptions []Subscription `json:"subscriptions"` 
	ReadingList []Article `json:"readingList"`
	LastUpdated string `json:"lastUpdated"`
}
