package main

import (
	"fmt"
	"sync"
)

type Number struct {
	Value int64
	Mutex sync.Mutex
}

func (n *Number) Add(val int64) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	n.Value = n.Value + val
}

func (n *Number) Get() int64 {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	return n.Value
}

func worker(number *Number, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		number.Add(1)
		//number.Value = number.Value + 1
	}
	wg.Done()
}

func main() {
	number := Number{Value: 0}
	wg := sync.WaitGroup{}
	workers := 10
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go worker(&number, &wg)
	}
	wg.Wait()
	fmt.Printf("number is %d", number.Get())
}
