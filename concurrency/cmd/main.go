package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printHello(i, &wg)
	}
	wg.Wait()
}

func printHello(s any, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

// main is already a go routine
// go -> lunch a new goroutine

// wait groups
// add a number of goroutines to you want to wait.
