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

type IPv6AddressPool struct {
	Id                types.String `tfsdk:"id"`
	Domain            types.String `tfsdk:"domain"`
	Name              types.String `tfsdk:"name"`
	Type              types.String `tfsdk:"type"`
	Description       types.String `tfsdk:"description"`
	StartAddress      types.String `tfsdk:"start_address"`
	NumberOfAddresses types.Int64  `tfsdk:"number_of_addresses"`
	Overridable       types.Bool   `tfsdk:"overridable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data IPv6AddressPool) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/ipv6addresspools"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data IPv6AddressPool) toBody(ctx context.Context, state IPv6AddressPool) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "IPv6AddressPool")
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.StartAddress.IsNull() {
		body, _ = sjson.Set(body, "ipv6StartAddress", data.StartAddress.ValueString())
	}
	if !data.NumberOfAddresses.IsNull() {
		body, _ = sjson.Set(body, "numberOfAddresses", data.NumberOfAddresses.ValueInt64())
	}
	if !data.Overridable.IsNull() {
		body, _ = sjson.Set(body, "overridable", data.Overridable.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *IPv6AddressPool) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("ipv6StartAddress"); value.Exists() {
		data.StartAddress = types.StringValue(value.String())
	} else {
		data.StartAddress = types.StringNull()
	}
	if value := res.Get("numberOfAddresses"); value.Exists() {
		data.NumberOfAddresses = types.Int64Value(value.Int())
	} else {
		data.NumberOfAddresses = types.Int64Null()
	}
	if value := res.Get("overridable"); value.Exists() {
		data.Overridable = types.BoolValue(value.Bool())
	} else {
		data.Overridable = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *IPv6AddressPool) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		if !data.Description.IsNull() && data.Description.ValueString() == "" {
			data.Description = types.StringValue("")
		} else {
			data.Description = types.StringNull()
		}
	}
	if value := res.Get("ipv6StartAddress"); value.Exists() && !data.StartAddress.IsNull() {
		data.StartAddress = types.StringValue(value.String())
	} else {
		data.StartAddress = types.StringNull()
	}
	if value := res.Get("numberOfAddresses"); value.Exists() && !data.NumberOfAddresses.IsNull() {
		data.NumberOfAddresses = types.Int64Value(value.Int())
	} else {
		data.NumberOfAddresses = types.Int64Null()
	}
	if value := res.Get("overridable"); value.Exists() && !data.Overridable.IsNull() {
		data.Overridable = types.BoolValue(value.Bool())
	} else {
		data.Overridable = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *IPv6AddressPool) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
