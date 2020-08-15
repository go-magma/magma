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

package main

import (
	"github.com/go-magma/magma/modules/feg/cloud/go/feg"
	feg_service "github.com/go-magma/magma/modules/feg/cloud/go/services/feg"
	"github.com/go-magma/magma/modules/feg/cloud/go/services/feg/obsidian/handlers"
	"github.com/go-magma/magma/orc8r/cloud/go/obsidian"
	"github.com/go-magma/magma/orc8r/cloud/go/service"

	"github.com/golang/glog"
)

func main() {
	srv, err := service.NewOrchestratorService(feg.ModuleName, feg_service.ServiceName)
	if err != nil {
		glog.Fatalf("Error creating feg service %s", err)
	}
	obsidian.AttachHandlers(srv.EchoServer, handlers.GetHandlers())
	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error while running service and echo server: %s", err)
	}
}
