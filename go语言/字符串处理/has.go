package main

import (
	"fmt"
	"strings"
)

func h1() {
	s1 := "chihuo@golang"
	if strings.HasPrefix(s1, "chihuo") {
		fmt.Printf("%s has prefix chihuo\n", s1)
	}
	if strings.HasSuffix(s1, "golang") {
		fmt.Printf("%s has suffix golang\n", s1)
	}
}
