package main

import (
	"context"
	"os"
	"server/rpc"
)

type Server struct {
}

func (s Server) Exit(ctx context.Context, empty *rpc.Empty) (*rpc.Empty, error) {
	defer os.Exit(0)
	return &rpc.Empty{}, nil
}

func (s Server) Hello(ctx context.Context, request *rpc.Empty) (*rpc.HelloResponse, error) {
	resp := rpc.HelloResponse{Hello: "hello client."}
	return &resp, nil
}
