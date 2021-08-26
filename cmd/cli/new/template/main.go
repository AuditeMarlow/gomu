package template

var MainFNC = `package main

import (
{{if .Jaeger }}	"{{.Dir}}/debug/trace/jaeger"
{{end}}	"{{.Dir}}/handler"

{{if .Jaeger }}	ot "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
{{end}}	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
)

var (
	service = "{{lower .Alias}}"
	version = "latest"
)

func main() {
{{if .Jaeger}}	// Create tracer
	tracer, closer, err := jaeger.NewTracer(
		jaeger.Name(service),
		jaeger.FromEnv(true),
		jaeger.GlobalTracer(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

{{end}}	// Create function
	fnc := micro.NewFunction(
		micro.Name(service),
		micro.Version(version),
{{if .Jaeger}}		micro.WrapCall(ot.NewCallWrapper(tracer)),
		micro.WrapClient(ot.NewClientWrapper(tracer)),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(tracer)),
{{end}}	)
	fnc.Init()

	// Handle function
	fnc.Handle(new(handler.{{title .Alias}}))

	// Run function
	if err := fnc.Run(); err != nil {
		log.Fatal(err)
	}
}
`

var MainSRV = `package main

import (
{{if .Jaeger }}	"{{.Dir}}/debug/trace/jaeger"
{{end}}	"{{.Dir}}/handler"
	pb "{{.Dir}}/proto"

{{if .Jaeger }}	ot "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
{{end}}	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
)

var (
	service = "{{lower .Alias}}"
	version = "latest"
)

func main() {
{{if .Jaeger}}	// Create tracer
	tracer, closer, err := jaeger.NewTracer(
		jaeger.Name(service),
		jaeger.FromEnv(true),
		jaeger.GlobalTracer(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

{{end}}	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
{{if .Jaeger}}		micro.WrapCall(ot.NewCallWrapper(tracer)),
		micro.WrapClient(ot.NewClientWrapper(tracer)),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer)),
		micro.WrapSubscriber(ot.NewSubscriberWrapper(tracer)),
{{end}}	)
	srv.Init()

	// Register handler
	pb.Register{{title .Alias}}Handler(srv.Server(), new(handler.{{title .Alias}}))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
`
