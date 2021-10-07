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
func ConfigTemplate(projectType string, path string) error {
	templateFile, ok := DefaultTemplates[projectType]
	if !ok {
		return fmt.Errorf("Project type is not supported: %s", projectType)
	}

	cfg := Config{
		Org:     "test",
		Repo:    "test",
		Tag:     "test",
		Type:    projectType,
		Out:     "path_to_artifact_or_build_directory",
		Install: DefaultInstalls[projectType],
		Image:   DefaultImages[projectType],
		Build:   DefaultBuilds[projectType],
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
