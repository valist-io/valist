package build

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/go-playground/validator/v10"
	"github.com/valist-io/valist/internal/core/types"
)

// Config contains valist build settings.
type Config struct {
	Type      string            `yaml:"type,omitempty" validate:"project_type"`
	Name      string            `yaml:"name" validate:"required,lowercase,shortname"`
	Tag       string            `yaml:"tag" validate:"required,acceptable_characters"`
	Artifacts map[string]string `yaml:"artifacts" validate:"platforms"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("shortname", ValidateShortname)                        //nolint:errcheck
	validate.RegisterValidation("acceptable_characters", ValidateAcceptableCharacters) //nolint:errcheck
	validate.RegisterValidation("project_type", ValidateProjectType)                   //nolint:errcheck
	validate.RegisterValidation("platforms", ValidateArtifacts)                        //nolint:errcheck
}

func (c Config) Validate() error {
	return validate.Struct(c)
}

func ValidateShortname(fl validator.FieldLevel) bool {
	return types.RegexPath.MatchString(fl.Field().String())
}

func ValidateAcceptableCharacters(fl validator.FieldLevel) bool {
	return types.RegexAcceptableCharacters.MatchString(fl.Field().String())
}

func ValidateProjectType(fl validator.FieldLevel) bool {
	for _, projectType := range types.ProjectTypes {
		if projectType == fl.Field().String() || fl.Field().String() == "" {
			return true
		}
	}
	return false
}

func ValidateArtifacts(fl validator.FieldLevel) bool {
	valid := true
	for iter := fl.Field().MapRange(); iter.Next(); {
		valid = types.RegexPlatformArchitecture.MatchString(iter.Key().String()) // linux/amd64
		if !valid {
			fmt.Println("Invalid os/arch in platforms")
			break
		}

		valid = types.RegexPath.MatchString(iter.Value().String()) // bin/linux/amd64/valist
		if !valid {
			fmt.Println("Invalid path to artifact")
			break
		}
	}
	return valid
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

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}

	return c.Validate()
}
