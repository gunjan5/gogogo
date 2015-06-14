package main

import (
	"fmt"
	"os"
)

const (
	//constant pi
	PI = 3.14
)

func main() {
	user := os.Getenv("USER")

	fmt.Print("\n\n\n system user: ", user, "\n\n\n\n Pi: ", PI, "\n\n\n")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
