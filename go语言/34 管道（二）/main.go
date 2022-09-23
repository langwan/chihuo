package main

import (
	"fmt"
	"time"
)

func worker1(id int, in chan bool, out chan bool) {
	fmt.Printf("worker %d start\n", id)
	<-in
	fmt.Printf("worker %d quit\n", id)
	out <- true
}

func worker2(id int, in <-chan bool, out chan<- bool) {
	fmt.Printf("worker %d start\n", id)
	<-in
	fmt.Printf("worker %d quit\n", id)
	out <- true
}

func worker3(id int, in <-chan bool, out chan<- bool) {
	fmt.Printf("worker %d start\n", id)
	for {
		fmt.Printf("worker %d doing\n", id)
		time.Sleep(200 * time.Millisecond)
		select {
		case <-in:
			out <- true
			return
		default:
		}
	}
}

/*
chan 读写
chan<-只写
<-chan 只读
*/
func main() {
	in := make(chan bool)
	out := make(chan bool)
	workers := 3

	for i := 0; i < workers; i++ {
		go worker3(i, in, out)
	}

	go func() {
		time.Sleep(time.Second)
		//close(in)
		for i := 0; i < workers; i++ {
			in <- true
		}
	}()
	count := 0
	for count < workers {
		<-out
		count++
	}

	fmt.Println("ok")
}
