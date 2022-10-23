package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	src, _ := os.Open("samples/1.mp4")
	dst, _ := os.Create("samples/2.mp4")
	nw, _ := io.Copy(dst, src)
	fmt.Printf("nw = %d\n", nw)
}
