package storage

import (
	"fmt"

	"artifact-store/internal/config"
	"artifact-store/internal/storage/backend"
)

// Initializes the underlying storage system
// type Initializer interface {
// 	Initialize() error
// }
//
// // Writes the file bytes contents to the underlying system
// type Writer interface {
// 	Write(fileBytes []byte) error
// 	// TODO: need to define the model
// 	// Write(file_package Package) (int, error)
// }
//
// // Read the file byte contents from the underlying system
// type Reader interface {
// 	Read(location string) ([]byte, error)
// 	// TODO: need to define the model
// 	// Read(location string) (Package, error)
// }
//
// type Terminator interface {
// 	Close() error
// }

type Storage interface {
	Read(location string) ([]byte, error)
	Write(bytes []byte) error
}

// type Storage struct {
// 	FileSystem backend.FileSystem
// }

func New(config config.StorageConfig) (Storage, error) {
	switch config.Backend {
		case "fs":
			return backend.NewFSBackend(config.Fs.Path), nil
		default:
			return nil, fmt.Errorf("Unknown or unsupported storage backend: %v", config.Backend)
	}
}

// func (stg* Storage) Create(config config.StorageConfig) error {
// 	for _, s := range config.Enabled {
// 		switch(s) {
// 			case "fs":
// 				stg.FileSystem = backend.FileSystem{}
// 				fsPath := config.Fs.Path
// 				err := stg.FileSystem.Initialize(fsPath)
// 				if err != nil {
// 					return fmt.Errorf("could not initialize file system storage: %v", err)
// 				}
// 			default:
// 				return fmt.Errorf("unknown or unsupported storage backend '%v'", s)
// 		}
// 	}
// 	return nil
// }
//
// func (stg* Storage) ToString() string {
// 	output := fmt.Sprintf("storage : ")
// 	return output
// }
