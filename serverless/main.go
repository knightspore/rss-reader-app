package main

import (
	"net/http"

	"github.com/knightspore/rss-reader-app/serverless/handler"
)

func main() {

	handler.CreateTypes("./../frontend/src/types/server.ts")

	// HTML Server with Router to Functions
	http.HandleFunc("/subscriptions/add", handler.AddSubscriptionHandler)
	http.HandleFunc("/subscriptions/remove", handler.RemoveSubscriptionHandler)
	http.HandleFunc("/posts/read", handler.GetPostTextHandler)

	// Start Server
	http.ListenAndServe(":8080", nil)

}
