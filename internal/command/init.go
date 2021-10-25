package command

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// Init creates a new valist config.
func Init(ctx context.Context, rpath string) error {
	valist := Config{
		Name:      rpath,
		Tag:       "0.0.1",
		Artifacts: map[string]string{"linux/amd64": "path_to_bin"},
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	vpath := filepath.Join(cwd, "valist.yml")
	if err := valist.Load(vpath); err != os.ErrNotExist {
		return fmt.Errorf("project already exists: %s", vpath)
	}

	// create will do nothing if org and repo already exist
	if err := Create(ctx, rpath); err != nil {
		return err
	}

	return valist.Save(vpath)
}
