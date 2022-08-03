package vo

import (
	"encoding/xml"
	"fmt"

	"github.com/knightspore/rss-reader-app/backend/module/Network"
)

type Subscription struct {
	ID string `xml:"channel>id" json:"id"`
	Title string `xml:"channel>title" json:"title"`
	Description string `xml:"channel>description" json:"description"`
	URL        string   `xml:"channel>link" json:"url"`
	LastUpdated string `xml:"channel>updated" json:"lastUpdated"`
	Muted bool `json:"channel>muted"`
	Icon string `json:"channel>icon"`
	Articles []Article `xml:"channel>item" json:"articles"`
}

func NewSubscription(url string, title string) (Subscription, error) {

	data, err := Network.Fetch(url)
	if err != nil {
		fmt.Printf("Error Fetching Feed URL: %s", url)
		return Subscription{}, err	
	}

	var s Subscription
	err = xml.Unmarshal(data, &s)

	if err != nil {
		fmt.Printf("Error Unmarshaling XML: %s", url)
	}

	if title != "" {
		s.Title = title
	}

	if len(s.ID) == 0 {
		s.ID = url
	}

	for _, a := range s.Articles {
		a.Parent = s.ID
	}

	return s, err

}

func (s *Subscription) LatestArticles() *[]Article {

	data, err := Network.Fetch(s.URL)
	if err != nil {
		fmt.Printf("Error Fetching Feed URL: %s", s.URL)
	}

	var a []Article
	err = xml.Unmarshal(data, &a)

	if err != nil {
		fmt.Printf("Error UnMarhshaling XML: %s", s.URL)
	}

	// TODO: Only return NEW Articles

	return &a

}

