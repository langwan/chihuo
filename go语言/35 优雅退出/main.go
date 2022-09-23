package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("main start")

	defer func() {
		fmt.Println("bye main from defer")
		quit()
	}()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for s := range sig {
			switch s {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGHUP:
				quit()
				if i, ok := s.(syscall.Signal); ok {
					os.Exit(int(i))
				} else {
					os.Exit(0)
				}
			}
		}
	}()

	wait := make(chan bool)
	go func() {
		for {
			time.Sleep(5000 * time.Millisecond)
			close(wait)
		}
	}()
	<-wait

	fmt.Println("main end")
}

func quit() {
	fmt.Println("\n成功退出")
}
