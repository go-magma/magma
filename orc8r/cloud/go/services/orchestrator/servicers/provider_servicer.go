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

package servicers

import (
	"context"
	"fmt"

	"github.com/go-magma/magma/lib/go/definitions"
	"github.com/go-magma/magma/lib/go/protos"
	streamer_protos "github.com/go-magma/magma/orc8r/cloud/go/services/streamer/protos"
	"github.com/go-magma/magma/orc8r/cloud/go/services/streamer/providers"
)

type orchestratorStreamProviderServicer struct{}

func NewOrchestratorStreamProviderServicer() streamer_protos.StreamProviderServer {
	return &orchestratorStreamProviderServicer{}
}

func (s *orchestratorStreamProviderServicer) GetUpdates(ctx context.Context, req *protos.StreamRequest) (*protos.DataUpdateBatch, error) {
	var streamer providers.StreamProvider
	switch req.GetStreamName() {
	case definitions.MconfigStreamName:
		streamer = &providers.MconfigProvider{}
	default:
		return nil, fmt.Errorf("GetUpdates failed: unknown stream name provided: %s", req.GetStreamName())
	}

	update, err := streamer.GetUpdates(req.GetGatewayId(), req.GetExtraArgs())
	if err != nil {
		// Note: return blank err to properly receive EAGAINs from mconfig provider
		return nil, err
	}
	return &protos.DataUpdateBatch{Updates: update}, nil
}
