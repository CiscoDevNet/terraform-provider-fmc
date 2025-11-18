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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ExtendedAccessList struct {
	Id      types.String                `tfsdk:"id"`
	Domain  types.String                `tfsdk:"domain"`
	Name    types.String                `tfsdk:"name"`
	Type    types.String                `tfsdk:"type"`
	Entries []ExtendedAccessListEntries `tfsdk:"entries"`
}

type ExtendedAccessListEntries struct {
	Action                     types.String                                          `tfsdk:"action"`
	LogLevel                   types.String                                          `tfsdk:"log_level"`
	Logging                    types.String                                          `tfsdk:"logging"`
	LogInterval                types.Int64                                           `tfsdk:"log_interval"`
	SourceNetworkLiterals      []ExtendedAccessListEntriesSourceNetworkLiterals      `tfsdk:"source_network_literals"`
	DestinationNetworkLiterals []ExtendedAccessListEntriesDestinationNetworkLiterals `tfsdk:"destination_network_literals"`
	SourceNetworkObjects       []ExtendedAccessListEntriesSourceNetworkObjects       `tfsdk:"source_network_objects"`
	DestinationNetworkObjects  []ExtendedAccessListEntriesDestinationNetworkObjects  `tfsdk:"destination_network_objects"`
	SourceSgtObjects           []ExtendedAccessListEntriesSourceSgtObjects           `tfsdk:"source_sgt_objects"`
	SourcePortObjects          []ExtendedAccessListEntriesSourcePortObjects          `tfsdk:"source_port_objects"`
	DestinationPortObjects     []ExtendedAccessListEntriesDestinationPortObjects     `tfsdk:"destination_port_objects"`
	DestinationPortLiterals    []ExtendedAccessListEntriesDestinationPortLiterals    `tfsdk:"destination_port_literals"`
	SourcePortLiterals         []ExtendedAccessListEntriesSourcePortLiterals         `tfsdk:"source_port_literals"`
}

type ExtendedAccessListEntriesSourceNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
	Type  types.String `tfsdk:"type"`
}
type ExtendedAccessListEntriesDestinationNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
	Type  types.String `tfsdk:"type"`
}
type ExtendedAccessListEntriesSourceNetworkObjects struct {
	Id types.String `tfsdk:"id"`
}
type ExtendedAccessListEntriesDestinationNetworkObjects struct {
	Id types.String `tfsdk:"id"`
}
type ExtendedAccessListEntriesSourceSgtObjects struct {
	Id types.String `tfsdk:"id"`
}
type ExtendedAccessListEntriesSourcePortObjects struct {
	Id types.String `tfsdk:"id"`
}
type ExtendedAccessListEntriesDestinationPortObjects struct {
	Id types.String `tfsdk:"id"`
}
type ExtendedAccessListEntriesDestinationPortLiterals struct {
	Type     types.String `tfsdk:"type"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
	IcmpType types.String `tfsdk:"icmp_type"`
	IcmpCode types.String `tfsdk:"icmp_code"`
}
type ExtendedAccessListEntriesSourcePortLiterals struct {
	Type     types.String `tfsdk:"type"`
	Port     types.String `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
	IcmpType types.String `tfsdk:"icmp_type"`
	IcmpCode types.String `tfsdk:"icmp_code"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionCreateExtendedAccessList = version.Must(version.NewVersion("7.2"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ExtendedAccessList) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/extendedaccesslists"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ExtendedAccessList) toBody(ctx context.Context, state ExtendedAccessList) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []any{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.LogLevel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logLevel", item.LogLevel.ValueString())
			}
			if !item.Logging.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logging", item.Logging.ValueString())
			}
			if !item.LogInterval.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logInterval", item.LogInterval.ValueInt64())
			}
			if len(item.SourceNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.literals", []any{})
				for _, childItem := range item.SourceNetworkLiterals {
					itemChildBody := ""
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.literals.-1", itemChildBody)
				}
			}
			if len(item.DestinationNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.literals", []any{})
				for _, childItem := range item.DestinationNetworkLiterals {
					itemChildBody := ""
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.literals.-1", itemChildBody)
				}
			}
			if len(item.SourceNetworkObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.objects", []any{})
				for _, childItem := range item.SourceNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationNetworkObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.objects", []any{})
				for _, childItem := range item.DestinationNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.SourceSgtObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "securityGroupTags.objects", []any{})
				for _, childItem := range item.SourceSgtObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "SecurityGroupTag")
					itemBody, _ = sjson.SetRaw(itemBody, "securityGroupTags.objects.-1", itemChildBody)
				}
			}
			if len(item.SourcePortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourcePorts.objects", []any{})
				for _, childItem := range item.SourcePortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationPortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationPorts.objects", []any{})
				for _, childItem := range item.DestinationPortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationPortLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationPorts.literals", []any{})
				for _, childItem := range item.DestinationPortLiterals {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Port.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "port", childItem.Port.ValueString())
					}
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "protocol", childItem.Protocol.ValueString())
					}
					if !childItem.IcmpType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "icmpType", childItem.IcmpType.ValueString())
					}
					if !childItem.IcmpCode.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "code", childItem.IcmpCode.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.literals.-1", itemChildBody)
				}
			}
			if len(item.SourcePortLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourcePorts.literals", []any{})
				for _, childItem := range item.SourcePortLiterals {
					itemChildBody := ""
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Port.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "port", childItem.Port.ValueString())
					}
					if !childItem.Protocol.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "protocol", childItem.Protocol.ValueString())
					}
					if !childItem.IcmpType.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "icmpType", childItem.IcmpType.ValueString())
					}
					if !childItem.IcmpCode.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "code", childItem.IcmpCode.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.literals.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ExtendedAccessList) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]ExtendedAccessListEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ExtendedAccessListEntries{}
			if value := res.Get("action"); value.Exists() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("logLevel"); value.Exists() {
				data.LogLevel = types.StringValue(value.String())
			} else {
				data.LogLevel = types.StringValue("INFORMATIONAL")
			}
			if value := res.Get("logging"); value.Exists() {
				data.Logging = types.StringValue(value.String())
			} else {
				data.Logging = types.StringNull()
			}
			if value := res.Get("logInterval"); value.Exists() {
				data.LogInterval = types.Int64Value(value.Int())
			} else {
				data.LogInterval = types.Int64Value(300)
			}
			if value := res.Get("sourceNetworks.literals"); value.Exists() {
				data.SourceNetworkLiterals = make([]ExtendedAccessListEntriesSourceNetworkLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ExtendedAccessListEntriesSourceNetworkLiterals{}
					if value := res.Get("value"); value.Exists() {
						data.Value = types.StringValue(value.String())
					} else {
						data.Value = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).SourceNetworkLiterals = append((*parent).SourceNetworkLiterals, data)
					return true
				})
			}
			if value := res.Get("destinationNetworks.literals"); value.Exists() {
				data.DestinationNetworkLiterals = make([]ExtendedAccessListEntriesDestinationNetworkLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ExtendedAccessListEntriesDestinationNetworkLiterals{}
					if value := res.Get("value"); value.Exists() {
						data.Value = types.StringValue(value.String())
					} else {
						data.Value = types.StringNull()
					}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					(*parent).DestinationNetworkLiterals = append((*parent).DestinationNetworkLiterals, data)
					return true
				})
			}
			if value := res.Get("sourceNetworks.objects"); value.Exists() {
				data.SourceNetworkObjects = make([]ExtendedAccessListEntriesSourceNetworkObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ExtendedAccessListEntriesSourceNetworkObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).SourceNetworkObjects = append((*parent).SourceNetworkObjects, data)
					return true
				})
			}
			if value := res.Get("destinationNetworks.objects"); value.Exists() {
				data.DestinationNetworkObjects = make([]ExtendedAccessListEntriesDestinationNetworkObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ExtendedAccessListEntriesDestinationNetworkObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).DestinationNetworkObjects = append((*parent).DestinationNetworkObjects, data)
					return true
				})
			}
			if value := res.Get("securityGroupTags.objects"); value.Exists() {
				data.SourceSgtObjects = make([]ExtendedAccessListEntriesSourceSgtObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ExtendedAccessListEntriesSourceSgtObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).SourceSgtObjects = append((*parent).SourceSgtObjects, data)
					return true
				})
			}
			if value := res.Get("sourcePorts.objects"); value.Exists() {
				data.SourcePortObjects = make([]ExtendedAccessListEntriesSourcePortObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ExtendedAccessListEntriesSourcePortObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).SourcePortObjects = append((*parent).SourcePortObjects, data)
					return true
				})
			}
			if value := res.Get("destinationPorts.objects"); value.Exists() {
				data.DestinationPortObjects = make([]ExtendedAccessListEntriesDestinationPortObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ExtendedAccessListEntriesDestinationPortObjects{}
					if value := res.Get("id"); value.Exists() {
						data.Id = types.StringValue(value.String())
					} else {
						data.Id = types.StringNull()
					}
					(*parent).DestinationPortObjects = append((*parent).DestinationPortObjects, data)
					return true
				})
			}
			if value := res.Get("destinationPorts.literals"); value.Exists() {
				data.DestinationPortLiterals = make([]ExtendedAccessListEntriesDestinationPortLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ExtendedAccessListEntriesDestinationPortLiterals{}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("port"); value.Exists() {
						data.Port = types.StringValue(value.String())
					} else {
						data.Port = types.StringNull()
					}
					if value := res.Get("protocol"); value.Exists() {
						data.Protocol = types.StringValue(value.String())
					} else {
						data.Protocol = types.StringNull()
					}
					if value := res.Get("icmpType"); value.Exists() {
						data.IcmpType = types.StringValue(value.String())
					} else {
						data.IcmpType = types.StringNull()
					}
					if value := res.Get("code"); value.Exists() {
						data.IcmpCode = types.StringValue(value.String())
					} else {
						data.IcmpCode = types.StringNull()
					}
					(*parent).DestinationPortLiterals = append((*parent).DestinationPortLiterals, data)
					return true
				})
			}
			if value := res.Get("sourcePorts.literals"); value.Exists() {
				data.SourcePortLiterals = make([]ExtendedAccessListEntriesSourcePortLiterals, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := ExtendedAccessListEntriesSourcePortLiterals{}
					if value := res.Get("type"); value.Exists() {
						data.Type = types.StringValue(value.String())
					} else {
						data.Type = types.StringNull()
					}
					if value := res.Get("port"); value.Exists() {
						data.Port = types.StringValue(value.String())
					} else {
						data.Port = types.StringNull()
					}
					if value := res.Get("protocol"); value.Exists() {
						data.Protocol = types.StringValue(value.String())
					} else {
						data.Protocol = types.StringNull()
					}
					if value := res.Get("icmpType"); value.Exists() {
						data.IcmpType = types.StringValue(value.String())
					} else {
						data.IcmpType = types.StringNull()
					}
					if value := res.Get("code"); value.Exists() {
						data.IcmpCode = types.StringValue(value.String())
					} else {
						data.IcmpCode = types.StringNull()
					}
					(*parent).SourcePortLiterals = append((*parent).SourcePortLiterals, data)
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
func (data *ExtendedAccessList) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	{
		l := len(res.Get("entries").Array())
		tflog.Debug(ctx, fmt.Sprintf("entries array resizing from %d to %d", len(data.Entries), l))
		for i := len(data.Entries); i < l; i++ {
			data.Entries = append(data.Entries, ExtendedAccessListEntries{})
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
		if value := res.Get("logLevel"); value.Exists() && !data.LogLevel.IsNull() {
			data.LogLevel = types.StringValue(value.String())
		} else if data.LogLevel.ValueString() != "INFORMATIONAL" {
			data.LogLevel = types.StringNull()
		}
		if value := res.Get("logging"); value.Exists() && !data.Logging.IsNull() {
			data.Logging = types.StringValue(value.String())
		} else {
			data.Logging = types.StringNull()
		}
		if value := res.Get("logInterval"); value.Exists() && !data.LogInterval.IsNull() {
			data.LogInterval = types.Int64Value(value.Int())
		} else if data.LogInterval.ValueInt64() != 300 {
			data.LogInterval = types.Int64Null()
		}
		for i := 0; i < len(data.SourceNetworkLiterals); i++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.SourceNetworkLiterals[i].Value.ValueString()}

			parent := &data
			data := (*parent).SourceNetworkLiterals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourceNetworks.literals").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourceNetworkLiterals[%d] = %+v",
					i,
					(*parent).SourceNetworkLiterals[i],
				))
				(*parent).SourceNetworkLiterals = slices.Delete((*parent).SourceNetworkLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).SourceNetworkLiterals[i] = data
		}
		for i := 0; i < len(data.DestinationNetworkLiterals); i++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.DestinationNetworkLiterals[i].Value.ValueString()}

			parent := &data
			data := (*parent).DestinationNetworkLiterals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("destinationNetworks.literals").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationNetworkLiterals[%d] = %+v",
					i,
					(*parent).DestinationNetworkLiterals[i],
				))
				(*parent).DestinationNetworkLiterals = slices.Delete((*parent).DestinationNetworkLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).DestinationNetworkLiterals[i] = data
		}
		for i := 0; i < len(data.SourceNetworkObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceNetworkObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourceNetworkObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourceNetworks.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourceNetworkObjects[%d] = %+v",
					i,
					(*parent).SourceNetworkObjects[i],
				))
				(*parent).SourceNetworkObjects = slices.Delete((*parent).SourceNetworkObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SourceNetworkObjects[i] = data
		}
		for i := 0; i < len(data.DestinationNetworkObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationNetworkObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationNetworkObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("destinationNetworks.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationNetworkObjects[%d] = %+v",
					i,
					(*parent).DestinationNetworkObjects[i],
				))
				(*parent).DestinationNetworkObjects = slices.Delete((*parent).DestinationNetworkObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationNetworkObjects[i] = data
		}
		for i := 0; i < len(data.SourceSgtObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourceSgtObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourceSgtObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("securityGroupTags.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourceSgtObjects[%d] = %+v",
					i,
					(*parent).SourceSgtObjects[i],
				))
				(*parent).SourceSgtObjects = slices.Delete((*parent).SourceSgtObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SourceSgtObjects[i] = data
		}
		for i := 0; i < len(data.SourcePortObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SourcePortObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).SourcePortObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourcePorts.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourcePortObjects[%d] = %+v",
					i,
					(*parent).SourcePortObjects[i],
				))
				(*parent).SourcePortObjects = slices.Delete((*parent).SourcePortObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SourcePortObjects[i] = data
		}
		for i := 0; i < len(data.DestinationPortObjects); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.DestinationPortObjects[i].Id.ValueString()}

			parent := &data
			data := (*parent).DestinationPortObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("destinationPorts.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationPortObjects[%d] = %+v",
					i,
					(*parent).DestinationPortObjects[i],
				))
				(*parent).DestinationPortObjects = slices.Delete((*parent).DestinationPortObjects, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).DestinationPortObjects[i] = data
		}
		for i := 0; i < len(data.DestinationPortLiterals); i++ {
			keys := [...]string{"type", "port", "protocol", "icmpType", "code"}
			keyValues := [...]string{data.DestinationPortLiterals[i].Type.ValueString(), data.DestinationPortLiterals[i].Port.ValueString(), data.DestinationPortLiterals[i].Protocol.ValueString(), data.DestinationPortLiterals[i].IcmpType.ValueString(), data.DestinationPortLiterals[i].IcmpCode.ValueString()}

			parent := &data
			data := (*parent).DestinationPortLiterals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("destinationPorts.literals").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing DestinationPortLiterals[%d] = %+v",
					i,
					(*parent).DestinationPortLiterals[i],
				))
				(*parent).DestinationPortLiterals = slices.Delete((*parent).DestinationPortLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
				data.Port = types.StringValue(value.String())
			} else {
				data.Port = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("icmpType"); value.Exists() && !data.IcmpType.IsNull() {
				data.IcmpType = types.StringValue(value.String())
			} else {
				data.IcmpType = types.StringNull()
			}
			if value := res.Get("code"); value.Exists() && !data.IcmpCode.IsNull() {
				data.IcmpCode = types.StringValue(value.String())
			} else {
				data.IcmpCode = types.StringNull()
			}
			(*parent).DestinationPortLiterals[i] = data
		}
		for i := 0; i < len(data.SourcePortLiterals); i++ {
			keys := [...]string{"type", "port", "protocol", "icmpType", "code"}
			keyValues := [...]string{data.SourcePortLiterals[i].Type.ValueString(), data.SourcePortLiterals[i].Port.ValueString(), data.SourcePortLiterals[i].Protocol.ValueString(), data.SourcePortLiterals[i].IcmpType.ValueString(), data.SourcePortLiterals[i].IcmpCode.ValueString()}

			parent := &data
			data := (*parent).SourcePortLiterals[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("sourcePorts.literals").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SourcePortLiterals[%d] = %+v",
					i,
					(*parent).SourcePortLiterals[i],
				))
				(*parent).SourcePortLiterals = slices.Delete((*parent).SourcePortLiterals, i, i+1)
				i--

				continue
			}
			if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
				data.Port = types.StringValue(value.String())
			} else {
				data.Port = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringNull()
			}
			if value := res.Get("icmpType"); value.Exists() && !data.IcmpType.IsNull() {
				data.IcmpType = types.StringValue(value.String())
			} else {
				data.IcmpType = types.StringNull()
			}
			if value := res.Get("code"); value.Exists() && !data.IcmpCode.IsNull() {
				data.IcmpCode = types.StringValue(value.String())
			} else {
				data.IcmpCode = types.StringNull()
			}
			(*parent).SourcePortLiterals[i] = data
		}
		(*parent).Entries[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ExtendedAccessList) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
