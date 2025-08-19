package https

import (
	"akkeDNSII/internals/config"
	"akkeDNSII/internals/control"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

// HTTPSServer implements the Server interface for HTTPS
type HTTPSServer struct {
	addr    string
	server  *http.Server
	tlsCert string
	tlsKey  string
}

// HTTPSResponse represents the JSON response for HTTPS
type HTTPSResponse struct {
	Change bool `json:"change"`
}

// NewHTTPSServer creates a new HTTPS server
func NewHTTPSServer(cfg *config.Config) *HTTPSServer {
	return &HTTPSServer{
		addr:    cfg.ServerAddr,
		tlsCert: cfg.TlsCert,
		tlsKey:  cfg.TlsKey,
	}
}

// Start implements Server.Start for HTTPS
func (s *HTTPSServer) Start() error {
	// Create Chi router
	r := chi.NewRouter()

	// Define our GET endpoint
	r.Get("/", RootHandler)

	// Create the HTTP server
	s.server = &http.Server{
		Addr:    s.addr,
		Handler: r,
	}

	// Start the server
	return s.server.ListenAndServeTLS(s.tlsCert, s.tlsKey)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Endpoint %s has been hit by agent\n", r.URL.Path)

	// Check if we should transition
	shouldChange := control.Manager.CheckAndReset()
	response := HTTPSResponse{
		Change: shouldChange,
	}
	if shouldChange {
		log.Printf("HTTPS: Sending transition signal (change=true)")
	} else {
		log.Printf("HTTPS: Normal response (change=false)")
	}
	
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// Stop implements Server.Stop for HTTPS
func (s *HTTPSServer) Stop() error {
	// If there's no server, nothing to stop
	if s.server == nil {
		return nil
	}

	// Give the server 5 seconds to shut down gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.server.Shutdown(ctx)
}
