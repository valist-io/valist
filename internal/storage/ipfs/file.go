package ipfs

import (
	"io/fs"
	"time"

	files "github.com/ipfs/go-ipfs-files"
)

type File struct {
	name string
	file files.File
}

func (f *File) Stat() (fs.FileInfo, error) {
	return &FileInfo{f.name, f.file}, nil
}

func (f *File) Read(p []byte) (int, error) {
	return f.file.Read(p)
}

func (f *File) Close() error {
	return f.file.Close()
}

type FileInfo struct {
	name string
	node files.Node
}

func (fi *FileInfo) Name() string {
	return fi.name
}

func (fi *FileInfo) Size() int64 {
	size, _ := fi.node.Size()
	return size
}

func (fi *FileInfo) Mode() FileMode {
	return 0
}

func (fi *FileInfo) ModTime() time.Time {
	return time.Now()
}

func (fi *FileInfo) IsDir() bool {
	_, ok := fi.node.(files.Directory)
	return ok
}

func (fi *FileInfo) Sys() interface{} {
	return nil
}
