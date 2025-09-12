package config

// Config holds all application configuration
type Config struct {
	ClientAddr string
	ServerAddr string
	Timing     TimingConfig
	// TODO Add fields for Protocol, TlsKey, and TlsCert (all string)
}

type TimingConfig struct {
	// TODO Add fields for Delay (time.Duration) and Jitter (int)
}
