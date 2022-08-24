package parse

import (
	"encoding/xml"
)

type RSSFeed struct {
	ID          string    `xml:"channel>id" json:"id"`
	Title       string    `xml:"channel>title" json:"title"`
	Description string    `xml:"channel>description" json:"description"`
	URL         string    `xml:"channel>link" json:"url"`
	LastUpdated string    `xml:"channel>updated" json:"lastUpdated"`
	Muted       bool      `json:"muted"`
	Icon        string    `json:"icon"`
	Articles    []RSSItem `xml:"channel>item" json:"articles"`
}

type RSSItem struct {
	ID          string `xml:"guid" json:"id"`
	PubDate     string `xml:"pubDate" json:"pubDate"`
	PubEpoch    int64  `xml:"pubEpoch" json:"pubEpoch"`
	Title       string `xml:"title" json:"title"`
	Description string `xml:"description" json:"description"`
	Content     string `xml:"content" json:"content"`
	URL         string `xml:"link" json:"url"`
	IsRead      bool   `json:"read"`
	Parent      string `json:"parent"`
	Icon 			string `json:"icon"`
}

func UnmarshalRSS(doc []byte) (RSSFeed, error) {
	var f RSSFeed
	err := xml.Unmarshal(doc, &f)
	if err != nil {
		return RSSFeed{}, err
	}
	return f, nil
}