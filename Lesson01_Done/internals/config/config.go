package config

import "time"

// Config holds all application configuration
type Config struct {
	ClientAddr string
	ServerAddr string
	Timing     TimingConfig
	Protocol   string // this will be the starting protocol
	TlsKey     string
	TlsCert    string
}

type TimingConfig struct {
	Delay  time.Duration // Base delay between cycles
	Jitter int           // Jitter percentage (0-100)}
}
