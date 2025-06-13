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
	"slices"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type RouteMap struct {
	Id          types.String      `tfsdk:"id"`
	Domain      types.String      `tfsdk:"domain"`
	Name        types.String      `tfsdk:"name"`
	Type        types.String      `tfsdk:"type"`
	Overridable types.Bool        `tfsdk:"overridable"`
	Entries     []RouteMapEntries `tfsdk:"entries"`
}

type RouteMapEntries struct {
	Action                                          types.String                                     `tfsdk:"action"`
	MatchSecurityZones                              []RouteMapEntriesMatchSecurityZones              `tfsdk:"match_security_zones"`
	MatchInterfaceNames                             types.List                                       `tfsdk:"match_interface_names"`
	MatchIpv4AddressAccessLists                     []RouteMapEntriesMatchIpv4AddressAccessLists     `tfsdk:"match_ipv4_address_access_lists"`
	MatchIpv4AddressPrefixLists                     []RouteMapEntriesMatchIpv4AddressPrefixLists     `tfsdk:"match_ipv4_address_prefix_lists"`
	MatchIpv4NextHopAccessLists                     []RouteMapEntriesMatchIpv4NextHopAccessLists     `tfsdk:"match_ipv4_next_hop_access_lists"`
	MatchIpv4NextHopPrefixLists                     []RouteMapEntriesMatchIpv4NextHopPrefixLists     `tfsdk:"match_ipv4_next_hop_prefix_lists"`
	MatchIpv4RouteSourceAccessLists                 []RouteMapEntriesMatchIpv4RouteSourceAccessLists `tfsdk:"match_ipv4_route_source_access_lists"`
	MatchIpv4RouteSourcePrefixLists                 []RouteMapEntriesMatchIpv4RouteSourcePrefixLists `tfsdk:"match_ipv4_route_source_prefix_lists"`
	MatchIpv6AddressExtendedAccessListId            types.String                                     `tfsdk:"match_ipv6_address_extended_access_list_id"`
	MatchIpv6AddressPrefixListId                    types.String                                     `tfsdk:"match_ipv6_address_prefix_list_id"`
	MatchIpv6NextHopExtendedAccessListId            types.String                                     `tfsdk:"match_ipv6_next_hop_extended_access_list_id"`
	MatchIpv6NextHopPrefixListId                    types.String                                     `tfsdk:"match_ipv6_next_hop_prefix_list_id"`
	MatchIpv6RouteSourceExtendedAccessListId        types.String                                     `tfsdk:"match_ipv6_route_source_extended_access_list_id"`
	MatchIpv6RouteSourcePrefixListId                types.String                                     `tfsdk:"match_ipv6_route_source_prefix_list_id"`
	MatchBgpAsPathLists                             []RouteMapEntriesMatchBgpAsPathLists             `tfsdk:"match_bgp_as_path_lists"`
	MatchBgpCommunityLists                          []RouteMapEntriesMatchBgpCommunityLists          `tfsdk:"match_bgp_community_lists"`
	MatchBgpExtendedCommunityLists                  []RouteMapEntriesMatchBgpExtendedCommunityLists  `tfsdk:"match_bgp_extended_community_lists"`
	MatchBgpPolicyLists                             []RouteMapEntriesMatchBgpPolicyLists             `tfsdk:"match_bgp_policy_lists"`
	MatchMetricRouteValues                          types.List                                       `tfsdk:"match_metric_route_values"`
	MatchTagValues                                  types.List                                       `tfsdk:"match_tag_values"`
	MatchRouteTypeExternal1                         types.Bool                                       `tfsdk:"match_route_type_external_1"`
	MatchRouteTypeExternal2                         types.Bool                                       `tfsdk:"match_route_type_external_2"`
	MatchRouteTypeInternal                          types.Bool                                       `tfsdk:"match_route_type_internal"`
	MatchRouteTypeLocal                             types.Bool                                       `tfsdk:"match_route_type_local"`
	MatchRouteTypeNssaExternal1                     types.Bool                                       `tfsdk:"match_route_type_nssa_external_1"`
	MatchRouteTypeNssaExternal2                     types.Bool                                       `tfsdk:"match_route_type_nssa_external_2"`
	SetMetricBandwidth                              types.Int64                                      `tfsdk:"set_metric_bandwidth"`
	SetMetricType                                   types.String                                     `tfsdk:"set_metric_type"`
	SetBgpAsPathPrepend                             types.List                                       `tfsdk:"set_bgp_as_path_prepend"`
	SetBgpAsPathPrependLastAs                       types.Int64                                      `tfsdk:"set_bgp_as_path_prepend_last_as"`
	SetBgpAsPathConvertRouteTagIntoAsPath           types.Bool                                       `tfsdk:"set_bgp_as_path_convert_route_tag_into_as_path"`
	SetBgpCommunityNone                             types.Bool                                       `tfsdk:"set_bgp_community_none"`
	SetBgpCommunitySpecificCommunity                types.Int64                                      `tfsdk:"set_bgp_community_specific_community"`
	SetBgpCommunityAddToExistingCommunities         types.Bool                                       `tfsdk:"set_bgp_community_add_to_existing_communities"`
	SetBgpCommunityInternet                         types.Bool                                       `tfsdk:"set_bgp_community_internet"`
	SetBgpCommunityNoAdvertise                      types.Bool                                       `tfsdk:"set_bgp_community_no_advertise"`
	SetBgpCommunityNoExport                         types.Bool                                       `tfsdk:"set_bgp_community_no_export"`
	SetBgpCommunityRouteTarget                      types.String                                     `tfsdk:"set_bgp_community_route_target"`
	SetBgpCommunityAddToExistingExtendedCommunities types.Bool                                       `tfsdk:"set_bgp_community_add_to_existing_extended_communities"`
	SetBgpAutomaticTag                              types.Bool                                       `tfsdk:"set_bgp_automatic_tag"`
	SetBgpLocalPreference                           types.Int64                                      `tfsdk:"set_bgp_local_preference"`
	SetBgpWeight                                    types.Int64                                      `tfsdk:"set_bgp_weight"`
	SetBgpOrigin                                    types.String                                     `tfsdk:"set_bgp_origin"`
	SetBgpIpv4NextHop                               types.String                                     `tfsdk:"set_bgp_ipv4_next_hop"`
	SetBgpIpv4NextHopSpecificIp                     types.List                                       `tfsdk:"set_bgp_ipv4_next_hop_specific_ip"`
	SetBgpIpv4PrefixListId                          types.String                                     `tfsdk:"set_bgp_ipv4_prefix_list_id"`
	SetBgpIpv6NextHop                               types.String                                     `tfsdk:"set_bgp_ipv6_next_hop"`
	SetBgpIpv6NextHopSpecificIp                     types.List                                       `tfsdk:"set_bgp_ipv6_next_hop_specific_ip"`
	SetBgpIpv6PrefixListId                          types.String                                     `tfsdk:"set_bgp_ipv6_prefix_list_id"`
}

type RouteMapEntriesMatchSecurityZones struct {
	Id types.String `tfsdk:"id"`
}
type RouteMapEntriesMatchIpv4AddressAccessLists struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesMatchIpv4AddressPrefixLists struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesMatchIpv4NextHopAccessLists struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesMatchIpv4NextHopPrefixLists struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesMatchIpv4RouteSourceAccessLists struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesMatchIpv4RouteSourcePrefixLists struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesMatchBgpAsPathLists struct {
	Id types.String `tfsdk:"id"`
}
type RouteMapEntriesMatchBgpCommunityLists struct {
	Id types.String `tfsdk:"id"`
}
type RouteMapEntriesMatchBgpExtendedCommunityLists struct {
	Id types.String `tfsdk:"id"`
}
type RouteMapEntriesMatchBgpPolicyLists struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data RouteMap) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/routemaps"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data RouteMap) toBody(ctx context.Context, state RouteMap) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Overridable.IsNull() {
		body, _ = sjson.Set(body, "overridable", data.Overridable.ValueBool())
	}
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if len(item.MatchSecurityZones) > 0 {
				itemBody, _ = sjson.Set(itemBody, "interfaces", []interface{}{})
				for _, childItem := range item.MatchSecurityZones {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interfaces.-1", itemChildBody)
				}
			}
			if !item.MatchInterfaceNames.IsNull() {
				var values []string
				item.MatchInterfaceNames.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "interfaceNames", values)
			}
			if len(item.MatchIpv4AddressAccessLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4AccessListAddresses", []interface{}{})
				for _, childItem := range item.MatchIpv4AddressAccessLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4AccessListAddresses.-1", itemChildBody)
				}
			}
			if len(item.MatchIpv4AddressPrefixLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4PrefixListAddresses", []interface{}{})
				for _, childItem := range item.MatchIpv4AddressPrefixLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4PrefixListAddresses.-1", itemChildBody)
				}
			}
			if len(item.MatchIpv4NextHopAccessLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4AccessListNextHops", []interface{}{})
				for _, childItem := range item.MatchIpv4NextHopAccessLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4AccessListNextHops.-1", itemChildBody)
				}
			}
			if len(item.MatchIpv4NextHopPrefixLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4PrefixListNexthops", []interface{}{})
				for _, childItem := range item.MatchIpv4NextHopPrefixLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4PrefixListNexthops.-1", itemChildBody)
				}
			}
			if len(item.MatchIpv4RouteSourceAccessLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4AccessListRouteSources", []interface{}{})
				for _, childItem := range item.MatchIpv4RouteSourceAccessLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4AccessListRouteSources.-1", itemChildBody)
				}
			}
			if len(item.MatchIpv4RouteSourcePrefixLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4PrefixListRouteSources", []interface{}{})
				for _, childItem := range item.MatchIpv4RouteSourcePrefixLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4PrefixListRouteSources.-1", itemChildBody)
				}
			}
			if !item.MatchIpv6AddressExtendedAccessListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6AccessListAddresses.0.id", item.MatchIpv6AddressExtendedAccessListId.ValueString())
			}
			if !item.MatchIpv6AddressPrefixListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6PrefixListAddresses.0.id", item.MatchIpv6AddressPrefixListId.ValueString())
			}
			if !item.MatchIpv6NextHopExtendedAccessListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6AccessListNextHops.0.id", item.MatchIpv6NextHopExtendedAccessListId.ValueString())
			}
			if !item.MatchIpv6NextHopPrefixListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6PrefixListNexthops.0.id", item.MatchIpv6NextHopPrefixListId.ValueString())
			}
			if !item.MatchIpv6RouteSourceExtendedAccessListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6AccessListRouteSources.0.id", item.MatchIpv6RouteSourceExtendedAccessListId.ValueString())
			}
			if !item.MatchIpv6RouteSourcePrefixListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipv6PrefixListRouteSources.0.id", item.MatchIpv6RouteSourcePrefixListId.ValueString())
			}
			if len(item.MatchBgpAsPathLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "asPathLists", []interface{}{})
				for _, childItem := range item.MatchBgpAsPathLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "asPathLists.-1", itemChildBody)
				}
			}
			if len(item.MatchBgpCommunityLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "communityLists", []interface{}{})
				for _, childItem := range item.MatchBgpCommunityLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "communityLists.-1", itemChildBody)
				}
			}
			if len(item.MatchBgpExtendedCommunityLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "extendedCommunityLists", []interface{}{})
				for _, childItem := range item.MatchBgpExtendedCommunityLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "extendedCommunityLists.-1", itemChildBody)
				}
			}
			if len(item.MatchBgpPolicyLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "policyLists", []interface{}{})
				for _, childItem := range item.MatchBgpPolicyLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "policyLists.-1", itemChildBody)
				}
			}
			if !item.MatchMetricRouteValues.IsNull() {
				var values []int64
				item.MatchMetricRouteValues.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "metricRouteValues", values)
			}
			if !item.MatchTagValues.IsNull() {
				var values []int64
				item.MatchTagValues.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "tagValues", values)
			}
			if !item.MatchRouteTypeExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeExternal1", item.MatchRouteTypeExternal1.ValueBool())
			}
			if !item.MatchRouteTypeExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeExternal2", item.MatchRouteTypeExternal2.ValueBool())
			}
			if !item.MatchRouteTypeInternal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeInternal", item.MatchRouteTypeInternal.ValueBool())
			}
			if !item.MatchRouteTypeLocal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeLocal", item.MatchRouteTypeLocal.ValueBool())
			}
			if !item.MatchRouteTypeNssaExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeNSSAExternal1", item.MatchRouteTypeNssaExternal1.ValueBool())
			}
			if !item.MatchRouteTypeNssaExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeNSSAExternal2", item.MatchRouteTypeNssaExternal2.ValueBool())
			}
			if !item.SetMetricBandwidth.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "metricBandwidth", item.SetMetricBandwidth.ValueInt64())
			}
			if !item.SetMetricType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "metricType", item.SetMetricType.ValueString())
			}
			if !item.SetBgpAsPathPrepend.IsNull() {
				var values []int64
				item.SetBgpAsPathPrepend.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "asPathPrependASPath", values)
			}
			if !item.SetBgpAsPathPrependLastAs.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "asPathPrependLastASToASPath", item.SetBgpAsPathPrependLastAs.ValueInt64())
			}
			if !item.SetBgpAsPathConvertRouteTagIntoAsPath.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "convertRouteTagIntoASPath", item.SetBgpAsPathConvertRouteTagIntoAsPath.ValueBool())
			}
			if !item.SetBgpCommunityNone.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "specificCommunityNone", item.SetBgpCommunityNone.ValueBool())
			}
			if !item.SetBgpCommunitySpecificCommunity.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "communityListSetting", item.SetBgpCommunitySpecificCommunity.ValueInt64())
			}
			if !item.SetBgpCommunityAddToExistingCommunities.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "addToExistingCommunity", item.SetBgpCommunityAddToExistingCommunities.ValueBool())
			}
			if !item.SetBgpCommunityInternet.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "communityListSettingInternet", item.SetBgpCommunityInternet.ValueBool())
			}
			if !item.SetBgpCommunityNoAdvertise.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "communityListSettingNoAdvertise", item.SetBgpCommunityNoAdvertise.ValueBool())
			}
			if !item.SetBgpCommunityNoExport.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "communityListSettingNoExport", item.SetBgpCommunityNoExport.ValueBool())
			}
			if !item.SetBgpCommunityRouteTarget.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "extendedCommunityRouteTarget", item.SetBgpCommunityRouteTarget.ValueString())
			}
			if !item.SetBgpCommunityAddToExistingExtendedCommunities.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "extendedCommunityAdditive", item.SetBgpCommunityAddToExistingExtendedCommunities.ValueBool())
			}
			if !item.SetBgpAutomaticTag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "automaticTagSetting", item.SetBgpAutomaticTag.ValueBool())
			}
			if !item.SetBgpLocalPreference.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "localPreferenceSetting", item.SetBgpLocalPreference.ValueInt64())
			}
			if !item.SetBgpWeight.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "weightSetting", item.SetBgpWeight.ValueInt64())
			}
			if !item.SetBgpOrigin.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "originSetting", item.SetBgpOrigin.ValueString())
			}
			if !item.SetBgpIpv4NextHop.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nextHopIPV4Setting", item.SetBgpIpv4NextHop.ValueString())
			}
			if !item.SetBgpIpv4NextHopSpecificIp.IsNull() {
				var values []string
				item.SetBgpIpv4NextHopSpecificIp.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "specificIPsIPV4Setting", values)
			}
			if !item.SetBgpIpv4PrefixListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "prefixListIPV4Setting.id", item.SetBgpIpv4PrefixListId.ValueString())
			}
			if !item.SetBgpIpv6NextHop.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "nextHopIPV6Setting", item.SetBgpIpv6NextHop.ValueString())
			}
			if !item.SetBgpIpv6NextHopSpecificIp.IsNull() {
				var values []string
				item.SetBgpIpv6NextHopSpecificIp.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "specificIPsIPV6Setting", values)
			}
			if !item.SetBgpIpv6PrefixListId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "prefixListIPV6Setting.id", item.SetBgpIpv6PrefixListId.ValueString())
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *RouteMap) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("overridable"); value.Exists() {
		data.Overridable = types.BoolValue(value.Bool())
	} else {
		data.Overridable = types.BoolNull()
	}
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]RouteMapEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := RouteMapEntries{}
			if value := res.Get("action"); value.Exists() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("interfaces"); value.Exists() {
				data.MatchSecurityZones = make([]RouteMapEntriesMatchSecurityZones, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchSecurityZones{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).MatchSecurityZones = append((*parent).MatchSecurityZones, data)
					return true
				})
			}
			if value := res.Get("interfaceNames"); value.Exists() {
				data.MatchInterfaceNames = helpers.GetStringList(value.Array())
			} else {
				data.MatchInterfaceNames = types.ListNull(types.StringType)
			}
			if value := res.Get("ipv4AccessListAddresses"); value.Exists() {
				data.MatchIpv4AddressAccessLists = make([]RouteMapEntriesMatchIpv4AddressAccessLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchIpv4AddressAccessLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).MatchIpv4AddressAccessLists = append((*parent).MatchIpv4AddressAccessLists, data)
					return true
				})
			}
			if value := res.Get("ipv4PrefixListAddresses"); value.Exists() {
				data.MatchIpv4AddressPrefixLists = make([]RouteMapEntriesMatchIpv4AddressPrefixLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchIpv4AddressPrefixLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).MatchIpv4AddressPrefixLists = append((*parent).MatchIpv4AddressPrefixLists, data)
					return true
				})
			}
			if value := res.Get("ipv4AccessListNextHops"); value.Exists() {
				data.MatchIpv4NextHopAccessLists = make([]RouteMapEntriesMatchIpv4NextHopAccessLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchIpv4NextHopAccessLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).MatchIpv4NextHopAccessLists = append((*parent).MatchIpv4NextHopAccessLists, data)
					return true
				})
			}
			if value := res.Get("ipv4PrefixListNexthops"); value.Exists() {
				data.MatchIpv4NextHopPrefixLists = make([]RouteMapEntriesMatchIpv4NextHopPrefixLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchIpv4NextHopPrefixLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).MatchIpv4NextHopPrefixLists = append((*parent).MatchIpv4NextHopPrefixLists, data)
					return true
				})
			}
			if value := res.Get("ipv4AccessListRouteSources"); value.Exists() {
				data.MatchIpv4RouteSourceAccessLists = make([]RouteMapEntriesMatchIpv4RouteSourceAccessLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchIpv4RouteSourceAccessLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).MatchIpv4RouteSourceAccessLists = append((*parent).MatchIpv4RouteSourceAccessLists, data)
					return true
				})
			}
			if value := res.Get("ipv4PrefixListRouteSources"); value.Exists() {
				data.MatchIpv4RouteSourcePrefixLists = make([]RouteMapEntriesMatchIpv4RouteSourcePrefixLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchIpv4RouteSourcePrefixLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).MatchIpv4RouteSourcePrefixLists = append((*parent).MatchIpv4RouteSourcePrefixLists, data)
					return true
				})
			}
			if value := res.Get("ipv6AccessListAddresses.0.id"); value.Exists() {
				data.MatchIpv6AddressExtendedAccessListId = types.StringValue(value.String())
			} else {
				data.MatchIpv6AddressExtendedAccessListId = types.StringNull()
			}
			if value := res.Get("ipv6PrefixListAddresses.0.id"); value.Exists() {
				data.MatchIpv6AddressPrefixListId = types.StringValue(value.String())
			} else {
				data.MatchIpv6AddressPrefixListId = types.StringNull()
			}
			if value := res.Get("ipv6AccessListNextHops.0.id"); value.Exists() {
				data.MatchIpv6NextHopExtendedAccessListId = types.StringValue(value.String())
			} else {
				data.MatchIpv6NextHopExtendedAccessListId = types.StringNull()
			}
			if value := res.Get("ipv6PrefixListNexthops.0.id"); value.Exists() {
				data.MatchIpv6NextHopPrefixListId = types.StringValue(value.String())
			} else {
				data.MatchIpv6NextHopPrefixListId = types.StringNull()
			}
			if value := res.Get("ipv6AccessListRouteSources.0.id"); value.Exists() {
				data.MatchIpv6RouteSourceExtendedAccessListId = types.StringValue(value.String())
			} else {
				data.MatchIpv6RouteSourceExtendedAccessListId = types.StringNull()
			}
			if value := res.Get("ipv6PrefixListRouteSources.0.id"); value.Exists() {
				data.MatchIpv6RouteSourcePrefixListId = types.StringValue(value.String())
			} else {
				data.MatchIpv6RouteSourcePrefixListId = types.StringNull()
			}
			if value := res.Get("asPathLists"); value.Exists() {
				data.MatchBgpAsPathLists = make([]RouteMapEntriesMatchBgpAsPathLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchBgpAsPathLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).MatchBgpAsPathLists = append((*parent).MatchBgpAsPathLists, data)
					return true
				})
			}
			if value := res.Get("communityLists"); value.Exists() {
				data.MatchBgpCommunityLists = make([]RouteMapEntriesMatchBgpCommunityLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchBgpCommunityLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).MatchBgpCommunityLists = append((*parent).MatchBgpCommunityLists, data)
					return true
				})
			}
			if value := res.Get("extendedCommunityLists"); value.Exists() {
				data.MatchBgpExtendedCommunityLists = make([]RouteMapEntriesMatchBgpExtendedCommunityLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchBgpExtendedCommunityLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).MatchBgpExtendedCommunityLists = append((*parent).MatchBgpExtendedCommunityLists, data)
					return true
				})
			}
			if value := res.Get("policyLists"); value.Exists() {
				data.MatchBgpPolicyLists = make([]RouteMapEntriesMatchBgpPolicyLists, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesMatchBgpPolicyLists{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).MatchBgpPolicyLists = append((*parent).MatchBgpPolicyLists, data)
					return true
				})
			}
			if value := res.Get("metricRouteValues"); value.Exists() {
				data.MatchMetricRouteValues = helpers.GetInt64List(value.Array())
			} else {
				data.MatchMetricRouteValues = types.ListNull(types.Int64Type)
			}
			if value := res.Get("tagValues"); value.Exists() {
				data.MatchTagValues = helpers.GetInt64List(value.Array())
			} else {
				data.MatchTagValues = types.ListNull(types.Int64Type)
			}
			if value := res.Get("routeTypeExternal1"); value.Exists() {
				data.MatchRouteTypeExternal1 = types.BoolValue(value.Bool())
			} else {
				data.MatchRouteTypeExternal1 = types.BoolNull()
			}
			if value := res.Get("routeTypeExternal2"); value.Exists() {
				data.MatchRouteTypeExternal2 = types.BoolValue(value.Bool())
			} else {
				data.MatchRouteTypeExternal2 = types.BoolNull()
			}
			if value := res.Get("routeTypeInternal"); value.Exists() {
				data.MatchRouteTypeInternal = types.BoolValue(value.Bool())
			} else {
				data.MatchRouteTypeInternal = types.BoolNull()
			}
			if value := res.Get("routeTypeLocal"); value.Exists() {
				data.MatchRouteTypeLocal = types.BoolValue(value.Bool())
			} else {
				data.MatchRouteTypeLocal = types.BoolNull()
			}
			if value := res.Get("routeTypeNSSAExternal1"); value.Exists() {
				data.MatchRouteTypeNssaExternal1 = types.BoolValue(value.Bool())
			} else {
				data.MatchRouteTypeNssaExternal1 = types.BoolNull()
			}
			if value := res.Get("routeTypeNSSAExternal2"); value.Exists() {
				data.MatchRouteTypeNssaExternal2 = types.BoolValue(value.Bool())
			} else {
				data.MatchRouteTypeNssaExternal2 = types.BoolNull()
			}
			if value := res.Get("metricBandwidth"); value.Exists() {
				data.SetMetricBandwidth = types.Int64Value(value.Int())
			} else {
				data.SetMetricBandwidth = types.Int64Null()
			}
			if value := res.Get("metricType"); value.Exists() {
				data.SetMetricType = types.StringValue(value.String())
			} else {
				data.SetMetricType = types.StringNull()
			}
			if value := res.Get("asPathPrependASPath"); value.Exists() {
				data.SetBgpAsPathPrepend = helpers.GetInt64List(value.Array())
			} else {
				data.SetBgpAsPathPrepend = types.ListNull(types.Int64Type)
			}
			if value := res.Get("asPathPrependLastASToASPath"); value.Exists() {
				data.SetBgpAsPathPrependLastAs = types.Int64Value(value.Int())
			} else {
				data.SetBgpAsPathPrependLastAs = types.Int64Null()
			}
			if value := res.Get("convertRouteTagIntoASPath"); value.Exists() {
				data.SetBgpAsPathConvertRouteTagIntoAsPath = types.BoolValue(value.Bool())
			} else {
				data.SetBgpAsPathConvertRouteTagIntoAsPath = types.BoolNull()
			}
			if value := res.Get("specificCommunityNone"); value.Exists() {
				data.SetBgpCommunityNone = types.BoolValue(value.Bool())
			} else {
				data.SetBgpCommunityNone = types.BoolNull()
			}
			if value := res.Get("communityListSetting"); value.Exists() {
				data.SetBgpCommunitySpecificCommunity = types.Int64Value(value.Int())
			} else {
				data.SetBgpCommunitySpecificCommunity = types.Int64Null()
			}
			if value := res.Get("addToExistingCommunity"); value.Exists() {
				data.SetBgpCommunityAddToExistingCommunities = types.BoolValue(value.Bool())
			} else {
				data.SetBgpCommunityAddToExistingCommunities = types.BoolNull()
			}
			if value := res.Get("communityListSettingInternet"); value.Exists() {
				data.SetBgpCommunityInternet = types.BoolValue(value.Bool())
			} else {
				data.SetBgpCommunityInternet = types.BoolNull()
			}
			if value := res.Get("communityListSettingNoAdvertise"); value.Exists() {
				data.SetBgpCommunityNoAdvertise = types.BoolValue(value.Bool())
			} else {
				data.SetBgpCommunityNoAdvertise = types.BoolNull()
			}
			if value := res.Get("communityListSettingNoExport"); value.Exists() {
				data.SetBgpCommunityNoExport = types.BoolValue(value.Bool())
			} else {
				data.SetBgpCommunityNoExport = types.BoolNull()
			}
			if value := res.Get("extendedCommunityRouteTarget"); value.Exists() {
				data.SetBgpCommunityRouteTarget = types.StringValue(value.String())
			} else {
				data.SetBgpCommunityRouteTarget = types.StringNull()
			}
			if value := res.Get("extendedCommunityAdditive"); value.Exists() {
				data.SetBgpCommunityAddToExistingExtendedCommunities = types.BoolValue(value.Bool())
			} else {
				data.SetBgpCommunityAddToExistingExtendedCommunities = types.BoolNull()
			}
			if value := res.Get("automaticTagSetting"); value.Exists() {
				data.SetBgpAutomaticTag = types.BoolValue(value.Bool())
			} else {
				data.SetBgpAutomaticTag = types.BoolNull()
			}
			if value := res.Get("localPreferenceSetting"); value.Exists() {
				data.SetBgpLocalPreference = types.Int64Value(value.Int())
			} else {
				data.SetBgpLocalPreference = types.Int64Null()
			}
			if value := res.Get("weightSetting"); value.Exists() {
				data.SetBgpWeight = types.Int64Value(value.Int())
			} else {
				data.SetBgpWeight = types.Int64Null()
			}
			if value := res.Get("originSetting"); value.Exists() {
				data.SetBgpOrigin = types.StringValue(value.String())
			} else {
				data.SetBgpOrigin = types.StringNull()
			}
			if value := res.Get("nextHopIPV4Setting"); value.Exists() {
				data.SetBgpIpv4NextHop = types.StringValue(value.String())
			} else {
				data.SetBgpIpv4NextHop = types.StringNull()
			}
			if value := res.Get("specificIPsIPV4Setting"); value.Exists() {
				data.SetBgpIpv4NextHopSpecificIp = helpers.GetStringList(value.Array())
			} else {
				data.SetBgpIpv4NextHopSpecificIp = types.ListNull(types.StringType)
			}
			if value := res.Get("prefixListIPV4Setting.id"); value.Exists() {
				data.SetBgpIpv4PrefixListId = types.StringValue(value.String())
			} else {
				data.SetBgpIpv4PrefixListId = types.StringNull()
			}
			if value := res.Get("nextHopIPV6Setting"); value.Exists() {
				data.SetBgpIpv6NextHop = types.StringValue(value.String())
			} else {
				data.SetBgpIpv6NextHop = types.StringNull()
			}
			if value := res.Get("specificIPsIPV6Setting"); value.Exists() {
				data.SetBgpIpv6NextHopSpecificIp = helpers.GetStringList(value.Array())
			} else {
				data.SetBgpIpv6NextHopSpecificIp = types.ListNull(types.StringType)
			}
			if value := res.Get("prefixListIPV6Setting.id"); value.Exists() {
				data.SetBgpIpv6PrefixListId = types.StringValue(value.String())
			} else {
				data.SetBgpIpv6PrefixListId = types.StringNull()
			}
			(*parent).Entries = append((*parent).Entries, data)
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
func (data *RouteMap) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("overridable"); value.Exists() && !data.Overridable.IsNull() {
		data.Overridable = types.BoolValue(value.Bool())
	} else {
		data.Overridable = types.BoolNull()
	}
	{
		l := len(res.Get("entries").Array())
		tflog.Debug(ctx, fmt.Sprintf("entries array resizing from %d to %d", len(data.Entries), l))
		for i := len(data.Entries); i < l; i++ {
			data.Entries = append(data.Entries, RouteMapEntries{})
		}
		if len(data.Entries) > l {
			data.Entries = data.Entries[:l]
		}
	}
	for i := range data.Entries {
		parent := &data
		data := (*parent).Entries[i]
		parentRes := &res
		res := parentRes.Get(fmt.Sprintf("entries.%d", i))
		if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
			data.Action = types.StringValue(value.String())
		} else {
			data.Action = types.StringNull()
		}
		for i := 0; i < len(data.MatchSecurityZones); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.MatchSecurityZones[i].Id.ValueString()}

			parent := &data
			data := (*parent).MatchSecurityZones[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("interfaces").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchSecurityZones[%d] = %+v",
					i,
					(*parent).MatchSecurityZones[i],
				))
				(*parent).MatchSecurityZones = slices.Delete((*parent).MatchSecurityZones, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).MatchSecurityZones[i] = data
		}
		if value := res.Get("interfaceNames"); value.Exists() && !data.MatchInterfaceNames.IsNull() {
			data.MatchInterfaceNames = helpers.GetStringList(value.Array())
		} else {
			data.MatchInterfaceNames = types.ListNull(types.StringType)
		}
		for i := 0; i < len(data.MatchIpv4AddressAccessLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.MatchIpv4AddressAccessLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).MatchIpv4AddressAccessLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv4AccessListAddresses").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchIpv4AddressAccessLists[%d] = %+v",
					i,
					(*parent).MatchIpv4AddressAccessLists[i],
				))
				(*parent).MatchIpv4AddressAccessLists = slices.Delete((*parent).MatchIpv4AddressAccessLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).MatchIpv4AddressAccessLists[i] = data
		}
		for i := 0; i < len(data.MatchIpv4AddressPrefixLists); i++ {
			keys := [...]string{"id", "type"}
			keyValues := [...]string{data.MatchIpv4AddressPrefixLists[i].Id.ValueString(), data.MatchIpv4AddressPrefixLists[i].Type.ValueString()}

			parent := &data
			data := (*parent).MatchIpv4AddressPrefixLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv4PrefixListAddresses").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchIpv4AddressPrefixLists[%d] = %+v",
					i,
					(*parent).MatchIpv4AddressPrefixLists[i],
				))
				(*parent).MatchIpv4AddressPrefixLists = slices.Delete((*parent).MatchIpv4AddressPrefixLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).MatchIpv4AddressPrefixLists[i] = data
		}
		for i := 0; i < len(data.MatchIpv4NextHopAccessLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.MatchIpv4NextHopAccessLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).MatchIpv4NextHopAccessLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv4AccessListNextHops").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchIpv4NextHopAccessLists[%d] = %+v",
					i,
					(*parent).MatchIpv4NextHopAccessLists[i],
				))
				(*parent).MatchIpv4NextHopAccessLists = slices.Delete((*parent).MatchIpv4NextHopAccessLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).MatchIpv4NextHopAccessLists[i] = data
		}
		for i := 0; i < len(data.MatchIpv4NextHopPrefixLists); i++ {
			keys := [...]string{"id", "type"}
			keyValues := [...]string{data.MatchIpv4NextHopPrefixLists[i].Id.ValueString(), data.MatchIpv4NextHopPrefixLists[i].Type.ValueString()}

			parent := &data
			data := (*parent).MatchIpv4NextHopPrefixLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv4PrefixListNexthops").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchIpv4NextHopPrefixLists[%d] = %+v",
					i,
					(*parent).MatchIpv4NextHopPrefixLists[i],
				))
				(*parent).MatchIpv4NextHopPrefixLists = slices.Delete((*parent).MatchIpv4NextHopPrefixLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).MatchIpv4NextHopPrefixLists[i] = data
		}
		for i := 0; i < len(data.MatchIpv4RouteSourceAccessLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.MatchIpv4RouteSourceAccessLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).MatchIpv4RouteSourceAccessLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv4AccessListRouteSources").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchIpv4RouteSourceAccessLists[%d] = %+v",
					i,
					(*parent).MatchIpv4RouteSourceAccessLists[i],
				))
				(*parent).MatchIpv4RouteSourceAccessLists = slices.Delete((*parent).MatchIpv4RouteSourceAccessLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).MatchIpv4RouteSourceAccessLists[i] = data
		}
		for i := 0; i < len(data.MatchIpv4RouteSourcePrefixLists); i++ {
			keys := [...]string{"id", "type"}
			keyValues := [...]string{data.MatchIpv4RouteSourcePrefixLists[i].Id.ValueString(), data.MatchIpv4RouteSourcePrefixLists[i].Type.ValueString()}

			parent := &data
			data := (*parent).MatchIpv4RouteSourcePrefixLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv4PrefixListRouteSources").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchIpv4RouteSourcePrefixLists[%d] = %+v",
					i,
					(*parent).MatchIpv4RouteSourcePrefixLists[i],
				))
				(*parent).MatchIpv4RouteSourcePrefixLists = slices.Delete((*parent).MatchIpv4RouteSourcePrefixLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).MatchIpv4RouteSourcePrefixLists[i] = data
		}
		if value := res.Get("ipv6AccessListAddresses.0.id"); value.Exists() && !data.MatchIpv6AddressExtendedAccessListId.IsNull() {
			data.MatchIpv6AddressExtendedAccessListId = types.StringValue(value.String())
		} else {
			data.MatchIpv6AddressExtendedAccessListId = types.StringNull()
		}
		if value := res.Get("ipv6PrefixListAddresses.0.id"); value.Exists() && !data.MatchIpv6AddressPrefixListId.IsNull() {
			data.MatchIpv6AddressPrefixListId = types.StringValue(value.String())
		} else {
			data.MatchIpv6AddressPrefixListId = types.StringNull()
		}
		if value := res.Get("ipv6AccessListNextHops.0.id"); value.Exists() && !data.MatchIpv6NextHopExtendedAccessListId.IsNull() {
			data.MatchIpv6NextHopExtendedAccessListId = types.StringValue(value.String())
		} else {
			data.MatchIpv6NextHopExtendedAccessListId = types.StringNull()
		}
		if value := res.Get("ipv6PrefixListNexthops.0.id"); value.Exists() && !data.MatchIpv6NextHopPrefixListId.IsNull() {
			data.MatchIpv6NextHopPrefixListId = types.StringValue(value.String())
		} else {
			data.MatchIpv6NextHopPrefixListId = types.StringNull()
		}
		if value := res.Get("ipv6AccessListRouteSources.0.id"); value.Exists() && !data.MatchIpv6RouteSourceExtendedAccessListId.IsNull() {
			data.MatchIpv6RouteSourceExtendedAccessListId = types.StringValue(value.String())
		} else {
			data.MatchIpv6RouteSourceExtendedAccessListId = types.StringNull()
		}
		if value := res.Get("ipv6PrefixListRouteSources.0.id"); value.Exists() && !data.MatchIpv6RouteSourcePrefixListId.IsNull() {
			data.MatchIpv6RouteSourcePrefixListId = types.StringValue(value.String())
		} else {
			data.MatchIpv6RouteSourcePrefixListId = types.StringNull()
		}
		for i := 0; i < len(data.MatchBgpAsPathLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.MatchBgpAsPathLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).MatchBgpAsPathLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("asPathLists").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchBgpAsPathLists[%d] = %+v",
					i,
					(*parent).MatchBgpAsPathLists[i],
				))
				(*parent).MatchBgpAsPathLists = slices.Delete((*parent).MatchBgpAsPathLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).MatchBgpAsPathLists[i] = data
		}
		for i := 0; i < len(data.MatchBgpCommunityLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.MatchBgpCommunityLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).MatchBgpCommunityLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("communityLists").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchBgpCommunityLists[%d] = %+v",
					i,
					(*parent).MatchBgpCommunityLists[i],
				))
				(*parent).MatchBgpCommunityLists = slices.Delete((*parent).MatchBgpCommunityLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).MatchBgpCommunityLists[i] = data
		}
		for i := 0; i < len(data.MatchBgpExtendedCommunityLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.MatchBgpExtendedCommunityLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).MatchBgpExtendedCommunityLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("extendedCommunityLists").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchBgpExtendedCommunityLists[%d] = %+v",
					i,
					(*parent).MatchBgpExtendedCommunityLists[i],
				))
				(*parent).MatchBgpExtendedCommunityLists = slices.Delete((*parent).MatchBgpExtendedCommunityLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).MatchBgpExtendedCommunityLists[i] = data
		}
		for i := 0; i < len(data.MatchBgpPolicyLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.MatchBgpPolicyLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).MatchBgpPolicyLists[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("policyLists").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing MatchBgpPolicyLists[%d] = %+v",
					i,
					(*parent).MatchBgpPolicyLists[i],
				))
				(*parent).MatchBgpPolicyLists = slices.Delete((*parent).MatchBgpPolicyLists, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).MatchBgpPolicyLists[i] = data
		}
		if value := res.Get("metricRouteValues"); value.Exists() && !data.MatchMetricRouteValues.IsNull() {
			data.MatchMetricRouteValues = helpers.GetInt64List(value.Array())
		} else {
			data.MatchMetricRouteValues = types.ListNull(types.Int64Type)
		}
		if value := res.Get("tagValues"); value.Exists() && !data.MatchTagValues.IsNull() {
			data.MatchTagValues = helpers.GetInt64List(value.Array())
		} else {
			data.MatchTagValues = types.ListNull(types.Int64Type)
		}
		if value := res.Get("routeTypeExternal1"); value.Exists() && !data.MatchRouteTypeExternal1.IsNull() {
			data.MatchRouteTypeExternal1 = types.BoolValue(value.Bool())
		} else {
			data.MatchRouteTypeExternal1 = types.BoolNull()
		}
		if value := res.Get("routeTypeExternal2"); value.Exists() && !data.MatchRouteTypeExternal2.IsNull() {
			data.MatchRouteTypeExternal2 = types.BoolValue(value.Bool())
		} else {
			data.MatchRouteTypeExternal2 = types.BoolNull()
		}
		if value := res.Get("routeTypeInternal"); value.Exists() && !data.MatchRouteTypeInternal.IsNull() {
			data.MatchRouteTypeInternal = types.BoolValue(value.Bool())
		} else {
			data.MatchRouteTypeInternal = types.BoolNull()
		}
		if value := res.Get("routeTypeLocal"); value.Exists() && !data.MatchRouteTypeLocal.IsNull() {
			data.MatchRouteTypeLocal = types.BoolValue(value.Bool())
		} else {
			data.MatchRouteTypeLocal = types.BoolNull()
		}
		if value := res.Get("routeTypeNSSAExternal1"); value.Exists() && !data.MatchRouteTypeNssaExternal1.IsNull() {
			data.MatchRouteTypeNssaExternal1 = types.BoolValue(value.Bool())
		} else {
			data.MatchRouteTypeNssaExternal1 = types.BoolNull()
		}
		if value := res.Get("routeTypeNSSAExternal2"); value.Exists() && !data.MatchRouteTypeNssaExternal2.IsNull() {
			data.MatchRouteTypeNssaExternal2 = types.BoolValue(value.Bool())
		} else {
			data.MatchRouteTypeNssaExternal2 = types.BoolNull()
		}
		if value := res.Get("metricBandwidth"); value.Exists() && !data.SetMetricBandwidth.IsNull() {
			data.SetMetricBandwidth = types.Int64Value(value.Int())
		} else {
			data.SetMetricBandwidth = types.Int64Null()
		}
		if value := res.Get("metricType"); value.Exists() && !data.SetMetricType.IsNull() {
			data.SetMetricType = types.StringValue(value.String())
		} else {
			data.SetMetricType = types.StringNull()
		}
		if value := res.Get("asPathPrependASPath"); value.Exists() && !data.SetBgpAsPathPrepend.IsNull() {
			data.SetBgpAsPathPrepend = helpers.GetInt64List(value.Array())
		} else {
			data.SetBgpAsPathPrepend = types.ListNull(types.Int64Type)
		}
		if value := res.Get("asPathPrependLastASToASPath"); value.Exists() && !data.SetBgpAsPathPrependLastAs.IsNull() {
			data.SetBgpAsPathPrependLastAs = types.Int64Value(value.Int())
		} else {
			data.SetBgpAsPathPrependLastAs = types.Int64Null()
		}
		if value := res.Get("convertRouteTagIntoASPath"); value.Exists() && !data.SetBgpAsPathConvertRouteTagIntoAsPath.IsNull() {
			data.SetBgpAsPathConvertRouteTagIntoAsPath = types.BoolValue(value.Bool())
		} else {
			data.SetBgpAsPathConvertRouteTagIntoAsPath = types.BoolNull()
		}
		if value := res.Get("specificCommunityNone"); value.Exists() && !data.SetBgpCommunityNone.IsNull() {
			data.SetBgpCommunityNone = types.BoolValue(value.Bool())
		} else {
			data.SetBgpCommunityNone = types.BoolNull()
		}
		if value := res.Get("communityListSetting"); value.Exists() && !data.SetBgpCommunitySpecificCommunity.IsNull() {
			data.SetBgpCommunitySpecificCommunity = types.Int64Value(value.Int())
		} else {
			data.SetBgpCommunitySpecificCommunity = types.Int64Null()
		}
		if value := res.Get("addToExistingCommunity"); value.Exists() && !data.SetBgpCommunityAddToExistingCommunities.IsNull() {
			data.SetBgpCommunityAddToExistingCommunities = types.BoolValue(value.Bool())
		} else {
			data.SetBgpCommunityAddToExistingCommunities = types.BoolNull()
		}
		if value := res.Get("communityListSettingInternet"); value.Exists() && !data.SetBgpCommunityInternet.IsNull() {
			data.SetBgpCommunityInternet = types.BoolValue(value.Bool())
		} else {
			data.SetBgpCommunityInternet = types.BoolNull()
		}
		if value := res.Get("communityListSettingNoAdvertise"); value.Exists() && !data.SetBgpCommunityNoAdvertise.IsNull() {
			data.SetBgpCommunityNoAdvertise = types.BoolValue(value.Bool())
		} else {
			data.SetBgpCommunityNoAdvertise = types.BoolNull()
		}
		if value := res.Get("communityListSettingNoExport"); value.Exists() && !data.SetBgpCommunityNoExport.IsNull() {
			data.SetBgpCommunityNoExport = types.BoolValue(value.Bool())
		} else {
			data.SetBgpCommunityNoExport = types.BoolNull()
		}
		if value := res.Get("extendedCommunityRouteTarget"); value.Exists() && !data.SetBgpCommunityRouteTarget.IsNull() {
			data.SetBgpCommunityRouteTarget = types.StringValue(value.String())
		} else {
			data.SetBgpCommunityRouteTarget = types.StringNull()
		}
		if value := res.Get("extendedCommunityAdditive"); value.Exists() && !data.SetBgpCommunityAddToExistingExtendedCommunities.IsNull() {
			data.SetBgpCommunityAddToExistingExtendedCommunities = types.BoolValue(value.Bool())
		} else {
			data.SetBgpCommunityAddToExistingExtendedCommunities = types.BoolNull()
		}
		if value := res.Get("automaticTagSetting"); value.Exists() && !data.SetBgpAutomaticTag.IsNull() {
			data.SetBgpAutomaticTag = types.BoolValue(value.Bool())
		} else {
			data.SetBgpAutomaticTag = types.BoolNull()
		}
		if value := res.Get("localPreferenceSetting"); value.Exists() && !data.SetBgpLocalPreference.IsNull() {
			data.SetBgpLocalPreference = types.Int64Value(value.Int())
		} else {
			data.SetBgpLocalPreference = types.Int64Null()
		}
		if value := res.Get("weightSetting"); value.Exists() && !data.SetBgpWeight.IsNull() {
			data.SetBgpWeight = types.Int64Value(value.Int())
		} else {
			data.SetBgpWeight = types.Int64Null()
		}
		if value := res.Get("originSetting"); value.Exists() && !data.SetBgpOrigin.IsNull() {
			data.SetBgpOrigin = types.StringValue(value.String())
		} else {
			data.SetBgpOrigin = types.StringNull()
		}
		if value := res.Get("nextHopIPV4Setting"); value.Exists() && !data.SetBgpIpv4NextHop.IsNull() {
			data.SetBgpIpv4NextHop = types.StringValue(value.String())
		} else {
			data.SetBgpIpv4NextHop = types.StringNull()
		}
		if value := res.Get("specificIPsIPV4Setting"); value.Exists() && !data.SetBgpIpv4NextHopSpecificIp.IsNull() {
			data.SetBgpIpv4NextHopSpecificIp = helpers.GetStringList(value.Array())
		} else {
			data.SetBgpIpv4NextHopSpecificIp = types.ListNull(types.StringType)
		}
		if value := res.Get("prefixListIPV4Setting.id"); value.Exists() && !data.SetBgpIpv4PrefixListId.IsNull() {
			data.SetBgpIpv4PrefixListId = types.StringValue(value.String())
		} else {
			data.SetBgpIpv4PrefixListId = types.StringNull()
		}
		if value := res.Get("nextHopIPV6Setting"); value.Exists() && !data.SetBgpIpv6NextHop.IsNull() {
			data.SetBgpIpv6NextHop = types.StringValue(value.String())
		} else {
			data.SetBgpIpv6NextHop = types.StringNull()
		}
		if value := res.Get("specificIPsIPV6Setting"); value.Exists() && !data.SetBgpIpv6NextHopSpecificIp.IsNull() {
			data.SetBgpIpv6NextHopSpecificIp = helpers.GetStringList(value.Array())
		} else {
			data.SetBgpIpv6NextHopSpecificIp = types.ListNull(types.StringType)
		}
		if value := res.Get("prefixListIPV6Setting.id"); value.Exists() && !data.SetBgpIpv6PrefixListId.IsNull() {
			data.SetBgpIpv6PrefixListId = types.StringValue(value.String())
		} else {
			data.SetBgpIpv6PrefixListId = types.StringNull()
		}
		(*parent).Entries[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *RouteMap) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

func (data RouteMap) adjustBody(ctx context.Context, req string) string {

	for i := range len(data.Entries) {
		req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.sequence", i), i+1)

		if gjson.Get(req, fmt.Sprintf("entries.%d.ipv6AccessListAddresses", i)).Exists() {
			req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.ipv6AccessListAddresses.0.type", i), "ExtendedAccessList")
		}

		if gjson.Get(req, fmt.Sprintf("entries.%d.ipv6AccessListNextHops", i)).Exists() {
			req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.ipv6AccessListNextHops.0.type", i), "ExtendedAccessList")
		}

		if gjson.Get(req, fmt.Sprintf("entries.%d.ipv6AccessListRouteSources", i)).Exists() {
			req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.ipv6AccessListRouteSources.0.type", i), "ExtendedAccessList")
		}

		if gjson.Get(req, fmt.Sprintf("entries.%d.ipv6PrefixListAddresses", i)).Exists() {
			req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.ipv6PrefixListAddresses.0.type", i), "IPV6PrefixList")
		}

		if gjson.Get(req, fmt.Sprintf("entries.%d.ipv6PrefixListNexthops", i)).Exists() {
			req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.ipv6PrefixListNexthops.0.type", i), "IPV6PrefixList")
		}

		if gjson.Get(req, fmt.Sprintf("entries.%d.ipv6PrefixListRouteSources", i)).Exists() {
			req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.ipv6PrefixListRouteSources.0.type", i), "IPV6PrefixList")
		}

	}

	return req
}
