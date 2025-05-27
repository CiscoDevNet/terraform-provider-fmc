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
		keys := [...]string{"sequence", "action"}
		keyValues := [...]string{strconv.FormatInt(data.Entries[i].SequenceNumber.ValueInt64(), 10), data.Entries[i].Action.ValueString()}

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
