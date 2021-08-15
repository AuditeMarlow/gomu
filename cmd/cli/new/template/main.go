package template

var MainFNC = `package main

import (
	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"

	"{{.Dir}}/handler"
)

func main() {
	// Create function
	fnc := micro.NewFunction(
		micro.Name("{{lower .Alias}}"),
		micro.Version("latest"),
	)
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
	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"

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
		log.Fatal(err)
	}
}
`
