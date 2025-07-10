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

type SLAMonitor struct {
	Id                 types.String                   `tfsdk:"id"`
	Domain             types.String                   `tfsdk:"domain"`
	Name               types.String                   `tfsdk:"name"`
	Type               types.String                   `tfsdk:"type"`
	Description        types.String                   `tfsdk:"description"`
	SlaMonitorId       types.Int64                    `tfsdk:"sla_monitor_id"`
	Timeout            types.Int64                    `tfsdk:"timeout"`
	Frequency          types.Int64                    `tfsdk:"frequency"`
	Threshold          types.Int64                    `tfsdk:"threshold"`
	DataSize           types.Int64                    `tfsdk:"data_size"`
	Tos                types.Int64                    `tfsdk:"tos"`
	NumberOfPackets    types.Int64                    `tfsdk:"number_of_packets"`
	MonitorAddress     types.String                   `tfsdk:"monitor_address"`
	SelectedInterfaces []SLAMonitorSelectedInterfaces `tfsdk:"selected_interfaces"`
}

type SLAMonitorSelectedInterfaces struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SLAMonitor) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/slamonitors"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SLAMonitor) toBody(ctx context.Context, state SLAMonitor) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "SLAMonitor")
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.SlaMonitorId.IsNull() {
		body, _ = sjson.Set(body, "slaId", data.SlaMonitorId.ValueInt64())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, "timeout", data.Timeout.ValueInt64())
	}
	if !data.Frequency.IsNull() {
		body, _ = sjson.Set(body, "frequency", data.Frequency.ValueInt64())
	}
	if !data.Threshold.IsNull() {
		body, _ = sjson.Set(body, "threshold", data.Threshold.ValueInt64())
	}
	if !data.DataSize.IsNull() {
		body, _ = sjson.Set(body, "dataSize", data.DataSize.ValueInt64())
	}
	if !data.Tos.IsNull() {
		body, _ = sjson.Set(body, "tos", data.Tos.ValueInt64())
	}
	if !data.NumberOfPackets.IsNull() {
		body, _ = sjson.Set(body, "noOfPackets", data.NumberOfPackets.ValueInt64())
	}
	if !data.MonitorAddress.IsNull() {
		body, _ = sjson.Set(body, "monitorAddress", data.MonitorAddress.ValueString())
	}
	if len(data.SelectedInterfaces) > 0 {
		body, _ = sjson.Set(body, "interfaceObjects", []interface{}{})
		for _, item := range data.SelectedInterfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "interfaceObjects.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SLAMonitor) fromBody(ctx context.Context, res gjson.Result) {
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
		data.SelectedInterfaces = make([]SLAMonitorSelectedInterfaces, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := SLAMonitorSelectedInterfaces{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SelectedInterfaces = append((*parent).SelectedInterfaces, data)
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
func (data *SLAMonitor) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SLAMonitor) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
