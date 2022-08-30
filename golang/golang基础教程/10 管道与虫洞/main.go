package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int)
var sum = 0

func worker(wg *sync.WaitGroup) {
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				wg.Done()
				return
			}
			sum = sum + num
		}
	}
}

func producer() {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go worker(&wg)
	go producer()
	wg.Wait()
	fmt.Printf("sum is %d", sum)
}
