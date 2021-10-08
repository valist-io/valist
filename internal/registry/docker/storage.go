package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/valist-io/valist/internal/core/types"
	"github.com/valist-io/valist/internal/storage"
)

func (h *handler) writeBlob(uuid string, r io.Reader) error {
	path := filepath.Join(os.TempDir(), uuid, "blob")

	blob, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer blob.Close()

	size, err := io.Copy(blob, r)
	if err != nil {
		return err
	}

	h.uploads[uuid] += size
	return nil
}

func (h *handler) findBlob(ctx context.Context, orgName, repoName, digest string) ([]string, error) {
	if p, ok := h.blobs[digest]; ok {
		return p, nil
	}

	raw := fmt.Sprintf("%s/%s/latest", orgName, repoName)
	res, err := h.client.ResolvePath(ctx, raw)
	if err != nil {
		return nil, err
	}

	data, err := h.client.Storage().ReadFile(ctx, res.Release.ReleaseCID)
	if err != nil {
		return nil, err
	}

	var meta types.ReleaseMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	artifact, ok := meta.Artifacts[digest]
	if !ok {
		return nil, fmt.Errorf("artifact not found")
	}

	return artifact.Providers, nil
}

func (h *handler) loadBlob(ctx context.Context, orgName, repoName, digest string) (storage.File, error) {
	p, err := h.findBlob(ctx, orgName, repoName, digest)
	if err != nil {
		return nil, err
	}

	return h.client.Storage().Open(ctx, p...)
}

func (h *handler) loadManifest(ctx context.Context, orgName, repoName, ref string) (storage.File, error) {
	raw := fmt.Sprintf("%s/%s/%s", orgName, repoName, ref)
	res, err := h.client.ResolvePath(ctx, raw)
	if err != nil {
		return nil, err
	}

	data, err := h.client.Storage().ReadFile(ctx, res.Release.ReleaseCID)
	if err != nil {
		return nil, err
	}

	var meta types.ReleaseMeta
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	artifact, ok := meta.Artifacts[ref]
	if !ok {
		return nil, fmt.Errorf("artifact not found")
	}

	return h.client.Storage().Open(ctx, artifact.Providers...)
}
