package main

import (
	"akkeDNSII/internals/config"
	"akkeDNSII/internals/models"
	"fmt"
)

func main() {
	agentCfg := config.Config{
		// Set protocol to HTTPS
	}

	_, err := models.NewAgent(&agentCfg)
	// if there is an error, use fmt.Println to display it

}
