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

type DeviceHAPairFailoverInterfaceMACAddress struct {
	Id                types.String `tfsdk:"id"`
	Domain            types.String `tfsdk:"domain"`
	HaPairId          types.String `tfsdk:"ha_pair_id"`
	Type              types.String `tfsdk:"type"`
	InterfaceName     types.String `tfsdk:"interface_name"`
	InterfaceId       types.String `tfsdk:"interface_id"`
	InterfaceType     types.String `tfsdk:"interface_type"`
	ActiveMacAddress  types.String `tfsdk:"active_mac_address"`
	StandbyMacAddress types.String `tfsdk:"standby_mac_address"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceHAPairFailoverInterfaceMACAddress) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devicehapairs/ftddevicehapairs/%v/failoverinterfacemacaddressconfigs", url.QueryEscape(data.HaPairId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceHAPairFailoverInterfaceMACAddress) toBody(ctx context.Context, state DeviceHAPairFailoverInterfaceMACAddress) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.InterfaceName.IsNull() {
		body, _ = sjson.Set(body, "physicalInterface.name", data.InterfaceName.ValueString())
	}
	if !data.InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "physicalInterface.id", data.InterfaceId.ValueString())
	}
	if !data.InterfaceType.IsNull() {
		body, _ = sjson.Set(body, "physicalInterface.type", data.InterfaceType.ValueString())
	}
	if !data.ActiveMacAddress.IsNull() {
		body, _ = sjson.Set(body, "failoverActiveMac", data.ActiveMacAddress.ValueString())
	}
	if !data.StandbyMacAddress.IsNull() {
		body, _ = sjson.Set(body, "failoverStandbyMac", data.StandbyMacAddress.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceHAPairFailoverInterfaceMACAddress) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("physicalInterface.name"); value.Exists() {
		data.InterfaceName = types.StringValue(value.String())
	} else {
		data.InterfaceName = types.StringNull()
	}
	if value := res.Get("physicalInterface.id"); value.Exists() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if value := res.Get("physicalInterface.type"); value.Exists() {
		data.InterfaceType = types.StringValue(value.String())
	} else {
		data.InterfaceType = types.StringNull()
	}
	if value := res.Get("failoverActiveMac"); value.Exists() {
		data.ActiveMacAddress = types.StringValue(value.String())
	} else {
		data.ActiveMacAddress = types.StringNull()
	}
	if value := res.Get("failoverStandbyMac"); value.Exists() {
		data.StandbyMacAddress = types.StringValue(value.String())
	} else {
		data.StandbyMacAddress = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceHAPairFailoverInterfaceMACAddress) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("physicalInterface.name"); value.Exists() && !data.InterfaceName.IsNull() {
		data.InterfaceName = types.StringValue(value.String())
	} else {
		data.InterfaceName = types.StringNull()
	}
	if value := res.Get("physicalInterface.id"); value.Exists() && !data.InterfaceId.IsNull() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if value := res.Get("physicalInterface.type"); value.Exists() && !data.InterfaceType.IsNull() {
		data.InterfaceType = types.StringValue(value.String())
	} else {
		data.InterfaceType = types.StringNull()
	}
	if value := res.Get("failoverActiveMac"); value.Exists() && !data.ActiveMacAddress.IsNull() {
		data.ActiveMacAddress = types.StringValue(value.String())
	} else {
		data.ActiveMacAddress = types.StringNull()
	}
	if value := res.Get("failoverStandbyMac"); value.Exists() && !data.StandbyMacAddress.IsNull() {
		data.StandbyMacAddress = types.StringValue(value.String())
	} else {
		data.StandbyMacAddress = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceHAPairFailoverInterfaceMACAddress) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
