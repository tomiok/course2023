package main

import (
	"context"
	"fmt"
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
