package main_test

import (
	"testing"

	"github.com/knightspore/rss-reader-app/backend/parse"
	"github.com/knightspore/rss-reader-app/backend/vo"
)

func TestCreateRSSFeedSubscription(t *testing.T) {

	rss, err := parse.NewFeed("https://www.groundup.org.za/sitenews/rss/")
	if err != nil {
		t.Log("Error creating new RSS Feed")
		t.Fatal(err)
	}

	s, _, err := vo.NewSubscriptionFromJSON(rss.JSON)
	if err != nil {
		t.Fatalf("Error in vo.Subscription: %+v", err)
	}

	want := "GroundUp News"
	got := s.Title
	if got != want {
		t.Fatalf("Error parsing RSS Feed (Title) - Want: %s / Got: %s", want, got)
	}

	want_len := 15
	got_len := len(s.Articles)
	if got_len != want_len {
		t.Fatalf("Error parsing RSS Feed (Article Length) - Want: %d / Got: %d", want_len, got_len)
	}
}

func TestCreateAtomFeedSubscription(t *testing.T) {

	atom, err := parse.NewFeed("https://www.groundup.org.za/sitenews/atom/")
	if err != nil {
		t.Log("Error creating new Atom Feed")
		t.Fatal(err)
	}

	s, _, err := vo.NewSubscriptionFromJSON(atom.JSON)
	if err != nil {
		t.Fatalf("Err")
	}

	want := "GroundUp News"
	got := s.Title
	if got != want {
		t.Fatalf("Error parsing Atom Feed (Title) - Want: %s / Got: %s", want, got)
	}

	want_len := 15
	got_len := len(s.Articles)
	if got_len != want_len {
		t.Fatalf("Error parsing Atom Feed (Article Length) - Want: %d / Got: %d", want_len, got_len)
	}

}
