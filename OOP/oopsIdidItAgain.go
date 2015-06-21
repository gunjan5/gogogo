package main

import (
	"fmt"
)

var ()

type Rect struct {
	wid int
	hei int
}

type Rects []Rect

func (r *Rect) Area() int {
	return r.wid * r.hei
}

func (rs Rects) Area() int {
	var a int

	for _, r := range rs {
		a += r.Area()
	}
	return a
}

func main() {
	r := Rect{wid: 10, hei: 5}
	t := Rect{wid: 11, hei: 6}
	y := Rect{wid: 12, hei: 7}

	rs := Rects{r, t, y}

	fmt.Println("area: ", r.Area())
	fmt.Println("\n\n area: ", rs.Area())
}
