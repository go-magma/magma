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

package plugin

import (
	"github.com/go-magma/devmand/cloud/go/devmand"
	"github.com/go-magma/devmand/cloud/go/services/devmand/obsidian/models"
	"github.com/go-magma/magma/orc8r/cloud/go/obsidian"
	"github.com/go-magma/magma/orc8r/cloud/go/serde"
	"github.com/go-magma/magma/orc8r/cloud/go/services/configurator"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd"
	"github.com/go-magma/magma/orc8r/cloud/go/services/state"
	"github.com/go-magma/magma/orc8r/cloud/go/services/state/indexer"
	"github.com/go-magma/magma/orc8r/cloud/go/services/streamer/providers"
	"github.com/go-magma/magma/lib/go/registry"
	"github.com/go-magma/magma/lib/go/service/config"
)

// DevmandOrchestratorPlugin is the orchestrator plugin for devmand
type DevmandOrchestratorPlugin struct{}

// GetName gets the name of the devmand module
func (*DevmandOrchestratorPlugin) GetName() string {
	return devmand.ModuleName
}

// GetServices gets the devmand service locations
func (*DevmandOrchestratorPlugin) GetServices() []registry.ServiceLocation {
	serviceLocations, err := registry.LoadServiceRegistryConfig(devmand.ModuleName)
	if err != nil {
		return []registry.ServiceLocation{}
	}
	return serviceLocations
}

// GetSerdes gets the devmand serializers and deserializers
func (*DevmandOrchestratorPlugin) GetSerdes() []serde.Serde {
	return []serde.Serde{
		state.NewStateSerde(devmand.SymphonyDeviceStateType, &models.SymphonyDeviceState{}),
		configurator.NewNetworkEntityConfigSerde(devmand.SymphonyDeviceType, &models.SymphonyDeviceConfig{}),
	}
}

func (*DevmandOrchestratorPlugin) GetMconfigBuilders() []configurator.MconfigBuilder {
	return []configurator.MconfigBuilder{
		&Builder{},
	}
}

// GetMetricsProfiles gets the metricsd profiles
func (*DevmandOrchestratorPlugin) GetMetricsProfiles(metricsConfig *config.ConfigMap) []metricsd.MetricsProfile {
	return []metricsd.MetricsProfile{}
}

// GetObsidianHandlers gets the devmand obsidian handlers
func (*DevmandOrchestratorPlugin) GetObsidianHandlers(metricsConfig *config.ConfigMap) []obsidian.Handler {
	return []obsidian.Handler{}
}

// GetStreamerProviders gets the stream providers
func (*DevmandOrchestratorPlugin) GetStreamerProviders() []providers.StreamProvider {
	return []providers.StreamProvider{}
}

func (*DevmandOrchestratorPlugin) GetStateIndexers() []indexer.Indexer {
	return []indexer.Indexer{}
}
