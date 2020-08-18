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

package test_init

import (
	"testing"

	"github.com/go-magma/magma/modules/feg/cloud/go/feg"
	"github.com/go-magma/magma/modules/feg/cloud/go/protos"
	"github.com/go-magma/magma/modules/feg/cloud/go/services/feg_relay"
	"github.com/go-magma/magma/modules/feg/cloud/go/services/feg_relay/servicers"
	"github.com/go-magma/magma/orc8r/cloud/go/test_utils"

	"golang.org/x/net/context"
)

// A little Go "polymorphism" magic for testing
type testFegProxyServer struct {
	servicers.FegToGwRelayServer
}

func (srv *testFegProxyServer) CancelLocation(
	ctx context.Context,
	req *protos.CancelLocationRequest,
) (*protos.CancelLocationAnswer, error) {
	return srv.CancelLocationUnverified(ctx, req)
}

func StartTestService(t *testing.T) {
	srv, lis := test_utils.NewTestService(t, feg.ModuleName, feg_relay.ServiceName)
	protos.RegisterS6AGatewayServiceServer(srv.GrpcServer, &testFegProxyServer{})
	go srv.RunTest(lis)
}
