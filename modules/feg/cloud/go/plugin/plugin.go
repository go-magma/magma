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

// Package plugin exposes the OrchestratorPlugin implementation for the module.
// This is so unit tests can register the plugin without building and loading
// it from disk.
package plugin

import (
	"github.com/go-magma/magma/lib/go/registry"
	"github.com/go-magma/magma/lib/go/service/config"
	"github.com/go-magma/magma/modules/feg/cloud/go/feg"
	"github.com/go-magma/magma/modules/feg/cloud/go/services/feg/obsidian/models"
	"github.com/go-magma/magma/orc8r/cloud/go/obsidian"
	"github.com/go-magma/magma/orc8r/cloud/go/serde"
	"github.com/go-magma/magma/orc8r/cloud/go/services/configurator"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd"
	"github.com/go-magma/magma/orc8r/cloud/go/services/state/indexer"
	"github.com/go-magma/magma/orc8r/cloud/go/services/streamer/providers"
)

// FegOrchestratorPlugin is an implementation of OrchestratorPlugin for the
// feg module
type FegOrchestratorPlugin struct{}

func (*FegOrchestratorPlugin) GetName() string {
	return feg.ModuleName
}

func (*FegOrchestratorPlugin) GetServices() []registry.ServiceLocation {
	serviceLocations, err := registry.LoadServiceRegistryConfig(feg.ModuleName)
	if err != nil {
		return []registry.ServiceLocation{}
	}
	return serviceLocations
}

func (*FegOrchestratorPlugin) GetSerdes() []serde.Serde {
	return []serde.Serde{
		// configurator serdes
		configurator.NewNetworkConfigSerde(feg.FegNetworkType, &models.NetworkFederationConfigs{}),
		configurator.NewNetworkConfigSerde(feg.FederatedNetworkType, &models.FederatedNetworkConfigs{}),
		configurator.NewNetworkEntityConfigSerde(feg.FegGatewayType, &models.GatewayFederationConfigs{}),
	}
}

func (*FegOrchestratorPlugin) GetMconfigBuilders() []configurator.MconfigBuilder {
	return []configurator.MconfigBuilder{
		&Builder{},
	}
}

func (*FegOrchestratorPlugin) GetMetricsProfiles(metricsConfig *config.ConfigMap) []metricsd.MetricsProfile {
	return []metricsd.MetricsProfile{}
}

func (*FegOrchestratorPlugin) GetObsidianHandlers(metricsConfig *config.ConfigMap) []obsidian.Handler {
	return []obsidian.Handler{}
}

func (*FegOrchestratorPlugin) GetStreamerProviders() []providers.StreamProvider {
	return []providers.StreamProvider{}
}

func (*FegOrchestratorPlugin) GetStateIndexers() []indexer.Indexer {
	return []indexer.Indexer{}
}
