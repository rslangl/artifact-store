package api

import (
	// "encoding/json"
	// "fmt"
	// "io"
	// "slices"
	// "net/http"
	// openapi_types "github.com/oapi-codegen/runtime/types"
	//
	"artifacts/internal/storage"
	//"artifacts/internal/storage/storage_error"
)

type Server struct{
	storageHandler storage.Storage
}

func NewServer(storageHandler storage.Storage) Server {
	return Server{
		storageHandler: storageHandler,
	}
}

