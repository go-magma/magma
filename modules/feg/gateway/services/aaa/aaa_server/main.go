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

// Package main implements WiFi AAA server
package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"

	managed_configs "github.com/go-magma/magma/gateway/go/mconfig"
	"github.com/go-magma/magma/lib/go/service"
	fegprotos "github.com/go-magma/magma/modules/feg/cloud/go/protos"
	"github.com/go-magma/magma/modules/feg/cloud/go/protos/mconfig"
	"github.com/go-magma/magma/modules/feg/gateway/registry"
	"github.com/go-magma/magma/modules/feg/gateway/services/aaa/protos"
	"github.com/go-magma/magma/modules/feg/gateway/services/aaa/servicers"
	"github.com/go-magma/magma/modules/feg/gateway/services/aaa/store"
	lteprotos "github.com/go-magma/magma/modules/lte/cloud/go/protos"
)

const (
	AAAServiceName = "aaa_server"
	Version        = "0.1"
)

func main() {
	flag.Parse() // for glog

	// Create a shared Session Table
	sessions := store.NewMemorySessionTable()

	// Create the EAP AKA Provider service
	srv, err := service.NewServiceWithOptions(registry.ModuleName, registry.AAA_SERVER)
	if err != nil {
		glog.Fatalf("Error creating AAA service: %s", err)
	}
	aaaConfigs := &mconfig.AAAConfig{}
	err = managed_configs.GetServiceConfigs(AAAServiceName, aaaConfigs)
	if err != nil {
		glog.Warningf("Error getting AAA Server service configs: %s", err)
		aaaConfigs = nil
	}
	acct, _ := servicers.NewAccountingService(sessions, proto.Clone(aaaConfigs).(*mconfig.AAAConfig))
	protos.RegisterAccountingServer(srv.GrpcServer, acct)
	lteprotos.RegisterAbortSessionResponderServer(srv.GrpcServer, acct)
	fegprotos.RegisterSwxGatewayServiceServer(srv.GrpcServer, acct)

	auth, _ := servicers.NewEapAuthenticator(sessions, aaaConfigs, acct)
	protos.RegisterAuthenticatorServer(srv.GrpcServer, auth)

	// Starts built in radius server if built with this option
	startBuiltInRadius(aaaConfigs, auth, acct)

	glog.Infof("Starting AAA Service v%s.", Version)
	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error running AAA service: %s", err)
	}
}
