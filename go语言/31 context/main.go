package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	worker3()
}

func worker1() {
	deep := 10
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	go handle(ctx, 500*time.Millisecond, deep)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func worker2() {
	deep := 10
	ctx, cancel := context.WithCancel(context.Background())
	go handle(ctx, 500*time.Millisecond, deep)
	time.Sleep(1 * time.Second)
	cancel()
}

func worker3() {
	deep := 10

	ctx := context.WithValue(context.Background(), "token", "chihuo")

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	go handle(ctx, 500*time.Millisecond, deep)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration, deep int) {
	if deep > 0 {
		time.Sleep(200 * time.Millisecond)
		go handle(ctx, duration, deep-1)
	}

	if ctx.Value("token") != nil {
		fmt.Printf("token is %s\n", ctx.Value("token"))
	}

	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Printf("process request with %v, %d\n",
			duration, deep)
	}
}
