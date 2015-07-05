package main

import . "fmt"

func main() {
	Println(message())
	print("yello", "humans")
	Println("I'm", peek())
	p := Sprintf("ayy %v", peek())
	Printf(p)
}

func message() (string, string) {
	return "hello", "world"
}

func peek() (msg string) {
	msg = "pikachu"
	return
}
func print(v ...interface{}) {
	Println(v...)
}
