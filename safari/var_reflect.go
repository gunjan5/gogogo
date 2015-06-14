package main

import (
	"fmt"
	"reflect"
)

var (
	name, course            string
	module                  float64
	name1, course1, module1 = "Gunjan", "Docker", 4.5
	blood                   = "A neg" //note, we use = in var and := if var is declared and initialized inside a func
)

func main() {
	blood1 := "A neg"
	fmt.Println("Name: ", name, "and type: ", reflect.TypeOf(name))
	fmt.Println("Module: ", module, "and type: ", reflect.TypeOf(module))

	fmt.Println("Name: ", name1, "blood type: ", blood1, "and type: ", reflect.TypeOf(blood1))
	fmt.Println("Module1: ", module1, "and type: ", reflect.TypeOf(module1))
}
