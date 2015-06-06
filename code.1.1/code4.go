package main

import "fmt"

func main(){
	var radius float32
	fmt.Println("Enter Radius:")
	fmt.Scanf("%f", &radius)
	
	var len float32
	len = 2 * 3.14 * radius
	
	fmt.Println("len:", len)
	fmt.Println("rad:", radius)
}