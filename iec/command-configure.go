package main

import (
	"github.com/urfave/cli"
)

var commandConfigure = cli.Command{
	Name:        "configure",
	ArgsUsage:   "",
	Description: "",
	HideHelp:    true,
	Action:      callConfigure,
	Flags:       []cli.Flag{},
}

func callConfigure(c *cli.Context) error {
	return nil
}
