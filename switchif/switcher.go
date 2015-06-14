package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	switch "docker" {

	case "abc":
		fmt.Print("\nooye")
	case "docker":
		fmt.Print("\nyaaaaaay!!")
		fallthrough
	case "bloody":
		fmt.Print("\nbloody hell")
	default:
		fmt.Print("\ngo home")

	}

	newTest()
	fmt.Print("\n")
}

func newTest() {

	switch tmp := random(); tmp {

	case 0, 2, 4, 6, 8:
		fmt.Print("\neven ", tmp)
	case 1, 3, 5, 7, 9:
		fmt.Print("\nOdd ", tmp)

	}

}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
