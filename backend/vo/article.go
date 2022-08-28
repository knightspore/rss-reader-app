package vo

import (
	"encoding/json"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/knightspore/rss-reader-app/backend/util"
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
	Icon        string `json:"icon"`
}

func NewArticleFromJSON(j []byte) (Article, error) {
	var a Article
	err := json.Unmarshal(j, &a)
	if err != nil {
		return a, err
	}
	return a, nil
}

func (a *Article) Read() (string, error) {

	if a.Content != "" {
		return a.Content, nil
	}

	data, err := util.Fetch(a.URL)
	if err != nil {
		return "", err
	}

	// TODO: LCA for Body

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(string(data))
	if err != nil {
		return "", err
	}

	return markdown, err

}
