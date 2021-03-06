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

// File registry.go provides a metrics exporter registry by forwarding calls to
// the service registry.

package metricsd

import (
	"github.com/go-magma/magma/lib/go/registry"
	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd/exporters"
)

// GetMetricsExporters returns all registered metrics exporters.
func GetMetricsExporters() []exporters.Exporter {
	services := registry.FindServices(orc8r.MetricsExporterLabel)

	var ret []exporters.Exporter
	for _, s := range services {
		ret = append(ret, exporters.NewRemoteExporter(s))
	}

	return ret
}
