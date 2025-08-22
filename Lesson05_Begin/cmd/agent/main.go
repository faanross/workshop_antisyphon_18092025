package main

import (
	"akkeDNSII/internals/config"
	"akkeDNSII/internals/models"
	"akkeDNSII/internals/runloop"
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

	// Call agent's factory function
	comm, err := models.NewAgent(cfg)
	if err != nil {
		log.Fatalf("Failed to create communicator: %v", err)
	}

	// Create context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Start run loop in goroutine
	go func() {
		log.Printf("Starting %s client run loop", cfg.Protocol)
		log.Printf("Delay: %v, Jitter: %d%%", cfg.Timing.Delay, cfg.Timing.Jitter)

		// TODO call runloop.RunLoop + error check
	}()

	// Wait for interrupt signal
	// TODO create sigChan channel
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	log.Println("Shutting down client...")
	// Todo call cancel()
}
