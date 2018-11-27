package main

import (
	"github.com/urfave/cli"
)

var commandGenerateKeys = cli.Command{
	Name:        "generate-keys",
	ArgsUsage:   "",
	Description: "Generate keys",
	HideHelp:    true,
	Action:      callGenerateKey,
	Flags:       []cli.Flag{},
}

func callGenerateKey(c *cli.Context) error {
	return nil
}
