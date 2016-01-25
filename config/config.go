package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gregarmer/r53dyndns/utils"
	"log"
	"os"
	"os/user"
	"strings"
)

const configFile = ".r53dyndns"

type Config struct {
	AwsAccessKey string `json:"aws_access_key"`
	AwsSecretKey string `json:"aws_secret_key"`
	ZoneId       string `json:"zone_id"`
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

	if c.ZoneId == "" {
		return errors.New("missing ZoneId, cannot continue")
	}

	return nil
}

func LoadConfig(configFile string) *Config {
	// expand tilde to users home dir
	if strings.HasPrefix(configFile, "~") {
		u, _ := user.Current()
		configFile = strings.Replace(configFile, "~", u.HomeDir, 1)
	}

	log.Printf("config file is %s", configFile)

	// make sure the config file actually exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		utils.Fatalf("Error: couldn't load config from %s\n\n"+
			"If you haven't configured r53dyndns yet, place a "+
			"file that looks like this in ~/.r53dyndns:\n\n%s",
			configFile, SampleConfig())
	}

	// load config
	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		utils.Fatalf("error: %s", err)
	}

	// pre-flight check
	utils.CheckErr(config.PreFlight())

	return &config
}

func SampleConfig() string {
	config := Config{}
	fileJson, _ := json.MarshalIndent(config, "", "  ")
	return fmt.Sprintf("%s", fileJson)
}
