package main

import (
	"app/rpc"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	certificate, err := tls.LoadX509KeyPair("../tls/server.crt", "../tls/server.key")
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../tls/ca.crt")
	if err != nil {
		panic(err)

	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		panic("AppendCertsFromPEM failed")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))

	if err != nil {
		log.Fatalf("启动grpc server失败")
		return
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	rpc.RegisterServerServer(grpcServer, Server{})

	log.Println("service start")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("启动grpc server失败")
	}
}
