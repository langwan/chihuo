package main

import (
	"fmt"
	"sdk/model"
)

func main() {
	user := model.User{Name: "chihuo"}
	fmt.Printf("app user name is %s\n", user.Name)
}
