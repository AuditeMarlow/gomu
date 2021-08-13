package call

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/urfave/cli/v2"
)

func Run(ctx *cli.Context) error {
	args := ctx.Args().Slice()
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

	request := c.NewRequest(service, endpoint, creq, client.WithContentType("application/json"))
	response := map[string]string{}

	if err := c.Call(context.Background(), request, &response); err != nil {
		return err
	}

	b, err := json.Marshal(response)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
