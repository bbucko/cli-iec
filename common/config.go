package common

import (
	"github.com/bbucko/cli-iec/common/keys"
	"github.com/bbucko/cli-iec/ini-repo"
	"github.com/go-ini/ini"
	"github.com/urfave/cli"
	"strings"
)

func OpenConfig(c *cli.Context, namespace string, jurisdiction string) (config *Configuration, err error) {
	cfg, err := ini_repo.OpenConfig()

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
