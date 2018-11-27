package common

import (
	"fmt"
	"github.com/bbucko/cli-iec/jwt"
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

	err = ensureSection(cfg, "keys")
	if err != nil {
		return
	}

	err = saveFile(cfg)
	if err != nil {
		return
	}

	iecSection := cfg.Section("iec")
	_ = cfg.Section("jwt")
	keysSection := cfg.Section("keys")

	config = NewConfiguration(namespace, jurisdiction)
	keyName := iecSection.Key(config.configKey("jwtKeyName")).MustString("default")
	config.key = NewRSAKey(keyName, keysSection)
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

func NewRSAKey(keyName string, section *ini.Section) (key *jwt.RSAKey) {
	key = new(jwt.RSAKey)
	key.KeyName = keyName
	key.PublicKey = section.Key(key.PublicKeySectionName()).String()
	key.PrivateKey = section.Key(key.PrivateKeySectionName()).String()
	return key
}

type Configuration struct {
	namespace    string
	jurisdiction string
	key          *jwt.RSAKey
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
