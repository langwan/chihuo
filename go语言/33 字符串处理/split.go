package main

import (
	"fmt"
	"strings"
)

func s1() {
	s1 := "chihuo@golang"
	arr := strings.Split(s1, "@")
	fmt.Printf("arr is %v\n", arr)
}
