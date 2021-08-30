package git

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	billy "github.com/go-git/go-billy/v5"
	files "github.com/ipfs/go-ipfs-files"
)

var (
	_ billy.File       = (*file)(nil)
	_ billy.Filesystem = (*filesystem)(nil)
)

var ErrReadOnly = errors.New("read-only filesystem")

type file struct {
	files.File
}

func (f *file) Name() string {
	return ""
}

func (f *file) Lock() error {
	return nil
}

func (f *file) Unlock() error {
	return nil
}

func (f *file) ReadAt(b []byte, off int64) (n int, err error) {
	return 0, fmt.Errorf("ReadAt not supported")
}

func (f *file) Truncate(size int64) error {
	return ErrReadOnly
}

func (f *file) Write(p []byte) (n int, err error) {
	return 0, ErrReadOnly
}

type filesystem struct{}

func (fs *filesystem) Create(filename string) (billy.File, error) {
	return nil, ErrReadOnly
}

func (fs *filesystem) Open(filename string) (billy.File, error) {
	return nil, nil
}

func (fs *filesystem) OpenFile(filename string, flag int, perm os.FileMode) (billy.File, error) {
	return fs.Open(filename)
}

func (fs *filesystem) Stat(filename string) (os.FileInfo, error) {
	return nil, nil
}

func (fs *filesystem) Rename(oldpath, newpath string) error {
	return ErrReadOnly
}

func (fs *filesystem) Remove(filename string) error {
	return ErrReadOnly
}

func (fs *filesystem) Join(elem ...string) string {
	return filepath.Join(elem...)
}

func (fs *filesystem) TempFile(dir, prefix string) (billy.File, error) {
	return nil, ErrReadOnly
}

func (fs *filesystem) ReadDir(path string) ([]os.FileInfo, error) {
	return nil, nil
}

func (fs *filesystem) MkdirAll(filename string, perm os.FileMode) error {
	return ErrReadOnly
}

func (fs *filesystem) Lstat(filename string) (os.FileInfo, error) {
	return nil, nil
}

func (fs *filesystem) Symlink(target, link string) error {
	return ErrReadOnly
}

func (fs *filesystem) Readlink(link string) (string, error) {
	return "", nil
}

func (fs *filesystem) Chroot(path string) (billy.Filesystem, error) {
	return nil, nil
}

func (fs *filesystem) Root() string {
	return ""
}
