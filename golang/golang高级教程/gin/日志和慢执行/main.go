package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

// 全部默认(生产环境勿用 仅限于开发)
func worker1() {

	g := gin.Default()

	g.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	g.GET("/panic", func(c *gin.Context) {
		panic("panic")
	})

	g.Run(":8080")
}

// 日志(包含访问和异常)写入到文件里
func worker2() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = f

	g := gin.New()
	g.Use(gin.Logger(), gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {

		msg := fmt.Sprintf("panic %s\n", debug.Stack())
		gin.DefaultWriter.Write([]byte(msg))

		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	g.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	g.GET("/panic", func(c *gin.Context) {
		panic("panic")
	})

	g.Run(":8080")
}

// 慢执行日志
func worker3() {

	f, _ := os.Create("custom.log")
	gin.DefaultWriter = f

	g := gin.New()
	g.Use(func(c *gin.Context) {
		t := time.Now()

		// before request

		c.Next()

		// after request
		latency := time.Since(t)

		// access the status we are sending
		status := c.Writer.Status()
		if latency > 20000 {
			s := fmt.Sprintf("slow %s %s %s %d\n", c.Request.RequestURI, t, latency, status)
			gin.DefaultWriter.Write([]byte(s))
		} else {
			s := fmt.Sprintf("%s %s %s %d\n", c.Request.RequestURI, t, latency, status)
			gin.DefaultWriter.Write([]byte(s))
		}

	}, gin.Recovery())

	g.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

	g.GET("/panic", func(c *gin.Context) {
		panic("panic")
	})

	g.Run(":8080")
}

func main() {
	worker3()
}
