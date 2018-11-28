package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/bbucko/cli-iec/common/jwtoken"
)

var commandGenerateToken = cli.Command{
	Name:        "token",
	ArgsUsage:   "[name] [jurisdiction] [clientId] [clientIdClaim] [authGroups] [authGroupsClaim]",
	Description: "",
	HideHelp:    true,
	Action:      callGenerateToken,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "namespace",
			Usage:       "Namespace name",
			EnvVar:      "",
			Hidden:      false,
			Value:       "",
			Destination: nil,
		},
		cli.StringFlag{
			Name:        "name",
			Usage:       "Key name",
			EnvVar:      "",
			Hidden:      true,
			Value:       "",
			Destination: nil,
		},
		cli.StringFlag{
			Name:        "jurisdiction",
			Usage:       "Namespace jurisdiction",
			EnvVar:      "",
			Hidden:      true,
			Value:       "na",
			Destination: nil,
		},
		cli.StringFlag{
			Name:        "clientId",
			Usage:       "Client id",
			EnvVar:      "",
			Hidden:      false,
			Value:       "",
			Destination: nil,
		},
		cli.StringFlag{
			Name:        "clientIdClaim",
			Usage:       "Client id claim",
			EnvVar:      "",
			Hidden:      true,
			Value:       "clientId",
			Destination: nil,
		},
		cli.StringFlag{
			Name:        "authGroups",
			Usage:       "Auth groups",
			EnvVar:      "",
			Hidden:      false,
			Value:       "",
			Destination: nil,
		},
		cli.StringFlag{
			Name:        "authGroupsClaim",
			Usage:       "Auth groups claim",
			EnvVar:      "",
			Hidden:      true,
			Value:       "authGroups",
			Destination: nil,
		},
	},
}



func callGenerateToken(c *cli.Context) error {
	params := jwtoken.JWTParams{
		c.String("namespace"),
		c.String("jurisdiction"),
		c.String("clientId"),
		c.String("clientIdClaim"),
		c.String("authGroups"),
		c.String("authGroupsClaim"),
	}
	signingKey := jwtoken.GetPrivateKey(c, params)
	token, err := jwtoken.CreateToken(params, signingKey)

	if err != nil {
		fmt.Errorf("Error generating token %v", err)
	} else {
		fmt.Println("JWT Token:", token)
	}

	return nil
}
