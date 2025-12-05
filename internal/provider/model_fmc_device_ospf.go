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
	"strconv"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceOSPF struct {
	Id                               types.String                 `tfsdk:"id"`
	Domain                           types.String                 `tfsdk:"domain"`
	VrfId                            types.String                 `tfsdk:"vrf_id"`
	DeviceId                         types.String                 `tfsdk:"device_id"`
	Type                             types.String                 `tfsdk:"type"`
	ProcessId                        types.Int64                  `tfsdk:"process_id"`
	RouterId                         types.String                 `tfsdk:"router_id"`
	Rfc1583Compatible                types.Bool                   `tfsdk:"rfc_1583_compatible"`
	LogAdjacencyChanges              types.String                 `tfsdk:"log_adjacency_changes"`
	IgnoreLsaMospf                   types.Bool                   `tfsdk:"ignore_lsa_mospf"`
	AdministrativeDistanceInterArea  types.Int64                  `tfsdk:"administrative_distance_inter_area"`
	AdministrativeDistanceIntraArea  types.Int64                  `tfsdk:"administrative_distance_intra_area"`
	AdministrativeDistanceExternal   types.Int64                  `tfsdk:"administrative_distance_external"`
	TimerLsaGroup                    types.Int64                  `tfsdk:"timer_lsa_group"`
	DefaultRouteAlwaysAdvertise      types.Bool                   `tfsdk:"default_route_always_advertise"`
	DefaultRouteMetric               types.Int64                  `tfsdk:"default_route_metric"`
	DefaultRouteMetricType           types.String                 `tfsdk:"default_route_metric_type"`
	DefaultRouteRouteMapId           types.String                 `tfsdk:"default_route_route_map_id"`
	NonStopForwarding                types.Bool                   `tfsdk:"non_stop_forwarding"`
	NonStopForwardingMechanism       types.String                 `tfsdk:"non_stop_forwarding_mechanism"`
	NonStopForwardingHelperMode      types.Bool                   `tfsdk:"non_stop_forwarding_helper_mode"`
	NonStopForwardingRestartInterval types.Int64                  `tfsdk:"non_stop_forwarding_restart_interval"`
	NonStopForwardingCapability      types.Bool                   `tfsdk:"non_stop_forwarding_capability"`
	NonStopForwardingStrictMode      types.Bool                   `tfsdk:"non_stop_forwarding_strict_mode"`
	Areas                            []DeviceOSPFAreas            `tfsdk:"areas"`
	Redistributions                  []DeviceOSPFRedistributions  `tfsdk:"redistributions"`
	FilterRules                      []DeviceOSPFFilterRules      `tfsdk:"filter_rules"`
	SummaryAddresses                 []DeviceOSPFSummaryAddresses `tfsdk:"summary_addresses"`
	Neighbors                        []DeviceOSPFNeighbors        `tfsdk:"neighbors"`
}

type DeviceOSPFAreas struct {
	Id                     types.String                      `tfsdk:"id"`
	Type                   types.String                      `tfsdk:"type"`
	NoSummary              types.Bool                        `tfsdk:"no_summary"`
	NoRedistribution       types.Bool                        `tfsdk:"no_redistribution"`
	DefaultRouteMetricType types.String                      `tfsdk:"default_route_metric_type"`
	DefaultRouteMetric     types.Int64                       `tfsdk:"default_route_metric"`
	Networks               []DeviceOSPFAreasNetworks         `tfsdk:"networks"`
	Authentication         types.String                      `tfsdk:"authentication"`
	DefaultCost            types.Int64                       `tfsdk:"default_cost"`
	Ranges                 []DeviceOSPFAreasRanges           `tfsdk:"ranges"`
	VirtualLinks           []DeviceOSPFAreasVirtualLinks     `tfsdk:"virtual_links"`
	InterAreaFilters       []DeviceOSPFAreasInterAreaFilters `tfsdk:"inter_area_filters"`
}

type DeviceOSPFRedistributions struct {
	RedistributeProtocol types.String `tfsdk:"redistribute_protocol"`
	AsNumber             types.Int64  `tfsdk:"as_number"`
	MatchExternal1       types.Bool   `tfsdk:"match_external_1"`
	MatchExternal2       types.Bool   `tfsdk:"match_external_2"`
	MatchInternal        types.Bool   `tfsdk:"match_internal"`
	MatchNssaExternal1   types.Bool   `tfsdk:"match_nssa_external_1"`
	MatchNssaExternal2   types.Bool   `tfsdk:"match_nssa_external_2"`
	ProcessId            types.Int64  `tfsdk:"process_id"`
	Subnets              types.Bool   `tfsdk:"subnets"`
	Metric               types.Int64  `tfsdk:"metric"`
	MetricType           types.String `tfsdk:"metric_type"`
	Tag                  types.Int64  `tfsdk:"tag"`
	RouteMapId           types.String `tfsdk:"route_map_id"`
}

type DeviceOSPFFilterRules struct {
	AccessListId     types.String `tfsdk:"access_list_id"`
	TrafficDirection types.String `tfsdk:"traffic_direction"`
	RoutingProcess   types.String `tfsdk:"routing_process"`
	RoutingProcessId types.Int64  `tfsdk:"routing_process_id"`
	InterfaceId      types.String `tfsdk:"interface_id"`
}

type DeviceOSPFSummaryAddresses struct {
	Networks  []DeviceOSPFSummaryAddressesNetworks `tfsdk:"networks"`
	Tag       types.Int64                          `tfsdk:"tag"`
	Advertise types.Bool                           `tfsdk:"advertise"`
}

type DeviceOSPFNeighbors struct {
	InterfaceId    types.String `tfsdk:"interface_id"`
	NeighborHostId types.String `tfsdk:"neighbor_host_id"`
}

type DeviceOSPFAreasNetworks struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}
type DeviceOSPFAreasRanges struct {
	NetworkObjectId types.String `tfsdk:"network_object_id"`
	Advertise       types.Bool   `tfsdk:"advertise"`
}
type DeviceOSPFAreasVirtualLinks struct {
	PeerRouterHostId           types.String                                        `tfsdk:"peer_router_host_id"`
	HelloInterval              types.Int64                                         `tfsdk:"hello_interval"`
	TransmitDelay              types.Int64                                         `tfsdk:"transmit_delay"`
	RetransmitInterval         types.Int64                                         `tfsdk:"retransmit_interval"`
	DeadInterval               types.Int64                                         `tfsdk:"dead_interval"`
	AuthenticationPassword     types.String                                        `tfsdk:"authentication_password"`
	AuthenticationAreaPassword types.String                                        `tfsdk:"authentication_area_password"`
	AuthenticationAreaMd5s     []DeviceOSPFAreasVirtualLinksAuthenticationAreaMd5s `tfsdk:"authentication_area_md5s"`
	AuthenticationMd5s         []DeviceOSPFAreasVirtualLinksAuthenticationMd5s     `tfsdk:"authentication_md5s"`
	AuthenticationKeyChainId   types.String                                        `tfsdk:"authentication_key_chain_id"`
}
type DeviceOSPFAreasInterAreaFilters struct {
	PrefixListId    types.String `tfsdk:"prefix_list_id"`
	PrefixListName  types.String `tfsdk:"prefix_list_name"`
	FilterDirection types.String `tfsdk:"filter_direction"`
}

type DeviceOSPFSummaryAddressesNetworks struct {
	Id types.String `tfsdk:"id"`
}

type DeviceOSPFAreasVirtualLinksAuthenticationAreaMd5s struct {
	Id  types.Int64  `tfsdk:"id"`
	Key types.String `tfsdk:"key"`
}
type DeviceOSPFAreasVirtualLinksAuthenticationMd5s struct {
	Id  types.Int64  `tfsdk:"id"`
	Key types.String `tfsdk:"key"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionDeviceOSPF = version.Must(version.NewVersion("7.6"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceOSPF) getPath() string {
	if data.VrfId.ValueString() != "" {
		return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/virtualrouters/%v/ospfv2routes", url.QueryEscape(data.DeviceId.ValueString()), url.QueryEscape(data.VrfId.ValueString()))
	} else {
		return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/ospfv2routes", url.QueryEscape(data.DeviceId.ValueString()))
	}
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceOSPF) toBody(ctx context.Context, state DeviceOSPF) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.ProcessId.IsNull() {
		body, _ = sjson.Set(body, "processId", data.ProcessId.ValueInt64())
	}
	if !data.RouterId.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.routerId", data.RouterId.ValueString())
	}
	if !data.Rfc1583Compatible.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.rfc1583Compatible", data.Rfc1583Compatible.ValueBool())
	}
	if !data.LogAdjacencyChanges.IsNull() {
		body, _ = sjson.Set(body, "logAdjacencyChanges.logType", data.LogAdjacencyChanges.ValueString())
	}
	if !data.IgnoreLsaMospf.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.ignoreLsaMospf", data.IgnoreLsaMospf.ValueBool())
	}
	if !data.AdministrativeDistanceInterArea.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.administrativeDistance.interArea", data.AdministrativeDistanceInterArea.ValueInt64())
	}
	if !data.AdministrativeDistanceIntraArea.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.administrativeDistance.intraArea", data.AdministrativeDistanceIntraArea.ValueInt64())
	}
	if !data.AdministrativeDistanceExternal.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.administrativeDistance.external", data.AdministrativeDistanceExternal.ValueInt64())
	}
	if !data.TimerLsaGroup.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.timers.lsaGroup", data.TimerLsaGroup.ValueInt64())
	}
	if !data.DefaultRouteAlwaysAdvertise.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.defaultInformationOriginate.alwaysAdvertise", data.DefaultRouteAlwaysAdvertise.ValueBool())
	}
	if !data.DefaultRouteMetric.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.defaultInformationOriginate.routeMetric.metricValue", data.DefaultRouteMetric.ValueInt64())
	}
	if !data.DefaultRouteMetricType.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.defaultInformationOriginate.routeMetric.metricType", data.DefaultRouteMetricType.ValueString())
	}
	if !data.DefaultRouteRouteMapId.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.defaultInformationOriginate.routeMap.id", data.DefaultRouteRouteMapId.ValueString())
	}
	if !data.NonStopForwarding.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.nsfGracefulRestart.nsfEnable", data.NonStopForwarding.ValueBool())
	}
	if !data.NonStopForwardingMechanism.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.nsfGracefulRestart.nsfMechanism", data.NonStopForwardingMechanism.ValueString())
	}
	if !data.NonStopForwardingHelperMode.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.nsfGracefulRestart.nsfHelper", data.NonStopForwardingHelperMode.ValueBool())
	}
	if !data.NonStopForwardingRestartInterval.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.nsfGracefulRestart.nsfRestartIntraval", data.NonStopForwardingRestartInterval.ValueInt64())
	}
	if !data.NonStopForwardingCapability.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.nsfGracefulRestart.nsfCapability", data.NonStopForwardingCapability.ValueBool())
	}
	if !data.NonStopForwardingStrictMode.IsNull() {
		body, _ = sjson.Set(body, "processConfiguration.nsfGracefulRestart.nsfEnforceGlobal", data.NonStopForwardingStrictMode.ValueBool())
	}
	if len(data.Areas) > 0 {
		body, _ = sjson.Set(body, "areas", []any{})
		for _, item := range data.Areas {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaId", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaType.type", item.Type.ValueString())
			}
			if !item.NoSummary.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaType.noSummary", item.NoSummary.ValueBool())
			}
			if !item.NoRedistribution.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaType.noRedistribution", item.NoRedistribution.ValueBool())
			}
			if !item.DefaultRouteMetricType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaType.routeMetric.metricType", item.DefaultRouteMetricType.ValueString())
			}
			if !item.DefaultRouteMetric.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "areaType.routeMetric.metricValue", item.DefaultRouteMetric.ValueInt64())
			}
			if len(item.Networks) > 0 {
				itemBody, _ = sjson.Set(itemBody, "areaNetworks", []any{})
				for _, childItem := range item.Networks {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "areaNetworks.-1", itemChildBody)
				}
			}
			if !item.Authentication.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authentication", item.Authentication.ValueString())
			}
			if !item.DefaultCost.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "defaultCost", item.DefaultCost.ValueInt64())
			}
			if len(item.Ranges) > 0 {
				itemBody, _ = sjson.Set(itemBody, "areaRanges", []any{})
				for _, childItem := range item.Ranges {
					itemChildBody := ""
					if !childItem.NetworkObjectId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "addressNetwork.id", childItem.NetworkObjectId.ValueString())
					}
					if !childItem.Advertise.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "advertise", childItem.Advertise.ValueBool())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "areaRanges.-1", itemChildBody)
				}
			}
			if len(item.VirtualLinks) > 0 {
				itemBody, _ = sjson.Set(itemBody, "virtualLinks", []any{})
				for _, childItem := range item.VirtualLinks {
					itemChildBody := ""
					if !childItem.PeerRouterHostId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "routerId", childItem.PeerRouterHostId.ValueString())
					}
					if !childItem.HelloInterval.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "helloInterval", childItem.HelloInterval.ValueInt64())
					}
					if !childItem.TransmitDelay.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "transmitDelay", childItem.TransmitDelay.ValueInt64())
					}
					if !childItem.RetransmitInterval.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "retransmitInterval", childItem.RetransmitInterval.ValueInt64())
					}
					if !childItem.DeadInterval.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "deadInterval", childItem.DeadInterval.ValueInt64())
					}
					if !childItem.AuthenticationPassword.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "authentication.passwdAuth.authKey", childItem.AuthenticationPassword.ValueString())
					}
					if !childItem.AuthenticationAreaPassword.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "authentication.areaAuth.passwdAuth.authKey", childItem.AuthenticationAreaPassword.ValueString())
					}
					if len(childItem.AuthenticationAreaMd5s) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "authentication.areaAuth.md5AuthList", []any{})
						for _, childChildItem := range childItem.AuthenticationAreaMd5s {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "md5KeyId", childChildItem.Id.ValueInt64())
							}
							if !childChildItem.Key.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "md5Key", childChildItem.Key.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "authentication.areaAuth.md5AuthList.-1", itemChildChildBody)
						}
					}
					if len(childItem.AuthenticationMd5s) > 0 {
						itemChildBody, _ = sjson.Set(itemChildBody, "authentication.md5AuthList", []any{})
						for _, childChildItem := range childItem.AuthenticationMd5s {
							itemChildChildBody := ""
							if !childChildItem.Id.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "md5KeyId", childChildItem.Id.ValueInt64())
							}
							if !childChildItem.Key.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "md5Key", childChildItem.Key.ValueString())
							}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "authentication.md5AuthList.-1", itemChildChildBody)
						}
					}
					if !childItem.AuthenticationKeyChainId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "authentication.keyChain.authKey.id", childItem.AuthenticationKeyChainId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "virtualLinks.-1", itemChildBody)
				}
			}
			if len(item.InterAreaFilters) > 0 {
				itemBody, _ = sjson.Set(itemBody, "filterList", []any{})
				for _, childItem := range item.InterAreaFilters {
					itemChildBody := ""
					if !childItem.PrefixListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "prefixList.id", childItem.PrefixListId.ValueString())
					}
					if !childItem.PrefixListName.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "prefixList.name", childItem.PrefixListName.ValueString())
					}
					if !childItem.FilterDirection.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "filterDirection", childItem.FilterDirection.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "filterList.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "areas.-1", itemBody)
		}
	}
	if len(data.Redistributions) > 0 {
		body, _ = sjson.Set(body, "redistributeProtocols", []any{})
		for _, item := range data.Redistributions {
			itemBody := ""
			if !item.RedistributeProtocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.RedistributeProtocol.ValueString())
			}
			if !item.AsNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "asNumber", item.AsNumber.ValueInt64())
			}
			if !item.MatchExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchExternal1", item.MatchExternal1.ValueBool())
			}
			if !item.MatchExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchExternal2", item.MatchExternal2.ValueBool())
			}
			if !item.MatchInternal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchInternal", item.MatchInternal.ValueBool())
			}
			if !item.MatchNssaExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchNssaExternal1", item.MatchNssaExternal1.ValueBool())
			}
			if !item.MatchNssaExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchNssaExternal2", item.MatchNssaExternal2.ValueBool())
			}
			if !item.ProcessId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "processId", item.ProcessId.ValueInt64())
			}
			if !item.Subnets.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "subnets", item.Subnets.ValueBool())
			}
			if !item.Metric.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMetric.metricValue", item.Metric.ValueInt64())
			}
			if !item.MetricType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMetric.metricType", item.MetricType.ValueString())
			}
			if !item.Tag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tagNumber", item.Tag.ValueInt64())
			}
			if !item.RouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMap.id", item.RouteMapId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "redistributeProtocols.-1", itemBody)
		}
	}
	if len(data.FilterRules) > 0 {
		body, _ = sjson.Set(body, "filterRules", []any{})
		for _, item := range data.FilterRules {
			itemBody := ""
			if !item.AccessListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "accessList.id", item.AccessListId.ValueString())
			}
			if !item.TrafficDirection.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.TrafficDirection.ValueString())
			}
			if !item.RoutingProcess.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.RoutingProcess.ValueString())
			}
			if !item.RoutingProcessId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "processId", item.RoutingProcessId.ValueInt64())
			}
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "inInterface.id", item.InterfaceId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "filterRules.-1", itemBody)
		}
	}
	if len(data.SummaryAddresses) > 0 {
		body, _ = sjson.Set(body, "summaryAddresses", []any{})
		for _, item := range data.SummaryAddresses {
			itemBody := ""
			if len(item.Networks) > 0 {
				itemBody, _ = sjson.Set(itemBody, "summaryNetwork", []any{})
				for _, childItem := range item.Networks {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "summaryNetwork.-1", itemChildBody)
				}
			}
			if !item.Tag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tagNumber", item.Tag.ValueInt64())
			}
			if !item.Advertise.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "advertise", item.Advertise.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "summaryAddresses.-1", itemBody)
		}
	}
	if len(data.Neighbors) > 0 {
		body, _ = sjson.Set(body, "neighbors", []any{})
		for _, item := range data.Neighbors {
			itemBody := ""
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborInterface.id", item.InterfaceId.ValueString())
			}
			if !item.NeighborHostId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipAddress.id", item.NeighborHostId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "neighbors.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceOSPF) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("processId"); value.Exists() {
		data.ProcessId = types.Int64Value(value.Int())
	} else {
		data.ProcessId = types.Int64Null()
	}
	if value := res.Get("processConfiguration.routerId"); value.Exists() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	if value := res.Get("processConfiguration.rfc1583Compatible"); value.Exists() {
		data.Rfc1583Compatible = types.BoolValue(value.Bool())
	} else {
		data.Rfc1583Compatible = types.BoolValue(false)
	}
	if value := res.Get("logAdjacencyChanges.logType"); value.Exists() {
		data.LogAdjacencyChanges = types.StringValue(value.String())
	} else {
		data.LogAdjacencyChanges = types.StringNull()
	}
	if value := res.Get("processConfiguration.ignoreLsaMospf"); value.Exists() {
		data.IgnoreLsaMospf = types.BoolValue(value.Bool())
	} else {
		data.IgnoreLsaMospf = types.BoolValue(false)
	}
	if value := res.Get("processConfiguration.administrativeDistance.interArea"); value.Exists() {
		data.AdministrativeDistanceInterArea = types.Int64Value(value.Int())
	} else {
		data.AdministrativeDistanceInterArea = types.Int64Value(110)
	}
	if value := res.Get("processConfiguration.administrativeDistance.intraArea"); value.Exists() {
		data.AdministrativeDistanceIntraArea = types.Int64Value(value.Int())
	} else {
		data.AdministrativeDistanceIntraArea = types.Int64Value(110)
	}
	if value := res.Get("processConfiguration.administrativeDistance.external"); value.Exists() {
		data.AdministrativeDistanceExternal = types.Int64Value(value.Int())
	} else {
		data.AdministrativeDistanceExternal = types.Int64Value(110)
	}
	if value := res.Get("processConfiguration.timers.lsaGroup"); value.Exists() {
		data.TimerLsaGroup = types.Int64Value(value.Int())
	} else {
		data.TimerLsaGroup = types.Int64Value(240)
	}
	if value := res.Get("processConfiguration.defaultInformationOriginate.alwaysAdvertise"); value.Exists() {
		data.DefaultRouteAlwaysAdvertise = types.BoolValue(value.Bool())
	} else {
		data.DefaultRouteAlwaysAdvertise = types.BoolNull()
	}
	if value := res.Get("processConfiguration.defaultInformationOriginate.routeMetric.metricValue"); value.Exists() {
		data.DefaultRouteMetric = types.Int64Value(value.Int())
	} else {
		data.DefaultRouteMetric = types.Int64Null()
	}
	if value := res.Get("processConfiguration.defaultInformationOriginate.routeMetric.metricType"); value.Exists() {
		data.DefaultRouteMetricType = types.StringValue(value.String())
	} else {
		data.DefaultRouteMetricType = types.StringNull()
	}
	if value := res.Get("processConfiguration.defaultInformationOriginate.routeMap.id"); value.Exists() {
		data.DefaultRouteRouteMapId = types.StringValue(value.String())
	} else {
		data.DefaultRouteRouteMapId = types.StringNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfEnable"); value.Exists() {
		data.NonStopForwarding = types.BoolValue(value.Bool())
	} else {
		data.NonStopForwarding = types.BoolNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfMechanism"); value.Exists() {
		data.NonStopForwardingMechanism = types.StringValue(value.String())
	} else {
		data.NonStopForwardingMechanism = types.StringNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfHelper"); value.Exists() {
		data.NonStopForwardingHelperMode = types.BoolValue(value.Bool())
	} else {
		data.NonStopForwardingHelperMode = types.BoolNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfRestartIntraval"); value.Exists() {
		data.NonStopForwardingRestartInterval = types.Int64Value(value.Int())
	} else {
		data.NonStopForwardingRestartInterval = types.Int64Null()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfCapability"); value.Exists() {
		data.NonStopForwardingCapability = types.BoolValue(value.Bool())
	} else {
		data.NonStopForwardingCapability = types.BoolNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfEnforceGlobal"); value.Exists() {
		data.NonStopForwardingStrictMode = types.BoolValue(value.Bool())
	} else {
		data.NonStopForwardingStrictMode = types.BoolNull()
	}
	if value := res.Get("areas"); value.Exists() {
		data.Areas = make([]DeviceOSPFAreas, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceOSPFAreas{}
			if value := res.Get("areaId"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("areaType.type"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("areaType.noSummary"); value.Exists() {
				data.NoSummary = types.BoolValue(value.Bool())
			} else {
				data.NoSummary = types.BoolNull()
			}
			if value := res.Get("areaType.noRedistribution"); value.Exists() {
				data.NoRedistribution = types.BoolValue(value.Bool())
			} else {
				data.NoRedistribution = types.BoolNull()
			}
			if value := res.Get("areaType.routeMetric.metricType"); value.Exists() {
				data.DefaultRouteMetricType = types.StringValue(value.String())
			} else {
				data.DefaultRouteMetricType = types.StringNull()
			}
			if value := res.Get("areaType.routeMetric.metricValue"); value.Exists() {
				data.DefaultRouteMetric = types.Int64Value(value.Int())
			} else {
				data.DefaultRouteMetric = types.Int64Null()
			}
			if value := res.Get("areaNetworks"); value.Exists() {
				data.Networks = make([]DeviceOSPFAreasNetworks, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceOSPFAreasNetworks{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("name"); value.Exists() {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					(*parent).Networks = append((*parent).Networks, data)
					return true
				})
			}
			if value := res.Get("authentication"); value.Exists() {
				data.Authentication = types.StringValue(value.String())
			} else {
				data.Authentication = types.StringNull()
			}
			if value := res.Get("defaultCost"); value.Exists() {
				data.DefaultCost = types.Int64Value(value.Int())
			} else {
				data.DefaultCost = types.Int64Null()
			}
			if value := res.Get("areaRanges"); value.Exists() {
				data.Ranges = make([]DeviceOSPFAreasRanges, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceOSPFAreasRanges{}
					if value := res.Get("addressNetwork.id"); value.Exists() {
						data.NetworkObjectId = types.StringValue(value.String())
					} else {
						data.NetworkObjectId = types.StringNull()
					}
					if value := res.Get("advertise"); value.Exists() {
						data.Advertise = types.BoolValue(value.Bool())
					} else {
						data.Advertise = types.BoolNull()
					}
					(*parent).Ranges = append((*parent).Ranges, data)
					return true
				})
			}
			if value := res.Get("virtualLinks"); value.Exists() {
				data.VirtualLinks = make([]DeviceOSPFAreasVirtualLinks, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceOSPFAreasVirtualLinks{}
					if value := res.Get("routerId"); value.Exists() {
						data.PeerRouterHostId = types.StringValue(value.String())
					} else {
						data.PeerRouterHostId = types.StringNull()
					}
					if value := res.Get("helloInterval"); value.Exists() {
						data.HelloInterval = types.Int64Value(value.Int())
					} else {
						data.HelloInterval = types.Int64Value(10)
					}
					if value := res.Get("transmitDelay"); value.Exists() {
						data.TransmitDelay = types.Int64Value(value.Int())
					} else {
						data.TransmitDelay = types.Int64Value(1)
					}
					if value := res.Get("retransmitInterval"); value.Exists() {
						data.RetransmitInterval = types.Int64Value(value.Int())
					} else {
						data.RetransmitInterval = types.Int64Value(5)
					}
					if value := res.Get("deadInterval"); value.Exists() {
						data.DeadInterval = types.Int64Value(value.Int())
					} else {
						data.DeadInterval = types.Int64Value(40)
					}
					if value := res.Get("authentication.passwdAuth.authKey"); value.Exists() {
						data.AuthenticationPassword = types.StringValue(value.String())
					} else {
						data.AuthenticationPassword = types.StringNull()
					}
					if value := res.Get("authentication.areaAuth.passwdAuth.authKey"); value.Exists() {
						data.AuthenticationAreaPassword = types.StringValue(value.String())
					} else {
						data.AuthenticationAreaPassword = types.StringNull()
					}
					if value := res.Get("authentication.areaAuth.md5AuthList"); value.Exists() {
						data.AuthenticationAreaMd5s = make([]DeviceOSPFAreasVirtualLinksAuthenticationAreaMd5s, 0)
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := DeviceOSPFAreasVirtualLinksAuthenticationAreaMd5s{}
							if value := res.Get("md5KeyId"); value.Exists() {
								data.Id = types.Int64Value(value.Int())
							} else {
								data.Id = types.Int64Null()
							}
							if value := res.Get("md5Key"); value.Exists() {
								data.Key = types.StringValue(value.String())
							} else {
								data.Key = types.StringNull()
							}
							(*parent).AuthenticationAreaMd5s = append((*parent).AuthenticationAreaMd5s, data)
							return true
						})
					}
					if value := res.Get("authentication.md5AuthList"); value.Exists() {
						data.AuthenticationMd5s = make([]DeviceOSPFAreasVirtualLinksAuthenticationMd5s, 0)
						value.ForEach(func(k, res gjson.Result) bool {
							parent := &data
							data := DeviceOSPFAreasVirtualLinksAuthenticationMd5s{}
							if value := res.Get("md5KeyId"); value.Exists() {
								data.Id = types.Int64Value(value.Int())
							} else {
								data.Id = types.Int64Null()
							}
							if value := res.Get("md5Key"); value.Exists() {
								data.Key = types.StringValue(value.String())
							} else {
								data.Key = types.StringNull()
							}
							(*parent).AuthenticationMd5s = append((*parent).AuthenticationMd5s, data)
							return true
						})
					}
					if value := res.Get("authentication.keyChain.authKey.id"); value.Exists() {
						data.AuthenticationKeyChainId = types.StringValue(value.String())
					} else {
						data.AuthenticationKeyChainId = types.StringNull()
					}
					(*parent).VirtualLinks = append((*parent).VirtualLinks, data)
					return true
				})
			}
			if value := res.Get("filterList"); value.Exists() {
				data.InterAreaFilters = make([]DeviceOSPFAreasInterAreaFilters, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceOSPFAreasInterAreaFilters{}
					if value := res.Get("prefixList.id"); value.Exists() {
						data.PrefixListId = types.StringValue(value.String())
					} else {
						data.PrefixListId = types.StringNull()
					}
					if value := res.Get("prefixList.name"); value.Exists() {
						data.PrefixListName = types.StringValue(value.String())
					} else {
						data.PrefixListName = types.StringNull()
					}
					if value := res.Get("filterDirection"); value.Exists() {
						data.FilterDirection = types.StringValue(value.String())
					} else {
						data.FilterDirection = types.StringNull()
					}
					(*parent).InterAreaFilters = append((*parent).InterAreaFilters, data)
					return true
				})
			}
			(*parent).Areas = append((*parent).Areas, data)
			return true
		})
	}
	if value := res.Get("redistributeProtocols"); value.Exists() {
		data.Redistributions = make([]DeviceOSPFRedistributions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceOSPFRedistributions{}
			if value := res.Get("type"); value.Exists() {
				data.RedistributeProtocol = types.StringValue(value.String())
			} else {
				data.RedistributeProtocol = types.StringNull()
			}
			if value := res.Get("asNumber"); value.Exists() {
				data.AsNumber = types.Int64Value(value.Int())
			} else {
				data.AsNumber = types.Int64Null()
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
			if value := res.Get("matchInternal"); value.Exists() {
				data.MatchInternal = types.BoolValue(value.Bool())
			} else {
				data.MatchInternal = types.BoolNull()
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
			if value := res.Get("processId"); value.Exists() {
				data.ProcessId = types.Int64Value(value.Int())
			} else {
				data.ProcessId = types.Int64Null()
			}
			if value := res.Get("subnets"); value.Exists() {
				data.Subnets = types.BoolValue(value.Bool())
			} else {
				data.Subnets = types.BoolNull()
			}
			if value := res.Get("routeMetric.metricValue"); value.Exists() {
				data.Metric = types.Int64Value(value.Int())
			} else {
				data.Metric = types.Int64Null()
			}
			if value := res.Get("routeMetric.metricType"); value.Exists() {
				data.MetricType = types.StringValue(value.String())
			} else {
				data.MetricType = types.StringValue("TYPE_2")
			}
			if value := res.Get("tagNumber"); value.Exists() {
				data.Tag = types.Int64Value(value.Int())
			} else {
				data.Tag = types.Int64Null()
			}
			if value := res.Get("routeMap.id"); value.Exists() {
				data.RouteMapId = types.StringValue(value.String())
			} else {
				data.RouteMapId = types.StringNull()
			}
			(*parent).Redistributions = append((*parent).Redistributions, data)
			return true
		})
	}
	if value := res.Get("filterRules"); value.Exists() {
		data.FilterRules = make([]DeviceOSPFFilterRules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceOSPFFilterRules{}
			if value := res.Get("accessList.id"); value.Exists() {
				data.AccessListId = types.StringValue(value.String())
			} else {
				data.AccessListId = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.TrafficDirection = types.StringValue(value.String())
			} else {
				data.TrafficDirection = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() {
				data.RoutingProcess = types.StringValue(value.String())
			} else {
				data.RoutingProcess = types.StringNull()
			}
			if value := res.Get("processId"); value.Exists() {
				data.RoutingProcessId = types.Int64Value(value.Int())
			} else {
				data.RoutingProcessId = types.Int64Null()
			}
			if value := res.Get("inInterface.id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			(*parent).FilterRules = append((*parent).FilterRules, data)
			return true
		})
	}
	if value := res.Get("summaryAddresses"); value.Exists() {
		data.SummaryAddresses = make([]DeviceOSPFSummaryAddresses, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceOSPFSummaryAddresses{}
			if value := res.Get("summaryNetwork"); value.Exists() {
				data.Networks = make([]DeviceOSPFSummaryAddressesNetworks, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceOSPFSummaryAddressesNetworks{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).Networks = append((*parent).Networks, data)
					return true
				})
			}
			if value := res.Get("tagNumber"); value.Exists() {
				data.Tag = types.Int64Value(value.Int())
			} else {
				data.Tag = types.Int64Null()
			}
			if value := res.Get("advertise"); value.Exists() {
				data.Advertise = types.BoolValue(value.Bool())
			} else {
				data.Advertise = types.BoolNull()
			}
			(*parent).SummaryAddresses = append((*parent).SummaryAddresses, data)
			return true
		})
	}
	if value := res.Get("neighbors"); value.Exists() {
		data.Neighbors = make([]DeviceOSPFNeighbors, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceOSPFNeighbors{}
			if value := res.Get("neighborInterface.id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			if value := res.Get("ipAddress.id"); value.Exists() {
				data.NeighborHostId = types.StringValue(value.String())
			} else {
				data.NeighborHostId = types.StringNull()
			}
			(*parent).Neighbors = append((*parent).Neighbors, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceOSPF) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("processId"); value.Exists() && !data.ProcessId.IsNull() {
		data.ProcessId = types.Int64Value(value.Int())
	} else {
		data.ProcessId = types.Int64Null()
	}
	if value := res.Get("processConfiguration.routerId"); value.Exists() && !data.RouterId.IsNull() {
		data.RouterId = types.StringValue(value.String())
	} else {
		data.RouterId = types.StringNull()
	}
	if value := res.Get("processConfiguration.rfc1583Compatible"); value.Exists() && !data.Rfc1583Compatible.IsNull() {
		data.Rfc1583Compatible = types.BoolValue(value.Bool())
	} else if data.Rfc1583Compatible.ValueBool() != false {
		data.Rfc1583Compatible = types.BoolNull()
	}
	if value := res.Get("logAdjacencyChanges.logType"); value.Exists() && !data.LogAdjacencyChanges.IsNull() {
		data.LogAdjacencyChanges = types.StringValue(value.String())
	} else {
		data.LogAdjacencyChanges = types.StringNull()
	}
	if value := res.Get("processConfiguration.ignoreLsaMospf"); value.Exists() && !data.IgnoreLsaMospf.IsNull() {
		data.IgnoreLsaMospf = types.BoolValue(value.Bool())
	} else if data.IgnoreLsaMospf.ValueBool() != false {
		data.IgnoreLsaMospf = types.BoolNull()
	}
	if value := res.Get("processConfiguration.administrativeDistance.interArea"); value.Exists() && !data.AdministrativeDistanceInterArea.IsNull() {
		data.AdministrativeDistanceInterArea = types.Int64Value(value.Int())
	} else if data.AdministrativeDistanceInterArea.ValueInt64() != 110 {
		data.AdministrativeDistanceInterArea = types.Int64Null()
	}
	if value := res.Get("processConfiguration.administrativeDistance.intraArea"); value.Exists() && !data.AdministrativeDistanceIntraArea.IsNull() {
		data.AdministrativeDistanceIntraArea = types.Int64Value(value.Int())
	} else if data.AdministrativeDistanceIntraArea.ValueInt64() != 110 {
		data.AdministrativeDistanceIntraArea = types.Int64Null()
	}
	if value := res.Get("processConfiguration.administrativeDistance.external"); value.Exists() && !data.AdministrativeDistanceExternal.IsNull() {
		data.AdministrativeDistanceExternal = types.Int64Value(value.Int())
	} else if data.AdministrativeDistanceExternal.ValueInt64() != 110 {
		data.AdministrativeDistanceExternal = types.Int64Null()
	}
	if value := res.Get("processConfiguration.timers.lsaGroup"); value.Exists() && !data.TimerLsaGroup.IsNull() {
		data.TimerLsaGroup = types.Int64Value(value.Int())
	} else if data.TimerLsaGroup.ValueInt64() != 240 {
		data.TimerLsaGroup = types.Int64Null()
	}
	if value := res.Get("processConfiguration.defaultInformationOriginate.alwaysAdvertise"); value.Exists() && !data.DefaultRouteAlwaysAdvertise.IsNull() {
		data.DefaultRouteAlwaysAdvertise = types.BoolValue(value.Bool())
	} else {
		data.DefaultRouteAlwaysAdvertise = types.BoolNull()
	}
	if value := res.Get("processConfiguration.defaultInformationOriginate.routeMetric.metricValue"); value.Exists() && !data.DefaultRouteMetric.IsNull() {
		data.DefaultRouteMetric = types.Int64Value(value.Int())
	} else {
		data.DefaultRouteMetric = types.Int64Null()
	}
	if value := res.Get("processConfiguration.defaultInformationOriginate.routeMetric.metricType"); value.Exists() && !data.DefaultRouteMetricType.IsNull() {
		data.DefaultRouteMetricType = types.StringValue(value.String())
	} else {
		data.DefaultRouteMetricType = types.StringNull()
	}
	if value := res.Get("processConfiguration.defaultInformationOriginate.routeMap.id"); value.Exists() && !data.DefaultRouteRouteMapId.IsNull() {
		data.DefaultRouteRouteMapId = types.StringValue(value.String())
	} else {
		data.DefaultRouteRouteMapId = types.StringNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfEnable"); value.Exists() && !data.NonStopForwarding.IsNull() {
		data.NonStopForwarding = types.BoolValue(value.Bool())
	} else {
		data.NonStopForwarding = types.BoolNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfMechanism"); value.Exists() && !data.NonStopForwardingMechanism.IsNull() {
		data.NonStopForwardingMechanism = types.StringValue(value.String())
	} else {
		data.NonStopForwardingMechanism = types.StringNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfHelper"); value.Exists() && !data.NonStopForwardingHelperMode.IsNull() {
		data.NonStopForwardingHelperMode = types.BoolValue(value.Bool())
	} else {
		data.NonStopForwardingHelperMode = types.BoolNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfRestartIntraval"); value.Exists() && !data.NonStopForwardingRestartInterval.IsNull() {
		data.NonStopForwardingRestartInterval = types.Int64Value(value.Int())
	} else {
		data.NonStopForwardingRestartInterval = types.Int64Null()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfCapability"); value.Exists() && !data.NonStopForwardingCapability.IsNull() {
		data.NonStopForwardingCapability = types.BoolValue(value.Bool())
	} else {
		data.NonStopForwardingCapability = types.BoolNull()
	}
	if value := res.Get("processConfiguration.nsfGracefulRestart.nsfEnforceGlobal"); value.Exists() && !data.NonStopForwardingStrictMode.IsNull() {
		data.NonStopForwardingStrictMode = types.BoolValue(value.Bool())
	} else {
		data.NonStopForwardingStrictMode = types.BoolNull()
	}
	for i := 0; i < len(data.Areas); i++ {
		keys := [...]string{"areaId"}
		keyValues := [...]string{data.Areas[i].Id.ValueString()}

		parent := &data
		data := (*parent).Areas[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("areas").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Areas[%d] = %+v",
				i,
				(*parent).Areas[i],
			))
			(*parent).Areas = slices.Delete((*parent).Areas, i, i+1)
			i--

			continue
		}
		if value := res.Get("areaId"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("areaType.type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("areaType.noSummary"); value.Exists() && !data.NoSummary.IsNull() {
			data.NoSummary = types.BoolValue(value.Bool())
		} else {
			data.NoSummary = types.BoolNull()
		}
		if value := res.Get("areaType.noRedistribution"); value.Exists() && !data.NoRedistribution.IsNull() {
			data.NoRedistribution = types.BoolValue(value.Bool())
		} else {
			data.NoRedistribution = types.BoolNull()
		}
		if value := res.Get("areaType.routeMetric.metricType"); value.Exists() && !data.DefaultRouteMetricType.IsNull() {
			data.DefaultRouteMetricType = types.StringValue(value.String())
		} else {
			data.DefaultRouteMetricType = types.StringNull()
		}
		if value := res.Get("areaType.routeMetric.metricValue"); value.Exists() && !data.DefaultRouteMetric.IsNull() {
			data.DefaultRouteMetric = types.Int64Value(value.Int())
		} else {
			data.DefaultRouteMetric = types.Int64Null()
		}
		for i := 0; i < len(data.Networks); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Networks[i].Id.ValueString()}

			parent := &data
			data := (*parent).Networks[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("areaNetworks").ForEach(
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
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).Networks[i] = data
		}
		if value := res.Get("authentication"); value.Exists() && !data.Authentication.IsNull() {
			data.Authentication = types.StringValue(value.String())
		} else {
			data.Authentication = types.StringNull()
		}
		if value := res.Get("defaultCost"); value.Exists() && !data.DefaultCost.IsNull() {
			data.DefaultCost = types.Int64Value(value.Int())
		} else {
			data.DefaultCost = types.Int64Null()
		}
		for i := 0; i < len(data.Ranges); i++ {
			keys := [...]string{"addressNetwork.id"}
			keyValues := [...]string{data.Ranges[i].NetworkObjectId.ValueString()}

			parent := &data
			data := (*parent).Ranges[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("areaRanges").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Ranges[%d] = %+v",
					i,
					(*parent).Ranges[i],
				))
				(*parent).Ranges = slices.Delete((*parent).Ranges, i, i+1)
				i--

				continue
			}
			if value := res.Get("addressNetwork.id"); value.Exists() && !data.NetworkObjectId.IsNull() {
				data.NetworkObjectId = types.StringValue(value.String())
			} else {
				data.NetworkObjectId = types.StringNull()
			}
			if value := res.Get("advertise"); value.Exists() && !data.Advertise.IsNull() {
				data.Advertise = types.BoolValue(value.Bool())
			} else {
				data.Advertise = types.BoolNull()
			}
			(*parent).Ranges[i] = data
		}
		for i := 0; i < len(data.VirtualLinks); i++ {
			keys := [...]string{"routerId"}
			keyValues := [...]string{data.VirtualLinks[i].PeerRouterHostId.ValueString()}

			parent := &data
			data := (*parent).VirtualLinks[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("virtualLinks").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing VirtualLinks[%d] = %+v",
					i,
					(*parent).VirtualLinks[i],
				))
				(*parent).VirtualLinks = slices.Delete((*parent).VirtualLinks, i, i+1)
				i--

				continue
			}
			if value := res.Get("routerId"); value.Exists() && !data.PeerRouterHostId.IsNull() {
				data.PeerRouterHostId = types.StringValue(value.String())
			} else {
				data.PeerRouterHostId = types.StringNull()
			}
			if value := res.Get("helloInterval"); value.Exists() && !data.HelloInterval.IsNull() {
				data.HelloInterval = types.Int64Value(value.Int())
			} else if data.HelloInterval.ValueInt64() != 10 {
				data.HelloInterval = types.Int64Null()
			}
			if value := res.Get("transmitDelay"); value.Exists() && !data.TransmitDelay.IsNull() {
				data.TransmitDelay = types.Int64Value(value.Int())
			} else if data.TransmitDelay.ValueInt64() != 1 {
				data.TransmitDelay = types.Int64Null()
			}
			if value := res.Get("retransmitInterval"); value.Exists() && !data.RetransmitInterval.IsNull() {
				data.RetransmitInterval = types.Int64Value(value.Int())
			} else if data.RetransmitInterval.ValueInt64() != 5 {
				data.RetransmitInterval = types.Int64Null()
			}
			if value := res.Get("deadInterval"); value.Exists() && !data.DeadInterval.IsNull() {
				data.DeadInterval = types.Int64Value(value.Int())
			} else if data.DeadInterval.ValueInt64() != 40 {
				data.DeadInterval = types.Int64Null()
			}
			if value := res.Get("authentication.passwdAuth.authKey"); value.Exists() && !data.AuthenticationPassword.IsNull() {
				data.AuthenticationPassword = types.StringValue(value.String())
			} else {
				data.AuthenticationPassword = types.StringNull()
			}
			if value := res.Get("authentication.areaAuth.passwdAuth.authKey"); value.Exists() && !data.AuthenticationAreaPassword.IsNull() {
				data.AuthenticationAreaPassword = types.StringValue(value.String())
			} else {
				data.AuthenticationAreaPassword = types.StringNull()
			}
			for i := 0; i < len(data.AuthenticationAreaMd5s); i++ {
				keys := [...]string{"md5KeyId"}
				keyValues := [...]string{strconv.FormatInt(data.AuthenticationAreaMd5s[i].Id.ValueInt64(), 10)}

				parent := &data
				data := (*parent).AuthenticationAreaMd5s[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("authentication.areaAuth.md5AuthList").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing AuthenticationAreaMd5s[%d] = %+v",
						i,
						(*parent).AuthenticationAreaMd5s[i],
					))
					(*parent).AuthenticationAreaMd5s = slices.Delete((*parent).AuthenticationAreaMd5s, i, i+1)
					i--

					continue
				}
				if value := res.Get("md5KeyId"); value.Exists() && !data.Id.IsNull() {
					data.Id = types.Int64Value(value.Int())
				} else {
					data.Id = types.Int64Null()
				}
				if value := res.Get("md5Key"); value.Exists() && !data.Key.IsNull() {
					data.Key = types.StringValue(value.String())
				} else {
					data.Key = types.StringNull()
				}
				(*parent).AuthenticationAreaMd5s[i] = data
			}
			for i := 0; i < len(data.AuthenticationMd5s); i++ {
				keys := [...]string{"md5KeyId"}
				keyValues := [...]string{strconv.FormatInt(data.AuthenticationMd5s[i].Id.ValueInt64(), 10)}

				parent := &data
				data := (*parent).AuthenticationMd5s[i]
				parentRes := &res
				var res gjson.Result

				parentRes.Get("authentication.md5AuthList").ForEach(
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
					tflog.Debug(ctx, fmt.Sprintf("removing AuthenticationMd5s[%d] = %+v",
						i,
						(*parent).AuthenticationMd5s[i],
					))
					(*parent).AuthenticationMd5s = slices.Delete((*parent).AuthenticationMd5s, i, i+1)
					i--

					continue
				}
				if value := res.Get("md5KeyId"); value.Exists() && !data.Id.IsNull() {
					data.Id = types.Int64Value(value.Int())
				} else {
					data.Id = types.Int64Null()
				}
				if value := res.Get("md5Key"); value.Exists() && !data.Key.IsNull() {
					data.Key = types.StringValue(value.String())
				} else {
					data.Key = types.StringNull()
				}
				(*parent).AuthenticationMd5s[i] = data
			}
			if value := res.Get("authentication.keyChain.authKey.id"); value.Exists() && !data.AuthenticationKeyChainId.IsNull() {
				data.AuthenticationKeyChainId = types.StringValue(value.String())
			} else {
				data.AuthenticationKeyChainId = types.StringNull()
			}
			(*parent).VirtualLinks[i] = data
		}
		for i := 0; i < len(data.InterAreaFilters); i++ {
			keys := [...]string{"prefixList.id"}
			keyValues := [...]string{data.InterAreaFilters[i].PrefixListId.ValueString()}

			parent := &data
			data := (*parent).InterAreaFilters[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("filterList").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing InterAreaFilters[%d] = %+v",
					i,
					(*parent).InterAreaFilters[i],
				))
				(*parent).InterAreaFilters = slices.Delete((*parent).InterAreaFilters, i, i+1)
				i--

				continue
			}
			if value := res.Get("prefixList.id"); value.Exists() && !data.PrefixListId.IsNull() {
				data.PrefixListId = types.StringValue(value.String())
			} else {
				data.PrefixListId = types.StringNull()
			}
			if value := res.Get("prefixList.name"); value.Exists() && !data.PrefixListName.IsNull() {
				data.PrefixListName = types.StringValue(value.String())
			} else {
				data.PrefixListName = types.StringNull()
			}
			if value := res.Get("filterDirection"); value.Exists() && !data.FilterDirection.IsNull() {
				data.FilterDirection = types.StringValue(value.String())
			} else {
				data.FilterDirection = types.StringNull()
			}
			(*parent).InterAreaFilters[i] = data
		}
		(*parent).Areas[i] = data
	}
	for i := 0; i < len(data.Redistributions); i++ {
		keys := [...]string{"type"}
		keyValues := [...]string{data.Redistributions[i].RedistributeProtocol.ValueString()}

		parent := &data
		data := (*parent).Redistributions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("redistributeProtocols").ForEach(
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
		if value := res.Get("type"); value.Exists() && !data.RedistributeProtocol.IsNull() {
			data.RedistributeProtocol = types.StringValue(value.String())
		} else {
			data.RedistributeProtocol = types.StringNull()
		}
		if value := res.Get("asNumber"); value.Exists() && !data.AsNumber.IsNull() {
			data.AsNumber = types.Int64Value(value.Int())
		} else {
			data.AsNumber = types.Int64Null()
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
		if value := res.Get("matchInternal"); value.Exists() && !data.MatchInternal.IsNull() {
			data.MatchInternal = types.BoolValue(value.Bool())
		} else {
			data.MatchInternal = types.BoolNull()
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
		if value := res.Get("processId"); value.Exists() && !data.ProcessId.IsNull() {
			data.ProcessId = types.Int64Value(value.Int())
		} else {
			data.ProcessId = types.Int64Null()
		}
		if value := res.Get("subnets"); value.Exists() && !data.Subnets.IsNull() {
			data.Subnets = types.BoolValue(value.Bool())
		} else {
			data.Subnets = types.BoolNull()
		}
		if value := res.Get("routeMetric.metricValue"); value.Exists() && !data.Metric.IsNull() {
			data.Metric = types.Int64Value(value.Int())
		} else {
			data.Metric = types.Int64Null()
		}
		if value := res.Get("routeMetric.metricType"); value.Exists() && !data.MetricType.IsNull() {
			data.MetricType = types.StringValue(value.String())
		} else if data.MetricType.ValueString() != "TYPE_2" {
			data.MetricType = types.StringNull()
		}
		if value := res.Get("tagNumber"); value.Exists() && !data.Tag.IsNull() {
			data.Tag = types.Int64Value(value.Int())
		} else {
			data.Tag = types.Int64Null()
		}
		if value := res.Get("routeMap.id"); value.Exists() && !data.RouteMapId.IsNull() {
			data.RouteMapId = types.StringValue(value.String())
		} else {
			data.RouteMapId = types.StringNull()
		}
		(*parent).Redistributions[i] = data
	}
	for i := 0; i < len(data.FilterRules); i++ {
		keys := [...]string{"accessList.id", "type"}
		keyValues := [...]string{data.FilterRules[i].AccessListId.ValueString(), data.FilterRules[i].TrafficDirection.ValueString()}

		parent := &data
		data := (*parent).FilterRules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("filterRules").ForEach(
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
		if value := res.Get("accessList.id"); value.Exists() && !data.AccessListId.IsNull() {
			data.AccessListId = types.StringValue(value.String())
		} else {
			data.AccessListId = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.TrafficDirection.IsNull() {
			data.TrafficDirection = types.StringValue(value.String())
		} else {
			data.TrafficDirection = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.RoutingProcess.IsNull() {
			data.RoutingProcess = types.StringValue(value.String())
		} else {
			data.RoutingProcess = types.StringNull()
		}
		if value := res.Get("processId"); value.Exists() && !data.RoutingProcessId.IsNull() {
			data.RoutingProcessId = types.Int64Value(value.Int())
		} else {
			data.RoutingProcessId = types.Int64Null()
		}
		if value := res.Get("inInterface.id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		(*parent).FilterRules[i] = data
	}
	{
		l := len(res.Get("summaryAddresses").Array())
		tflog.Debug(ctx, fmt.Sprintf("summaryAddresses array resizing from %d to %d", len(data.SummaryAddresses), l))
		for i := len(data.SummaryAddresses); i < l; i++ {
			data.SummaryAddresses = append(data.SummaryAddresses, DeviceOSPFSummaryAddresses{})
		}
		if len(data.SummaryAddresses) > l {
			data.SummaryAddresses = data.SummaryAddresses[:l]
		}
	}
	for i := range data.SummaryAddresses {
		parent := &data
		data := (*parent).SummaryAddresses[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("summaryAddresses.%d", i))
		for i := 0; i < len(data.Networks); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Networks[i].Id.ValueString()}

			parent := &data
			data := (*parent).Networks[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("summaryNetwork").ForEach(
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
		if value := res.Get("tagNumber"); value.Exists() && !data.Tag.IsNull() {
			data.Tag = types.Int64Value(value.Int())
		} else {
			data.Tag = types.Int64Null()
		}
		if value := res.Get("advertise"); value.Exists() && !data.Advertise.IsNull() {
			data.Advertise = types.BoolValue(value.Bool())
		} else {
			data.Advertise = types.BoolNull()
		}
		(*parent).SummaryAddresses[i] = data
	}
	for i := 0; i < len(data.Neighbors); i++ {
		keys := [...]string{"neighborInterface.id", "ipAddress.id"}
		keyValues := [...]string{data.Neighbors[i].InterfaceId.ValueString(), data.Neighbors[i].NeighborHostId.ValueString()}

		parent := &data
		data := (*parent).Neighbors[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("neighbors").ForEach(
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
		if value := res.Get("ipAddress.id"); value.Exists() && !data.NeighborHostId.IsNull() {
			data.NeighborHostId = types.StringValue(value.String())
		} else {
			data.NeighborHostId = types.StringNull()
		}
		(*parent).Neighbors[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceOSPF) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

func (data DeviceOSPF) adjustBody(ctx context.Context, req string) string {
	// odd -> 1, even -> 2
	req, _ = sjson.Set(req, "enableProcess", fmt.Sprintf("PROCESS_%v", 2-data.ProcessId.ValueInt64()%2))
	if len(data.Redistributions) > 0 {
		for i, item := range data.Redistributions {
			if item.RedistributeProtocol.ValueString() == "RedistributeOSPF" {
				// pairing consecutive numbers (1<>2, 3<>4):
				var peerProcessId = ((data.ProcessId.ValueInt64() - 1) ^ 1) + 1
				req, _ = sjson.Set(req, fmt.Sprintf("redistributeProtocols.%d.processId", i), peerProcessId)
			}
		}
	}
	return req
}
