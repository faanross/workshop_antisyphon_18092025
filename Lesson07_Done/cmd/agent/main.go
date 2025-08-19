package main

import (
	"akkeDNSII/internals/config"
	"akkeDNSII/internals/models"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
)

const pathToYAML = "./configs/config.yaml"

func main() {
	// Command line flag for config file path
	configPath := flag.String("config", pathToYAML, "path to configuration file")
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	comm, err := models.NewAgent(cfg)
	if err != nil {
		log.Fatalf("Failed to create communicator: %v", err)
	}

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	comm.Send(ctx)
	
	//Start run loop in goroutine
	//go func() {
	//	log.Printf("Starting %s client run loop", cfg.Protocol)
	//	log.Printf("Delay: %v, Jitter: %d%%", cfg.Timing.Delay, cfg.Timing.Jitter)
	//
	//	if err := runloop.RunLoop(ctx, comm, cfg); err != nil {
	//		log.Printf("Run loop error: %v", err)
	//	}
	//}()

	// Wait for interrupt signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down client...")
	cancel() // This will cause the run loop to exit
}
