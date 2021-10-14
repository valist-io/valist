package git

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"

	billy "github.com/go-git/go-billy/v5"

	"github.com/valist-io/valist/internal/storage"
)

var (
	_ billy.File       = (*storageFile)(nil)
	_ billy.Filesystem = (*storageFS)(nil)
)

var ErrReadOnly = errors.New("read-only filesystem")

type storageFile struct {
	storage.File
	name string
}

func (f *storageFile) Name() string {
	return f.name
}

func (f *storageFile) Lock() error {
	return nil
}

func (f *storageFile) Unlock() error {
	return nil
}

func (f *storageFile) ReadAt(b []byte, off int64) (n int, err error) {
	return 0, fmt.Errorf("ReadAt not supported")
}

func (f *storageFile) Truncate(size int64) error {
	return ErrReadOnly
}

func (f *storageFile) Write(p []byte) (n int, err error) {
	return 0, ErrReadOnly
}

type storageFS struct {
	storage storage.Provider
	root    string
}

func (fs *storageFS) Open(filename string) (billy.File, error) {
	f, err := fs.storage.Open(context.Background(), fs.Join(fs.root, filename))
	if err != nil {
		return nil, err
	}

	return &storageFile{f, filename}, nil
}

func (fs *storageFS) Stat(filename string) (os.FileInfo, error) {
	f, err := fs.storage.Open(context.Background(), fs.Join(fs.root, filename))
	if err != nil {
		return nil, err
	}

	return f.Stat()
}

func (fs *storageFS) Lstat(filename string) (os.FileInfo, error) {
	return fs.Stat(filename)
}

func (fs *storageFS) OpenFile(filename string, flag int, perm os.FileMode) (billy.File, error) {
	return fs.Open(filename)
}

func (fs *storageFS) ReadDir(filename string) ([]os.FileInfo, error) {
	return fs.storage.ReadDir(context.Background(), fs.Join(fs.root, filename))
}

func (fs *storageFS) Join(elem ...string) string {
	return path.Join(elem...)
}

func (fs *storageFS) Create(filename string) (billy.File, error) {
	return nil, ErrReadOnly
}

func (fs *storageFS) Rename(oldpath, newpath string) error {
	return ErrReadOnly
}

func (fs *storageFS) Remove(filename string) error {
	return ErrReadOnly
}

func (fs *storageFS) TempFile(dir, prefix string) (billy.File, error) {
	return nil, ErrReadOnly
}

func (fs *storageFS) MkdirAll(filename string, perm os.FileMode) error {
	return ErrReadOnly
}

func (fs *storageFS) Symlink(target, link string) error {
	return ErrReadOnly
}

func (fs *storageFS) Readlink(link string) (string, error) {
	return "", nil
}

func (fs *storageFS) Chroot(path string) (billy.Filesystem, error) {
	return nil, fmt.Errorf("Chroot not supported")
}

func (fs *storageFS) Root() string {
	return ""
}
