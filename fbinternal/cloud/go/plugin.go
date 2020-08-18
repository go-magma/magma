package main

import (
	fbinternalp "github.com/go-magma/fbinternal/cloud/go/plugin"
	"github.com/go-magma/magma/orc8r/cloud/go/plugin"
)

func main() {}

func GetOrchestratorPlugin() plugin.OrchestratorPlugin {
	return &fbinternalp.FbinternalOrchestratorPlugin{}
}
