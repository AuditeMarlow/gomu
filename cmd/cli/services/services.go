package services

import (
	"fmt"

	"github.com/asim/go-micro/v3"
	"github.com/urfave/cli/v2"
)

func NewCommand(alias string) *cli.Command {
	return &cli.Command{
		Name:   "services",
		Usage:  "List services in the registry",
		Action: Run,
	}
}

func Run(ctx *cli.Context) error {
	srvs, err := micro.NewService().Options().Registry.ListServices()
	if err != nil {
		return err
	}

	for _, srv := range srvs {
		fmt.Println(srv.Name)
	}

	return nil
}
