package config

import (
	"encoding/json"
	"errors"
	"github.com/gregarmer/r53dyndns/utils"
	"os"
	"os/user"
	"path/filepath"
)

const configFile = ".r53dyndns"

type Config struct {
	AwsAccessKey string `json:"aws_access_key"`
	AwsSecretKey string `json:"aws_secret_key"`
}

func (c *Config) Copy() Config {
	return *c
}

func (c *Config) PreFlight() error {
	if c.AwsAccessKey == "" {
		return errors.New("missing AwsAccessKey, cannot continue")
	}

	if c.AwsSecretKey == "" {
		return errors.New("missing AwsSecretKey, cannot continue")
	}

	return nil
}

func GetConfigPath() string {
	u, _ := user.Current()
	return filepath.Join(u.HomeDir, configFile)
}

func LoadConfig(configFile string) *Config {
	// make sure the config file actually exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		utils.Fatalf("couldn't load config from %s", configFile)
	}

	// init config if needed
	defaultConfigPath := GetConfigPath()
	if configFile == defaultConfigPath {
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			utils.Fatalf("please create a config file containing your AWS creds.")
		}
	}

	// load config
	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		utils.Fatalf("error: %s", err)
	}

	// pre-flight check (s3 keys, access to postgres etc)
	utils.CheckErr(config.PreFlight())

	return &config
}
