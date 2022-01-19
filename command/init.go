package command

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/valist-io/valist/prompt"
	"github.com/valist-io/valist/publish"
)

var ErrProjectExist = errors.New("valist.yml already exists")

// Init creates a new valist config.
func Init(ctx context.Context, rpath string, wizard bool) error {
	pub := publish.Config{
		Name:      rpath,
		Tag:       "0.0.1",
		Artifacts: make(map[string]string),
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	vpath := filepath.Join(cwd, "valist.yml")
	if err := pub.Load(vpath); err == nil {
		return ErrProjectExist
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if wizard {
		// run the interactive wizard to define valist.yml
		// this should produce a config that serves as a
		// good starting point for any type of project

		multi, err := prompt.ReleaseMultiArch().Run()
		if err != nil {
			return err
		}

		switch multi {
		case "y", "Y":
			// add supported multi arch install platforms
			pub.Artifacts["linux/amd64"] = "path_to_bin"
			pub.Artifacts["linux/arm64"] = "path_to_bin"
			pub.Artifacts["darwin/amd64"] = "path_to_bin"
			pub.Artifacts["darwin/arm64"] = "path_to_bin"
			pub.Artifacts["windows/amd64"] = "path_to_bin"
		default:
			// add some example artifacts
			pub.Artifacts["exe"] = "path_to_bin"
			pub.Artifacts["www"] = "path_to_web"
			pub.Artifacts["img"] = "path_to_img"
		}
	}

	return pub.Save(vpath)
}
