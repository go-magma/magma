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

package protos_test

import (
	"testing"

	"github.com/go-magma/magma/modules/lte/cloud/go/protos/mconfig"
	"github.com/go-magma/magma/modules/lte/cloud/go/services/cellular/protos"

	"github.com/stretchr/testify/assert"
)

func TestGetNetworkServiceName(t *testing.T) {
	name, err := protos.GetNetworkServiceName(protos.NetworkEPCConfig_ENFORCEMENT)
	assert.NoError(t, err)
	assert.Equal(t, name, "policy_enforcement")

	_, err = protos.GetNetworkServiceName(99999999)
	assert.Error(t, err)
}

func TestGetNetworkServiceEnum(t *testing.T) {
	enum, err := protos.GetNetworkServiceEnum("policy_enforcement")
	assert.NoError(t, err)
	assert.Equal(t, enum, protos.NetworkEPCConfig_ENFORCEMENT)

	_, err = protos.GetNetworkServiceEnum("unknown enum")
	assert.Error(t, err)
}

func TestGetPipelineDServicesConfig(t *testing.T) {
	// Non-default service set -> that subset seen in mconfig
	apps, err := protos.GetPipelineDServicesConfig([]protos.NetworkEPCConfig_NetworkServices{
		protos.NetworkEPCConfig_ENFORCEMENT,
	})
	assert.NoError(t, err)
	assert.Equal(t, apps, []mconfig.PipelineD_NetworkServices{
		mconfig.PipelineD_ENFORCEMENT,
	})

	// Default service set -> default set seen in mconfig
	apps, err = protos.GetPipelineDServicesConfig([]protos.NetworkEPCConfig_NetworkServices{})
	assert.NoError(t, err)
	assert.Equal(t, apps, []mconfig.PipelineD_NetworkServices{
		mconfig.PipelineD_ENFORCEMENT,
	})

	// Unrecognized service -> err
	_, err = protos.GetPipelineDServicesConfig([]protos.NetworkEPCConfig_NetworkServices{99999})
	assert.Error(t, err, "Unknown pipeline service enum: 99999")
}
