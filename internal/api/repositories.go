package api

import (
	"encoding/json"
	// "fmt"
	// "io"
	// "slices"
	"net/http"
	//openapi_types "github.com/oapi-codegen/runtime/types"

	// "artifacts/internal/storage"
	// "artifacts/internal/storage/storage_error"
)

func (s Server) GetRepositories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Error{
		Code: new(int(http.StatusOK)),
		Message: new(string("")),
	})
	return
}

func (s Server) AddRepository(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Error{
		Code: new(int(http.StatusCreated)),
		Message: new(string("Repository created")),
	})
	return
}
