package main

import (
	"app/rpc"
	"context"
	"encoding/base64"
	"encoding/json"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type Account struct {
}

func getAccountToken(ctx context.Context) (at *AccountToken, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Aborted, "token error.")
	}
	token := ""
	if value, ok := md["token"]; ok {
		token = value[0]
		if len(token) == 0 {
			return nil, status.Errorf(codes.Aborted, "token error.")
		}
	} else {
		return nil, status.Errorf(codes.Aborted, "token error.")
	}
	ss := strings.Split(token, ".")
	at = &AccountToken{}
	payload, err := base64.RawURLEncoding.DecodeString(ss[1])
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "token error.")
	}
	err = json.Unmarshal(payload, at)
	return at, nil
}

func (a Account) Doing(ctx context.Context, empty *rpc.Empty) (*rpc.DoingResponse, error) {
	at, err := getAccountToken(ctx)
	if err != nil {
		return &rpc.DoingResponse{}, err
	}
	return &rpc.DoingResponse{Name: at.Name}, nil
}
