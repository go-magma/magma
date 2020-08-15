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

package test_init

import (
	"testing"

	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd"
	"github.com/go-magma/magma/orc8r/cloud/go/services/metricsd/protos"
	"github.com/go-magma/magma/orc8r/cloud/go/test_utils"
)

func StartTestServiceInternal(t *testing.T, exporter protos.MetricsExporterServer) {
	srv, lis := test_utils.NewTestService(t, orc8r.ModuleName, metricsd.ServiceName)
	protos.RegisterMetricsExporterServer(srv.GrpcServer, exporter)
	go srv.RunTest(lis)
}
