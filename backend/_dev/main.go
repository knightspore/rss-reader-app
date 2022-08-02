package main

import (
	"fmt"

	subscription "github.com/knightspore/rss-reader-app/backend/module/Subscription"
	"github.com/knightspore/rss-reader-app/backend/vo"
)

func main() {

	urls := [15]string{
		"http://techcrunch.com/feed/",
		"https://www.wired.com/feed/category/backchannel/latest/rss",
		"https://www.protocol.com/feeds/feed.rss",
		"https://timdaub.github.io/atom.xml",
		"https://flaviocopes.com/index.xml",
		"https://www.emgoto.com/rss.xml",
		"https://kentcdodds.com/blog/rss.xml",
		"https://steve-yegge.blogspot.com/atom.xml",
		"https://jovicailic.org/feed/",
		"https://machinelearningmastery.com/feed/",
		"http://lesswrong.com/.rss",
		"https://www.bellingcat.com/feed",
		"https://twobithistory.org/feed.xml",
		"https://ciaran.co.za/rss.xml",
		"https://news.ycombinator.com/rss",
	}

	var subs []vo.Subscription

	for _, url := range urls {
		s, err := subscription.Create(url)
		if err != nil {
			fmt.Println(err)
		} else {
			subs = append(subs, s)
		}
	}

	var user vo.User
	user.ID = "0001"
	user.Subscriptions = append(user.Subscriptions, subs)

	fmt.Println(user.ReadingList[:10])
	fmt.Printf("Last Updated At: "%s time.Now())

}
