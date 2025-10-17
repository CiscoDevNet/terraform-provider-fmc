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

type Chassis struct {
	Id              types.String `tfsdk:"id"`
	Domain          types.String `tfsdk:"domain"`
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	HostName        types.String `tfsdk:"host_name"`
	RegistrationKey types.String `tfsdk:"registration_key"`
	DeviceGroupId   types.String `tfsdk:"device_group_id"`
	NatId           types.String `tfsdk:"nat_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data Chassis) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/chassis/fmcmanagedchassis"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data Chassis) toBody(ctx context.Context, state Chassis) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "chassisName", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "Device")
	if !data.HostName.IsNull() {
		body, _ = sjson.Set(body, "chassisHostName", data.HostName.ValueString())
	}
	if !data.RegistrationKey.IsNull() {
		body, _ = sjson.Set(body, "regKey", data.RegistrationKey.ValueString())
	}
	if !data.DeviceGroupId.IsNull() {
		body, _ = sjson.Set(body, "deviceGroup.id", data.DeviceGroupId.ValueString())
	}
	if !data.NatId.IsNull() {
		body, _ = sjson.Set(body, "natID", data.NatId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *Chassis) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("chassisName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("chassisHostName"); value.Exists() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("deviceGroup.id"); value.Exists() {
		data.DeviceGroupId = types.StringValue(value.String())
	} else {
		data.DeviceGroupId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *Chassis) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("chassisName"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("chassisHostName"); value.Exists() && !data.HostName.IsNull() {
		data.HostName = types.StringValue(value.String())
	} else {
		data.HostName = types.StringNull()
	}
	if value := res.Get("deviceGroup.id"); value.Exists() && !data.DeviceGroupId.IsNull() {
		data.DeviceGroupId = types.StringValue(value.String())
	} else {
		data.DeviceGroupId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *Chassis) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns
