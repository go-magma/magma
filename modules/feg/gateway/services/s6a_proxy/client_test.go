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

package s6a_proxy_test

import (
	"testing"

	"github.com/go-magma/magma/modules/feg/cloud/go/protos"
	"github.com/go-magma/magma/modules/feg/gateway/registry"
	"github.com/go-magma/magma/modules/feg/gateway/service_health"
	"github.com/go-magma/magma/modules/feg/gateway/services/s6a_proxy"
	"github.com/go-magma/magma/modules/feg/gateway/services/s6a_proxy/servicers/test"
	"github.com/go-magma/magma/modules/feg/gateway/services/s6a_proxy/test_init"
)

func TestS6aProxyClient(t *testing.T) {
	err := test_init.StartTestService(t)
	if err != nil {
		t.Fatal(err)
		return
	}

	req := &protos.AuthenticationInformationRequest{
		UserName:                   test.TEST_IMSI,
		VisitedPlmn:                []byte(test.TEST_PLMN_ID),
		NumRequestedEutranVectors:  3,
		ImmediateResponsePreferred: true,
	}
	// AIR
	r, err := s6a_proxy.AuthenticationInformation(req)
	if err != nil {
		t.Fatalf("GRPC AIR Error: %v", err)
		return
	}
	t.Logf("GRPC AIA: %s", r)
	if r.ErrorCode != protos.ErrorCode_UNDEFINED {
		t.Errorf("Unexpected AIA Error Code: %d", r.ErrorCode)
	}
	if len(r.EutranVectors) != 3 {
		t.Errorf("Unexpected Number of EutranVectors: %d, Expected: 3", len(r.EutranVectors))
	}

	ulReq := &protos.UpdateLocationRequest{
		UserName:           test.TEST_IMSI,
		VisitedPlmn:        []byte(test.TEST_PLMN_ID),
		SkipSubscriberData: false,
		InitialAttach:      true,
	}
	// ULR
	ulResp, err := s6a_proxy.UpdateLocation(ulReq)
	if err != nil {
		t.Fatalf("GRPC ULR Error: %v", err)
		return
	}
	t.Logf("GRPC ULA: %s", ulResp)
	if ulResp.ErrorCode != protos.ErrorCode_UNDEFINED {
		t.Errorf("Unexpected ULA Error Code: %d", r.ErrorCode)
	}

	puReq := &protos.PurgeUERequest{
		UserName: test.TEST_IMSI,
	}
	// PUR
	puResp, err := s6a_proxy.PurgeUE(puReq)
	if err != nil {
		t.Fatalf("GRPC PUR Error: %v", err)
	}
	t.Logf("GRPC PUA: %s", puResp)
	if puResp.ErrorCode != protos.ErrorCode_SUCCESS {
		t.Errorf("Unexpected PUA Error Code: %d", r.ErrorCode)
	}

	// Disable connections and ensure subsequent requests fail
	disableReq := &protos.DisableMessage{
		DisablePeriodSecs: 10,
	}
	err = service_health.Disable(registry.S6A_PROXY, disableReq)
	if err != nil {
		t.Fatalf("GRPC ServiceHealth Disable Error: %v", err)
		return
	}

	// AIR should fail
	_, err = s6a_proxy.AuthenticationInformation(req)
	if err == nil {
		t.Errorf("AIR Succeeded, but should have failed due to disabled connections")
	}

	// Enable connections
	err = service_health.Enable(registry.S6A_PROXY)
	if err != nil {
		t.Fatalf("GRPC ServiceHealth Enable Error: %v", err)
		return
	}

	// ULR should pass now
	ulResp, err = s6a_proxy.UpdateLocation(ulReq)
	if err != nil {
		t.Fatalf("GRPC ULR Error: %v", err)
		return
	}
	t.Logf("GRPC ULA: %s", ulResp)
	if ulResp.ErrorCode != protos.ErrorCode_UNDEFINED {
		t.Errorf("Unexpected ULA Error Code: %d", ulResp.ErrorCode)
	}
}
