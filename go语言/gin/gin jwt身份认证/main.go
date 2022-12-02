package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/langwan/go-jwt-hs256"
	"strings"
)

type AccountToken struct {
	Name string `json:"name"`
}

func main() {
	jwt.Secret = "123456"
	g := gin.Default()
	g.Use(func(c *gin.Context) {
		if c.Request.RequestURI == "/login" {
			return
		}
		if token, ok := c.Request.Header["Token"]; ok {
			err := jwt.Verify(token[0])
			if err != nil {
				c.AbortWithStatusJSON(403, "403 Forbidden")
				return
			}
		} else {
			c.AbortWithStatusJSON(403, "403 Forbidden")
		}
	})
	g.POST("/login", func(c *gin.Context) {
		authToken := AccountToken{Name: "chihuo"}
		sign, _ := jwt.Sign(authToken)
		c.JSON(200, struct {
			Token string `json:"token"`
		}{Token: sign})
	})
	g.POST("/doing", func(c *gin.Context) {
		at, err := getToken(c)
		if err != nil {
			c.AbortWithError(500, err)
		} else {
			c.JSON(200, at)
		}
	})
	g.Run(":8080")
}

func getToken(c *gin.Context) (at *AccountToken, err error) {
	if token, ok := c.Request.Header["Token"]; ok {
		ss := strings.Split(token[0], ".")
		at := &AccountToken{}
		payload, err := base64.RawURLEncoding.DecodeString(ss[1])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(payload, at)
		return at, nil
	} else {
		return nil, errors.New("auth token not find")
	}
}
