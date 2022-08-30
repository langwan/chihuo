package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const filename = "./preferences.json"

type Preferences struct {
	Name    string /* 配置别名 */
	Home    string /* 工作目录 */
	Workers int32  /* 工作单元数 */
}

func main() {
	p := Preferences{Name: "主配置", Home: "/home", Workers: 16}
	marshal, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("json marshal err: %v", err)
		return
	}
	err = ioutil.WriteFile(filename, marshal, 0644)
	if err != nil {
		fmt.Printf("write file err: %v", err)
		return
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file err: %v", err)
		return
	}

	rp := Preferences{}

	err = json.Unmarshal(data, &rp)
	if err != nil {
		fmt.Printf("json unmarshal err: %v", err)
		return
	}
	fmt.Printf("preferences is: %v", rp)
}
