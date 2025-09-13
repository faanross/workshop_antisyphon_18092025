package control

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// TransitionManager handles the global transition state
type TransitionManager struct {
	mu sync.RWMutex
	// TODO create our flag - a bool called shouldTransition
	shouldTransition bool
}

// Global instance
var Manager = &TransitionManager{
	shouldTransition: false,
}

// CheckAndReset atomically checks if transition is needed and resets the flag
// This ensures the transition signal is consumed only once
func (tm *TransitionManager) CheckAndReset() bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if tm.shouldTransition {
		// TODO immediately reset shouldTransition to false
		log.Printf("Transition signal consumed and reset")
		return true
	}

	return false
}

// TriggerTransition sets the transition flag
func (tm *TransitionManager) TriggerTransition() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// TODO set flag tm.shouldTransition to true
	log.Printf("Transition triggered")
}

func handleSwitch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO call Manager.TriggerTransition()

	response := "Protocol transition triggered"

	json.NewEncoder(w).Encode(response)
}

// StartControlAPI starts the control API server on port 8080
func StartControlAPI() {

	// TODO call http.HandleFunc, set endpoint as /switch (arg 1) and call handler handleSwitch (arg 2)

	log.Println("Starting Control API on :8080")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Printf("Control API error: %v", err)
		}
	}()
}
