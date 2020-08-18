package main

import (
	"log"

	"github.com/go-magma/fbinternal/cloud/go/fbinternal"
	"github.com/go-magma/fbinternal/cloud/go/protos"
	"github.com/go-magma/fbinternal/cloud/go/services/vpnservice"
	"github.com/go-magma/fbinternal/cloud/go/services/vpnservice/servicers"
	"github.com/go-magma/magma/orc8r/cloud/go/service"
)

const taKeyPath = "/var/opt/magma/certs/vpn_ta.key"

func main() {
	srv, err := service.NewOrchestratorService(fbinternal.ModuleName, vpnservice.ServiceName)
	if err != nil {
		log.Fatalf("Error creating service: %s", err)
	}

	// Add servicers to the service
	servicer := servicers.NewVPNServicer(taKeyPath)
	protos.RegisterVPNServiceServer(srv.GrpcServer, servicer)

	// Run the service
	err = srv.Run()
	if err != nil {
		log.Fatalf("Error running service: %s", err)
	}
}
