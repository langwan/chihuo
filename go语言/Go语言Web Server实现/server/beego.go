package main

import (
	"github.com/beego/beego/v2/server/web"
	beecontext "github.com/beego/beego/v2/server/web/context"
	"os"
)

func beegoServer() {
	web.Get("/profile", func(ctx *beecontext.Context) {
		data, err := os.ReadFile("./html/profile.html")
		if err != nil {
			return
		}
		ctx.WriteString(string(data))
	})
	web.Get("/home", func(ctx *beecontext.Context) {
		data, err := os.ReadFile("./html/home.html")
		if err != nil {
			return
		}
		ctx.WriteString(string(data))
	})
	web.Run(":8100")
}
