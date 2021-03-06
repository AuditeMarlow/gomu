package cli

import (
	"github.com/auditemarlow/gomu/cmd"
	"github.com/auditemarlow/gomu/cmd/cli/call"
	"github.com/auditemarlow/gomu/cmd/cli/describe"
	"github.com/auditemarlow/gomu/cmd/cli/new"
	"github.com/auditemarlow/gomu/cmd/cli/run"
	"github.com/auditemarlow/gomu/cmd/cli/services"
	"github.com/auditemarlow/gomu/cmd/cli/stream"
)

var (
	alias string = cmd.DefaultCmd.App().Name
)

func init() {
	cmd.Register(
		call.NewCommand(alias),
		describe.NewCommand(alias),
		new.NewCommand(alias),
		run.NewCommand(alias),
		services.NewCommand(alias),
		stream.NewCommand(alias),
	)
}
