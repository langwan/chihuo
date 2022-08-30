package main

import (
	"app/rpc"
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"runtime/debug"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))

	if err != nil {
		log.Fatal().Msg("启动grpc server失败")
		return
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor()))

	rpc.RegisterServerServer(grpcServer, Server{})

	log.Info().Msg("service start")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Msg("启动grpc server失败")
	}
}

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		st := time.Now()
		defer func() {
			if recoverError := recover(); recoverError != nil {
				log.Error().Str("method", info.FullMethod).Interface("recover", recoverError).Bytes("stack", debug.Stack()).Interface("req", req).Interface("resp", resp).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			} else {
				log.Info().Str("method", info.FullMethod).Interface("req", req).Interface("resp", resp).Err(err).TimeDiff("runtime", time.Now(), st).Send()
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}
