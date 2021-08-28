package docker

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	cid "github.com/ipfs/go-cid"
	files "github.com/ipfs/go-ipfs-files"

	"github.com/valist-io/registry/internal/core/types"
)

type Storage struct {
	client  types.CoreAPI
	blobs   map[string]cid.Cid
	uploads map[string]int64
}

func NewStorage(client types.CoreAPI) *Storage {
	return &Storage{
		client:  client,
		blobs:   make(map[string]cid.Cid),
		uploads: make(map[string]int64),
	}
}

func (s *Storage) StartUpload() (string, error) {
	path, err := os.MkdirTemp("", "")
	if err != nil {
		return "", err
	}

	return filepath.Base(path), nil
}

func (s *Storage) WriteUpload(uuid string, r io.Reader) (int64, error) {
	path := filepath.Join(os.TempDir(), uuid, "blob")

	blob, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer blob.Close()

	size, err := io.Copy(blob, r)
	if err != nil {
		return 0, err
	}

	s.uploads[uuid] += size
	return size, nil
}

func (s *Storage) FinishUpload(ctx context.Context, uuid, digest string) error {
	path := filepath.Join(os.TempDir(), uuid, "blob")

	id, err := s.client.WriteFilePath(ctx, path)
	if err != nil {
		return err
	}
	defer os.RemoveAll(filepath.Dir(path))

	s.blobs[digest] = id
	delete(s.uploads, uuid)

	return nil
}

func (s *Storage) GetBlob(ctx context.Context, orgName, repoName, digest string) (files.File, error) {
	if id, ok := s.blobs[digest]; ok {
		return s.client.GetFile(ctx, id)
	}

	raw := fmt.Sprintf("%s/%s/latest/blobs/%s", orgName, repoName, digest)
	res, err := s.client.ResolvePath(ctx, raw)
	if err != nil {
		return nil, err
	}

	file, ok := res.File.(files.File)
	if !ok {
		return nil, fmt.Errorf("not found")
	}

	return file, nil
}
