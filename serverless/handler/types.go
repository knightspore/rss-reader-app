package handler

import (
	"fmt"

	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func CreateTypes() {
	converter := typescriptify.New().
		Add(Feed{}).
		Add(Event{}).
		Add(PostContent{})

	err := converter.ConvertToFile("../shared/types.ts")
	if err != nil {
		fmt.Println(err)
	}
}
