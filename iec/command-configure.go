package main

import (
	"github.com/bbucko/cli-iec/common"
	"github.com/urfave/cli"
	"log"
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
	config, err := common.OpenConfig(c, "ns", "js")

	if err != nil {
		log.Fatal("Error encountered during initialization of configuration: ", err)
		return err
	}

	log.Print("Config: ", config)

	return nil
}
