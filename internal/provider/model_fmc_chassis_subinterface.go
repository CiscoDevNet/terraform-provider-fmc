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

type ChassisSubinterface struct {
	Id                  types.String `tfsdk:"id"`
	Domain              types.String `tfsdk:"domain"`
	ChassisId           types.String `tfsdk:"chassis_id"`
	Type                types.String `tfsdk:"type"`
	Name                types.String `tfsdk:"name"`
	ParentInterfaceName types.String `tfsdk:"parent_interface_name"`
	ParentInterfaceId   types.String `tfsdk:"parent_interface_id"`
	SubInterfaceId      types.Int64  `tfsdk:"sub_interface_id"`
	VlanId              types.Int64  `tfsdk:"vlan_id"`
	PortType            types.String `tfsdk:"port_type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ChassisSubinterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/chassis/fmcmanagedchassis/%v/subinterfaces", url.QueryEscape(data.ChassisId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ChassisSubinterface) toBody(ctx context.Context, state ChassisSubinterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "SubInterface")
	if !data.ParentInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "parentInterface.name", data.ParentInterfaceName.ValueString())
	}
	if !data.ParentInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "parentInterface.id", data.ParentInterfaceId.ValueString())
	}
	if !data.SubInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "subIntfId", data.SubInterfaceId.ValueInt64())
	}
	if !data.VlanId.IsNull() {
		body, _ = sjson.Set(body, "vlanId", data.VlanId.ValueInt64())
	}
	if !data.PortType.IsNull() {
		body, _ = sjson.Set(body, "portType", data.PortType.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ChassisSubinterface) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("parentInterface.name"); value.Exists() {
		data.ParentInterfaceName = types.StringValue(value.String())
	} else {
		data.ParentInterfaceName = types.StringNull()
	}
	if value := res.Get("parentInterface.id"); value.Exists() {
		data.ParentInterfaceId = types.StringValue(value.String())
	} else {
		data.ParentInterfaceId = types.StringNull()
	}
	if value := res.Get("subIntfId"); value.Exists() {
		data.SubInterfaceId = types.Int64Value(value.Int())
	} else {
		data.SubInterfaceId = types.Int64Null()
	}
	if value := res.Get("vlanId"); value.Exists() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("portType"); value.Exists() {
		data.PortType = types.StringValue(value.String())
	} else {
		data.PortType = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ChassisSubinterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("parentInterface.name"); value.Exists() && !data.ParentInterfaceName.IsNull() {
		data.ParentInterfaceName = types.StringValue(value.String())
	} else {
		data.ParentInterfaceName = types.StringNull()
	}
	if value := res.Get("parentInterface.id"); value.Exists() && !data.ParentInterfaceId.IsNull() {
		data.ParentInterfaceId = types.StringValue(value.String())
	} else {
		data.ParentInterfaceId = types.StringNull()
	}
	if value := res.Get("subIntfId"); value.Exists() && !data.SubInterfaceId.IsNull() {
		data.SubInterfaceId = types.Int64Value(value.Int())
	} else {
		data.SubInterfaceId = types.Int64Null()
	}
	if value := res.Get("vlanId"); value.Exists() && !data.VlanId.IsNull() {
		data.VlanId = types.Int64Value(value.Int())
	} else {
		data.VlanId = types.Int64Null()
	}
	if value := res.Get("portType"); value.Exists() && !data.PortType.IsNull() {
		data.PortType = types.StringValue(value.String())
	} else {
		data.PortType = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ChassisSubinterface) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.Name.IsUnknown() {
		if value := res.Get("name"); value.Exists() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
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
