package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type handler func(conn net.Conn)

var handlers = make(map[string]handler)

func server() {
	get("/profile", func(conn net.Conn) {
		data, err := os.ReadFile("./html/profile.html")
		if err != nil {
			return
		}
		resp(conn, data)
	})
	get("/home", func(conn net.Conn) {
		data, err := os.ReadFile("./html/home.html")
		if err != nil {
			return
		}
		resp(conn, data)
	})
	listen, err := net.Listen("tcp", ":8100")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}

		message = strings.TrimSpace(message)
		if strings.HasPrefix(message, "GET ") {
			sp := strings.Split(message, " ")
			if handler, ok := handlers[sp[1]]; ok {
				handler(conn)
				conn.Close()
			}
		}

	}

}

func get(uri string, handler handler) {
	handlers[uri] = handler
}

func resp(conn net.Conn, body []byte) {
	content := fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n%s", string(body))
	conn.Write([]byte(content))
}
