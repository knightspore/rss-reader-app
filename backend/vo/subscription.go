package vo

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/araddon/dateparse"
	"github.com/knightspore/rss-reader-app/backend/parse"
	"github.com/knightspore/rss-reader-app/backend/util"
	"github.com/segmentio/fasthash/fnv1a"
)

type Subscription struct {
	ID          string    `xml:"channel>id" json:"id"`
	Title       string    `xml:"channel>title" json:"title"`
	Description string    `xml:"channel>description" json:"description"`
	URL         string    `xml:"channel>link" json:"url"`
	LastUpdated string    `xml:"channel>updated" json:"lastUpdated"`
	Muted       bool      `json:"muted"`
	Icon        string    `json:"icon"`
	Articles    []string `json:"articles"`
}

type SubscriptionArticles struct {
	Items []Article `xml:"channel>item" json:"articleList"`
}

func NewSubscription(url string, title string) (Subscription, []Article, error) {

	feed, err := parse.NewFeed(url)
	if err != nil {
		return Subscription{}, []Article{}, err
	}

	sub, arts, err := NewSubscriptionFromJSON(feed.JSON)

	// Fill in Missing Sub Data
	if sub.Title != "" {
		sub.Title = title
	}
	if len(sub.ID) == 0 {
		sub.ID = url
	}
	if len(sub.URL) == 0 {
		sub.URL = url
	}
	if len(sub.Icon) == 0 {
		sub.Icon = "https://www.google.com/s2/favicons?domain=" + sub.URL + "&sz=48"
	}

	var populated []Article

	// Handle Articles
	for _, a := range arts.Items {
		a.Parent, a.ParentID, a.Icon = sub.Title, sub.ID, sub.Icon

		h := fnv1a.HashString32(sub.URL+a.Title)
		a.ID = strconv.FormatUint(uint64(h), 10)

		t, err := dateparse.ParseAny(a.PubDate)
		if err != nil {
			fmt.Printf("Error Parsing Time: %s / %s\n", a.Title, a.Parent)
		}

		ts := t.Unix()
		a.PubEpoch = ts

		sub.Articles = append(sub.Articles, a.ID)
		populated = append(populated, a)
	}

	return sub, populated, err

}

func NewSubscriptionFromJSON(j []byte) (Subscription, SubscriptionArticles, error) {
	var s Subscription
	var a SubscriptionArticles
	err := json.Unmarshal(j, &s)
	if err != nil {
		return s, a, err
	}
	err = json.Unmarshal(j, &a)
	if err != nil {
		return s, a, err
	}
	return s, a, nil
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
