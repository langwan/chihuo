package main

import (
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var chunks = 10

	workers := 10

	jobs := make(chan int, chunks)

	wg := sync.WaitGroup{}
	wg.Add(chunks)

	for i := 0; i < workers; i++ {
		go worker(i, jobs, &wg)
	}

	scheduler(jobs, chunks)

	wg.Wait()
}
