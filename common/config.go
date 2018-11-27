package common

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
)

func OpenConfig(c *cli.Context, namespace string, jurisdiction string) (config *Configuration, err error) {
	cfg, err := openFile()
	if err != nil {
		return
	}

	err = ensureSection(cfg, "iec")
	if err != nil {
		return
	}

	err = ensureSection(cfg, "jwt")
	if err != nil {
		return
	}

	err = saveFile(cfg)
	if err != nil {
		return
	}

	config = NewConfiguration(namespace, jurisdiction, cfg.Section("iec"))

	return
}

func SaveConfig(c *cli.Context, configuration *Configuration) (err error) {
	cfg, err := openFile()
	if err != nil {
		return
	}

	section := cfg.Section("iec")
	configuration.UpdateSection(section)
	return nil
}

func ensureSection(cfg *ini.File, sectionName string) (err error) {
	section := cfg.Section(sectionName)
	if section == nil {
		section, err = cfg.NewSection("iec")
		if err != nil {
			return
		}
	}
	return nil
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

type Configuration struct {
	namespace    string
	jurisdiction string
	keyName      string
}

func NewConfiguration(namespace string, jurisdiction string, section *ini.Section) (config *Configuration) {
	config = new(Configuration)
	config.namespace = namespace
	config.jurisdiction = jurisdiction
	config.keyName = section.Key(config.key("jwtKeyName")).String()
	return config
}

func (c *Configuration) UpdateSection(section *ini.Section) {

}

func (c *Configuration) key(key string) string {
	return fmt.Sprint(c.namespace, "_", c.jurisdiction, "_", key)
}
