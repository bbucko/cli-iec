package common

import (
	"github.com/go-ini/ini"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestReadingFilledConfigFile(t *testing.T) {
	wd, _ := os.Getwd()

	_ = os.Setenv("AKAMAI_CLI_HOME", filepath.Join(wd, "testdata"))
	config, err := OpenConfig(nil, "ns", "js")
	if err != nil {
		t.Fatal(err)
	}

	if config.namespace != "ns" {
		t.Fatal("namespace:", config.namespace, "does not match", "ns")
	}

	if config.jurisdiction != "js" {
		t.Fatal("jurisdiction:", config.jurisdiction, "does not match", "js")
	}

	if config.key.keyName != "key1" {
		t.Fatal("key name:", config.key.keyName, "does not match", "key1")
	}

	if config.key.publicKey != "abc" {
		t.Fatal("public key name:", config.key.publicKey, "does not match", "abc")
	}

	if config.key.privateKey != "def" {
		t.Fatal("private key name:", config.key.privateKey, "does not match", "def")
	}
}

func TestReadingEmptyConfigFile(t *testing.T) {
	tempHome, _ := ioutil.TempDir("", "akamai_cli_")
	configFile := filepath.Join(tempHome, "config")
	os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0666)
	defer os.RemoveAll(tempHome)
	_ = os.Setenv("AKAMAI_CLI_HOME", tempHome)

	config, _ := OpenConfig(nil, "ns", "js")

	if config.namespace != "ns" {
		t.Fatal("namespace:", config.namespace, "does not match", "ns")
	}

	if config.jurisdiction != "js" {
		t.Fatal("jurisdiction:", config.jurisdiction, "does not match", "js")
	}

	if config.key.keyName != "default" {
		t.Fatal("default public key:", config.key.keyName, "does not match", "default")
	}

	cfg, _ := ini.Load(configFile)
	if cfg.Section("iec") == nil {
		t.Fatal("missing section: ", "iec")
	}

	if cfg.Section("jwt") == nil {
		t.Fatal("missing section: ", "jwt")
	}
}
