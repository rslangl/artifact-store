package storage

import (
	"fmt"

	"artifacts/internal/config"
	"artifacts/internal/storage/backend"
)

type Storage interface {
	Read(repository string, resource string, version string) ([]byte, error)
	Write(repository string, name string, bytes []byte) error
}

func New(config config.StorageConfig) (Storage, error) {
	switch config.Backend {
		case "fs":
			fsBackend, err := backend.NewFSBackend(config.Fs.Path)
			if err != nil {
				return nil, fmt.Errorf("Error occurred initializing file system backend: %v", err)
			}
			return fsBackend, nil
		default:
			return nil, fmt.Errorf("Unknown or unsupported storage backend: %v", config.Backend)
	}
}
