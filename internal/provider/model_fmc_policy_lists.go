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
	"maps"
	"slices"

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type PolicyLists struct {
	Id     types.String                `tfsdk:"id"`
	Domain types.String                `tfsdk:"domain"`
	Items  map[string]PolicyListsItems `tfsdk:"items"`
}

type PolicyListsItems struct {
	Id                             types.String                                     `tfsdk:"id"`
	Type                           types.String                                     `tfsdk:"type"`
	Action                         types.String                                     `tfsdk:"action"`
	Interfaces                     []PolicyListsItemsInterfaces                     `tfsdk:"interfaces"`
	InterfaceNames                 types.List                                       `tfsdk:"interface_names"`
	AddressStandardAccessLists     []PolicyListsItemsAddressStandardAccessLists     `tfsdk:"address_standard_access_lists"`
	AddressIpv4PrefixLists         []PolicyListsItemsAddressIpv4PrefixLists         `tfsdk:"address_ipv4_prefix_lists"`
	NextHopStandardAccessLists     []PolicyListsItemsNextHopStandardAccessLists     `tfsdk:"next_hop_standard_access_lists"`
	NextHopIpv4PrefixLists         []PolicyListsItemsNextHopIpv4PrefixLists         `tfsdk:"next_hop_ipv4_prefix_lists"`
	RouteSourceStandardAccessLists []PolicyListsItemsRouteSourceStandardAccessLists `tfsdk:"route_source_standard_access_lists"`
	RouteSourceIpv4PrefixLists     []PolicyListsItemsRouteSourceIpv4PrefixLists     `tfsdk:"route_source_ipv4_prefix_lists"`
	AsPaths                        []PolicyListsItemsAsPaths                        `tfsdk:"as_paths"`
	CommunityLists                 []PolicyListsItemsCommunityLists                 `tfsdk:"community_lists"`
	ExtendedCommunityLists         []PolicyListsItemsExtendedCommunityLists         `tfsdk:"extended_community_lists"`
	MatchCommunityExactly          types.Bool                                       `tfsdk:"match_community_exactly"`
	Metric                         types.Int64                                      `tfsdk:"metric"`
	Tag                            types.Int64                                      `tfsdk:"tag"`
}

type PolicyListsItemsInterfaces struct {
	Id types.String `tfsdk:"id"`
}
type PolicyListsItemsAddressStandardAccessLists struct {
	Id types.String `tfsdk:"id"`
}
type PolicyListsItemsAddressIpv4PrefixLists struct {
	Id types.String `tfsdk:"id"`
}
type PolicyListsItemsNextHopStandardAccessLists struct {
	Id types.String `tfsdk:"id"`
}
type PolicyListsItemsNextHopIpv4PrefixLists struct {
	Id types.String `tfsdk:"id"`
}
type PolicyListsItemsRouteSourceStandardAccessLists struct {
	Id types.String `tfsdk:"id"`
}
type PolicyListsItemsRouteSourceIpv4PrefixLists struct {
	Id types.String `tfsdk:"id"`
}
type PolicyListsItemsAsPaths struct {
	Id types.String `tfsdk:"id"`
}
type PolicyListsItemsCommunityLists struct {
	Id types.String `tfsdk:"id"`
}
type PolicyListsItemsExtendedCommunityLists struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkCreatePolicyLists = version.Must(version.NewVersion("999"))
var minFMCVersionBulkDeletePolicyLists = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data PolicyLists) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/policylists"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data PolicyLists) toBody(ctx context.Context, state PolicyLists) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "items", []any{})
		for key, item := range data.Items {
			itemBody, _ := sjson.Set("{}", "name", key)
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if len(item.Interfaces) > 0 {
				itemBody, _ = sjson.Set(itemBody, "interfaces", []any{})
				for _, childItem := range item.Interfaces {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interfaces.-1", itemChildBody)
				}
			}
			if !item.InterfaceNames.IsNull() {
				var values []string
				item.InterfaceNames.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "interfaceNames", values)
			}
			if len(item.AddressStandardAccessLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "standardAccessListAddresses", []any{})
				for _, childItem := range item.AddressStandardAccessLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "standardAccessListAddresses.-1", itemChildBody)
				}
			}
			if len(item.AddressIpv4PrefixLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4PrefixListAddresses", []any{})
				for _, childItem := range item.AddressIpv4PrefixLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4PrefixListAddresses.-1", itemChildBody)
				}
			}
			if len(item.NextHopStandardAccessLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "standardAccessListNextHops", []any{})
				for _, childItem := range item.NextHopStandardAccessLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "standardAccessListNextHops.-1", itemChildBody)
				}
			}
			if len(item.NextHopIpv4PrefixLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4PrefixListNexthops", []any{})
				for _, childItem := range item.NextHopIpv4PrefixLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4PrefixListNexthops.-1", itemChildBody)
				}
			}
			if len(item.RouteSourceStandardAccessLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "standardAccessListRouteSources", []any{})
				for _, childItem := range item.RouteSourceStandardAccessLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "standardAccessListRouteSources.-1", itemChildBody)
				}
			}
			if len(item.RouteSourceIpv4PrefixLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4PrefixListRouteSources", []any{})
				for _, childItem := range item.RouteSourceIpv4PrefixLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv4PrefixListRouteSources.-1", itemChildBody)
				}
			}
			if len(item.AsPaths) > 0 {
				itemBody, _ = sjson.Set(itemBody, "asPathLists", []any{})
				for _, childItem := range item.AsPaths {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "asPathLists.-1", itemChildBody)
				}
			}
			if len(item.CommunityLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "communityLists", []any{})
				for _, childItem := range item.CommunityLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "communityLists.-1", itemChildBody)
				}
			}
			if len(item.ExtendedCommunityLists) > 0 {
				itemBody, _ = sjson.Set(itemBody, "extendedCommunityLists", []any{})
				for _, childItem := range item.ExtendedCommunityLists {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "extendedCommunityLists.-1", itemChildBody)
				}
			}
			if !item.MatchCommunityExactly.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "matchCommunityExactly", item.MatchCommunityExactly.ValueBool())
			}
			if !item.Metric.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "metric", item.Metric.ValueInt64())
			}
			if !item.Tag.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tag", item.Tag.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *PolicyLists) fromBody(ctx context.Context, res gjson.Result) {
	// Build lookup map for O(1) access
	itemsByName := make(map[string]gjson.Result)
	res.Get("items").ForEach(func(_, v gjson.Result) bool {
		if name := v.Get("name").String(); name != "" {
			itemsByName[name] = v
		}
		return true
	})
	for k := range data.Items {
		parent := &data
		data := (*parent).Items[k]
		res, found := itemsByName[k]
		if !found {
			tflog.Debug(ctx, fmt.Sprintf("subresource not found, removing: name=%v", k))
			delete((*parent).Items, k)
			continue
		}
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
		if value := res.Get("action"); value.Exists() {
			data.Action = types.StringValue(value.String())
		} else {
			data.Action = types.StringNull()
		}
		if value := res.Get("interfaces"); value.Exists() {
			data.Interfaces = make([]PolicyListsItemsInterfaces, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsInterfaces{}
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
			data.AddressStandardAccessLists = make([]PolicyListsItemsAddressStandardAccessLists, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsAddressStandardAccessLists{}
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
			data.AddressIpv4PrefixLists = make([]PolicyListsItemsAddressIpv4PrefixLists, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsAddressIpv4PrefixLists{}
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
			data.NextHopStandardAccessLists = make([]PolicyListsItemsNextHopStandardAccessLists, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsNextHopStandardAccessLists{}
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
			data.NextHopIpv4PrefixLists = make([]PolicyListsItemsNextHopIpv4PrefixLists, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsNextHopIpv4PrefixLists{}
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
			data.RouteSourceStandardAccessLists = make([]PolicyListsItemsRouteSourceStandardAccessLists, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsRouteSourceStandardAccessLists{}
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
			data.RouteSourceIpv4PrefixLists = make([]PolicyListsItemsRouteSourceIpv4PrefixLists, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsRouteSourceIpv4PrefixLists{}
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
			data.AsPaths = make([]PolicyListsItemsAsPaths, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsAsPaths{}
				if value := res.Get("id"); value.Exists() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).AsPaths = append((*parent).AsPaths, data)
				return true
			})
		}
		if value := res.Get("communityLists"); value.Exists() {
			data.CommunityLists = make([]PolicyListsItemsCommunityLists, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsCommunityLists{}
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
			data.ExtendedCommunityLists = make([]PolicyListsItemsExtendedCommunityLists, 0, len(value.Array()))
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := PolicyListsItemsExtendedCommunityLists{}
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
		(*parent).Items[k] = data
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *PolicyLists) fromBodyPartial(ctx context.Context, res gjson.Result) {
	// Build lookup map for O(1) access by id
	itemsById := make(map[string]gjson.Result)
	res.Get("items").ForEach(func(_, v gjson.Result) bool {
		if id := v.Get("id").String(); id != "" {
			itemsById[id] = v
		}
		return true
	})
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		if data.Id.ValueString() == "" {
			continue
		}
		res, _ := itemsById[data.Id.ValueString()]
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
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
		interfacesArray := res.Get("interfaces")
		for i := 0; i < len(data.Interfaces); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Interfaces[i].Id.ValueString()}

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
		addressStandardAccessListsArray := res.Get("standardAccessListAddresses")
		for i := 0; i < len(data.AddressStandardAccessLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.AddressStandardAccessLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).AddressStandardAccessLists[i]
			var res gjson.Result

			addressStandardAccessListsArray.ForEach(
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
		addressIpv4PrefixListsArray := res.Get("ipv4PrefixListAddresses")
		for i := 0; i < len(data.AddressIpv4PrefixLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.AddressIpv4PrefixLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).AddressIpv4PrefixLists[i]
			var res gjson.Result

			addressIpv4PrefixListsArray.ForEach(
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
		nextHopStandardAccessListsArray := res.Get("standardAccessListNextHops")
		for i := 0; i < len(data.NextHopStandardAccessLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.NextHopStandardAccessLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).NextHopStandardAccessLists[i]
			var res gjson.Result

			nextHopStandardAccessListsArray.ForEach(
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
		nextHopIpv4PrefixListsArray := res.Get("ipv4PrefixListNexthops")
		for i := 0; i < len(data.NextHopIpv4PrefixLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.NextHopIpv4PrefixLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).NextHopIpv4PrefixLists[i]
			var res gjson.Result

			nextHopIpv4PrefixListsArray.ForEach(
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
		routeSourceStandardAccessListsArray := res.Get("standardAccessListRouteSources")
		for i := 0; i < len(data.RouteSourceStandardAccessLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.RouteSourceStandardAccessLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).RouteSourceStandardAccessLists[i]
			var res gjson.Result

			routeSourceStandardAccessListsArray.ForEach(
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
		routeSourceIpv4PrefixListsArray := res.Get("ipv4PrefixListRouteSources")
		for i := 0; i < len(data.RouteSourceIpv4PrefixLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.RouteSourceIpv4PrefixLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).RouteSourceIpv4PrefixLists[i]
			var res gjson.Result

			routeSourceIpv4PrefixListsArray.ForEach(
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
		asPathsArray := res.Get("asPathLists")
		for i := 0; i < len(data.AsPaths); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.AsPaths[i].Id.ValueString()}

			parent := &data
			data := (*parent).AsPaths[i]
			var res gjson.Result

			asPathsArray.ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing AsPaths[%d] = %+v",
					i,
					(*parent).AsPaths[i],
				))
				(*parent).AsPaths = slices.Delete((*parent).AsPaths, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).AsPaths[i] = data
		}
		communityListsArray := res.Get("communityLists")
		for i := 0; i < len(data.CommunityLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.CommunityLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).CommunityLists[i]
			var res gjson.Result

			communityListsArray.ForEach(
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
		extendedCommunityListsArray := res.Get("extendedCommunityLists")
		for i := 0; i < len(data.ExtendedCommunityLists); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.ExtendedCommunityLists[i].Id.ValueString()}

			parent := &data
			data := (*parent).ExtendedCommunityLists[i]
			var res gjson.Result

			extendedCommunityListsArray.ForEach(
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
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *PolicyLists) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	// Build lookup maps for O(1) access
	itemsByName := make(map[string]gjson.Result)
	itemsById := make(map[string]gjson.Result)
	res.Get("items").ForEach(func(_, v gjson.Result) bool {
		if name := v.Get("name").String(); name != "" {
			itemsByName[name] = v
		}
		if id := v.Get("id").String(); id != "" {
			itemsById[id] = v
		}
		return true
	})
	for i, val := range data.Items {
		var r gjson.Result
		if val.Id.IsUnknown() {
			r = itemsByName[i]
		} else if val.Id.ValueString() != "" {
			r = itemsById[val.Id.ValueString()]
		}
		if v := data.Items[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Items[i] = v
		}
		if v := data.Items[i]; v.Type.IsUnknown() {
			if value := r.Get("type"); value.Exists() {
				v.Type = types.StringValue(value.String())
			} else {
				v.Type = types.StringNull()
			}
			data.Items[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

func (data *PolicyLists) Clone() PolicyLists {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data PolicyLists) toBodyNonBulk(ctx context.Context, state PolicyLists) string {
	// This is one-by-one update, so only one element to update is expected
	if len(data.Items) > 1 {
		tflog.Error(ctx, "Found more than one element to change. Only one will be changed.")
	}

	// Utilize existing toBody function
	body := data.toBody(ctx, state)

	// Get first element only
	return gjson.Get(body, "0").String()
}

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
