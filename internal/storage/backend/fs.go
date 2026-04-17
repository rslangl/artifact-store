package backend

import (
	"os"
	"io/fs"
)

type FileSystem struct {
	path fs.FS
}

// Implementation of the `Initializer` interface
func (fs *FileSystem) Initialize(path string) error {
	// TODO: ensure path exists, with read and write permisions
	fs.path = os.DirFS(path)
	return nil
}

// Implementation of the `Writer` interface
func (fs *FileSystem) Write(fileBytes []byte) error {
	return nil
}

// Implementation of the `Reader` interface
func (fs *FileSystem) Read(path string) error {
	return nil
}

// Implementation of the `Terminator` interface
func (fs *FileSystem) Close() error {
	return nil
}
