package common

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	keyName string
}

func cliConfigPath() (configPath string) {
	cliHome := os.Getenv("AKAMAI_CLI_HOME")
	if cliHome == "" {
		home, _ := homedir.Dir()
		cliHome = filepath.Join(home, ".akamai-cli")
	}
	configPath = filepath.Join(cliHome, "config")
	return
}

func openFile() (configFile *ini.File, err error) {
	configPath := cliConfigPath()
	log.Print("Loading INI file from: ", configPath)
	configFile, err = ini.Load(configPath)
	if err != nil {
		return
	}
	return
}

func saveFile(cfg *ini.File) (err error) {
	configPath := cliConfigPath()
	err = cfg.SaveTo(configPath)
	if err != nil {
		return
	}
	return nil
}

func OpenConfig(c *cli.Context, namespace string, jurisdiction string) (config *Configuration, err error) {
	cfg, err := openFile()
	if err != nil {
		return
	}

	section := cfg.Section("iec")
	if section == nil {
		section, err = cfg.NewSection("iec")
	}

	if err != nil {
		return
	}

	err = saveFile(cfg)
	if err != nil {
		return
	}

	config = new(Configuration)
	config.keyName = section.Key(fmt.Sprint(namespace, "_", jurisdiction, "_", "key")).String()

	return
}
