package main

import (
	"fmt"
)

func main() {

	var x int

	fmt.Scanf("%d", &x)

	if msg := "Hi!, "; x >= 22 {
		fmt.Println(msg, "yooo")
	} else if x >= 18 && x < 22 {
		fmt.Println(msg, "meh")
	} else {
		fmt.Println(msg, "get out")
	}

}
