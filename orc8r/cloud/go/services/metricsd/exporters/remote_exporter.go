package exporters

import (
	"context"

	merrors "github.com/go-magma/magma/lib/go/errors"
	"github.com/go-magma/magma/lib/go/registry"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd/protos"

	"github.com/golang/glog"
)

// remoteExporter identifies a remote metrics exporter.
type remoteExporter struct {
	// service name of the exporter
	// should always be uppercase to match service registry convention
	service string
}

func NewRemoteExporter(serviceName string) Exporter {
	return &remoteExporter{service: serviceName}
}

func (r *remoteExporter) Submit(metrics []MetricAndContext) error {
	c, err := r.getExporterClient()
	if err != nil {
		return err
	}
	_, err = c.Submit(context.Background(), &protos.SubmitMetricsRequest{Metrics: MakeProtoMetrics(metrics)})
	return err
}

func (r *remoteExporter) getExporterClient() (protos.MetricsExporterClient, error) {
	conn, err := registry.GetConnection(r.service)
	if err != nil {
		initErr := merrors.NewInitError(err, r.service)
		glog.Error(initErr)
		return nil, initErr
	}
	return protos.NewMetricsExporterClient(conn), nil
}
