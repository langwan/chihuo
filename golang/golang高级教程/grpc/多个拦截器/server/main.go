package main

import (
	"app/rpc"
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog"
	"os"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))

	if err != nil {
		log.Fatal().Msg("启动grpc server失败")
		return
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(LogUnaryServerInterceptor(), AuthUnaryServerInterceptor())))
	//grpcServer := grpc.NewServer(grpc.UnaryInterceptor(LogUnaryServerInterceptor()), grpc.UnaryInterceptor(LogUnaryServerInterceptor()))
	rpc.RegisterServerServer(grpcServer, Server{})

	log.Info().Msg("service start")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Msg("启动grpc server失败")
	}
}

func LogUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Info().Str("method", info.FullMethod).Msg("LogUnaryServerInterceptor")
		resp, err = handler(ctx, req)
		return resp, err
	}
}

func AuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Info().Str("method", info.FullMethod).Msg("AuthUnaryServerInterceptor")
		resp, err = handler(ctx, req)
		return resp, err
	}
}
