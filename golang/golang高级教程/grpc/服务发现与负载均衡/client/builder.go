package main

import "google.golang.org/grpc/resolver"

type ChihuoBuilder struct {
	addrs map[string][]string
}

func (b ChihuoBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &ChihuoResolver{}
	paths := b.addrs[target.URL.Path]
	addrs := make([]resolver.Address, len(paths))
	for i, s := range paths {
		addrs[i] = resolver.Address{Addr: s}
	}
	cc.UpdateState(resolver.State{Addresses: addrs})
	return r, nil
}

func (c ChihuoBuilder) Scheme() string {
	return "chihuo"
}
