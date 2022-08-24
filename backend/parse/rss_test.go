package parse_test

import (
	_ "embed"
	"testing"

	"github.com/knightspore/rss-reader-app/backend/parse"
)

func TestRSSUnmarshal(t *testing.T) {

	rss, err := parse.NewFeed("https://www.wired.com/feed/category/business/latest/rss")
	if err != nil {
		t.Log("Error creating new RSS Feed")
		t.Fatal(err)
	}

	_, err = parse.UnmarshalRSS(rss.Doc)
	if err != nil {
		t.Fatal(err)
	}

}