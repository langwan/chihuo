package main

import (
	"app/pb"
	"context"
	"fmt"
	"github.com/langwan/langgo/core/rpc"
	"google.golang.org/grpc"
	"io"
)

func main() {
	conn, err := rpc.NewClient(nil, "127.0.0.1:8000", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	ServerClient := pb.NewLanguageClient(conn)

	stream, err := ServerClient.Register(context.Background())

	err = stream.Send(&pb.PluginMessage{
		Id:   "",
		Cmd:  "register",
		Body: "en",
	})
	if err != nil {
		panic(err)
		return
	}

	for {
		recv, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("server say bye")
			return
		} else if err != nil {
			fmt.Printf("recv %v", err)
			continue
		} else {
			stream.Send(&pb.PluginMessage{
				Id:   recv.Id,
				Cmd:  "",
				Body: "hello chihuo 111111",
			})
		}
	}
}
