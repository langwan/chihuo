package main

import (
	"github.com/client9/reopen"
	"github.com/rs/zerolog"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker1() {
	f, err := reopen.NewFileWriter("./worker.log")
	if err != nil {
		log.Fatalf("Unable to set output log: %s", err)
	}
	log.SetOutput(f)

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)
	go func() {
		for {
			<-sigusr1
			f.Reopen()
		}
	}()

	index := 0

	for {
		log.Printf("log %d", index)
		index++
		time.Sleep(200 * time.Millisecond)
	}

}

func worker2() {
	f, err := reopen.NewFileWriter("./worker.log")
	if err != nil {
		log.Fatalf("Unable to set output log: %s", err)
	}
	logger := zerolog.New(f)

	sighup := make(chan os.Signal, 1)
	signal.Notify(sighup, syscall.SIGUSR1)
	go func() {
		for {
			<-sighup
			f.Reopen()
		}
	}()

	index := 0

	for {
		logger.Info().Msgf("log %d", index)
		index++
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	worker2()
}
