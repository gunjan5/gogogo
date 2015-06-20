package main

import (
	"fmt"
)

var ()

type Rect struct {
	wid int
	hei int
}

func (r *Rect) Area() int {
	return r.wid * r.hei
}

func main() {
	r := Rect{wid: 10, hei: 5}
	fmt.Println("area: ", r.Area())
}
