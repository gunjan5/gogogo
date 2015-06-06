package main

import "fmt"

func main(){
	s:=[]int{1,2,3}
	s[1]=22
	fmt.Println(s[1])
	
	var s1[]int
	s1 = make([]int,15,25)
	s1 = append(s1,11111)
	s1=append(s1,s...)
	
	s2:=s1[5:18:25]
	fmt.Println(s1, "len", len(s1), "cap", cap(s1), "s2" , s2)
copy( s2, s1[4:18])
fmt.Println(s2) //??

a:=[3]int{1,2,3}
fmt.Println(a) 

}