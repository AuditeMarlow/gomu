package stream

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/urfave/cli/v2"
)

func Client(ctx *cli.Context) error {
	args := ctx.Args().Slice()
	if len(args) < 2 {
		return cli.ShowSubcommandHelp(ctx)
	}

	service := args[0]
	endpoint := args[1]
	requests := args[2:]

	srv := micro.NewService()
	srv.Init()
	c := srv.Client()

	var r interface{}
	request := c.NewRequest(service, endpoint, r, client.WithContentType("application/json"))
	stream, err := c.Stream(ctx.Context, request)
	if err != nil {
		return err
	}

	for _, req := range requests {
		fmt.Printf("req = %+v\n", req)
		d := json.NewDecoder(strings.NewReader(req))
		d.UseNumber()

		var creq map[string]interface{}
		if err := d.Decode(&creq); err != nil {
			return err
		}

		if err := stream.Send(creq); err != nil {
			return err
		}

		time.Sleep(time.Millisecond * 250)
	}

	var rsp map[string]interface{}
	stream.Recv(rsp)
	fmt.Println(rsp)

	return nil
}
