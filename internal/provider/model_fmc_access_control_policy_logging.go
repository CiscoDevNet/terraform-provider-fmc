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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type AccessControlPolicyLogging struct {
	Id                                 types.String `tfsdk:"id"`
	Domain                             types.String `tfsdk:"domain"`
	AccessControlPolicyId              types.String `tfsdk:"access_control_policy_id"`
	Type                               types.String `tfsdk:"type"`
	SyslogAlertId                      types.String `tfsdk:"syslog_alert_id"`
	SyslogUsePlatformSettings          types.Bool   `tfsdk:"syslog_use_platform_settings"`
	SyslogPlatformSettingsSeverity     types.String `tfsdk:"syslog_platform_settings_severity"`
	IpsEventsSendSyslog                types.Bool   `tfsdk:"ips_events_send_syslog"`
	IpsEventsSyslogSeverity            types.String `tfsdk:"ips_events_syslog_severity"`
	IpsEventsSyslogAlertId             types.String `tfsdk:"ips_events_syslog_alert_id"`
	FileAndMalwareEventsSendSyslog     types.Bool   `tfsdk:"file_and_malware_events_send_syslog"`
	FileAndMalwareEventsSyslogSeverity types.String `tfsdk:"file_and_malware_events_syslog_severity"`
	FileAndMalwareEventsSyslogAlertId  types.String `tfsdk:"file_and_malware_events_syslog_alert_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data AccessControlPolicyLogging) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/accesspolicies/%v/loggingsettings", url.QueryEscape(data.AccessControlPolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data AccessControlPolicyLogging) toBody(ctx context.Context, state AccessControlPolicyLogging) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.SyslogAlertId.IsNull() {
		body, _ = sjson.Set(body, "syslogConfig.id", data.SyslogAlertId.ValueString())
	}
	if !data.SyslogUsePlatformSettings.IsNull() {
		body, _ = sjson.Set(body, "syslogConfigFromPlatformSetting", data.SyslogUsePlatformSettings.ValueBool())
	}
	if !data.SyslogPlatformSettingsSeverity.IsNull() {
		body, _ = sjson.Set(body, "severityForPlatformSettingSyslogConfig", data.SyslogPlatformSettingsSeverity.ValueString())
	}
	if !data.IpsEventsSendSyslog.IsNull() {
		body, _ = sjson.Set(body, "enableipsSyslog", data.IpsEventsSendSyslog.ValueBool())
	}
	if !data.IpsEventsSyslogSeverity.IsNull() {
		body, _ = sjson.Set(body, "ipsSyslogSeverity", data.IpsEventsSyslogSeverity.ValueString())
	}
	if !data.IpsEventsSyslogAlertId.IsNull() {
		body, _ = sjson.Set(body, "ipsSyslogConfig.id", data.IpsEventsSyslogAlertId.ValueString())
	}
	if !data.FileAndMalwareEventsSendSyslog.IsNull() {
		body, _ = sjson.Set(body, "enableFileAndMalwareSyslog", data.FileAndMalwareEventsSendSyslog.ValueBool())
	}
	if !data.FileAndMalwareEventsSyslogSeverity.IsNull() {
		body, _ = sjson.Set(body, "fileAndMalwareSyslogSeverity", data.FileAndMalwareEventsSyslogSeverity.ValueString())
	}
	if !data.FileAndMalwareEventsSyslogAlertId.IsNull() {
		body, _ = sjson.Set(body, "fileAndMalwareSyslogConfig.id", data.FileAndMalwareEventsSyslogAlertId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *AccessControlPolicyLogging) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("syslogConfig.id"); value.Exists() {
		data.SyslogAlertId = types.StringValue(value.String())
	} else {
		data.SyslogAlertId = types.StringNull()
	}
	if value := res.Get("syslogConfigFromPlatformSetting"); value.Exists() {
		data.SyslogUsePlatformSettings = types.BoolValue(value.Bool())
	} else {
		data.SyslogUsePlatformSettings = types.BoolNull()
	}
	if value := res.Get("severityForPlatformSettingSyslogConfig"); value.Exists() {
		data.SyslogPlatformSettingsSeverity = types.StringValue(value.String())
	} else {
		data.SyslogPlatformSettingsSeverity = types.StringNull()
	}
	if value := res.Get("enableipsSyslog"); value.Exists() {
		data.IpsEventsSendSyslog = types.BoolValue(value.Bool())
	} else {
		data.IpsEventsSendSyslog = types.BoolNull()
	}
	if value := res.Get("ipsSyslogSeverity"); value.Exists() {
		data.IpsEventsSyslogSeverity = types.StringValue(value.String())
	} else {
		data.IpsEventsSyslogSeverity = types.StringNull()
	}
	if value := res.Get("ipsSyslogConfig.id"); value.Exists() {
		data.IpsEventsSyslogAlertId = types.StringValue(value.String())
	} else {
		data.IpsEventsSyslogAlertId = types.StringNull()
	}
	if value := res.Get("enableFileAndMalwareSyslog"); value.Exists() {
		data.FileAndMalwareEventsSendSyslog = types.BoolValue(value.Bool())
	} else {
		data.FileAndMalwareEventsSendSyslog = types.BoolNull()
	}
	if value := res.Get("fileAndMalwareSyslogSeverity"); value.Exists() {
		data.FileAndMalwareEventsSyslogSeverity = types.StringValue(value.String())
	} else {
		data.FileAndMalwareEventsSyslogSeverity = types.StringNull()
	}
	if value := res.Get("fileAndMalwareSyslogConfig.id"); value.Exists() {
		data.FileAndMalwareEventsSyslogAlertId = types.StringValue(value.String())
	} else {
		data.FileAndMalwareEventsSyslogAlertId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *AccessControlPolicyLogging) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("syslogConfig.id"); value.Exists() && !data.SyslogAlertId.IsNull() {
		data.SyslogAlertId = types.StringValue(value.String())
	} else {
		data.SyslogAlertId = types.StringNull()
	}
	if value := res.Get("syslogConfigFromPlatformSetting"); value.Exists() && !data.SyslogUsePlatformSettings.IsNull() {
		data.SyslogUsePlatformSettings = types.BoolValue(value.Bool())
	} else {
		data.SyslogUsePlatformSettings = types.BoolNull()
	}
	if value := res.Get("severityForPlatformSettingSyslogConfig"); value.Exists() && !data.SyslogPlatformSettingsSeverity.IsNull() {
		data.SyslogPlatformSettingsSeverity = types.StringValue(value.String())
	} else {
		data.SyslogPlatformSettingsSeverity = types.StringNull()
	}
	if value := res.Get("enableipsSyslog"); value.Exists() && !data.IpsEventsSendSyslog.IsNull() {
		data.IpsEventsSendSyslog = types.BoolValue(value.Bool())
	} else {
		data.IpsEventsSendSyslog = types.BoolNull()
	}
	if value := res.Get("ipsSyslogSeverity"); value.Exists() && !data.IpsEventsSyslogSeverity.IsNull() {
		data.IpsEventsSyslogSeverity = types.StringValue(value.String())
	} else {
		data.IpsEventsSyslogSeverity = types.StringNull()
	}
	if value := res.Get("ipsSyslogConfig.id"); value.Exists() && !data.IpsEventsSyslogAlertId.IsNull() {
		data.IpsEventsSyslogAlertId = types.StringValue(value.String())
	} else {
		data.IpsEventsSyslogAlertId = types.StringNull()
	}
	if value := res.Get("enableFileAndMalwareSyslog"); value.Exists() && !data.FileAndMalwareEventsSendSyslog.IsNull() {
		data.FileAndMalwareEventsSendSyslog = types.BoolValue(value.Bool())
	} else {
		data.FileAndMalwareEventsSendSyslog = types.BoolNull()
	}
	if value := res.Get("fileAndMalwareSyslogSeverity"); value.Exists() && !data.FileAndMalwareEventsSyslogSeverity.IsNull() {
		data.FileAndMalwareEventsSyslogSeverity = types.StringValue(value.String())
	} else {
		data.FileAndMalwareEventsSyslogSeverity = types.StringNull()
	}
	if value := res.Get("fileAndMalwareSyslogConfig.id"); value.Exists() && !data.FileAndMalwareEventsSyslogAlertId.IsNull() {
		data.FileAndMalwareEventsSyslogAlertId = types.StringValue(value.String())
	} else {
		data.FileAndMalwareEventsSyslogAlertId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *AccessControlPolicyLogging) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data AccessControlPolicyLogging) toBodyPutDelete(ctx context.Context) string {
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

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides
