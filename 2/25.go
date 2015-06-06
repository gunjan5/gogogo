package main

import "fmt"

func main(){
a:=[2][2]int{
{1,2},
{2,2},
}

b:=[2][3]int{
{1,2,3},
{4,5},
}

fmt.Println(a,b)


c:=[2][2]int{
{a[0][0]+b[1][1], a[1][0]},
{2,3},
}

fmt.Println(c)

var x [2][2][2]int
fmt.Println(x[1][1][1])
}