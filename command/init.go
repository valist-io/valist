package command

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/valist-io/valist/prompt"
)

// Init creates a new valist config.
func Init(ctx context.Context, rpath string, wizard bool) error {
	valist := Config{
		Name:      rpath,
		Tag:       "0.0.1",
		Artifacts: make(map[string]string),
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	vpath := filepath.Join(cwd, "valist.yml")
	if err := valist.Load(vpath); err != os.ErrNotExist {
		return fmt.Errorf("project already exists: %s", vpath)
	}

	if wizard {
		// run the interactive wizard to define valist.yml
		// this should produce a config that serves as a
		// good starting point for any type of project

		multi, err := prompt.ReleaseMultiArch().Run()
		if err != nil {
			return err
		}

		switch multi[0] {
		case 'y', 'Y':
			// add supported multi arch install platforms
			valist.Artifacts["linux/amd64"] = "path_to_bin"
			valist.Artifacts["linux/arm64"] = "path_to_bin"
			valist.Artifacts["darwin/amd64"] = "path_to_bin"
			valist.Artifacts["darwin/arm64"] = "path_to_bin"
			valist.Artifacts["windows/amd64"] = "path_to_bin"
		default:
			// add some example artifacts
			valist.Artifacts["exe"] = "path_to_bin"
			valist.Artifacts["www"] = "path_to_web"
			valist.Artifacts["img"] = "path_to_img"
		}
	}

	// create will do nothing if org and repo already exist
	if err := Create(ctx, rpath); err != nil {
		return err
	}

	return valist.Save(vpath)
}
