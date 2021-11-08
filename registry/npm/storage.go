package npm

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/valist-io/valist/core/types"
	"github.com/valist-io/valist/storage"
)

const MetaFileName = "doc.json"

func (h *handler) fetchPackage(ctx context.Context, org, repo, tag string) (string, error) {
	res, err := h.client.ResolvePath(ctx, fmt.Sprintf("%s/%s/%s", org, repo, tag))
	if err != nil {
		return "", err
	}

	meta, err := h.client.GetReleaseMeta(ctx, res.Release.ReleaseCID)
	if err != nil {
		return "", err
	}

	artifact, ok := meta.Artifacts[MetaFileName]
	if !ok {
		return "", fmt.Errorf("artifact not found")
	}

	return artifact.Provider, nil
}

func (h *handler) latestVersions(ctx context.Context, org, repo string) (map[string]Package, error) {
	path, err := h.fetchPackage(ctx, org, repo, "latest")
	if err == types.ErrReleaseNotExist {
		return make(map[string]Package), nil
	}

	if err != nil {
		return nil, err
	}

	data, err := h.client.Storage().ReadFile(ctx, path)
	if err != nil {
		return nil, err
	}

	var meta Metadata
	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	return meta.Versions, nil
}

func (h *handler) loadPackage(ctx context.Context, org, repo, tag string) (storage.File, error) {
	path, err := h.fetchPackage(ctx, org, repo, tag)
	if err != nil {
		return nil, err
	}

	return h.client.Storage().Open(ctx, path)
}

func (h *handler) writeAttachment(ctx context.Context, data string) (string, error) {
	var tarData bytes.Buffer
	buf := bytes.NewBufferString(data)
	dec := base64.NewDecoder(base64.StdEncoding, buf)

	if _, err := io.Copy(&tarData, dec); err != nil {
		return "", err
	}

	return h.client.Storage().Write(ctx, tarData.Bytes())
}
