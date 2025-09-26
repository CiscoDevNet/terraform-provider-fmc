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

type FTDPlatformSettingsSyslogSettingsSyslogID struct {
	Id                                  types.String `tfsdk:"id"`
	Domain                              types.String `tfsdk:"domain"`
	FtdPlatformSettingsId               types.String `tfsdk:"ftd_platform_settings_id"`
	FtdPlatformSettingsSyslogSettingsId types.String `tfsdk:"ftd_platform_settings_syslog_settings_id"`
	Type                                types.String `tfsdk:"type"`
	SyslogId                            types.Int64  `tfsdk:"syslog_id"`
	LoggingLevel                        types.String `tfsdk:"logging_level"`
	Enabled                             types.Bool   `tfsdk:"enabled"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionFTDPlatformSettingsSyslogSettingsSyslogID = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsSyslogSettingsSyslogID) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/syslog/syslogsettings/%v/syslogids", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()), url.QueryEscape(data.FtdPlatformSettingsSyslogSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsSyslogSettingsSyslogID) toBody(ctx context.Context, state FTDPlatformSettingsSyslogSettingsSyslogID) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.SyslogId.IsNull() {
		body, _ = sjson.Set(body, "syslogId", data.SyslogId.ValueInt64())
	}
	if !data.LoggingLevel.IsNull() {
		body, _ = sjson.Set(body, "logLevel", data.LoggingLevel.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsSyslogSettingsSyslogID) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("syslogId"); value.Exists() {
		data.SyslogId = types.Int64Value(value.Int())
	} else {
		data.SyslogId = types.Int64Null()
	}
	if value := res.Get("logLevel"); value.Exists() {
		data.LoggingLevel = types.StringValue(value.String())
	} else {
		data.LoggingLevel = types.StringValue("DEFAULT")
	}
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolValue(true)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FTDPlatformSettingsSyslogSettingsSyslogID) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("syslogId"); value.Exists() && !data.SyslogId.IsNull() {
		data.SyslogId = types.Int64Value(value.Int())
	} else {
		data.SyslogId = types.Int64Null()
	}
	if value := res.Get("logLevel"); value.Exists() && !data.LoggingLevel.IsNull() {
		data.LoggingLevel = types.StringValue(value.String())
	} else if data.LoggingLevel.ValueString() != "DEFAULT" {
		data.LoggingLevel = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else if data.Enabled.ValueBool() != true {
		data.Enabled = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsSyslogSettingsSyslogID) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

func (data FTDPlatformSettingsSyslogSettingsSyslogID) adjustBody(ctx context.Context, req string) string {
	// FMCBUG CSCwr26361 FMC API: FTD Platform Settings Syslog Settings Syslog ID 'enabled' field value gets inverted
	if !data.Enabled.IsNull() {
		req, _ = sjson.Set(req, "enabled", !data.Enabled.ValueBool())
	}
	return req
}
