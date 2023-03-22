package h

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context, body any) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "", "body": body})
}

func Fail(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
}

func FailMessage(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": message})
}

func FailCode(c *gin.Context, code int, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"code": code, "message": err})
}

func Forbidden(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1, "message": "无权限访问"})
}
