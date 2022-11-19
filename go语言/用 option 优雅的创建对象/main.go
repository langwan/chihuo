package main

import (
	"app/lib/hello"
	"fmt"
	"time"
)

func main() {
	h := hello.New()
	fmt.Println(h)

	h = hello.New(hello.WithConfig("./config.yml"))
	fmt.Println(h)

	h = hello.New(hello.WithConfig("./config.yml"), hello.WithName("chihuo"), hello.WithVersion("1.0.0"))
	fmt.Println(h)

	h = &hello.Instance{
		Message:    "hello",
		Name:       "chihuo",
		LatestTime: time.Now(),
		Version:    "1.0.0",
	}
	fmt.Println(h)
}
