package main

import (
	"app/pb"
	"app/sdk"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
	"github.com/langwan/langgo/core/rpc"
	"github.com/langwan/langgo/helpers/uuid"
	"os"
	"sync"
	"syscall"
	"time"
)

var plugins sync.Map
var requests sync.Map

func call(pluginName string, cmd string, timeout time.Duration, req interface{}) (interface{}, error) {
	j, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	ctx := sdk.CallContext{
		Id:       uuid.String2(),
		Cmd:      cmd,
		Return:   make(chan bool),
		Request:  string(j),
		Response: "",
	}
	if data, ok := plugins.Load(pluginName); ok {
		fifo := data.(chan *sdk.CallContext)
		fifo <- &ctx
		select {
		case <-ctx.Return:
			return ctx.Response, nil
		case <-time.After(timeout):
		}

	}
	return nil, errors.New("timeout")
}

func main() {
	langgo.Run()
	core.SignalHandle(&core.SignalHandler{
		Sig: syscall.SIGINT,
		F: func() {
			os.Exit(int(syscall.SIGINT))
		},
	})
	go func() {

		for {

			plugins.Range(func(key, value any) bool {
				response, err := call(key.(string), "hello world", 3*time.Second, nil)
				if err != nil {
					fmt.Printf("call err %v\n", err)
				}
				fmt.Printf("response %v\n", response)
				return true
			})
			time.Sleep(time.Second)
		}

	}()
	cg := rpc.NewServer(nil)
	gs, err := cg.Server()
	pb.RegisterLanguageServer(gs, &Server{})
	err = cg.Run("127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

}
