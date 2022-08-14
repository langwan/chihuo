package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {

	for i := 0; i < 50; i++ {
		fmt.Printf("worker is %d\n", id)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	fmt.Println("执行...")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go worker(1, &wg)
	fmt.Println("等待...")
	wg.Wait()
	fmt.Println("结束")
}
