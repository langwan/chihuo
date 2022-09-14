package main

import (
	"bufio"
	"client/rpc"
	"context"
	"google.golang.org/grpc"
	"os/exec"
	"strings"
	"testing"
)

func server(dir string, t *testing.T, foo func()) {
	quit := make(chan bool)
	run := make(chan bool)

	go func() {
		cmd := exec.Command("go", "run", ".")
		cmd.Dir = dir
		stdout, err := cmd.StderrPipe()

		if err != nil {
			t.Fatal(err)
		}

		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
		//time.Sleep(1000 * time.Millisecond)
		//run <- true
		go func() {
			buf := bufio.NewReader(stdout)
			for {
				line, _, _ := buf.ReadLine()
				if index := strings.Index(string(line), "service start"); index != -1 {
					run <- true
				}
			}
		}()

		if err := cmd.Wait(); err != nil {
			quit <- false
		}

		quit <- true
	}()
	<-run
	foo()

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
		return
	}
	ServerClient := rpc.NewServerClient(conn)
	ServerClient.Exit(context.Background(), &rpc.Empty{})
	q := <-quit
	if q {
		t.Log("ok")
	} else {
		t.Fail()
	}
}

func TestHello(t *testing.T) {
	server("../server", t, func() {
		conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
		if err != nil {
			t.Fatal(err)
			return
		}

		ServerClient := rpc.NewServerClient(conn)
		helloResponse, err := ServerClient.Hello(context.Background(), &rpc.Empty{})

		if err != nil {
			t.Fatal(err)
			return
		}

		t.Log(helloResponse, err)
		
	})
}
