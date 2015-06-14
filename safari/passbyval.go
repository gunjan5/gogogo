package main

import "fmt"

var (
	name  = "Gunjan"
	lname = "Patel"
)

func main() {

	IQ := 24

	fmt.Println("my name is ", name, " ", lname, " and my IQ is ", IQ, ":(")

	takeIQenhancingPills(IQ)
	fmt.Println("my name is ", name, " ", lname, " and my IQ is still", IQ, ":(")

}

func takeIQenhancingPills(IQ int) int {
	IQ = IQ + 100

	fmt.Println("my name is ", name, " ", lname, " and my IQ is ", IQ, ":)")

	return IQ

}
