package services

import (
	"fmt"
	"sort"

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

	var services []string
	for _, srv := range srvs {
		services = append(services, srv.Name)
	}

	sort.Strings(services)
	for _, srv := range services {
		fmt.Println(srv)
	}

	return nil
}
