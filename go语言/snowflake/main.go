package main

import (
	"app/snowflake"
	"fmt"
)

func main() {
	sf, err := snowflake.New(1)
	if err != nil {
		fmt.Errorf("snowflake new error %v\n", err)
		return
	}
	id := sf.Gen()
	fmt.Printf("%64b id\n", -1^(-1<<63))
	fmt.Printf("%64b id\n", id)
	fmt.Printf("%64d id int64\n", id)

	id = sf.Gen()
	fmt.Printf("%64b next id\n", id)
	fmt.Printf("%64d next id int64\n", id)
}
