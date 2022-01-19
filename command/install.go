package command

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/valist-io/valist/core"
)

// Install downloads a binary artifact to the valist bin path.
func Install(ctx context.Context, rpath string) error {
	client := ctx.Value(ClientKey).(*core.Client)
	config := ctx.Value(ConfigKey).(*core.Config)

	if strings.Count(rpath, "/") < 2 {
		rpath += "/latest"
	}

	logger.Notice("Fetching from distributed storage...")
	res, err := client.ResolvePath(ctx, rpath)
	if err != nil {
		return err
	}
	if res.Release == nil {
		return fmt.Errorf("invalid release path: %s", rpath)
	}

	platform := runtime.GOOS + "/" + runtime.GOARCH
	artifact, ok := res.Release.Artifacts[platform]
	if !ok {
		return fmt.Errorf("target platform %s not found in release", platform)
	}
	data, err := client.ReadFile(ctx, artifact.Provider)
	if err != nil {
		return err
	}

	binPath := config.InstallPath()
	exePath := filepath.Join(binPath, res.ProjectName)

	logger.Info("Installing for target platform %s", platform)
	if err := os.MkdirAll(binPath, 0744); err != nil {
		return err
	}
	if err := os.WriteFile(exePath, data, 0744); err != nil {
		return err
	}

	logger.Info("Successfully installed: %s\n", exePath)
	if !strings.Contains(os.Getenv("PATH"), binPath) {
		logger.Info("%s not detected in $PATH. Add to path or run:", binPath)
		logger.Info(`export PATH="$PATH:%s"`, binPath)
	}

	return nil
}
