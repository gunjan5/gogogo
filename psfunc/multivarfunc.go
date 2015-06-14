package main

import (
	"fmt"
)

const (
	k = iota * 2
)

func main() {

	goog := bestStock(405, 345, 346, 678, 234, 435, 456, 567)
	fmt.Println(goog)
	fmt.Println(k)
}

func bestStock(stock ...int) int {
	best := stock[0]
	for _, i := range stock {
		if i > best {
			best = i
		}
	}
	return best
}
