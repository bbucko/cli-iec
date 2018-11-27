package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
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

var keysRepository = make(map[string]Keys)

func callCreateKeys(c *cli.Context) error {
	name := c.String("name")
	bits := c.Int("bits")

	keys, er := createKeys(name, bits)
	if er != nil {
		log.Fatal("Creating keys failed for ")
		return er
	}

	log.Printf("%s", keys.Private)
	log.Printf("%s", keys.Public)

	keys.persist()
	return nil
}

type Keys struct {
	Name    string
	Private string
	Public  string
}

func FetchKeysByName(name string) (Keys, error) {
	log.Printf("Searching repository for keys with name: [%v]", name)

	return keysRepository[name], nil
}

func createKeys(name string, bits int) (Keys, error) {
	log.Printf("Creating keys with name [%v] with size [%v] bits", name, bits)
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Println(err)
		return Keys{}, err
	}
	privateKeyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	publicKeyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}
	return Keys{name, string(pem.EncodeToMemory(privateKeyPem)), string(pem.EncodeToMemory(publicKeyPem))}, nil
}

func (keys Keys) persist() {
	log.Printf("Saving keys to repository: %v", keys)

	log.Println(keys.Public)

	keysRepository[keys.Name] = keys
}

func (keys Keys) publicKeyName() string {
	return fmt.Sprintf("%v_pub", keys.Name)
}

func (keys Keys) privateKeyName() string {
	return fmt.Sprintf("%v_prv", keys.Name)
}

func (keys Keys) String() string {
	return fmt.Sprintf("[name: %v]", keys.Name)
}
