package main

import "context"

type creds struct {
}

var token string

func (c creds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"token": token}, nil
}

func (c creds) RequireTransportSecurity() bool {
	return false
}
