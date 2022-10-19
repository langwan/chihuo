package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadFileByHttp(t *testing.T) {
	fp := "./samples/1.mp4"
	name := filepath.Base(fp)
	http.HandleFunc("/"+name, func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, fp)
	})
	err := http.ListenAndServe(":8100", nil)
	assert.NoError(t, err)
}

func TestDownloadFileByGin(t *testing.T) {
	fp := "./samples/1.mp4"
	name := filepath.Base(fp)
	g := gin.Default()

	g.GET("/"+name, func(context *gin.Context) {
		http.ServeFile(context.Writer, context.Request, fp)
	})
	g.StaticFile("/sf_"+name, fp)

	err := g.Run(":8100")
	err = http.ListenAndServe(":8100", nil)
	assert.NoError(t, err)
}

func TestDownloadFileStreamByHttp(t *testing.T) {
	fp := "./samples/1.mp4"
	name := filepath.Base(fp)
	http.HandleFunc("/"+name, func(writer http.ResponseWriter, request *http.Request) {
		f, _ := os.Open(fp)
		defer f.Close()
		info, _ := f.Stat()
		http.ServeContent(writer, request, name, info.ModTime(), f)
	})
	err := http.ListenAndServe(":8100", nil)
	assert.NoError(t, err)
}

func TestDownloadDirByGin(t *testing.T) {
	dir := "./samples"
	g := gin.Default()
	g.Static("/samples", dir)
	err := g.Run(":8100")
	err = http.ListenAndServe(":8100", nil)
	assert.NoError(t, err)
}

type MyStreamer struct {
	File *os.File
}

func (m *MyStreamer) Read(p []byte) (n int, err error) {
	n, err = m.File.Read(p)
	fmt.Printf("my streamer reads %d, err = %v\n", n, err)
	return n, err
}

func (m *MyStreamer) Seek(offset int64, whence int) (int64, error) {
	n, err := m.File.Seek(offset, whence)
	fmt.Printf("my streamer seek offset %d\n", n)
	return n, err
}

func TestDownloadFileMyStreamByHttp(t *testing.T) {
	fp := "./samples/1.mp4"
	name := filepath.Base(fp)

	http.HandleFunc("/"+name, func(writer http.ResponseWriter, request *http.Request) {
		f, _ := os.Open(fp)
		defer f.Close()
		myStreamer := MyStreamer{File: f}
		info, _ := f.Stat()
		http.ServeContent(writer, request, name, info.ModTime(), &myStreamer)
	})
	err := http.ListenAndServe(":8100", nil)
	assert.NoError(t, err)
}
