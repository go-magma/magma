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
	"github.com/go-magma/cwf/cloud/go/cwf"
	"github.com/go-magma/cwf/cloud/go/services/cwf/obsidian/models"
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

// CwfOrchestratorPlugin implements OrchestratorPlugin for the CWF module
type CwfOrchestratorPlugin struct{}

func (*CwfOrchestratorPlugin) GetName() string {
	return cwf.ModuleName
}

func (*CwfOrchestratorPlugin) GetServices() []registry.ServiceLocation {
	serviceLocations, err := registry.LoadServiceRegistryConfig(cwf.ModuleName)
	if err != nil {
		return []registry.ServiceLocation{}
	}
	return serviceLocations
}

func (*CwfOrchestratorPlugin) GetSerdes() []serde.Serde {
	return []serde.Serde{
		configurator.NewNetworkConfigSerde(cwf.CwfNetworkType, &models.NetworkCarrierWifiConfigs{}),
		configurator.NewNetworkEntityConfigSerde(cwf.CwfGatewayType, &models.GatewayCwfConfigs{}),
		state.NewStateSerde(cwf.CwfSubscriberDirectoryType, &models.CwfSubscriberDirectoryRecord{}),
		state.NewStateSerde(cwf.CwfClusterHealthType, &models.CarrierWifiNetworkClusterStatus{}),
		state.NewStateSerde(cwf.CwfGatewayHealthType, &models.CarrierWifiGatewayHealthStatus{}),
	}
}

func (*CwfOrchestratorPlugin) GetMconfigBuilders() []configurator.MconfigBuilder {
	return []configurator.MconfigBuilder{
		&Builder{},
	}
}

func (*CwfOrchestratorPlugin) GetMetricsProfiles(metricsConfig *config.ConfigMap) []metricsd.MetricsProfile {
	return []metricsd.MetricsProfile{}
}

func (*CwfOrchestratorPlugin) GetObsidianHandlers(metricsConfig *config.ConfigMap) []obsidian.Handler {
	return []obsidian.Handler{}
}

func (*CwfOrchestratorPlugin) GetStreamerProviders() []providers.StreamProvider {
	return []providers.StreamProvider{}
}

func (*CwfOrchestratorPlugin) GetStateIndexers() []indexer.Indexer {
	return []indexer.Indexer{}
}
