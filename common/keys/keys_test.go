package keys

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
	"testing"
)

func TestKeysInTempDir(t *testing.T) {
	env := os.Getenv("AKAMAI_CLI_HOME")
	os.Setenv("AKAMAI_CLI_HOME", os.TempDir())

	t.Run("TestFetchRSAKeyByName", func(t *testing.T) {
		keyName := randomName(t)
		key := tryCreateRSAKey(t, keyName, 128)
		key.Persist()
		fetchedKey, er := FetchRSAKeyByName(keyName)
		if er != nil {
			t.Fatal("Error fetching created key")
		}
		if er != nil {
			t.Fatal("Error fetching created key")
		}
		if key != fetchedKey {
			t.Fatal("Created key is different than persisted")
		}
	})
	t.Run("TestRSAKeyFormat", func(t *testing.T) {
		key := tryCreateRSAKey(t, randomName(t), 128)

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
	})
	t.Run("TestKeysPublicKeySectionName", func(t *testing.T) {
		key := tryCreateRSAKey(t, "some-name", 128)
		expectedName := "some-name_public"

		if !(key.PublicKeySectionName() == expectedName) {
			t.Fatalf("Public key section name format error, actual [%v] != expected [%v]", key.PublicKeySectionName(), expectedName)
		}
	})
	t.Run("TestKeysPrivateKeySectionName", func(t *testing.T) {
		keyName := randomName(t)
		key := tryCreateRSAKey(t, keyName, 128)
		expectedName := fmt.Sprintf("%v_private", keyName)

		if !(key.PrivateKeySectionName() == expectedName) {
			t.Fatalf("Private key section name format error, actual [%v] != expected [%v]", key.PrivateKeySectionName(), expectedName)
		}
	})

	os.Setenv("AKAMAI_CLI_HOME", env)
}

func tryCreateRSAKey(t *testing.T, name string, bits int) RSAKey {
	key, er := CreateRSAKey(name, bits)
	if er != nil {
		t.Fatal("Error creating key")
	}
	return key
}

func randomName(t *testing.T) string {
	value, err := uuid.NewRandom()
	if err != nil {
		t.Fatal("Error when generate random name")
	}
	return value.String()
}
