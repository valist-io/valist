package build

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/go-playground/validator/v10"
	"github.com/valist-io/valist/internal/core/types"
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
	Type      string            `yaml:"type" validate:"required,alphanum"`
	Org       string            `yaml:"org" validate:"required,alphanum"`
	Repo      string            `yaml:"repo" validate:"required,alphanum"`
	Tag       string            `yaml:"tag" validate:"required,printascii"`
	Meta      string            `yaml:"meta,omitempty" validate:"printascii"`
	Image     string            `yaml:"image,omitempty" validate:"printascii"`
	Build     string            `yaml:"build,omitempty" validate:"printascii"`
	Install   string            `yaml:"install,omitempty" validate:"printascii"`
	Out       string            `yaml:"out,omitempty" validate:"printascii"`
	Platforms map[string]string `yaml:"platforms,omitempty" validate:"printascii"`
}

var validate *validator.Validate

func (c Config) Validate() error {
	if validate == nil {
		validate = validator.New()
		validate.RegisterValidation("platforms", ValidatePlatforms)
	}
	return validate.Struct(c)
}

func ValidatePlatforms(fl validator.FieldLevel) bool {
	iter := fl.Field().MapRange()
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		err := validate.Var(key, "required,printascii")
		if err != nil {
			return false
		}

		err = validate.Var(value, "required,printascii")
		if err != nil {
			return false
		}
	}
	return true
}

func (c Config) Save(path string) error {
	yamlData, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	if err = c.Validate(); err != nil {
		return err
	}

	return os.WriteFile(path, yamlData, 0644)
}

func (c *Config) Load(path string) error {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}

	return c.Validate()
}
