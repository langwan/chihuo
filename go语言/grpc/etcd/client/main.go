package main

import (
	"client/rpc"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	"google.golang.org/grpc"
)

const etcdUrl = "http://localhost:2379"
const serviceName = "chihuo/server"

func main() {
	//bd := &ChihuoBuilder{addrs: map[string][]string{"/api": []string{"localhost:8001", "localhost:8002", "localhost:8003"}}}
	//resolver.Register(bd)
	etcdClient, err := clientv3.NewFromURL(etcdUrl)
	if err != nil {
		panic(err)
	}
	etcdResolver, err := resolver.NewBuilder(etcdClient)

	conn, err := grpc.Dial(fmt.Sprintf("etcd:///%s", serviceName), grpc.WithResolvers(etcdResolver), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))

	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	ServerClient := rpc.NewServerClient(conn)

	for {
		helloRespone, err := ServerClient.Hello(context.Background(), &rpc.Empty{})
		if err != nil {
			fmt.Printf("err: %v", err)
			return
		}

		log.Println(helloRespone, err)
		time.Sleep(500 * time.Millisecond)
	}

}
