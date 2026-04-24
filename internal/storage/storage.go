package storage

import (
	"fmt"

	"artifact-store/internal/config"
	"artifact-store/internal/storage/backend"
)

type Storage interface {
	Read(location string) ([]byte, error)
	Write(bytes []byte) error
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

// func (stg* Storage) ToString() string {
// 	output := fmt.Sprintf("storage : ")
// 	return output
// }
