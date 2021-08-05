package build

import (
	"encoding/json"
	"os"
)

type PackageJson struct {
	Name    string
	Version string
}

func ParsePackageJson(path string) (PackageJson, error) {
	var packageJson PackageJson

	data, err := os.ReadFile(path)
	if err != nil {
		return packageJson, err
	}

	if err := json.Unmarshal(data, &packageJson); err != nil {
		return packageJson, err
	}

	return packageJson, nil
}
