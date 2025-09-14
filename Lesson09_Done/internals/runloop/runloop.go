package runloop

import (
	"akkeDNSII/internals/config"
	"akkeDNSII/internals/https"
	"akkeDNSII/internals/models"
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

func RunLoop(ctx context.Context, comm models.Agent, cfg *config.Config) error {

	for {
		// Check if context is cancelled
		select {
		case <-ctx.Done():
			log.Println("Run loop cancelled")
			return nil
		default:
		}

		response, err := comm.Send(ctx)
		if err != nil {
			log.Printf("Error sending request: %v", err)
			// Don't exit - just sleep and try again
			time.Sleep(cfg.Timing.Delay)
			continue // Skip to next iteration
		}

		// BASED ON PROTOCOL, HANDLE PARSING DIFFERENTLY

		switch cfg.Protocol {
		case "https":
			// Parse and display response
			var httpsResp https.HTTPSResponse
			if err := json.Unmarshal(response, &httpsResp); err != nil {
				log.Fatalf("Failed to parse response: %v", err)
			}

			log.Printf("Received response: change=%v", httpsResp.Change)
		case "dns":
			ipAddr := string(response)
			log.Printf("Received response: IP=%v", ipAddr)

		}

		// Calculate sleep duration with jitter
		sleepDuration := CalculateSleepDuration(time.Duration(cfg.Timing.Delay), cfg.Timing.Jitter)
		log.Printf("Sleeping for %v", sleepDuration)

		// Sleep with cancellation support
		select {
		case <-time.After(sleepDuration):
			// Continue to next iteration
		case <-ctx.Done():
			log.Println("Run loop cancelled")
			return nil
		}
	}
}

// CalculateSleepDuration calculates the actual sleep time with jitter
func CalculateSleepDuration(baseDelay time.Duration, jitterPercent int) time.Duration {
	if jitterPercent == 0 {
		return baseDelay
	}

	// Calculate jitter range
	jitterRange := float64(baseDelay) * float64(jitterPercent) / 100.0

	// Random value between -jitterRange and +jitterRange
	jitter := (rand.Float64()*2 - 1) * jitterRange

	// Calculate final duration
	finalDuration := float64(baseDelay) + jitter

	// Ensure we don't go negative
	if finalDuration < 0 {
		finalDuration = 0
	}

	return time.Duration(finalDuration)
}
