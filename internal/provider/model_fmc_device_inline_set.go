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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceInlineSet struct {
	Id                   types.String                    `tfsdk:"id"`
	Domain               types.String                    `tfsdk:"domain"`
	DeviceId             types.String                    `tfsdk:"device_id"`
	Type                 types.String                    `tfsdk:"type"`
	Name                 types.String                    `tfsdk:"name"`
	Mtu                  types.Int64                     `tfsdk:"mtu"`
	FailSafe             types.Bool                      `tfsdk:"fail_safe"`
	BypassStandby        types.Bool                      `tfsdk:"bypass_standby"`
	BypassForce          types.Bool                      `tfsdk:"bypass_force"`
	InterfacePairs       []DeviceInlineSetInterfacePairs `tfsdk:"interface_pairs"`
	TapMode              types.Bool                      `tfsdk:"tap_mode"`
	PropagateLinkState   types.Bool                      `tfsdk:"propagate_link_state"`
	StrictTcpEnforcement types.Bool                      `tfsdk:"strict_tcp_enforcement"`
	SnortFailOpenBusy    types.Bool                      `tfsdk:"snort_fail_open_busy"`
	SnortFailOpenDown    types.Bool                      `tfsdk:"snort_fail_open_down"`
}

type DeviceInlineSetInterfacePairs struct {
	FirstInterfaceId    types.String `tfsdk:"first_interface_id"`
	FirstInterfaceName  types.String `tfsdk:"first_interface_name"`
	FirstInterfaceType  types.String `tfsdk:"first_interface_type"`
	SecondInterfaceId   types.String `tfsdk:"second_interface_id"`
	SecondInterfaceName types.String `tfsdk:"second_interface_name"`
	SecondInterfaceType types.String `tfsdk:"second_interface_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceInlineSet) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/inlinesets", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceInlineSet) toBody(ctx context.Context, state DeviceInlineSet) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Mtu.IsNull() {
		body, _ = sjson.Set(body, "mtu", data.Mtu.ValueInt64())
	}
	if !data.FailSafe.IsNull() {
		body, _ = sjson.Set(body, "failSafe", data.FailSafe.ValueBool())
	}
	if !data.BypassStandby.IsNull() {
		body, _ = sjson.Set(body, "standBy", data.BypassStandby.ValueBool())
	}
	if !data.BypassForce.IsNull() {
		body, _ = sjson.Set(body, "bypass", data.BypassForce.ValueBool())
	}
	if len(data.InterfacePairs) > 0 {
		var interfacePairsBody strings.Builder
		interfacePairsBody.WriteString("[")
		for _, item := range data.InterfacePairs {
			itemBody := ""
			if !item.FirstInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "first.id", item.FirstInterfaceId.ValueString())
			}
			if !item.FirstInterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "first.name", item.FirstInterfaceName.ValueString())
			}
			if !item.FirstInterfaceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "first.type", item.FirstInterfaceType.ValueString())
			}
			if !item.SecondInterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "second.id", item.SecondInterfaceId.ValueString())
			}
			if !item.SecondInterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "second.name", item.SecondInterfaceName.ValueString())
			}
			if !item.SecondInterfaceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "second.type", item.SecondInterfaceType.ValueString())
			}
			if itemBody != "" {
				if interfacePairsBody.Len() > 1 {
					interfacePairsBody.WriteString(",")
				}
				interfacePairsBody.WriteString(itemBody)
			}
		}
		interfacePairsBody.WriteString("]")
		body, _ = sjson.SetRaw(body, "inlinepairs", interfacePairsBody.String())
	}
	if !data.TapMode.IsNull() {
		body, _ = sjson.Set(body, "tapMode", data.TapMode.ValueBool())
	}
	if !data.PropagateLinkState.IsNull() {
		body, _ = sjson.Set(body, "propogateLinkState", data.PropagateLinkState.ValueBool())
	}
	if !data.StrictTcpEnforcement.IsNull() {
		body, _ = sjson.Set(body, "strictTCPEnforcement", data.StrictTcpEnforcement.ValueBool())
	}
	if !data.SnortFailOpenBusy.IsNull() {
		body, _ = sjson.Set(body, "failOpenSnortBusy", data.SnortFailOpenBusy.ValueBool())
	}
	if !data.SnortFailOpenDown.IsNull() {
		body, _ = sjson.Set(body, "failOpenSnortDown", data.SnortFailOpenDown.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceInlineSet) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("mtu"); value.Exists() {
		data.Mtu = types.Int64Value(value.Int())
	} else {
		data.Mtu = types.Int64Value(1500)
	}
	if value := res.Get("failSafe"); value.Exists() {
		data.FailSafe = types.BoolValue(value.Bool())
	} else {
		data.FailSafe = types.BoolNull()
	}
	if value := res.Get("standBy"); value.Exists() {
		data.BypassStandby = types.BoolValue(value.Bool())
	} else {
		data.BypassStandby = types.BoolNull()
	}
	if value := res.Get("bypass"); value.Exists() {
		data.BypassForce = types.BoolValue(value.Bool())
	} else {
		data.BypassForce = types.BoolNull()
	}
	if value := res.Get("inlinepairs"); value.Exists() {
		data.InterfacePairs = make([]DeviceInlineSetInterfacePairs, 0, int(value.Get("#").Int()))
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceInlineSetInterfacePairs{}
			if value := res.Get("first.id"); value.Exists() {
				data.FirstInterfaceId = types.StringValue(value.String())
			} else {
				data.FirstInterfaceId = types.StringNull()
			}
			if value := res.Get("first.name"); value.Exists() {
				data.FirstInterfaceName = types.StringValue(value.String())
			} else {
				data.FirstInterfaceName = types.StringNull()
			}
			if value := res.Get("first.type"); value.Exists() {
				data.FirstInterfaceType = types.StringValue(value.String())
			} else {
				data.FirstInterfaceType = types.StringNull()
			}
			if value := res.Get("second.id"); value.Exists() {
				data.SecondInterfaceId = types.StringValue(value.String())
			} else {
				data.SecondInterfaceId = types.StringNull()
			}
			if value := res.Get("second.name"); value.Exists() {
				data.SecondInterfaceName = types.StringValue(value.String())
			} else {
				data.SecondInterfaceName = types.StringNull()
			}
			if value := res.Get("second.type"); value.Exists() {
				data.SecondInterfaceType = types.StringValue(value.String())
			} else {
				data.SecondInterfaceType = types.StringNull()
			}
			(*parent).InterfacePairs = append((*parent).InterfacePairs, data)
			return true
		})
	}
	if value := res.Get("tapMode"); value.Exists() {
		data.TapMode = types.BoolValue(value.Bool())
	} else {
		data.TapMode = types.BoolNull()
	}
	if value := res.Get("propogateLinkState"); value.Exists() {
		data.PropagateLinkState = types.BoolValue(value.Bool())
	} else {
		data.PropagateLinkState = types.BoolNull()
	}
	if value := res.Get("strictTCPEnforcement"); value.Exists() {
		data.StrictTcpEnforcement = types.BoolValue(value.Bool())
	} else {
		data.StrictTcpEnforcement = types.BoolNull()
	}
	if value := res.Get("failOpenSnortBusy"); value.Exists() {
		data.SnortFailOpenBusy = types.BoolValue(value.Bool())
	} else {
		data.SnortFailOpenBusy = types.BoolNull()
	}
	if value := res.Get("failOpenSnortDown"); value.Exists() {
		data.SnortFailOpenDown = types.BoolValue(value.Bool())
	} else {
		data.SnortFailOpenDown = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceInlineSet) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("mtu"); value.Exists() && !data.Mtu.IsNull() {
		data.Mtu = types.Int64Value(value.Int())
	} else if data.Mtu.ValueInt64() != 1500 {
		data.Mtu = types.Int64Null()
	}
	if value := res.Get("failSafe"); value.Exists() && !data.FailSafe.IsNull() {
		data.FailSafe = types.BoolValue(value.Bool())
	} else {
		data.FailSafe = types.BoolNull()
	}
	if value := res.Get("standBy"); value.Exists() && !data.BypassStandby.IsNull() {
		data.BypassStandby = types.BoolValue(value.Bool())
	} else {
		data.BypassStandby = types.BoolNull()
	}
	if value := res.Get("bypass"); value.Exists() && !data.BypassForce.IsNull() {
		data.BypassForce = types.BoolValue(value.Bool())
	} else {
		data.BypassForce = types.BoolNull()
	}
	interfacePairsArray := res.Get("inlinepairs")
	for i := 0; i < len(data.InterfacePairs); i++ {
		keys := [...]string{"first.id", "second.id"}
		keyValues := [...]string{data.InterfacePairs[i].FirstInterfaceId.ValueString(), data.InterfacePairs[i].SecondInterfaceId.ValueString()}

		parent := &data
		data := (*parent).InterfacePairs[i]
		var res gjson.Result

		interfacePairsArray.ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing InterfacePairs[%d] = %+v",
				i,
				(*parent).InterfacePairs[i],
			))
			(*parent).InterfacePairs = slices.Delete((*parent).InterfacePairs, i, i+1)
			i--

			continue
		}
		if value := res.Get("first.id"); value.Exists() && !data.FirstInterfaceId.IsNull() {
			data.FirstInterfaceId = types.StringValue(value.String())
		} else {
			data.FirstInterfaceId = types.StringNull()
		}
		if value := res.Get("first.name"); value.Exists() && !data.FirstInterfaceName.IsNull() {
			data.FirstInterfaceName = types.StringValue(value.String())
		} else {
			data.FirstInterfaceName = types.StringNull()
		}
		if value := res.Get("first.type"); value.Exists() && !data.FirstInterfaceType.IsNull() {
			data.FirstInterfaceType = types.StringValue(value.String())
		} else {
			data.FirstInterfaceType = types.StringNull()
		}
		if value := res.Get("second.id"); value.Exists() && !data.SecondInterfaceId.IsNull() {
			data.SecondInterfaceId = types.StringValue(value.String())
		} else {
			data.SecondInterfaceId = types.StringNull()
		}
		if value := res.Get("second.name"); value.Exists() && !data.SecondInterfaceName.IsNull() {
			data.SecondInterfaceName = types.StringValue(value.String())
		} else {
			data.SecondInterfaceName = types.StringNull()
		}
		if value := res.Get("second.type"); value.Exists() && !data.SecondInterfaceType.IsNull() {
			data.SecondInterfaceType = types.StringValue(value.String())
		} else {
			data.SecondInterfaceType = types.StringNull()
		}
		(*parent).InterfacePairs[i] = data
	}
	if value := res.Get("tapMode"); value.Exists() && !data.TapMode.IsNull() {
		data.TapMode = types.BoolValue(value.Bool())
	} else {
		data.TapMode = types.BoolNull()
	}
	if value := res.Get("propogateLinkState"); value.Exists() && !data.PropagateLinkState.IsNull() {
		data.PropagateLinkState = types.BoolValue(value.Bool())
	} else {
		data.PropagateLinkState = types.BoolNull()
	}
	if value := res.Get("strictTCPEnforcement"); value.Exists() && !data.StrictTcpEnforcement.IsNull() {
		data.StrictTcpEnforcement = types.BoolValue(value.Bool())
	} else {
		data.StrictTcpEnforcement = types.BoolNull()
	}
	if value := res.Get("failOpenSnortBusy"); value.Exists() && !data.SnortFailOpenBusy.IsNull() {
		data.SnortFailOpenBusy = types.BoolValue(value.Bool())
	} else {
		data.SnortFailOpenBusy = types.BoolNull()
	}
	if value := res.Get("failOpenSnortDown"); value.Exists() && !data.SnortFailOpenDown.IsNull() {
		data.SnortFailOpenDown = types.BoolValue(value.Bool())
	} else {
		data.SnortFailOpenDown = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceInlineSet) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides
