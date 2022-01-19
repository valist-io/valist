package publish

import (
	"os"
	"regexp"

	"gopkg.in/yaml.v3"

	"github.com/go-playground/validator/v10"
)

var (
	RegexShortname            = regexp.MustCompile(`^[0-9a-z-_]+$`)
	RegexPath                 = regexp.MustCompile(`^[0-9A-z\-_\/\.]+$`)
	RegexAcceptableCharacters = regexp.MustCompile(`^[0-9A-z-_\\/\. ]*$`)
)

var validate *validator.Validate

//nolint:errcheck
func init() {
	validate = validator.New()
	validate.RegisterValidation("shortname", func(fl validator.FieldLevel) bool {
		return RegexPath.MatchString(fl.Field().String())
	})
	validate.RegisterValidation("acceptable_characters", func(fl validator.FieldLevel) bool {
		return RegexAcceptableCharacters.MatchString(fl.Field().String())
	})
}

// Config contains project settings.
type Config struct {
	Name      string            `yaml:"name"      validate:"required,lowercase,shortname"`
	Tag       string            `yaml:"tag"       validate:"required,acceptable_characters"`
	Artifacts map[string]string `yaml:"artifacts" validate:"required,dive,keys,shortname,endkeys,shortname"`
}

// Validate returns any validation errors.
func (c Config) Validate() error {
	return validate.Struct(c)
}

func (c Config) Save(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	if err = c.Validate(); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (c *Config) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}
	return c.Validate()
}
