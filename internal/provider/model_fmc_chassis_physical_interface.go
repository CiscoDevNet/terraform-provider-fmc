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

type ChassisPhysicalInterface struct {
	Id              types.String `tfsdk:"id"`
	Domain          types.String `tfsdk:"domain"`
	ChassisId       types.String `tfsdk:"chassis_id"`
	Type            types.String `tfsdk:"type"`
	Name            types.String `tfsdk:"name"`
	PortType        types.String `tfsdk:"port_type"`
	AdminState      types.String `tfsdk:"admin_state"`
	AutoNegotiation types.Bool   `tfsdk:"auto_negotiation"`
	Duplex          types.String `tfsdk:"duplex"`
	Speed           types.String `tfsdk:"speed"`
	FecMode         types.String `tfsdk:"fec_mode"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ChassisPhysicalInterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/chassis/fmcmanagedchassis/%v/physicalinterfaces", url.QueryEscape(data.ChassisId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ChassisPhysicalInterface) toBody(ctx context.Context, state ChassisPhysicalInterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "PhysicalInterface")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.PortType.IsNull() {
		body, _ = sjson.Set(body, "portType", data.PortType.ValueString())
	}
	if !data.AdminState.IsNull() {
		body, _ = sjson.Set(body, "adminState", data.AdminState.ValueString())
	}
	if !data.AutoNegotiation.IsNull() {
		body, _ = sjson.Set(body, "hardware.autoNegState", data.AutoNegotiation.ValueBool())
	}
	if !data.Duplex.IsNull() {
		body, _ = sjson.Set(body, "hardware.duplex", data.Duplex.ValueString())
	}
	if !data.Speed.IsNull() {
		body, _ = sjson.Set(body, "hardware.speed", data.Speed.ValueString())
	}
	if !data.FecMode.IsNull() {
		body, _ = sjson.Set(body, "hardware.fecMode", data.FecMode.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ChassisPhysicalInterface) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("portType"); value.Exists() {
		data.PortType = types.StringValue(value.String())
	} else {
		data.PortType = types.StringNull()
	}
	if value := res.Get("adminState"); value.Exists() {
		data.AdminState = types.StringValue(value.String())
	} else {
		data.AdminState = types.StringValue("ENABLED")
	}
	if value := res.Get("hardware.autoNegState"); value.Exists() {
		data.AutoNegotiation = types.BoolValue(value.Bool())
	} else {
		data.AutoNegotiation = types.BoolNull()
	}
	if value := res.Get("hardware.duplex"); value.Exists() {
		data.Duplex = types.StringValue(value.String())
	} else {
		data.Duplex = types.StringNull()
	}
	if value := res.Get("hardware.speed"); value.Exists() {
		data.Speed = types.StringValue(value.String())
	} else {
		data.Speed = types.StringNull()
	}
	if value := res.Get("hardware.fecMode"); value.Exists() {
		data.FecMode = types.StringValue(value.String())
	} else {
		data.FecMode = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ChassisPhysicalInterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("portType"); value.Exists() && !data.PortType.IsNull() {
		data.PortType = types.StringValue(value.String())
	} else {
		data.PortType = types.StringNull()
	}
	if value := res.Get("adminState"); value.Exists() && !data.AdminState.IsNull() {
		data.AdminState = types.StringValue(value.String())
	} else if data.AdminState.ValueString() != "ENABLED" {
		data.AdminState = types.StringNull()
	}
	if value := res.Get("hardware.autoNegState"); value.Exists() && !data.AutoNegotiation.IsNull() {
		data.AutoNegotiation = types.BoolValue(value.Bool())
	} else {
		data.AutoNegotiation = types.BoolNull()
	}
	if value := res.Get("hardware.duplex"); value.Exists() && !data.Duplex.IsNull() {
		data.Duplex = types.StringValue(value.String())
	} else {
		data.Duplex = types.StringNull()
	}
	if value := res.Get("hardware.speed"); value.Exists() && !data.Speed.IsNull() {
		data.Speed = types.StringValue(value.String())
	} else {
		data.Speed = types.StringNull()
	}
	if value := res.Get("hardware.fecMode"); value.Exists() && !data.FecMode.IsNull() {
		data.FecMode = types.StringValue(value.String())
	} else {
		data.FecMode = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ChassisPhysicalInterface) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// toBodyPutDelete generates minimal required body to reset the resource to its default state.
func (data ChassisPhysicalInterface) toBodyPutDelete(ctx context.Context, state ChassisPhysicalInterface) string {
	body := ""
	body, _ = sjson.Set(body, "type", "PhysicalInterface")
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	return body
}
