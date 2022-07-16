package handler

import (
	"fmt"
	"time"

	readability "github.com/go-shiori/go-readability"
)

func main() {

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

	fmt.Println(post.Title)
	fmt.Println("")
	fmt.Println(post.Byline)
	fmt.Println("")
	fmt.Println(post.Site)
	fmt.Println("")
	fmt.Println(post.Favicon)
	fmt.Println("")
	fmt.Println("")
	fmt.Println(post.Text)

}
