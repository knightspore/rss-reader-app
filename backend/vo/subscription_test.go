package vo_test

import (
	"fmt"
	"testing"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func TestNewSubscription(t *testing.T) {
	s, arts, err := vo.NewSubscription("https://ciaran.co.za/rss.xml", "RSS Test")
	if err != nil {
		t.Fatal(err)
	}

	if len(arts) == 0 || s.Title == "" {
		fmt.Println(err)
	}

	s, arts, err = vo.NewSubscription("https://www.groundup.org.za/sitenews/atom/", "Atom Test")
	if err != nil {
		t.Fatal(err)
	}

	if len(arts) == 0 || s.Title == "" {
		fmt.Println(err)
	}
}
