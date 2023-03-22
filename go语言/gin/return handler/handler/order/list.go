package order

import (
	"app/h"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListResponse struct {
	Items []string `json:"items"`
}

func List(c *gin.Context, req *h.Empty) (*ListResponse, error) {
	c.JSON(http.StatusOK, gin.H{"message": "gin output c.json"})
	return nil, nil
}
