package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Printf("start\n")
	sig := make(chan os.Signal)
	die := make(chan bool)
	signal.Notify(sig, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for {
			s := <-sig
			fmt.Printf("recv sig %d\n", s)
			if s == syscall.SIGUSR1 {
				die <- true
			}
		}
	}()
	<-die
	fmt.Printf("exit\n")
}
