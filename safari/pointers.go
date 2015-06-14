package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10.0000000000000000
	b := 5
	ptr := &a

	//& to reference * to dereference
	fmt.Println("A: ", a, "Type: ", reflect.TypeOf(a), "location in memory: ", &a, ptr, *ptr)

	fmt.Println("B: ", b, "Type: ", reflect.TypeOf(b))

	c := int(a) + b
	fmt.Println("C: ", c, "Type: ", reflect.TypeOf(c))

}
