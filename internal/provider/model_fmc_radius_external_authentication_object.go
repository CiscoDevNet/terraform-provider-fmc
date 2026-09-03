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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type RadiusExternalAuthenticationObject struct {
	Id                          types.String `tfsdk:"id"`
	Domain                      types.String `tfsdk:"domain"`
	Name                        types.String `tfsdk:"name"`
	Type                        types.String `tfsdk:"type"`
	Description                 types.String `tfsdk:"description"`
	ServerAddress               types.String `tfsdk:"server_address"`
	ServerPort                  types.String `tfsdk:"server_port"`
	Key                         types.String `tfsdk:"key"`
	BackupServerAddress         types.String `tfsdk:"backup_server_address"`
	BackupServerPort            types.String `tfsdk:"backup_server_port"`
	BackupKey                   types.String `tfsdk:"backup_key"`
	Timeout                     types.Int64  `tfsdk:"timeout"`
	Retries                     types.Int64  `tfsdk:"retries"`
	MessageAuthenticatorEnabled types.Bool   `tfsdk:"message_authenticator_enabled"`
	CliAccessUserList           types.String `tfsdk:"cli_access_user_list"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data RadiusExternalAuthenticationObject) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/users/externalauths/authconfigobjects/radiusconfigobjects"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data RadiusExternalAuthenticationObject) toBody(ctx context.Context, state RadiusExternalAuthenticationObject) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "RADIUSConfigObject")
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.ServerAddress.IsNull() {
		body, _ = sjson.Set(body, "serverAddress", data.ServerAddress.ValueString())
	}
	if !data.ServerPort.IsNull() {
		body, _ = sjson.Set(body, "serverPort", data.ServerPort.ValueString())
	}
	if !data.Key.IsNull() {
		body, _ = sjson.Set(body, "secretKey", data.Key.ValueString())
	}
	if !data.BackupServerAddress.IsNull() {
		body, _ = sjson.Set(body, "backupServerAddress", data.BackupServerAddress.ValueString())
	}
	if !data.BackupServerPort.IsNull() {
		body, _ = sjson.Set(body, "backupServerPort", data.BackupServerPort.ValueString())
	}
	if !data.BackupKey.IsNull() {
		body, _ = sjson.Set(body, "backupServerSecretKey", data.BackupKey.ValueString())
	}
	if !data.Timeout.IsNull() {
		body, _ = sjson.Set(body, "timeout", data.Timeout.ValueInt64())
	}
	if !data.Retries.IsNull() {
		body, _ = sjson.Set(body, "retries", data.Retries.ValueInt64())
	}
	if !data.MessageAuthenticatorEnabled.IsNull() {
		body, _ = sjson.Set(body, "isMAEnabled", data.MessageAuthenticatorEnabled.ValueBool())
	}
	if !data.CliAccessUserList.IsNull() {
		body, _ = sjson.Set(body, "cliAccessUserList", data.CliAccessUserList.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *RadiusExternalAuthenticationObject) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("serverAddress"); value.Exists() {
		data.ServerAddress = types.StringValue(value.String())
	} else {
		data.ServerAddress = types.StringNull()
	}
	if value := res.Get("serverPort"); value.Exists() {
		data.ServerPort = types.StringValue(value.String())
	} else {
		data.ServerPort = types.StringValue("1812")
	}
	if value := res.Get("backupServerAddress"); value.Exists() {
		data.BackupServerAddress = types.StringValue(value.String())
	} else {
		data.BackupServerAddress = types.StringNull()
	}
	if value := res.Get("backupServerPort"); value.Exists() {
		data.BackupServerPort = types.StringValue(value.String())
	} else {
		data.BackupServerPort = types.StringNull()
	}
	if value := res.Get("timeout"); value.Exists() {
		data.Timeout = types.Int64Value(value.Int())
	} else {
		data.Timeout = types.Int64Value(30)
	}
	if value := res.Get("retries"); value.Exists() {
		data.Retries = types.Int64Value(value.Int())
	} else {
		data.Retries = types.Int64Value(3)
	}
	if value := res.Get("isMAEnabled"); value.Exists() {
		data.MessageAuthenticatorEnabled = types.BoolValue(value.Bool())
	} else {
		data.MessageAuthenticatorEnabled = types.BoolValue(true)
	}
	if value := res.Get("cliAccessUserList"); value.Exists() {
		data.CliAccessUserList = types.StringValue(value.String())
	} else {
		data.CliAccessUserList = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *RadiusExternalAuthenticationObject) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("serverAddress"); value.Exists() && !data.ServerAddress.IsNull() {
		data.ServerAddress = types.StringValue(value.String())
	} else {
		data.ServerAddress = types.StringNull()
	}
	if value := res.Get("serverPort"); value.Exists() && !data.ServerPort.IsNull() {
		data.ServerPort = types.StringValue(value.String())
	} else if data.ServerPort.ValueString() != "1812" {
		data.ServerPort = types.StringNull()
	}
	if value := res.Get("backupServerAddress"); value.Exists() && !data.BackupServerAddress.IsNull() {
		data.BackupServerAddress = types.StringValue(value.String())
	} else {
		data.BackupServerAddress = types.StringNull()
	}
	if value := res.Get("backupServerPort"); value.Exists() && !data.BackupServerPort.IsNull() {
		data.BackupServerPort = types.StringValue(value.String())
	} else {
		data.BackupServerPort = types.StringNull()
	}
	if value := res.Get("timeout"); value.Exists() && !data.Timeout.IsNull() {
		data.Timeout = types.Int64Value(value.Int())
	} else if data.Timeout.ValueInt64() != 30 {
		data.Timeout = types.Int64Null()
	}
	if value := res.Get("retries"); value.Exists() && !data.Retries.IsNull() {
		data.Retries = types.Int64Value(value.Int())
	} else if data.Retries.ValueInt64() != 3 {
		data.Retries = types.Int64Null()
	}
	if value := res.Get("isMAEnabled"); value.Exists() && !data.MessageAuthenticatorEnabled.IsNull() {
		data.MessageAuthenticatorEnabled = types.BoolValue(value.Bool())
	} else if data.MessageAuthenticatorEnabled.ValueBool() != true {
		data.MessageAuthenticatorEnabled = types.BoolNull()
	}
	if value := res.Get("cliAccessUserList"); value.Exists() && !data.CliAccessUserList.IsNull() {
		data.CliAccessUserList = types.StringValue(value.String())
	} else {
		data.CliAccessUserList = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *RadiusExternalAuthenticationObject) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
