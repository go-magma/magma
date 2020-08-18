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

package pluginimpl

import (
	"github.com/go-magma/magma/lib/go/definitions"
	"github.com/go-magma/magma/lib/go/registry"
	"github.com/go-magma/magma/lib/go/service/config"
	"github.com/go-magma/magma/orc8r/cloud/go/obsidian"
	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	"github.com/go-magma/magma/orc8r/cloud/go/serde"
	"github.com/go-magma/magma/orc8r/cloud/go/services/configurator"
	"github.com/go-magma/magma/orc8r/cloud/go/services/device"
	"github.com/go-magma/magma/orc8r/cloud/go/services/directoryd"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd/collection"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd/exporters"
	"github.com/go-magma/magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	"github.com/go-magma/magma/orc8r/cloud/go/services/state"
	"github.com/go-magma/magma/orc8r/cloud/go/services/state/indexer"
	"github.com/go-magma/magma/orc8r/cloud/go/services/streamer/providers"
)

// BaseOrchestratorPlugin is the OrchestratorPlugin for the orc8r module
type BaseOrchestratorPlugin struct{}

func (*BaseOrchestratorPlugin) GetName() string {
	return orc8r.ModuleName
}

func (*BaseOrchestratorPlugin) GetServices() []registry.ServiceLocation {
	serviceLocations, err := registry.LoadServiceRegistryConfig(orc8r.ModuleName)
	if err != nil {
		return []registry.ServiceLocation{}
	}
	return serviceLocations
}

func (*BaseOrchestratorPlugin) GetSerdes() []serde.Serde {
	return []serde.Serde{
		// State service serdes
		state.NewStateSerde(orc8r.GatewayStateType, &models.GatewayStatus{}),
		// For checkin_cli.py to test cloud < - > gateway connection
		state.NewStateSerde(state.StringMapSerdeType, &state.StringToStringMap{}),
		// For DirectoryD records
		state.NewStateSerde(orc8r.DirectoryRecordType, &directoryd.DirectoryRecord{}),

		// Device service serdes
		serde.NewBinarySerde(device.SerdeDomain, orc8r.AccessGatewayRecordType, &models.GatewayDevice{}),

		// Config manager serdes
		configurator.NewNetworkConfigSerde(orc8r.DnsdNetworkType, &models.NetworkDNSConfig{}),
		configurator.NewNetworkConfigSerde(orc8r.NetworkFeaturesConfig, &models.NetworkFeatures{}),

		configurator.NewNetworkEntityConfigSerde(orc8r.MagmadGatewayType, &models.MagmadGatewayConfigs{}),
		configurator.NewNetworkEntityConfigSerde(orc8r.UpgradeReleaseChannelEntityType, &models.ReleaseChannel{}),
		configurator.NewNetworkEntityConfigSerde(orc8r.UpgradeTierEntityType, &models.Tier{}),
	}
}

func (*BaseOrchestratorPlugin) GetMconfigBuilders() []configurator.MconfigBuilder {
	return []configurator.MconfigBuilder{
		&BaseOrchestratorMconfigBuilder{},
		&DnsdMconfigBuilder{},
	}
}

func (*BaseOrchestratorPlugin) GetMetricsProfiles(metricsConfig *config.ConfigMap) []metricsd.MetricsProfile {
	return getMetricsProfiles()
}

func (*BaseOrchestratorPlugin) GetObsidianHandlers(metricsConfig *config.ConfigMap) []obsidian.Handler {
	return []obsidian.Handler{}
}

func (*BaseOrchestratorPlugin) GetStreamerProviders() []providers.StreamProvider {
	return []providers.StreamProvider{
		providers.NewRemoteProvider(definitions.StreamerServiceName, definitions.MconfigStreamName),
	}
}

func (*BaseOrchestratorPlugin) GetStateIndexers() []indexer.Indexer {
	// TODO(hcgatewood): fix this once k8s polling is enabled -- for now, hard-coding this single indexer as the only remote indexer
	return []indexer.Indexer{
		// From orc8r/cloud/go/services/directoryd/servicers/indexer_servicer.go
		indexer.NewRemoteIndexer(directoryd.ServiceName, 1, orc8r.DirectoryRecordType),
	}
}

const (
	ProfileNamePrometheus = "prometheus"
	ProfileNameExportAll  = "exportall"
)

func getMetricsProfiles() []metricsd.MetricsProfile {
	// Controller profile - 1 collector for each service
	services := registry.ListControllerServices()

	deviceCollectors := []collection.MetricCollector{&collection.DiskUsageMetricCollector{}, &collection.ProcMetricsCollector{}}
	allCollectors := make([]collection.MetricCollector, 0, len(services)+len(deviceCollectors))

	for _, s := range services {
		allCollectors = append(allCollectors, collection.NewCloudServiceMetricCollector(s))
	}
	for _, c := range deviceCollectors {
		allCollectors = append(allCollectors, c)
	}

	// Prometheus profile - Exports all service metric to Prometheus
	prometheusProfile := metricsd.MetricsProfile{
		Name:       ProfileNamePrometheus,
		Collectors: allCollectors,
		Exporters:  []exporters.Exporter{exporters.NewRemoteExporter(metricsd.ServiceName)},
	}

	// ExportAllProfile - Exports to all exporters
	exportAllProfile := metricsd.MetricsProfile{
		Name:       ProfileNameExportAll,
		Collectors: allCollectors,
		Exporters:  []exporters.Exporter{exporters.NewRemoteExporter(metricsd.ServiceName)},
	}

	return []metricsd.MetricsProfile{
		prometheusProfile,
		exportAllProfile,
	}
}
