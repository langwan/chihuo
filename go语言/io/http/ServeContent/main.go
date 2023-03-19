package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	enableBuffer = flag.Bool("buffer", false, "enable buffer")
)

func main() {
	flag.Parse()
	src := "https://chihuo-test.oss-cn-hangzhou.aliyuncs.com/6543f38b-6cb3-4d66-9974-21cbd6301018.mp4"
	name := "sample.mp4"
	addr := ":8100"
	http.HandleFunc("/"+name, func(writer http.ResponseWriter, request *http.Request) {
		var reader io.ReadSeeker
		reader = NewHttpReadSeeker(http.DefaultClient, src)
		buf := make([]byte, 1024*1024*2)
		reader = NewBufferedReadSeeker(reader, buf)
		http.ServeContent(writer, request, name, time.Now(), reader)
	})
	fmt.Printf("http server %s\n", addr)
	fmt.Printf("http url http://127.0.0.1%s/%s\n", addr, name)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
