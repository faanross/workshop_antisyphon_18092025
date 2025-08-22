package dns

import (
	"akkeDNSII/internals/config"
	"github.com/miekg/dns"
	"log"
	"net"
)

// DNSServer implements the Server interface for DNS
type DNSServer struct {
	// TODO add field for addr of type string
	// TODO add server for addr of type *dns.Server

}

// NewDNSServer creates a new DNS server
func NewDNSServer(cfg *config.Config) *DNSServer {
	return &DNSServer{
		// TODO set addr equal to value from config
	}
}

// Start implements Server.Start for DNS
func (s *DNSServer) Start() error {
	// Create and configure the DNS server
	s.server = &dns.Server{
		Addr: s.addr,
		// TODO set Net field equal to "udp"
		// TODO set Handler field equal to s.handleDNSRequest method passed as argument to dns.HandlerFunc()
	}

	// Start server
	return s.server.ListenAndServe()
}

// handleDNSRequest is our DNS Server's handler
func (s *DNSServer) handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	// Create response message
	m := new(dns.Msg)
	// TODO set the message as a reply
	// TODO set the answer as being authoritative

	// Process each question
	for _, question := range r.Question {
		// We only handle A records for now

		// TODO add conditional logic so that if it's not A record skip to next iteration with continue

		// Log the query
		log.Printf("DNS query for: %s", question.Name)

		// For now, always return 42.42.42.42
		rr := &dns.A{
			Hdr: dns.RR_Header{
				Name:   question.Name,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    300,
			},
			// TODO set answer equal to 42.42.42.42
		}
		m.Answer = append(m.Answer, rr)
	}

	// Send response
	// TODO send response (m) using WriteMsg()
}

// Stop implements Server.Stop for DNS
func (s *DNSServer) Stop() error {
	if s.server == nil {
		return nil
	}
	log.Println("Stopping DNS server...")
	// TODO return a call to Shutdown() on s.server
}
