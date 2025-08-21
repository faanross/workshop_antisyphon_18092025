package https

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

// HTTPSAgent implements the Communicator interface for HTTPS
type HTTPSAgent struct {
	serverAddr string
	// TODO add client field of type *http.Client
}

// NewHTTPSAgent creates a new HTTPS agent
func NewHTTPSAgent(serverAddr string) *HTTPSAgent {
	// Create TLS config that accepts self-signed certificates
	tlsConfig := &tls.Config{
		// TODO ensure our tlsConfig allows self-signed certs
	}

	// Create HTTP client with custom TLS config
	client := &http.Client{
		Transport: &http.Transport{
			// TODO assign tlsConfig as our TLSClientConfig field
		},
	}

	return &HTTPSAgent{
		serverAddr: serverAddr,
		client:     client,
	}
}

// Send implements Communicator.Send for HTTPS
func (c *HTTPSAgent) Send(ctx context.Context) ([]byte, error) {
	// Construct the URL
	// TODO use Sprintf to construct the https url

	// Create GET request
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Send request
	// TODO send the request using client.Do(), returns a response and error
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("server returned status %d: %s", resp.StatusCode, body)
	}

	// Read response body
	// TODO use io.ReadAll to read response body, returns body content and error
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	// Return the raw JSON as message data
	return body, nil
}
