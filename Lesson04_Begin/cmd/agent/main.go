package main

import (
	"akkeDNSII/internals/config"
	"akkeDNSII/internals/https"
	"akkeDNSII/internals/models"
	"context"
	"encoding/json"
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

	// TODO pass the config to our models.NewAgent() factory function
	if err != nil {
		log.Fatalf("Failed to create communicator: %v", err)
	}

	// TEMPORARY CODE JUST TO TEST!
	// Send a test message

	log.Printf("Sending request to %s server...", cfg.Protocol)

	// TODO Call the Send() method on the comm instance
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Parse and display response
	var httpsResp https.HTTPSResponse
	if err := json.Unmarshal(response, &httpsResp); err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}

	log.Printf("Received response: change=%v", httpsResp.Change)
}
