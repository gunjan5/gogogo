package main

import (
	"fmt"
	"strings"
)

func main() {
	text1 := "gunjan patel"
	text2 := "is crazy"

	fmt.Println(text1, text2, "\n\n")

	t1, t2 := changeTitleCaseIt1(text1, text2)

	fmt.Println(t1, t2, "\n\n")
}

func changeTitleCaseIt1(text1, text2 string) (t1_why_is_this_here, t2_why_is_this_here string) {

	text1 = strings.Title(text1)
	text2 = strings.ToUpper(text2)

	return text1, text2

}
