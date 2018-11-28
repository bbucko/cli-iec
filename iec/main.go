package main

import (
	akamai "github.com/akamai/cli-common-golang"
	"github.com/urfave/cli"
	"os"
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
		commandConfigure,
		commandSubscribe,
		commandPublish,
	}

	return commands, nil
}

func main() {
	akamai.CreateApp(
		"iec",
		"",
		"",
		"0.0.1",
		"default",
		commandLocator,
	)
	akamai.App.Run(os.Args)
}
