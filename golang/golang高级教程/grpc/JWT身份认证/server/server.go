package main

import (
	"app/rpc"
	"context"
	"github.com/langwan/go-jwt-hs256"
)

type Server struct {
}

type AccountToken struct {
	Id   string
	Name string
}

func (s Server) Login(ctx context.Context, request *rpc.LoginRequest) (*rpc.LoginResponse, error) {
	payload := AccountToken{
		Id:   "1000",
		Name: "chihuo",
	}
	sign, err := jwt.Sign(payload)
	if err != nil {
		return nil, err
	}
	return &rpc.LoginResponse{Token: sign}, nil
}

func (s Server) Hello(ctx context.Context, request *rpc.Empty) (*rpc.HelloResponse, error) {
	resp := rpc.HelloResponse{Hello: "hello client."}
	return &resp, nil
}
