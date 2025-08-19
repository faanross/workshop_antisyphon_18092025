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
	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parsing config file: %w", err)
	}

	// Optional, but good proactive -> Validate the configuration
	if err := cfg.ValidateConfig(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &cfg, nil
}

// ValidateConfig checks if the configuration is valid
func (c *Config) ValidateConfig() error {
	if c.ClientAddr == "" {
		return fmt.Errorf("agent address cannot be empty")
	}

	if c.ServerAddr == "" {
		return fmt.Errorf("server address cannot be empty")
	}

	if c.Timing.Delay <= 0 {
		return fmt.Errorf("delay must be positive")
	}

	if c.Timing.Jitter < 0 || c.Timing.Jitter > 100 {
		return fmt.Errorf("jitter must be between 0 and 100")
	}

	if c.TlsCert == "" {
		return fmt.Errorf("tls cert cannot be empty")
	}

	if c.TlsKey == "" {
		return fmt.Errorf("tls cert cannot be empty")
	}

	return nil
}
