package cli

import (
	"github.com/auditemarlow/gomu/cmd"
	"github.com/auditemarlow/gomu/cmd/cli/call"
	"github.com/auditemarlow/gomu/cmd/cli/new"
	"github.com/auditemarlow/gomu/cmd/cli/run"
	"github.com/auditemarlow/gomu/cmd/cli/stream"
)

var (
	alias string = cmd.DefaultCmd.App().Name
)

func init() {
	cmd.Register(
		call.NewCommand(alias),
		new.NewCommand(alias),
		run.NewCommand(alias),
		stream.NewCommand(alias),
	)
}
