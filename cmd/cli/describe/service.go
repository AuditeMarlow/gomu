package describe

import (
	"encoding/json"
	"fmt"

	"github.com/asim/go-micro/v3"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func Service(ctx *cli.Context) error {
	args := ctx.Args().Slice()
	if len(args) < 1 {
		return cli.ShowSubcommandHelp(ctx)
	}
	if ctx.String("format") != "json" && ctx.String("format") != "yaml" {
		return cli.ShowSubcommandHelp(ctx)
	}

	srvs, err := micro.NewService().Options().Registry.GetService(args[0])
	if err != nil {
		return err
	}
	if len(srvs) == 0 {
		return fmt.Errorf("service %s not found", args[0])
	}

	for _, srv := range srvs {
		var b []byte
		var err error
		if ctx.String("format") == "json" {
			b, err = json.MarshalIndent(srv, "", "  ")
		} else if ctx.String("format") == "yaml" {
			b, err = yaml.Marshal(srv)
		}
		if err != nil {
			return err
		}
		fmt.Println(string(b))
	}

	return nil
}
