package vo

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/knightspore/rss-reader-app/backend/module/Network"
)

type Article struct {
	ID string `xml:"guid" json:"id"`
	PubDate     string `xml:"pubdate"`
	Title string `xml:"title" json:"title"`
	Description string `xml:"description" json:"description"`
	Content string `xml:"content" json:"content"`
	URL string `xml:"link" json:"url"`
	Read bool `json:"read"`
}

func (a *Article) getContent() (string, error) {
	data, err :=	Network.Fetch(a.URL)
	if err != nil  {
		return "Error Fetching Article Content", err
	}
	return string(data), err
}

func (a *Article) Markdown() (string, error) {
	data, err := a.getContent()
	if err != nil {
		return "Error Convering Article to Markdown", err
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(data)

	return markdown, err
}