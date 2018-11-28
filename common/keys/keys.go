package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/bbucko/cli-iec/ini-repo"
	"log"
)

type RSAKey struct {
	KeyName    string
	PrivateKey string
	PublicKey  string
}

func FetchRSAKeyByName(name string) (key RSAKey, err error) {
	log.Printf("Searching repository for RSA key with name: [%v]", name)
	privateKey, err := ini_repo.GetValue("keys", fmt.Sprintf("%v_private", name))
	if err != nil {
		log.Fatal(err)
		return
	}
	publicKey, err := ini_repo.GetValue("keys", fmt.Sprintf("%v_public", name))
	if err != nil {
		log.Fatal(err)
		return
	}
	return RSAKey{name, privateKey, publicKey}, nil
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
	ini_repo.Persist("keys", key.PublicKeySectionName(), key.PublicKey)
	ini_repo.Persist("keys", key.PrivateKeySectionName(), key.PrivateKey)
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
