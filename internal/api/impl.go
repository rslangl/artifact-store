package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"artifact-store/internal/storage"
)

type Server struct{
	storage *storage.Storage
}

func NewServer(stg *storage.Storage) Server {
	return Server{
		storage: stg,
	}
}

func (s Server) GetCharts(w http.ResponseWriter, r *http.Request) {
	data, err := s.storage.FileSystem.Read("") // TODO: constant for root path for Helm charts
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (s Server) GetChart(w http.ResponseWriter, r *http.Request, name string, version string) {
	path := path.Join("", name)  // TODO: handle version for chart
	data, err := s.storage.FileSystem.Read(path)
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

func (Server) GetChart(w http.ResponseWriter, r *http.Request, name string, version string) {
	res := Chart{Id: new(int64(1337)), Name: new(string("test-chart-2"))}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
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
