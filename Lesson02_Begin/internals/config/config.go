package config

import "time"

// Config holds all application configuration
type Config struct {
	ClientAddr string `yaml:"client"`
	// TODO Add YAML tags to all other fields
	// Make sure they match exactly with labels in the actual YAML file
	ServerAddr string
	Timing     TimingConfig
	Protocol   string
	TlsKey     string
	TlsCert    string
}

type TimingConfig struct {
	Delay  time.Duration
	Jitter int
}
