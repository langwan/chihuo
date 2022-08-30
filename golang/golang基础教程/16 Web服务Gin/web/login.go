package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Body    string `json:"body"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	req := LoginRequest{}
	c.ShouldBind(&req)
	if req.Name == "chihuo" && req.Password == "password" {
		resp := WebResponse{Code: 0, Message: "ok", Body: "1000000"}
		c.JSON(http.StatusOK, resp)
	} else {
		resp := WebResponse{Code: -1, Message: "login failed", Body: ""}
		c.AbortWithStatusJSON(http.StatusOK, resp)
	}
}
