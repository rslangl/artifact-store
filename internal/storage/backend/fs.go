package backend

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

type FileSystem struct {
	Path fs.FS
}

func NewFSBackend(path string) *FileSystem {
	return &FileSystem{
		Path: os.DirFS(path),
	}
}

// Implementation of the `Initializer` interface
// func (f *FileSystem) Initialize(path string) error {
// 	// TODO: ensure path exists, with read and write permisions
// 	f.Path = os.DirFS(path)
// 	return nil
// }

// Implementation of the `Writer` interface
func (f *FileSystem) Write(bytes []byte) error {
	// TODO: create path if not exists (requires more parameters)
	return nil
}

// Implementation of the `Reader` interface
func (f *FileSystem) Read(resource string) ([]byte, error) {
	bytes, err := fs.ReadFile(f.Path, path.Clean(resource))
	if err != nil {
		return nil, fmt.Errorf("File read error for '%v': %v", resource, err)
	}
	return bytes, nil
}

// Implementation of the `Terminator` interface
func (fs *FileSystem) Close() error {
	return nil
}
