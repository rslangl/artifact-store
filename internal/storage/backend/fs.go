package backend

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type FileSystem struct {
	Path fs.FS
}

func NewFSBackend(path string) (*FileSystem, error) {
	if err := os.MkdirAll(path, 0744); err != nil {
		return nil, err
	}
	return &FileSystem{
		Path: os.DirFS(path),
	}, nil
}

// Implementation of the `Writer` interface
func (f *FileSystem) Write(bytes []byte) error { // TODO: define type `artifact` or similar instead
	// TODO: create path if not exists (requires more parameters)
	return nil
}

// Implementation of the `Reader` interface
func (f *FileSystem) Read(resource string, version string) ([]byte, error) {
	dir, err := f.Path.Open(filepath.Join(resource, version))
	if err != nil {
		return nil, err
	}
	if rd, ok := dir.(fs.ReadDirFile); ok {
		entries, err := rd.ReadDir(-1)
		if err != nil {
			dir.Close()
			return nil, err
		}
		if len(entries) == 0 {
			dir.Close()
			return nil, fs.ErrNotExist
		}
		file := path.Join(resource, version, entries[0].Name())
		dir.Close()
		fileBytes, err := fs.ReadFile(f.Path, path.Clean(file))
		if err != nil {
			return nil, err
		}
		return fileBytes, nil
	}
	return nil, nil
}
