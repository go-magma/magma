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

//go:generate sh $MAGMA_ROOT/lib/go/protos/gen_protos.sh ../../../protos -I$MAGMA_ROOT/protos/prometheus -I$MAGMA_ROOT/protos
//go:generate sh $MAGMA_ROOT/lib/go/protos/gen_protos_sub.sh ../../../protos oai
package protos
