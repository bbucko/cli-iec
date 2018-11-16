package main

import (
	"github.com/urfave/cli"
	"os"

	akamai "github.com/akamai/cli-common-golang"
)

var commandLocator akamai.CommandLocator = func() ([]cli.Command, error) {
	commands := []cli.Command{
		{
			Name:        "list",
			Description: "List commands",
			Action:      akamai.CmdList,
		},
		{
			Name:         "help",
			Description:  "Displays help information",
			ArgsUsage:    "[command] [sub-command]",
			Action:       akamai.CmdHelp,
			BashComplete: akamai.DefaultAutoComplete,
		},
	}

	return commands, nil
}

func main() {
	akamai.CreateApp(
		"iec-connect",
		"",
		"",
		"0.0.1",
		"default",
		commandLocator,
	)
	akamai.App.Run(os.Args)
}
