package template

var HandlerFNC = `package handler

import (
	"context"

	log "github.com/asim/go-micro/v3/logger"

	pb "{{.Dir}}/proto"
)

type {{title .Alias}} struct{}

func (e *{{title .Alias}}) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received {{title .Alias}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
`

var HandlerSRV = `package handler

import (
	"context"
	"io"
	"time"

	log "github.com/asim/go-micro/v3/logger"

	pb "{{.Dir}}/proto"
)

type {{title .Alias}} struct{}

func (e *{{title .Alias}}) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received {{title .Alias}}.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *{{title .Alias}}) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.{{title .Alias}}_ServerStreamStream) error {
	log.Infof("Received {{title .Alias}}.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *{{title .Alias}}) BidiStream(ctx context.Context, stream pb.{{title .Alias}}_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
`
