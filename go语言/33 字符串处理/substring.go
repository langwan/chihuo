package main

import "fmt"

func ss1() {
	s1 := "chihuo@golang"
	s2 := s1[6:len(s1)]
	fmt.Printf("sub string is %s\n", s2)
}
