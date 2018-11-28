package ini_repo

import (
	"github.com/go-ini/ini"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func OpenConfig() (config *ini.File, err error) {
	cfg, err := openFile()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = ensureSection(cfg, "iec")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = ensureSection(cfg, "jwt")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = ensureSection(cfg, "keys")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = saveFile(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	return cfg, nil
}

func Persist(sectionName string, key string, value string) (err error) {
	cfg, err := openFile()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = ensureSection(cfg, sectionName)
	if err != nil {
		log.Fatal(err)
		return
	}
	section, err := cfg.GetSection(sectionName)
	if err != nil {
		log.Fatal(err)
		return
	}
	section.NewKey(key, value)
	saveFile(cfg)
	return
}

func GetValue(sectionName string, keyName string) (value string, err error) {
	cfg, err := openFile()
	if err != nil {
		log.Fatal(err)
		return
	}
	section, err := cfg.GetSection(sectionName)
	if err != nil {
		log.Fatal(err)
		return
	}

	key, err := section.GetKey(keyName)
	if err != nil {
		log.Fatal(err)
		return
	}
	return key.Value(), nil
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

func saveFile(cfg *ini.File) (err error) {
	configPath := cliConfigPath()
	err = cfg.SaveTo(configPath)
	if err != nil {
		return
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
	ensurePathExists(configPath)
	configFile, err = ini.Load(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	return configFile, nil
}

func ensurePathExists(configPath string) (err error) {
	path := strings.Split(configPath, "/")
	err = os.MkdirAll(strings.Join(path[:len(path)-1], "/"), 0777)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = os.OpenFile(configPath, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
