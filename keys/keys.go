package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

// FIXME: remove and use ini file implementation
var keysRepository = make(map[string]RSAKey)

type RSAKey struct {
	KeyName    string
	PrivateKey string
	PublicKey  string
}

func FetchRSAKeyByName(name string) (RSAKey, error) {
	log.Printf("Searching repository for RSA key with name: [%v]", name)
	return keysRepository[name], nil
}

func CreateRSAKey(name string, bits int) (RSAKey, error) {
	log.Printf("Creating RSA key with name [%v] with size [%v] bits", name, bits)
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Println(err)
		return RSAKey{}, err
	}
	privateKeyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	publicKeyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}
	return RSAKey{name, string(pem.EncodeToMemory(privateKeyPem)), string(pem.EncodeToMemory(publicKeyPem))}, nil
}

func (key RSAKey) Persist() {
	log.Printf("Saving RSA key: %v", key)
	keysRepository[key.KeyName] = key
}

func (key RSAKey) PublicKeySectionName() string {
	return fmt.Sprintf("%v_public", key.KeyName)
}

func (key RSAKey) PrivateKeySectionName() string {
	return fmt.Sprintf("%v_private", key.KeyName)
}

func (key RSAKey) String() string {
	return fmt.Sprintf("[name: %v]", key.KeyName)
}
