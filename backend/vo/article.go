package vo

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/knightspore/rss-reader-app/backend/module/Network"
)

type Article struct {
	ID          string `xml:"guid" json:"id"`
	PubDate     string `xml:"pubDate" json:"pubDate"`
	PubEpoch    int64  `xml:"pubEpoch" json:"pubEpoch"`
	Title       string `xml:"title" json:"title"`
	Description string `xml:"description" json:"description"`
	Content     string `xml:"content" json:"content"`
	URL         string `xml:"link" json:"url"`
	IsRead      bool   `json:"read"`
	Parent      string `json:"parent"`
}

func (a *Article) Get() (string, error) {
	data, err := Network.Fetch(a.URL)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (a *Article) Read() (string, error) {
	data, err := a.Get()
	if err != nil {
		return "", err
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(string(data))
	if err != nil {
		return "", err
	}

	return markdown, err
}