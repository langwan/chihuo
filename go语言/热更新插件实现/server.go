package main

import (
	"app/pb"
	"app/sdk"
	"io"
)

type Server struct {
}

func (s Server) Register(stream pb.Language_RegisterServer) error {
	var name string
	defer func() {
		if len(name) != 0 {
			if data, ok := plugins.LoadAndDelete(name); ok {
				f := data.(chan *sdk.CallContext)
				close(f)
			}
		}
	}()

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if req.Cmd == "register" {
			name = req.GetBody()
			if data, ok := plugins.LoadAndDelete(req.GetBody()); ok {
				f := data.(chan *sdk.CallContext)
				close(f)
			}
			fifo := make(chan *sdk.CallContext, 10)
			plugins.Store(req.GetBody(), fifo)
			go func() {
				for {
					ctx, ok := <-fifo
					if !ok {
						return
					} else {
						requests.Store(ctx.Id, ctx)
						stream.Send(&pb.ServerMessage{
							Id:   ctx.Id,
							Cmd:  ctx.Cmd,
							Body: ctx.Request,
						})
					}
				}
			}()
		} else {
			if data, ok := requests.Load(req.GetId()); ok {
				if cc, ok := data.(*sdk.CallContext); ok {
					cc.Response = req.GetBody()
					close(cc.Return)
				}
			}
		}
	}
}

//func (s Server) Register(request *pb.RegisterRequest, stream pb.Language_RegisterServer) error {
//	var fifo = make(chan string, 10)
//	if data, ok := plugins.Load(request.GetName()); ok {
//		f := data.(chan string)
//		close(f)
//	}
//	plugins.Store(request.GetName(), fifo)
//	for {
//		content, ok := <-fifo
//		if !ok {
//			return nil
//		}
//		stream.Send(&pb.Message{Content: content})
//		return nil
//	}
//}
