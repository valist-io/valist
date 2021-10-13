package storage

import (
	"context"
	"io"
	"io/fs"
)

type File interface {
	io.Reader
	io.Closer
	io.Seeker
	// Stat returns FileInfo.
	Stat() (fs.FileInfo, error)
}

type Provider interface {
	// Open returns the named file.
	Open(context.Context, string) (File, error)
	// ReadDir returns a list of files in the given directory path.
	ReadDir(context.Context, string) ([]fs.FileInfo, error)
	// ReadFile reads the file with the given path.
	ReadFile(context.Context, string) ([]byte, error)
	// WriteFile writes the contents of the given file path.
	WriteFile(context.Context, string) (string, error)
	// Write writes the given contents to a file.
	Write(context.Context, []byte) (string, error)
}
