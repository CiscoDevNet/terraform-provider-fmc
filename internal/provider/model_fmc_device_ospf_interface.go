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
	"strconv"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceOSPFInterface struct {
	Id                         types.String                                `tfsdk:"id"`
	Domain                     types.String                                `tfsdk:"domain"`
	VrfId                      types.String                                `tfsdk:"vrf_id"`
	DeviceId                   types.String                                `tfsdk:"device_id"`
	Type                       types.String                                `tfsdk:"type"`
	InterfaceId                types.String                                `tfsdk:"interface_id"`
	DefaultCost                types.Int64                                 `tfsdk:"default_cost"`
	Priority                   types.Int64                                 `tfsdk:"priority"`
	MtuMissmatchIgnore         types.Bool                                  `tfsdk:"mtu_missmatch_ignore"`
	HelloInterval              types.Int64                                 `tfsdk:"hello_interval"`
	TransmitDelay              types.Int64                                 `tfsdk:"transmit_delay"`
	RetransmitInterval         types.Int64                                 `tfsdk:"retransmit_interval"`
	DeadInterval               types.Int64                                 `tfsdk:"dead_interval"`
	HelloMultiplier            types.Int64                                 `tfsdk:"hello_multiplier"`
	PointToPoint               types.Bool                                  `tfsdk:"point_to_point"`
	Bfd                        types.Bool                                  `tfsdk:"bfd"`
	AuthenticationPassword     types.String                                `tfsdk:"authentication_password"`
	AuthenticationAreaPassword types.String                                `tfsdk:"authentication_area_password"`
	AuthenticationAreaMd5s     []DeviceOSPFInterfaceAuthenticationAreaMd5s `tfsdk:"authentication_area_md5s"`
	AuthenticationMd5s         []DeviceOSPFInterfaceAuthenticationMd5s     `tfsdk:"authentication_md5s"`
	AuthenticationKeyChainId   types.String                                `tfsdk:"authentication_key_chain_id"`
}

type DeviceOSPFInterfaceAuthenticationAreaMd5s struct {
	Id  types.Int64  `tfsdk:"id"`
	Key types.String `tfsdk:"key"`
}

type DeviceOSPFInterfaceAuthenticationMd5s struct {
	Id  types.Int64  `tfsdk:"id"`
	Key types.String `tfsdk:"key"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionDeviceOSPFInterface = version.Must(version.NewVersion("7.6"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceOSPFInterface) getPath() string {
	if data.VrfId.ValueString() != "" {
		return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/virtualrouters/%v/ospfinterface", url.QueryEscape(data.DeviceId.ValueString()), url.QueryEscape(data.VrfId.ValueString()))
	} else {
		return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/routing/ospfinterface", url.QueryEscape(data.DeviceId.ValueString()))
	}
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceOSPFInterface) toBody(ctx context.Context, state DeviceOSPFInterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "deviceInterface.id", data.InterfaceId.ValueString())
	}
	if !data.DefaultCost.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.packetCost", data.DefaultCost.ValueInt64())
	}
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.priority", data.Priority.ValueInt64())
	}
	if !data.MtuMissmatchIgnore.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ignoreMtuMismatch", data.MtuMissmatchIgnore.ValueBool())
	}
	if !data.HelloInterval.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ospfInterval.helloInterval.helloInterval", data.HelloInterval.ValueInt64())
	}
	if !data.TransmitDelay.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.transmitDelay", data.TransmitDelay.ValueInt64())
	}
	if !data.RetransmitInterval.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.retransmitInterval", data.RetransmitInterval.ValueInt64())
	}
	if !data.DeadInterval.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ospfInterval.helloInterval.deadInterval", data.DeadInterval.ValueInt64())
	}
	if !data.HelloMultiplier.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ospfInterval.helloMultiplier.helloMultiplier", data.HelloMultiplier.ValueInt64())
	}
	if !data.PointToPoint.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ptpNonBroadcast", data.PointToPoint.ValueBool())
	}
	if !data.Bfd.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.enableBFD", data.Bfd.ValueBool())
	}
	if !data.AuthenticationPassword.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ospfAuthentication.passwdAuth.authKey", data.AuthenticationPassword.ValueString())
	}
	if !data.AuthenticationAreaPassword.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ospfAuthentication.areaAuth.passwdAuth.authKey", data.AuthenticationAreaPassword.ValueString())
	}
	if len(data.AuthenticationAreaMd5s) > 0 {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ospfAuthentication.areaAuth.md5AuthList", []any{})
		for _, item := range data.AuthenticationAreaMd5s {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "md5KeyId", item.Id.ValueInt64())
			}
			if !item.Key.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "md5Key", item.Key.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ospfProtocolConfiguration.ospfAuthentication.areaAuth.md5AuthList.-1", itemBody)
		}
	}
	if len(data.AuthenticationMd5s) > 0 {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ospfAuthentication.md5AuthList", []any{})
		for _, item := range data.AuthenticationMd5s {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "md5KeyId", item.Id.ValueInt64())
			}
			if !item.Key.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "md5Key", item.Key.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ospfProtocolConfiguration.ospfAuthentication.md5AuthList.-1", itemBody)
		}
	}
	if !data.AuthenticationKeyChainId.IsNull() {
		body, _ = sjson.Set(body, "ospfProtocolConfiguration.ospfAuthentication.keyChain.authKey.id", data.AuthenticationKeyChainId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceOSPFInterface) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("deviceInterface.id"); value.Exists() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if value := res.Get("ospfProtocolConfiguration.packetCost"); value.Exists() {
		data.DefaultCost = types.Int64Value(value.Int())
	} else {
		data.DefaultCost = types.Int64Value(10)
	}
	if value := res.Get("ospfProtocolConfiguration.priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Value(1)
	}
	if value := res.Get("ospfProtocolConfiguration.ignoreMtuMismatch"); value.Exists() {
		data.MtuMissmatchIgnore = types.BoolValue(value.Bool())
	} else {
		data.MtuMissmatchIgnore = types.BoolNull()
	}
	if value := res.Get("ospfProtocolConfiguration.ospfInterval.helloInterval.helloInterval"); value.Exists() {
		data.HelloInterval = types.Int64Value(value.Int())
	} else {
		data.HelloInterval = types.Int64Value(10)
	}
	if value := res.Get("ospfProtocolConfiguration.transmitDelay"); value.Exists() {
		data.TransmitDelay = types.Int64Value(value.Int())
	} else {
		data.TransmitDelay = types.Int64Value(1)
	}
	if value := res.Get("ospfProtocolConfiguration.retransmitInterval"); value.Exists() {
		data.RetransmitInterval = types.Int64Value(value.Int())
	} else {
		data.RetransmitInterval = types.Int64Value(5)
	}
	if value := res.Get("ospfProtocolConfiguration.ospfInterval.helloInterval.deadInterval"); value.Exists() {
		data.DeadInterval = types.Int64Value(value.Int())
	} else {
		data.DeadInterval = types.Int64Value(40)
	}
	if value := res.Get("ospfProtocolConfiguration.ospfInterval.helloMultiplier.helloMultiplier"); value.Exists() {
		data.HelloMultiplier = types.Int64Value(value.Int())
	} else {
		data.HelloMultiplier = types.Int64Null()
	}
	if value := res.Get("ospfProtocolConfiguration.ptpNonBroadcast"); value.Exists() {
		data.PointToPoint = types.BoolValue(value.Bool())
	} else {
		data.PointToPoint = types.BoolNull()
	}
	if value := res.Get("ospfProtocolConfiguration.enableBFD"); value.Exists() {
		data.Bfd = types.BoolValue(value.Bool())
	} else {
		data.Bfd = types.BoolNull()
	}
	if value := res.Get("ospfProtocolConfiguration.ospfAuthentication.passwdAuth.authKey"); value.Exists() {
		data.AuthenticationPassword = types.StringValue(value.String())
	} else {
		data.AuthenticationPassword = types.StringNull()
	}
	if value := res.Get("ospfProtocolConfiguration.ospfAuthentication.areaAuth.passwdAuth.authKey"); value.Exists() {
		data.AuthenticationAreaPassword = types.StringValue(value.String())
	} else {
		data.AuthenticationAreaPassword = types.StringNull()
	}
	if value := res.Get("ospfProtocolConfiguration.ospfAuthentication.areaAuth.md5AuthList"); value.Exists() {
		data.AuthenticationAreaMd5s = make([]DeviceOSPFInterfaceAuthenticationAreaMd5s, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceOSPFInterfaceAuthenticationAreaMd5s{}
			if value := res.Get("md5KeyId"); value.Exists() {
				data.Id = types.Int64Value(value.Int())
			} else {
				data.Id = types.Int64Null()
			}
			if value := res.Get("md5Key"); value.Exists() {
				data.Key = types.StringValue(value.String())
			} else {
				data.Key = types.StringNull()
			}
			(*parent).AuthenticationAreaMd5s = append((*parent).AuthenticationAreaMd5s, data)
			return true
		})
	}
	if value := res.Get("ospfProtocolConfiguration.ospfAuthentication.md5AuthList"); value.Exists() {
		data.AuthenticationMd5s = make([]DeviceOSPFInterfaceAuthenticationMd5s, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceOSPFInterfaceAuthenticationMd5s{}
			if value := res.Get("md5KeyId"); value.Exists() {
				data.Id = types.Int64Value(value.Int())
			} else {
				data.Id = types.Int64Null()
			}
			if value := res.Get("md5Key"); value.Exists() {
				data.Key = types.StringValue(value.String())
			} else {
				data.Key = types.StringNull()
			}
			(*parent).AuthenticationMd5s = append((*parent).AuthenticationMd5s, data)
			return true
		})
	}
	if value := res.Get("ospfProtocolConfiguration.ospfAuthentication.keyChain.authKey.id"); value.Exists() {
		data.AuthenticationKeyChainId = types.StringValue(value.String())
	} else {
		data.AuthenticationKeyChainId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceOSPFInterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("deviceInterface.id"); value.Exists() && !data.InterfaceId.IsNull() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if value := res.Get("ospfProtocolConfiguration.packetCost"); value.Exists() && !data.DefaultCost.IsNull() {
		data.DefaultCost = types.Int64Value(value.Int())
	} else if data.DefaultCost.ValueInt64() != 10 {
		data.DefaultCost = types.Int64Null()
	}
	if value := res.Get("ospfProtocolConfiguration.priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else if data.Priority.ValueInt64() != 1 {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("ospfProtocolConfiguration.ignoreMtuMismatch"); value.Exists() && !data.MtuMissmatchIgnore.IsNull() {
		data.MtuMissmatchIgnore = types.BoolValue(value.Bool())
	} else {
		data.MtuMissmatchIgnore = types.BoolNull()
	}
	if value := res.Get("ospfProtocolConfiguration.ospfInterval.helloInterval.helloInterval"); value.Exists() && !data.HelloInterval.IsNull() {
		data.HelloInterval = types.Int64Value(value.Int())
	} else if data.HelloInterval.ValueInt64() != 10 {
		data.HelloInterval = types.Int64Null()
	}
	if value := res.Get("ospfProtocolConfiguration.transmitDelay"); value.Exists() && !data.TransmitDelay.IsNull() {
		data.TransmitDelay = types.Int64Value(value.Int())
	} else if data.TransmitDelay.ValueInt64() != 1 {
		data.TransmitDelay = types.Int64Null()
	}
	if value := res.Get("ospfProtocolConfiguration.retransmitInterval"); value.Exists() && !data.RetransmitInterval.IsNull() {
		data.RetransmitInterval = types.Int64Value(value.Int())
	} else if data.RetransmitInterval.ValueInt64() != 5 {
		data.RetransmitInterval = types.Int64Null()
	}
	if value := res.Get("ospfProtocolConfiguration.ospfInterval.helloInterval.deadInterval"); value.Exists() && !data.DeadInterval.IsNull() {
		data.DeadInterval = types.Int64Value(value.Int())
	} else if data.DeadInterval.ValueInt64() != 40 {
		data.DeadInterval = types.Int64Null()
	}
	if value := res.Get("ospfProtocolConfiguration.ospfInterval.helloMultiplier.helloMultiplier"); value.Exists() && !data.HelloMultiplier.IsNull() {
		data.HelloMultiplier = types.Int64Value(value.Int())
	} else {
		data.HelloMultiplier = types.Int64Null()
	}
	if value := res.Get("ospfProtocolConfiguration.ptpNonBroadcast"); value.Exists() && !data.PointToPoint.IsNull() {
		data.PointToPoint = types.BoolValue(value.Bool())
	} else {
		data.PointToPoint = types.BoolNull()
	}
	if value := res.Get("ospfProtocolConfiguration.enableBFD"); value.Exists() && !data.Bfd.IsNull() {
		data.Bfd = types.BoolValue(value.Bool())
	} else {
		data.Bfd = types.BoolNull()
	}
	if value := res.Get("ospfProtocolConfiguration.ospfAuthentication.passwdAuth.authKey"); value.Exists() && !data.AuthenticationPassword.IsNull() {
		data.AuthenticationPassword = types.StringValue(value.String())
	} else {
		data.AuthenticationPassword = types.StringNull()
	}
	if value := res.Get("ospfProtocolConfiguration.ospfAuthentication.areaAuth.passwdAuth.authKey"); value.Exists() && !data.AuthenticationAreaPassword.IsNull() {
		data.AuthenticationAreaPassword = types.StringValue(value.String())
	} else {
		data.AuthenticationAreaPassword = types.StringNull()
	}
	for i := 0; i < len(data.AuthenticationAreaMd5s); i++ {
		keys := [...]string{"md5KeyId"}
		keyValues := [...]string{strconv.FormatInt(data.AuthenticationAreaMd5s[i].Id.ValueInt64(), 10)}

		parent := &data
		data := (*parent).AuthenticationAreaMd5s[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ospfProtocolConfiguration.ospfAuthentication.areaAuth.md5AuthList").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AuthenticationAreaMd5s[%d] = %+v",
				i,
				(*parent).AuthenticationAreaMd5s[i],
			))
			(*parent).AuthenticationAreaMd5s = slices.Delete((*parent).AuthenticationAreaMd5s, i, i+1)
			i--

			continue
		}
		if value := res.Get("md5KeyId"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.Int64Value(value.Int())
		} else {
			data.Id = types.Int64Null()
		}
		if value := res.Get("md5Key"); value.Exists() && !data.Key.IsNull() {
			data.Key = types.StringValue(value.String())
		} else {
			data.Key = types.StringNull()
		}
		(*parent).AuthenticationAreaMd5s[i] = data
	}
	for i := 0; i < len(data.AuthenticationMd5s); i++ {
		keys := [...]string{"md5KeyId"}
		keyValues := [...]string{strconv.FormatInt(data.AuthenticationMd5s[i].Id.ValueInt64(), 10)}

		parent := &data
		data := (*parent).AuthenticationMd5s[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ospfProtocolConfiguration.ospfAuthentication.md5AuthList").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AuthenticationMd5s[%d] = %+v",
				i,
				(*parent).AuthenticationMd5s[i],
			))
			(*parent).AuthenticationMd5s = slices.Delete((*parent).AuthenticationMd5s, i, i+1)
			i--

			continue
		}
		if value := res.Get("md5KeyId"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.Int64Value(value.Int())
		} else {
			data.Id = types.Int64Null()
		}
		if value := res.Get("md5Key"); value.Exists() && !data.Key.IsNull() {
			data.Key = types.StringValue(value.String())
		} else {
			data.Key = types.StringNull()
		}
		(*parent).AuthenticationMd5s[i] = data
	}
	if value := res.Get("ospfProtocolConfiguration.ospfAuthentication.keyChain.authKey.id"); value.Exists() && !data.AuthenticationKeyChainId.IsNull() {
		data.AuthenticationKeyChainId = types.StringValue(value.String())
	} else {
		data.AuthenticationKeyChainId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceOSPFInterface) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
