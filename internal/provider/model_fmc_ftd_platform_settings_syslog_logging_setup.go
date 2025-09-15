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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FTDPlatformSettingsSyslogLoggingSetup struct {
	Id                                    types.String                                                    `tfsdk:"id"`
	Domain                                types.String                                                    `tfsdk:"domain"`
	FtdPlatformSettingsId                 types.String                                                    `tfsdk:"ftd_platform_settings_id"`
	Type                                  types.String                                                    `tfsdk:"type"`
	EnableLogging                         types.Bool                                                      `tfsdk:"enable_logging"`
	EnableLoggingOnTheFailoverStandbyUnit types.Bool                                                      `tfsdk:"enable_logging_on_the_failover_standby_unit"`
	EmblemFormat                          types.Bool                                                      `tfsdk:"emblem_format"`
	SendDebugMessagesAsSyslog             types.Bool                                                      `tfsdk:"send_debug_messages_as_syslog"`
	InternalBufferMemorySize              types.Int64                                                     `tfsdk:"internal_buffer_memory_size"`
	FmcLoggingType                        types.String                                                    `tfsdk:"fmc_logging_type"`
	FmcLoggingLevel                       types.String                                                    `tfsdk:"fmc_logging_level"`
	FtpServerHostId                       types.String                                                    `tfsdk:"ftp_server_host_id"`
	FtpServerUsername                     types.String                                                    `tfsdk:"ftp_server_username"`
	FtpServerPassword                     types.String                                                    `tfsdk:"ftp_server_password"`
	FtpServerPath                         types.String                                                    `tfsdk:"ftp_server_path"`
	FtpServerInterfaceGroups              []FTDPlatformSettingsSyslogLoggingSetupFtpServerInterfaceGroups `tfsdk:"ftp_server_interface_groups"`
	Flash                                 types.Bool                                                      `tfsdk:"flash"`
	FlashMaximumSpace                     types.Int64                                                     `tfsdk:"flash_maximum_space"`
	FlashMinimumFreeSpace                 types.Int64                                                     `tfsdk:"flash_minimum_free_space"`
}

type FTDPlatformSettingsSyslogLoggingSetupFtpServerInterfaceGroups struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionCreateFTDPlatformSettingsSyslogLoggingSetup = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsSyslogLoggingSetup) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/syslog/basicloggingsetups", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsSyslogLoggingSetup) toBody(ctx context.Context, state FTDPlatformSettingsSyslogLoggingSetup) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.EnableLogging.IsNull() {
		body, _ = sjson.Set(body, "enableLogging", data.EnableLogging.ValueBool())
	}
	if !data.EnableLoggingOnTheFailoverStandbyUnit.IsNull() {
		body, _ = sjson.Set(body, "enableLoggingFailover", data.EnableLoggingOnTheFailoverStandbyUnit.ValueBool())
	}
	if !data.EmblemFormat.IsNull() {
		body, _ = sjson.Set(body, "enableEMBLEMFormat", data.EmblemFormat.ValueBool())
	}
	if !data.SendDebugMessagesAsSyslog.IsNull() {
		body, _ = sjson.Set(body, "enableDebugMesgAsSyslog", data.SendDebugMessagesAsSyslog.ValueBool())
	}
	if !data.InternalBufferMemorySize.IsNull() {
		body, _ = sjson.Set(body, "internalBuffMemSizeInBytes", data.InternalBufferMemorySize.ValueInt64())
	}
	if !data.FmcLoggingType.IsNull() {
		body, _ = sjson.Set(body, "loggingToFMCType", data.FmcLoggingType.ValueString())
	}
	if !data.FmcLoggingLevel.IsNull() {
		body, _ = sjson.Set(body, "loggingToFMCLogLevel", data.FmcLoggingLevel.ValueString())
	}
	if !data.FtpServerHostId.IsNull() {
		body, _ = sjson.Set(body, "ftpServerInfo.ipAddress.id", data.FtpServerHostId.ValueString())
	}
	if !data.FtpServerUsername.IsNull() {
		body, _ = sjson.Set(body, "ftpServerInfo.userName", data.FtpServerUsername.ValueString())
	}
	if !data.FtpServerPassword.IsNull() {
		body, _ = sjson.Set(body, "ftpServerInfo.password", data.FtpServerPassword.ValueString())
	}
	if !data.FtpServerPath.IsNull() {
		body, _ = sjson.Set(body, "ftpServerInfo.path", data.FtpServerPath.ValueString())
	}
	if len(data.FtpServerInterfaceGroups) > 0 {
		body, _ = sjson.Set(body, "ftpServerInfo.interfaceGroups", []any{})
		for _, item := range data.FtpServerInterfaceGroups {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ftpServerInfo.interfaceGroups.-1", itemBody)
		}
	}
	if !data.Flash.IsNull() {
		body, _ = sjson.Set(body, "enableFlash", data.Flash.ValueBool())
	}
	if !data.FlashMaximumSpace.IsNull() {
		body, _ = sjson.Set(body, "maxFlashForLoggingInKB", data.FlashMaximumSpace.ValueInt64())
	}
	if !data.FlashMinimumFreeSpace.IsNull() {
		body, _ = sjson.Set(body, "minFreePreservedSpaceInKB", data.FlashMinimumFreeSpace.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsSyslogLoggingSetup) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enableLogging"); value.Exists() {
		data.EnableLogging = types.BoolValue(value.Bool())
	} else {
		data.EnableLogging = types.BoolNull()
	}
	if value := res.Get("enableLoggingFailover"); value.Exists() {
		data.EnableLoggingOnTheFailoverStandbyUnit = types.BoolValue(value.Bool())
	} else {
		data.EnableLoggingOnTheFailoverStandbyUnit = types.BoolNull()
	}
	if value := res.Get("enableEMBLEMFormat"); value.Exists() {
		data.EmblemFormat = types.BoolValue(value.Bool())
	} else {
		data.EmblemFormat = types.BoolNull()
	}
	if value := res.Get("enableDebugMesgAsSyslog"); value.Exists() {
		data.SendDebugMessagesAsSyslog = types.BoolValue(value.Bool())
	} else {
		data.SendDebugMessagesAsSyslog = types.BoolNull()
	}
	if value := res.Get("internalBuffMemSizeInBytes"); value.Exists() {
		data.InternalBufferMemorySize = types.Int64Value(value.Int())
	} else {
		data.InternalBufferMemorySize = types.Int64Value(4096)
	}
	if value := res.Get("loggingToFMCType"); value.Exists() {
		data.FmcLoggingType = types.StringValue(value.String())
	} else {
		data.FmcLoggingType = types.StringValue("VPN")
	}
	if value := res.Get("loggingToFMCLogLevel"); value.Exists() {
		data.FmcLoggingLevel = types.StringValue(value.String())
	} else {
		data.FmcLoggingLevel = types.StringValue("ERR")
	}
	if value := res.Get("ftpServerInfo.ipAddress.id"); value.Exists() {
		data.FtpServerHostId = types.StringValue(value.String())
	} else {
		data.FtpServerHostId = types.StringNull()
	}
	if value := res.Get("ftpServerInfo.userName"); value.Exists() {
		data.FtpServerUsername = types.StringValue(value.String())
	} else {
		data.FtpServerUsername = types.StringNull()
	}
	if value := res.Get("ftpServerInfo.path"); value.Exists() {
		data.FtpServerPath = types.StringValue(value.String())
	} else {
		data.FtpServerPath = types.StringNull()
	}
	if value := res.Get("ftpServerInfo.interfaceGroups"); value.Exists() {
		data.FtpServerInterfaceGroups = make([]FTDPlatformSettingsSyslogLoggingSetupFtpServerInterfaceGroups, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsSyslogLoggingSetupFtpServerInterfaceGroups{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).FtpServerInterfaceGroups = append((*parent).FtpServerInterfaceGroups, data)
			return true
		})
	}
	if value := res.Get("enableFlash"); value.Exists() {
		data.Flash = types.BoolValue(value.Bool())
	} else {
		data.Flash = types.BoolNull()
	}
	if value := res.Get("maxFlashForLoggingInKB"); value.Exists() {
		data.FlashMaximumSpace = types.Int64Value(value.Int())
	} else {
		data.FlashMaximumSpace = types.Int64Value(3076)
	}
	if value := res.Get("minFreePreservedSpaceInKB"); value.Exists() {
		data.FlashMinimumFreeSpace = types.Int64Value(value.Int())
	} else {
		data.FlashMinimumFreeSpace = types.Int64Value(1024)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FTDPlatformSettingsSyslogLoggingSetup) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enableLogging"); value.Exists() && !data.EnableLogging.IsNull() {
		data.EnableLogging = types.BoolValue(value.Bool())
	} else {
		data.EnableLogging = types.BoolNull()
	}
	if value := res.Get("enableLoggingFailover"); value.Exists() && !data.EnableLoggingOnTheFailoverStandbyUnit.IsNull() {
		data.EnableLoggingOnTheFailoverStandbyUnit = types.BoolValue(value.Bool())
	} else {
		data.EnableLoggingOnTheFailoverStandbyUnit = types.BoolNull()
	}
	if value := res.Get("enableEMBLEMFormat"); value.Exists() && !data.EmblemFormat.IsNull() {
		data.EmblemFormat = types.BoolValue(value.Bool())
	} else {
		data.EmblemFormat = types.BoolNull()
	}
	if value := res.Get("enableDebugMesgAsSyslog"); value.Exists() && !data.SendDebugMessagesAsSyslog.IsNull() {
		data.SendDebugMessagesAsSyslog = types.BoolValue(value.Bool())
	} else {
		data.SendDebugMessagesAsSyslog = types.BoolNull()
	}
	if value := res.Get("internalBuffMemSizeInBytes"); value.Exists() && !data.InternalBufferMemorySize.IsNull() {
		data.InternalBufferMemorySize = types.Int64Value(value.Int())
	} else if data.InternalBufferMemorySize.ValueInt64() != 4096 {
		data.InternalBufferMemorySize = types.Int64Null()
	}
	if value := res.Get("loggingToFMCType"); value.Exists() && !data.FmcLoggingType.IsNull() {
		data.FmcLoggingType = types.StringValue(value.String())
	} else if data.FmcLoggingType.ValueString() != "VPN" {
		data.FmcLoggingType = types.StringNull()
	}
	if value := res.Get("loggingToFMCLogLevel"); value.Exists() && !data.FmcLoggingLevel.IsNull() {
		data.FmcLoggingLevel = types.StringValue(value.String())
	} else if data.FmcLoggingLevel.ValueString() != "ERR" {
		data.FmcLoggingLevel = types.StringNull()
	}
	if value := res.Get("ftpServerInfo.ipAddress.id"); value.Exists() && !data.FtpServerHostId.IsNull() {
		data.FtpServerHostId = types.StringValue(value.String())
	} else {
		data.FtpServerHostId = types.StringNull()
	}
	if value := res.Get("ftpServerInfo.userName"); value.Exists() && !data.FtpServerUsername.IsNull() {
		data.FtpServerUsername = types.StringValue(value.String())
	} else {
		data.FtpServerUsername = types.StringNull()
	}
	if value := res.Get("ftpServerInfo.path"); value.Exists() && !data.FtpServerPath.IsNull() {
		data.FtpServerPath = types.StringValue(value.String())
	} else {
		data.FtpServerPath = types.StringNull()
	}
	for i := 0; i < len(data.FtpServerInterfaceGroups); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.FtpServerInterfaceGroups[i].Id.ValueString()}

		parent := &data
		data := (*parent).FtpServerInterfaceGroups[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ftpServerInfo.interfaceGroups").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing FtpServerInterfaceGroups[%d] = %+v",
				i,
				(*parent).FtpServerInterfaceGroups[i],
			))
			(*parent).FtpServerInterfaceGroups = slices.Delete((*parent).FtpServerInterfaceGroups, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).FtpServerInterfaceGroups[i] = data
	}
	if value := res.Get("enableFlash"); value.Exists() && !data.Flash.IsNull() {
		data.Flash = types.BoolValue(value.Bool())
	} else {
		data.Flash = types.BoolNull()
	}
	if value := res.Get("maxFlashForLoggingInKB"); value.Exists() && !data.FlashMaximumSpace.IsNull() {
		data.FlashMaximumSpace = types.Int64Value(value.Int())
	} else if data.FlashMaximumSpace.ValueInt64() != 3076 {
		data.FlashMaximumSpace = types.Int64Null()
	}
	if value := res.Get("minFreePreservedSpaceInKB"); value.Exists() && !data.FlashMinimumFreeSpace.IsNull() {
		data.FlashMinimumFreeSpace = types.Int64Value(value.Int())
	} else if data.FlashMinimumFreeSpace.ValueInt64() != 1024 {
		data.FlashMinimumFreeSpace = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsSyslogLoggingSetup) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data FTDPlatformSettingsSyslogLoggingSetup) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	body, _ = sjson.Set(body, "loggingToFMCType", "VPN")
	body, _ = sjson.Set(body, "loggingToFMCLogLevel", "ERR")
	return body
}
