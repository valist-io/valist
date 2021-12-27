package command

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/valist-io/valist/core/client"
)

// Get downloads a binary artifact.
func Get(ctx context.Context, rpath, apath, opath string) error {
	client := ctx.Value(ClientKey).(*client.Client)

	if strings.Count(rpath, "/") < 2 {
		rpath += "/latest"
	}

	res, err := client.ResolvePath(ctx, rpath)
	if err != nil {
		return err
	}

	if res.Release == nil {
		return fmt.Errorf("invalid release path: %s", rpath)
	}

	fmt.Println("Fetching from distributed storage...")
	releaseMeta, err := client.GetReleaseMeta(ctx, res.Release.ReleaseCID)
	if err != nil {
		return err
	}

	// default to system platform if no artifact specified
	if apath == "" {
		apath = runtime.GOOS + "/" + runtime.GOARCH
	}

	artifact, ok := releaseMeta.Artifacts[apath]
	if !ok {
		return fmt.Errorf("%s not found in release", apath)
	}

	data, err := client.ReadFile(ctx, artifact.Provider)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// default to current directory if no output specified
	if opath == "" {
		opath = filepath.Join(cwd, strings.ReplaceAll(apath, string(filepath.Separator), "-"))
	}

	return os.WriteFile(opath, data, 0744)
}
