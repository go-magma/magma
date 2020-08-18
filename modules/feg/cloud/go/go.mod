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
module github.com/go-magma/magma/modules/feg/cloud/go

replace (
	github.com/go-magma/magma/gateway/go => ../../../../gateway/go
	github.com/go-magma/magma/lib/go => ../../../../lib/go
	github.com/go-magma/magma/lib/go/protos => ../../../../lib/go/protos
	github.com/go-magma/magma/orc8r/cloud/go => ../../../../orc8r/cloud/go
)

require (
	github.com/Masterminds/squirrel v1.1.1-0.20190513200039-d13326f0be73
	github.com/go-magma/magma/lib/go v0.0.0
	github.com/go-magma/magma/lib/go/protos v0.0.0
	github.com/go-magma/magma/modules/feg/cloud/go/protos v0.0.0-20200818162846-8db662db3aa7
	github.com/go-magma/magma/modules/lte/cloud/go v0.0.0-20200818162846-8db662db3aa7
	github.com/go-magma/magma/orc8r/cloud/go v0.0.0
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/strfmt v0.19.4
	github.com/go-openapi/swag v0.18.0
	github.com/go-openapi/validate v0.18.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.1
	github.com/labstack/echo v0.0.0-20181123063414-c54d9e8eed6c
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.2.1
	github.com/stretchr/testify v1.4.0
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/grpc v1.27.1
	google.golang.org/protobuf v1.25.0
)

go 1.13
