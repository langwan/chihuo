package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8100")
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte("GET /profile HTTP/1.1\r\nHost: 127.0.0.1\r\n\r\n"))
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	rn, err := conn.Read(buf)
	if err != nil {
		return
	}
	fmt.Println(string(buf[:rn]))
}
