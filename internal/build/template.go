package build

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed template
var templateFS embed.FS

// ConfigTemplate creates a config from a project template.
func ConfigTemplate(projectPath string, path string) error {
	templateFile, ok := DefaultTemplates[projectType]
	if !ok {
		return fmt.Errorf("Project type is not supported: %s", projectType)
	}

	cfg := Config{
		Name:      projectPath,
		Tag:       "0.0.1",
		Artifacts: make(map[string]string),
	}

	configTemplate, err := template.ParseFS(templateFS, "template/*.tmpl")
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write template to valist.yml
	return configTemplate.ExecuteTemplate(f, templateFile, cfg)
}
