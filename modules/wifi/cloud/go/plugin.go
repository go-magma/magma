package main

import (
	"github.com/go-magma/magma/orc8r/cloud/go/plugin"
	wifip "github.com/go-magma/wifi/cloud/go/plugin"
)

func main() {}

func GetOrchestratorPlugin() plugin.OrchestratorPlugin {
	return &wifip.WifiOrchestratorPlugin{}
}
