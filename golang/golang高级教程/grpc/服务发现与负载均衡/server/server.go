package main

import (
	"app/rpc"
	"context"
	"fmt"
)

type Server struct {
}

func (s Server) Hello(ctx context.Context, request *rpc.Empty) (*rpc.HelloResponse, error) {
	resp := rpc.HelloResponse{Hello: "hello client."}
	return &resp, nil
}

func (s Server) Register(ctx context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	resp := rpc.RegisterResponse{}
	resp.Uid = fmt.Sprintf("%s.%s", request.GetName(), request.GetPassword())
	return &resp, nil
}
