package main

import (
	"client/rpc"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(), grpc.w)

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	ServerClient := rpc.NewServerClient(conn)

	helloRespone, err := ServerClient.Hello(context.Background(), &rpc.Empty{})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(helloRespone, err)

	registerResponse, err := ServerClient.Register(context.Background(), &rpc.RegisterRequest{Name: "chihuo", Password: "123456"})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	_, err = ServerClient.PanicMethod(context.Background(), &rpc.Empty{})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(registerResponse, err)

}
