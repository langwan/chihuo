package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func HttpLog() gin.HandlerFunc {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("HttpLog init")
	return func(c *gin.Context) {
		log.Info().Str("uri", c.Request.RequestURI).Send()
	}
}

func HttpLogFunc(c *gin.Context) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("HttpLog init")
	log.Info().Str("uri", c.Request.RequestURI).Send()
}

func main() {
	g := gin.Default()
	g.Use(HttpLogFunc)
	g.GET("/login", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	g.Run(":8080")
}
