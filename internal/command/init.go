package command

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/valist-io/valist/internal/core/client"
	"github.com/valist-io/valist/internal/core/types"
)

// Init creates a new valist config.
func Init(ctx context.Context, rpath string) error {
	client := ctx.Value(ClientKey).(*client.Client)

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

	_, err = client.ResolvePath(ctx, rpath)
	if err == types.ErrRepoNotExist || err == types.ErrOrgNotExist {
		err = Create(ctx, rpath)
	}

	if err != nil {
		return err
	}

	return valist.Save(vpath)
}
