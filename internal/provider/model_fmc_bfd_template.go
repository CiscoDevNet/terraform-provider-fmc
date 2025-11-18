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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type BFDTemplate struct {
	Id                               types.String `tfsdk:"id"`
	Domain                           types.String `tfsdk:"domain"`
	Name                             types.String `tfsdk:"name"`
	Type                             types.String `tfsdk:"type"`
	HopType                          types.String `tfsdk:"hop_type"`
	Echo                             types.String `tfsdk:"echo"`
	IntervalType                     types.String `tfsdk:"interval_type"`
	Multiplier                       types.Int64  `tfsdk:"multiplier"`
	MinimumTransmit                  types.Int64  `tfsdk:"minimum_transmit"`
	MinimumReceive                   types.Int64  `tfsdk:"minimum_receive"`
	AuthenticationType               types.String `tfsdk:"authentication_type"`
	AuthenticationPassword           types.String `tfsdk:"authentication_password"`
	AuthenticationPasswordEncryption types.String `tfsdk:"authentication_password_encryption"`
	AuthenticationKeyId              types.Int64  `tfsdk:"authentication_key_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBFDTemplate = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data BFDTemplate) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/bfdtemplates"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data BFDTemplate) toBody(ctx context.Context, state BFDTemplate) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "BFDTemplate")
	if !data.HopType.IsNull() {
		body, _ = sjson.Set(body, "hopType", data.HopType.ValueString())
	}
	if !data.Echo.IsNull() {
		body, _ = sjson.Set(body, "echo", data.Echo.ValueString())
	}
	if !data.IntervalType.IsNull() {
		body, _ = sjson.Set(body, "txRxInterval", data.IntervalType.ValueString())
	}
	if !data.Multiplier.IsNull() {
		body, _ = sjson.Set(body, "txRxMultiplier", data.Multiplier.ValueInt64())
	}
	if !data.MinimumTransmit.IsNull() {
		body, _ = sjson.Set(body, "minTransmit", data.MinimumTransmit.ValueInt64())
	}
	if !data.MinimumReceive.IsNull() {
		body, _ = sjson.Set(body, "minReceive", data.MinimumReceive.ValueInt64())
	}
	if !data.AuthenticationType.IsNull() {
		body, _ = sjson.Set(body, "authentication.authType", data.AuthenticationType.ValueString())
	}
	if !data.AuthenticationPassword.IsNull() {
		body, _ = sjson.Set(body, "authentication.authKey", data.AuthenticationPassword.ValueString())
	}
	if !data.AuthenticationPasswordEncryption.IsNull() {
		body, _ = sjson.Set(body, "authentication.pwdEncryption", data.AuthenticationPasswordEncryption.ValueString())
	}
	if !data.AuthenticationKeyId.IsNull() {
		body, _ = sjson.Set(body, "authentication.authKeyId", data.AuthenticationKeyId.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *BFDTemplate) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("hopType"); value.Exists() {
		data.HopType = types.StringValue(value.String())
	} else {
		data.HopType = types.StringNull()
	}
	if value := res.Get("echo"); value.Exists() {
		data.Echo = types.StringValue(value.String())
	} else {
		data.Echo = types.StringNull()
	}
	if value := res.Get("txRxInterval"); value.Exists() {
		data.IntervalType = types.StringValue(value.String())
	} else {
		data.IntervalType = types.StringNull()
	}
	if value := res.Get("txRxMultiplier"); value.Exists() {
		data.Multiplier = types.Int64Value(value.Int())
	} else {
		data.Multiplier = types.Int64Null()
	}
	if value := res.Get("minTransmit"); value.Exists() {
		data.MinimumTransmit = types.Int64Value(value.Int())
	} else {
		data.MinimumTransmit = types.Int64Null()
	}
	if value := res.Get("minReceive"); value.Exists() {
		data.MinimumReceive = types.Int64Value(value.Int())
	} else {
		data.MinimumReceive = types.Int64Null()
	}
	if value := res.Get("authentication.authType"); value.Exists() {
		data.AuthenticationType = types.StringValue(value.String())
	} else {
		data.AuthenticationType = types.StringNull()
	}
	if value := res.Get("authentication.pwdEncryption"); value.Exists() {
		data.AuthenticationPasswordEncryption = types.StringValue(value.String())
	} else {
		data.AuthenticationPasswordEncryption = types.StringNull()
	}
	if value := res.Get("authentication.authKeyId"); value.Exists() {
		data.AuthenticationKeyId = types.Int64Value(value.Int())
	} else {
		data.AuthenticationKeyId = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *BFDTemplate) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("hopType"); value.Exists() && !data.HopType.IsNull() {
		data.HopType = types.StringValue(value.String())
	} else {
		data.HopType = types.StringNull()
	}
	if value := res.Get("echo"); value.Exists() && !data.Echo.IsNull() {
		data.Echo = types.StringValue(value.String())
	} else {
		data.Echo = types.StringNull()
	}
	if value := res.Get("txRxInterval"); value.Exists() && !data.IntervalType.IsNull() {
		data.IntervalType = types.StringValue(value.String())
	} else {
		data.IntervalType = types.StringNull()
	}
	if value := res.Get("txRxMultiplier"); value.Exists() && !data.Multiplier.IsNull() {
		data.Multiplier = types.Int64Value(value.Int())
	} else {
		data.Multiplier = types.Int64Null()
	}
	if value := res.Get("minTransmit"); value.Exists() && !data.MinimumTransmit.IsNull() {
		data.MinimumTransmit = types.Int64Value(value.Int())
	} else {
		data.MinimumTransmit = types.Int64Null()
	}
	if value := res.Get("minReceive"); value.Exists() && !data.MinimumReceive.IsNull() {
		data.MinimumReceive = types.Int64Value(value.Int())
	} else {
		data.MinimumReceive = types.Int64Null()
	}
	if value := res.Get("authentication.authType"); value.Exists() && !data.AuthenticationType.IsNull() {
		data.AuthenticationType = types.StringValue(value.String())
	} else {
		data.AuthenticationType = types.StringNull()
	}
	if value := res.Get("authentication.pwdEncryption"); value.Exists() && !data.AuthenticationPasswordEncryption.IsNull() {
		data.AuthenticationPasswordEncryption = types.StringValue(value.String())
	} else {
		data.AuthenticationPasswordEncryption = types.StringNull()
	}
	if value := res.Get("authentication.authKeyId"); value.Exists() && !data.AuthenticationKeyId.IsNull() {
		data.AuthenticationKeyId = types.Int64Value(value.Int())
	} else {
		data.AuthenticationKeyId = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *BFDTemplate) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
