package vo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"github.com/araddon/dateparse"
	"github.com/knightspore/rss-reader-app/backend/parse"
	"github.com/knightspore/rss-reader-app/backend/util"
)

type Subscription struct {
	ID          string    `xml:"channel>id" json:"id"`
	Title       string    `xml:"channel>title" json:"title"`
	Description string    `xml:"channel>description" json:"description"`
	URL         string    `xml:"channel>link" json:"url"`
	LastUpdated string    `xml:"channel>updated" json:"lastUpdated"`
	Muted       bool      `json:"muted"`
	Icon        string    `json:"icon"`
	Articles    []Article `xml:"channel>item" json:"articles"`
}

func NewSubscription(url string, title string) (Subscription, error) {

	feed, err := parse.NewFeed(url)
	if err != nil {
		return Subscription{}, err
	}

	s, err := NewSubscriptionFromJSON(feed.JSON)

	fmt.Printf("%+v",s)

	// Fill in Missing Data
	if title != "" {
		s.Title = title
	}
	if len(s.ID) == 0 {
		s.ID = url
	}
	if len(s.URL) == 0 {
		s.URL = url
	}
	if len(s.Icon) == 0 {
		s.Icon = "https://www.google.com/s2/favicons?domain=" + s.URL + "&sz=48"
	}

	// Fill in Article Data from Parent
	for i, a := range s.Articles {
		a.Parent = s.Title
		a.Icon = s.Icon
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

func NewSubscriptionFromJSON(j []byte) (Subscription, error) {
	var s Subscription
	err := json.Unmarshal(j, &s)
	if err != nil {
		return s, err
	}
	return s, nil
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
