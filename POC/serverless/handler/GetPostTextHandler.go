package handler

import (
	"encoding/json"
	"net/http"
	"time"

	readability "github.com/go-shiori/go-readability"
)

func GetPostTextHandler(w http.ResponseWriter, r *http.Request) {

	// Decode JSON

	var e Event
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create Readable Text from URL

	readable, err := readability.FromURL(e.ReadPostUrl, 30*time.Second)
	if err != nil {
		panic(err)
	}

	// Print the string
	post := PostContent{
		Title:   readable.Title,
		Byline:  readable.Byline,
		Site:    readable.SiteName,
		Text:    readable.TextContent,
		Favicon: readable.Favicon,
		URL:     e.ReadPostUrl,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

	// TODO: Add Hashed URLS to Small Database map[string]hash

}
