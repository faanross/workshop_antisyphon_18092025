package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// LoadConfig reads and parses the configuration file
func LoadConfig(path string) (*Config, error) {
	// We'll provide path to *.yaml to function when we call it
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening config file: %w", err)
	}
	defer file.Close()

	// instantiate struct to unmarshall yaml into
	// TODO Create an empty Config struct
	// TODO unmarshall YAML info struct

	// Optional, but good proactive -> Validate the configuration
	// TODO call the ValidateConfig() method

	return &cfg, nil
}

// ValidateConfig checks if the configuration is valid
func (c *Config) ValidateConfig() error {
	if c.ClientAddr == "" {
		return fmt.Errorf("agent address cannot be empty")
	}

	// TODO make sure ServerAddr is not blank

	if c.Timing.Delay <= 0 {
		return fmt.Errorf("delay must be positive")
	}

	// TODO make sure Timing.Jitter is between 0 and 100 inclusive

	if c.TlsCert == "" {
		return fmt.Errorf("tls cert cannot be empty")
	}

	// TODO make sure TlsKey is not blank

	return nil
}
