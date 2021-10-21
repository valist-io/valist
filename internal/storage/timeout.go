package storage

import (
	"context"
	"io/fs"
	"time"
)

type timeout struct {
	provider     Provider
	readTimeout  time.Duration
	writeTimeout time.Duration
}

// WithTimeout wraps a storage provider with default read and write timeouts.
func WithTimeout(provider Provider, readTimeout, writeTimeout time.Duration) Provider {
	return &timeout{provider, readTimeout, writeTimeout}
}

func (t *timeout) Open(ctx context.Context, fpath string) (File, error) {
	ctx, cancel := context.WithTimeout(ctx, t.readTimeout)
	defer cancel()
	return t.provider.Open(ctx, fpath)
}

func (t *timeout) ReadDir(ctx context.Context, fpath string) ([]fs.FileInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, t.readTimeout)
	defer cancel()
	return t.provider.ReadDir(ctx, fpath)
}

func (t *timeout) ReadFile(ctx context.Context, fpath string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, t.readTimeout)
	defer cancel()
	return t.provider.ReadFile(ctx, fpath)
}

func (t *timeout) WriteFile(ctx context.Context, fpath string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, t.writeTimeout)
	defer cancel()
	return t.provider.WriteFile(ctx, fpath)
}

func (t *timeout) Write(ctx context.Context, data []byte) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, t.writeTimeout)
	defer cancel()
	return t.provider.Write(ctx, data)
}
