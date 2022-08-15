package main

import "fmt"

type Item interface {
	Set(val string)
	Get() string
}

type CacheItem struct {
	Value string
}

func (ci *CacheItem) Set(val string) {
	ci.Value = val
}

func (ci *CacheItem) Get() string {
	return ci.Value
}

func main() {
	ci := CacheItem{}
	ci.Set("chihuo")
	fmt.Printf("ci value is %s", ci.Get())
}
