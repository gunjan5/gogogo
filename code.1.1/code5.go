package main

import "fmt"
/*
go tutorial 5
*/
func main(){
	var radius float32
	//print
	fmt.Println("Enter Radius:")
	//scan
	fmt.Scanf("%f", &radius)
	
	var len float32
	len = 2 * 3.14 * radius
	
	fmt.Println("len:", len)
	fmt.Println("rad:", radius)
}