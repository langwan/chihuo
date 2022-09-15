package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
)

const etcdUrl = "http://localhost:2379"
const serviceName = "chihuo/server"
const ttl = 10

var etcdClient *clientv3.Client

func etcdRegister(addr string) error {
	log.Printf("etcdRegister %s\b", addr)
	etcdClient, err := clientv3.NewFromURL(etcdUrl)

	if err != nil {
		return err
	}

	em, err := endpoints.NewManager(etcdClient, serviceName)
	if err != nil {
		return err
	}

	lease, _ := etcdClient.Grant(context.TODO(), ttl)

	err = em.AddEndpoint(context.TODO(), fmt.Sprintf("%s/%s", serviceName, addr), endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	//etcdClient.KeepAlive(context.TODO(), lease.ID)
	alive, err := etcdClient.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return err
	}

	go func() {
		for {
			<-alive
			fmt.Println("etcd server keep alive")
		}
	}()

	return nil
}

func etcdUnRegister(addr string) error {
	log.Printf("etcdUnRegister %s\b", addr)
	if etcdClient != nil {
		em, err := endpoints.NewManager(etcdClient, serviceName)
		if err != nil {
			return err
		}
		err = em.DeleteEndpoint(context.TODO(), fmt.Sprintf("%s/%s", serviceName, addr))
		if err != nil {
			return err
		}
		return err
	}

	return nil
}
