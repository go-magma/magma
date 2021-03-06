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

package s6a_proxy

import (
	"errors"
	"fmt"

	"github.com/go-magma/magma/modules/feg/cloud/go/protos"
	"github.com/go-magma/magma/modules/feg/cloud/go/services/feg_relay"
	"github.com/go-magma/magma/modules/feg/gateway/registry"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func getCloudConn() (*grpc.ClientConn, error) {
	conn, err := registry.Get().GetCloudConnection(feg_relay.ServiceName)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to establish connection to cloud FegToGwRelayClient: %s", err)
		glog.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return conn, nil
}

// GWS6AProxyCancelLocation forwards CLR to Controller
func GWS6AProxyCancelLocation(in *protos.CancelLocationRequest) (*protos.CancelLocationAnswer, error) {
	conn, err := getCloudConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := protos.NewS6AGatewayServiceClient(conn)
	return client.CancelLocation(context.Background(), in)
}

// GWS6AProxyReset forwards RSR to Controller
func GWS6AProxyReset(in *protos.ResetRequest) (*protos.ResetAnswer, error) {
	conn, err := getCloudConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := protos.NewS6AGatewayServiceClient(conn)
	return client.Reset(context.Background(), in)
}
