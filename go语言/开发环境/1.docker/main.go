package main

import (
	"flag"
	"fmt"
)

var (
	who = flag.String("who", "langwan", "who's name")
)

func main() {
	flag.Parse()
	fmt.Printf("hello %s\n", *who)
}
