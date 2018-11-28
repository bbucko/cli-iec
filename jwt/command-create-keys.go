package main

import (
	"github.com/bbucko/cli-iec/common/keys"
	"github.com/urfave/cli"
	"log"
)

var commandCreateKeys = cli.Command{
	Name:        "create-keys",
	ArgsUsage:   "",
	Description: "",
	HideHelp:    true,
	Action:      callCreateKeys,
	Flags: []cli.Flag{
		cli.IntFlag{"bits", "RSA key size (default value is 2048)", "AKAMAI_JWT_BITS", false, 2048, nil},
		cli.StringFlag{"name", "Name of the key (default vaue is myspace)", "AKAMAI_JWT_KEY_NAME", false, "myspace", nil},
	},
}

func callCreateKeys(c *cli.Context) error {
	name := c.String("name")
	bits := c.Int("bits")

	key, er := keys.CreateRSAKey(name, bits)
	if er != nil {
		log.Fatal("Creating keys failed for ")
		return er
	}

	log.Printf("%s", key.PrivateKey)
	log.Printf("%s", key.PublicKey)

	key.Persist()
	return nil
}
