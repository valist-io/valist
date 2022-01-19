package publish

import (
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

// Readme returns the contents of the project readme file if it exists.
func Readme(cwd string) ([]byte, error) {
	// TODO replace with regex or path matcher
	return os.ReadFile("README.md")
}

// Dependencies returns a list of dependencies for the project.
func Dependencies(cwd string) ([]string, error) {
	data, err := os.ReadFile(filepath.Join(cwd, "go.mod"))
	if err != nil {
		return nil, err
	}
	mod, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return nil, err
	}
	var dependencies []string
	for _, url := range mod.Require {
		dependencies = append(dependencies, url.Mod.String())
	}
	return dependencies, nil
}
