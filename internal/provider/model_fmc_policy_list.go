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

type PolicyList struct {
	Id                             types.String                               `tfsdk:"id"`
	Domain                         types.String                               `tfsdk:"domain"`
	Name                           types.String                               `tfsdk:"name"`
	Type                           types.String                               `tfsdk:"type"`
	Action                         types.String                               `tfsdk:"action"`
	Interfaces                     []PolicyListInterfaces                     `tfsdk:"interfaces"`
	InterfaceNames                 types.List                                 `tfsdk:"interface_names"`
	AddressStandardAccessLists     []PolicyListAddressStandardAccessLists     `tfsdk:"address_standard_access_lists"`
	AddressIpv4PrefixLists         []PolicyListAddressIpv4PrefixLists         `tfsdk:"address_ipv4_prefix_lists"`
	NextHopStandardAccessLists     []PolicyListNextHopStandardAccessLists     `tfsdk:"next_hop_standard_access_lists"`
	NextHopIpv4PrefixLists         []PolicyListNextHopIpv4PrefixLists         `tfsdk:"next_hop_ipv4_prefix_lists"`
	RouteSourceStandardAccessLists []PolicyListRouteSourceStandardAccessLists `tfsdk:"route_source_standard_access_lists"`
	RouteSourceIpv4PrefixLists     []PolicyListRouteSourceIpv4PrefixLists     `tfsdk:"route_source_ipv4_prefix_lists"`
	AsPathLists                    []PolicyListAsPathLists                    `tfsdk:"as_path_lists"`
	CommunityLists                 []PolicyListCommunityLists                 `tfsdk:"community_lists"`
	ExtendedCommunityLists         []PolicyListExtendedCommunityLists         `tfsdk:"extended_community_lists"`
	MatchCommunityExactly          types.Bool                                 `tfsdk:"match_community_exactly"`
	Metric                         types.Int64                                `tfsdk:"metric"`
	Tag                            types.Int64                                `tfsdk:"tag"`
}

type PolicyListInterfaces struct {
	Id types.String `tfsdk:"id"`
}

type PolicyListAddressStandardAccessLists struct {
	Id types.String `tfsdk:"id"`
}

type PolicyListAddressIpv4PrefixLists struct {
	Id types.String `tfsdk:"id"`
}

type PolicyListNextHopStandardAccessLists struct {
	Id types.String `tfsdk:"id"`
}

type PolicyListNextHopIpv4PrefixLists struct {
	Id types.String `tfsdk:"id"`
}

type PolicyListRouteSourceStandardAccessLists struct {
	Id types.String `tfsdk:"id"`
}

type PolicyListRouteSourceIpv4PrefixLists struct {
	Id types.String `tfsdk:"id"`
}

type PolicyListAsPathLists struct {
	Id types.String `tfsdk:"id"`
}

type PolicyListCommunityLists struct {
	Id types.String `tfsdk:"id"`
}

type PolicyListExtendedCommunityLists struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data PolicyList) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/policylists"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data PolicyList) toBody(ctx context.Context, state PolicyList) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Action.IsNull() {
		body, _ = sjson.Set(body, "action", data.Action.ValueString())
	}
	if len(data.Interfaces) > 0 {
		body, _ = sjson.Set(body, "interfaces", []interface{}{})
		for _, item := range data.Interfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "interfaces.-1", itemBody)
		}
	}
	if !data.InterfaceNames.IsNull() {
		var values []string
		data.InterfaceNames.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "interfaceNames", values)
	}
	if len(data.AddressStandardAccessLists) > 0 {
		body, _ = sjson.Set(body, "standardAccessListAddresses", []interface{}{})
		for _, item := range data.AddressStandardAccessLists {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "standardAccessListAddresses.-1", itemBody)
		}
	}
	if len(data.AddressIpv4PrefixLists) > 0 {
		body, _ = sjson.Set(body, "ipv4PrefixListAddresses", []interface{}{})
		for _, item := range data.AddressIpv4PrefixLists {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ipv4PrefixListAddresses.-1", itemBody)
		}
	}
	if len(data.NextHopStandardAccessLists) > 0 {
		body, _ = sjson.Set(body, "standardAccessListNextHops", []interface{}{})
		for _, item := range data.NextHopStandardAccessLists {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "standardAccessListNextHops.-1", itemBody)
		}
	}
	if len(data.NextHopIpv4PrefixLists) > 0 {
		body, _ = sjson.Set(body, "ipv4PrefixListNexthops", []interface{}{})
		for _, item := range data.NextHopIpv4PrefixLists {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ipv4PrefixListNexthops.-1", itemBody)
		}
	}
	if len(data.RouteSourceStandardAccessLists) > 0 {
		body, _ = sjson.Set(body, "standardAccessListRouteSources", []interface{}{})
		for _, item := range data.RouteSourceStandardAccessLists {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "standardAccessListRouteSources.-1", itemBody)
		}
	}
	if len(data.RouteSourceIpv4PrefixLists) > 0 {
		body, _ = sjson.Set(body, "ipv4PrefixListRouteSources", []interface{}{})
		for _, item := range data.RouteSourceIpv4PrefixLists {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ipv4PrefixListRouteSources.-1", itemBody)
		}
	}
	if len(data.AsPathLists) > 0 {
		body, _ = sjson.Set(body, "asPathLists", []interface{}{})
		for _, item := range data.AsPathLists {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "asPathLists.-1", itemBody)
		}
	}
	if len(data.CommunityLists) > 0 {
		body, _ = sjson.Set(body, "communityLists", []interface{}{})
		for _, item := range data.CommunityLists {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "communityLists.-1", itemBody)
		}
	}
	if len(data.ExtendedCommunityLists) > 0 {
		body, _ = sjson.Set(body, "extendedCommunityLists", []interface{}{})
		for _, item := range data.ExtendedCommunityLists {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "extendedCommunityLists.-1", itemBody)
		}
	}
	if !data.MatchCommunityExactly.IsNull() {
		body, _ = sjson.Set(body, "matchCommunityExactly", data.MatchCommunityExactly.ValueBool())
	}
	if !data.Metric.IsNull() {
		body, _ = sjson.Set(body, "metric", data.Metric.ValueInt64())
	}
	if !data.Tag.IsNull() {
		body, _ = sjson.Set(body, "tag", data.Tag.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *PolicyList) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("action"); value.Exists() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
	if value := res.Get("interfaces"); value.Exists() {
		data.Interfaces = make([]PolicyListInterfaces, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListInterfaces{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Interfaces = append((*parent).Interfaces, data)
			return true
		})
	}
	if value := res.Get("interfaceNames"); value.Exists() {
		data.InterfaceNames = helpers.GetStringList(value.Array())
	} else {
		data.InterfaceNames = types.ListNull(types.StringType)
	}
	if value := res.Get("standardAccessListAddresses"); value.Exists() {
		data.AddressStandardAccessLists = make([]PolicyListAddressStandardAccessLists, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListAddressStandardAccessLists{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).AddressStandardAccessLists = append((*parent).AddressStandardAccessLists, data)
			return true
		})
	}
	if value := res.Get("ipv4PrefixListAddresses"); value.Exists() {
		data.AddressIpv4PrefixLists = make([]PolicyListAddressIpv4PrefixLists, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListAddressIpv4PrefixLists{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).AddressIpv4PrefixLists = append((*parent).AddressIpv4PrefixLists, data)
			return true
		})
	}
	if value := res.Get("standardAccessListNextHops"); value.Exists() {
		data.NextHopStandardAccessLists = make([]PolicyListNextHopStandardAccessLists, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListNextHopStandardAccessLists{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).NextHopStandardAccessLists = append((*parent).NextHopStandardAccessLists, data)
			return true
		})
	}
	if value := res.Get("ipv4PrefixListNexthops"); value.Exists() {
		data.NextHopIpv4PrefixLists = make([]PolicyListNextHopIpv4PrefixLists, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListNextHopIpv4PrefixLists{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).NextHopIpv4PrefixLists = append((*parent).NextHopIpv4PrefixLists, data)
			return true
		})
	}
	if value := res.Get("standardAccessListRouteSources"); value.Exists() {
		data.RouteSourceStandardAccessLists = make([]PolicyListRouteSourceStandardAccessLists, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListRouteSourceStandardAccessLists{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).RouteSourceStandardAccessLists = append((*parent).RouteSourceStandardAccessLists, data)
			return true
		})
	}
	if value := res.Get("ipv4PrefixListRouteSources"); value.Exists() {
		data.RouteSourceIpv4PrefixLists = make([]PolicyListRouteSourceIpv4PrefixLists, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListRouteSourceIpv4PrefixLists{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).RouteSourceIpv4PrefixLists = append((*parent).RouteSourceIpv4PrefixLists, data)
			return true
		})
	}
	if value := res.Get("asPathLists"); value.Exists() {
		data.AsPathLists = make([]PolicyListAsPathLists, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListAsPathLists{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).AsPathLists = append((*parent).AsPathLists, data)
			return true
		})
	}
	if value := res.Get("communityLists"); value.Exists() {
		data.CommunityLists = make([]PolicyListCommunityLists, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListCommunityLists{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).CommunityLists = append((*parent).CommunityLists, data)
			return true
		})
	}
	if value := res.Get("extendedCommunityLists"); value.Exists() {
		data.ExtendedCommunityLists = make([]PolicyListExtendedCommunityLists, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := PolicyListExtendedCommunityLists{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).ExtendedCommunityLists = append((*parent).ExtendedCommunityLists, data)
			return true
		})
	}
	if value := res.Get("matchCommunityExactly"); value.Exists() {
		data.MatchCommunityExactly = types.BoolValue(value.Bool())
	} else {
		data.MatchCommunityExactly = types.BoolNull()
	}
	if value := res.Get("metric"); value.Exists() {
		data.Metric = types.Int64Value(value.Int())
	} else {
		data.Metric = types.Int64Null()
	}
	if value := res.Get("tag"); value.Exists() {
		data.Tag = types.Int64Value(value.Int())
	} else {
		data.Tag = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *PolicyList) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
		data.Action = types.StringValue(value.String())
	} else {
		data.Action = types.StringNull()
	}
	for i := 0; i < len(data.Interfaces); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Interfaces[i].Id.ValueString()}

		parent := &data
		data := (*parent).Interfaces[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing Interfaces[%d] = %+v",
				i,
				(*parent).Interfaces[i],
			))
			(*parent).Interfaces = slices.Delete((*parent).Interfaces, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).Interfaces[i] = data
	}
	if value := res.Get("interfaceNames"); value.Exists() && !data.InterfaceNames.IsNull() {
		data.InterfaceNames = helpers.GetStringList(value.Array())
	} else {
		data.InterfaceNames = types.ListNull(types.StringType)
	}
	for i := 0; i < len(data.AddressStandardAccessLists); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.AddressStandardAccessLists[i].Id.ValueString()}

		parent := &data
		data := (*parent).AddressStandardAccessLists[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("standardAccessListAddresses").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AddressStandardAccessLists[%d] = %+v",
				i,
				(*parent).AddressStandardAccessLists[i],
			))
			(*parent).AddressStandardAccessLists = slices.Delete((*parent).AddressStandardAccessLists, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).AddressStandardAccessLists[i] = data
	}
	for i := 0; i < len(data.AddressIpv4PrefixLists); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.AddressIpv4PrefixLists[i].Id.ValueString()}

		parent := &data
		data := (*parent).AddressIpv4PrefixLists[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing AddressIpv4PrefixLists[%d] = %+v",
				i,
				(*parent).AddressIpv4PrefixLists[i],
			))
			(*parent).AddressIpv4PrefixLists = slices.Delete((*parent).AddressIpv4PrefixLists, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).AddressIpv4PrefixLists[i] = data
	}
	for i := 0; i < len(data.NextHopStandardAccessLists); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.NextHopStandardAccessLists[i].Id.ValueString()}

		parent := &data
		data := (*parent).NextHopStandardAccessLists[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("standardAccessListNextHops").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing NextHopStandardAccessLists[%d] = %+v",
				i,
				(*parent).NextHopStandardAccessLists[i],
			))
			(*parent).NextHopStandardAccessLists = slices.Delete((*parent).NextHopStandardAccessLists, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).NextHopStandardAccessLists[i] = data
	}
	for i := 0; i < len(data.NextHopIpv4PrefixLists); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.NextHopIpv4PrefixLists[i].Id.ValueString()}

		parent := &data
		data := (*parent).NextHopIpv4PrefixLists[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing NextHopIpv4PrefixLists[%d] = %+v",
				i,
				(*parent).NextHopIpv4PrefixLists[i],
			))
			(*parent).NextHopIpv4PrefixLists = slices.Delete((*parent).NextHopIpv4PrefixLists, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).NextHopIpv4PrefixLists[i] = data
	}
	for i := 0; i < len(data.RouteSourceStandardAccessLists); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.RouteSourceStandardAccessLists[i].Id.ValueString()}

		parent := &data
		data := (*parent).RouteSourceStandardAccessLists[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("standardAccessListRouteSources").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing RouteSourceStandardAccessLists[%d] = %+v",
				i,
				(*parent).RouteSourceStandardAccessLists[i],
			))
			(*parent).RouteSourceStandardAccessLists = slices.Delete((*parent).RouteSourceStandardAccessLists, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).RouteSourceStandardAccessLists[i] = data
	}
	for i := 0; i < len(data.RouteSourceIpv4PrefixLists); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.RouteSourceIpv4PrefixLists[i].Id.ValueString()}

		parent := &data
		data := (*parent).RouteSourceIpv4PrefixLists[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing RouteSourceIpv4PrefixLists[%d] = %+v",
				i,
				(*parent).RouteSourceIpv4PrefixLists[i],
			))
			(*parent).RouteSourceIpv4PrefixLists = slices.Delete((*parent).RouteSourceIpv4PrefixLists, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).RouteSourceIpv4PrefixLists[i] = data
	}
	for i := 0; i < len(data.AsPathLists); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.AsPathLists[i].Id.ValueString()}

		parent := &data
		data := (*parent).AsPathLists[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing AsPathLists[%d] = %+v",
				i,
				(*parent).AsPathLists[i],
			))
			(*parent).AsPathLists = slices.Delete((*parent).AsPathLists, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).AsPathLists[i] = data
	}
	for i := 0; i < len(data.CommunityLists); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.CommunityLists[i].Id.ValueString()}

		parent := &data
		data := (*parent).CommunityLists[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing CommunityLists[%d] = %+v",
				i,
				(*parent).CommunityLists[i],
			))
			(*parent).CommunityLists = slices.Delete((*parent).CommunityLists, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).CommunityLists[i] = data
	}
	for i := 0; i < len(data.ExtendedCommunityLists); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.ExtendedCommunityLists[i].Id.ValueString()}

		parent := &data
		data := (*parent).ExtendedCommunityLists[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing ExtendedCommunityLists[%d] = %+v",
				i,
				(*parent).ExtendedCommunityLists[i],
			))
			(*parent).ExtendedCommunityLists = slices.Delete((*parent).ExtendedCommunityLists, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).ExtendedCommunityLists[i] = data
	}
	if value := res.Get("matchCommunityExactly"); value.Exists() && !data.MatchCommunityExactly.IsNull() {
		data.MatchCommunityExactly = types.BoolValue(value.Bool())
	} else {
		data.MatchCommunityExactly = types.BoolNull()
	}
	if value := res.Get("metric"); value.Exists() && !data.Metric.IsNull() {
		data.Metric = types.Int64Value(value.Int())
	} else {
		data.Metric = types.Int64Null()
	}
	if value := res.Get("tag"); value.Exists() && !data.Tag.IsNull() {
		data.Tag = types.Int64Value(value.Int())
	} else {
		data.Tag = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *PolicyList) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
