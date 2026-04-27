package api

import (
	"encoding/json"
	"fmt"
	"io"
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

func (s Server) AddChart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")  // TODO: set globally, or default?

	// Set max header size (50MB)
	r.Body = http.MaxBytesReader(w, r.Body, 50<<20) // TODO: set globally, or default?

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(NewError(http.StatusBadRequest, "Invalid form data"))
		return
	}

	file, fh, err := r.FormFile("chart")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(NewError(http.StatusBadRequest, "Missing chart"))
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(NewError(http.StatusBadRequest, "Malformed file upload"))
		return
	}

	if err := s.storageHandler.Write(fh.Filename, data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(NewError(http.StatusInternalServerError, "Error occurred durinf file IO"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode("")
}
