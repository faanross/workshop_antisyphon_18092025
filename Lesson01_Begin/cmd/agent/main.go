package main

import (
	"akkeDNSII/internals/config"
	"akkeDNSII/internals/models"
	"fmt"
)

func main() {
	agentCfg := config.Config{
		// TODO Set protocol to HTTPS
	}

	_, err := models.NewAgent(&agentCfg)
	// TODO if there is an error, use fmt.Println to display it

}
