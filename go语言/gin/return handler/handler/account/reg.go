package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type RegResponse struct {
	Message string
}

func Reg(c *gin.Context) {
	var req RegRequest
	c.ShouldBindJSON(&req)
	c.JSON(http.StatusOK, "call reg")
}
