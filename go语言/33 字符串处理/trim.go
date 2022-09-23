package main

import (
	"fmt"
	"strings"
)

func t1() {
	s1 := " chihuo@golang \n"
	s2 := strings.TrimSpace(s1)
	fmt.Printf("trim space '%s'\n", s2)
}
