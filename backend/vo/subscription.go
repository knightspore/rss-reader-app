package vo

import (
	"encoding/xml"
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/knightspore/rss-reader-app/backend/util"
)

type Subscription struct {
	ID          string    `xml:"channel>id" json:"id"`
	Title       string    `xml:"channel>title" json:"title"`
	Description string    `xml:"channel>description" json:"description"`
	URL         string    `xml:"channel>link" json:"url"`
	LastUpdated string    `xml:"channel>updated" json:"lastUpdated"`
	Muted       bool      `json:"channel>muted"`
	Icon        string    `json:"channel>icon"`
	Articles    []Article `xml:"channel>item" json:"articles"`
}

func NewSubscription(url string, title string) (Subscription, error) {

	data, err := util.Fetch(url)
	if err != nil {
		fmt.Printf("Error Fetching Feed URL: %s", url)
		return Subscription{}, err
	}

	var s Subscription
	err = xml.Unmarshal(data, &s)
	if err != nil {
		fmt.Printf("Error Unmarshaling XML: %s", url)
		return Subscription{}, err
	}

	// Fill in Missing Data

	if title != "" {
		s.Title = title
	}

	if len(s.ID) == 0 {
		s.ID = url
	}

	for i, a := range s.Articles {
		a.Parent = s.ID
		t, err := dateparse.ParseAny(a.PubDate)
		if err != nil {
			fmt.Printf("Error Parsing Time: %s / %s", a.Title, a.Parent)
		}

		ts := t.Unix()
		a.PubEpoch = ts

		s.Articles[i] = a
	}

	return s, err

}

func (s *Subscription) LatestArticles() *[]Article {

	data, err := util.Fetch(s.URL)
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
