package main

import (
	"github.com/urfave/cli"
)

var commandPublish = cli.Command{
	Name:        "publish",
	ArgsUsage:   "",
	Description: "",
	HideHelp:    true,
	Action:      callPublish,
	Flags:       []cli.Flag{},
}

func callPublish(c *cli.Context) error {
	return nil
}
