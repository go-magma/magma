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

// Package main implements Magma EAP SIM Service
package main

import (
	"flag"

	"github.com/golang/glog"

	managed_configs "github.com/go-magma/magma/gateway/go/mconfig"
	"github.com/go-magma/magma/lib/go/service"
	"github.com/go-magma/magma/modules/feg/cloud/go/protos/mconfig"
	"github.com/go-magma/magma/modules/feg/gateway/registry"
	"github.com/go-magma/magma/modules/feg/gateway/services/eap/protos"
	"github.com/go-magma/magma/modules/feg/gateway/services/eap/providers/sim"
	"github.com/go-magma/magma/modules/feg/gateway/services/eap/providers/sim/servicers"
	_ "github.com/go-magma/magma/modules/feg/gateway/services/eap/providers/sim/servicers/handlers"
)

func init() {
	flag.Parse()
}

func main() {
	// Create the EAP SIM Provider service
	srv, err := service.NewServiceWithOptions(registry.ModuleName, registry.EAP_SIM)
	if err != nil {
		glog.Fatalf("Error creating EAP SIM service: %s", err)
	}

	simConfigs := &mconfig.EapProviderConfig{}
	err = managed_configs.GetServiceConfigs(sim.EapSimServiceName, simConfigs)
	if err != nil {
		glog.Errorf("Error getting EAP SIM service configs: %s", err)
		simConfigs = nil
	}
	servicer, err := servicers.NewEapSimService(simConfigs)
	if err != nil {
		glog.Fatalf("failed to create EAP SIM Service: %v", err)
		return
	}
	protos.RegisterEapServiceServer(srv.GrpcServer, servicer)

	// Run the service
	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error running EAP SIM service: %s", err)
	}
}
