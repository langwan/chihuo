package order

import (
	"github.com/gin-gonic/gin"
)

type GetRequest struct {
	Id string `json:"id"`
}
type GetResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Get(c *gin.Context, req *GetRequest) (*GetResponse, error) {
	return &GetResponse{
		Id:   "1",
		Name: "order",
	}, nil
}
