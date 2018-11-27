package main

import (
	"github.com/urfave/cli"
)

var commandGenerateToken = cli.Command{
	Name:        "generate-token",
	ArgsUsage:   "",
	Description: "Generate token",
	HideHelp:    true,
	Action:      callGenerateToken,
	Flags:       []cli.Flag{},
}

func callGenerateToken(c *cli.Context) error {
	return nil
}
