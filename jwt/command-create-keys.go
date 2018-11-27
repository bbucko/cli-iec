package main

import (
	"github.com/urfave/cli"
)

var commandCreateKeys = cli.Command{
	Name:        "create-keys",
	ArgsUsage:   "",
	Description: "",
	HideHelp:    true,
	Action:      callCreateKeys,
	Flags:       []cli.Flag{},
}

func callCreateKeys(c *cli.Context) error {
	return nil
}
