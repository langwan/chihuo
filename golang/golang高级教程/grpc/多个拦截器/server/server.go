package main

import (
	"app/rpc"
	"context"
	"github.com/rs/zerolog/log"
)

type Server struct {
}

func (s Server) Hello(ctx context.Context, request *rpc.Empty) (*rpc.HelloResponse, error) {
	log.Info().Msg("call Hello")
	resp := rpc.HelloResponse{Hello: "hello client."}
	return &resp, nil
}
