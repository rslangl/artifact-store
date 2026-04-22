package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"artifact-store/internal/storage"
	"artifact-store/internal/config"
	"artifact-store/internal/service"
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

	// Setup context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

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

	// Create and run webservice
	svc := service.Create(cfg.Service)
	go func() {
		if err := svc.Run(); err != nil {
			log.Fatalf("Could not launch web service")
			stop()
		}
	}()

	<-ctx.Done()

	log.Printf("Terminated")
}
