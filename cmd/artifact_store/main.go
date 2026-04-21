package main

import (
	"artifact-store/internal/api"
	"artifact-store/internal/storage"
	"artifact-store/internal/config"
	"net/http"
	"flag"
	"log"
)

type CliOpts struct {
	configFile string
	debug bool
}

func main() {

	// Define and capture CLI args
	opts := CliOpts{}
	flag.StringVar(&opts.configFile, "config", "", "Path to config file")
	flag.BoolVar(&opts.debug, "debug", false, "Enable debug logging")
	flag.Parse()

	// Create and print runtime config
	cfg := &config.Config{}
	if err := cfg.Create(opts.configFile); err != nil {
		log.Fatalf("Could not create config: %v", err)
	}
	log.Printf("%v", cfg.ToString())

	// Create and print storage config
	stg := &storage.Storage{}
	if err := stg.Create(cfg.Storage); err != nil {
		log.Fatalf("Could not setup storage: %v", err)
	}
	log.Printf("%v", stg.ToString())

	// TODO: place webserver in subpackage
	server := api.NewServer()
	router := http.NewServeMux()
	handler := api.HandlerFromMux(server, router)
	service := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:8080",
	}
	service.ListenAndServe()
}
