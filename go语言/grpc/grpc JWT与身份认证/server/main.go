package main

import (
	"app/rpc"
	"context"
	"fmt"
	"github.com/langwan/go-jwt-hs256"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"strings"
)

func main() {

	jwt.Secret = "123456"

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))

	if err != nil {
		log.Fatal().Msg("启动grpc server失败")
		return
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptor()))
	rpc.RegisterAccountServer(grpcServer, Account{})
	rpc.RegisterServerServer(grpcServer, Server{})

	log.Info().Msg("service start")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Msg("启动grpc server失败")
	}
}

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		methods := strings.Split(info.FullMethod, "/")
		if methods[1] == "Account" {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, status.Errorf(codes.PermissionDenied, "permission denied")
			}
			token := ""
			if value, ok := md["token"]; ok {
				token = value[0]
				if len(token) == 0 {
					return nil, status.Errorf(codes.PermissionDenied, "permission denied")
				}
				fmt.Printf("server recv token %s\n", token)
			} else {
				return nil, status.Errorf(codes.PermissionDenied, "permission denied")
			}
			err = jwt.Verify(token)
			if err != nil {
				return nil, status.Errorf(codes.PermissionDenied, "permission denied")
			}

		}
		resp, err = handler(ctx, req)
		return resp, err
	}
}
