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
	"strconv"
	"testing"

	"github.com/go-magma/magma/lib/go/metrics"
	"github.com/go-magma/magma/modules/cwf/cloud/go/services/analytics"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestGetXAPCalculations(t *testing.T) {
	xapGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: activeUsersMetricName}, []string{analytics.DaysLabel, metrics.NetworkLabelName})
	calcs := getXAPCalculations([]int{1, 7, 30}, xapGauge, "metricName")
	for _, calc := range calcs {
		xapCalc := calc.(*analytics.XAPCalculation)
		assert.Equal(t, strconv.Itoa(xapCalc.Days), xapCalc.Labels[analytics.DaysLabel])
	}
}

func TestGetUserThroughputCalculations(t *testing.T) {
	userThroughputGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: userThroughputMetricName}, []string{analytics.DaysLabel, metrics.NetworkLabelName, analytics.DirectionLabel})
	calcs := getUserThroughputCalculations([]int{1, 7, 30}, userThroughputGauge, "metricName")
	for _, calc := range calcs {
		c := calc.(*analytics.UserThroughputCalculation)
		assert.Equal(t, strconv.Itoa(c.Days), c.Labels[analytics.DaysLabel])
	}
}

func TestGetUserConsumptionCalculations(t *testing.T) {
	userConsumptionGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: userConsumptionMetricName}, []string{analytics.DaysLabel, metrics.NetworkLabelName, analytics.DirectionLabel})
	calcs := getUserConsumptionCalculations([]int{1, 7, 30}, userConsumptionGauge, "metricName")
	for _, calc := range calcs {
		c := calc.(*analytics.UserConsumptionCalculation)
		assert.Equal(t, strconv.Itoa(c.Days), c.Labels[analytics.DaysLabel])
	}
}

func TestGetAPThroughputCalculations(t *testing.T) {
	apThroughputGauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{Name: apThroughputMetricName}, []string{analytics.DaysLabel, metrics.NetworkLabelName, analytics.DirectionLabel, analytics.APNLabel})
	calcs := getAPThroughputCalculations([]int{1, 7, 30}, apThroughputGauge, "metricName")
	for _, calc := range calcs {
		c := calc.(*analytics.APThroughputCalculation)
		assert.Equal(t, strconv.Itoa(c.Days), c.Labels[analytics.DaysLabel])
	}
}
