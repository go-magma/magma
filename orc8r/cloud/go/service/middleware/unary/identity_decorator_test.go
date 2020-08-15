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

package unary_test

import (
	"net"
	"testing"
	"time"

	"github.com/go-magma/magma/lib/go/protos"
	"github.com/go-magma/magma/lib/go/registry"
	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	"github.com/go-magma/magma/orc8r/cloud/go/serde"
	"github.com/go-magma/magma/orc8r/cloud/go/service"
	"github.com/go-magma/magma/orc8r/cloud/go/service/middleware/unary/test_utils"
	"github.com/go-magma/magma/orc8r/cloud/go/services/configurator"
	configuratorTestInit "github.com/go-magma/magma/orc8r/cloud/go/services/configurator/test_init"
	configuratorTestUtils "github.com/go-magma/magma/orc8r/cloud/go/services/configurator/test_utils"
	"github.com/go-magma/magma/orc8r/cloud/go/services/device"
	deviceTestInit "github.com/go-magma/magma/orc8r/cloud/go/services/device/test_init"
	"github.com/go-magma/magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	"github.com/go-magma/magma/orc8r/cloud/go/services/state"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const testAgHwID = "Test-AGW-Hw-Id"

type testStateServer struct {
	lastClientIdentity    *protos.Identity
	lastClientCertExpTime int64
}

func NewTestStateServer() (*testStateServer, error) {
	return &testStateServer{}, nil
}

func (srv *testStateServer) GetStates(ctx context.Context, req *protos.GetStatesRequest) (*protos.GetStatesResponse, error) {
	srv.lastClientIdentity =
		proto.Clone(protos.GetClientIdentity(ctx)).(*protos.Identity)
	return nil, nil
}

func (srv *testStateServer) ReportStates(ctx context.Context, req *protos.ReportStatesRequest) (*protos.ReportStatesResponse, error) {
	srv.lastClientIdentity = proto.Clone(protos.GetClientIdentity(ctx)).(*protos.Identity)
	srv.lastClientCertExpTime = protos.GetClientCertExpiration(ctx)
	return &protos.ReportStatesResponse{UnreportedStates: []*protos.IDAndError{}}, nil
}

func (srv *testStateServer) DeleteStates(ctx context.Context, req *protos.DeleteStatesRequest) (*protos.Void, error) {
	srv.lastClientIdentity =
		proto.Clone(protos.GetClientIdentity(ctx)).(*protos.Identity)
	return &protos.Void{}, nil
}

func (srv *testStateServer) SyncStates(ctx context.Context, req *protos.SyncStatesRequest) (*protos.SyncStatesResponse, error) {
	srv.lastClientIdentity = proto.Clone(protos.GetClientIdentity(ctx)).(*protos.Identity)
	srv.lastClientCertExpTime = protos.GetClientCertExpiration(ctx)
	return &protos.SyncStatesResponse{UnsyncedStates: []*protos.IDAndVersion{}}, nil
}

func TestIdentityInjector(t *testing.T) {
	configuratorTestInit.StartTestService(t)
	deviceTestInit.StartTestService(t)
	_ = serde.RegisterSerdes(
		serde.NewBinarySerde(device.SerdeDomain, orc8r.AccessGatewayRecordType, &models.GatewayDevice{}),
		state.NewStateSerde(orc8r.GatewayStateType, &models.GatewayStatus{}),
	)

	networkID := "identity_decorator_test_network"
	configuratorTestUtils.RegisterNetwork(t, networkID, "Identity Decorator Test")
	configuratorTestUtils.RegisterGateway(t, networkID, testAgHwID, &models.GatewayDevice{HardwareID: testAgHwID})

	// Create the service
	srv, err := service.NewTestOrchestratorService(t, orc8r.ModuleName, state.ServiceName)
	assert.NoError(t, err)

	// Add servicers to the service
	stateServer, err := NewTestStateServer()
	assert.NoError(t, err)
	protos.RegisterStateServiceServer(srv.GrpcServer, stateServer)

	l, err := net.Listen("tcp", "")
	assert.NoError(t, err)
	addr := l.Addr().String()
	// Run the service
	go srv.RunTest(l)

	conn, err := registry.GetClientConnection(context.Background(), addr)
	assert.NoError(t, err)
	stateClient := protos.NewStateServiceClient(conn)
	csn := test_utils.StartMockGwAccessControl(t, []string{testAgHwID})
	ctx := metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs("x-magma-client-cert-serial", csn[0]))

	gwState := &models.GatewayStatus{
		Meta: map[string]string{
			"foo": "bar",
		},
	}
	serializedGWStatus, err := serde.Serialize(state.SerdeDomain, orc8r.GatewayStateType, gwState)
	assert.NoError(t, err)
	states := []*protos.State{
		{
			Type:     orc8r.GatewayStateType,
			DeviceID: testAgHwID,
			Value:    serializedGWStatus,
		},
	}
	_, err = stateClient.ReportStates(
		ctx,
		&protos.ReportStatesRequest{States: states},
	)
	assert.NoError(t, err)
	identity := stateServer.lastClientIdentity
	assert.NotNil(t, identity)
	assert.True(t, time.Now().Unix() < stateServer.lastClientCertExpTime)

	cn := identity.ToCommonName()
	assert.NotNil(t, cn)
	assert.Equal(t, *cn, testAgHwID)

	gwid := identity.GetGateway()
	assert.NotNil(t, gwid)
	assert.Equal(t, gwid.HardwareId, testAgHwID)
	assert.Equal(t, gwid.NetworkId, networkID)
	assert.Equal(t, gwid.LogicalId, testAgHwID)

	// Test CTX without any Identification related headers (Identity should
	// not be injected by the middleware)
	_, err = stateClient.ReportStates(
		context.Background(),
		&protos.ReportStatesRequest{States: states},
	)
	assert.NoError(t, err)
	identity = stateServer.lastClientIdentity
	assert.Nil(t, identity)
	assert.Equal(t, int64(0), stateServer.lastClientCertExpTime)

	// Test empty x-magma-client-cert-serial header
	// Hack in the identity context
	ctx = metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs("x-magma-client-cert-serial", ""))
	_, err = stateClient.ReportStates(
		ctx,
		&protos.ReportStatesRequest{States: states},
	)
	assert.Error(t, err)
	assert.Equal(t, int64(0), stateServer.lastClientCertExpTime)

	// Test x-magma-client-cert-cn, but not x-magma-client-cert-serial headers
	ctx = metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs("x-magma-client-cert-cn", "bla bla bla"))
	_, err = stateClient.ReportStates(
		ctx,
		&protos.ReportStatesRequest{States: states},
	)
	assert.Error(t, err)
	assert.Equal(t, int64(0), stateServer.lastClientCertExpTime)

	// Unregister GW
	assert.NoError(
		t,
		configurator.DeleteEntity(networkID, orc8r.MagmadGatewayType, gwid.LogicalId))

	ctx = metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs("x-magma-client-cert-serial", csn[0]))

	// Expect PermissionDenied error now
	_, err = stateClient.ReportStates(
		ctx,
		&protos.ReportStatesRequest{States: states},
	)
	assert.Error(t, err)
	assert.Equal(t, int64(0), stateServer.lastClientCertExpTime)
	assert.Equal(
		t,
		"rpc error: code = PermissionDenied desc = Unregistered Gateway Test-AGW-Hw-Id",
		err.Error())
}
