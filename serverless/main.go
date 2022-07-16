package main

import (
	"net/http"

	"github.com/knightspore/rss-reader-app/serverless/handler"
)

func main() {

	handler.CreateTypes()

	// HTML Server with Router to Functions
	http.HandleFunc("/subscriptions/add", handler.AddSubscriptionHandler)
	http.HandleFunc("/subscriptions/remove", handler.RemoveSubscriptionHandler)
	// http.HandleFunc("/reading-list/generate", ListSubscriptionsHandler)

	// Start Server
	http.ListenAndServe(":8080", nil)

}
