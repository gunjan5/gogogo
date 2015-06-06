package main

import (
"fmt"
"strings"
)
func main(){


var dir string
fmt.Scanf("%s", &dir)

switch dir=strings.ToLower(dir); dir  {
case "n":
fmt.Println("Go N!")
case "s":
fmt.Println("Go S!")
case "e":
fmt.Println("Go E!")
case "w":
fmt.Println("Go W!")
default:
fmt.Println("Go TH!")
}
}