package build

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/valist-io/registry/internal/core/types"
)

var DefaultImages = map[string]string{
	types.ProjectTypeBinary: "gcc:bullseye",
	types.ProjectTypeNode:   "node:buster",
	types.ProjectTypeNPM:    "node:buster",
	types.ProjectTypeGo:     "golang:buster",
	types.ProjectTypeRust:   "rust:buster",
	types.ProjectTypePython: "python:buster",
	types.ProjectTypeDocker: "scratch",
	types.ProjectTypeCPP:    "gcc:bullseye",
	types.ProjectTypeStatic: "",
}

var DefaultInstalls = map[string]string{
	types.ProjectTypeBinary: "make install",
	types.ProjectTypeNode:   "npm install",
	types.ProjectTypeNPM:    "npm install",
	types.ProjectTypeGo:     "go get ./...",
	types.ProjectTypeRust:   "cargo install",
	types.ProjectTypePython: "pip install -r requirements.txt",
	types.ProjectTypeDocker: "",
	types.ProjectTypeCPP:    "make install",
	types.ProjectTypeStatic: "",
}

var DefaultBuilds = map[string]string{
	types.ProjectTypeBinary: "make build",
	types.ProjectTypeNode:   "npm run build",
	types.ProjectTypeNPM:    "npm run build",
	types.ProjectTypeGo:     "go build",
	types.ProjectTypeRust:   "cargo build",
	types.ProjectTypePython: "python3 -m build",
	types.ProjectTypeDocker: "",
	types.ProjectTypeCPP:    "make build",
	types.ProjectTypeStatic: "",
}

var DefaultTemplates = map[string]string{
	types.ProjectTypeBinary: "default.tmpl",
	types.ProjectTypeNode:   "node.tmpl",
	types.ProjectTypeNPM:    "npm.tmpl",
	types.ProjectTypeGo:     "go.tmpl",
	types.ProjectTypeRust:   "rust.tmpl",
	types.ProjectTypePython: "python.tmpl",
	types.ProjectTypeDocker: "docker.tmpl",
	types.ProjectTypeCPP:    "cpp.tmpl",
	types.ProjectTypeStatic: "static.tmpl",
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
