package main

import (
	"fmt"
	"github.com/langwan/go-jwt-hs256"
)

func main() {

	jwt.Secret = "123456"
	payload := struct {
		Name string `json:"name"`
	}{Name: "chihuo"}
	sign, err := jwt.Sign(payload)
	if err != nil {
		return
	}
	err = jwt.Verify(sign)
	if err != nil {
		return
	}
	fmt.Println("ok")
}
