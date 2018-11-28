package common

import (
	"github.com/bbucko/cli-iec/common/keys"
	"github.com/go-ini/ini"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
	"strings"
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
	jwtSection := cfg.Section("jwt")
	keysSection := cfg.Section("keys")

	jwtConfigName := iecSection.Key(configurationKey(namespace, jurisdiction, "jwtConfig")).MustString("default")
	keyName := jwtSection.Key(namedKey(jwtConfigName, "keyName")).MustString("default")
	key := NewRSAKey(keyName, keysSection)

	config = NewConfigurationWithRSAKey(namespace, jurisdiction, key)
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

func NewRSAKey(keyName string, section *ini.Section) (key *keys.RSAKey) {
	key = new(keys.RSAKey)
	key.KeyName = keyName
	key.PublicKey = section.Key(key.PublicKeySectionName()).String()
	key.PrivateKey = section.Key(key.PrivateKeySectionName()).String()
	return key
}

type Configuration struct {
	Namespace    string
	Jurisdiction string
	JwtConfig    JWTAuth
}

type JWTAuth struct {
	key *keys.RSAKey
}

func (config *JWTAuth) Key() *keys.RSAKey {
	return config.key
}

func (config *JWTAuth) ChangeKey(key *keys.RSAKey) error {
	config.key = key
	return nil
}

func NewConfigurationWithRSAKey(namespace string, jurisdiction string, key *keys.RSAKey) (config *Configuration) {
	config = new(Configuration)
	config.Namespace = namespace
	config.Jurisdiction = jurisdiction
	config.JwtConfig = JWTAuth{key: key}
	return config
}

func (c *Configuration) UpdateSection(section *ini.Section) {

}

func configurationKey(namespace string, jurisdiction, key string) string {
	return strings.Join([]string{namespace, jurisdiction, key}, "_")
}

func namedKey(name string, key string) string {
	return strings.Join([]string{name, key}, "_")
}
