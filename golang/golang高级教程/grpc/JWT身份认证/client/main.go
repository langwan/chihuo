package main

import (
	"client/rpc"
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(creds{}))

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	serverClient := rpc.NewServerClient(conn)

	helloResponse, err := serverClient.Hello(context.Background(), &rpc.Empty{})
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	log.Println(helloResponse, err)

	loginResponse, err := serverClient.Login(context.Background(), &rpc.LoginRequest{Name: "chihuo", Password: "123456"})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	token = loginResponse.GetToken()

	accountClient := rpc.NewAccountClient(conn)

	doingResponse, err := accountClient.Doing(context.Background(), &rpc.Empty{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(doingResponse.GetName())
}
