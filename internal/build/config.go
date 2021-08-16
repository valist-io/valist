package build

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	ProjectTypeBinary = "binary"
	ProjectTypeNode   = "node"
	ProjectTypeNPM    = "npm"
	ProjectTypeGo     = "go"
	ProjectTypeRust   = "rust"
	ProjectTypePython = "python"
	ProjectTypeDocker = "docker"
	ProjectTypeCPP    = "c++"
	ProjectTypeStatic = "static"
)

var DefaultImages = map[string]string{
	ProjectTypeBinary: "gcc:bullseye",
	ProjectTypeNode:   "node:buster",
	ProjectTypeNPM:    "node:buster",
	ProjectTypeGo:     "golang:buster",
	ProjectTypeRust:   "rust:buster",
	ProjectTypePython: "python:buster",
	ProjectTypeDocker: "scratch",
	ProjectTypeCPP:    "gcc:bullseye",
	ProjectTypeStatic: "",
}

var DefaultInstalls = map[string]string{
	ProjectTypeBinary: "make install",
	ProjectTypeNode:   "npm install",
	ProjectTypeNPM:    "npm install",
	ProjectTypeGo:     "go get ./...",
	ProjectTypeRust:   "cargo install",
	ProjectTypePython: "pip install -r requirements.txt",
	ProjectTypeDocker: "",
	ProjectTypeCPP:    "make install",
	ProjectTypeStatic: "",
}

var DefaultBuilds = map[string]string{
	ProjectTypeBinary: "make build",
	ProjectTypeNode:   "npm run build",
	ProjectTypeNPM:    "npm run build",
	ProjectTypeGo:     "go build",
	ProjectTypeRust:   "cargo build",
	ProjectTypePython: "python3 -m build",
	ProjectTypeDocker: "",
	ProjectTypeCPP:    "make build",
	ProjectTypeStatic: "",
}

var DefaultTemplates = map[string]string{
	ProjectTypeBinary: "default.tmpl",
	ProjectTypeNode:   "node.tmpl",
	ProjectTypeNPM:    "npm.tmpl",
	ProjectTypeGo:     "go.tmpl",
	ProjectTypeRust:   "rust.tmpl",
	ProjectTypePython: "python.tmpl",
	ProjectTypeDocker: "docker.tmpl",
	ProjectTypeCPP:    "cpp.tmpl",
	ProjectTypeStatic: "static.tmpl",
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
	Platforms map[string]string `yaml:"platforms,omitempty"`
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
