package build

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/go-playground/validator/v10"
	"github.com/valist-io/valist/internal/core/types"
)

// Config contains valist build settings.
type Config struct {
	Name      string            `yaml:"name"      validate:"required,lowercase,shortname"`
	Tag       string            `yaml:"tag"       validate:"required,acceptable_characters"`
	Artifacts map[string]string `yaml:"artifacts" validate:"required,artifacts"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("shortname", ValidateShortname)                        //nolint:errcheck
	validate.RegisterValidation("acceptable_characters", ValidateAcceptableCharacters) //nolint:errcheck
	validate.RegisterValidation("artifacts", ValidateArtifacts)                        //nolint:errcheck
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

func ValidateArtifacts(fl validator.FieldLevel) bool {
	for iter := fl.Field().MapRange(); iter.Next(); {
		if !types.RegexPath.MatchString(iter.Value().String()) {
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

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}

	return c.Validate()
}
