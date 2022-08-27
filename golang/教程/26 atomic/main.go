package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var c int64

func worker1(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		c++
	}
	wg.Done()
}

func worker2(wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		atomic.AddInt64(&c, 1)
	}
	wg.Done()
}

func main() {
	c = 0
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go worker2(&wg)
	}
	wg.Wait()
	fmt.Printf("c = %d", c)
}
