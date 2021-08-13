package template

var Main = `package main

import (
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"

	"{{.Dir}}/handler"
	pb "{{.Dir}}/proto"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("{{lower .Alias}}"),
		micro.Version("latest"),
	)
	srv.Init()

	// Register handler
	pb.Register{{title .Alias}}Handler(srv.Server(), new(handler.{{title .Alias}}))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
`
