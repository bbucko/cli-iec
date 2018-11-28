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

	if config.Namespace != "ns" {
		t.Fatal("Namespace:", config.Namespace, "does not match", "ns")
	}

	if config.Jurisdiction != "js" {
		t.Fatal("Jurisdiction:", config.Jurisdiction, "does not match", "js")
	}

	if config.JwtConfig.Key().KeyName != "key1" {
		t.Fatal("Key name:", config.JwtConfig.Key().KeyName, "does not match", "key1")
	}

	if config.JwtConfig.Key().PublicKey != "abc" {
		t.Fatal("public Key name:", config.JwtConfig.Key().PublicKey, "does not match", "abc")
	}

	if config.JwtConfig.Key().PrivateKey != "def" {
		t.Fatal("private Key name:", config.JwtConfig.Key().PrivateKey, "does not match", "def")
	}
}

func TestReadingEmptyConfigFile(t *testing.T) {
	tempHome, _ := ioutil.TempDir("", "akamai_cli_")
	configFile := filepath.Join(tempHome, "config")
	os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0666)
	defer os.RemoveAll(tempHome)
	_ = os.Setenv("AKAMAI_CLI_HOME", tempHome)

	config, _ := OpenConfig(nil, "ns", "js")

	if config.Namespace != "ns" {
		t.Fatal("Namespace:", config.Namespace, "does not match", "ns")
	}

	if config.Jurisdiction != "js" {
		t.Fatal("Jurisdiction:", config.Jurisdiction, "does not match", "js")
	}

	if config.JwtConfig.Key().KeyName != "default" {
		t.Fatal("default public Key:", config.JwtConfig.Key().KeyName, "does not match", "default")
	}

	cfg, _ := ini.Load(configFile)
	if cfg.Section("iec") == nil {
		t.Fatal("missing section: ", "iec")
	}

	if cfg.Section("jwt") == nil {
		t.Fatal("missing section: ", "jwt")
	}

	if cfg.Section("keys") == nil {
		t.Fatal("missing section: ", "keys")
	}
}
