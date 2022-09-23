package main

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"sync"
)

var sessions sync.Map

func main() {
	g := gin.Default()
	g.POST("/login", func(context *gin.Context) {
		token := uuid.NewV4().String()
		sessions.Store(token, token)
		context.JSON(200, struct {
			Token string `json:"token"`
		}{Token: token})
	})
	g.POST("/doing", func(context *gin.Context) {
		req := struct {
			Token string `json:"token"`
		}{}
		context.ShouldBind(&req)
		if _, ok := sessions.Load(req.Token); ok {
			context.JSON(200, struct {
				Message string `json:"message"`
			}{Message: "find"})
		} else {
			context.JSON(403, struct {
				Message string `json:"message"`
			}{Message: "not find"})
		}
	})
	g.Run(":8080")
}
