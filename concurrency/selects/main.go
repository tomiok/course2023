package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main3() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// new keyword - select.

	for {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			break
		}
	}
}

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
