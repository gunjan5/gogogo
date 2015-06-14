package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "gunjan patel"

	fmt.Println(text, "\n\n")

	changeTitleCaseIt(&text)

	fmt.Println(text, "\n\n")
}

func changeTitleCaseIt(text *string) {

	*text = strings.Title(*text)

}
