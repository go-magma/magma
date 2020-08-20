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

package storage

import (
	"time"

	"github.com/go-magma/magma/lib/go/protos"
	fegprotos "github.com/go-magma/magma/modules/feg/cloud/go/protos"
	"github.com/go-magma/magma/modules/feg/cloud/go/services/health"
	"github.com/go-magma/magma/orc8r/cloud/go/blobstore"
	"github.com/go-magma/magma/orc8r/cloud/go/clock"
)

// HealthToBlob converts a gatewayID and healthStats proto to a Blobstore blob
func HealthToBlob(gatewayID string, healthStats *fegprotos.HealthStats) (blobstore.Blob, error) {
	marshaledHealth, err := protos.Marshal(healthStats)
	if err != nil {
		return blobstore.Blob{}, err
	}
	return blobstore.Blob{
		Type:  health.HealthStatusType,
		Key:   gatewayID,
		Value: marshaledHealth,
	}, nil
}

// ClusterToBlob converts a clusterID and activeID to a Blobstore blob
func ClusterToBlob(clusterID string, activeID string) (blobstore.Blob, error) {
	clusterState := &fegprotos.ClusterState{
		ActiveGatewayLogicalId: activeID,
		Time:                   uint64(clock.Now().UnixNano()) / uint64(time.Millisecond),
	}
	marsheledCluster, err := protos.Marshal(clusterState)
	if err != nil {
		return blobstore.Blob{}, err
	}
	return blobstore.Blob{
		Type:  health.ClusterStatusType,
		Key:   clusterID,
		Value: marsheledCluster,
	}, nil
}
