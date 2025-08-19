package main

import (
	"akkeDNSII/internals/config"
	"akkeDNSII/internals/models"
	"fmt"
)

func main() {
	agentCfg := config.Config{
		Protocol: "https",
	}

	_, err := models.NewAgent(&agentCfg)
	if err != nil {
		fmt.Println(err)
	}

}
