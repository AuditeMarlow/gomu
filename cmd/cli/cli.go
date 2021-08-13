package cli

import (
	"github.com/auditemarlow/gomu/cmd"
	"github.com/auditemarlow/gomu/cmd/cli/call"
	"github.com/auditemarlow/gomu/cmd/cli/new"
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
	)
}
