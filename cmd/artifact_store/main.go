package main

import (
	"artifact-store/internal/api"
	//"artifact-store/internal/storage"
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

	opts := CliOpts{}

	flag.StringVar(&opts.configFile, "config", "", "Path to config file")
	flag.BoolVar(&opts.debug, "debug", false, "Enable debug logging")

	flag.Parse()

	cfg := &config.Config{}

	if err := cfg.Create(opts.configFile); err != nil {
		log.Fatalf("Could not create config: %v", err)
	}

	log.Printf("%v", cfg.ToString())

	// var fss FileSystemStorage = storage.Create()
	// fss.Initialize()
	//
	// var nas NasStorage = storage.Create()
	// nas.Initialize()
	//
	server := api.NewServer()
	router := http.NewServeMux()
	handler := api.HandlerFromMux(server, router)
	service := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:8080",
	}
	service.ListenAndServe()
}
