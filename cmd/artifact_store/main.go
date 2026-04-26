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

	opts := CliOpts{}
	flag.StringVar(&opts.configFile, "config", "", "Path to config file")
	flag.BoolVar(&opts.debug, "debug", false, "Enable debug logging")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.New(opts.configFile)
	if err != nil {
		log.Fatalf("Could not create config: %v", err)
	}
	log.Printf("%v", cfg.ToString())

	handler, err := storage.New(cfg.Storage)
	if err != nil {
		log.Fatalf("Could not initialize storage backend: %v", err)
	}

	svc := service.Create(cfg.Service, handler)
	go func() {
		if err := svc.Run(); err != nil {
			log.Fatalf("Could not launch web service")
			stop()
		}
	}()

	<-ctx.Done()

	log.Printf("Terminated")
}
