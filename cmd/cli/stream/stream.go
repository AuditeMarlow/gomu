package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/urfave/cli/v2"
)

func NewCommand(alias string) *cli.Command {
	return &cli.Command{
		Name:   "stream",
		Usage:  "Create a service stream, e.g. " + alias + " stream greeter Say.HelloStream '{\"name\": \"John\"}'",
		Action: Run,
	}
}

func Run(ctx *cli.Context) error {
	args := ctx.Args().Slice()
	if len(args) < 2 {
		return cli.ShowSubcommandHelp(ctx)
	}

	service := args[0]
	endpoint := args[1]

	if len(service) == 0 {
		fmt.Println("must provide a service name")
		return nil
	}
	if len(endpoint) == 0 {
		fmt.Println("must provide an endpoint")
		return nil
	}

	req := strings.Join(args[2:], " ")
	if len(req) == 0 {
		req = `{}`
	}

	d := json.NewDecoder(strings.NewReader(req))
	d.UseNumber()

	var creq map[string]interface{}

	if err := d.Decode(&creq); err != nil {
		return err
	}

	srv := micro.NewService()
	srv.Init()
	c := srv.Client()

	var r interface{}
	request := c.NewRequest(service, endpoint, r, client.WithContentType("application/json"))

	stream, err := c.Stream(context.Background(), request)
	if err != nil {
		return err
	}
	if err := stream.Send(creq); err != nil {
		return err
	}
	for stream.Error() == nil {
		rsp := map[string]interface{}{}
		err := stream.Recv(&rsp)
		if err != nil {
			return err
		}
		b, err := json.Marshal(rsp)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	}
	if stream.Error() != nil {
		return err
	}
	if err := stream.Close(); err != nil {
		return err
	}

	return nil
}
