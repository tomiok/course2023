package main

import (
	"fmt"
)

type Car struct {
	ID int
}

func main() {
	chCar := make(chan Car)

	go func() {
		defer close(chCar)
		for i := 0; i < 10; i++ {
			chCar <- Car{ID: i}
		}
	}()

	// non-blocking operation
	// this will read values until the channel is closed.

	for v := range chCar {
		fmt.Println(v)
	}
}
