package api

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	//"path"

	"artifact-store/internal/storage"
)

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

func (s Server) GetChart(w http.ResponseWriter, r *http.Request, name string, version string) {
	data, err := s.storageHandler.Read(name, version)
	if err != nil {
		if err == fs.ErrNotExist { // TODO: map to generic error types in case other backends are used
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(err) // TODO: define error type
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err) // TODO: define error type
		return
	}
	w.WriteHeader(http.StatusOK)
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
