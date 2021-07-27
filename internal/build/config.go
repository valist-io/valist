package build

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Define type for ValistConfig.
type ValistConfig struct {
	Type      string            `yaml:"type"`
	Org       string            `yaml:"org"`
	Repo      string            `yaml:"repo"`
	Tag       string            `yaml:"tag"`
	Meta      string            `yaml:"meta"`
	Image     string            `yaml:"image"`
	Build     string            `yaml:"build"`
	Install   string            `yaml:"install"`
	Out       string            `yaml:"out"`
	Artifacts map[string]string `yaml:"artifacts"`
}

func ParseValistConfig() ValistConfig {

	// Read yaml file from disk
	yamlFile, err := ioutil.ReadFile("valist.yml")

	// Print error if unable to read file
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
	}

	// Create valsit config object
	config := ValistConfig{}

	// Decode yaml data
	err = yaml.Unmarshal(yamlFile, &config)

	// Print error if unable to parse yaml file
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	return config
}
