package handler

import "fmt"

var in_mem_db = map[User][]Feed{}

func printDB() {
	for k, v := range in_mem_db {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func AddSubscription(e *Event) {
	// TODO: Handle "If doesn't already exist"
	if _, ok := in_mem_db[User(e.UserID)]; ok {
		in_mem_db[User(e.UserID)] = append(in_mem_db[User(e.UserID)], e.Feeds...)
	} else {
		in_mem_db[User(e.UserID)] = e.Feeds
	}
}

func RemoveSubscription(e *Event) {
	if _, ok := in_mem_db[User(e.UserID)]; ok {
		for i, f := range in_mem_db[User(e.UserID)] {
			if f.URL == e.Feeds[0].URL {
				in_mem_db[User(e.UserID)] = append(in_mem_db[User(e.UserID)][:i], in_mem_db[User(e.UserID)][i+1:]...)
			}
		}
	}
}
