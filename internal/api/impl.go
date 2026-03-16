package api

import (
	"encoding/json"
	"net/http"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetCharts(w http.ResponseWriter, r *http.Request) {
	res := Chart{Id: new(int64(1)), Name: new(string("test-chart"))}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
