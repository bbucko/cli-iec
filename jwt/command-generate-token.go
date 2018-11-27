package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli"
)

var commandGenerateToken = cli.Command{
	Name:        "token",
	ArgsUsage:   "",
	Description: "",
	HideHelp:    true,
	Action:      callGenerateToken,
	Flags:       []cli.Flag{},
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
	var signingKey = []byte("AllYourBase") // TODO get from file? stdout?

	var customClaims = map[string]string{ // TODO get from context
		"clientId":   "username",
		"authGroups": "testGroup",
	}
	claims := constructClaims(customClaims)
	token, err := createToken(claims, signingKey)

	if err != nil {
		fmt.Errorf("Error generating token %v", err)
	} else {
		fmt.Println("JWT Token:", token)
	}

	return nil
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
