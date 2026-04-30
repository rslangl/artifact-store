package backend

import (
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"

	"artifacts/internal/storage/storage_error"
)

type FileSystem struct {
	root string
	Path fs.FS
}

func NewFSBackend(path string) (*FileSystem, error) {
	if err := os.MkdirAll(path, 0744); err != nil {
		return nil, err
	}
	return &FileSystem{
		root: path,
		Path: os.DirFS(path),
	}, nil
}

// Implementation of the `Writer` interface
func (f *FileSystem) Write(repository string, name string, bytes []byte) error { // TODO: define type `artifact` or similar instead
	path := filepath.Join(f.root, repository, name)

	// TODO: to be implemented when we decide to add repositories
	// if err := os.MkdirAll(filepath.Dir(<f.root + repo>), 0o744); err != nil {
	// 	return err
	// }

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Implementation of the `Reader` interface
func (f *FileSystem) Read(repository string, resource string, version string) ([]byte, error) {
	dir, err := f.Path.Open(filepath.Join(repository, resource, version))
	if err != nil {
		return nil, storage_error.IOError
	}
	if rd, ok := dir.(fs.ReadDirFile); ok {
		entries, err := rd.ReadDir(-1)
		if err != nil {
			dir.Close()
			return nil, storage_error.IOError
		}
		if len(entries) == 0 {
			dir.Close()
			return nil, storage_error.NotFound
		}
		file := path.Join(resource, version, entries[0].Name())
		dir.Close()
		fileBytes, err := fs.ReadFile(f.Path, path.Clean(file))
		if err != nil {
			return nil, storage_error.IOError
		}
		return fileBytes, nil
	}
	return nil, nil
}
