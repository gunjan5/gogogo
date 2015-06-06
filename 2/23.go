package main

import "fmt"


var (
  x int
  px *int
  )
  
func main(){

x= 11
px=&x
*px=22
fmt.Println(x, px, *px, &x)
} 