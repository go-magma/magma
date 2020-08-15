/*
 Copyright (c) Facebook, Inc. and its affiliates.
 All rights reserved.

 This source code is licensed under the BSD-style license found in the
 LICENSE file in the root directory of this source tree.
*/

package test_init

import (
	"testing"

	"github.com/go-magma/magma/lib/go/definitions"
	"github.com/go-magma/magma/lib/go/protos"
	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	"github.com/go-magma/magma/orc8r/cloud/go/services/orchestrator"
	"github.com/go-magma/magma/orc8r/cloud/go/services/orchestrator/servicers"
	indexer_protos "github.com/go-magma/magma/orc8r/cloud/go/services/state/protos"
	streamer_protos "github.com/go-magma/magma/orc8r/cloud/go/services/streamer/protos"
	streamer_servicers "github.com/go-magma/magma/orc8r/cloud/go/services/streamer/servicers"
	"github.com/go-magma/magma/orc8r/cloud/go/test_utils"
)

type testStreamerServer struct {
	protos.StreamerServer
}

func (srv *testStreamerServer) GetUpdates(req *protos.StreamRequest, stream protos.Streamer_GetUpdatesServer) error {
	return streamer_servicers.GetUpdatesUnverified(req, stream)
}

func StartTestService(t *testing.T) {
	labels := map[string]string{
		orc8r.StreamProviderLabel: "true",
	}
	annotations := map[string]string{
		orc8r.StreamProviderStreamsAnnotation: definitions.MconfigStreamName,
	}
	srv, lis := test_utils.NewTestOrchestratorService(t, orc8r.ModuleName, orchestrator.ServiceName, labels, annotations)

	protos.RegisterStreamerServer(srv.GrpcServer, &testStreamerServer{})

	indexer_protos.RegisterIndexerServer(srv.GrpcServer, servicers.NewDirectoryIndexer())
	streamer_protos.RegisterStreamProviderServer(srv.GrpcServer, servicers.NewOrchestratorStreamProviderServicer())

	go srv.RunTest(lis)
}
