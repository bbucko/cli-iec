package jwt

import (
	"strings"
	"testing"
)

func TestFetchKeysByName(t *testing.T) {
	keyName := "my-keys"
	keys := tryCreateKeys(t, keyName, 128)

	keys.persist()
	fetchedKeys, er := FetchKeysByName(keyName)

	if er != nil {
		t.Fatal("Error fetching created key")
	}
	if keys != fetchedKeys {
		t.Fatal("Created keys is different than persisted")
	}
}

func TestKeysFormat(t *testing.T) {
	keys := tryCreateKeys(t, "some-name", 128)

	if !strings.Contains(keys.Private, "-----BEGIN RSA PRIVATE KEY-----") {
		t.Fatal("Private key format error")
	}
	if !strings.Contains(keys.Private, "-----END RSA PRIVATE KEY-----") {
		t.Fatal("Private key format error")
	}
	if !strings.Contains(keys.Public, "-----BEGIN RSA PUBLIC KEY-----") {
		t.Fatal("Public key format error")
	}
	if !strings.Contains(keys.Public, "-----END RSA PUBLIC KEY-----") {
		t.Fatal("Public key format error")
	}
}

func TestKeysPublicKeyName(t *testing.T) {
	keys := tryCreateKeys(t, "some-name", 128)
	expectedName := "some-name_pub"

	if !(keys.publicKeyName() == expectedName) {
		t.Fatalf("Public key name format error, actual [%v] != expected [%v]", keys.publicKeyName(), expectedName)
	}
}

func TestKeysPrivateKeyName(t *testing.T) {
	keys := tryCreateKeys(t, "some-name", 128)
	expectedName := "some-name_prv"

	if !(keys.privateKeyName() == expectedName) {
		t.Fatalf("Private key name format error, actual [%v] != expected [%v]", keys.privateKeyName(), expectedName)
	}
}

func tryCreateKeys(t *testing.T, name string, bits int) Keys {
	keys, er := createKeys(name, bits)
	if er != nil {
		t.Fatal("Error creating key")
	}
	return keys
}
