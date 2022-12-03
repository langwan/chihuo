package main

import (
	"github.com/gin-gonic/gin"
)

func ginServer() {
	g := gin.New()
	g.StaticFile("/profile", "./html/profile.html")
	g.StaticFile("/home", "./html/home.html")
	g.Run(":8100")
}
