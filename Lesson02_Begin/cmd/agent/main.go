package main

import (
	"akkeDNSII/internals/config"
	"flag"
	"log"
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

	log.Printf("Loaded configuration:\n")
	log.Printf("-> Client: %s\n", cfg.ClientAddr)
	log.Printf("-> Server: %s\n", cfg.ServerAddr)
	log.Printf("-> Delay: %v\n", cfg.Timing.Delay)
	log.Printf("-> Jitter: %d%%\n", cfg.Timing.Jitter)
	log.Printf("-> Starting Protocol: %s\n", cfg.Protocol)
}
