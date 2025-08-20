package models

// Agent defines the contract for agents
type Agent interface {
	// Send sends a message and waits for a response
	// TODO CREATE SEND() SIGNATURE HERE
}

// Server defines the contract for servers
type Server interface {
	// Start begins listening for requests
	// TODO CREATE START() SIGNATURE HERE

	// Stop gracefully shuts down the server
	// TODO CREATE STOP() SIGNATURE HERE
}
