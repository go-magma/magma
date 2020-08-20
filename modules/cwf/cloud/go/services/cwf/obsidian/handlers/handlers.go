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

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	merrors "github.com/go-magma/magma/lib/go/errors"
	"github.com/go-magma/magma/modules/cwf/cloud/go/cwf"
	cwfModels "github.com/go-magma/magma/modules/cwf/cloud/go/services/cwf/obsidian/models"
	fegModels "github.com/go-magma/magma/modules/feg/cloud/go/services/feg/obsidian/models"
	lteHandlers "github.com/go-magma/magma/modules/lte/cloud/go/services/lte/obsidian/handlers"
	policyModels "github.com/go-magma/magma/modules/lte/cloud/go/services/policydb/obsidian/models"
	"github.com/go-magma/magma/orc8r/cloud/go/models"
	"github.com/go-magma/magma/orc8r/cloud/go/obsidian"
	"github.com/go-magma/magma/orc8r/cloud/go/orc8r"
	"github.com/go-magma/magma/orc8r/cloud/go/services/configurator"
	"github.com/go-magma/magma/orc8r/cloud/go/services/directoryd"
	"github.com/go-magma/magma/orc8r/cloud/go/services/orchestrator/obsidian/handlers"
	orc8rModels "github.com/go-magma/magma/orc8r/cloud/go/services/orchestrator/obsidian/models"
	"github.com/go-magma/magma/orc8r/cloud/go/services/state"
	"github.com/go-magma/magma/orc8r/cloud/go/storage"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const (
	CwfNetworks                    = "cwf"
	ListNetworksPath               = obsidian.V1Root + CwfNetworks
	ManageNetworkPath              = ListNetworksPath + "/:network_id"
	ManageNetworkNamePath          = ManageNetworkPath + obsidian.UrlSep + "name"
	ManageNetworkDescriptionPath   = ManageNetworkPath + obsidian.UrlSep + "description"
	ManageNetworkFeaturesPath      = ManageNetworkPath + obsidian.UrlSep + "features"
	ManageNetworkDNSPath           = ManageNetworkPath + obsidian.UrlSep + "dns"
	ManageNetworkCarrierWifiPath   = ManageNetworkPath + obsidian.UrlSep + "carrier_wifi"
	ManageNetworkFederationPath    = ManageNetworkPath + obsidian.UrlSep + "federation"
	ManageNetworkSubscriberPath    = ManageNetworkPath + obsidian.UrlSep + "subscriber_config"
	ManageNetworkBaseNamesPath     = ManageNetworkSubscriberPath + obsidian.UrlSep + "base_names"
	ManageNetworkRuleNamesPath     = ManageNetworkSubscriberPath + obsidian.UrlSep + "rule_names"
	ManageNetworkBaseNamePath      = ManageNetworkBaseNamesPath + obsidian.UrlSep + ":base_name"
	ManageNetworkRuleNamePath      = ManageNetworkRuleNamesPath + obsidian.UrlSep + ":rule_id"
	ManageNetworkClusterStatusPath = ManageNetworkPath + obsidian.UrlSep + "cluster_status"
	ManageNetworkLiUesPath         = ManageNetworkPath + obsidian.UrlSep + ":li_ues"

	Gateways                      = "gateways"
	ListGatewaysPath              = ManageNetworkPath + obsidian.UrlSep + Gateways
	ManageGatewayPath             = ListGatewaysPath + obsidian.UrlSep + ":gateway_id"
	ManageGatewayNamePath         = ManageGatewayPath + obsidian.UrlSep + "name"
	ManageGatewayDescriptionPath  = ManageGatewayPath + obsidian.UrlSep + "description"
	ManageGatewayConfigPath       = ManageGatewayPath + obsidian.UrlSep + "magmad"
	ManageGatewayDevicePath       = ManageGatewayPath + obsidian.UrlSep + "device"
	ManageGatewayStatePath        = ManageGatewayPath + obsidian.UrlSep + "status"
	ManageGatewayTierPath         = ManageGatewayPath + obsidian.UrlSep + "tier"
	ManageGatewayCarrierWifiPath  = ManageGatewayPath + obsidian.UrlSep + "carrier_wifi"
	ManageGatewayHealthStatusPath = ManageGatewayPath + obsidian.UrlSep + "health_status"

	Subscribers                   = "subscribers"
	BaseSubscriberPath            = ManageNetworkPath + obsidian.UrlSep + Subscribers + obsidian.UrlSep + ":subscriber_id"
	SubscriberDirectoryRecordPath = BaseSubscriberPath + obsidian.UrlSep + "directory_record"
)

func GetHandlers() []obsidian.Handler {
	ret := []obsidian.Handler{
		handlers.GetListGatewaysHandler(ListGatewaysPath, cwf.CwfGatewayType, makeCwfGateways),
		{Path: ListGatewaysPath, Methods: obsidian.POST, HandlerFunc: createGateway},
		{Path: ManageGatewayPath, Methods: obsidian.GET, HandlerFunc: getGateway},
		{Path: ManageGatewayPath, Methods: obsidian.PUT, HandlerFunc: updateGateway},
		handlers.GetDeleteGatewayHandler(ManageGatewayPath, cwf.CwfGatewayType),

		{Path: ManageGatewayStatePath, Methods: obsidian.GET, HandlerFunc: handlers.GetStateHandler},
		{Path: ManageNetworkClusterStatusPath, Methods: obsidian.GET, HandlerFunc: getClusterStatusHandler},
		{Path: ManageGatewayHealthStatusPath, Methods: obsidian.GET, HandlerFunc: getHealthStatusHandler},
		{Path: SubscriberDirectoryRecordPath, Methods: obsidian.GET, HandlerFunc: getSubscriberDirectoryHandler},

		{Path: ManageNetworkBaseNamePath, Methods: obsidian.POST, HandlerFunc: lteHandlers.AddNetworkWideSubscriberBaseName},
		{Path: ManageNetworkRuleNamePath, Methods: obsidian.POST, HandlerFunc: lteHandlers.AddNetworkWideSubscriberRuleName},
		{Path: ManageNetworkBaseNamePath, Methods: obsidian.DELETE, HandlerFunc: lteHandlers.RemoveNetworkWideSubscriberBaseName},
		{Path: ManageNetworkRuleNamePath, Methods: obsidian.DELETE, HandlerFunc: lteHandlers.RemoveNetworkWideSubscriberRuleName},
	}

	ret = append(ret, handlers.GetTypedNetworkCRUDHandlers(ListNetworksPath, ManageNetworkPath, cwf.CwfNetworkType, &cwfModels.CwfNetwork{})...)

	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkNamePath, new(models.NetworkName), "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkDescriptionPath, new(models.NetworkDescription), "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkFeaturesPath, &orc8rModels.NetworkFeatures{}, orc8r.NetworkFeaturesConfig)...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkDNSPath, &orc8rModels.NetworkDNSConfig{}, orc8r.DnsdNetworkType)...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkCarrierWifiPath, &cwfModels.NetworkCarrierWifiConfigs{}, cwf.CwfNetworkType)...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkFederationPath, &fegModels.FederatedNetworkConfigs{}, cwf.CwfNetworkType)...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkSubscriberPath, &policyModels.NetworkSubscriberConfig{}, "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkRuleNamesPath, new(policyModels.RuleNames), "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkBaseNamesPath, new(policyModels.BaseNames), "")...)
	ret = append(ret, handlers.GetPartialNetworkHandlers(ManageNetworkLiUesPath, new(cwfModels.LiUes), "")...)

	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayNamePath, new(models.GatewayName))...)
	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayDescriptionPath, new(models.GatewayDescription))...)
	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayConfigPath, &orc8rModels.MagmadGatewayConfigs{})...)
	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayTierPath, new(orc8rModels.TierID))...)
	ret = append(ret, handlers.GetGatewayDeviceHandlers(ManageGatewayDevicePath)...)
	ret = append(ret, handlers.GetPartialGatewayHandlers(ManageGatewayCarrierWifiPath, &cwfModels.GatewayCwfConfigs{})...)

	return ret
}

func createGateway(c echo.Context) error {
	if nerr := handlers.CreateMagmadGatewayFromModel(c, &cwfModels.MutableCwfGateway{}); nerr != nil {
		return nerr
	}
	return c.NoContent(http.StatusCreated)
}

func getGateway(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	magmadModel, nerr := handlers.LoadMagmadGatewayModel(nid, gid)
	if nerr != nil {
		return nerr
	}

	ent, err := configurator.LoadEntity(
		nid, cwf.CwfGatewayType, gid,
		configurator.EntityLoadCriteria{LoadConfig: true, LoadAssocsFromThis: false},
	)
	if err != nil {
		return obsidian.HttpError(errors.Wrap(err, "failed to load cwf gateway"), http.StatusInternalServerError)
	}

	ret := &cwfModels.CwfGateway{
		ID:          magmadModel.ID,
		Name:        magmadModel.Name,
		Description: magmadModel.Description,
		Device:      magmadModel.Device,
		Status:      magmadModel.Status,
		Tier:        magmadModel.Tier,
		Magmad:      magmadModel.Magmad,
	}
	if ent.Config != nil {
		ret.CarrierWifi = ent.Config.(*cwfModels.GatewayCwfConfigs)
	}

	return c.JSON(http.StatusOK, ret)
}

func updateGateway(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}
	if nerr = handlers.UpdateMagmadGatewayFromModel(c, nid, gid, &cwfModels.MutableCwfGateway{}); nerr != nil {
		return nerr
	}
	return c.NoContent(http.StatusNoContent)
}

type cwfAndMagmadGateway struct {
	magmadGateway, cwfGateway configurator.NetworkEntity
}

func makeCwfGateways(
	entsByTK map[storage.TypeAndKey]configurator.NetworkEntity,
	devicesByID map[string]interface{},
	statusesByID map[string]*orc8rModels.GatewayStatus,
) map[string]handlers.GatewayModel {
	gatewayEntsByKey := map[string]*cwfAndMagmadGateway{}
	for tk, ent := range entsByTK {
		existing, found := gatewayEntsByKey[tk.Key]
		if !found {
			existing = &cwfAndMagmadGateway{}
			gatewayEntsByKey[tk.Key] = existing
		}

		switch ent.Type {
		case orc8r.MagmadGatewayType:
			existing.magmadGateway = ent
		case cwf.CwfGatewayType:
			existing.cwfGateway = ent
		}
	}

	ret := make(map[string]handlers.GatewayModel, len(gatewayEntsByKey))
	for key, ents := range gatewayEntsByKey {
		hwID := ents.magmadGateway.PhysicalID
		var devCasted *orc8rModels.GatewayDevice
		if devicesByID[hwID] != nil {
			devCasted = devicesByID[hwID].(*orc8rModels.GatewayDevice)
		}
		ret[key] = (&cwfModels.CwfGateway{}).FromBackendModels(ents.magmadGateway, ents.cwfGateway, devCasted, statusesByID[hwID])
	}
	return ret
}

func getSubscriberDirectoryHandler(c echo.Context) error {
	networkID, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	configuratorNetwork, err := configurator.LoadNetwork(networkID, false, false)
	if err != nil {
		return obsidian.HttpError(err, http.StatusNotFound)
	}
	if configuratorNetwork.Type != cwf.CwfNetworkType {
		return obsidian.HttpError(fmt.Errorf("NetworkID %s is not a CWF network", networkID), http.StatusBadRequest)
	}
	subscriberID := c.Param("subscriber_id")
	if subscriberID == "" {
		return obsidian.HttpError(fmt.Errorf("SubscriberID cannot be empty"), http.StatusBadRequest)
	}
	directoryState, err := state.GetState(networkID, orc8r.DirectoryRecordType, subscriberID)
	if err == merrors.ErrNotFound {
		return obsidian.HttpError(err, http.StatusNotFound)
	} else if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	cwfRecord, err := convertDirectoryRecordToSubscriberRecord(directoryState.ReportedState)
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, cwfRecord)
}

func convertDirectoryRecordToSubscriberRecord(iRecord interface{}) (*cwfModels.CwfSubscriberDirectoryRecord, error) {
	record, ok := iRecord.(*directoryd.DirectoryRecord)
	if !ok {
		return nil, fmt.Errorf("Could not convert retrieved state to DirectoryRecord")
	}
	b, err := json.Marshal(record.Identifiers)
	if err != nil {
		return nil, err
	}
	c := &cwfModels.CwfSubscriberDirectoryRecord{}
	err = json.Unmarshal(b, c)
	if err != nil {
		return nil, fmt.Errorf("Error converting DirectoryRecord to CWF Record: %s, %v", err, *record)
	}
	c.LocationHistory = record.LocationHistory
	return c, nil
}

func getClusterStatusHandler(c echo.Context) error {
	nid, nerr := obsidian.GetNetworkId(c)
	if nerr != nil {
		return nerr
	}
	network, err := configurator.LoadNetwork(nid, true, true)
	if err == merrors.ErrNotFound {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	if network.Type != cwf.CwfNetworkType {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("network %s is not a <%s> network", nid, cwf.CwfNetworkType))
	}
	reportedClusterStatus, err := state.GetState(nid, cwf.CwfClusterHealthType, "cluster")
	if err == merrors.ErrNotFound {
		return obsidian.HttpError(err, http.StatusNotFound)
	} else if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	clusterStatus, ok := reportedClusterStatus.ReportedState.(*cwfModels.CarrierWifiNetworkClusterStatus)
	if !ok {
		return obsidian.HttpError(fmt.Errorf("could not convert reported retrieved state to ClusterStatus"), http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, clusterStatus)
}

func getHealthStatusHandler(c echo.Context) error {
	nid, gid, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}
	pid, err := configurator.GetPhysicalIDOfEntity(nid, orc8r.MagmadGatewayType, gid)
	if err == merrors.ErrNotFound || len(pid) == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	reportedGatewayState, err := state.GetState(nid, cwf.CwfGatewayHealthType, gid)
	if err == merrors.ErrNotFound {
		return obsidian.HttpError(err, http.StatusNotFound)
	} else if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	healthState, ok := reportedGatewayState.ReportedState.(*cwfModels.CarrierWifiGatewayHealthStatus)
	if !ok {
		return obsidian.HttpError(fmt.Errorf("could not convert reported retrieved state to HealthStatus"), http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, healthState)
}
