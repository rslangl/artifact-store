package service

import (
	//"log"
	//"fmt"
	"net/http"
	"context"

	"artifact-store/internal/api"
	"artifact-store/internal/config"
	"artifact-store/internal/storage"
)

type WebService struct {
	httpServer *http.Server
}

func Create(config config.ServiceConfig, storage storage.Storage) WebService {
	server := api.NewServer(storage)
	router := http.NewServeMux()

	handler := api.HandlerFromMux(server, router)

	return WebService{
		httpServer: &http.Server{
			Handler: handler,
			Addr: config.Address,
		},
	}
}

func (svc* WebService) Run() error {
	return svc.httpServer.ListenAndServe()
}

func (svc* WebService) Shutdown(ctx context.Context) error {
	return svc.httpServer.Shutdown(ctx)
}
