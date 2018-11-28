package main

import (
	"github.com/akamai/cli-common-golang"
	"github.com/bbucko/cli-iec/common"
	"github.com/urfave/cli"
	"log"
	"time"
)

var commandConfigure = cli.Command{
	Name:        "configure",
	ArgsUsage:   "",
	Description: "",
	HideHelp:    true,
	Action:      callConfigure,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "namespace",
			Usage: "Namespace",
			Value: "ns",
		}, cli.StringFlag{
			Name:  "jurisdiction",
			Usage: "Jurisdiction (EU, NA, JP)",
			Value: "EU",
		}, cli.StringFlag{
			Name:  "auth",
			Usage: "Authentication method",
			Value: "jwt",
		}, cli.StringFlag{
			Name:  "hostname",
			Usage: "Hostname that clients will connect to",
			Value: "iec.kwapiszewski.com",
		}, cli.BoolFlag{
			Name:  "mqtt",
			Usage: "Enable MQTT over TLS",
		}, cli.BoolFlag{
			Name:  "ws",
			Usage: "Enable MQTT over WebSockets",
		}, cli.BoolFlag{
			Name:  "https",
			Usage: "Enable Gateway",
		}, cli.StringFlag{
			Name:  "jwtKey",
			Usage: "Name of the JWT Key that will be used for generating/validating authentication tokens",
			Value: "default",
		}, cli.BoolFlag{
			Name:  "activate",
			Usage: "If present this configuration will be activated",
		},
	},
}

func callConfigure(c *cli.Context) error {
	config, err := common.OpenConfig(c, "ns", "js")

	if err != nil {
		log.Fatal("Error encountered during initialization of configuration: ", err)
		return err
	}

	log.Print("Current configuration: ", config)

	dummyOperation("Uploading keys to JWT")
	dummyOperation("Configuring new Property")
	dummyOperation("Activating Property")
	dummyOperation("Creating new namespace in IEC")
	dummyOperation("Updating topics")
	dummyOperation("Activating namespace")

	err = common.SaveConfig(c, config)
	if err != nil {
		log.Fatal("Error encountered during saving of configuration: ", err)
		return err
	}

	return nil
}

func dummyOperation(operation string) {
	akamai_cli.StartSpinner(operation, "Done")
	time.Sleep(1 * time.Second)
	akamai_cli.StopSpinnerOk()
}
