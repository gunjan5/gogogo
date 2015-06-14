package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("bloody hello ", runtime.GOOS)
}
