package cli

import (
	"github.com/auditemarlow/gomu/cmd"
	"github.com/auditemarlow/gomu/cmd/cli/call"
	"github.com/auditemarlow/gomu/cmd/cli/new"
	"github.com/auditemarlow/gomu/cmd/cli/stream"
)

func init() {
	cmd.Register(
		call.NewCommand(),
		new.NewCommand(),
		stream.NewCommand(),
	)
}
