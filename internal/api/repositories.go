package api

import (
	"encoding/json"
	"net/http"
)

func (s Server) GetRepositories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	_ = json.NewEncoder(w).Encode(Error{
		Code:    new(int(http.StatusNotImplemented)),
		Message: new(string("Not implemented")),
	})
}

func (s Server) AddRepository(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(Error{
		Code:    new(int(http.StatusNotImplemented)),
		Message: new(string("Not implemented")),
	})
}
