package models

import (
	"akkeDNSII/internals/config"
	"akkeDNSII/internals/dns"
	"akkeDNSII/internals/https"
	"fmt"
)

// NewAgent creates a new communicator based on the protocol
func NewAgent(cfg *config.Config) (Agent, error) {
	switch cfg.Protocol {
	case "https":
		return https.NewHTTPSAgent(cfg.ServerAddr), nil
	case "dns":
		return dns.NewDNSAgent(cfg.ServerAddr), nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}

// NewServer creates a new server based on the protocol
func NewServer(cfg *config.Config) (Server, error) {
	switch cfg.Protocol {
	case "https":
		return https.NewHTTPSServer(cfg), nil
	case "dns":
		return dns.NewDNSServer(cfg), nil
	default:
		return nil, fmt.Errorf("unsupported protocol: %v", cfg.Protocol)
	}
}
