package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

// Dev Setup

var urls = [3]string{
	"https://flaviocopes.com/index.xml",
	"https://www.emgoto.com/rss.xml",
	"https://kentcdodds.com/blog/rss.xml",
}

var user vo.User

func main() {

	// Set Username
	UserNameSetup()
	AddTestArticles()

	list := getItem(0,4)

	for i, item := range list {
		content, err := item.Read()
		if err != nil {
			content = "Couldn't Read Article Content"
		}
		
		fmt.Println(">>> - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
		fmt.Println("|")
		fmt.Println("|  \033[1m" + item.Title + "\033[0m")
		fmt.Println("|  " + item.URL)
		fmt.Println("|")
		if (i == 0) {
		fmt.Println(content)
		fmt.Println("|")
		}
		fmt.Println("* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")

	}

}

func getItem(n,m int) []vo.Article {
	return user.ReadingList[n:m]
}

func UserNameSetup() {
	uuid := flag.String("n", "neo", "user name")
	flag.Parse()
	user.ID = *uuid
}

func AddTestArticles() {
	for _, url := range urls {
		s, err := vo.NewSubscription(url, url)
		if err != nil {
			fmt.Println(err)
		} else {
			user.Subscriptions = append(user.Subscriptions, s)
		}
	}

	user.LastUpdated = time.Now().String()
	user.RefreshReadingList()	
}