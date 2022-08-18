package main

import (
	"app/web"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", web.Login)

	r.Run()
}
