package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(workerId int, jobs <-chan int, wg *sync.WaitGroup) {
	for jobId := range jobs {
		fmt.Printf("worker %d, job %d\n", workerId, jobId)
		time.Sleep(time.Duration(rand.Intn(40)) * time.Millisecond)
		wg.Done()
	}
}

func scheduler(jobs chan int, len int) {
	for i := 0; i < len; i++ {
		jobs <- i
	}
}
