package main

import (
	"app/app"
	"encoding/json"
	"fmt"
)

func main() {

	mod := app.App{Name: "chihuo", Version: 2}

	fmt.Printf("mod name = %s ", mod.Name)

	jstr, _ := json.Marshal(mod)

	fmt.Printf("json string = %s", jstr)
}
