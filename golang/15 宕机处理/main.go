package main

import (
	"fmt"
	"time"
)

func worker() {
	panic("worker is destroy.")
}

func processing() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	worker()
}

func main() {
	for {
		processing()
		time.Sleep(time.Second)
	}
}
