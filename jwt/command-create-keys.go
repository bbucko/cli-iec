package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/urfave/cli"
	"log"
)

var bits int
var keyName string

var commandCreateKeys = cli.Command{
	Name:        "create-keys",
	ArgsUsage:   "",
	Description: "",
	HideHelp:    true,
	Action:      callCreateKeys,
	Flags: []cli.Flag{
		cli.IntFlag{"bits", "RSA key size (default value is 2048)", "AKAMAI_JWT_BITS", false, 2048, &bits},
		cli.StringFlag{"name", "Name of the key (default vaue is myspace)", "AKAMAI_JWT_KEY_NAME", false, "myspace", &keyName},
	},
}

func callCreateKeys(c *cli.Context) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, c.Int("bits"))
	if err != nil {
		log.Println(err)
		return err
	}

	privateKeyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	log.Printf("%s", pem.EncodeToMemory(privateKeyPem))

	publicKeyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}

	log.Printf("%s", pem.EncodeToMemory(publicKeyPem))

	return nil
}
