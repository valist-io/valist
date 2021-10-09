package storage

import (
	"context"
	"errors"
	"io/fs"
	"strings"
)

var (
	ErrInvalidPath = errors.New("invalid storage path")
	ErrNoProvider  = errors.New("storage provider not found")
)

type Storage struct {
	providers []Provider
}

// NewStorage returns a new storage manager with the given providers.
// Providers are priority sorted in the given order.
func NewStorage(providers ...Provider) (*Storage, error) {
	if len(providers) == 0 {
		return nil, ErrNoProvider
	}

	return &Storage{providers}, nil
}

// Provider returns the provider with the matching prefix.
func (s *Storage) Provider(prefix string) (Provider, error) {
	for _, provider := range s.providers {
		if provider.Prefix() == prefix {
			return provider, nil
		}
	}

	return nil, ErrNoProvider
}

// Open returns the named file.
func (s *Storage) Open(ctx context.Context, fpath string) (File, error) {
	parts := strings.Split(fpath, "/")
	if len(parts) < 2 {
		return nil, ErrNoProvider
	}

	provider, err := s.Provider(parts[1])
	if err != nil {
		return nil, err
	}

	return provider.Open(ctx, fpath)
}

// ReadDir returns a list of files in the given directory path.
func (s *Storage) ReadDir(ctx context.Context, fpath string) ([]fs.FileInfo, error) {
	parts := strings.Split(fpath, "/")
	if len(parts) < 2 {
		return nil, ErrNoProvider
	}

	provider, err := s.Provider(parts[1])
	if err != nil {
		return nil, err
	}

	return provider.ReadDir(ctx, fpath)
}

// ReadFile returns the contents of the file with the given path.
func (s *Storage) ReadFile(ctx context.Context, fpath string) ([]byte, error) {
	parts := strings.Split(fpath, "/")
	if len(parts) < 2 {
		return nil, ErrInvalidPath
	}

	provider, err := s.Provider(parts[1])
	if err != nil {
		return nil, err
	}

	return provider.ReadFile(ctx, fpath)
}

// WriteFile writes the contents of the given file path for all providers.
func (s *Storage) WriteFile(ctx context.Context, fpath string) ([]string, error) {
	if len(s.providers) == 0 {
		return nil, ErrNoProvider
	}

	set := make(map[string]bool)
	for _, provider := range s.providers {
		path, err := provider.WriteFile(ctx, fpath)
		if err != nil {
			return nil, err
		}
		set[path] = true
	}

	var paths []string
	for path := range set {
		paths = append(paths, path)
	}

	return paths, nil
}

// Write writes the given contents to a file for all providers.
func (s *Storage) Write(ctx context.Context, data []byte) ([]string, error) {
	if len(s.providers) == 0 {
		return nil, ErrNoProvider
	}

	set := make(map[string]bool)
	for _, provider := range s.providers {
		path, err := provider.Write(ctx, data)
		if err != nil {
			return nil, err
		}
		set[path] = true
	}

	var paths []string
	for path := range set {
		paths = append(paths, path)
	}

	return paths, nil
}
