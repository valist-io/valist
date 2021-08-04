package build

import (
	"embed"
	"os"
	"text/template"
)

//go:embed template
var configFS embed.FS

// ConfigTemplate creates a config from a project template.
func ConfigTemplate(projectType string, path string) error {
	type TemplateCfg struct {
		RenderMeta      bool
		RenderInstall   bool
		RenderPlatforms bool
		Config
	}

	cfg := TemplateCfg{}
	cfg.Type = projectType
	cfg.RenderMeta = true
	cfg.RenderPlatforms = true

	if projectType != "npm" {
		cfg.Out = "path_to_artifact_or_build_directory"
	}

	if projectType == "npm" {
		cfg.RenderMeta = false
		cfg.RenderPlatforms = false
	}

	if projectType == "static" || projectType == "go" {
		cfg.RenderInstall = false
	}

	cfg.Install = DefaultInstalls[projectType]
	cfg.Image = DefaultImages[projectType]
	cfg.Build = DefaultBuilds[projectType]

	configTemplate, err := template.New("default.tmpl").ParseFS(configFS, "template/*.tmpl")
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write template to valist.yml
	return configTemplate.Execute(f, cfg)
}
