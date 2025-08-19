package config

import "time"

// Config holds all application configuration
type Config struct {
	ClientAddr string       `yaml:"client"`
	ServerAddr string       `yaml:"server"`
	Timing     TimingConfig `yaml:"timing"`
	Protocol   string       `yaml:"protocol"` // this will be the starting protocol
	TlsKey     string       `yaml:"tls_key"`
	TlsCert    string       `yaml:"tls_cert"`
}

type TimingConfig struct {
	Delay  time.Duration `yaml:"delay"`  // Base delay between cycles
	Jitter int           `yaml:"jitter"` // Jitter percentage (0-100)}
}
