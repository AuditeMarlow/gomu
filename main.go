package main

import (
	"github.com/auditemarlow/gomu/cmd"

	// register commands
	_ "github.com/auditemarlow/gomu/cmd/cli"
)

func main() {
	cmd.Run()
}
