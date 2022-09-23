package main

import (
	"fmt"
	"github.com/langwan/go-jwt-hs256"
)

func main() {

	jwt.Secret = "123456"

	payload := struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}{Id: "1000", Name: "chihuo"}

	sign, err := jwt.Sign(payload)
	if err != nil {
		fmt.Printf("err %v\n", err)
		return
	}

	fmt.Printf("sign is %s\n", sign)

	err = jwt.Verify(sign)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("verify ok")
}
