package api

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetCharts(w http.ResponseWriter, r *http.Request) {
	// TODO: interface to storage backend for querying
	res := Chart{Id: new(int64(1)), Name: new(string("test-chart"))}
	w.WriteHeader(http.StatusOK)
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
