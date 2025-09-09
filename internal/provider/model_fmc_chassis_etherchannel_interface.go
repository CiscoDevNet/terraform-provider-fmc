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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ChassisEtherChannelInterface struct {
	Id                 types.String                                     `tfsdk:"id"`
	Domain             types.String                                     `tfsdk:"domain"`
	ChassisId          types.String                                     `tfsdk:"chassis_id"`
	Type               types.String                                     `tfsdk:"type"`
	Name               types.String                                     `tfsdk:"name"`
	EtherChannelId     types.Int64                                      `tfsdk:"ether_channel_id"`
	PortType           types.String                                     `tfsdk:"port_type"`
	AdminState         types.String                                     `tfsdk:"admin_state"`
	SelectedInterfaces []ChassisEtherChannelInterfaceSelectedInterfaces `tfsdk:"selected_interfaces"`
	AutoNegotiation    types.Bool                                       `tfsdk:"auto_negotiation"`
	Duplex             types.String                                     `tfsdk:"duplex"`
	Speed              types.String                                     `tfsdk:"speed"`
	LacpMode           types.String                                     `tfsdk:"lacp_mode"`
	LacpRate           types.String                                     `tfsdk:"lacp_rate"`
}

type ChassisEtherChannelInterfaceSelectedInterfaces struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ChassisEtherChannelInterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/chassis/fmcmanagedchassis/%v/etherchannelinterfaces", url.QueryEscape(data.ChassisId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ChassisEtherChannelInterface) toBody(ctx context.Context, state ChassisEtherChannelInterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "EtherChannelInterface")
	if !data.EtherChannelId.IsNull() {
		body, _ = sjson.Set(body, "etherChannelId", data.EtherChannelId.ValueInt64())
	}
	if !data.PortType.IsNull() {
		body, _ = sjson.Set(body, "portType", data.PortType.ValueString())
	}
	if !data.AdminState.IsNull() {
		body, _ = sjson.Set(body, "adminState", data.AdminState.ValueString())
	}
	if len(data.SelectedInterfaces) > 0 {
		body, _ = sjson.Set(body, "selectedInterfaces", []any{})
		for _, item := range data.SelectedInterfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "selectedInterfaces.-1", itemBody)
		}
	}
	if !data.AutoNegotiation.IsNull() {
		body, _ = sjson.Set(body, "hardware.autoNegState", data.AutoNegotiation.ValueBool())
	}
	if !data.Duplex.IsNull() {
		body, _ = sjson.Set(body, "hardware.duplex", data.Duplex.ValueString())
	}
	if !data.Speed.IsNull() {
		body, _ = sjson.Set(body, "hardware.speed", data.Speed.ValueString())
	}
	if !data.LacpMode.IsNull() {
		body, _ = sjson.Set(body, "lacpMode", data.LacpMode.ValueString())
	}
	if !data.LacpRate.IsNull() {
		body, _ = sjson.Set(body, "lacpRate", data.LacpRate.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ChassisEtherChannelInterface) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("etherChannelId"); value.Exists() {
		data.EtherChannelId = types.Int64Value(value.Int())
	} else {
		data.EtherChannelId = types.Int64Null()
	}
	if value := res.Get("portType"); value.Exists() {
		data.PortType = types.StringValue(value.String())
	} else {
		data.PortType = types.StringNull()
	}
	if value := res.Get("adminState"); value.Exists() {
		data.AdminState = types.StringValue(value.String())
	} else {
		data.AdminState = types.StringValue("ENABLED")
	}
	if value := res.Get("selectedInterfaces"); value.Exists() {
		data.SelectedInterfaces = make([]ChassisEtherChannelInterfaceSelectedInterfaces, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ChassisEtherChannelInterfaceSelectedInterfaces{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).SelectedInterfaces = append((*parent).SelectedInterfaces, data)
			return true
		})
	}
	if value := res.Get("hardware.autoNegState"); value.Exists() {
		data.AutoNegotiation = types.BoolValue(value.Bool())
	} else {
		data.AutoNegotiation = types.BoolNull()
	}
	if value := res.Get("hardware.duplex"); value.Exists() {
		data.Duplex = types.StringValue(value.String())
	} else {
		data.Duplex = types.StringNull()
	}
	if value := res.Get("hardware.speed"); value.Exists() {
		data.Speed = types.StringValue(value.String())
	} else {
		data.Speed = types.StringNull()
	}
	if value := res.Get("lacpMode"); value.Exists() {
		data.LacpMode = types.StringValue(value.String())
	} else {
		data.LacpMode = types.StringNull()
	}
	if value := res.Get("lacpRate"); value.Exists() {
		data.LacpRate = types.StringValue(value.String())
	} else {
		data.LacpRate = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ChassisEtherChannelInterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("etherChannelId"); value.Exists() && !data.EtherChannelId.IsNull() {
		data.EtherChannelId = types.Int64Value(value.Int())
	} else {
		data.EtherChannelId = types.Int64Null()
	}
	if value := res.Get("portType"); value.Exists() && !data.PortType.IsNull() {
		data.PortType = types.StringValue(value.String())
	} else {
		data.PortType = types.StringNull()
	}
	if value := res.Get("adminState"); value.Exists() && !data.AdminState.IsNull() {
		data.AdminState = types.StringValue(value.String())
	} else if data.AdminState.ValueString() != "ENABLED" {
		data.AdminState = types.StringNull()
	}
	for i := 0; i < len(data.SelectedInterfaces); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.SelectedInterfaces[i].Id.ValueString()}

		parent := &data
		data := (*parent).SelectedInterfaces[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("selectedInterfaces").ForEach(
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
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).SelectedInterfaces[i] = data
	}
	if value := res.Get("hardware.autoNegState"); value.Exists() && !data.AutoNegotiation.IsNull() {
		data.AutoNegotiation = types.BoolValue(value.Bool())
	} else {
		data.AutoNegotiation = types.BoolNull()
	}
	if value := res.Get("hardware.duplex"); value.Exists() && !data.Duplex.IsNull() {
		data.Duplex = types.StringValue(value.String())
	} else {
		data.Duplex = types.StringNull()
	}
	if value := res.Get("hardware.speed"); value.Exists() && !data.Speed.IsNull() {
		data.Speed = types.StringValue(value.String())
	} else {
		data.Speed = types.StringNull()
	}
	if value := res.Get("lacpMode"); value.Exists() && !data.LacpMode.IsNull() {
		data.LacpMode = types.StringValue(value.String())
	} else {
		data.LacpMode = types.StringNull()
	}
	if value := res.Get("lacpRate"); value.Exists() && !data.LacpRate.IsNull() {
		data.LacpRate = types.StringValue(value.String())
	} else {
		data.LacpRate = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ChassisEtherChannelInterface) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.Name.IsUnknown() {
		if value := res.Get("name"); value.Exists() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
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
