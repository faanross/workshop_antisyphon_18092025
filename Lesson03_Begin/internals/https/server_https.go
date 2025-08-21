package https

import (
	"akkeDNSII/internals/config"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

// HTTPSServer implements the Server interface for HTTPS
type HTTPSServer struct {
	addr   string
	server *http.Server
	// TODO add fields for tlsCert and tlsKey, both strings (path to the files)
}

// HTTPSResponse represents the JSON response for HTTPS
type HTTPSResponse struct {
	Change bool // TODO add tag here to allow for conversion to JSON
}

// NewHTTPSServer creates a new HTTPS server
func NewHTTPSServer(cfg *config.Config) *HTTPSServer {
	return &HTTPSServer{
		addr: cfg.ServerAddr,
		// TODO add fields for tlsCert and tlsKey
	}
}

// Start implements Server.Start for HTTPS
func (s *HTTPSServer) Start() error {
	// Create Chi router
	r := chi.NewRouter()

	// Define our GET endpoint
	// TODO define a GET endpoint at /, call the RootHandler function

	// Create the HTTP server
	s.server = &http.Server{
		// TODO assign fields Addr and Handler
		// Addr is equal to addr field from s
		// Handler is the Chi router object r
	}

	// Start the server
	return s.server.ListenAndServeTLS(s.tlsCert, s.tlsKey)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Endpoint %s has been hit by agent\n", r.URL.Path)

	// Create response with change set to false
	// TODO create variable response which is HTTPSResponse with Change field set to false

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
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
	// TODO create context.WithTimeout with a 5 second limit
	defer cancel()

	// TODO assign return value directly as calling Shutdown on our HTTPSServer's server field
	// Remember to pass context as argument

}
