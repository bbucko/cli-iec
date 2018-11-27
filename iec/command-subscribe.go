package main

import (
	"github.com/urfave/cli"
)

var commandSubscribe = cli.Command{
	Name:        "subscribe",
	ArgsUsage:   "",
	Description: "",
	HideHelp:    true,
	Action:      callSubscribe,
	Flags:       []cli.Flag{},
}

func callSubscribe(c *cli.Context) error {
	return nil
}
