package main

import (
	"github.com/urfave/cli"
)

var commandGenerateToken = cli.Command{
	Name:        "token",
	ArgsUsage:   "",
	Description: "",
	HideHelp:    true,
	Action:      callGenerateToken,
	Flags:       []cli.Flag{},
}

func callGenerateToken(c *cli.Context) error {
	return nil
}
