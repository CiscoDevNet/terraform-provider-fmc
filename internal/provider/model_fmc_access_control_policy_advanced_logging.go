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

type AccessControlPolicyAdvancedLogging struct {
	Id                     types.String `tfsdk:"id"`
	Domain                 types.String `tfsdk:"domain"`
	AccessControlPolicyId  types.String `tfsdk:"access_control_policy_id"`
	Type                   types.String `tfsdk:"type"`
	Enabled                types.Bool   `tfsdk:"enabled"`
	LogFormat              types.String `tfsdk:"log_format"`
	SendToOtherDestination types.Bool   `tfsdk:"send_to_other_destination"`
	SendToSplunk           types.Bool   `tfsdk:"send_to_splunk"`
	SyslogAlertId          types.String `tfsdk:"syslog_alert_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionAccessControlPolicyAdvancedLogging = version.Must(version.NewVersion("10.0"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data AccessControlPolicyAdvancedLogging) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/accesspolicies/%v/advancedloggingsettings", url.QueryEscape(data.AccessControlPolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data AccessControlPolicyAdvancedLogging) toBody(ctx context.Context, state AccessControlPolicyAdvancedLogging) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.LogFormat.IsNull() {
		body, _ = sjson.Set(body, "logOption", data.LogFormat.ValueString())
	}
	if !data.SendToOtherDestination.IsNull() {
		body, _ = sjson.Set(body, "useOverrideDestination", data.SendToOtherDestination.ValueBool())
	}
	if !data.SendToSplunk.IsNull() {
		body, _ = sjson.Set(body, "sendToSplunk", data.SendToSplunk.ValueBool())
	}
	if !data.SyslogAlertId.IsNull() {
		body, _ = sjson.Set(body, "syslogConfig.id", data.SyslogAlertId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *AccessControlPolicyAdvancedLogging) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("logOption"); value.Exists() {
		data.LogFormat = types.StringValue(value.String())
	} else {
		data.LogFormat = types.StringNull()
	}
	if value := res.Get("useOverrideDestination"); value.Exists() {
		data.SendToOtherDestination = types.BoolValue(value.Bool())
	} else {
		data.SendToOtherDestination = types.BoolNull()
	}
	if value := res.Get("sendToSplunk"); value.Exists() {
		data.SendToSplunk = types.BoolValue(value.Bool())
	} else {
		data.SendToSplunk = types.BoolNull()
	}
	if value := res.Get("syslogConfig.id"); value.Exists() {
		data.SyslogAlertId = types.StringValue(value.String())
	} else {
		data.SyslogAlertId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *AccessControlPolicyAdvancedLogging) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("logOption"); value.Exists() && !data.LogFormat.IsNull() {
		data.LogFormat = types.StringValue(value.String())
	} else {
		data.LogFormat = types.StringNull()
	}
	if value := res.Get("useOverrideDestination"); value.Exists() && !data.SendToOtherDestination.IsNull() {
		data.SendToOtherDestination = types.BoolValue(value.Bool())
	} else {
		data.SendToOtherDestination = types.BoolNull()
	}
	if value := res.Get("sendToSplunk"); value.Exists() && !data.SendToSplunk.IsNull() {
		data.SendToSplunk = types.BoolValue(value.Bool())
	} else {
		data.SendToSplunk = types.BoolNull()
	}
	if value := res.Get("syslogConfig.id"); value.Exists() && !data.SyslogAlertId.IsNull() {
		data.SyslogAlertId = types.StringValue(value.String())
	} else {
		data.SyslogAlertId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *AccessControlPolicyAdvancedLogging) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data AccessControlPolicyAdvancedLogging) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	body, _ = sjson.Set(body, "enabled", false)
	body, _ = sjson.Set(body, "logOption", "JSON")
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides
