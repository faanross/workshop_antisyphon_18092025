package dns

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
	"log"
)

// DNSAgent implements the CommunicatorAgent interface for DNS
type DNSAgent struct {
	serverAddr string
	client     *dns.Client
}

// NewDNSAgent creates a new DNS client
func NewDNSAgent(serverAddr string) *DNSAgent {
	return &DNSAgent{
		serverAddr: serverAddr,
		client:     new(dns.Client),
	}
}

// Send implements Communicator.Send for DNS
func (c *DNSAgent) Send(ctx context.Context) ([]byte, error) {
	// Create DNS query message
	m := new(dns.Msg)

	// For now, we'll query for a fixed domain
	domain := "www.thisdoesnotexist.com."
	m.SetQuestion(domain, dns.TypeA)
	log.Printf("Sending DNS query for: %s", domain)

	// Send query
	r, _, err := c.client.Exchange(m, c.serverAddr)
	if err != nil {
		return nil, fmt.Errorf("DNS exchange failed: %w", err)
	}

	// Check if we got an answer
	if len(r.Answer) == 0 {
		return nil, fmt.Errorf("no answer received")
	}

	// Extract the first A record
	for _, ans := range r.Answer {
		if a, ok := ans.(*dns.A); ok {
			// Return the IP address as string
			ipStr := a.A.String()
			log.Printf("Received DNS response: %s -> %s", domain, ipStr)
			return []byte(ipStr), nil
		}
	}

	return nil, fmt.Errorf("no A record in response")
}
