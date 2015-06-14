package main

import (
	"fmt"
)

func main() {

	goog := bestStock(405, 345, 346, 678, 234, 435, 456, 567)
	fmt.Println(goog)
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
