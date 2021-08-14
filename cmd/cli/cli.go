package cli

import (
	"github.com/auditemarlow/gomu/cmd"
	"github.com/auditemarlow/gomu/cmd/cli/call"
	"github.com/auditemarlow/gomu/cmd/cli/new"
	"github.com/auditemarlow/gomu/cmd/cli/stream"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:   "call",
			Usage:  "gomu call greeter Say.Hello '{\"name\": \"John\"}'",
			Action: call.Run,
		},
		&cli.Command{
			Name:   "new",
			Usage:  "gomu new greeter",
			Action: new.Run,
		},
		&cli.Command{
			Name:   "stream",
			Usage:  "gomu stream greeter Say.HelloStream '{\"name\": \"John\"}'",
			Action: stream.Run,
		},
	)
}
