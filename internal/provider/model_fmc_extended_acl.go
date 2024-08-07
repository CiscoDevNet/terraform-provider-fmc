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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ExtendedACL struct {
	Id          types.String         `tfsdk:"id"`
	Domain      types.String         `tfsdk:"domain"`
	Name        types.String         `tfsdk:"name"`
	Description types.String         `tfsdk:"description"`
	Entries     []ExtendedACLEntries `tfsdk:"entries"`
}

type ExtendedACLEntries struct {
	Action                     types.String                                   `tfsdk:"action"`
	LogLevel                   types.String                                   `tfsdk:"log_level"`
	Logging                    types.String                                   `tfsdk:"logging"`
	LogIntervalSeconds         types.Int64                                    `tfsdk:"log_interval_seconds"`
	SourceNetworkLiterals      []ExtendedACLEntriesSourceNetworkLiterals      `tfsdk:"source_network_literals"`
	DestinationNetworkLiterals []ExtendedACLEntriesDestinationNetworkLiterals `tfsdk:"destination_network_literals"`
	SourceNetworkObjects       []ExtendedACLEntriesSourceNetworkObjects       `tfsdk:"source_network_objects"`
	DestinationNetworkObjects  []ExtendedACLEntriesDestinationNetworkObjects  `tfsdk:"destination_network_objects"`
	SourcePortObjects          []ExtendedACLEntriesSourcePortObjects          `tfsdk:"source_port_objects"`
	DestinationPortObjects     []ExtendedACLEntriesDestinationPortObjects     `tfsdk:"destination_port_objects"`
}

type ExtendedACLEntriesSourceNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
	Type  types.String `tfsdk:"type"`
}
type ExtendedACLEntriesDestinationNetworkLiterals struct {
	Value types.String `tfsdk:"value"`
	Type  types.String `tfsdk:"type"`
}
type ExtendedACLEntriesSourceNetworkObjects struct {
	Id types.String `tfsdk:"id"`
}
type ExtendedACLEntriesDestinationNetworkObjects struct {
	Id types.String `tfsdk:"id"`
}
type ExtendedACLEntriesSourcePortObjects struct {
	Id types.String `tfsdk:"id"`
}
type ExtendedACLEntriesDestinationPortObjects struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ExtendedACL) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/extendedaccesslists"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ExtendedACL) toBody(ctx context.Context, state ExtendedACL) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []interface{}{})
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
			if !item.LogIntervalSeconds.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logInterval", item.LogIntervalSeconds.ValueInt64())
			}
			if len(item.SourceNetworkLiterals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.literals", []interface{}{})
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
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.literals", []interface{}{})
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
				itemBody, _ = sjson.Set(itemBody, "sourceNetworks.objects", []interface{}{})
				for _, childItem := range item.SourceNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourceNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationNetworkObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationNetworks.objects", []interface{}{})
				for _, childItem := range item.DestinationNetworkObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationNetworks.objects.-1", itemChildBody)
				}
			}
			if len(item.SourcePortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "sourcePorts.objects", []interface{}{})
				for _, childItem := range item.SourcePortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "sourcePorts.objects.-1", itemChildBody)
				}
			}
			if len(item.DestinationPortObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "destinationPorts.objects", []interface{}{})
				for _, childItem := range item.DestinationPortObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "destinationPorts.objects.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ExtendedACL) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]ExtendedACLEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := ExtendedACLEntries{}
			if cValue := v.Get("action"); cValue.Exists() {
				item.Action = types.StringValue(cValue.String())
			} else {
				item.Action = types.StringNull()
			}
			if cValue := v.Get("logLevel"); cValue.Exists() {
				item.LogLevel = types.StringValue(cValue.String())
			} else {
				item.LogLevel = types.StringValue("INFORMATIONAL")
			}
			if cValue := v.Get("logging"); cValue.Exists() {
				item.Logging = types.StringValue(cValue.String())
			} else {
				item.Logging = types.StringNull()
			}
			if cValue := v.Get("logInterval"); cValue.Exists() {
				item.LogIntervalSeconds = types.Int64Value(cValue.Int())
			} else {
				item.LogIntervalSeconds = types.Int64Value(300)
			}
			if cValue := v.Get("sourceNetworks.literals"); cValue.Exists() {
				item.SourceNetworkLiterals = make([]ExtendedACLEntriesSourceNetworkLiterals, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ExtendedACLEntriesSourceNetworkLiterals{}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					item.SourceNetworkLiterals = append(item.SourceNetworkLiterals, cItem)
					return true
				})
			}
			if cValue := v.Get("destinationNetworks.literals"); cValue.Exists() {
				item.DestinationNetworkLiterals = make([]ExtendedACLEntriesDestinationNetworkLiterals, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ExtendedACLEntriesDestinationNetworkLiterals{}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					item.DestinationNetworkLiterals = append(item.DestinationNetworkLiterals, cItem)
					return true
				})
			}
			if cValue := v.Get("sourceNetworks.objects"); cValue.Exists() {
				item.SourceNetworkObjects = make([]ExtendedACLEntriesSourceNetworkObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ExtendedACLEntriesSourceNetworkObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.SourceNetworkObjects = append(item.SourceNetworkObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("destinationNetworks.objects"); cValue.Exists() {
				item.DestinationNetworkObjects = make([]ExtendedACLEntriesDestinationNetworkObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ExtendedACLEntriesDestinationNetworkObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.DestinationNetworkObjects = append(item.DestinationNetworkObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("sourcePorts.objects"); cValue.Exists() {
				item.SourcePortObjects = make([]ExtendedACLEntriesSourcePortObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ExtendedACLEntriesSourcePortObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.SourcePortObjects = append(item.SourcePortObjects, cItem)
					return true
				})
			}
			if cValue := v.Get("destinationPorts.objects"); cValue.Exists() {
				item.DestinationPortObjects = make([]ExtendedACLEntriesDestinationPortObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := ExtendedACLEntriesDestinationPortObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					item.DestinationPortObjects = append(item.DestinationPortObjects, cItem)
					return true
				})
			}
			data.Entries = append(data.Entries, item)
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
func (data *ExtendedACL) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	{
		l := len(res.Get("entries").Array())
		tflog.Debug(ctx, fmt.Sprintf("entries array resizing from %d to %d", len(data.Entries), l))
		for i := len(data.Entries); i < l; i++ {
			data.Entries = append(data.Entries, ExtendedACLEntries{})
		}
		if len(data.Entries) > l {
			data.Entries = data.Entries[:l]
		}
	}
	for i := range data.Entries {
		r := res.Get(fmt.Sprintf("entries.%d", i))
		if value := r.Get("action"); value.Exists() && !data.Entries[i].Action.IsNull() {
			data.Entries[i].Action = types.StringValue(value.String())
		} else {
			data.Entries[i].Action = types.StringNull()
		}
		if value := r.Get("logLevel"); value.Exists() && !data.Entries[i].LogLevel.IsNull() {
			data.Entries[i].LogLevel = types.StringValue(value.String())
		} else if data.Entries[i].LogLevel.ValueString() != "INFORMATIONAL" {
			data.Entries[i].LogLevel = types.StringNull()
		}
		if value := r.Get("logging"); value.Exists() && !data.Entries[i].Logging.IsNull() {
			data.Entries[i].Logging = types.StringValue(value.String())
		} else {
			data.Entries[i].Logging = types.StringNull()
		}
		if value := r.Get("logInterval"); value.Exists() && !data.Entries[i].LogIntervalSeconds.IsNull() {
			data.Entries[i].LogIntervalSeconds = types.Int64Value(value.Int())
		} else if data.Entries[i].LogIntervalSeconds.ValueInt64() != 300 {
			data.Entries[i].LogIntervalSeconds = types.Int64Null()
		}
		for ci := 0; ci < len(data.Entries[i].SourceNetworkLiterals); ci++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.Entries[i].SourceNetworkLiterals[ci].Value.ValueString()}

			var cr gjson.Result
			r.Get("sourceNetworks.literals").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.Entries[i].SourceNetworkLiterals[%d] = %+v",
					ci,
					data.Entries[i].SourceNetworkLiterals[ci],
				))
				data.Entries[i].SourceNetworkLiterals = slices.Delete(data.Entries[i].SourceNetworkLiterals, ci, ci+1)
				ci--

				continue
			}
			if value := cr.Get("value"); value.Exists() && !data.Entries[i].SourceNetworkLiterals[ci].Value.IsNull() {
				data.Entries[i].SourceNetworkLiterals[ci].Value = types.StringValue(value.String())
			} else {
				data.Entries[i].SourceNetworkLiterals[ci].Value = types.StringNull()
			}
			if value := cr.Get("type"); value.Exists() && !data.Entries[i].SourceNetworkLiterals[ci].Type.IsNull() {
				data.Entries[i].SourceNetworkLiterals[ci].Type = types.StringValue(value.String())
			} else {
				data.Entries[i].SourceNetworkLiterals[ci].Type = types.StringNull()
			}
		}
		for ci := 0; ci < len(data.Entries[i].DestinationNetworkLiterals); ci++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.Entries[i].DestinationNetworkLiterals[ci].Value.ValueString()}

			var cr gjson.Result
			r.Get("destinationNetworks.literals").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.Entries[i].DestinationNetworkLiterals[%d] = %+v",
					ci,
					data.Entries[i].DestinationNetworkLiterals[ci],
				))
				data.Entries[i].DestinationNetworkLiterals = slices.Delete(data.Entries[i].DestinationNetworkLiterals, ci, ci+1)
				ci--

				continue
			}
			if value := cr.Get("value"); value.Exists() && !data.Entries[i].DestinationNetworkLiterals[ci].Value.IsNull() {
				data.Entries[i].DestinationNetworkLiterals[ci].Value = types.StringValue(value.String())
			} else {
				data.Entries[i].DestinationNetworkLiterals[ci].Value = types.StringNull()
			}
			if value := cr.Get("type"); value.Exists() && !data.Entries[i].DestinationNetworkLiterals[ci].Type.IsNull() {
				data.Entries[i].DestinationNetworkLiterals[ci].Type = types.StringValue(value.String())
			} else {
				data.Entries[i].DestinationNetworkLiterals[ci].Type = types.StringNull()
			}
		}
		for ci := 0; ci < len(data.Entries[i].SourceNetworkObjects); ci++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Entries[i].SourceNetworkObjects[ci].Id.ValueString()}

			var cr gjson.Result
			r.Get("sourceNetworks.objects").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.Entries[i].SourceNetworkObjects[%d] = %+v",
					ci,
					data.Entries[i].SourceNetworkObjects[ci],
				))
				data.Entries[i].SourceNetworkObjects = slices.Delete(data.Entries[i].SourceNetworkObjects, ci, ci+1)
				ci--

				continue
			}
			if value := cr.Get("id"); value.Exists() && !data.Entries[i].SourceNetworkObjects[ci].Id.IsNull() {
				data.Entries[i].SourceNetworkObjects[ci].Id = types.StringValue(value.String())
			} else {
				data.Entries[i].SourceNetworkObjects[ci].Id = types.StringNull()
			}
		}
		for ci := 0; ci < len(data.Entries[i].DestinationNetworkObjects); ci++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Entries[i].DestinationNetworkObjects[ci].Id.ValueString()}

			var cr gjson.Result
			r.Get("destinationNetworks.objects").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.Entries[i].DestinationNetworkObjects[%d] = %+v",
					ci,
					data.Entries[i].DestinationNetworkObjects[ci],
				))
				data.Entries[i].DestinationNetworkObjects = slices.Delete(data.Entries[i].DestinationNetworkObjects, ci, ci+1)
				ci--

				continue
			}
			if value := cr.Get("id"); value.Exists() && !data.Entries[i].DestinationNetworkObjects[ci].Id.IsNull() {
				data.Entries[i].DestinationNetworkObjects[ci].Id = types.StringValue(value.String())
			} else {
				data.Entries[i].DestinationNetworkObjects[ci].Id = types.StringNull()
			}
		}
		for ci := 0; ci < len(data.Entries[i].SourcePortObjects); ci++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Entries[i].SourcePortObjects[ci].Id.ValueString()}

			var cr gjson.Result
			r.Get("sourcePorts.objects").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.Entries[i].SourcePortObjects[%d] = %+v",
					ci,
					data.Entries[i].SourcePortObjects[ci],
				))
				data.Entries[i].SourcePortObjects = slices.Delete(data.Entries[i].SourcePortObjects, ci, ci+1)
				ci--

				continue
			}
			if value := cr.Get("id"); value.Exists() && !data.Entries[i].SourcePortObjects[ci].Id.IsNull() {
				data.Entries[i].SourcePortObjects[ci].Id = types.StringValue(value.String())
			} else {
				data.Entries[i].SourcePortObjects[ci].Id = types.StringNull()
			}
		}
		for ci := 0; ci < len(data.Entries[i].DestinationPortObjects); ci++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Entries[i].DestinationPortObjects[ci].Id.ValueString()}

			var cr gjson.Result
			r.Get("destinationPorts.objects").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.Entries[i].DestinationPortObjects[%d] = %+v",
					ci,
					data.Entries[i].DestinationPortObjects[ci],
				))
				data.Entries[i].DestinationPortObjects = slices.Delete(data.Entries[i].DestinationPortObjects, ci, ci+1)
				ci--

				continue
			}
			if value := cr.Get("id"); value.Exists() && !data.Entries[i].DestinationPortObjects[ci].Id.IsNull() {
				data.Entries[i].DestinationPortObjects[ci].Id = types.StringValue(value.String())
			} else {
				data.Entries[i].DestinationPortObjects[ci].Id = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ExtendedACL) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
