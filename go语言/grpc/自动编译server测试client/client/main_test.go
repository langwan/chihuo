package main

import (
	"bufio"
	"client/rpc"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func runServer(run chan bool, quit chan bool) (cmd *exec.Cmd) {
	ret := make(chan bool)

	go func() {
		cmd = exec.Command("go", "run", ".")
		cmd.Dir = "../server"
		stdout, err := cmd.StderrPipe()

		if err != nil {
			log.Fatal(err)
		}

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("server pid = %d\n", cmd.Process.Pid)
		//time.Sleep(10 * time.Second)
		ret <- true

		go func() {
			buf := bufio.NewReader(stdout)
			for {
				line, _, _ := buf.ReadLine()
				if index := strings.Index(string(line), "service start"); index != -1 {
					run <- true
				}
				if strings.TrimSpace(string(line)) != "" {
					fmt.Println(string(line))
				}
			}
		}()

		if err := cmd.Wait(); err != nil {
			quit <- false
		}
		fmt.Printf("pid = %d\n", cmd.Process.Pid)
		quit <- true
	}()
	<-ret
	return cmd
}

func TestHello(t *testing.T) {
	quit := make(chan bool)
	run := make(chan bool)
	runServer(run, quit)
	//cmd := runServer(run, quit)
	<-run
	fmt.Println("run")
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	ServerClient := rpc.NewServerClient(conn)
	helloResponse, err := ServerClient.Hello(context.Background(), &rpc.Empty{})
	if err != nil {
		fmt.Printf("err: %v", err)
		t.Error(err)
		return
	}

	log.Println(helloResponse, err)

	ServerClient.Exit(context.Background(), &rpc.Empty{})
	//cmd.Process.Kill()
	q := <-quit
	if q {
		t.Log("ok")
	} else {
		t.Fail()
	}
}
