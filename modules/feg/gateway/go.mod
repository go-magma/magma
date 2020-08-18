// Copyright 2020 The Magma Authors.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
module github.com/go-magma/magma/modules/feg/gateway

replace (
	github.com/fiorix/go-diameter/v4 => github.com/emakeev/go-diameter/v4 v4.0.7

	github.com/go-magma/magma/gateway/go => ../../../gateway/go
	github.com/go-magma/magma/lib/go => ../../../lib/go
	github.com/go-magma/magma/lib/go/protos => ../../../lib/go/protos
	github.com/go-magma/magma/modules/feg/cloud/go => ../cloud/go
	github.com/go-magma/magma/modules/feg/cloud/go/protos => ../cloud/go/protos
	github.com/go-magma/magma/modules/lte/cloud/go => ../../lte/cloud/go
	github.com/go-magma/magma/orc8r/cloud/go => ../../../orc8r/cloud/go
)

require (
	github.com/emakeev/snowflake v0.0.0-20200206205012-767080b052fe
	github.com/fiorix/go-diameter/v4 v4.0.1-0.20200120193412-55a1c21738f9
	github.com/go-magma/magma/gateway/go v0.0.0
	github.com/go-magma/magma/lib/go v0.0.0
	github.com/go-magma/magma/lib/go/protos v0.0.0
	github.com/go-magma/magma/modules/feg/cloud/go v0.0.0
	github.com/go-magma/magma/modules/feg/cloud/go/protos v0.0.0
	github.com/go-magma/magma/modules/lte/cloud/go v0.0.0
	github.com/go-magma/magma/orc8r/cloud/go v0.0.0
	github.com/go-openapi/swag v0.18.0
	github.com/go-redis/redis v6.14.1+incompatible
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.1
	github.com/ishidawataru/sctp v0.0.0-20190922091402-408ec287e38c
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.9.1
	github.com/shirou/gopsutil v2.20.3+incompatible
	github.com/stretchr/testify v1.4.0
	github.com/thoas/go-funk v0.7.0
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/grpc v1.27.1
	google.golang.org/protobuf v1.25.0
	layeh.com/radius v0.0.0-20200615152116-663b41c3bf86
)

go 1.13
