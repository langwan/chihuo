package main

import (
	"fmt"
	"io/ioutil"
)

func err1() (int, error) {
	data, err := ioutil.ReadFile("test.txt")

	if err != nil {
		return 0, err
	}

	fmt.Printf("read file body len = %d\n", len(data))

	return len(data), nil
}

func err2() {
	_, err := err1()
	if err != nil {
		fmt.Printf("err is \"%v\"", err)
		return
	}
	fmt.Printf("err2 ok.")
}

func err3() {
	l, _ := err1()
	fmt.Printf("data len is %d", l)
}

func main() {
	err3()
}
