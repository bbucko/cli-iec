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

	iecSection := cfg.Section("iec")
	jwtSection := cfg.Section("jwt")

	config = NewConfiguration(namespace, jurisdiction)
	keyName := iecSection.Key(config.configKey("jwtKeyName")).MustString("default")
	config.key = NewRSAKey(keyName, jwtSection)
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

type RSAKey struct {
	publicKey  string
	privateKey string
	keyName    string
}

func NewRSAKey(keyName string, section *ini.Section) (key *RSAKey) {
	key = new(RSAKey)
	key.keyName = keyName
	key.publicKey = section.Key(key.configKey("public")).String()
	key.privateKey = section.Key(key.configKey("private")).String()
	return key
}

func (k *RSAKey) configKey(key string) string {
	return fmt.Sprint(k.keyName, "_", key)
}

type Configuration struct {
	namespace    string
	jurisdiction string
	key          *RSAKey
}

func NewConfiguration(namespace string, jurisdiction string) (config *Configuration) {
	config = new(Configuration)
	config.namespace = namespace
	config.jurisdiction = jurisdiction
	return config
}

func (c *Configuration) UpdateSection(section *ini.Section) {

}

func (c *Configuration) configKey(key string) string {
	return fmt.Sprint(c.namespace, "_", c.jurisdiction, "_", key)
}
