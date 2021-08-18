package build

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/valist-io/registry/internal/core"
)

var DefaultImages = map[string]string{
	core.ProjectTypeBinary: "gcc:bullseye",
	core.ProjectTypeNode:   "node:buster",
	core.ProjectTypeNPM:    "node:buster",
	core.ProjectTypeGo:     "golang:buster",
	core.ProjectTypeRust:   "rust:buster",
	core.ProjectTypePython: "python:buster",
	core.ProjectTypeDocker: "scratch",
	core.ProjectTypeCPP:    "gcc:bullseye",
	core.ProjectTypeStatic: "",
}

var DefaultInstalls = map[string]string{
	core.ProjectTypeBinary: "make install",
	core.ProjectTypeNode:   "npm install",
	core.ProjectTypeNPM:    "npm install",
	core.ProjectTypeGo:     "go get ./...",
	core.ProjectTypeRust:   "cargo install",
	core.ProjectTypePython: "pip install -r requirements.txt",
	core.ProjectTypeDocker: "",
	core.ProjectTypeCPP:    "make install",
	core.ProjectTypeStatic: "",
}

var DefaultBuilds = map[string]string{
	core.ProjectTypeBinary: "make build",
	core.ProjectTypeNode:   "npm run build",
	core.ProjectTypeNPM:    "npm run build",
	core.ProjectTypeGo:     "go build",
	core.ProjectTypeRust:   "cargo build",
	core.ProjectTypePython: "python3 -m build",
	core.ProjectTypeDocker: "",
	core.ProjectTypeCPP:    "make build",
	core.ProjectTypeStatic: "",
}

var DefaultTemplates = map[string]string{
	core.ProjectTypeBinary: "default.tmpl",
	core.ProjectTypeNode:   "node.tmpl",
	core.ProjectTypeNPM:    "npm.tmpl",
	core.ProjectTypeGo:     "go.tmpl",
	core.ProjectTypeRust:   "rust.tmpl",
	core.ProjectTypePython: "python.tmpl",
	core.ProjectTypeDocker: "docker.tmpl",
	core.ProjectTypeCPP:    "cpp.tmpl",
	core.ProjectTypeStatic: "static.tmpl",
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
