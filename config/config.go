package config

import (
	"os"
	"log"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Username string
	Password string
	Connection string
}

// Reads info from config file
func ReadConfig() Config {
	var configfile = "config/config.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}