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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SLAMonitors struct {
	Id     types.String                `tfsdk:"id"`
	Domain types.String                `tfsdk:"domain"`
	Items  map[string]SLAMonitorsItems `tfsdk:"items"`
}

type SLAMonitorsItems struct {
	Id                 types.String                         `tfsdk:"id"`
	Type               types.String                         `tfsdk:"type"`
	Description        types.String                         `tfsdk:"description"`
	SlaMonitorId       types.Int64                          `tfsdk:"sla_monitor_id"`
	Timeout            types.Int64                          `tfsdk:"timeout"`
	Frequency          types.Int64                          `tfsdk:"frequency"`
	Threshold          types.Int64                          `tfsdk:"threshold"`
	DataSize           types.Int64                          `tfsdk:"data_size"`
	Tos                types.Int64                          `tfsdk:"tos"`
	NumberOfPackets    types.Int64                          `tfsdk:"number_of_packets"`
	MonitorAddress     types.String                         `tfsdk:"monitor_address"`
	SelectedInterfaces []SLAMonitorsItemsSelectedInterfaces `tfsdk:"selected_interfaces"`
}

type SLAMonitorsItemsSelectedInterfaces struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkDeleteSLAMonitors = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SLAMonitors) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/slamonitors"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SLAMonitors) toBody(ctx context.Context, state SLAMonitors) string {
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
			itemBody, _ = sjson.Set(itemBody, "type", "SLAMonitor")
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.SlaMonitorId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "slaId", item.SlaMonitorId.ValueInt64())
			}
			if !item.Timeout.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "timeout", item.Timeout.ValueInt64())
			}
			if !item.Frequency.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "frequency", item.Frequency.ValueInt64())
			}
			if !item.Threshold.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "threshold", item.Threshold.ValueInt64())
			}
			if !item.DataSize.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dataSize", item.DataSize.ValueInt64())
			}
			if !item.Tos.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "tos", item.Tos.ValueInt64())
			}
			if !item.NumberOfPackets.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "noOfPackets", item.NumberOfPackets.ValueInt64())
			}
			if !item.MonitorAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "monitorAddress", item.MonitorAddress.ValueString())
			}
			if len(item.SelectedInterfaces) > 0 {
				itemBody, _ = sjson.Set(itemBody, "interfaceObjects", []any{})
				for _, childItem := range item.SelectedInterfaces {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interfaceObjects.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SLAMonitors) fromBody(ctx context.Context, res gjson.Result) {
	for k := range data.Items {
		parent := &data
		data := (*parent).Items[k]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("name").String() == k {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if !res.Exists() {
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
		if value := res.Get("description"); value.Exists() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("slaId"); value.Exists() {
			data.SlaMonitorId = types.Int64Value(value.Int())
		} else {
			data.SlaMonitorId = types.Int64Null()
		}
		if value := res.Get("timeout"); value.Exists() {
			data.Timeout = types.Int64Value(value.Int())
		} else {
			data.Timeout = types.Int64Value(5000)
		}
		if value := res.Get("frequency"); value.Exists() {
			data.Frequency = types.Int64Value(value.Int())
		} else {
			data.Frequency = types.Int64Value(60)
		}
		if value := res.Get("threshold"); value.Exists() {
			data.Threshold = types.Int64Value(value.Int())
		} else {
			data.Threshold = types.Int64Value(5000)
		}
		if value := res.Get("dataSize"); value.Exists() {
			data.DataSize = types.Int64Value(value.Int())
		} else {
			data.DataSize = types.Int64Value(28)
		}
		if value := res.Get("tos"); value.Exists() {
			data.Tos = types.Int64Value(value.Int())
		} else {
			data.Tos = types.Int64Value(0)
		}
		if value := res.Get("noOfPackets"); value.Exists() {
			data.NumberOfPackets = types.Int64Value(value.Int())
		} else {
			data.NumberOfPackets = types.Int64Value(1)
		}
		if value := res.Get("monitorAddress"); value.Exists() {
			data.MonitorAddress = types.StringValue(value.String())
		} else {
			data.MonitorAddress = types.StringNull()
		}
		if value := res.Get("interfaceObjects"); value.Exists() {
			data.SelectedInterfaces = make([]SLAMonitorsItemsSelectedInterfaces, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := SLAMonitorsItemsSelectedInterfaces{}
				if value := res.Get("id"); value.Exists() {
					data.Id = types.StringValue(value.String())
				} else {
					data.Id = types.StringNull()
				}
				(*parent).SelectedInterfaces = append((*parent).SelectedInterfaces, data)
				return true
			})
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
func (data *SLAMonitors) fromBodyPartial(ctx context.Context, res gjson.Result) {
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == data.Id.ValueString() && data.Id.ValueString() != "" {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
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
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("slaId"); value.Exists() && !data.SlaMonitorId.IsNull() {
			data.SlaMonitorId = types.Int64Value(value.Int())
		} else {
			data.SlaMonitorId = types.Int64Null()
		}
		if value := res.Get("timeout"); value.Exists() && !data.Timeout.IsNull() {
			data.Timeout = types.Int64Value(value.Int())
		} else if data.Timeout.ValueInt64() != 5000 {
			data.Timeout = types.Int64Null()
		}
		if value := res.Get("frequency"); value.Exists() && !data.Frequency.IsNull() {
			data.Frequency = types.Int64Value(value.Int())
		} else if data.Frequency.ValueInt64() != 60 {
			data.Frequency = types.Int64Null()
		}
		if value := res.Get("threshold"); value.Exists() && !data.Threshold.IsNull() {
			data.Threshold = types.Int64Value(value.Int())
		} else if data.Threshold.ValueInt64() != 5000 {
			data.Threshold = types.Int64Null()
		}
		if value := res.Get("dataSize"); value.Exists() && !data.DataSize.IsNull() {
			data.DataSize = types.Int64Value(value.Int())
		} else if data.DataSize.ValueInt64() != 28 {
			data.DataSize = types.Int64Null()
		}
		if value := res.Get("tos"); value.Exists() && !data.Tos.IsNull() {
			data.Tos = types.Int64Value(value.Int())
		} else if data.Tos.ValueInt64() != 0 {
			data.Tos = types.Int64Null()
		}
		if value := res.Get("noOfPackets"); value.Exists() && !data.NumberOfPackets.IsNull() {
			data.NumberOfPackets = types.Int64Value(value.Int())
		} else if data.NumberOfPackets.ValueInt64() != 1 {
			data.NumberOfPackets = types.Int64Null()
		}
		if value := res.Get("monitorAddress"); value.Exists() && !data.MonitorAddress.IsNull() {
			data.MonitorAddress = types.StringValue(value.String())
		} else {
			data.MonitorAddress = types.StringNull()
		}
		for i := 0; i < len(data.SelectedInterfaces); i++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.SelectedInterfaces[i].Id.ValueString()}

			parent := &data
			data := (*parent).SelectedInterfaces[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("interfaceObjects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing SelectedInterfaces[%d] = %+v",
					i,
					(*parent).SelectedInterfaces[i],
				))
				(*parent).SelectedInterfaces = slices.Delete((*parent).SelectedInterfaces, i, i+1)
				i--

				continue
			}
			if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SelectedInterfaces[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SLAMonitors) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	for i, val := range data.Items {
		var r gjson.Result
		res.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if val.Id.IsUnknown() {
					if v.Get("name").String() == i {
						r = v
						return false // break ForEach
					}
				} else {
					if v.Get("id").String() == val.Id.ValueString() && val.Id.ValueString() != "" {
						r = v
						return false // break ForEach
					}
				}

				return true
			},
		)
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

func (data *SLAMonitors) Clone() SLAMonitors {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data SLAMonitors) toBodyNonBulk(ctx context.Context, state SLAMonitors) string {
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
