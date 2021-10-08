package ipfs

import (
	"io/fs"
	"time"

	files "github.com/ipfs/go-ipfs-files"
)

type file struct {
	name string
	file files.File
}

func (f *file) Stat() (fs.FileInfo, error) {
	return &fileInfo{f.name, f.file}, nil
}

func (f *file) Seek(offset int64, whence int) (int64, error) {
	return f.file.Seek(offset, whence)
}

func (f *file) Read(p []byte) (int, error) {
	return f.file.Read(p)
}

func (f *file) Close() error {
	return f.file.Close()
}

type fileInfo struct {
	name string
	node files.Node
}

func (fi *fileInfo) Name() string {
	return fi.name
}

func (fi *fileInfo) Size() int64 {
	size, _ := fi.node.Size()
	return size
}

func (fi *fileInfo) Mode() fs.FileMode {
	return 0
}

func (fi *fileInfo) ModTime() time.Time {
	return time.Now()
}

func (fi *fileInfo) IsDir() bool {
	_, ok := fi.node.(files.Directory)
	return ok
}

func (fi *fileInfo) Sys() interface{} {
	return nil
}
