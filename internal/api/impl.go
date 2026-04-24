package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

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
	data, err := s.storageHandler.Read(".") // TODO: constant for root path for Helm charts
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (s Server) GetChart(w http.ResponseWriter, r *http.Request, name string, version string) {
	path := path.Join("", name)  // TODO: handle version for chart
	data, err := s.storageHandler.Read(path)
	if err != nil {
		// TODO: handle 404 Not Found
		w.WriteHeader(http.StatusInternalServerError)
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
