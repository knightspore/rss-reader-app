package vo_test

import (
	"testing"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

func TestNewSubscription(t *testing.T) {
	_, err := vo.NewSubscription("https://ciaran.co.za/rss.xml", "RSS Test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = vo.NewSubscription("https://www.groundup.org.za/sitenews/atom/", "Atom Test")
	if err != nil {
		t.Fatal(err)
	}
}