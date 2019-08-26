package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// XDConfig is the configuration
type XDConfig struct {
	Scripts map[string]string `json:"scripts" yaml:"scripts"`
}

// ReadConfig read the XD config
func ReadConfig() (*XDConfig, error) {
	configFile, err := findConfigFile()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read configuration file '%s': %v", configFile, err)
	}

	config := XDConfig{}

	if strings.HasSuffix(configFile, ".json") {
		err = json.Unmarshal(b, &config)
	} else {
		err = yaml.Unmarshal(b, &config)
	}

	if err != nil {
		return nil, fmt.Errorf("invalid format in configuration file '%s': %v", configFile, err)
	}
	return &config, nil
}

func findConfigFile() (string, error) {
	if fileExists(".xd.yml") {
		return ".xd.yml", nil
	}
	if fileExists(".xd.yaml") {
		return ".xd.yaml", nil
	}
	if fileExists(".xd.json") {
		return ".xd.json", nil
	}
	return "", fmt.Errorf("configuration file not found in current directoy")
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
