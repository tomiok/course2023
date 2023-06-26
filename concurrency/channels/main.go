package main

import "fmt"

// channels

func main() {
	// channel declaration
	chInt := make(chan int)
	//chIntFixed := make(chan int,10)

	// that channels with capacity behaves different that the ones with no cap.
	//
	//chStr := make(chan string)
	//chStrFixed := make(chan string,1)

	// write into a channel

	// new operator!! -> <-
	go func() {
		chInt <- 1
		close(chInt) // close the channels when you finish the writing
	}()

	// reading the channel
	value := <-chInt

	fmt.Println(value) // print "hello world"
}

// blocking and non-blocking operations
// when you have a blocking operation, you need to use a new goroutine
