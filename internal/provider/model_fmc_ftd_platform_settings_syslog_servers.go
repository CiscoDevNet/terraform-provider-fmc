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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FTDPlatformSettingsSyslogServers struct {
	Id                                        types.String                                    `tfsdk:"id"`
	Domain                                    types.String                                    `tfsdk:"domain"`
	FtdPlatformSettingsId                     types.String                                    `tfsdk:"ftd_platform_settings_id"`
	Type                                      types.String                                    `tfsdk:"type"`
	AllowUserTrafficWhenTcpSyslogServerIsDown types.Bool                                      `tfsdk:"allow_user_traffic_when_tcp_syslog_server_is_down"`
	MessageQueueSize                          types.Int64                                     `tfsdk:"message_queue_size"`
	SyslogServers                             []FTDPlatformSettingsSyslogServersSyslogServers `tfsdk:"syslog_servers"`
}

type FTDPlatformSettingsSyslogServersSyslogServers struct {
	NetworkObjectId        types.String                                                    `tfsdk:"network_object_id"`
	Protocol               types.String                                                    `tfsdk:"protocol"`
	Port                   types.Int64                                                     `tfsdk:"port"`
	EmblemFormat           types.Bool                                                      `tfsdk:"emblem_format"`
	SecureSyslog           types.Bool                                                      `tfsdk:"secure_syslog"`
	UseManagementInterface types.Bool                                                      `tfsdk:"use_management_interface"`
	InterfaceLiterals      types.Set                                                       `tfsdk:"interface_literals"`
	InterfaceObjects       []FTDPlatformSettingsSyslogServersSyslogServersInterfaceObjects `tfsdk:"interface_objects"`
}

type FTDPlatformSettingsSyslogServersSyslogServersInterfaceObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionFTDPlatformSettingsSyslogServers = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsSyslogServers) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/syslog/servers", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsSyslogServers) toBody(ctx context.Context, state FTDPlatformSettingsSyslogServers) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.AllowUserTrafficWhenTcpSyslogServerIsDown.IsNull() {
		body, _ = sjson.Set(body, "allowUserTrafficWhenDown", data.AllowUserTrafficWhenTcpSyslogServerIsDown.ValueBool())
	}
	if !data.MessageQueueSize.IsNull() {
		body, _ = sjson.Set(body, "messageSizeQueue", data.MessageQueueSize.ValueInt64())
	}
	if len(data.SyslogServers) > 0 {
		body, _ = sjson.Set(body, "servers", []any{})
		for _, item := range data.SyslogServers {
			itemBody := ""
			if !item.NetworkObjectId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipAddress.object.id", item.NetworkObjectId.ValueString())
			}
			if !item.Protocol.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "protocol", item.Protocol.ValueString())
			}
			if !item.Port.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "port", item.Port.ValueInt64())
			}
			if !item.EmblemFormat.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logMessagesInEMBLEM", item.EmblemFormat.ValueBool())
			}
			if !item.SecureSyslog.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "secureSyslog", item.SecureSyslog.ValueBool())
			}
			if !item.UseManagementInterface.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isMgmtReachable", item.UseManagementInterface.ValueBool())
			}
			if !item.InterfaceLiterals.IsNull() {
				var values []string
				item.InterfaceLiterals.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "interfaces.literals", values)
			}
			if len(item.InterfaceObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "interfaces.objects", []any{})
				for _, childItem := range item.InterfaceObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interfaces.objects.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "servers.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsSyslogServers) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("allowUserTrafficWhenDown"); value.Exists() {
		data.AllowUserTrafficWhenTcpSyslogServerIsDown = types.BoolValue(value.Bool())
	} else {
		data.AllowUserTrafficWhenTcpSyslogServerIsDown = types.BoolValue(true)
	}
	if value := res.Get("messageSizeQueue"); value.Exists() {
		data.MessageQueueSize = types.Int64Value(value.Int())
	} else {
		data.MessageQueueSize = types.Int64Value(512)
	}
	if value := res.Get("servers"); value.Exists() {
		data.SyslogServers = make([]FTDPlatformSettingsSyslogServersSyslogServers, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsSyslogServersSyslogServers{}
			if value := res.Get("ipAddress.object.id"); value.Exists() {
				data.NetworkObjectId = types.StringValue(value.String())
			} else {
				data.NetworkObjectId = types.StringNull()
			}
			if value := res.Get("protocol"); value.Exists() {
				data.Protocol = types.StringValue(value.String())
			} else {
				data.Protocol = types.StringValue("TCP")
			}
			if value := res.Get("port"); value.Exists() {
				data.Port = types.Int64Value(value.Int())
			} else {
				data.Port = types.Int64Value(1470)
			}
			if value := res.Get("logMessagesInEMBLEM"); value.Exists() {
				data.EmblemFormat = types.BoolValue(value.Bool())
			} else {
				data.EmblemFormat = types.BoolNull()
			}
			if value := res.Get("secureSyslog"); value.Exists() {
				data.SecureSyslog = types.BoolValue(value.Bool())
			} else {
				data.SecureSyslog = types.BoolNull()
			}
			if value := res.Get("isMgmtReachable"); value.Exists() {
				data.UseManagementInterface = types.BoolValue(value.Bool())
			} else {
				data.UseManagementInterface = types.BoolValue(true)
			}
			if value := res.Get("interfaces.literals"); value.Exists() {
				data.InterfaceLiterals = helpers.GetStringSet(value.Array())
			} else {
				data.InterfaceLiterals = types.SetNull(types.StringType)
			}
			if value := res.Get("interfaces.objects"); value.Exists() {
				data.InterfaceObjects = make([]FTDPlatformSettingsSyslogServersSyslogServersInterfaceObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := FTDPlatformSettingsSyslogServersSyslogServersInterfaceObjects{}
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
					if value := res.Get("name"); value.Exists() {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					(*parent).InterfaceObjects = append((*parent).InterfaceObjects, data)
					return true
				})
			}
			(*parent).SyslogServers = append((*parent).SyslogServers, data)
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
func (data *FTDPlatformSettingsSyslogServers) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("allowUserTrafficWhenDown"); value.Exists() && !data.AllowUserTrafficWhenTcpSyslogServerIsDown.IsNull() {
		data.AllowUserTrafficWhenTcpSyslogServerIsDown = types.BoolValue(value.Bool())
	} else if data.AllowUserTrafficWhenTcpSyslogServerIsDown.ValueBool() != true {
		data.AllowUserTrafficWhenTcpSyslogServerIsDown = types.BoolNull()
	}
	if value := res.Get("messageSizeQueue"); value.Exists() && !data.MessageQueueSize.IsNull() {
		data.MessageQueueSize = types.Int64Value(value.Int())
	} else if data.MessageQueueSize.ValueInt64() != 512 {
		data.MessageQueueSize = types.Int64Null()
	}
	for i := 0; i < len(data.SyslogServers); i++ {
		keys := [...]string{"ipAddress.object.id"}
		keyValues := [...]string{data.SyslogServers[i].NetworkObjectId.ValueString()}

		parent := &data
		data := (*parent).SyslogServers[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("servers").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing SyslogServers[%d] = %+v",
				i,
				(*parent).SyslogServers[i],
			))
			(*parent).SyslogServers = slices.Delete((*parent).SyslogServers, i, i+1)
			i--

			continue
		}
		if value := res.Get("ipAddress.object.id"); value.Exists() && !data.NetworkObjectId.IsNull() {
			data.NetworkObjectId = types.StringValue(value.String())
		} else {
			data.NetworkObjectId = types.StringNull()
		}
		if value := res.Get("protocol"); value.Exists() && !data.Protocol.IsNull() {
			data.Protocol = types.StringValue(value.String())
		} else if data.Protocol.ValueString() != "TCP" {
			data.Protocol = types.StringNull()
		}
		if value := res.Get("port"); value.Exists() && !data.Port.IsNull() {
			data.Port = types.Int64Value(value.Int())
		} else if data.Port.ValueInt64() != 1470 {
			data.Port = types.Int64Null()
		}
		if value := res.Get("logMessagesInEMBLEM"); value.Exists() && !data.EmblemFormat.IsNull() {
			data.EmblemFormat = types.BoolValue(value.Bool())
		} else {
			data.EmblemFormat = types.BoolNull()
		}
		if value := res.Get("secureSyslog"); value.Exists() && !data.SecureSyslog.IsNull() {
			data.SecureSyslog = types.BoolValue(value.Bool())
		} else {
			data.SecureSyslog = types.BoolNull()
		}
		if value := res.Get("isMgmtReachable"); value.Exists() && !data.UseManagementInterface.IsNull() {
			data.UseManagementInterface = types.BoolValue(value.Bool())
		} else if data.UseManagementInterface.ValueBool() != true {
			data.UseManagementInterface = types.BoolNull()
		}
		if value := res.Get("interfaces.literals"); value.Exists() && !data.InterfaceLiterals.IsNull() {
			data.InterfaceLiterals = helpers.GetStringSet(value.Array())
		} else {
			data.InterfaceLiterals = types.SetNull(types.StringType)
		}
		for i := 0; i < len(data.InterfaceObjects); i++ {
			keys := [...]string{"id", "type", "name"}
			keyValues := [...]string{data.InterfaceObjects[i].Id.ValueString(), data.InterfaceObjects[i].Type.ValueString(), data.InterfaceObjects[i].Name.ValueString()}

			parent := &data
			data := (*parent).InterfaceObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("interfaces.objects").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing InterfaceObjects[%d] = %+v",
					i,
					(*parent).InterfaceObjects[i],
				))
				(*parent).InterfaceObjects = slices.Delete((*parent).InterfaceObjects, i, i+1)
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
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).InterfaceObjects[i] = data
		}
		(*parent).SyslogServers[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsSyslogServers) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data FTDPlatformSettingsSyslogServers) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	return body
}

// End of section. //template:end toBodyPutDelete
