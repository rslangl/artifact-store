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

func NewFSBackend(path string) (*FileSystem, error) {
	if err := os.MkdirAll(path, 0744); err != nil {
		return nil, fmt.Errorf("Could not create root path for file system backend path '%v': %v", path, err)
	}
	return &FileSystem{
		Path: os.DirFS(path),
	}, nil
}

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
