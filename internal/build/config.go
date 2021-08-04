package build

import (
	"os"

	"gopkg.in/yaml.v3"
)

var DefaultImages = map[string]string{
	"binary": "gcc:bullseye",
	"node":   "node:buster",
	"npm":    "node:buster",
	"go":     "golang:buster",
	"rust":   "rust:buster",
	"python": "python:buster",
	"docker": "scratch",
	"c++":    "gcc:bullseye",
	"static": "",
}

var DefaultInstalls = map[string]string{
	"binary": "make install",
	"node":   "npm install",
	"npm":    "npm install",
	"go":     "go get ./...",
	"rust":   "cargo install",
	"python": "pip install -r requirements.txt",
	"docker": "",
	"c++":    "make install",
	"static": "",
}

var DefaultBuilds = map[string]string{
	"binary": "make build",
	"node":   "npm run build",
	"npm":    "npm run build",
	"go":     "go build",
	"rust":   "cargo build",
	"python": "python3 -m build",
	"docker": "",
	"c++":    "make build",
	"static": "",
}

// Config contains valist build settings.
type Config struct {
	Type      string            `yaml:"type"`
	Org       string            `yaml:"org"`
	Repo      string            `yaml:"repo"`
	Tag       string            `yaml:"tag"`
	Meta      string            `yaml:"meta,omitempty"`
	Image     string            `yaml:"image,omitempty"`
	Build     string            `yaml:"build,omitempty"`
	Install   string            `yaml:"install,omitempty"`
	Out       string            `yaml:"out,omitempty"`
	Platforms map[string]string `yaml:"artifacts,omitempty"`
}

func (c Config) Save(path string) error {
	yamlData, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(path, yamlData, 0644)
}

func (c *Config) Load(path string) error {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(yamlFile, c)
}
