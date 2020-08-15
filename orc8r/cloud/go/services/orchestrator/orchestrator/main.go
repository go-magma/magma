/*
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/go-magma/magma/orc8r/cloud/go/obsidian"
	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	"github.com/go-magma/magma/orc8r/cloud/go/service"
	exporter_protos "github.com/go-magma/magma/orc8r/cloud/go/services/metricsd/protos"
	"github.com/go-magma/magma/orc8r/cloud/go/services/orchestrator"
	"github.com/go-magma/magma/orc8r/cloud/go/services/orchestrator/obsidian/handlers"
	"github.com/go-magma/magma/orc8r/cloud/go/services/orchestrator/servicers"
	indexer_protos "github.com/go-magma/magma/orc8r/cloud/go/services/state/protos"
	streamer_protos "github.com/go-magma/magma/orc8r/cloud/go/services/streamer/protos"

	"github.com/golang/glog"
)

func main() {
	srv, err := service.NewOrchestratorService(orc8r.ModuleName, orchestrator.ServiceName)
	if err != nil {
		glog.Fatalf("Error creating orchestrator service %s", err)
	}

	obsidian.AttachHandlers(srv.EchoServer, handlers.GetObsidianHandlers())

	exporterServicer := servicers.NewPushExporterServicer(srv.Config.MustGetStrings(orchestrator.PrometheusPushAddresses))
	exporter_protos.RegisterMetricsExporterServer(srv.GrpcServer, exporterServicer)
	indexer_protos.RegisterIndexerServer(srv.GrpcServer, servicers.NewDirectoryIndexer())
	streamer_protos.RegisterStreamProviderServer(srv.GrpcServer, servicers.NewOrchestratorStreamProviderServicer())

	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error while running service and echo server: %s", err)
	}
}
