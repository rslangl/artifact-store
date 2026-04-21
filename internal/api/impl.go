package api

import (
	"net/http"
	"encoding/json"
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
