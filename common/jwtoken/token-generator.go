package jwtoken

import (
	"errors"
	"fmt"
	"github.com/bbucko/cli-iec/common"
	"github.com/bbucko/cli-iec/common/keys"
	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/cli"
	"log"
)

type Token struct {
	Raw       string                 // The raw token.  Populated when you Parse a token
	Method    jwt.SigningMethod      // The signing method used or to be used
	Header    map[string]interface{} // The first segment of the token
	Claims    jwt.Claims             // The second segment of the token
	Signature string                 // The third segment of the token.  Populated when you Parse a token
	Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
}

type JWTParams struct {
	Namespace       string
	Jurisdiction    string
	ClientId        string
	ClientIdClaim   string
	AuthGroups      string
	AuthGroupsClaim string
}

func GenerateToken(keyName string, params JWTParams) (token string, err error) {
	var signingKey, _ = keys.FetchRSAKeyByName(keyName)

	token, err = CreateToken(params, []byte(signingKey.PrivateKey))
	if err != nil {
		log.Fatal(err)
		return
	}

	return token, err
}

func GetPrivateKey(c *cli.Context, params JWTParams) []byte {
	var config, _ = common.OpenConfig(c, params.Namespace, params.Jurisdiction)
	var privateKey = config.JwtConfig.Key().PrivateKey

	fmt.Println("Private Key:", privateKey)
	return []byte(privateKey)
}

func constructClaims(params JWTParams) jwt.Claims {
	var customClaims = jwt.MapClaims{
		params.ClientIdClaim:   params.ClientId,
		params.AuthGroupsClaim: params.AuthGroups,
	}
	//fmt.Printf("%v\n", customClaims[params.ClientIdClaim])
	//customClaims[params.ClientIdClaim] = params.ClientId

	return customClaims
}

func CreateToken(params JWTParams, signBytes []byte) (string, error) {
	claims := constructClaims(params)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	if err != nil {
		return "", errors.New("Error parsing RSA private key from PEM!")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)
	return ss, err
}
