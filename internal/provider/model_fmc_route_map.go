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
	"strconv"

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
	SequenceNumber             types.Int64                                 `tfsdk:"sequence_number"`
	Action                     types.String                                `tfsdk:"action"`
	SecurityZones              []RouteMapEntriesSecurityZones              `tfsdk:"security_zones"`
	Ipv4AccessListAddresses    []RouteMapEntriesIpv4AccessListAddresses    `tfsdk:"ipv4_access_list_addresses"`
	Ipv4AccessListNextHops     []RouteMapEntriesIpv4AccessListNextHops     `tfsdk:"ipv4_access_list_next_hops"`
	Ipv4AccessListRouteSources []RouteMapEntriesIpv4AccessListRouteSources `tfsdk:"ipv4_access_list_route_sources"`
	Ipv6AccessListAddresses    []RouteMapEntriesIpv6AccessListAddresses    `tfsdk:"ipv6_access_list_addresses"`
	Ipv6AccessListNextHops     []RouteMapEntriesIpv6AccessListNextHops     `tfsdk:"ipv6_access_list_next_hops"`
	Ipv6AccessListRouteSources []RouteMapEntriesIpv6AccessListRouteSources `tfsdk:"ipv6_access_list_route_sources"`
	MetricRouteValues          types.List                                  `tfsdk:"metric_route_values"`
	TagValues                  types.List                                  `tfsdk:"tag_values"`
	RouteTypeExternal1         types.Bool                                  `tfsdk:"route_type_external1"`
	RouteTypeExternal2         types.Bool                                  `tfsdk:"route_type_external2"`
	RouteTypeInternal          types.Bool                                  `tfsdk:"route_type_internal"`
	RouteTypeLocal             types.Bool                                  `tfsdk:"route_type_local"`
	RouteTypeNSSAExternal1     types.Bool                                  `tfsdk:"route_type_n_s_s_a_external1"`
	RouteTypeNSSAExternal2     types.Bool                                  `tfsdk:"route_type_n_s_s_a_external2"`
}

type RouteMapEntriesSecurityZones struct {
	Id types.String `tfsdk:"id"`
}
type RouteMapEntriesIpv4AccessListAddresses struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesIpv4AccessListNextHops struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesIpv4AccessListRouteSources struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesIpv6AccessListAddresses struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesIpv6AccessListNextHops struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type RouteMapEntriesIpv6AccessListRouteSources struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
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
			if !item.SequenceNumber.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "sequence", item.SequenceNumber.ValueInt64())
			}
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if len(item.SecurityZones) > 0 {
				itemBody, _ = sjson.Set(itemBody, "interfaces", []interface{}{})
				for _, childItem := range item.SecurityZones {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interfaces.-1", itemChildBody)
				}
			}
			if len(item.Ipv4AccessListAddresses) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4AccessListAddresses", []interface{}{})
				for _, childItem := range item.Ipv4AccessListAddresses {
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
			if len(item.Ipv4AccessListNextHops) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4AccessListNextHops", []interface{}{})
				for _, childItem := range item.Ipv4AccessListNextHops {
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
			if len(item.Ipv4AccessListRouteSources) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv4AccessListRouteSources", []interface{}{})
				for _, childItem := range item.Ipv4AccessListRouteSources {
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
			if len(item.Ipv6AccessListAddresses) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv6AccessListAddresses", []interface{}{})
				for _, childItem := range item.Ipv6AccessListAddresses {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv6AccessListAddresses.-1", itemChildBody)
				}
			}
			if len(item.Ipv6AccessListNextHops) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv6AccessListNextHops", []interface{}{})
				for _, childItem := range item.Ipv6AccessListNextHops {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv6AccessListNextHops.-1", itemChildBody)
				}
			}
			if len(item.Ipv6AccessListRouteSources) > 0 {
				itemBody, _ = sjson.Set(itemBody, "ipv6AccessListRouteSources", []interface{}{})
				for _, childItem := range item.Ipv6AccessListRouteSources {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "ipv6AccessListRouteSources.-1", itemChildBody)
				}
			}
			if !item.MetricRouteValues.IsNull() {
				var values []int64
				item.MetricRouteValues.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "metricRouteValues", values)
			}
			if !item.TagValues.IsNull() {
				var values []int64
				item.TagValues.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "tagValues", values)
			}
			if !item.RouteTypeExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeExternal1", item.RouteTypeExternal1.ValueBool())
			}
			if !item.RouteTypeExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeExternal2", item.RouteTypeExternal2.ValueBool())
			}
			if !item.RouteTypeInternal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeInternal", item.RouteTypeInternal.ValueBool())
			}
			if !item.RouteTypeLocal.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeLocal", item.RouteTypeLocal.ValueBool())
			}
			if !item.RouteTypeNSSAExternal1.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeNSSAExternal1", item.RouteTypeNSSAExternal1.ValueBool())
			}
			if !item.RouteTypeNSSAExternal2.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "routeTypeNSSAExternal2", item.RouteTypeNSSAExternal2.ValueBool())
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
			if value := res.Get("sequence"); value.Exists() {
				data.SequenceNumber = types.Int64Value(value.Int())
			} else {
				data.SequenceNumber = types.Int64Null()
			}
			if value := res.Get("action"); value.Exists() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("interfaces"); value.Exists() {
				data.SecurityZones = make([]RouteMapEntriesSecurityZones, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesSecurityZones{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).SecurityZones = append((*parent).SecurityZones, data)
					return true
				})
			}
			if value := res.Get("ipv4AccessListAddresses"); value.Exists() {
				data.Ipv4AccessListAddresses = make([]RouteMapEntriesIpv4AccessListAddresses, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesIpv4AccessListAddresses{}
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
					(*parent).Ipv4AccessListAddresses = append((*parent).Ipv4AccessListAddresses, data)
					return true
				})
			}
			if value := res.Get("ipv4AccessListNextHops"); value.Exists() {
				data.Ipv4AccessListNextHops = make([]RouteMapEntriesIpv4AccessListNextHops, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesIpv4AccessListNextHops{}
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
					(*parent).Ipv4AccessListNextHops = append((*parent).Ipv4AccessListNextHops, data)
					return true
				})
			}
			if value := res.Get("ipv4AccessListRouteSources"); value.Exists() {
				data.Ipv4AccessListRouteSources = make([]RouteMapEntriesIpv4AccessListRouteSources, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesIpv4AccessListRouteSources{}
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
					(*parent).Ipv4AccessListRouteSources = append((*parent).Ipv4AccessListRouteSources, data)
					return true
				})
			}
			if value := res.Get("ipv6AccessListAddresses"); value.Exists() {
				data.Ipv6AccessListAddresses = make([]RouteMapEntriesIpv6AccessListAddresses, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesIpv6AccessListAddresses{}
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
					(*parent).Ipv6AccessListAddresses = append((*parent).Ipv6AccessListAddresses, data)
					return true
				})
			}
			if value := res.Get("ipv6AccessListNextHops"); value.Exists() {
				data.Ipv6AccessListNextHops = make([]RouteMapEntriesIpv6AccessListNextHops, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesIpv6AccessListNextHops{}
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
					(*parent).Ipv6AccessListNextHops = append((*parent).Ipv6AccessListNextHops, data)
					return true
				})
			}
			if value := res.Get("ipv6AccessListRouteSources"); value.Exists() {
				data.Ipv6AccessListRouteSources = make([]RouteMapEntriesIpv6AccessListRouteSources, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := RouteMapEntriesIpv6AccessListRouteSources{}
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
					(*parent).Ipv6AccessListRouteSources = append((*parent).Ipv6AccessListRouteSources, data)
					return true
				})
			}
			if value := res.Get("metricRouteValues"); value.Exists() {
				data.MetricRouteValues = helpers.GetInt64List(value.Array())
			} else {
				data.MetricRouteValues = types.ListNull(types.Int64Type)
			}
			if value := res.Get("tagValues"); value.Exists() {
				data.TagValues = helpers.GetInt64List(value.Array())
			} else {
				data.TagValues = types.ListNull(types.Int64Type)
			}
			if value := res.Get("routeTypeExternal1"); value.Exists() {
				data.RouteTypeExternal1 = types.BoolValue(value.Bool())
			} else {
				data.RouteTypeExternal1 = types.BoolNull()
			}
			if value := res.Get("routeTypeExternal2"); value.Exists() {
				data.RouteTypeExternal2 = types.BoolValue(value.Bool())
			} else {
				data.RouteTypeExternal2 = types.BoolNull()
			}
			if value := res.Get("routeTypeInternal"); value.Exists() {
				data.RouteTypeInternal = types.BoolValue(value.Bool())
			} else {
				data.RouteTypeInternal = types.BoolNull()
			}
			if value := res.Get("routeTypeLocal"); value.Exists() {
				data.RouteTypeLocal = types.BoolValue(value.Bool())
			} else {
				data.RouteTypeLocal = types.BoolNull()
			}
			if value := res.Get("routeTypeNSSAExternal1"); value.Exists() {
				data.RouteTypeNSSAExternal1 = types.BoolValue(value.Bool())
			} else {
				data.RouteTypeNSSAExternal1 = types.BoolNull()
			}
			if value := res.Get("routeTypeNSSAExternal2"); value.Exists() {
				data.RouteTypeNSSAExternal2 = types.BoolValue(value.Bool())
			} else {
				data.RouteTypeNSSAExternal2 = types.BoolNull()
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
	for i := 0; i < len(data.Entries); i++ {
		keys := [...]string{"sequence", "action", "routeTypeExternal1", "routeTypeExternal2", "routeTypeInternal", "routeTypeLocal", "routeTypeNSSAExternal1", "routeTypeNSSAExternal2"}
		keyValues := [...]string{strconv.FormatInt(data.Entries[i].SequenceNumber.ValueInt64(), 10), data.Entries[i].Action.ValueString(), strconv.FormatBool(data.Entries[i].RouteTypeExternal1.ValueBool()), strconv.FormatBool(data.Entries[i].RouteTypeExternal2.ValueBool()), strconv.FormatBool(data.Entries[i].RouteTypeInternal.ValueBool()), strconv.FormatBool(data.Entries[i].RouteTypeLocal.ValueBool()), strconv.FormatBool(data.Entries[i].RouteTypeNSSAExternal1.ValueBool()), strconv.FormatBool(data.Entries[i].RouteTypeNSSAExternal2.ValueBool())}

		parent := &data
		data := (*parent).Entries[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("entries").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Entries[%d] = %+v",
				i,
				(*parent).Entries[i],
			))
			(*parent).Entries = slices.Delete((*parent).Entries, i, i+1)
			i--

			continue
		}
		if value := res.Get("sequence"); value.Exists() && !data.SequenceNumber.IsNull() {
			data.SequenceNumber = types.Int64Value(value.Int())
		} else {
			data.SequenceNumber = types.Int64Null()
		}
		if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
			data.Action = types.StringValue(value.String())
		} else {
			data.Action = types.StringNull()
		}
		for i := 0; i < len(data.SecurityZones); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SecurityZones[i].Id.ValueString()}

			parent := &data
			data := (*parent).SecurityZones[i]
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
				tflog.Debug(ctx, fmt.Sprintf("removing SecurityZones[%d] = %+v",
					i,
					(*parent).SecurityZones[i],
				))
				(*parent).SecurityZones = slices.Delete((*parent).SecurityZones, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SecurityZones[i] = data
		}
		for i := 0; i < len(data.Ipv4AccessListAddresses); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Ipv4AccessListAddresses[i].Id.ValueString()}

			parent := &data
			data := (*parent).Ipv4AccessListAddresses[i]
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
				tflog.Debug(ctx, fmt.Sprintf("removing Ipv4AccessListAddresses[%d] = %+v",
					i,
					(*parent).Ipv4AccessListAddresses[i],
				))
				(*parent).Ipv4AccessListAddresses = slices.Delete((*parent).Ipv4AccessListAddresses, i, i+1)
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
			(*parent).Ipv4AccessListAddresses[i] = data
		}
		for i := 0; i < len(data.Ipv4AccessListNextHops); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Ipv4AccessListNextHops[i].Id.ValueString()}

			parent := &data
			data := (*parent).Ipv4AccessListNextHops[i]
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
				tflog.Debug(ctx, fmt.Sprintf("removing Ipv4AccessListNextHops[%d] = %+v",
					i,
					(*parent).Ipv4AccessListNextHops[i],
				))
				(*parent).Ipv4AccessListNextHops = slices.Delete((*parent).Ipv4AccessListNextHops, i, i+1)
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
			(*parent).Ipv4AccessListNextHops[i] = data
		}
		for i := 0; i < len(data.Ipv4AccessListRouteSources); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Ipv4AccessListRouteSources[i].Id.ValueString()}

			parent := &data
			data := (*parent).Ipv4AccessListRouteSources[i]
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
				tflog.Debug(ctx, fmt.Sprintf("removing Ipv4AccessListRouteSources[%d] = %+v",
					i,
					(*parent).Ipv4AccessListRouteSources[i],
				))
				(*parent).Ipv4AccessListRouteSources = slices.Delete((*parent).Ipv4AccessListRouteSources, i, i+1)
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
			(*parent).Ipv4AccessListRouteSources[i] = data
		}
		for i := 0; i < len(data.Ipv6AccessListAddresses); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Ipv6AccessListAddresses[i].Id.ValueString()}

			parent := &data
			data := (*parent).Ipv6AccessListAddresses[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv6AccessListAddresses").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Ipv6AccessListAddresses[%d] = %+v",
					i,
					(*parent).Ipv6AccessListAddresses[i],
				))
				(*parent).Ipv6AccessListAddresses = slices.Delete((*parent).Ipv6AccessListAddresses, i, i+1)
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
			(*parent).Ipv6AccessListAddresses[i] = data
		}
		for i := 0; i < len(data.Ipv6AccessListNextHops); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Ipv6AccessListNextHops[i].Id.ValueString()}

			parent := &data
			data := (*parent).Ipv6AccessListNextHops[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv6AccessListNextHops").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Ipv6AccessListNextHops[%d] = %+v",
					i,
					(*parent).Ipv6AccessListNextHops[i],
				))
				(*parent).Ipv6AccessListNextHops = slices.Delete((*parent).Ipv6AccessListNextHops, i, i+1)
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
			(*parent).Ipv6AccessListNextHops[i] = data
		}
		for i := 0; i < len(data.Ipv6AccessListRouteSources); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Ipv6AccessListRouteSources[i].Id.ValueString()}

			parent := &data
			data := (*parent).Ipv6AccessListRouteSources[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("ipv6AccessListRouteSources").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Ipv6AccessListRouteSources[%d] = %+v",
					i,
					(*parent).Ipv6AccessListRouteSources[i],
				))
				(*parent).Ipv6AccessListRouteSources = slices.Delete((*parent).Ipv6AccessListRouteSources, i, i+1)
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
			(*parent).Ipv6AccessListRouteSources[i] = data
		}
		if value := res.Get("metricRouteValues"); value.Exists() && !data.MetricRouteValues.IsNull() {
			data.MetricRouteValues = helpers.GetInt64List(value.Array())
		} else {
			data.MetricRouteValues = types.ListNull(types.Int64Type)
		}
		if value := res.Get("tagValues"); value.Exists() && !data.TagValues.IsNull() {
			data.TagValues = helpers.GetInt64List(value.Array())
		} else {
			data.TagValues = types.ListNull(types.Int64Type)
		}
		if value := res.Get("routeTypeExternal1"); value.Exists() && !data.RouteTypeExternal1.IsNull() {
			data.RouteTypeExternal1 = types.BoolValue(value.Bool())
		} else {
			data.RouteTypeExternal1 = types.BoolNull()
		}
		if value := res.Get("routeTypeExternal2"); value.Exists() && !data.RouteTypeExternal2.IsNull() {
			data.RouteTypeExternal2 = types.BoolValue(value.Bool())
		} else {
			data.RouteTypeExternal2 = types.BoolNull()
		}
		if value := res.Get("routeTypeInternal"); value.Exists() && !data.RouteTypeInternal.IsNull() {
			data.RouteTypeInternal = types.BoolValue(value.Bool())
		} else {
			data.RouteTypeInternal = types.BoolNull()
		}
		if value := res.Get("routeTypeLocal"); value.Exists() && !data.RouteTypeLocal.IsNull() {
			data.RouteTypeLocal = types.BoolValue(value.Bool())
		} else {
			data.RouteTypeLocal = types.BoolNull()
		}
		if value := res.Get("routeTypeNSSAExternal1"); value.Exists() && !data.RouteTypeNSSAExternal1.IsNull() {
			data.RouteTypeNSSAExternal1 = types.BoolValue(value.Bool())
		} else {
			data.RouteTypeNSSAExternal1 = types.BoolNull()
		}
		if value := res.Get("routeTypeNSSAExternal2"); value.Exists() && !data.RouteTypeNSSAExternal2.IsNull() {
			data.RouteTypeNSSAExternal2 = types.BoolValue(value.Bool())
		} else {
			data.RouteTypeNSSAExternal2 = types.BoolNull()
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
