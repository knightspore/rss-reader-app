package handler

import (
	"fmt"
	"net/http"
	"time"

	readability "github.com/go-shiori/go-readability"
)

func GetPostTextHandler(w http.ResponseWriter, r *http.Request) {

	url := "https://www.thelegalartist.com/blog/on-bill-wattersons-refusal-to-license-calvin-and-hobbes"

	readable, err := readability.FromURL(url, 30*time.Second)
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
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", post)))

}
