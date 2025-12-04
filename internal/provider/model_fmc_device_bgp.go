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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceBGP struct {
	Id                             types.String                      `tfsdk:"id"`
	Domain                         types.String                      `tfsdk:"domain"`
	VrfId                          types.String                      `tfsdk:"vrf_id"`
	DeviceId                       types.String                      `tfsdk:"device_id"`
	Name                           types.String                      `tfsdk:"name"`
	Type                           types.String                      `tfsdk:"type"`
	AsNumber                       types.String                      `tfsdk:"as_number"`
	Ipv4AddressFamilyId            types.String                      `tfsdk:"ipv4_address_family_id"`
	Ipv4AddressFamilyType          types.String                      `tfsdk:"ipv4_address_family_type"`
	Ipv4LearnedRouteMapId          types.String                      `tfsdk:"ipv4_learned_route_map_id"`
	Ipv4DefaultInformationOrginate types.Bool                        `tfsdk:"ipv4_default_information_orginate"`
	Ipv4AutoSummary                types.Bool                        `tfsdk:"ipv4_auto_summary"`
	Ipv4SuppressInactiveRoutes     types.Bool                        `tfsdk:"ipv4_suppress_inactive_routes"`
	Ipv4Synchronization            types.Bool                        `tfsdk:"ipv4_synchronization"`
	Ipv4RedistributeIbgpIntoIgp    types.Bool                        `tfsdk:"ipv4_redistribute_ibgp_into_igp"`
	Ipv4ExternalDistance           types.Int64                       `tfsdk:"ipv4_external_distance"`
	Ipv4InternalDistance           types.Int64                       `tfsdk:"ipv4_internal_distance"`
	Ipv4LocalDistance              types.Int64                       `tfsdk:"ipv4_local_distance"`
	Ipv4NumberOfIbgpPaths          types.Int64                       `tfsdk:"ipv4_number_of_ibgp_paths"`
	Ipv4NumberOfEbgpPaths          types.Int64                       `tfsdk:"ipv4_number_of_ebgp_paths"`
	Ipv4Neighbors                  []DeviceBGPIpv4Neighbors          `tfsdk:"ipv4_neighbors"`
	Ipv4AggregateAddresses         []DeviceBGPIpv4AggregateAddresses `tfsdk:"ipv4_aggregate_addresses"`
	Ipv4Filterings                 []DeviceBGPIpv4Filterings         `tfsdk:"ipv4_filterings"`
	Ipv4Networks                   []DeviceBGPIpv4Networks           `tfsdk:"ipv4_networks"`
	Ipv4Redistributions            []DeviceBGPIpv4Redistributions    `tfsdk:"ipv4_redistributions"`
	Ipv4RouteInjections            []DeviceBGPIpv4RouteInjections    `tfsdk:"ipv4_route_injections"`
	Ipv4ImportRouteTargets         types.List                        `tfsdk:"ipv4_import_route_targets"`
	Ipv4ExportRouteTargets         types.List                        `tfsdk:"ipv4_export_route_targets"`
	Ipv4ImportGlobalVrfRouteMapId  types.String                      `tfsdk:"ipv4_import_global_vrf_route_map_id"`
	Ipv4ExportGlobalVrfRouteMapId  types.String                      `tfsdk:"ipv4_export_global_vrf_route_map_id"`
	Ipv4ImportUserVrfRouteMapId    types.String                      `tfsdk:"ipv4_import_user_vrf_route_map_id"`
	Ipv4ExportUserVrfRouteMapId    types.String                      `tfsdk:"ipv4_export_user_vrf_route_map_id"`
}

type DeviceBGPIpv4Neighbors struct {
	Address                         types.String                                `tfsdk:"address"`
	RemoteAs                        types.String                                `tfsdk:"remote_as"`
	BfdFallover                     types.String                                `tfsdk:"bfd_fallover"`
	UpdateSourceInterfaceId         types.String                                `tfsdk:"update_source_interface_id"`
	EnableAddress                   types.Bool                                  `tfsdk:"enable_address"`
	AsOverride                      types.Bool                                  `tfsdk:"as_override"`
	GracefulRestart                 types.Bool                                  `tfsdk:"graceful_restart"`
	ShutdownAdministratively        types.Bool                                  `tfsdk:"shutdown_administratively"`
	Description                     types.String                                `tfsdk:"description"`
	FilterAccessLists               []DeviceBGPIpv4NeighborsFilterAccessLists   `tfsdk:"filter_access_lists"`
	FilterRouteMaps                 []DeviceBGPIpv4NeighborsFilterRouteMaps     `tfsdk:"filter_route_maps"`
	FilterPrefixLists               []DeviceBGPIpv4NeighborsFilterPrefixLists   `tfsdk:"filter_prefix_lists"`
	FilterAsPaths                   []DeviceBGPIpv4NeighborsFilterAsPaths       `tfsdk:"filter_as_paths"`
	FilterMaximumPrefixes           types.Int64                                 `tfsdk:"filter_maximum_prefixes"`
	FilterWarningOnly               types.Bool                                  `tfsdk:"filter_warning_only"`
	FilterThresholdValue            types.Int64                                 `tfsdk:"filter_threshold_value"`
	FilterRestartInterval           types.Int64                                 `tfsdk:"filter_restart_interval"`
	RoutesAdvertisementInterval     types.Int64                                 `tfsdk:"routes_advertisement_interval"`
	RoutesRemovePrivateAs           types.Bool                                  `tfsdk:"routes_remove_private_as"`
	RoutesGenerateDefaultRouteMapId types.String                                `tfsdk:"routes_generate_default_route_map_id"`
	RoutesAdvertiseMaps             []DeviceBGPIpv4NeighborsRoutesAdvertiseMaps `tfsdk:"routes_advertise_maps"`
	KeepaliveInterval               types.Int64                                 `tfsdk:"keepalive_interval"`
	HoldTime                        types.Int64                                 `tfsdk:"hold_time"`
	MinimumHoldTime                 types.Int64                                 `tfsdk:"minimum_hold_time"`
	AuthenticationPassword          types.String                                `tfsdk:"authentication_password"`
	SendCommunityAttribute          types.Bool                                  `tfsdk:"send_community_attribute"`
	NextHopSelf                     types.Bool                                  `tfsdk:"next_hop_self"`
	DisableConnectionVerification   types.Bool                                  `tfsdk:"disable_connection_verification"`
	TcpPathMtuDiscovery             types.Bool                                  `tfsdk:"tcp_path_mtu_discovery"`
	MaxHopCount                     types.Int64                                 `tfsdk:"max_hop_count"`
	TcpTransportMode                types.Bool                                  `tfsdk:"tcp_transport_mode"`
	Weight                          types.Int64                                 `tfsdk:"weight"`
	Version                         types.String                                `tfsdk:"version"`
	CustomizedLocalAsNumber         types.String                                `tfsdk:"customized_local_as_number"`
	CustomizedNoPrepend             types.Bool                                  `tfsdk:"customized_no_prepend"`
	CustomizedReplaceAs             types.Bool                                  `tfsdk:"customized_replace_as"`
	CustomizedAcceptBothAs          types.Bool                                  `tfsdk:"customized_accept_both_as"`
}

type DeviceBGPIpv4AggregateAddresses struct {
	GenerateAs     types.Bool   `tfsdk:"generate_as"`
	Filter         types.Bool   `tfsdk:"filter"`
	NetworkId      types.String `tfsdk:"network_id"`
	AdvertiseMapId types.String `tfsdk:"advertise_map_id"`
	AttributeMapId types.String `tfsdk:"attribute_map_id"`
	SuppressMapId  types.String `tfsdk:"suppress_map_id"`
}

type DeviceBGPIpv4Filterings struct {
	AccessListId types.String `tfsdk:"access_list_id"`
	Direction    types.String `tfsdk:"direction"`
	Protocol     types.String `tfsdk:"protocol"`
	ProcessId    types.String `tfsdk:"process_id"`
}

type DeviceBGPIpv4Networks struct {
	NetworkId  types.String `tfsdk:"network_id"`
	RouteMapId types.String `tfsdk:"route_map_id"`
}

type DeviceBGPIpv4Redistributions struct {
	SourceProtocol     types.String `tfsdk:"source_protocol"`
	RouteMapId         types.String `tfsdk:"route_map_id"`
	Metric             types.Int64  `tfsdk:"metric"`
	ProcessId          types.String `tfsdk:"process_id"`
	MatchExternal1     types.Bool   `tfsdk:"match_external1"`
	MatchExternal2     types.Bool   `tfsdk:"match_external2"`
	MatchInternal      types.Bool   `tfsdk:"match_internal"`
	MatchNssaExternal1 types.Bool   `tfsdk:"match_nssa_external1"`
	MatchNssaExternal2 types.Bool   `tfsdk:"match_nssa_external2"`
}

type DeviceBGPIpv4RouteInjections struct {
	InjectRouteMapId  types.String `tfsdk:"inject_route_map_id"`
	ExistRouteMapId   types.String `tfsdk:"exist_route_map_id"`
	InheritAttributes types.String `tfsdk:"inherit_attributes"`
}

type DeviceBGPIpv4NeighborsFilterAccessLists struct {
	AccessListId    types.String `tfsdk:"access_list_id"`
	UpdateDirection types.String `tfsdk:"update_direction"`
}
type DeviceBGPIpv4NeighborsFilterRouteMaps struct {
	RouteMapId      types.String `tfsdk:"route_map_id"`
	UpdateDirection types.String `tfsdk:"update_direction"`
}
type DeviceBGPIpv4NeighborsFilterPrefixLists struct {
	PrefixListId    types.String `tfsdk:"prefix_list_id"`
	UpdateDirection types.String `tfsdk:"update_direction"`
}
type DeviceBGPIpv4NeighborsFilterAsPaths struct {
	AsPathId        types.String `tfsdk:"as_path_id"`
	AsPathName      types.String `tfsdk:"as_path_name"`
	UpdateDirection types.String `tfsdk:"update_direction"`
}
type DeviceBGPIpv4NeighborsRoutesAdvertiseMaps struct {
	AdvertiseMapId     types.String `tfsdk:"advertise_map_id"`
	UseExistMap        types.Bool   `tfsdk:"use_exist_map"`
	ExistNonexistMapId types.String `tfsdk:"exist_nonexist_map_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceBGP) getPath() string {
	if data.VrfId.ValueString() != "" {
		return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/virtualrouters/%v/bgp", url.QueryEscape(data.DeviceId.ValueString()), url.QueryEscape(data.VrfId.ValueString()))
	} else {
		return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/bgp", url.QueryEscape(data.DeviceId.ValueString()))
	}
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceBGP) toBody(ctx context.Context, state DeviceBGP) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Ipv4AddressFamilyId.IsNull() && !data.Ipv4AddressFamilyId.IsUnknown() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.id", data.Ipv4AddressFamilyId.ValueString())
	}
	if !data.Ipv4LearnedRouteMapId.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.aftableMap.id", data.Ipv4LearnedRouteMapId.ValueString())
	}
	if !data.Ipv4DefaultInformationOrginate.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.defaultInformationOrginate", data.Ipv4DefaultInformationOrginate.ValueBool())
	}
	if !data.Ipv4AutoSummary.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.autoSummary", data.Ipv4AutoSummary.ValueBool())
	}
	if !data.Ipv4SuppressInactiveRoutes.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.bgpSupressInactive", data.Ipv4SuppressInactiveRoutes.ValueBool())
	}
	if !data.Ipv4Synchronization.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.synchronization", data.Ipv4Synchronization.ValueBool())
	}
	if !data.Ipv4RedistributeIbgpIntoIgp.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.bgpRedistributeInternal", data.Ipv4RedistributeIbgpIntoIgp.ValueBool())
	}
	if !data.Ipv4ExternalDistance.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distance.externalDistance", data.Ipv4ExternalDistance.ValueInt64())
	}
	if !data.Ipv4InternalDistance.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distance.internalDistance", data.Ipv4InternalDistance.ValueInt64())
	}
	if !data.Ipv4LocalDistance.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distance.localDistance", data.Ipv4LocalDistance.ValueInt64())
	}
	if !data.Ipv4NumberOfIbgpPaths.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.ibgp", data.Ipv4NumberOfIbgpPaths.ValueInt64())
	}
	if !data.Ipv4NumberOfEbgpPaths.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.ebgp", data.Ipv4NumberOfEbgpPaths.ValueInt64())
	}
	if len(data.Ipv4Neighbors) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.neighbors", []any{})
		for _, item := range data.Ipv4Neighbors {
			itemBody := ""
			if !item.Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4Address", item.Address.ValueString())
			}
			if !item.RemoteAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "remoteAs", item.RemoteAs.ValueString())
			}
			if !item.BfdFallover.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.fallOverBFD", item.BfdFallover.ValueString())
			}
			if !item.UpdateSourceInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.updateSource.id", item.UpdateSourceInterfaceId.ValueString())
			}
			if !item.EnableAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.enableAddress", item.EnableAddress.ValueBool())
			}
			if !item.AsOverride.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.asOverride", item.AsOverride.ValueBool())
			}
			if !item.GracefulRestart.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborHaMode.disable", item.GracefulRestart.ValueBool())
			}
			if !item.ShutdownAdministratively.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.shutdown", item.ShutdownAdministratively.ValueBool())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborGeneral.description", item.Description.ValueString())
			}
			if len(item.FilterAccessLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborDistributeLists", []any{})
				for _, childItem := range item.FilterAccessLists {
					itemChildBody := ""
					if !childItem.AccessListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "accessList.id", childItem.AccessListId.ValueString())
					}
					if !childItem.UpdateDirection.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "filterUpdateAction", childItem.UpdateDirection.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "neighborFiltering.neighborDistributeLists.-1", itemChildBody)
				}
			}
			if len(item.FilterRouteMaps) > 0 {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborRouteMap", []any{})
				for _, childItem := range item.FilterRouteMaps {
					itemChildBody := ""
					if !childItem.RouteMapId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "routeMap.id", childItem.RouteMapId.ValueString())
					}
					if !childItem.UpdateDirection.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "filterUpdateAction", childItem.UpdateDirection.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "neighborFiltering.neighborRouteMap.-1", itemChildBody)
				}
			}
			if len(item.FilterPrefixLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.ipv4PrefixListFilter", []any{})
				for _, childItem := range item.FilterPrefixLists {
					itemChildBody := ""
					if !childItem.PrefixListId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "ipv4PrefixList.id", childItem.PrefixListId.ValueString())
					}
					if !childItem.UpdateDirection.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "filterUpdateAction", childItem.UpdateDirection.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "neighborFiltering.ipv4PrefixListFilter.-1", itemChildBody)
				}
			}
			if len(item.FilterAsPaths) > 0 {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborFilterList", []any{})
				for _, childItem := range item.FilterAsPaths {
					itemChildBody := ""
					if !childItem.AsPathId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "asPathList.id", childItem.AsPathId.ValueString())
					}
					if !childItem.AsPathName.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "asPathList.name", childItem.AsPathName.ValueString())
					}
					if !childItem.UpdateDirection.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "filterUpdateAction", childItem.UpdateDirection.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "neighborFiltering.neighborFilterList.-1", itemChildBody)
				}
			}
			if !item.FilterMaximumPrefixes.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborMaximumPrefix.maxPrefixLimit", item.FilterMaximumPrefixes.ValueInt64())
			}
			if !item.FilterWarningOnly.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborMaximumPrefix.warningOnly", item.FilterWarningOnly.ValueBool())
			}
			if !item.FilterThresholdValue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborMaximumPrefix.thresholdValue", item.FilterThresholdValue.ValueInt64())
			}
			if !item.FilterRestartInterval.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborMaximumPrefix.restartInterval", item.FilterRestartInterval.ValueInt64())
			}
			if !item.RoutesAdvertisementInterval.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborRoutes.advertisementInterval", item.RoutesAdvertisementInterval.ValueInt64())
			}
			if !item.RoutesRemovePrivateAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborRoutes.removePrivateAs", item.RoutesRemovePrivateAs.ValueBool())
			}
			if !item.RoutesGenerateDefaultRouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborFiltering.neighborDefaultOriginate.routeMap.id", item.RoutesGenerateDefaultRouteMapId.ValueString())
			}
			if len(item.RoutesAdvertiseMaps) > 0 {
				itemBody, _ = sjson.Set(itemBody, "neighborRoutes.neighborAdvertiseMaps", []any{})
				for _, childItem := range item.RoutesAdvertiseMaps {
					itemChildBody := ""
					if !childItem.AdvertiseMapId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "routeMap.id", childItem.AdvertiseMapId.ValueString())
					}
					if !childItem.UseExistMap.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "existMap", childItem.UseExistMap.ValueBool())
					}
					if !childItem.ExistNonexistMapId.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "existRouteMap.id", childItem.ExistNonexistMapId.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "neighborRoutes.neighborAdvertiseMaps.-1", itemChildBody)
				}
			}
			if !item.KeepaliveInterval.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborTimers.keepAliveInterval", item.KeepaliveInterval.ValueInt64())
			}
			if !item.HoldTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborTimers.holdTime", item.HoldTime.ValueInt64())
			}
			if !item.MinimumHoldTime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborTimers.minimumHoldTime", item.MinimumHoldTime.ValueInt64())
			}
			if !item.AuthenticationPassword.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborSecret", item.AuthenticationPassword.ValueString())
			}
			if !item.SendCommunityAttribute.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.sendCommunity", item.SendCommunityAttribute.ValueBool())
			}
			if !item.NextHopSelf.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.nextHopSelf", item.NextHopSelf.ValueBool())
			}
			if !item.DisableConnectionVerification.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborHops.disableConnectedCheck", item.DisableConnectionVerification.ValueBool())
			}
			if !item.TcpPathMtuDiscovery.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborTransportPathMTUDiscovery.disable", item.TcpPathMtuDiscovery.ValueBool())
			}
			if !item.MaxHopCount.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborHops.maxHopCount", item.MaxHopCount.ValueInt64())
			}
			if !item.TcpTransportMode.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborTransportConnectionMode.establishTCPSession", item.TcpTransportMode.ValueBool())
			}
			if !item.Weight.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborWeight", item.Weight.ValueInt64())
			}
			if !item.Version.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborAdvanced.neighborVersion", item.Version.ValueString())
			}
			if !item.CustomizedLocalAsNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborLocalAs.asNumber", item.CustomizedLocalAsNumber.ValueString())
			}
			if !item.CustomizedNoPrepend.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborLocalAs.noPrepend", item.CustomizedNoPrepend.ValueBool())
			}
			if !item.CustomizedReplaceAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborLocalAs.replaceAs", item.CustomizedReplaceAs.ValueBool())
			}
			if !item.CustomizedAcceptBothAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "neighborLocalAs.dualAs", item.CustomizedAcceptBothAs.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.neighbors.-1", itemBody)
		}
	}
	if len(data.Ipv4AggregateAddresses) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.aggregateAddressesIPv4s", []any{})
		for _, item := range data.Ipv4AggregateAddresses {
			itemBody := ""
			if !item.GenerateAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "asSet", item.GenerateAs.ValueBool())
			}
			if !item.Filter.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "summaryOnly", item.Filter.ValueBool())
			}
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4AggregateNetwork.id", item.NetworkId.ValueString())
			}
			if !item.AdvertiseMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "advertiseMap.id", item.AdvertiseMapId.ValueString())
			}
			if !item.AttributeMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "attributeMap.id", item.AttributeMapId.ValueString())
			}
			if !item.SuppressMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "suppressMap.id", item.SuppressMapId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.aggregateAddressesIPv4s.-1", itemBody)
		}
	}
	if len(data.Ipv4Filterings) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.distributeLists", []any{})
		for _, item := range data.Ipv4Filterings {
			itemBody := ""
			if !item.AccessListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "accessList.id", item.AccessListId.ValueString())
			}
			if !item.Direction.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Direction.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol.protocol", item.Protocol.ValueString())
			}
			if !item.ProcessId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol.processId", item.ProcessId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.distributeLists.-1", itemBody)
		}
	}
	if len(data.Ipv4Networks) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.networks", []any{})
		for _, item := range data.Ipv4Networks {
			itemBody := ""
			if !item.NetworkId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv4Address.id", item.NetworkId.ValueString())
			}
			if !item.RouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMap.id", item.RouteMapId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.networks.-1", itemBody)
		}
	}
	if len(data.Ipv4Redistributions) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.redistributeProtocols", []any{})
		for _, item := range data.Ipv4Redistributions {
			itemBody := ""
			if !item.SourceProtocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.SourceProtocol.ValueString())
			}
			if !item.RouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMap.id", item.RouteMapId.ValueString())
			}
			if !item.Metric.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeMetric.metricValue", item.Metric.ValueInt64())
			}
			if !item.ProcessId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "processId", item.ProcessId.ValueString())
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
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.redistributeProtocols.-1", itemBody)
		}
	}
	if len(data.Ipv4RouteInjections) > 0 {
		body, _ = sjson.Set(body, "addressFamilyIPv4.injectMaps", []any{})
		for _, item := range data.Ipv4RouteInjections {
			itemBody := ""
			if !item.InjectRouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "injectMap.id", item.InjectRouteMapId.ValueString())
			}
			if !item.ExistRouteMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "existMap.id", item.ExistRouteMapId.ValueString())
			}
			if !item.InheritAttributes.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "copyAttributes", item.InheritAttributes.ValueString())
			}
			body, _ = sjson.SetRaw(body, "addressFamilyIPv4.injectMaps.-1", itemBody)
		}
	}
	if !data.Ipv4ImportRouteTargets.IsNull() {
		var values []string
		data.Ipv4ImportRouteTargets.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "addressFamilyIPv4.routeImportExport.importRouteTargets", values)
	}
	if !data.Ipv4ExportRouteTargets.IsNull() {
		var values []string
		data.Ipv4ExportRouteTargets.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "addressFamilyIPv4.routeImportExport.exportRouteTargets", values)
	}
	if !data.Ipv4ImportGlobalVrfRouteMapId.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.routeImportExport.globalVrfImportRouteMap.id", data.Ipv4ImportGlobalVrfRouteMapId.ValueString())
	}
	if !data.Ipv4ExportGlobalVrfRouteMapId.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.routeImportExport.globalVrfExportRouteMap.id", data.Ipv4ExportGlobalVrfRouteMapId.ValueString())
	}
	if !data.Ipv4ImportUserVrfRouteMapId.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.routeImportExport.userVrfImportRouteMap.id", data.Ipv4ImportUserVrfRouteMapId.ValueString())
	}
	if !data.Ipv4ExportUserVrfRouteMapId.IsNull() {
		body, _ = sjson.Set(body, "addressFamilyIPv4.routeImportExport.userVrfExportRouteMap.id", data.Ipv4ExportUserVrfRouteMapId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceBGP) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
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
	if value := res.Get("addressFamilyIPv4.id"); value.Exists() {
		data.Ipv4AddressFamilyId = types.StringValue(value.String())
	} else {
		data.Ipv4AddressFamilyId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.type"); value.Exists() {
		data.Ipv4AddressFamilyType = types.StringValue(value.String())
	} else {
		data.Ipv4AddressFamilyType = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.aftableMap.id"); value.Exists() {
		data.Ipv4LearnedRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4LearnedRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.defaultInformationOrginate"); value.Exists() {
		data.Ipv4DefaultInformationOrginate = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DefaultInformationOrginate = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.autoSummary"); value.Exists() {
		data.Ipv4AutoSummary = types.BoolValue(value.Bool())
	} else {
		data.Ipv4AutoSummary = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpSupressInactive"); value.Exists() {
		data.Ipv4SuppressInactiveRoutes = types.BoolValue(value.Bool())
	} else {
		data.Ipv4SuppressInactiveRoutes = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.synchronization"); value.Exists() {
		data.Ipv4Synchronization = types.BoolValue(value.Bool())
	} else {
		data.Ipv4Synchronization = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpRedistributeInternal"); value.Exists() {
		data.Ipv4RedistributeIbgpIntoIgp = types.BoolValue(value.Bool())
	} else {
		data.Ipv4RedistributeIbgpIntoIgp = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.distance.externalDistance"); value.Exists() {
		data.Ipv4ExternalDistance = types.Int64Value(value.Int())
	} else {
		data.Ipv4ExternalDistance = types.Int64Value(20)
	}
	if value := res.Get("addressFamilyIPv4.distance.internalDistance"); value.Exists() {
		data.Ipv4InternalDistance = types.Int64Value(value.Int())
	} else {
		data.Ipv4InternalDistance = types.Int64Value(200)
	}
	if value := res.Get("addressFamilyIPv4.distance.localDistance"); value.Exists() {
		data.Ipv4LocalDistance = types.Int64Value(value.Int())
	} else {
		data.Ipv4LocalDistance = types.Int64Value(200)
	}
	if value := res.Get("addressFamilyIPv4.ibgp"); value.Exists() {
		data.Ipv4NumberOfIbgpPaths = types.Int64Value(value.Int())
	} else {
		data.Ipv4NumberOfIbgpPaths = types.Int64Value(1)
	}
	if value := res.Get("addressFamilyIPv4.ebgp"); value.Exists() {
		data.Ipv4NumberOfEbgpPaths = types.Int64Value(value.Int())
	} else {
		data.Ipv4NumberOfEbgpPaths = types.Int64Value(1)
	}
	if value := res.Get("addressFamilyIPv4.neighbors"); value.Exists() {
		data.Ipv4Neighbors = make([]DeviceBGPIpv4Neighbors, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4Neighbors{}
			if value := res.Get("ipv4Address"); value.Exists() {
				data.Address = types.StringValue(value.String())
			} else {
				data.Address = types.StringNull()
			}
			if value := res.Get("remoteAs"); value.Exists() {
				data.RemoteAs = types.StringValue(value.String())
			} else {
				data.RemoteAs = types.StringNull()
			}
			if value := res.Get("neighborGeneral.fallOverBFD"); value.Exists() {
				data.BfdFallover = types.StringValue(value.String())
			} else {
				data.BfdFallover = types.StringValue("NONE")
			}
			if value := res.Get("neighborGeneral.updateSource.id"); value.Exists() {
				data.UpdateSourceInterfaceId = types.StringValue(value.String())
			} else {
				data.UpdateSourceInterfaceId = types.StringNull()
			}
			if value := res.Get("neighborGeneral.enableAddress"); value.Exists() {
				data.EnableAddress = types.BoolValue(value.Bool())
			} else {
				data.EnableAddress = types.BoolValue(false)
			}
			if value := res.Get("neighborGeneral.asOverride"); value.Exists() {
				data.AsOverride = types.BoolValue(value.Bool())
			} else {
				data.AsOverride = types.BoolNull()
			}
			if value := res.Get("neighborHaMode.disable"); value.Exists() {
				data.GracefulRestart = types.BoolValue(value.Bool())
			} else {
				data.GracefulRestart = types.BoolNull()
			}
			if value := res.Get("neighborGeneral.shutdown"); value.Exists() {
				data.ShutdownAdministratively = types.BoolValue(value.Bool())
			} else {
				data.ShutdownAdministratively = types.BoolValue(false)
			}
			if value := res.Get("neighborGeneral.description"); value.Exists() {
				data.Description = types.StringValue(value.String())
			} else {
				data.Description = types.StringNull()
			}
			if value := res.Get("neighborFiltering.neighborDistributeLists"); value.Exists() {
				data.FilterAccessLists = make([]DeviceBGPIpv4NeighborsFilterAccessLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceBGPIpv4NeighborsFilterAccessLists{}
					if value := res.Get("accessList.id"); value.Exists() {
						data.AccessListId = types.StringValue(value.String())
					} else {
						data.AccessListId = types.StringNull()
					}
					if value := res.Get("filterUpdateAction"); value.Exists() {
						data.UpdateDirection = types.StringValue(value.String())
					} else {
						data.UpdateDirection = types.StringNull()
					}
					(*parent).FilterAccessLists = append((*parent).FilterAccessLists, data)
					return true
				})
			}
			if value := res.Get("neighborFiltering.neighborRouteMap"); value.Exists() {
				data.FilterRouteMaps = make([]DeviceBGPIpv4NeighborsFilterRouteMaps, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceBGPIpv4NeighborsFilterRouteMaps{}
					if value := res.Get("routeMap.id"); value.Exists() {
						data.RouteMapId = types.StringValue(value.String())
					} else {
						data.RouteMapId = types.StringNull()
					}
					if value := res.Get("filterUpdateAction"); value.Exists() {
						data.UpdateDirection = types.StringValue(value.String())
					} else {
						data.UpdateDirection = types.StringNull()
					}
					(*parent).FilterRouteMaps = append((*parent).FilterRouteMaps, data)
					return true
				})
			}
			if value := res.Get("neighborFiltering.ipv4PrefixListFilter"); value.Exists() {
				data.FilterPrefixLists = make([]DeviceBGPIpv4NeighborsFilterPrefixLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceBGPIpv4NeighborsFilterPrefixLists{}
					if value := res.Get("ipv4PrefixList.id"); value.Exists() {
						data.PrefixListId = types.StringValue(value.String())
					} else {
						data.PrefixListId = types.StringNull()
					}
					if value := res.Get("filterUpdateAction"); value.Exists() {
						data.UpdateDirection = types.StringValue(value.String())
					} else {
						data.UpdateDirection = types.StringNull()
					}
					(*parent).FilterPrefixLists = append((*parent).FilterPrefixLists, data)
					return true
				})
			}
			if value := res.Get("neighborFiltering.neighborFilterList"); value.Exists() {
				data.FilterAsPaths = make([]DeviceBGPIpv4NeighborsFilterAsPaths, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceBGPIpv4NeighborsFilterAsPaths{}
					if value := res.Get("asPathList.id"); value.Exists() {
						data.AsPathId = types.StringValue(value.String())
					} else {
						data.AsPathId = types.StringNull()
					}
					if value := res.Get("asPathList.name"); value.Exists() {
						data.AsPathName = types.StringValue(value.String())
					} else {
						data.AsPathName = types.StringNull()
					}
					if value := res.Get("filterUpdateAction"); value.Exists() {
						data.UpdateDirection = types.StringValue(value.String())
					} else {
						data.UpdateDirection = types.StringNull()
					}
					(*parent).FilterAsPaths = append((*parent).FilterAsPaths, data)
					return true
				})
			}
			if value := res.Get("neighborFiltering.neighborMaximumPrefix.maxPrefixLimit"); value.Exists() {
				data.FilterMaximumPrefixes = types.Int64Value(value.Int())
			} else {
				data.FilterMaximumPrefixes = types.Int64Null()
			}
			if value := res.Get("neighborFiltering.neighborMaximumPrefix.warningOnly"); value.Exists() {
				data.FilterWarningOnly = types.BoolValue(value.Bool())
			} else {
				data.FilterWarningOnly = types.BoolNull()
			}
			if value := res.Get("neighborFiltering.neighborMaximumPrefix.thresholdValue"); value.Exists() {
				data.FilterThresholdValue = types.Int64Value(value.Int())
			} else {
				data.FilterThresholdValue = types.Int64Null()
			}
			if value := res.Get("neighborFiltering.neighborMaximumPrefix.restartInterval"); value.Exists() {
				data.FilterRestartInterval = types.Int64Value(value.Int())
			} else {
				data.FilterRestartInterval = types.Int64Null()
			}
			if value := res.Get("neighborRoutes.advertisementInterval"); value.Exists() {
				data.RoutesAdvertisementInterval = types.Int64Value(value.Int())
			} else {
				data.RoutesAdvertisementInterval = types.Int64Value(0)
			}
			if value := res.Get("neighborRoutes.removePrivateAs"); value.Exists() {
				data.RoutesRemovePrivateAs = types.BoolValue(value.Bool())
			} else {
				data.RoutesRemovePrivateAs = types.BoolValue(false)
			}
			if value := res.Get("neighborFiltering.neighborDefaultOriginate.routeMap.id"); value.Exists() {
				data.RoutesGenerateDefaultRouteMapId = types.StringValue(value.String())
			} else {
				data.RoutesGenerateDefaultRouteMapId = types.StringNull()
			}
			if value := res.Get("neighborRoutes.neighborAdvertiseMaps"); value.Exists() {
				data.RoutesAdvertiseMaps = make([]DeviceBGPIpv4NeighborsRoutesAdvertiseMaps, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := DeviceBGPIpv4NeighborsRoutesAdvertiseMaps{}
					if value := res.Get("routeMap.id"); value.Exists() {
						data.AdvertiseMapId = types.StringValue(value.String())
					} else {
						data.AdvertiseMapId = types.StringNull()
					}
					if value := res.Get("existMap"); value.Exists() {
						data.UseExistMap = types.BoolValue(value.Bool())
					} else {
						data.UseExistMap = types.BoolNull()
					}
					if value := res.Get("existRouteMap.id"); value.Exists() {
						data.ExistNonexistMapId = types.StringValue(value.String())
					} else {
						data.ExistNonexistMapId = types.StringNull()
					}
					(*parent).RoutesAdvertiseMaps = append((*parent).RoutesAdvertiseMaps, data)
					return true
				})
			}
			if value := res.Get("neighborTimers.keepAliveInterval"); value.Exists() {
				data.KeepaliveInterval = types.Int64Value(value.Int())
			} else {
				data.KeepaliveInterval = types.Int64Null()
			}
			if value := res.Get("neighborTimers.holdTime"); value.Exists() {
				data.HoldTime = types.Int64Value(value.Int())
			} else {
				data.HoldTime = types.Int64Null()
			}
			if value := res.Get("neighborTimers.minimumHoldTime"); value.Exists() {
				data.MinimumHoldTime = types.Int64Value(value.Int())
			} else {
				data.MinimumHoldTime = types.Int64Null()
			}
			if value := res.Get("neighborAdvanced.neighborSecret"); value.Exists() {
				data.AuthenticationPassword = types.StringValue(value.String())
			} else {
				data.AuthenticationPassword = types.StringNull()
			}
			if value := res.Get("neighborAdvanced.sendCommunity"); value.Exists() {
				data.SendCommunityAttribute = types.BoolValue(value.Bool())
			} else {
				data.SendCommunityAttribute = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.nextHopSelf"); value.Exists() {
				data.NextHopSelf = types.BoolValue(value.Bool())
			} else {
				data.NextHopSelf = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.neighborHops.disableConnectedCheck"); value.Exists() {
				data.DisableConnectionVerification = types.BoolValue(value.Bool())
			} else {
				data.DisableConnectionVerification = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.neighborTransportPathMTUDiscovery.disable"); value.Exists() {
				data.TcpPathMtuDiscovery = types.BoolValue(value.Bool())
			} else {
				data.TcpPathMtuDiscovery = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.neighborHops.maxHopCount"); value.Exists() {
				data.MaxHopCount = types.Int64Value(value.Int())
			} else {
				data.MaxHopCount = types.Int64Value(1)
			}
			if value := res.Get("neighborAdvanced.neighborTransportConnectionMode.establishTCPSession"); value.Exists() {
				data.TcpTransportMode = types.BoolValue(value.Bool())
			} else {
				data.TcpTransportMode = types.BoolValue(false)
			}
			if value := res.Get("neighborAdvanced.neighborWeight"); value.Exists() {
				data.Weight = types.Int64Value(value.Int())
			} else {
				data.Weight = types.Int64Value(0)
			}
			if value := res.Get("neighborAdvanced.neighborVersion"); value.Exists() {
				data.Version = types.StringValue(value.String())
			} else {
				data.Version = types.StringValue("0")
			}
			if value := res.Get("neighborLocalAs.asNumber"); value.Exists() {
				data.CustomizedLocalAsNumber = types.StringValue(value.String())
			} else {
				data.CustomizedLocalAsNumber = types.StringNull()
			}
			if value := res.Get("neighborLocalAs.noPrepend"); value.Exists() {
				data.CustomizedNoPrepend = types.BoolValue(value.Bool())
			} else {
				data.CustomizedNoPrepend = types.BoolNull()
			}
			if value := res.Get("neighborLocalAs.replaceAs"); value.Exists() {
				data.CustomizedReplaceAs = types.BoolValue(value.Bool())
			} else {
				data.CustomizedReplaceAs = types.BoolNull()
			}
			if value := res.Get("neighborLocalAs.dualAs"); value.Exists() {
				data.CustomizedAcceptBothAs = types.BoolValue(value.Bool())
			} else {
				data.CustomizedAcceptBothAs = types.BoolNull()
			}
			(*parent).Ipv4Neighbors = append((*parent).Ipv4Neighbors, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.aggregateAddressesIPv4s"); value.Exists() {
		data.Ipv4AggregateAddresses = make([]DeviceBGPIpv4AggregateAddresses, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4AggregateAddresses{}
			if value := res.Get("asSet"); value.Exists() {
				data.GenerateAs = types.BoolValue(value.Bool())
			} else {
				data.GenerateAs = types.BoolNull()
			}
			if value := res.Get("summaryOnly"); value.Exists() {
				data.Filter = types.BoolValue(value.Bool())
			} else {
				data.Filter = types.BoolNull()
			}
			if value := res.Get("ipv4AggregateNetwork.id"); value.Exists() {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("advertiseMap.id"); value.Exists() {
				data.AdvertiseMapId = types.StringValue(value.String())
			} else {
				data.AdvertiseMapId = types.StringNull()
			}
			if value := res.Get("attributeMap.id"); value.Exists() {
				data.AttributeMapId = types.StringValue(value.String())
			} else {
				data.AttributeMapId = types.StringNull()
			}
			if value := res.Get("suppressMap.id"); value.Exists() {
				data.SuppressMapId = types.StringValue(value.String())
			} else {
				data.SuppressMapId = types.StringNull()
			}
			(*parent).Ipv4AggregateAddresses = append((*parent).Ipv4AggregateAddresses, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.distributeLists"); value.Exists() {
		data.Ipv4Filterings = make([]DeviceBGPIpv4Filterings, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4Filterings{}
			if value := res.Get("accessList.id"); value.Exists() {
				data.AccessListId = types.StringValue(value.String())
			} else {
				data.AccessListId = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() {
				data.Direction = types.StringValue(value.String())
			} else {
				data.Direction = types.StringNull()
			}
			if value := res.Get("protocol.protocol"); value.Exists() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("protocol.processId"); value.Exists() {
				data.ProcessId = types.StringValue(value.String())
			} else {
				data.ProcessId = types.StringNull()
			}
			(*parent).Ipv4Filterings = append((*parent).Ipv4Filterings, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.networks"); value.Exists() {
		data.Ipv4Networks = make([]DeviceBGPIpv4Networks, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4Networks{}
			if value := res.Get("ipv4Address.id"); value.Exists() {
				data.NetworkId = types.StringValue(value.String())
			} else {
				data.NetworkId = types.StringNull()
			}
			if value := res.Get("routeMap.id"); value.Exists() {
				data.RouteMapId = types.StringValue(value.String())
			} else {
				data.RouteMapId = types.StringNull()
			}
			(*parent).Ipv4Networks = append((*parent).Ipv4Networks, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.redistributeProtocols"); value.Exists() {
		data.Ipv4Redistributions = make([]DeviceBGPIpv4Redistributions, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4Redistributions{}
			if value := res.Get("type"); value.Exists() {
				data.SourceProtocol = types.StringValue(value.String())
			} else {
				data.SourceProtocol = types.StringNull()
			}
			if value := res.Get("routeMap.id"); value.Exists() {
				data.RouteMapId = types.StringValue(value.String())
			} else {
				data.RouteMapId = types.StringNull()
			}
			if value := res.Get("routeMetric.metricValue"); value.Exists() {
				data.Metric = types.Int64Value(value.Int())
			} else {
				data.Metric = types.Int64Null()
			}
			if value := res.Get("processId"); value.Exists() {
				data.ProcessId = types.StringValue(value.String())
			} else {
				data.ProcessId = types.StringNull()
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
			(*parent).Ipv4Redistributions = append((*parent).Ipv4Redistributions, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.injectMaps"); value.Exists() {
		data.Ipv4RouteInjections = make([]DeviceBGPIpv4RouteInjections, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBGPIpv4RouteInjections{}
			if value := res.Get("injectMap.id"); value.Exists() {
				data.InjectRouteMapId = types.StringValue(value.String())
			} else {
				data.InjectRouteMapId = types.StringNull()
			}
			if value := res.Get("existMap.id"); value.Exists() {
				data.ExistRouteMapId = types.StringValue(value.String())
			} else {
				data.ExistRouteMapId = types.StringNull()
			}
			if value := res.Get("copyAttributes"); value.Exists() {
				data.InheritAttributes = types.StringValue(value.String())
			} else {
				data.InheritAttributes = types.StringValue("true")
			}
			(*parent).Ipv4RouteInjections = append((*parent).Ipv4RouteInjections, data)
			return true
		})
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.importRouteTargets"); value.Exists() {
		data.Ipv4ImportRouteTargets = helpers.GetStringList(value.Array())
	} else {
		data.Ipv4ImportRouteTargets = types.ListNull(types.StringType)
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.exportRouteTargets"); value.Exists() {
		data.Ipv4ExportRouteTargets = helpers.GetStringList(value.Array())
	} else {
		data.Ipv4ExportRouteTargets = types.ListNull(types.StringType)
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.globalVrfImportRouteMap.id"); value.Exists() {
		data.Ipv4ImportGlobalVrfRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4ImportGlobalVrfRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.globalVrfExportRouteMap.id"); value.Exists() {
		data.Ipv4ExportGlobalVrfRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4ExportGlobalVrfRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.userVrfImportRouteMap.id"); value.Exists() {
		data.Ipv4ImportUserVrfRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4ImportUserVrfRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.userVrfExportRouteMap.id"); value.Exists() {
		data.Ipv4ExportUserVrfRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4ExportUserVrfRouteMapId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceBGP) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
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
	if value := res.Get("addressFamilyIPv4.id"); value.Exists() && !data.Ipv4AddressFamilyId.IsNull() {
		data.Ipv4AddressFamilyId = types.StringValue(value.String())
	} else {
		data.Ipv4AddressFamilyId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.type"); value.Exists() && !data.Ipv4AddressFamilyType.IsNull() {
		data.Ipv4AddressFamilyType = types.StringValue(value.String())
	} else {
		data.Ipv4AddressFamilyType = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.aftableMap.id"); value.Exists() && !data.Ipv4LearnedRouteMapId.IsNull() {
		data.Ipv4LearnedRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4LearnedRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.defaultInformationOrginate"); value.Exists() && !data.Ipv4DefaultInformationOrginate.IsNull() {
		data.Ipv4DefaultInformationOrginate = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DefaultInformationOrginate = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.autoSummary"); value.Exists() && !data.Ipv4AutoSummary.IsNull() {
		data.Ipv4AutoSummary = types.BoolValue(value.Bool())
	} else {
		data.Ipv4AutoSummary = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpSupressInactive"); value.Exists() && !data.Ipv4SuppressInactiveRoutes.IsNull() {
		data.Ipv4SuppressInactiveRoutes = types.BoolValue(value.Bool())
	} else {
		data.Ipv4SuppressInactiveRoutes = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.synchronization"); value.Exists() && !data.Ipv4Synchronization.IsNull() {
		data.Ipv4Synchronization = types.BoolValue(value.Bool())
	} else {
		data.Ipv4Synchronization = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.bgpRedistributeInternal"); value.Exists() && !data.Ipv4RedistributeIbgpIntoIgp.IsNull() {
		data.Ipv4RedistributeIbgpIntoIgp = types.BoolValue(value.Bool())
	} else {
		data.Ipv4RedistributeIbgpIntoIgp = types.BoolNull()
	}
	if value := res.Get("addressFamilyIPv4.distance.externalDistance"); value.Exists() && !data.Ipv4ExternalDistance.IsNull() {
		data.Ipv4ExternalDistance = types.Int64Value(value.Int())
	} else if data.Ipv4ExternalDistance.ValueInt64() != 20 {
		data.Ipv4ExternalDistance = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.distance.internalDistance"); value.Exists() && !data.Ipv4InternalDistance.IsNull() {
		data.Ipv4InternalDistance = types.Int64Value(value.Int())
	} else if data.Ipv4InternalDistance.ValueInt64() != 200 {
		data.Ipv4InternalDistance = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.distance.localDistance"); value.Exists() && !data.Ipv4LocalDistance.IsNull() {
		data.Ipv4LocalDistance = types.Int64Value(value.Int())
	} else if data.Ipv4LocalDistance.ValueInt64() != 200 {
		data.Ipv4LocalDistance = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.ibgp"); value.Exists() && !data.Ipv4NumberOfIbgpPaths.IsNull() {
		data.Ipv4NumberOfIbgpPaths = types.Int64Value(value.Int())
	} else if data.Ipv4NumberOfIbgpPaths.ValueInt64() != 1 {
		data.Ipv4NumberOfIbgpPaths = types.Int64Null()
	}
	if value := res.Get("addressFamilyIPv4.ebgp"); value.Exists() && !data.Ipv4NumberOfEbgpPaths.IsNull() {
		data.Ipv4NumberOfEbgpPaths = types.Int64Value(value.Int())
	} else if data.Ipv4NumberOfEbgpPaths.ValueInt64() != 1 {
		data.Ipv4NumberOfEbgpPaths = types.Int64Null()
	}
	for i := 0; i < len(data.Ipv4Neighbors); i++ {
		keys := [...]string{"ipv4Address"}
		keyValues := [...]string{data.Ipv4Neighbors[i].Address.ValueString()}

		parent := &data
		data := (*parent).Ipv4Neighbors[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.neighbors").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4Neighbors[%d] = %+v",
				i,
				(*parent).Ipv4Neighbors[i],
			))
			(*parent).Ipv4Neighbors = slices.Delete((*parent).Ipv4Neighbors, i, i+1)
			i--

			continue
		}
		if value := res.Get("ipv4Address"); value.Exists() && !data.Address.IsNull() {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("remoteAs"); value.Exists() && !data.RemoteAs.IsNull() {
			data.RemoteAs = types.StringValue(value.String())
		} else {
			data.RemoteAs = types.StringNull()
		}
		if value := res.Get("neighborGeneral.fallOverBFD"); value.Exists() && !data.BfdFallover.IsNull() {
			data.BfdFallover = types.StringValue(value.String())
		} else if data.BfdFallover.ValueString() != "NONE" {
			data.BfdFallover = types.StringNull()
		}
		if value := res.Get("neighborGeneral.updateSource.id"); value.Exists() && !data.UpdateSourceInterfaceId.IsNull() {
			data.UpdateSourceInterfaceId = types.StringValue(value.String())
		} else {
			data.UpdateSourceInterfaceId = types.StringNull()
		}
		if value := res.Get("neighborGeneral.enableAddress"); value.Exists() && !data.EnableAddress.IsNull() {
			data.EnableAddress = types.BoolValue(value.Bool())
		} else if data.EnableAddress.ValueBool() != false {
			data.EnableAddress = types.BoolNull()
		}
		if value := res.Get("neighborGeneral.asOverride"); value.Exists() && !data.AsOverride.IsNull() {
			data.AsOverride = types.BoolValue(value.Bool())
		} else {
			data.AsOverride = types.BoolNull()
		}
		if value := res.Get("neighborHaMode.disable"); value.Exists() && !data.GracefulRestart.IsNull() {
			data.GracefulRestart = types.BoolValue(value.Bool())
		} else {
			data.GracefulRestart = types.BoolNull()
		}
		if value := res.Get("neighborGeneral.shutdown"); value.Exists() && !data.ShutdownAdministratively.IsNull() {
			data.ShutdownAdministratively = types.BoolValue(value.Bool())
		} else if data.ShutdownAdministratively.ValueBool() != false {
			data.ShutdownAdministratively = types.BoolNull()
		}
		if value := res.Get("neighborGeneral.description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		for i := 0; i < len(data.FilterAccessLists); i++ {
			keys := [...]string{"accessList.id", "filterUpdateAction"}
			keyValues := [...]string{data.FilterAccessLists[i].AccessListId.ValueString(), data.FilterAccessLists[i].UpdateDirection.ValueString()}

			parent := &data
			data := (*parent).FilterAccessLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("neighborFiltering.neighborDistributeLists").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing FilterAccessLists[%d] = %+v",
					i,
					(*parent).FilterAccessLists[i],
				))
				(*parent).FilterAccessLists = slices.Delete((*parent).FilterAccessLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("accessList.id"); value.Exists() && !data.AccessListId.IsNull() {
				data.AccessListId = types.StringValue(value.String())
			} else {
				data.AccessListId = types.StringNull()
			}
			if value := res.Get("filterUpdateAction"); value.Exists() && !data.UpdateDirection.IsNull() {
				data.UpdateDirection = types.StringValue(value.String())
			} else {
				data.UpdateDirection = types.StringNull()
			}
			(*parent).FilterAccessLists[i] = data
		}
		for i := 0; i < len(data.FilterRouteMaps); i++ {
			keys := [...]string{"routeMap.id", "filterUpdateAction"}
			keyValues := [...]string{data.FilterRouteMaps[i].RouteMapId.ValueString(), data.FilterRouteMaps[i].UpdateDirection.ValueString()}

			parent := &data
			data := (*parent).FilterRouteMaps[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("neighborFiltering.neighborRouteMap").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing FilterRouteMaps[%d] = %+v",
					i,
					(*parent).FilterRouteMaps[i],
				))
				(*parent).FilterRouteMaps = slices.Delete((*parent).FilterRouteMaps, i, i+1)
				i--

				continue
			}
			if value := res.Get("routeMap.id"); value.Exists() && !data.RouteMapId.IsNull() {
				data.RouteMapId = types.StringValue(value.String())
			} else {
				data.RouteMapId = types.StringNull()
			}
			if value := res.Get("filterUpdateAction"); value.Exists() && !data.UpdateDirection.IsNull() {
				data.UpdateDirection = types.StringValue(value.String())
			} else {
				data.UpdateDirection = types.StringNull()
			}
			(*parent).FilterRouteMaps[i] = data
		}
		for i := 0; i < len(data.FilterPrefixLists); i++ {
			keys := [...]string{"ipv4PrefixList.id", "filterUpdateAction"}
			keyValues := [...]string{data.FilterPrefixLists[i].PrefixListId.ValueString(), data.FilterPrefixLists[i].UpdateDirection.ValueString()}

			parent := &data
			data := (*parent).FilterPrefixLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("neighborFiltering.ipv4PrefixListFilter").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing FilterPrefixLists[%d] = %+v",
					i,
					(*parent).FilterPrefixLists[i],
				))
				(*parent).FilterPrefixLists = slices.Delete((*parent).FilterPrefixLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("ipv4PrefixList.id"); value.Exists() && !data.PrefixListId.IsNull() {
				data.PrefixListId = types.StringValue(value.String())
			} else {
				data.PrefixListId = types.StringNull()
			}
			if value := res.Get("filterUpdateAction"); value.Exists() && !data.UpdateDirection.IsNull() {
				data.UpdateDirection = types.StringValue(value.String())
			} else {
				data.UpdateDirection = types.StringNull()
			}
			(*parent).FilterPrefixLists[i] = data
		}
		for i := 0; i < len(data.FilterAsPaths); i++ {
			keys := [...]string{"asPathList.id", "asPathList.name", "filterUpdateAction"}
			keyValues := [...]string{data.FilterAsPaths[i].AsPathId.ValueString(), data.FilterAsPaths[i].AsPathName.ValueString(), data.FilterAsPaths[i].UpdateDirection.ValueString()}

			parent := &data
			data := (*parent).FilterAsPaths[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("neighborFiltering.neighborFilterList").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing FilterAsPaths[%d] = %+v",
					i,
					(*parent).FilterAsPaths[i],
				))
				(*parent).FilterAsPaths = slices.Delete((*parent).FilterAsPaths, i, i+1)
				i--

				continue
			}
			if value := res.Get("asPathList.id"); value.Exists() && !data.AsPathId.IsNull() {
				data.AsPathId = types.StringValue(value.String())
			} else {
				data.AsPathId = types.StringNull()
			}
			if value := res.Get("asPathList.name"); value.Exists() && !data.AsPathName.IsNull() {
				data.AsPathName = types.StringValue(value.String())
			} else {
				data.AsPathName = types.StringNull()
			}
			if value := res.Get("filterUpdateAction"); value.Exists() && !data.UpdateDirection.IsNull() {
				data.UpdateDirection = types.StringValue(value.String())
			} else {
				data.UpdateDirection = types.StringNull()
			}
			(*parent).FilterAsPaths[i] = data
		}
		if value := res.Get("neighborFiltering.neighborMaximumPrefix.maxPrefixLimit"); value.Exists() && !data.FilterMaximumPrefixes.IsNull() {
			data.FilterMaximumPrefixes = types.Int64Value(value.Int())
		} else {
			data.FilterMaximumPrefixes = types.Int64Null()
		}
		if value := res.Get("neighborFiltering.neighborMaximumPrefix.warningOnly"); value.Exists() && !data.FilterWarningOnly.IsNull() {
			data.FilterWarningOnly = types.BoolValue(value.Bool())
		} else {
			data.FilterWarningOnly = types.BoolNull()
		}
		if value := res.Get("neighborFiltering.neighborMaximumPrefix.thresholdValue"); value.Exists() && !data.FilterThresholdValue.IsNull() {
			data.FilterThresholdValue = types.Int64Value(value.Int())
		} else {
			data.FilterThresholdValue = types.Int64Null()
		}
		if value := res.Get("neighborFiltering.neighborMaximumPrefix.restartInterval"); value.Exists() && !data.FilterRestartInterval.IsNull() {
			data.FilterRestartInterval = types.Int64Value(value.Int())
		} else {
			data.FilterRestartInterval = types.Int64Null()
		}
		if value := res.Get("neighborRoutes.advertisementInterval"); value.Exists() && !data.RoutesAdvertisementInterval.IsNull() {
			data.RoutesAdvertisementInterval = types.Int64Value(value.Int())
		} else if data.RoutesAdvertisementInterval.ValueInt64() != 0 {
			data.RoutesAdvertisementInterval = types.Int64Null()
		}
		if value := res.Get("neighborRoutes.removePrivateAs"); value.Exists() && !data.RoutesRemovePrivateAs.IsNull() {
			data.RoutesRemovePrivateAs = types.BoolValue(value.Bool())
		} else if data.RoutesRemovePrivateAs.ValueBool() != false {
			data.RoutesRemovePrivateAs = types.BoolNull()
		}
		if value := res.Get("neighborFiltering.neighborDefaultOriginate.routeMap.id"); value.Exists() && !data.RoutesGenerateDefaultRouteMapId.IsNull() {
			data.RoutesGenerateDefaultRouteMapId = types.StringValue(value.String())
		} else {
			data.RoutesGenerateDefaultRouteMapId = types.StringNull()
		}
		for i := 0; i < len(data.RoutesAdvertiseMaps); i++ {
			keys := [...]string{"routeMap.id", "existRouteMap.id"}
			keyValues := [...]string{data.RoutesAdvertiseMaps[i].AdvertiseMapId.ValueString(), data.RoutesAdvertiseMaps[i].ExistNonexistMapId.ValueString()}

			parent := &data
			data := (*parent).RoutesAdvertiseMaps[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("neighborRoutes.neighborAdvertiseMaps").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing RoutesAdvertiseMaps[%d] = %+v",
					i,
					(*parent).RoutesAdvertiseMaps[i],
				))
				(*parent).RoutesAdvertiseMaps = slices.Delete((*parent).RoutesAdvertiseMaps, i, i+1)
				i--

				continue
			}
			if value := res.Get("routeMap.id"); value.Exists() && !data.AdvertiseMapId.IsNull() {
				data.AdvertiseMapId = types.StringValue(value.String())
			} else {
				data.AdvertiseMapId = types.StringNull()
			}
			if value := res.Get("existMap"); value.Exists() && !data.UseExistMap.IsNull() {
				data.UseExistMap = types.BoolValue(value.Bool())
			} else {
				data.UseExistMap = types.BoolNull()
			}
			if value := res.Get("existRouteMap.id"); value.Exists() && !data.ExistNonexistMapId.IsNull() {
				data.ExistNonexistMapId = types.StringValue(value.String())
			} else {
				data.ExistNonexistMapId = types.StringNull()
			}
			(*parent).RoutesAdvertiseMaps[i] = data
		}
		if value := res.Get("neighborTimers.keepAliveInterval"); value.Exists() && !data.KeepaliveInterval.IsNull() {
			data.KeepaliveInterval = types.Int64Value(value.Int())
		} else {
			data.KeepaliveInterval = types.Int64Null()
		}
		if value := res.Get("neighborTimers.holdTime"); value.Exists() && !data.HoldTime.IsNull() {
			data.HoldTime = types.Int64Value(value.Int())
		} else {
			data.HoldTime = types.Int64Null()
		}
		if value := res.Get("neighborTimers.minimumHoldTime"); value.Exists() && !data.MinimumHoldTime.IsNull() {
			data.MinimumHoldTime = types.Int64Value(value.Int())
		} else {
			data.MinimumHoldTime = types.Int64Null()
		}
		if value := res.Get("neighborAdvanced.neighborSecret"); value.Exists() && !data.AuthenticationPassword.IsNull() {
			data.AuthenticationPassword = types.StringValue(value.String())
		} else {
			data.AuthenticationPassword = types.StringNull()
		}
		if value := res.Get("neighborAdvanced.sendCommunity"); value.Exists() && !data.SendCommunityAttribute.IsNull() {
			data.SendCommunityAttribute = types.BoolValue(value.Bool())
		} else if data.SendCommunityAttribute.ValueBool() != false {
			data.SendCommunityAttribute = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.nextHopSelf"); value.Exists() && !data.NextHopSelf.IsNull() {
			data.NextHopSelf = types.BoolValue(value.Bool())
		} else if data.NextHopSelf.ValueBool() != false {
			data.NextHopSelf = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.neighborHops.disableConnectedCheck"); value.Exists() && !data.DisableConnectionVerification.IsNull() {
			data.DisableConnectionVerification = types.BoolValue(value.Bool())
		} else if data.DisableConnectionVerification.ValueBool() != false {
			data.DisableConnectionVerification = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.neighborTransportPathMTUDiscovery.disable"); value.Exists() && !data.TcpPathMtuDiscovery.IsNull() {
			data.TcpPathMtuDiscovery = types.BoolValue(value.Bool())
		} else if data.TcpPathMtuDiscovery.ValueBool() != false {
			data.TcpPathMtuDiscovery = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.neighborHops.maxHopCount"); value.Exists() && !data.MaxHopCount.IsNull() {
			data.MaxHopCount = types.Int64Value(value.Int())
		} else if data.MaxHopCount.ValueInt64() != 1 {
			data.MaxHopCount = types.Int64Null()
		}
		if value := res.Get("neighborAdvanced.neighborTransportConnectionMode.establishTCPSession"); value.Exists() && !data.TcpTransportMode.IsNull() {
			data.TcpTransportMode = types.BoolValue(value.Bool())
		} else if data.TcpTransportMode.ValueBool() != false {
			data.TcpTransportMode = types.BoolNull()
		}
		if value := res.Get("neighborAdvanced.neighborWeight"); value.Exists() && !data.Weight.IsNull() {
			data.Weight = types.Int64Value(value.Int())
		} else if data.Weight.ValueInt64() != 0 {
			data.Weight = types.Int64Null()
		}
		if value := res.Get("neighborAdvanced.neighborVersion"); value.Exists() && !data.Version.IsNull() {
			data.Version = types.StringValue(value.String())
		} else if data.Version.ValueString() != "0" {
			data.Version = types.StringNull()
		}
		if value := res.Get("neighborLocalAs.asNumber"); value.Exists() && !data.CustomizedLocalAsNumber.IsNull() {
			data.CustomizedLocalAsNumber = types.StringValue(value.String())
		} else {
			data.CustomizedLocalAsNumber = types.StringNull()
		}
		if value := res.Get("neighborLocalAs.noPrepend"); value.Exists() && !data.CustomizedNoPrepend.IsNull() {
			data.CustomizedNoPrepend = types.BoolValue(value.Bool())
		} else {
			data.CustomizedNoPrepend = types.BoolNull()
		}
		if value := res.Get("neighborLocalAs.replaceAs"); value.Exists() && !data.CustomizedReplaceAs.IsNull() {
			data.CustomizedReplaceAs = types.BoolValue(value.Bool())
		} else {
			data.CustomizedReplaceAs = types.BoolNull()
		}
		if value := res.Get("neighborLocalAs.dualAs"); value.Exists() && !data.CustomizedAcceptBothAs.IsNull() {
			data.CustomizedAcceptBothAs = types.BoolValue(value.Bool())
		} else {
			data.CustomizedAcceptBothAs = types.BoolNull()
		}
		(*parent).Ipv4Neighbors[i] = data
	}
	for i := 0; i < len(data.Ipv4AggregateAddresses); i++ {
		keys := [...]string{"ipv4AggregateNetwork.id", "advertiseMap.id", "attributeMap.id", "suppressMap.id"}
		keyValues := [...]string{data.Ipv4AggregateAddresses[i].NetworkId.ValueString(), data.Ipv4AggregateAddresses[i].AdvertiseMapId.ValueString(), data.Ipv4AggregateAddresses[i].AttributeMapId.ValueString(), data.Ipv4AggregateAddresses[i].SuppressMapId.ValueString()}

		parent := &data
		data := (*parent).Ipv4AggregateAddresses[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.aggregateAddressesIPv4s").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4AggregateAddresses[%d] = %+v",
				i,
				(*parent).Ipv4AggregateAddresses[i],
			))
			(*parent).Ipv4AggregateAddresses = slices.Delete((*parent).Ipv4AggregateAddresses, i, i+1)
			i--

			continue
		}
		if value := res.Get("asSet"); value.Exists() && !data.GenerateAs.IsNull() {
			data.GenerateAs = types.BoolValue(value.Bool())
		} else {
			data.GenerateAs = types.BoolNull()
		}
		if value := res.Get("summaryOnly"); value.Exists() && !data.Filter.IsNull() {
			data.Filter = types.BoolValue(value.Bool())
		} else {
			data.Filter = types.BoolNull()
		}
		if value := res.Get("ipv4AggregateNetwork.id"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("advertiseMap.id"); value.Exists() && !data.AdvertiseMapId.IsNull() {
			data.AdvertiseMapId = types.StringValue(value.String())
		} else {
			data.AdvertiseMapId = types.StringNull()
		}
		if value := res.Get("attributeMap.id"); value.Exists() && !data.AttributeMapId.IsNull() {
			data.AttributeMapId = types.StringValue(value.String())
		} else {
			data.AttributeMapId = types.StringNull()
		}
		if value := res.Get("suppressMap.id"); value.Exists() && !data.SuppressMapId.IsNull() {
			data.SuppressMapId = types.StringValue(value.String())
		} else {
			data.SuppressMapId = types.StringNull()
		}
		(*parent).Ipv4AggregateAddresses[i] = data
	}
	for i := 0; i < len(data.Ipv4Filterings); i++ {
		keys := [...]string{"accessList.id", "type", "protocol.protocol", "protocol.processId"}
		keyValues := [...]string{data.Ipv4Filterings[i].AccessListId.ValueString(), data.Ipv4Filterings[i].Direction.ValueString(), data.Ipv4Filterings[i].Protocol.ValueString(), data.Ipv4Filterings[i].ProcessId.ValueString()}

		parent := &data
		data := (*parent).Ipv4Filterings[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.distributeLists").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4Filterings[%d] = %+v",
				i,
				(*parent).Ipv4Filterings[i],
			))
			(*parent).Ipv4Filterings = slices.Delete((*parent).Ipv4Filterings, i, i+1)
			i--

			continue
		}
		if value := res.Get("accessList.id"); value.Exists() && !data.AccessListId.IsNull() {
			data.AccessListId = types.StringValue(value.String())
		} else {
			data.AccessListId = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Direction.IsNull() {
			data.Direction = types.StringValue(value.String())
		} else {
			data.Direction = types.StringNull()
		}
		if value := res.Get("protocol.protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("protocol.processId"); value.Exists() && !data.ProcessId.IsNull() {
			data.ProcessId = types.StringValue(value.String())
		} else {
			data.ProcessId = types.StringNull()
		}
		(*parent).Ipv4Filterings[i] = data
	}
	for i := 0; i < len(data.Ipv4Networks); i++ {
		keys := [...]string{"ipv4Address.id", "routeMap.id"}
		keyValues := [...]string{data.Ipv4Networks[i].NetworkId.ValueString(), data.Ipv4Networks[i].RouteMapId.ValueString()}

		parent := &data
		data := (*parent).Ipv4Networks[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.networks").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4Networks[%d] = %+v",
				i,
				(*parent).Ipv4Networks[i],
			))
			(*parent).Ipv4Networks = slices.Delete((*parent).Ipv4Networks, i, i+1)
			i--

			continue
		}
		if value := res.Get("ipv4Address.id"); value.Exists() && !data.NetworkId.IsNull() {
			data.NetworkId = types.StringValue(value.String())
		} else {
			data.NetworkId = types.StringNull()
		}
		if value := res.Get("routeMap.id"); value.Exists() && !data.RouteMapId.IsNull() {
			data.RouteMapId = types.StringValue(value.String())
		} else {
			data.RouteMapId = types.StringNull()
		}
		(*parent).Ipv4Networks[i] = data
	}
	for i := 0; i < len(data.Ipv4Redistributions); i++ {
		keys := [...]string{"type", "processId"}
		keyValues := [...]string{data.Ipv4Redistributions[i].SourceProtocol.ValueString(), data.Ipv4Redistributions[i].ProcessId.ValueString()}

		parent := &data
		data := (*parent).Ipv4Redistributions[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.redistributeProtocols").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4Redistributions[%d] = %+v",
				i,
				(*parent).Ipv4Redistributions[i],
			))
			(*parent).Ipv4Redistributions = slices.Delete((*parent).Ipv4Redistributions, i, i+1)
			i--

			continue
		}
		if value := res.Get("type"); value.Exists() && !data.SourceProtocol.IsNull() {
			data.SourceProtocol = types.StringValue(value.String())
		} else {
			data.SourceProtocol = types.StringNull()
		}
		if value := res.Get("routeMap.id"); value.Exists() && !data.RouteMapId.IsNull() {
			data.RouteMapId = types.StringValue(value.String())
		} else {
			data.RouteMapId = types.StringNull()
		}
		if value := res.Get("routeMetric.metricValue"); value.Exists() && !data.Metric.IsNull() {
			data.Metric = types.Int64Value(value.Int())
		} else {
			data.Metric = types.Int64Null()
		}
		if value := res.Get("processId"); value.Exists() && !data.ProcessId.IsNull() {
			data.ProcessId = types.StringValue(value.String())
		} else {
			data.ProcessId = types.StringNull()
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
		(*parent).Ipv4Redistributions[i] = data
	}
	for i := 0; i < len(data.Ipv4RouteInjections); i++ {
		keys := [...]string{"injectMap.id"}
		keyValues := [...]string{data.Ipv4RouteInjections[i].InjectRouteMapId.ValueString()}

		parent := &data
		data := (*parent).Ipv4RouteInjections[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("addressFamilyIPv4.injectMaps").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4RouteInjections[%d] = %+v",
				i,
				(*parent).Ipv4RouteInjections[i],
			))
			(*parent).Ipv4RouteInjections = slices.Delete((*parent).Ipv4RouteInjections, i, i+1)
			i--

			continue
		}
		if value := res.Get("injectMap.id"); value.Exists() && !data.InjectRouteMapId.IsNull() {
			data.InjectRouteMapId = types.StringValue(value.String())
		} else {
			data.InjectRouteMapId = types.StringNull()
		}
		if value := res.Get("existMap.id"); value.Exists() && !data.ExistRouteMapId.IsNull() {
			data.ExistRouteMapId = types.StringValue(value.String())
		} else {
			data.ExistRouteMapId = types.StringNull()
		}
		if value := res.Get("copyAttributes"); value.Exists() && !data.InheritAttributes.IsNull() {
			data.InheritAttributes = types.StringValue(value.String())
		} else if data.InheritAttributes.ValueString() != "true" {
			data.InheritAttributes = types.StringNull()
		}
		(*parent).Ipv4RouteInjections[i] = data
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.importRouteTargets"); value.Exists() && !data.Ipv4ImportRouteTargets.IsNull() {
		data.Ipv4ImportRouteTargets = helpers.GetStringList(value.Array())
	} else {
		data.Ipv4ImportRouteTargets = types.ListNull(types.StringType)
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.exportRouteTargets"); value.Exists() && !data.Ipv4ExportRouteTargets.IsNull() {
		data.Ipv4ExportRouteTargets = helpers.GetStringList(value.Array())
	} else {
		data.Ipv4ExportRouteTargets = types.ListNull(types.StringType)
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.globalVrfImportRouteMap.id"); value.Exists() && !data.Ipv4ImportGlobalVrfRouteMapId.IsNull() {
		data.Ipv4ImportGlobalVrfRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4ImportGlobalVrfRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.globalVrfExportRouteMap.id"); value.Exists() && !data.Ipv4ExportGlobalVrfRouteMapId.IsNull() {
		data.Ipv4ExportGlobalVrfRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4ExportGlobalVrfRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.userVrfImportRouteMap.id"); value.Exists() && !data.Ipv4ImportUserVrfRouteMapId.IsNull() {
		data.Ipv4ImportUserVrfRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4ImportUserVrfRouteMapId = types.StringNull()
	}
	if value := res.Get("addressFamilyIPv4.routeImportExport.userVrfExportRouteMap.id"); value.Exists() && !data.Ipv4ExportUserVrfRouteMapId.IsNull() {
		data.Ipv4ExportUserVrfRouteMapId = types.StringValue(value.String())
	} else {
		data.Ipv4ExportUserVrfRouteMapId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceBGP) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Name.IsUnknown() {
		if value := res.Get("name"); value.Exists() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
	}
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.AsNumber.IsUnknown() {
		if value := res.Get("asNumber"); value.Exists() {
			data.AsNumber = types.StringValue(value.String())
		} else {
			data.AsNumber = types.StringNull()
		}
	}
	if data.Ipv4AddressFamilyId.IsUnknown() {
		if value := res.Get("addressFamilyIPv4.id"); value.Exists() {
			data.Ipv4AddressFamilyId = types.StringValue(value.String())
		} else {
			data.Ipv4AddressFamilyId = types.StringNull()
		}
	}
	if data.Ipv4AddressFamilyType.IsUnknown() {
		if value := res.Get("addressFamilyIPv4.type"); value.Exists() {
			data.Ipv4AddressFamilyType = types.StringValue(value.String())
		} else {
			data.Ipv4AddressFamilyType = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// End of section. //template:end toBodyNonBulk
