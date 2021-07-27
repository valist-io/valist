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
	Image     string            `yaml:"image,omitempty"`
	Build     string            `yaml:"build,omitempty"`
	Install   string            `yaml:"install,omitempty"`
	Out       string            `yaml:"out,omitempty"`
	Artifacts map[string]string `yaml:"artifacts,omitempty"`
}

var defaultImages = map[string]string{
	"binary": "gcc:bullseye",
	"node":   "node:buster",
	"go":     "golang:buster",
	"rust":   "rust:buster",
	"python": "python:buster",
	"docker": "scratch",
	"c++":    "gcc:bullseye",
	"static": "",
}

var defaultInstalls = map[string]string{
	"binary": "make install",
	"node":   "npm install",
	"go":     "go get",
	"rust":   "cargo install",
	"python": "pip install -r requirements.txt",
	"docker": "",
	"c++":    "make install",
	"static": "",
}

var defaultBuilds = map[string]string{
	"binary": "make build",
	"node":   "npm run build",
	"go":     "go build",
	"rust":   "cargo build",
	"python": "python3 -m build",
	"docker": "",
	"c++":    "make build",
	"static": "",
}

func CreateValistConfig(
	projectType string,
	orgName string,
	repoName string,
	tag string,
	meta string,
	build string,
	install string,
	out string,
	artifacts map[string]string,
) {
	valistConfig := ValistConfig{
		Type:      projectType,
		Org:       orgName,
		Repo:      repoName,
		Tag:       tag,
		Meta:      meta,
		Build:     build,
		Install:   install,
		Out:       out,
		Artifacts: artifacts,
	}

	yamlData, err := yaml.Marshal(valistConfig)

	if err != nil {
		fmt.Printf("Could not create yaml object: %s\n", err)
	}

	ioutil.WriteFile("valist.yml", yamlData, 0644)
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
