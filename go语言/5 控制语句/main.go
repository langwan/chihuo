package main

import "fmt"

var items []int
var m = make(map[int]int)

func main() {

	for i := 0; i < 10; i++ {
		items = append(items, i)
		m[i] = i
	}

	fmt.Println("items:")
	for _, v := range items {
		fmt.Printf("\titem %d\n", v)
	}
	fmt.Println("\nm:")
	for k, v := range m {
		fmt.Printf("\tk = %d, v = %d\n", k, v)
	}

}
