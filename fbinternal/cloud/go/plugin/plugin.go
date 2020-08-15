package plugin

import (
	fbinternal_service "github.com/go-magma/fbinternal/cloud/go/services/fbinternal"
	"github.com/go-magma/fbinternal/cloud/go/services/testcontroller"
	"github.com/go-magma/magma/orc8r/cloud/go/services/state/indexer"

	"github.com/go-magma/fbinternal/cloud/go/fbinternal"
	"github.com/go-magma/fbinternal/cloud/go/services/testcontroller/obsidian/models"
	"github.com/go-magma/magma/orc8r/cloud/go/obsidian"
	"github.com/go-magma/magma/orc8r/cloud/go/serde"
	"github.com/go-magma/magma/orc8r/cloud/go/services/configurator"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd/collection"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd/exporters"
	"github.com/go-magma/magma/orc8r/cloud/go/services/streamer/providers"
	"github.com/go-magma/magma/lib/go/registry"
	"github.com/go-magma/magma/lib/go/service/config"
)

type FbinternalOrchestratorPlugin struct{}

func (*FbinternalOrchestratorPlugin) GetName() string {
	return fbinternal.ModuleName
}

func (*FbinternalOrchestratorPlugin) GetServices() []registry.ServiceLocation {
	serviceLocations, err := registry.LoadServiceRegistryConfig(fbinternal.ModuleName)
	if err != nil {
		return []registry.ServiceLocation{}
	}
	return serviceLocations
}

func (*FbinternalOrchestratorPlugin) GetSerdes() []serde.Serde {
	return []serde.Serde{
		configurator.NewNetworkConfigSerde(fbinternal.TestControllerNetworkType, &models.TestConfig{}),
		serde.NewBinarySerde(testcontroller.SerdeDomain, testcontroller.EnodedTestCaseType, &models.EnodebdTestConfig{}),
		serde.NewBinarySerde(testcontroller.SerdeDomain, testcontroller.EnodedTestExcludeTraffic, &models.EnodebdTestConfig{}),
	}
}

func (*FbinternalOrchestratorPlugin) GetMconfigBuilders() []configurator.MconfigBuilder {
	return []configurator.MconfigBuilder{}
}

func (*FbinternalOrchestratorPlugin) GetMetricsProfiles(metricsConfig *config.ConfigMap) []metricsd.MetricsProfile {
	return getMetricsProfiles()
}

func (*FbinternalOrchestratorPlugin) GetObsidianHandlers(metricsConfig *config.ConfigMap) []obsidian.Handler {
	return []obsidian.Handler{}
}

func (*FbinternalOrchestratorPlugin) GetStreamerProviders() []providers.StreamProvider {
	return []providers.StreamProvider{}
}

func (*FbinternalOrchestratorPlugin) GetStateIndexers() []indexer.Indexer {
	return []indexer.Indexer{}
}

const (
	ProfileNameFacebook   = "facebook"
	ProfileNameSys        = "sys"
	ProfileNamePrometheus = "fbprometheus"
	ProfileNameExportAll  = "fbexportall"
)

func getMetricsProfiles() []metricsd.MetricsProfile {
	// Sys profile - collectors for disk usage and metricsd
	sysProfile := metricsd.MetricsProfile{
		Name: ProfileNameSys,
		Collectors: []collection.MetricCollector{
			&collection.DiskUsageMetricCollector{},
			collection.NewCloudServiceMetricCollector(metricsd.ServiceName),
		},
		Exporters: []exporters.Exporter{exporters.NewRemoteExporter(fbinternal_service.ServiceName)},
	}

	// Facebook profile - 1 collector for each service
	allServices := registry.ListControllerServices()
	controllerCollectors := make([]collection.MetricCollector, 0, len(allServices)+1)
	for _, srv := range allServices {
		controllerCollectors = append(controllerCollectors, collection.NewCloudServiceMetricCollector(srv))
	}

	controllerCollectors = append(controllerCollectors, &collection.DiskUsageMetricCollector{})
	facebookProfile := metricsd.MetricsProfile{
		Name:       ProfileNameFacebook,
		Collectors: controllerCollectors,
		Exporters: []exporters.Exporter{
			exporters.NewRemoteExporter(fbinternal_service.ServiceName),
		},
	}

	prometheusProfile := metricsd.MetricsProfile{
		Name:       ProfileNamePrometheus,
		Collectors: controllerCollectors,
		Exporters: []exporters.Exporter{
			exporters.NewRemoteExporter(fbinternal_service.ServiceName),
			exporters.NewRemoteExporter(metricsd.ServiceName),
		},
	}

	allExporterProfile := metricsd.MetricsProfile{
		Name:       ProfileNameExportAll,
		Collectors: controllerCollectors,
		Exporters: []exporters.Exporter{
			exporters.NewRemoteExporter(fbinternal_service.ServiceName),
			exporters.NewRemoteExporter(metricsd.ServiceName),
		},
	}

	return []metricsd.MetricsProfile{
		sysProfile,
		facebookProfile,
		prometheusProfile,
		allExporterProfile,
	}
}
