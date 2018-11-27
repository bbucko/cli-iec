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

	cfg, _ := ini.Load(configFile)
	if cfg.Section("iec") == nil {
		t.Fatal("missing section: ", "iec")
	}

	if cfg.Section("jwt") == nil {
		t.Fatal("missing section: ", "jwt")
	}
}
