package common

import (
	"fmt"
	"github.com/bbucko/cli-iec/ini-repo"
	"github.com/bbucko/cli-iec/keys"
	"github.com/go-ini/ini"
	"github.com/urfave/cli"
)

func OpenConfig(c *cli.Context, namespace string, jurisdiction string) (config *Configuration, err error) {
	cfg, err := ini_repo.OpenConfig()

	iecSection := cfg.Section("iec")
	_ = cfg.Section("jwt")
	keysSection := cfg.Section("keys")

	config = NewConfiguration(namespace, jurisdiction)
	keyName := iecSection.Key(config.configKey("jwtKeyName")).MustString("default")
	config.key = NewRSAKey(keyName, keysSection)
	return
}

func SaveConfig(c *cli.Context, configuration *Configuration) (err error) {
	cfg, err := ini_repo.OpenConfig()
	if err != nil {
		return
	}
	section := cfg.Section("iec")
	configuration.UpdateSection(section)
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
	namespace    string
	jurisdiction string
	key          *keys.RSAKey
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
