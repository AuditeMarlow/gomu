package template

var HandlerFNC = `package handler

import (
	"context"

	log "github.com/asim/go-micro/v3/logger"

	pb "{{.Dir}}/proto"
)

type {{title .Alias}} struct{}

func (e *{{title .Alias}}) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Infof("Received {{title .Alias}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
`

var HandlerSRV = `package handler

import (
	"context"
	"time"

	log "github.com/asim/go-micro/v3/logger"

	pb "{{.Dir}}/proto"
)

type {{title .Alias}} struct{}

func (e *{{title .Alias}}) Call(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Infof("Received {{title .Alias}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *{{title .Alias}}) Stream(ctx context.Context, req *pb.StreamRequest, stream pb.{{title .Alias}}_StreamStream) error {
	log.Infof("Received {{title .Alias}}.Stream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.StreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}
`
