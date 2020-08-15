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

// Access Control Manager is a service which stores, manages and verifies
// operator Identity objects and their rights to access (read/write) Entities.
package main

import (
	"github.com/go-magma/magma/orc8r/cloud/go/blobstore"
	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	"github.com/go-magma/magma/orc8r/cloud/go/service"
	"github.com/go-magma/magma/orc8r/cloud/go/services/accessd"
	"github.com/go-magma/magma/orc8r/cloud/go/services/accessd/protos"
	"github.com/go-magma/magma/orc8r/cloud/go/services/accessd/servicers"
	"github.com/go-magma/magma/orc8r/cloud/go/services/accessd/storage"
	"github.com/go-magma/magma/orc8r/cloud/go/sqorc"
	storage2 "github.com/go-magma/magma/orc8r/cloud/go/storage"

	"github.com/golang/glog"
)

func main() {
	// Create the service
	srv, err := service.NewOrchestratorService(orc8r.ModuleName, accessd.ServiceName)
	if err != nil {
		glog.Fatalf("Error creating service: %s", err)
	}

	// Init storage
	db, err := sqorc.Open(storage2.SQLDriver, storage2.DatabaseSource)
	if err != nil {
		glog.Fatalf("Failed to connect to database: %s", err)
	}
	fact := blobstore.NewEntStorage(storage.AccessdTableBlobstore, db, sqorc.GetSqlBuilder())
	err = fact.InitializeFactory()
	if err != nil {
		glog.Fatalf("Error initializing accessd database: %s", err)
	}
	store := storage.NewAccessdBlobstore(fact)

	// Add servicers to the service
	accessdServer := servicers.NewAccessdServer(store)
	protos.RegisterAccessControlManagerServer(srv.GrpcServer, accessdServer)

	// Run the service
	err = srv.Run()
	if err != nil {
		glog.Fatalf("Error running service: %s", err)
	}
}
