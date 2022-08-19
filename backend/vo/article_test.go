package vo_test

import (
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
