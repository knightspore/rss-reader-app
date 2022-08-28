package parse

import (
	"encoding/xml"
)

type AtomFeed struct {
	ID          string    `xml:"id" json:"id"`
	Title       string    `xml:"title" json:"title"`
	Description string    `xml:"description" json:"description"`
	URL         string    `xml:"link" json:"url"`
	LastUpdated string    `xml:"updated" json:"lastUpdated"`
	Muted       bool      `json:"muted"`
	Icon        string    `json:"icon"`
	Articles    []AtomEntry `xml:"entry" json:"articleList"`
}

type AtomEntry struct {
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

func UnmarshalAtom(doc []byte) (AtomFeed, error) {
	var f AtomFeed
	err := xml.Unmarshal(doc, &f)
	if err != nil {
		return AtomFeed{}, err
	}
	return f, nil
}