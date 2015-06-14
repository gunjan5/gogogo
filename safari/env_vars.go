package main

import (
	"fmt"
	"os"
)

func main() {
	user := os.Getenv("USER")

	fmt.Print("system user: ", user, "\n\n\n\n")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}
