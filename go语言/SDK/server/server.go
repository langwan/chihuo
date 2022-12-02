package main

import (
	"fmt"
	"sdk/model"
)

func main() {
	user := model.User{Name: "chihuo"}
	fmt.Printf("server user name is %s\n", user.Name)
}
