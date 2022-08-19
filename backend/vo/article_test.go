package vo_test

import (
	"testing"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

var article = vo.Article{
	ID:          "123",
	PubDate:     "Fri, 19 Aug 2022 11:29:09 +0000",
	Title:       "Test Article",
	Description: "This is a test article",
	Content:     "This is the content of the article",
	URL:         "http://test.com",
	IsRead:      false,
	Parent:      "https://test.com/xml",
}

func TestTimestamp(t *testing.T) {
	want := "1660908549"

	got, err := article.Timestamp()
	if err != nil {
		t.Errorf("Error parsing timestamp: %s", err)
	}

	if got != want {
		t.Errorf("Expected %s, got %s", want, got)
	}

}
