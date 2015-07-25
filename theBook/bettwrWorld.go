package main

import . "fmt"

func main() {
	Println("I'm a ", peek())
	p := Sprintf("ayy %v", peek())
	Printf(p)
}

func peek() (msg string) {
	msg = "pikachu"
	return
}
