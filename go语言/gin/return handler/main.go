package main

import (
	"app/h"
	"app/handler/account"
	"app/handler/order"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	{
		h.POST(g, "acc/login", account.Login)
		g.POST("acc/reg", account.Reg)
	}
	{
		h.POST(g, "order/list", order.List)
		h.POST(g, "order/get", order.Get)
	}

	g.Run(":8000")
}
