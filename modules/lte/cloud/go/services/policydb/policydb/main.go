/*
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/go-magma/magma/modules/lte/cloud/go/lte"
	"github.com/go-magma/magma/modules/lte/cloud/go/protos"
	"github.com/go-magma/magma/modules/lte/cloud/go/services/policydb"
	"github.com/go-magma/magma/modules/lte/cloud/go/services/policydb/obsidian/handlers"
	"github.com/go-magma/magma/modules/lte/cloud/go/services/policydb/servicers"
	"github.com/go-magma/magma/orc8r/cloud/go/obsidian"
	"github.com/go-magma/magma/orc8r/cloud/go/service"

	"github.com/golang/glog"
)

func main() {
	// Create the service
	srv, err := service.NewOrchestratorService(lte.ModuleName, policydb.ServiceName)
	if err != nil {
		glog.Fatalf("Error creating service: %s", err)
	}
	assignmentServicer := servicers.NewPolicyAssignmentServer()
	protos.RegisterPolicyAssignmentControllerServer(srv.GrpcServer, assignmentServicer)
	obsidian.AttachHandlers(srv.EchoServer, handlers.GetHandlers())
	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error while running service and echo server: %s", err)
	}
}
