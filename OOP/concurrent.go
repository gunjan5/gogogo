package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()
		time.Sleep(4 * time.Second)
		fmt.Println("hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("ay ay ay")
	}()

	waitGrp.Wait()
}
