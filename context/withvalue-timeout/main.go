package main

import (
	"context"
	"fmt"
	"time"
)

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func doSomethingCooler(ctx context.Context) {
	rID := ctx.Value("request-id")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("doSomethingCooler: context is done (timeout)")
			return
		default:
			fmt.Printf("doing something cooler with %s\n", rID)
		}
		time.Sleep(500 * time.Millisecond)

	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	ctx = enrichContext(ctx)
	go doSomethingCooler(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("main: context is done (timeout)")
	}

	time.Sleep(2 * time.Second)
}
