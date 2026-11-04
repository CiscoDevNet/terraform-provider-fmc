// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceEIGRP struct {
	Id                                        types.String                   `tfsdk:"id"`
	Domain                                    types.String                   `tfsdk:"domain"`
	DeviceId                                  types.String                   `tfsdk:"device_id"`
	Type                                      types.String                   `tfsdk:"type"`
	AsNumber                                  types.String                   `tfsdk:"as_number"`
	Networks                                  []DeviceEIGRPNetworks          `tfsdk:"networks"`
	AutoSummary                               types.Bool                     `tfsdk:"auto_summary"`
	PassiveInterfacesAll                      types.Bool                     `tfsdk:"passive_interfaces_all"`
	PassiveInterfaces                         []DeviceEIGRPPassiveInterfaces `tfsdk:"passive_interfaces"`
	Neighbors                                 []DeviceEIGRPNeighbors         `tfsdk:"neighbors"`
	FilterRules                               []DeviceEIGRPFilterRules       `tfsdk:"filter_rules"`
	Redistributions                           []DeviceEIGRPRedistributions   `tfsdk:"redistributions"`
	SummaryAddresses                          []DeviceEIGRPSummaryAddresses  `tfsdk:"summary_addresses"`
	Interfaces                                []DeviceEIGRPInterfaces        `tfsdk:"interfaces"`
	RouterId                                  types.String                   `tfsdk:"router_id"`
	DefaultInformationInStandardAccessListId  types.String                   `tfsdk:"default_information_in_standard_access_list_id"`
	DefaultInformationOutStandardAccessListId types.String                   `tfsdk:"default_information_out_standard_access_list_id"`
	AdministrativeDistanceInternal            types.Int64                    `tfsdk:"administrative_distance_internal"`
	AdministrativeDistanceExternal            types.Int64                    `tfsdk:"administrative_distance_external"`
	LogNeighborChanges                        types.Bool                     `tfsdk:"log_neighbor_changes"`
	LogNeighborWarnings                       types.Bool                     `tfsdk:"log_neighbor_warnings"`
	LogNeighborWarningsInterval               types.Int64                    `tfsdk:"log_neighbor_warnings_interval"`
	StubReceiveOnly                           types.Bool                     `tfsdk:"stub_receive_only"`
	StubAdvertiseConnected                    types.Bool                     `tfsdk:"stub_advertise_connected"`
	StubAdvertiseRedistributed                types.Bool                     `tfsdk:"stub_advertise_redistributed"`
	StubAdvertiseStatic                       types.Bool                     `tfsdk:"stub_advertise_static"`
	StubAdvertiseSummary                      types.Bool                     `tfsdk:"stub_advertise_summary"`
	DefaultMetricBandwidth                    types.Int64                    `tfsdk:"default_metric_bandwidth"`
	DefaultMetricDelayTime                    types.Int64                    `tfsdk:"default_metric_delay_time"`
	DefaultMetricReliability                  types.Int64                    `tfsdk:"default_metric_reliability"`
	DefaultMetricLoading                      types.Int64                    `tfsdk:"default_metric_loading"`
	DefaultMetricMtu                          types.Int64                    `tfsdk:"default_metric_mtu"`
	ClusterIpv4PoolId                         types.String                   `tfsdk:"cluster_ipv4_pool_id"`
}

type DeviceEIGRPNetworks struct {
	Id types.String `tfsdk:"id"`
}

type DeviceEIGRPPassiveInterfaces struct {
	Id types.String `tfsdk:"id"`
}

type DeviceEIGRPNeighbors struct {
	InterfaceId types.String `tfsdk:"interface_id"`
	HostId      types.String `tfsdk:"host_id"`
}

type DeviceEIGRPFilterRules struct {
	FilterDirection      types.String `tfsdk:"filter_direction"`
	StandardAccessListId types.String `tfsdk:"standard_access_list_id"`
	InterfaceId          types.String `tfsdk:"interface_id"`
	RoutingProcess       types.String `tfsdk:"routing_process"`
	RoutingProcessId     types.String `tfsdk:"routing_process_id"`
}

type DeviceEIGRPRedistributions struct {
	RedistributeProtocol types.String `tfsdk:"redistribute_protocol"`
	ProcessId            types.String `tfsdk:"process_id"`
	RouteMapId           types.String `tfsdk:"route_map_id"`
	MatchInternal        types.Bool   `tfsdk:"match_internal"`
	MatchExternal1       types.Bool   `tfsdk:"match_external_1"`
	MatchExternal2       types.Bool   `tfsdk:"match_external_2"`
	MatchNssaExternal1   types.Bool   `tfsdk:"match_nssa_external_1"`
	MatchNssaExternal2   types.Bool   `tfsdk:"match_nssa_external_2"`
	MetricBandwidth      types.Int64  `tfsdk:"metric_bandwidth"`
	MetricDelay          types.Int64  `tfsdk:"metric_delay"`
	MetricReliability    types.Int64  `tfsdk:"metric_reliability"`
	MetricLoading        types.Int64  `tfsdk:"metric_loading"`
	MetricMtu            types.Int64  `tfsdk:"metric_mtu"`
}

type DeviceEIGRPSummaryAddresses struct {
	InterfaceId            types.String `tfsdk:"interface_id"`
	NetworkId              types.String `tfsdk:"network_id"`
	AdministrativeDistance types.Int64  `tfsdk:"administrative_distance"`
}

type DeviceEIGRPInterfaces struct {
	InterfaceId           types.String `tfsdk:"interface_id"`
	InterfaceType         types.String `tfsdk:"interface_type"`
	HelloInterval         types.Int64  `tfsdk:"hello_interval"`
	HoldTime              types.Int64  `tfsdk:"hold_time"`
	SplitHorizon          types.Bool   `tfsdk:"split_horizon"`
	DelayTime             types.Int64  `tfsdk:"delay_time"`
	Md5Authentication     types.Bool   `tfsdk:"md5_authentication"`
	AuthenticationKeyType types.String `tfsdk:"authentication_key_type"`
	AuthenticationKeyId   types.Int64  `tfsdk:"authentication_key_id"`
	AuthenticationKey     types.String `tfsdk:"authentication_key"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceEIGRP) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/eigrproutes", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceEIGRP) toBody(ctx context.Context, state DeviceEIGRP) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.AsNumber.IsNull() {
		body, _ = sjson.Set(body, "asNumber", data.AsNumber.ValueString())
	}
	if len(data.Networks) > 0 {
		var networksBody strings.Builder
		networksBody.WriteString("[")
		for _, item := range data.Networks {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if itemBody != "" {
				if networksBody.Len() > 1 {
					networksBody.WriteString(",")
				}
				networksBody.WriteString(itemBody)
			}
		}
		networksBody.WriteString("]")
		body, _ = sjson.SetRaw(body, "basicConfiguration.networks", networksBody.String())
	}
	if !data.AutoSummary.IsNull() {
		body, _ = sjson.Set(body, "basicConfiguration.autoSummary", data.AutoSummary.ValueBool())
	}
	if !data.PassiveInterfacesAll.IsNull() {
		body, _ = sjson.Set(body, "basicConfiguration.passiveInterfaces.allInterfaces", data.PassiveInterfacesAll.ValueBool())
	}
	if len(data.PassiveInterfaces) > 0 {
		var passiveInterfacesBody strings.Builder
		passiveInterfacesBody.WriteString("[")
		for _, item := range data.PassiveInterfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if itemBody != "" {
				if passiveInterfacesBody.Len() > 1 {
					passiveInterfacesBody.WriteString(",")
				}
				passiveInterfacesBody.WriteString(itemBody)
			}
		}
		passiveInterfacesBody.WriteString("]")
		body, _ = sjson.SetRaw(body, "basicConfiguration.passiveInterfaces.selectedInterfaces", passiveInterfacesBody.String())
	}
	if len(data.Neighbors) > 0 {
		var neighborsBody strings.Builder
		neighborsBody.WriteString("[")
		for _, item := range data.Neighbors {
			itemBody := ""
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborInterface.id", item.InterfaceId.ValueString())
			}
			if !item.HostId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAddress.id", item.HostId.ValueString())
			}
			if itemBody != "" {
				if neighborsBody.Len() > 1 {
					neighborsBody.WriteString(",")
				}
				neighborsBody.WriteString(itemBody)
			}
		}
		neighborsBody.WriteString("]")
		body, _ = sjson.SetRaw(body, "neighbors", neighborsBody.String())
	}
	if len(data.FilterRules) > 0 {
		var filterRulesBody strings.Builder
		filterRulesBody.WriteString("[")
		for _, item := range data.FilterRules {
			itemBody := ""
			if !item.FilterDirection.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "filterDirection", item.FilterDirection.ValueString())
			}
			if !item.StandardAccessListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "accessList.id", item.StandardAccessListId.ValueString())
			}
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceInterface.id", item.InterfaceId.ValueString())
			}
			if !item.RoutingProcess.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.RoutingProcess.ValueString())
			}
			if !item.RoutingProcessId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "processId", item.RoutingProcessId.ValueString())
			}
			if itemBody != "" {
				if filterRulesBody.Len() > 1 {
					filterRulesBody.WriteString(",")
				}
				filterRulesBody.WriteString(itemBody)
			}
		}
		filterRulesBody.WriteString("]")
		body, _ = sjson.SetRaw(body, "distributionList", filterRulesBody.String())
	}
	if len(data.Redistributions) > 0 {
		var redistributionsBody strings.Builder
		redistributionsBody.WriteString("[")
		for _, item := range data.Redistributions {
			itemBody := ""
			if !item.RedistributeProtocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.RedistributeProtocol.ValueString())
			}
			if !item.ProcessId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "processId", item.ProcessId.ValueString())
			}
			if !item.RouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMap.id", item.RouteMapId.ValueString())
			}
			if !item.MatchInternal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchInternal", item.MatchInternal.ValueBool())
			}
			if !item.MatchExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchExternal1", item.MatchExternal1.ValueBool())
			}
			if !item.MatchExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchExternal2", item.MatchExternal2.ValueBool())
			}
			if !item.MatchNssaExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchNssaExternal1", item.MatchNssaExternal1.ValueBool())
			}
			if !item.MatchNssaExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchNssaExternal2", item.MatchNssaExternal2.ValueBool())
			}
			if !item.MetricBandwidth.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMetric.bandwidth", item.MetricBandwidth.ValueInt64())
			}
			if !item.MetricDelay.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMetric.delay", item.MetricDelay.ValueInt64())
			}
			if !item.MetricReliability.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMetric.reliability", item.MetricReliability.ValueInt64())
			}
			if !item.MetricLoading.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMetric.effectiveBandwidth", item.MetricLoading.ValueInt64())
			}
			if !item.MetricMtu.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMetric.pathMTU", item.MetricMtu.ValueInt64())
			}
			if itemBody != "" {
				if redistributionsBody.Len() > 1 {
					redistributionsBody.WriteString(",")
				}
				redistributionsBody.WriteString(itemBody)
			}
		}
		redistributionsBody.WriteString("]")
		body, _ = sjson.SetRaw(body, "redistributeProtocols", redistributionsBody.String())
	}
	if len(data.SummaryAddresses) > 0 {
		var summaryAddressesBody strings.Builder
		summaryAddressesBody.WriteString("[")
		for _, item := range data.SummaryAddresses {
			itemBody := ""
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceInterface.id", item.InterfaceId.ValueString())
			}
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "summaryAddress.id", item.NetworkId.ValueString())
			}
			if !item.AdministrativeDistance.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "adminDistance", item.AdministrativeDistance.ValueInt64())
			}
			if itemBody != "" {
				if summaryAddressesBody.Len() > 1 {
					summaryAddressesBody.WriteString(",")
				}
				summaryAddressesBody.WriteString(itemBody)
			}
		}
		summaryAddressesBody.WriteString("]")
		body, _ = sjson.SetRaw(body, "summaryAddressList", summaryAddressesBody.String())
	}
	if len(data.Interfaces) > 0 {
		var interfacesBody strings.Builder
		interfacesBody.WriteString("[")
		for _, item := range data.Interfaces {
			itemBody := ""
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceInterface.id", item.InterfaceId.ValueString())
			}
			if !item.InterfaceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "deviceInterface.type", item.InterfaceType.ValueString())
			}
			if !item.HelloInterval.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eigrpProtocolConfiguration.timers.helloInterval", item.HelloInterval.ValueInt64())
			}
			if !item.HoldTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eigrpProtocolConfiguration.timers.holdTime", item.HoldTime.ValueInt64())
			}
			if !item.SplitHorizon.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eigrpProtocolConfiguration.splitHorizon", item.SplitHorizon.ValueBool())
			}
			if !item.DelayTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eigrpProtocolConfiguration.timers.delay", item.DelayTime.ValueInt64())
			}
			if !item.Md5Authentication.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eigrpProtocolConfiguration.authentication.md5auth", item.Md5Authentication.ValueBool())
			}
			if !item.AuthenticationKeyType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eigrpProtocolConfiguration.authentication.pwdEncryption", item.AuthenticationKeyType.ValueString())
			}
			if !item.AuthenticationKeyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eigrpProtocolConfiguration.authentication.keyId", item.AuthenticationKeyId.ValueInt64())
			}
			if !item.AuthenticationKey.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "eigrpProtocolConfiguration.authentication.password", item.AuthenticationKey.ValueString())
			}
			if itemBody != "" {
				if interfacesBody.Len() > 1 {
					interfacesBody.WriteString(",")
				}
				interfacesBody.WriteString(itemBody)
			}
		}
		interfacesBody.WriteString("]")
		body, _ = sjson.SetRaw(body, "eigrpInterfaces", interfacesBody.String())
	}
	if !data.RouterId.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.routerId", data.RouterId.ValueString())
	}
	if !data.DefaultInformationInStandardAccessListId.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.defaultRouteDistributionIn.id", data.DefaultInformationInStandardAccessListId.ValueString())
	}
	if !data.DefaultInformationOutStandardAccessListId.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.defaultRouteDistributionOut.id", data.DefaultInformationOutStandardAccessListId.ValueString())
	}
	if !data.AdministrativeDistanceInternal.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.administrativeDistance.internal", data.AdministrativeDistanceInternal.ValueInt64())
	}
	if !data.AdministrativeDistanceExternal.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.administrativeDistance.external", data.AdministrativeDistanceExternal.ValueInt64())
	}
	if !data.LogNeighborChanges.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.logNeighborChanges", data.LogNeighborChanges.ValueBool())
	}
	if !data.LogNeighborWarnings.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.enableLogNeighborWarnings", data.LogNeighborWarnings.ValueBool())
	}
	if !data.LogNeighborWarningsInterval.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.logNeighborWarningInterval", data.LogNeighborWarningsInterval.ValueInt64())
	}
	if !data.StubReceiveOnly.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.stub.receiveOnly", data.StubReceiveOnly.ValueBool())
	}
	if !data.StubAdvertiseConnected.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.stub.advertise.connected", data.StubAdvertiseConnected.ValueBool())
	}
	if !data.StubAdvertiseRedistributed.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.stub.advertise.redistributed", data.StubAdvertiseRedistributed.ValueBool())
	}
	if !data.StubAdvertiseStatic.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.stub.advertise.static", data.StubAdvertiseStatic.ValueBool())
	}
	if !data.StubAdvertiseSummary.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.stub.advertise.summary", data.StubAdvertiseSummary.ValueBool())
	}
	if !data.DefaultMetricBandwidth.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.routeMetric.bandwidth", data.DefaultMetricBandwidth.ValueInt64())
	}
	if !data.DefaultMetricDelayTime.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.routeMetric.delay", data.DefaultMetricDelayTime.ValueInt64())
	}
	if !data.DefaultMetricReliability.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.routeMetric.reliability", data.DefaultMetricReliability.ValueInt64())
	}
	if !data.DefaultMetricLoading.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.routeMetric.effectiveBandwidth", data.DefaultMetricLoading.ValueInt64())
	}
	if !data.DefaultMetricMtu.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.routeMetric.pathMTU", data.DefaultMetricMtu.ValueInt64())
	}
	if !data.ClusterIpv4PoolId.IsNull() {
		body, _ = sjson.Set(body, "advancedConfiguration.clusterPool.id", data.ClusterIpv4PoolId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceEIGRP) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("asNumber"); value.Exists() {
		data.AsNumber = types.StringValue(value.String())
	} else {
		data.AsNumber = types.StringNull()
	}
	if value := res.Get("basicConfiguration.networks"); value.Exists() {
		data.Networks = make([]DeviceEIGRPNetworks, 0, int(value.Get("#").Int()))
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceEIGRPNetworks{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Networks = append((*parent).Networks, data)
			return true
		})
	}
	if value := res.Get("basicConfiguration.autoSummary"); value.Exists() {
		data.AutoSummary = types.BoolValue(value.Bool())
	} else {
		data.AutoSummary = types.BoolValue(false)
	}
	if value := res.Get("basicConfiguration.passiveInterfaces.allInterfaces"); value.Exists() {
		data.PassiveInterfacesAll = types.BoolValue(value.Bool())
	} else {
		data.PassiveInterfacesAll = types.BoolNull()
	}
	if value := res.Get("basicConfiguration.passiveInterfaces.selectedInterfaces"); value.Exists() {
		data.PassiveInterfaces = make([]DeviceEIGRPPassiveInterfaces, 0, int(value.Get("#").Int()))
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceEIGRPPassiveInterfaces{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).PassiveInterfaces = append((*parent).PassiveInterfaces, data)
			return true
		})
	}
	if value := res.Get("neighbors"); value.Exists() {
		data.Neighbors = make([]DeviceEIGRPNeighbors, 0, int(value.Get("#").Int()))
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceEIGRPNeighbors{}
			if value := res.Get("neighborInterface.id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			if value := res.Get("neighborAddress.id"); value.Exists() {
				data.HostId = types.StringValue(value.String())
			} else {
				data.HostId = types.StringNull()
			}
			(*parent).Neighbors = append((*parent).Neighbors, data)
			return true
		})
	}
	if value := res.Get("distributionList"); value.Exists() {
		data.FilterRules = make([]DeviceEIGRPFilterRules, 0, int(value.Get("#").Int()))
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceEIGRPFilterRules{}
			if value := res.Get("filterDirection"); value.Exists() {
				data.FilterDirection = types.StringValue(value.String())
			} else {
				data.FilterDirection = types.StringNull()
			}
			if value := res.Get("accessList.id"); value.Exists() {
				data.StandardAccessListId = types.StringValue(value.String())
			} else {
				data.StandardAccessListId = types.StringNull()
			}
			if value := res.Get("deviceInterface.id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() {
				data.RoutingProcess = types.StringValue(value.String())
			} else {
				data.RoutingProcess = types.StringNull()
			}
			if value := res.Get("processId"); value.Exists() {
				data.RoutingProcessId = types.StringValue(value.String())
			} else {
				data.RoutingProcessId = types.StringNull()
			}
			(*parent).FilterRules = append((*parent).FilterRules, data)
			return true
		})
	}
	if value := res.Get("redistributeProtocols"); value.Exists() {
		data.Redistributions = make([]DeviceEIGRPRedistributions, 0, int(value.Get("#").Int()))
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceEIGRPRedistributions{}
			if value := res.Get("protocol"); value.Exists() {
				data.RedistributeProtocol = types.StringValue(value.String())
			} else {
				data.RedistributeProtocol = types.StringNull()
			}
			if value := res.Get("processId"); value.Exists() {
				data.ProcessId = types.StringValue(value.String())
			} else {
				data.ProcessId = types.StringNull()
			}
			if value := res.Get("routeMap.id"); value.Exists() {
				data.RouteMapId = types.StringValue(value.String())
			} else {
				data.RouteMapId = types.StringNull()
			}
			if value := res.Get("matchInternal"); value.Exists() {
				data.MatchInternal = types.BoolValue(value.Bool())
			} else {
				data.MatchInternal = types.BoolNull()
			}
			if value := res.Get("matchExternal1"); value.Exists() {
				data.MatchExternal1 = types.BoolValue(value.Bool())
			} else {
				data.MatchExternal1 = types.BoolNull()
			}
			if value := res.Get("matchExternal2"); value.Exists() {
				data.MatchExternal2 = types.BoolValue(value.Bool())
			} else {
				data.MatchExternal2 = types.BoolNull()
			}
			if value := res.Get("matchNssaExternal1"); value.Exists() {
				data.MatchNssaExternal1 = types.BoolValue(value.Bool())
			} else {
				data.MatchNssaExternal1 = types.BoolNull()
			}
			if value := res.Get("matchNssaExternal2"); value.Exists() {
				data.MatchNssaExternal2 = types.BoolValue(value.Bool())
			} else {
				data.MatchNssaExternal2 = types.BoolNull()
			}
			if value := res.Get("routeMetric.bandwidth"); value.Exists() {
				data.MetricBandwidth = types.Int64Value(value.Int())
			} else {
				data.MetricBandwidth = types.Int64Null()
			}
			if value := res.Get("routeMetric.delay"); value.Exists() {
				data.MetricDelay = types.Int64Value(value.Int())
			} else {
				data.MetricDelay = types.Int64Null()
			}
			if value := res.Get("routeMetric.reliability"); value.Exists() {
				data.MetricReliability = types.Int64Value(value.Int())
			} else {
				data.MetricReliability = types.Int64Null()
			}
			if value := res.Get("routeMetric.effectiveBandwidth"); value.Exists() {
				data.MetricLoading = types.Int64Value(value.Int())
			} else {
				data.MetricLoading = types.Int64Null()
			}
			if value := res.Get("routeMetric.pathMTU"); value.Exists() {
				data.MetricMtu = types.Int64Value(value.Int())
			} else {
				data.MetricMtu = types.Int64Null()
			}
			(*parent).Redistributions = append((*parent).Redistributions, data)
			return true
		})
	}
	if value := res.Get("summaryAddressList"); value.Exists() {
		data.SummaryAddresses = make([]DeviceEIGRPSummaryAddresses, 0, int(value.Get("#").Int()))
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceEIGRPSummaryAddresses{}
			if value := res.Get("deviceInterface.id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			if value := res.Get("summaryAddress.id"); value.Exists() {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("adminDistance"); value.Exists() {
				data.AdministrativeDistance = types.Int64Value(value.Int())
			} else {
				data.AdministrativeDistance = types.Int64Null()
			}
			(*parent).SummaryAddresses = append((*parent).SummaryAddresses, data)
			return true
		})
	}
	if value := res.Get("eigrpInterfaces"); value.Exists() {
		data.Interfaces = make([]DeviceEIGRPInterfaces, 0, int(value.Get("#").Int()))
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceEIGRPInterfaces{}
			if value := res.Get("deviceInterface.id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			if value := res.Get("deviceInterface.type"); value.Exists() {
				data.InterfaceType = types.StringValue(value.String())
			} else {
				data.InterfaceType = types.StringNull()
			}
			if value := res.Get("eigrpProtocolConfiguration.timers.helloInterval"); value.Exists() {
				data.HelloInterval = types.Int64Value(value.Int())
			} else {
				data.HelloInterval = types.Int64Value(5)
			}
			if value := res.Get("eigrpProtocolConfiguration.timers.holdTime"); value.Exists() {
				data.HoldTime = types.Int64Value(value.Int())
			} else {
				data.HoldTime = types.Int64Value(15)
			}
			if value := res.Get("eigrpProtocolConfiguration.splitHorizon"); value.Exists() {
				data.SplitHorizon = types.BoolValue(value.Bool())
			} else {
				data.SplitHorizon = types.BoolNull()
			}
			if value := res.Get("eigrpProtocolConfiguration.timers.delay"); value.Exists() {
				data.DelayTime = types.Int64Value(value.Int())
			} else {
				data.DelayTime = types.Int64Null()
			}
			if value := res.Get("eigrpProtocolConfiguration.authentication.md5auth"); value.Exists() {
				data.Md5Authentication = types.BoolValue(value.Bool())
			} else {
				data.Md5Authentication = types.BoolNull()
			}
			if value := res.Get("eigrpProtocolConfiguration.authentication.pwdEncryption"); value.Exists() {
				data.AuthenticationKeyType = types.StringValue(value.String())
			} else {
				data.AuthenticationKeyType = types.StringNull()
			}
			if value := res.Get("eigrpProtocolConfiguration.authentication.keyId"); value.Exists() {
				data.AuthenticationKeyId = types.Int64Value(value.Int())
			} else {
				data.AuthenticationKeyId = types.Int64Null()
			}
			(*parent).Interfaces = append((*parent).Interfaces, data)
			return true
		})
	}
	if value := res.Get("advancedConfiguration.routerId"); value.Exists() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	if value := res.Get("advancedConfiguration.defaultRouteDistributionIn.id"); value.Exists() {
		data.DefaultInformationInStandardAccessListId = types.StringValue(value.String())
	} else {
		data.DefaultInformationInStandardAccessListId = types.StringNull()
	}
	if value := res.Get("advancedConfiguration.defaultRouteDistributionOut.id"); value.Exists() {
		data.DefaultInformationOutStandardAccessListId = types.StringValue(value.String())
	} else {
		data.DefaultInformationOutStandardAccessListId = types.StringNull()
	}
	if value := res.Get("advancedConfiguration.administrativeDistance.internal"); value.Exists() {
		data.AdministrativeDistanceInternal = types.Int64Value(value.Int())
	} else {
		data.AdministrativeDistanceInternal = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.administrativeDistance.external"); value.Exists() {
		data.AdministrativeDistanceExternal = types.Int64Value(value.Int())
	} else {
		data.AdministrativeDistanceExternal = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.logNeighborChanges"); value.Exists() {
		data.LogNeighborChanges = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborChanges = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.enableLogNeighborWarnings"); value.Exists() {
		data.LogNeighborWarnings = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborWarnings = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.logNeighborWarningInterval"); value.Exists() {
		data.LogNeighborWarningsInterval = types.Int64Value(value.Int())
	} else {
		data.LogNeighborWarningsInterval = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.stub.receiveOnly"); value.Exists() {
		data.StubReceiveOnly = types.BoolValue(value.Bool())
	} else {
		data.StubReceiveOnly = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.stub.advertise.connected"); value.Exists() {
		data.StubAdvertiseConnected = types.BoolValue(value.Bool())
	} else {
		data.StubAdvertiseConnected = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.stub.advertise.redistributed"); value.Exists() {
		data.StubAdvertiseRedistributed = types.BoolValue(value.Bool())
	} else {
		data.StubAdvertiseRedistributed = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.stub.advertise.static"); value.Exists() {
		data.StubAdvertiseStatic = types.BoolValue(value.Bool())
	} else {
		data.StubAdvertiseStatic = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.stub.advertise.summary"); value.Exists() {
		data.StubAdvertiseSummary = types.BoolValue(value.Bool())
	} else {
		data.StubAdvertiseSummary = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.routeMetric.bandwidth"); value.Exists() {
		data.DefaultMetricBandwidth = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricBandwidth = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.routeMetric.delay"); value.Exists() {
		data.DefaultMetricDelayTime = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricDelayTime = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.routeMetric.reliability"); value.Exists() {
		data.DefaultMetricReliability = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricReliability = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.routeMetric.effectiveBandwidth"); value.Exists() {
		data.DefaultMetricLoading = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricLoading = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.routeMetric.pathMTU"); value.Exists() {
		data.DefaultMetricMtu = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricMtu = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.clusterPool.id"); value.Exists() {
		data.ClusterIpv4PoolId = types.StringValue(value.String())
	} else {
		data.ClusterIpv4PoolId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceEIGRP) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("asNumber"); value.Exists() && !data.AsNumber.IsNull() {
		data.AsNumber = types.StringValue(value.String())
	} else {
		data.AsNumber = types.StringNull()
	}
	networksArray := res.Get("basicConfiguration.networks")
	for i := 0; i < len(data.Networks); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Networks[i].Id.ValueString()}

		parent := &data
		data := (*parent).Networks[i]
		var res gjson.Result

		networksArray.ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Networks[%d] = %+v",
				i,
				(*parent).Networks[i],
			))
			(*parent).Networks = slices.Delete((*parent).Networks, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).Networks[i] = data
	}
	if value := res.Get("basicConfiguration.autoSummary"); value.Exists() && !data.AutoSummary.IsNull() {
		data.AutoSummary = types.BoolValue(value.Bool())
	} else if data.AutoSummary.ValueBool() != false {
		data.AutoSummary = types.BoolNull()
	}
	if value := res.Get("basicConfiguration.passiveInterfaces.allInterfaces"); value.Exists() && !data.PassiveInterfacesAll.IsNull() {
		data.PassiveInterfacesAll = types.BoolValue(value.Bool())
	} else {
		data.PassiveInterfacesAll = types.BoolNull()
	}
	passiveInterfacesArray := res.Get("basicConfiguration.passiveInterfaces.selectedInterfaces")
	for i := 0; i < len(data.PassiveInterfaces); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.PassiveInterfaces[i].Id.ValueString()}

		parent := &data
		data := (*parent).PassiveInterfaces[i]
		var res gjson.Result

		passiveInterfacesArray.ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing PassiveInterfaces[%d] = %+v",
				i,
				(*parent).PassiveInterfaces[i],
			))
			(*parent).PassiveInterfaces = slices.Delete((*parent).PassiveInterfaces, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).PassiveInterfaces[i] = data
	}
	neighborsArray := res.Get("neighbors")
	for i := 0; i < len(data.Neighbors); i++ {
		keys := [...]string{"neighborAddress.id"}
		keyValues := [...]string{data.Neighbors[i].HostId.ValueString()}

		parent := &data
		data := (*parent).Neighbors[i]
		var res gjson.Result

		neighborsArray.ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Neighbors[%d] = %+v",
				i,
				(*parent).Neighbors[i],
			))
			(*parent).Neighbors = slices.Delete((*parent).Neighbors, i, i+1)
			i--

			continue
		}
		if value := res.Get("neighborInterface.id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("neighborAddress.id"); value.Exists() && !data.HostId.IsNull() {
			data.HostId = types.StringValue(value.String())
		} else {
			data.HostId = types.StringNull()
		}
		(*parent).Neighbors[i] = data
	}
	filterRulesArray := res.Get("distributionList")
	for i := 0; i < len(data.FilterRules); i++ {
		keys := [...]string{"filterDirection", "accessList.id", "deviceInterface.id", "protocol", "processId"}
		keyValues := [...]string{data.FilterRules[i].FilterDirection.ValueString(), data.FilterRules[i].StandardAccessListId.ValueString(), data.FilterRules[i].InterfaceId.ValueString(), data.FilterRules[i].RoutingProcess.ValueString(), data.FilterRules[i].RoutingProcessId.ValueString()}

		parent := &data
		data := (*parent).FilterRules[i]
		var res gjson.Result

		filterRulesArray.ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing FilterRules[%d] = %+v",
				i,
				(*parent).FilterRules[i],
			))
			(*parent).FilterRules = slices.Delete((*parent).FilterRules, i, i+1)
			i--

			continue
		}
		if value := res.Get("filterDirection"); value.Exists() && !data.FilterDirection.IsNull() {
			data.FilterDirection = types.StringValue(value.String())
		} else {
			data.FilterDirection = types.StringNull()
		}
		if value := res.Get("accessList.id"); value.Exists() && !data.StandardAccessListId.IsNull() {
			data.StandardAccessListId = types.StringValue(value.String())
		} else {
			data.StandardAccessListId = types.StringNull()
		}
		if value := res.Get("deviceInterface.id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.RoutingProcess.IsNull() {
			data.RoutingProcess = types.StringValue(value.String())
		} else {
			data.RoutingProcess = types.StringNull()
		}
		if value := res.Get("processId"); value.Exists() && !data.RoutingProcessId.IsNull() {
			data.RoutingProcessId = types.StringValue(value.String())
		} else {
			data.RoutingProcessId = types.StringNull()
		}
		(*parent).FilterRules[i] = data
	}
	redistributionsArray := res.Get("redistributeProtocols")
	for i := 0; i < len(data.Redistributions); i++ {
		keys := [...]string{"protocol", "processId"}
		keyValues := [...]string{data.Redistributions[i].RedistributeProtocol.ValueString(), data.Redistributions[i].ProcessId.ValueString()}

		parent := &data
		data := (*parent).Redistributions[i]
		var res gjson.Result

		redistributionsArray.ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Redistributions[%d] = %+v",
				i,
				(*parent).Redistributions[i],
			))
			(*parent).Redistributions = slices.Delete((*parent).Redistributions, i, i+1)
			i--

			continue
		}
		if value := res.Get("protocol"); value.Exists() && !data.RedistributeProtocol.IsNull() {
			data.RedistributeProtocol = types.StringValue(value.String())
		} else {
			data.RedistributeProtocol = types.StringNull()
		}
		if value := res.Get("processId"); value.Exists() && !data.ProcessId.IsNull() {
			data.ProcessId = types.StringValue(value.String())
		} else {
			data.ProcessId = types.StringNull()
		}
		if value := res.Get("routeMap.id"); value.Exists() && !data.RouteMapId.IsNull() {
			data.RouteMapId = types.StringValue(value.String())
		} else {
			data.RouteMapId = types.StringNull()
		}
		if value := res.Get("matchInternal"); value.Exists() && !data.MatchInternal.IsNull() {
			data.MatchInternal = types.BoolValue(value.Bool())
		} else {
			data.MatchInternal = types.BoolNull()
		}
		if value := res.Get("matchExternal1"); value.Exists() && !data.MatchExternal1.IsNull() {
			data.MatchExternal1 = types.BoolValue(value.Bool())
		} else {
			data.MatchExternal1 = types.BoolNull()
		}
		if value := res.Get("matchExternal2"); value.Exists() && !data.MatchExternal2.IsNull() {
			data.MatchExternal2 = types.BoolValue(value.Bool())
		} else {
			data.MatchExternal2 = types.BoolNull()
		}
		if value := res.Get("matchNssaExternal1"); value.Exists() && !data.MatchNssaExternal1.IsNull() {
			data.MatchNssaExternal1 = types.BoolValue(value.Bool())
		} else {
			data.MatchNssaExternal1 = types.BoolNull()
		}
		if value := res.Get("matchNssaExternal2"); value.Exists() && !data.MatchNssaExternal2.IsNull() {
			data.MatchNssaExternal2 = types.BoolValue(value.Bool())
		} else {
			data.MatchNssaExternal2 = types.BoolNull()
		}
		if value := res.Get("routeMetric.bandwidth"); value.Exists() && !data.MetricBandwidth.IsNull() {
			data.MetricBandwidth = types.Int64Value(value.Int())
		} else {
			data.MetricBandwidth = types.Int64Null()
		}
		if value := res.Get("routeMetric.delay"); value.Exists() && !data.MetricDelay.IsNull() {
			data.MetricDelay = types.Int64Value(value.Int())
		} else {
			data.MetricDelay = types.Int64Null()
		}
		if value := res.Get("routeMetric.reliability"); value.Exists() && !data.MetricReliability.IsNull() {
			data.MetricReliability = types.Int64Value(value.Int())
		} else {
			data.MetricReliability = types.Int64Null()
		}
		if value := res.Get("routeMetric.effectiveBandwidth"); value.Exists() && !data.MetricLoading.IsNull() {
			data.MetricLoading = types.Int64Value(value.Int())
		} else {
			data.MetricLoading = types.Int64Null()
		}
		if value := res.Get("routeMetric.pathMTU"); value.Exists() && !data.MetricMtu.IsNull() {
			data.MetricMtu = types.Int64Value(value.Int())
		} else {
			data.MetricMtu = types.Int64Null()
		}
		(*parent).Redistributions[i] = data
	}
	summaryAddressesArray := res.Get("summaryAddressList")
	for i := 0; i < len(data.SummaryAddresses); i++ {
		keys := [...]string{"deviceInterface.id", "summaryAddress.id"}
		keyValues := [...]string{data.SummaryAddresses[i].InterfaceId.ValueString(), data.SummaryAddresses[i].NetworkId.ValueString()}

		parent := &data
		data := (*parent).SummaryAddresses[i]
		var res gjson.Result

		summaryAddressesArray.ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing SummaryAddresses[%d] = %+v",
				i,
				(*parent).SummaryAddresses[i],
			))
			(*parent).SummaryAddresses = slices.Delete((*parent).SummaryAddresses, i, i+1)
			i--

			continue
		}
		if value := res.Get("deviceInterface.id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("summaryAddress.id"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("adminDistance"); value.Exists() && !data.AdministrativeDistance.IsNull() {
			data.AdministrativeDistance = types.Int64Value(value.Int())
		} else {
			data.AdministrativeDistance = types.Int64Null()
		}
		(*parent).SummaryAddresses[i] = data
	}
	interfacesArray := res.Get("eigrpInterfaces")
	for i := 0; i < len(data.Interfaces); i++ {
		keys := [...]string{"deviceInterface.id"}
		keyValues := [...]string{data.Interfaces[i].InterfaceId.ValueString()}

		parent := &data
		data := (*parent).Interfaces[i]
		var res gjson.Result

		interfacesArray.ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Interfaces[%d] = %+v",
				i,
				(*parent).Interfaces[i],
			))
			(*parent).Interfaces = slices.Delete((*parent).Interfaces, i, i+1)
			i--

			continue
		}
		if value := res.Get("deviceInterface.id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("deviceInterface.type"); value.Exists() && !data.InterfaceType.IsNull() {
			data.InterfaceType = types.StringValue(value.String())
		} else {
			data.InterfaceType = types.StringNull()
		}
		if value := res.Get("eigrpProtocolConfiguration.timers.helloInterval"); value.Exists() && !data.HelloInterval.IsNull() {
			data.HelloInterval = types.Int64Value(value.Int())
		} else if data.HelloInterval.ValueInt64() != 5 {
			data.HelloInterval = types.Int64Null()
		}
		if value := res.Get("eigrpProtocolConfiguration.timers.holdTime"); value.Exists() && !data.HoldTime.IsNull() {
			data.HoldTime = types.Int64Value(value.Int())
		} else if data.HoldTime.ValueInt64() != 15 {
			data.HoldTime = types.Int64Null()
		}
		if value := res.Get("eigrpProtocolConfiguration.splitHorizon"); value.Exists() && !data.SplitHorizon.IsNull() {
			data.SplitHorizon = types.BoolValue(value.Bool())
		} else {
			data.SplitHorizon = types.BoolNull()
		}
		if value := res.Get("eigrpProtocolConfiguration.timers.delay"); value.Exists() && !data.DelayTime.IsNull() {
			data.DelayTime = types.Int64Value(value.Int())
		} else {
			data.DelayTime = types.Int64Null()
		}
		if value := res.Get("eigrpProtocolConfiguration.authentication.md5auth"); value.Exists() && !data.Md5Authentication.IsNull() {
			data.Md5Authentication = types.BoolValue(value.Bool())
		} else {
			data.Md5Authentication = types.BoolNull()
		}
		if value := res.Get("eigrpProtocolConfiguration.authentication.pwdEncryption"); value.Exists() && !data.AuthenticationKeyType.IsNull() {
			data.AuthenticationKeyType = types.StringValue(value.String())
		} else {
			data.AuthenticationKeyType = types.StringNull()
		}
		if value := res.Get("eigrpProtocolConfiguration.authentication.keyId"); value.Exists() && !data.AuthenticationKeyId.IsNull() {
			data.AuthenticationKeyId = types.Int64Value(value.Int())
		} else {
			data.AuthenticationKeyId = types.Int64Null()
		}
		(*parent).Interfaces[i] = data
	}
	if value := res.Get("advancedConfiguration.routerId"); value.Exists() && !data.RouterId.IsNull() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	if value := res.Get("advancedConfiguration.defaultRouteDistributionIn.id"); value.Exists() && !data.DefaultInformationInStandardAccessListId.IsNull() {
		data.DefaultInformationInStandardAccessListId = types.StringValue(value.String())
	} else {
		data.DefaultInformationInStandardAccessListId = types.StringNull()
	}
	if value := res.Get("advancedConfiguration.defaultRouteDistributionOut.id"); value.Exists() && !data.DefaultInformationOutStandardAccessListId.IsNull() {
		data.DefaultInformationOutStandardAccessListId = types.StringValue(value.String())
	} else {
		data.DefaultInformationOutStandardAccessListId = types.StringNull()
	}
	if value := res.Get("advancedConfiguration.administrativeDistance.internal"); value.Exists() && !data.AdministrativeDistanceInternal.IsNull() {
		data.AdministrativeDistanceInternal = types.Int64Value(value.Int())
	} else {
		data.AdministrativeDistanceInternal = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.administrativeDistance.external"); value.Exists() && !data.AdministrativeDistanceExternal.IsNull() {
		data.AdministrativeDistanceExternal = types.Int64Value(value.Int())
	} else {
		data.AdministrativeDistanceExternal = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.logNeighborChanges"); value.Exists() && !data.LogNeighborChanges.IsNull() {
		data.LogNeighborChanges = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborChanges = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.enableLogNeighborWarnings"); value.Exists() && !data.LogNeighborWarnings.IsNull() {
		data.LogNeighborWarnings = types.BoolValue(value.Bool())
	} else {
		data.LogNeighborWarnings = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.logNeighborWarningInterval"); value.Exists() && !data.LogNeighborWarningsInterval.IsNull() {
		data.LogNeighborWarningsInterval = types.Int64Value(value.Int())
	} else {
		data.LogNeighborWarningsInterval = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.stub.receiveOnly"); value.Exists() && !data.StubReceiveOnly.IsNull() {
		data.StubReceiveOnly = types.BoolValue(value.Bool())
	} else {
		data.StubReceiveOnly = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.stub.advertise.connected"); value.Exists() && !data.StubAdvertiseConnected.IsNull() {
		data.StubAdvertiseConnected = types.BoolValue(value.Bool())
	} else {
		data.StubAdvertiseConnected = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.stub.advertise.redistributed"); value.Exists() && !data.StubAdvertiseRedistributed.IsNull() {
		data.StubAdvertiseRedistributed = types.BoolValue(value.Bool())
	} else {
		data.StubAdvertiseRedistributed = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.stub.advertise.static"); value.Exists() && !data.StubAdvertiseStatic.IsNull() {
		data.StubAdvertiseStatic = types.BoolValue(value.Bool())
	} else {
		data.StubAdvertiseStatic = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.stub.advertise.summary"); value.Exists() && !data.StubAdvertiseSummary.IsNull() {
		data.StubAdvertiseSummary = types.BoolValue(value.Bool())
	} else {
		data.StubAdvertiseSummary = types.BoolNull()
	}
	if value := res.Get("advancedConfiguration.routeMetric.bandwidth"); value.Exists() && !data.DefaultMetricBandwidth.IsNull() {
		data.DefaultMetricBandwidth = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricBandwidth = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.routeMetric.delay"); value.Exists() && !data.DefaultMetricDelayTime.IsNull() {
		data.DefaultMetricDelayTime = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricDelayTime = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.routeMetric.reliability"); value.Exists() && !data.DefaultMetricReliability.IsNull() {
		data.DefaultMetricReliability = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricReliability = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.routeMetric.effectiveBandwidth"); value.Exists() && !data.DefaultMetricLoading.IsNull() {
		data.DefaultMetricLoading = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricLoading = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.routeMetric.pathMTU"); value.Exists() && !data.DefaultMetricMtu.IsNull() {
		data.DefaultMetricMtu = types.Int64Value(value.Int())
	} else {
		data.DefaultMetricMtu = types.Int64Null()
	}
	if value := res.Get("advancedConfiguration.clusterPool.id"); value.Exists() && !data.ClusterIpv4PoolId.IsNull() {
		data.ClusterIpv4PoolId = types.StringValue(value.String())
	} else {
		data.ClusterIpv4PoolId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceEIGRP) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk

// Section below is generated&owned by "gen/generator.go". //template:begin findObjectsToBeReplaced

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

// End of section. //template:end clearItemIds

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides
