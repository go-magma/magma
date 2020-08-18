package main

import (
	"flag"

	"github.com/go-magma/fbinternal/cloud/go/services/download/servicers"
	"github.com/go-magma/magma/orc8r/cloud/go/plugin"
	"github.com/go-magma/magma/lib/go/registry"
)

func main() {
	flag.Parse()
	plugin.LoadAllPluginsFatalOnError(&plugin.DefaultOrchestratorPluginLoader{})
	registry.MustPopulateServices()

	servicers.Run()
}
