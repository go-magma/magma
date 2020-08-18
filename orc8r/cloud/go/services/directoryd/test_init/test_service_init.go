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

	"github.com/go-magma/magma/lib/go/protos"
	"github.com/go-magma/magma/orc8r/cloud/go/blobstore"
	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	"github.com/go-magma/magma/orc8r/cloud/go/services/directoryd"
	"github.com/go-magma/magma/orc8r/cloud/go/services/directoryd/servicers"
	"github.com/go-magma/magma/orc8r/cloud/go/services/directoryd/storage"
	"github.com/go-magma/magma/orc8r/cloud/go/sqorc"
	"github.com/go-magma/magma/orc8r/cloud/go/test_utils"

	"github.com/stretchr/testify/assert"
)

func StartTestService(t *testing.T) {
	// Create service
	srv, lis := test_utils.NewTestService(t, orc8r.ModuleName, directoryd.ServiceName)

	// Init storage
	db, err := sqorc.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	fact := blobstore.NewSQLBlobStorageFactory(storage.DirectorydTableBlobstore, db, sqorc.GetSqlBuilder())
	err = fact.InitializeFactory()
	assert.NoError(t, err)
	store := storage.NewDirectorydBlobstore(fact)

	// Add servicers
	directoryServicer, err := servicers.NewDirectoryLookupServicer(store)
	assert.NoError(t, err)
	protos.RegisterDirectoryLookupServer(srv.GrpcServer, directoryServicer)
	protos.RegisterGatewayDirectoryServiceServer(srv.GrpcServer, servicers.NewDirectoryUpdateServicer())

	// Run service
	go srv.RunTest(lis)
}
