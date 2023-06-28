package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx = NewValueInCtx(ctx, "key1", "value1")

	result := ctx.Value("key1")

	fmt.Println(result)

	// with new context
	//ctx2 := NewValueInCtx(context.Background(), "key2", "value2")
	//
	//req := http.Request{}
	//
	//req.Context()
}

func NewValueInCtx(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func Cancellation() {
	ctx, fn := context.WithCancel(context.Background())
	defer fn()

	select {
	case <-ctx.Done():
	// finish the program
	default:
		fmt.Println("still working")
	}
}

func Timeout() {
	ctx, fn := context.WithTimeout(context.Background(), 5*time.Minute)
	defer fn()

	select {
	case <-ctx.Done():
	// finish the program
	default:
		fmt.Println("still working")
	}
}
