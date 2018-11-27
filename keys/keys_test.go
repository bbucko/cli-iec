package keys

import (
	"strings"
	"testing"
)

func TestFetchRSAKeyByName(t *testing.T) {
	keyName := "my-key"
	key := tryCreateRSAKey(t, keyName, 128)

	key.Persist()
	fetchedKeys, er := FetchRSAKeyByName(keyName)

	if er != nil {
		t.Fatal("Error fetching created key")
	}
	if key != fetchedKeys {
		t.Fatal("Created key is different than persisted")
	}
}

func TestRSAKeyFormat(t *testing.T) {
	key := tryCreateRSAKey(t, "some-name", 128)

	if !strings.Contains(key.PrivateKey, "-----BEGIN RSA PRIVATE KEY-----") {
		t.Fatal("Private key format error")
	}
	if !strings.Contains(key.PrivateKey, "-----END RSA PRIVATE KEY-----") {
		t.Fatal("Private key format error")
	}
	if !strings.Contains(key.PublicKey, "-----BEGIN RSA PUBLIC KEY-----") {
		t.Fatal("Public key format error")
	}
	if !strings.Contains(key.PublicKey, "-----END RSA PUBLIC KEY-----") {
		t.Fatal("Public key format error")
	}
}

func TestKeysPublicKeySectionName(t *testing.T) {
	key := tryCreateRSAKey(t, "some-name", 128)
	expectedName := "some-name_public"

	if !(key.PublicKeySectionName() == expectedName) {
		t.Fatalf("Public key section name format error, actual [%v] != expected [%v]", key.PublicKeySectionName(), expectedName)
	}
}

func TestKeysPrivateKeySectionName(t *testing.T) {
	key := tryCreateRSAKey(t, "some-name", 128)
	expectedName := "some-name_private"

	if !(key.PrivateKeySectionName() == expectedName) {
		t.Fatalf("Private key section name format error, actual [%v] != expected [%v]", key.PrivateKeySectionName(), expectedName)
	}
}

func tryCreateRSAKey(t *testing.T, name string, bits int) RSAKey {
	key, er := CreateRSAKey(name, bits)
	if er != nil {
		t.Fatal("Error creating key")
	}
	return key
}
