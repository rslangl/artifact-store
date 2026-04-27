package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"path"

	"artifact-store/internal/storage"
	"artifact-store/internal/storage/storage_error"
)

type ArtifactErrorMessage interface {
	ResourceNotFound(resource string, version string)
	InternalServerError(resource string, version string)
}

type ArtifactError struct {
	ErrorMessage string
}

func (e* ArtifactError) ResourceNotFound(resource string, version string) {
	e.ErrorMessage = fmt.Sprintf("The requested resource '%v:%v' was not found", resource, version)
}

func (e* ArtifactError) InternalServerError(resource string, version string) {
	e.ErrorMessage = fmt.Sprintf("Internal server error occurred while fetching the requested resource '%v:%v'", resource, version)
}

type Server struct{
	storageHandler storage.Storage
}

func NewServer(storageHandler storage.Storage) Server {
	return Server{
		storageHandler: storageHandler,
	}
}

func (s Server) GetCharts(w http.ResponseWriter, r *http.Request) {
	data, err := s.storageHandler.Read("", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (s Server) GetChart(w http.ResponseWriter, r *http.Request, resource string, version string) {
	data, err := s.storageHandler.Read(resource, version)
	if err != nil {
		e := &ArtifactError{}
		if err == storage_error.NotFound {
			e.ResourceNotFound(resource, version)
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(e)
			return
		}
		e.InternalServerError(resource, version)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func (Server) GetChartVersions(w http.ResponseWriter, r *http.Request, name string) {
	// TODO: interface to storage backend for querying
	res := fmt.Sprintf("%v", r)
	_ = json.NewEncoder(w).Encode(res)
}

func (Server) AddChart(w http.ResponseWriter, r *http.Request) {
	// TODO: interface to storage backend for creating
	res := fmt.Sprintf("%v", r)
	_ = json.NewEncoder(w).Encode(res)
}
