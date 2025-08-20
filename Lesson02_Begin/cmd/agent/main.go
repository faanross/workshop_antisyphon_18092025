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
	// TODO Call Parse() on flag

	// Load configuration
	// TODO Call our new LoadConfig function
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Loaded configuration:\n")
	log.Printf("-> Client: %s\n", cfg.ClientAddr)
	// TODO Add Print statement for Server
	// TODO Add Print statement for Delay
	log.Printf("-> Jitter: %d%%\n", cfg.Timing.Jitter)
	// TODO Add Print statement for Starting Protocol
}
