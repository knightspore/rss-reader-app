package handler

import (
	"fmt"

	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func CreateTypes(p string) {
	converter := typescriptify.New().
		Add(Feed{}).
		Add(Event{}).
		Add(PostContent{})

	converter.CreateInterface = true

	err := converter.ConvertToFile(p)
	if err != nil {
		fmt.Println(err)
	}
}
