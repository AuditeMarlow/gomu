package stream

import (
	"github.com/urfave/cli/v2"
)

func NewCommand(alias string) *cli.Command {
	return &cli.Command{
		Name:  "stream",
		Usage: "Create a service stream",
		Subcommands: []*cli.Command{
			{
				Name:    "bidi",
				Aliases: []string{"b"},
				Usage:   "Create a bidirectional service stream, e.g. " + alias + " stream bidirectional helloworld Helloworld.PingPong '{\"stroke\": 1}' '{\"stroke\": 2}'",
				Action:  Bidirectional,
			},
			{
				Name:    "client",
				Aliases: []string{"c"},
				Usage:   "Create a client service stream, e.g. " + alias + " stream client helloworld Helloworld.ClientStream '{\"stroke\": 1}' '{\"stroke\": 2}'",
				Action:  Client,
			},
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "Create a server service stream, e.g. " + alias + " stream server helloworld Helloworld.ServerStream '{\"count\": 10}'",
				Action:  Server,
			},
		},
	}
}
