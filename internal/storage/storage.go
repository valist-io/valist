package storage

import (
	"context"
	"io/fs"
)

type Storage interface {
	// Mkdir returns a new empty directory.
	Mkdir() Directory
	// Open opens the named file.
	Open(context.Context, string) (fs.File, error)
	// ReadDir returns a list of files in the given directory path.
	ReadDir(context.Context, string) ([]fs.FileInfo, error)
	// ReadFile reads the file with the given path.
	ReadFile(context.Context, string) ([]byte, error)
	// Write writes the given contents to a file.
	Write(context.Context, []byte) (string, error)
	// WriteFile writes the contents of the given file path.
	WriteFile(context.Context, string) (string, error)
}

type Directory interface {
	// Add adds the given file to the directory at the given path.
	Add(context.Context, string, string) error
	// Remove removes the file with the given path from the directory.
	Remove(context.Context, string) error
	// Path returns the directory path.
	Path() string
}
