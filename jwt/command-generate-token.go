package main

import (
	"errors"
	"fmt"
	"github.com/akamai/cli-iec/common/keys"
	"github.com/bbucko/cli-iec/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli"
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

type IEClaims struct {
	ClientId  string `json:"clientId"`
	AuthGroups string  `json:"groups"`
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

type JWTParams struct {
	namespace string
	jurisdiction string
	clientId string
	clientIdClaim string
	authGroups string
	authGroupsClaim string
}

func GenerateToken(keyName string, params JWTParams) (string, error) {
	var signingKey, _ = keys.FetchRSAKeyByName(keyName)

	token, err := createToken(params, []byte(signingKey.PrivateKey))
	return token, err
}

func getPrivateKey(c *cli.Context, params JWTParams) []byte {
	var config, _ = common.OpenConfig(c, params.namespace, params.jurisdiction)
	var privateKey = config.JwtConfig.Key().PrivateKey

	fmt.Println("Private Key:", privateKey)
	return []byte(privateKey)
}

func constructClaims(params JWTParams) jwt.Claims {
	var customClaims = jwt.MapClaims{
		params.clientIdClaim:   params.clientId,
		params.authGroupsClaim: params.authGroups,
	}
	fmt.Printf("%v\n", customClaims[params.clientIdClaim])
	customClaims[params.clientIdClaim] = params.clientId

	return IEClaims{
		params.clientId,
		params.authGroups,
		jwt.StandardClaims{
			ExpiresAt: 25000,
		},
	}
}

func createToken(params JWTParams, signBytes []byte) (string, error) {
	claims := constructClaims(params)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	if err != nil {
		return "", errors.New("Error parsing RSA private key from PEM!")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)
	return ss, err
}

func callGenerateToken(c *cli.Context) error {
	params := JWTParams{
		c.String("namespace"),
		c.String("jurisdiction"),
		c.String("clientId"),
		c.String("clientIdClaim"),
		c.String("authGroups"),
		c.String("authGroupsClaim"),
	}
	signingKey := getPrivateKey(c, params)
	token, err := createToken(params, signingKey)

	if err != nil {
		fmt.Errorf("Error generating token %v", err)
	} else {
		fmt.Println("JWT Token:", token)
	}

	return nil
}