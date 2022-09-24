package main

import "github.com/panjf2000/ants"

var onlinePool *ants.Pool
var offlinePool *ants.Pool

func init() {

	onlinePool, _ = ants.NewPool(10, ants.WithNonblocking(true))

	offlinePool, _ = ants.NewPool(2, ants.WithNonblocking(true))

}
