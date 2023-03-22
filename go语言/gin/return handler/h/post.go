package h

import (
	"github.com/gin-gonic/gin"
	"io"
	"reflect"
)

type Empty struct {
}

func POST(g *gin.Engine, relativePath string, handler any) {
	g.POST(relativePath, func(c *gin.Context) {
		method := reflect.ValueOf(handler)
		parameter := method.Type().In(1)
		req := reflect.New(parameter.Elem()).Interface()
		err := c.ShouldBindJSON(req)
		if err != io.EOF && err != nil {
			Fail(c, err)
			return
		}
		in := make([]reflect.Value, 0)
		in = append(in, reflect.ValueOf(c))
		in = append(in, reflect.ValueOf(req))
		call := method.Call(in)
		if !call[1].IsNil() {
			callErr := call[1].Interface().(error)
			Fail(c, callErr)
			return
		}
		if !call[0].IsNil() {
			Ok(c, call[0].Interface())
		}
	})
}
