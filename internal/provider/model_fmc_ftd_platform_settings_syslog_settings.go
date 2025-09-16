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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FTDPlatformSettingsSyslogSettings struct {
	Id                            types.String `tfsdk:"id"`
	Domain                        types.String `tfsdk:"domain"`
	FtdPlatformSettingsId         types.String `tfsdk:"ftd_platform_settings_id"`
	Type                          types.String `tfsdk:"type"`
	Facility                      types.String `tfsdk:"facility"`
	TimestampFormat               types.String `tfsdk:"timestamp_format"`
	DeviceIdType                  types.String `tfsdk:"device_id_type"`
	DeviceIdUserDefinedId         types.String `tfsdk:"device_id_user_defined_id"`
	DeviceIdInterfaceId           types.String `tfsdk:"device_id_interface_id"`
	AllSyslogMessages             types.Bool   `tfsdk:"all_syslog_messages"`
	AllSyslogMessagesLoggingLevel types.String `tfsdk:"all_syslog_messages_logging_level"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionFTDPlatformSettingsSyslogSettings = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsSyslogSettings) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/syslog/syslogsettings", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsSyslogSettings) toBody(ctx context.Context, state FTDPlatformSettingsSyslogSettings) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Facility.IsNull() {
		body, _ = sjson.Set(body, "facility", data.Facility.ValueString())
	}
	if !data.TimestampFormat.IsNull() {
		body, _ = sjson.Set(body, "timeStampFormat", data.TimestampFormat.ValueString())
	}
	if !data.DeviceIdType.IsNull() {
		body, _ = sjson.Set(body, "syslogDeviceIdConfig.deviceIdConfigType", data.DeviceIdType.ValueString())
	}
	if !data.DeviceIdUserDefinedId.IsNull() {
		body, _ = sjson.Set(body, "syslogDeviceIdConfig.userDefinedId", data.DeviceIdUserDefinedId.ValueString())
	}
	if !data.DeviceIdInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "syslogDeviceIdConfig.interface.id", data.DeviceIdInterfaceId.ValueString())
	}
	if !data.AllSyslogMessages.IsNull() {
		body, _ = sjson.Set(body, "syslogEnableAll.enable", data.AllSyslogMessages.ValueBool())
	}
	if !data.AllSyslogMessagesLoggingLevel.IsNull() {
		body, _ = sjson.Set(body, "syslogEnableAll.loggingLevel", data.AllSyslogMessagesLoggingLevel.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsSyslogSettings) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("facility"); value.Exists() {
		data.Facility = types.StringValue(value.String())
	} else {
		data.Facility = types.StringValue("LOCAL4")
	}
	if value := res.Get("timeStampFormat"); value.Exists() {
		data.TimestampFormat = types.StringValue(value.String())
	} else {
		data.TimestampFormat = types.StringNull()
	}
	if value := res.Get("syslogDeviceIdConfig.deviceIdConfigType"); value.Exists() {
		data.DeviceIdType = types.StringValue(value.String())
	} else {
		data.DeviceIdType = types.StringNull()
	}
	if value := res.Get("syslogDeviceIdConfig.userDefinedId"); value.Exists() {
		data.DeviceIdUserDefinedId = types.StringValue(value.String())
	} else {
		data.DeviceIdUserDefinedId = types.StringNull()
	}
	if value := res.Get("syslogDeviceIdConfig.interface.id"); value.Exists() {
		data.DeviceIdInterfaceId = types.StringValue(value.String())
	} else {
		data.DeviceIdInterfaceId = types.StringNull()
	}
	if value := res.Get("syslogEnableAll.enable"); value.Exists() {
		data.AllSyslogMessages = types.BoolValue(value.Bool())
	} else {
		data.AllSyslogMessages = types.BoolNull()
	}
	if value := res.Get("syslogEnableAll.loggingLevel"); value.Exists() {
		data.AllSyslogMessagesLoggingLevel = types.StringValue(value.String())
	} else {
		data.AllSyslogMessagesLoggingLevel = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FTDPlatformSettingsSyslogSettings) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("facility"); value.Exists() && !data.Facility.IsNull() {
		data.Facility = types.StringValue(value.String())
	} else if data.Facility.ValueString() != "LOCAL4" {
		data.Facility = types.StringNull()
	}
	if value := res.Get("timeStampFormat"); value.Exists() && !data.TimestampFormat.IsNull() {
		data.TimestampFormat = types.StringValue(value.String())
	} else {
		data.TimestampFormat = types.StringNull()
	}
	if value := res.Get("syslogDeviceIdConfig.deviceIdConfigType"); value.Exists() && !data.DeviceIdType.IsNull() {
		data.DeviceIdType = types.StringValue(value.String())
	} else {
		data.DeviceIdType = types.StringNull()
	}
	if value := res.Get("syslogDeviceIdConfig.userDefinedId"); value.Exists() && !data.DeviceIdUserDefinedId.IsNull() {
		data.DeviceIdUserDefinedId = types.StringValue(value.String())
	} else {
		data.DeviceIdUserDefinedId = types.StringNull()
	}
	if value := res.Get("syslogDeviceIdConfig.interface.id"); value.Exists() && !data.DeviceIdInterfaceId.IsNull() {
		data.DeviceIdInterfaceId = types.StringValue(value.String())
	} else {
		data.DeviceIdInterfaceId = types.StringNull()
	}
	if value := res.Get("syslogEnableAll.enable"); value.Exists() && !data.AllSyslogMessages.IsNull() {
		data.AllSyslogMessages = types.BoolValue(value.Bool())
	} else {
		data.AllSyslogMessages = types.BoolNull()
	}
	if value := res.Get("syslogEnableAll.loggingLevel"); value.Exists() && !data.AllSyslogMessagesLoggingLevel.IsNull() {
		data.AllSyslogMessagesLoggingLevel = types.StringValue(value.String())
	} else {
		data.AllSyslogMessagesLoggingLevel = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsSyslogSettings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
func (data FTDPlatformSettingsSyslogSettings) toBodyPutDelete(ctx context.Context) string {
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
