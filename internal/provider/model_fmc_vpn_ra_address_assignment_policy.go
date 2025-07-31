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

type VPNRAAddressAssignmentPolicy struct {
	Id                                   types.String `tfsdk:"id"`
	Domain                               types.String `tfsdk:"domain"`
	VpnRaId                              types.String `tfsdk:"vpn_ra_id"`
	Type                                 types.String `tfsdk:"type"`
	Ipv4UseAuthorizationServer           types.Bool   `tfsdk:"ipv4_use_authorization_server"`
	Ipv4UseDhcp                          types.Bool   `tfsdk:"ipv4_use_dhcp"`
	Ipv4InternalAddressPool              types.Bool   `tfsdk:"ipv4_internal_address_pool"`
	Ipv4InternalAddressPoolReuseInterval types.Int64  `tfsdk:"ipv4_internal_address_pool_reuse_interval"`
	Ipv6UseAuthorizationServer           types.Bool   `tfsdk:"ipv6_use_authorization_server"`
	Ipv6InternalAddressPool              types.Bool   `tfsdk:"ipv6_internal_address_pool"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNRAAddressAssignmentPolicy) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ravpns/%v/addressassignmentsettings", url.QueryEscape(data.VpnRaId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNRAAddressAssignmentPolicy) toBody(ctx context.Context, state VPNRAAddressAssignmentPolicy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "RaVpnAddressAssignmentSetting")
	if !data.Ipv4UseAuthorizationServer.IsNull() {
		body, _ = sjson.Set(body, "useAuthorizationServerForIPv4", data.Ipv4UseAuthorizationServer.ValueBool())
	}
	if !data.Ipv4UseDhcp.IsNull() {
		body, _ = sjson.Set(body, "useDHCP", data.Ipv4UseDhcp.ValueBool())
	}
	if !data.Ipv4InternalAddressPool.IsNull() {
		body, _ = sjson.Set(body, "useInternalAddressPoolForIPv4", data.Ipv4InternalAddressPool.ValueBool())
	}
	if !data.Ipv4InternalAddressPoolReuseInterval.IsNull() {
		body, _ = sjson.Set(body, "ipAddressReuseInterval", data.Ipv4InternalAddressPoolReuseInterval.ValueInt64())
	}
	if !data.Ipv6UseAuthorizationServer.IsNull() {
		body, _ = sjson.Set(body, "useAuthorizationServerForIPv6", data.Ipv6UseAuthorizationServer.ValueBool())
	}
	if !data.Ipv6InternalAddressPool.IsNull() {
		body, _ = sjson.Set(body, "useInternalAddressPoolForIPv6", data.Ipv6InternalAddressPool.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNRAAddressAssignmentPolicy) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("useAuthorizationServerForIPv4"); value.Exists() {
		data.Ipv4UseAuthorizationServer = types.BoolValue(value.Bool())
	} else {
		data.Ipv4UseAuthorizationServer = types.BoolValue(true)
	}
	if value := res.Get("useDHCP"); value.Exists() {
		data.Ipv4UseDhcp = types.BoolValue(value.Bool())
	} else {
		data.Ipv4UseDhcp = types.BoolValue(true)
	}
	if value := res.Get("useInternalAddressPoolForIPv4"); value.Exists() {
		data.Ipv4InternalAddressPool = types.BoolValue(value.Bool())
	} else {
		data.Ipv4InternalAddressPool = types.BoolValue(true)
	}
	if value := res.Get("ipAddressReuseInterval"); value.Exists() {
		data.Ipv4InternalAddressPoolReuseInterval = types.Int64Value(value.Int())
	} else {
		data.Ipv4InternalAddressPoolReuseInterval = types.Int64Value(0)
	}
	if value := res.Get("useAuthorizationServerForIPv6"); value.Exists() {
		data.Ipv6UseAuthorizationServer = types.BoolValue(value.Bool())
	} else {
		data.Ipv6UseAuthorizationServer = types.BoolValue(true)
	}
	if value := res.Get("useInternalAddressPoolForIPv6"); value.Exists() {
		data.Ipv6InternalAddressPool = types.BoolValue(value.Bool())
	} else {
		data.Ipv6InternalAddressPool = types.BoolValue(true)
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNRAAddressAssignmentPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("useAuthorizationServerForIPv4"); value.Exists() && !data.Ipv4UseAuthorizationServer.IsNull() {
		data.Ipv4UseAuthorizationServer = types.BoolValue(value.Bool())
	} else if data.Ipv4UseAuthorizationServer.ValueBool() != true {
		data.Ipv4UseAuthorizationServer = types.BoolNull()
	}
	if value := res.Get("useDHCP"); value.Exists() && !data.Ipv4UseDhcp.IsNull() {
		data.Ipv4UseDhcp = types.BoolValue(value.Bool())
	} else if data.Ipv4UseDhcp.ValueBool() != true {
		data.Ipv4UseDhcp = types.BoolNull()
	}
	if value := res.Get("useInternalAddressPoolForIPv4"); value.Exists() && !data.Ipv4InternalAddressPool.IsNull() {
		data.Ipv4InternalAddressPool = types.BoolValue(value.Bool())
	} else if data.Ipv4InternalAddressPool.ValueBool() != true {
		data.Ipv4InternalAddressPool = types.BoolNull()
	}
	if value := res.Get("ipAddressReuseInterval"); value.Exists() && !data.Ipv4InternalAddressPoolReuseInterval.IsNull() {
		data.Ipv4InternalAddressPoolReuseInterval = types.Int64Value(value.Int())
	} else if data.Ipv4InternalAddressPoolReuseInterval.ValueInt64() != 0 {
		data.Ipv4InternalAddressPoolReuseInterval = types.Int64Null()
	}
	if value := res.Get("useAuthorizationServerForIPv6"); value.Exists() && !data.Ipv6UseAuthorizationServer.IsNull() {
		data.Ipv6UseAuthorizationServer = types.BoolValue(value.Bool())
	} else if data.Ipv6UseAuthorizationServer.ValueBool() != true {
		data.Ipv6UseAuthorizationServer = types.BoolNull()
	}
	if value := res.Get("useInternalAddressPoolForIPv6"); value.Exists() && !data.Ipv6InternalAddressPool.IsNull() {
		data.Ipv6InternalAddressPool = types.BoolValue(value.Bool())
	} else if data.Ipv6InternalAddressPool.ValueBool() != true {
		data.Ipv6InternalAddressPool = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNRAAddressAssignmentPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
func (data VPNRAAddressAssignmentPolicy) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	return body
}

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
