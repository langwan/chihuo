package main

import (
	"fmt"
	"time"
)

func main() {
	wait := make(chan bool)
	go func() {
		for {
			onlinePool.Submit(func() {
				fmt.Println("online send")
				time.Sleep(500 * time.Millisecond)
			})
		}
	}()
	go func() {
		for {

			offlinePool.Submit(func() {
				fmt.Println("offline send")
				time.Sleep(500 * time.Millisecond)
			})
		}
	}()
	<-wait
}
