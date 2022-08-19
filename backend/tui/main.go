package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"github.com/knightspore/rss-reader-app/backend/vo"
)

// Dev Setup

var urls = []string{
	"https://www.emgoto.com/rss.xml",
	"https://kentcdodds.com/blog/rss.xml",
	"https://ciaran.co.za/rss.xml",
}

var user vo.User

func main() {

	// Set Username
	UserNameSetup() // Working on API
	AddTestArticles() // Working on API

	list := getItem(0, 5) // TODO: Add to API

	// Print JSON
	b, err := json.Marshal(list)
	if err != nil {
		fmt.Println("Error Marshaling JSON")
	}

	fmt.Println(string(b))


}

func getItem(n, m int) []vo.Article {
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
			fmt.Println("Error Parsing URL:" + url)
			fmt.Println(err)
		}
		user.Subscriptions = append(user.Subscriptions, s)
	}

	user.LastUpdated = time.Now().String()
	user.RefreshReadingList()
}
