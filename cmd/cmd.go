package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type Cmd interface {
	App() *cli.App
	Run() error
}

type command struct {
	app *cli.App
}

func (c *command) App() *cli.App {
	return c.app
}

func (c *command) Run() error {
	return c.app.Run(os.Args)
}

var (
	DefaultCmd Cmd = New()

	name        string = "gomu"
	description string = "A go-micro helper tool"
)

func New() *command {
	c := &command{}
	c.app = cli.NewApp()
	c.app.Name = name
	c.app.Usage = description

	return c
}

func Register(cmds ...*cli.Command) {
	a := DefaultCmd.App()
	a.Commands = append(a.Commands, cmds...)
}

func Run() {
	if err := DefaultCmd.Run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
