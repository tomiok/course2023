package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	ch := make(chan string)
	chQuit := make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		for {
			select {
			case value, ok := <-ch:
				fmt.Println(value, ok)

			case <-chQuit:
				fmt.Println("out")
			}
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- "Goroutine : " + strconv.Itoa(i)
		}(i)
	}

	wg.Wait()
	chQuit <- true
	close(ch)
}
