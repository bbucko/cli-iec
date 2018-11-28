package main

import (
	"errors"
	"fmt"
	"github.com/bbucko/cli-iec/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli"
)

var commandGenerateToken cli.Command = cli.Command{
	Name:        "token",
	ArgsUsage:   "[name] [jurisdiction]",
	Description: "",
	HideHelp:    true,
	Action:      callGenerateToken,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "Namespace name",
			EnvVar:      "",
			Hidden:      false,
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
			Value:       "client-id",
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
			Value:       "auth-groups",
			Destination: nil,
		},
	},
}

type IEClaims struct {
	customClaims map[string]string
	jwt.StandardClaims
}

type Token struct {
	Raw       string                 // The raw token.  Populated when you Parse a token
	Method    jwt.SigningMethod      // The signing method used or to be used
	Header    map[string]interface{} // The first segment of the token
	Claims    jwt.Claims             // The second segment of the token
	Signature string                 // The third segment of the token.  Populated when you Parse a token
	Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
}

func callGenerateToken(c *cli.Context) error {
	var signingKey = getPublicKey(c, c.String("name"), c.String("jurisdiction"))

	var customClaims = map[string]string{
		c.String("clientIdClaim"):   c.String("clientId"),
		c.String("authGroupsClaim"): c.String("authGroups"),
	}
	claims := constructClaims(customClaims)
	fmt.Println("JWT Token Claims:", customClaims)

	token, err := createToken(claims, signingKey)

	if err != nil {
		fmt.Errorf("Error generating token %v", err)
	} else {
		fmt.Println("JWT Token:", token)
	}

	return nil
}

func getPublicKey(c *cli.Context, name string, jurisdiction string) []byte {
	config, err := common.OpenConfig(c, name, jurisdiction)

	fmt.Println(config, err)

	return []byte("AllYourBase") // TODO: get the public key from config
}

func constructClaims(customClaims map[string]string) IEClaims {
	return IEClaims{
		customClaims,
		jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
}

func createToken(claims IEClaims, signingKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)

	if err != nil {
		return "", errors.New("Error generating token!")
	}
	return ss, nil
}
