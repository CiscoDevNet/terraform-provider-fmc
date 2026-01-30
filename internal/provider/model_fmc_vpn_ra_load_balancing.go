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

type VPNRALoadBalancing struct {
	Id                               types.String `tfsdk:"id"`
	Domain                           types.String `tfsdk:"domain"`
	VpnRaId                          types.String `tfsdk:"vpn_ra_id"`
	Type                             types.String `tfsdk:"type"`
	Enabled                          types.Bool   `tfsdk:"enabled"`
	Ipv4GroupAddress                 types.String `tfsdk:"ipv4_group_address"`
	Ipv6GroupAddress                 types.String `tfsdk:"ipv6_group_address"`
	InterfaceId                      types.String `tfsdk:"interface_id"`
	Port                             types.Int64  `tfsdk:"port"`
	Ipsec                            types.Bool   `tfsdk:"ipsec"`
	IpsecEncryptionKey               types.String `tfsdk:"ipsec_encryption_key"`
	SendFqdnToPeerDevicesInsteadOfIp types.Bool   `tfsdk:"send_fqdn_to_peer_devices_instead_of_ip"`
	Ikev2RedirectPhase               types.String `tfsdk:"ikev2_redirect_phase"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNRALoadBalancing) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ravpns/%v/loadbalancesettings", url.QueryEscape(data.VpnRaId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNRALoadBalancing) toBody(ctx context.Context, state VPNRALoadBalancing) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "LoadBalanacing")
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enableVpnLoadBalancing", data.Enabled.ValueBool())
	}
	if !data.Ipv4GroupAddress.IsNull() {
		body, _ = sjson.Set(body, "groupSettings.groupIPv4Address", data.Ipv4GroupAddress.ValueString())
	}
	if !data.Ipv6GroupAddress.IsNull() {
		body, _ = sjson.Set(body, "groupSettings.groupIPv6Address", data.Ipv6GroupAddress.ValueString())
	}
	if !data.InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "groupSettings.communicationInterface.id", data.InterfaceId.ValueString())
	}
	if !data.Port.IsNull() {
		body, _ = sjson.Set(body, "groupSettings.communicationUdpPort", data.Port.ValueInt64())
	}
	if !data.Ipsec.IsNull() {
		body, _ = sjson.Set(body, "groupSettings.ipsecEncryption.enable", data.Ipsec.ValueBool())
	}
	if !data.IpsecEncryptionKey.IsNull() {
		body, _ = sjson.Set(body, "groupSettings.ipsecEncryption.encryptionKey", data.IpsecEncryptionKey.ValueString())
	}
	if !data.SendFqdnToPeerDevicesInsteadOfIp.IsNull() {
		body, _ = sjson.Set(body, "redirectSettings.redirectUsingFqdn", data.SendFqdnToPeerDevicesInsteadOfIp.ValueBool())
	}
	if !data.Ikev2RedirectPhase.IsNull() {
		body, _ = sjson.Set(body, "redirectSettings.ikev2RedirectPhase", data.Ikev2RedirectPhase.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNRALoadBalancing) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enableVpnLoadBalancing"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("groupSettings.groupIPv4Address"); value.Exists() {
		data.Ipv4GroupAddress = types.StringValue(value.String())
	} else {
		data.Ipv4GroupAddress = types.StringNull()
	}
	if value := res.Get("groupSettings.groupIPv6Address"); value.Exists() {
		data.Ipv6GroupAddress = types.StringValue(value.String())
	} else {
		data.Ipv6GroupAddress = types.StringNull()
	}
	if value := res.Get("groupSettings.communicationInterface.id"); value.Exists() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if value := res.Get("groupSettings.communicationUdpPort"); value.Exists() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get("groupSettings.ipsecEncryption.enable"); value.Exists() {
		data.Ipsec = types.BoolValue(value.Bool())
	} else {
		data.Ipsec = types.BoolNull()
	}
	if value := res.Get("redirectSettings.redirectUsingFqdn"); value.Exists() {
		data.SendFqdnToPeerDevicesInsteadOfIp = types.BoolValue(value.Bool())
	} else {
		data.SendFqdnToPeerDevicesInsteadOfIp = types.BoolNull()
	}
	if value := res.Get("redirectSettings.ikev2RedirectPhase"); value.Exists() {
		data.Ikev2RedirectPhase = types.StringValue(value.String())
	} else {
		data.Ikev2RedirectPhase = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNRALoadBalancing) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enableVpnLoadBalancing"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("groupSettings.groupIPv4Address"); value.Exists() && !data.Ipv4GroupAddress.IsNull() {
		data.Ipv4GroupAddress = types.StringValue(value.String())
	} else {
		data.Ipv4GroupAddress = types.StringNull()
	}
	if value := res.Get("groupSettings.groupIPv6Address"); value.Exists() && !data.Ipv6GroupAddress.IsNull() {
		data.Ipv6GroupAddress = types.StringValue(value.String())
	} else {
		data.Ipv6GroupAddress = types.StringNull()
	}
	if value := res.Get("groupSettings.communicationInterface.id"); value.Exists() && !data.InterfaceId.IsNull() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if value := res.Get("groupSettings.communicationUdpPort"); value.Exists() && !data.Port.IsNull() {
		data.Port = types.Int64Value(value.Int())
	} else {
		data.Port = types.Int64Null()
	}
	if value := res.Get("groupSettings.ipsecEncryption.enable"); value.Exists() && !data.Ipsec.IsNull() {
		data.Ipsec = types.BoolValue(value.Bool())
	} else {
		data.Ipsec = types.BoolNull()
	}
	if value := res.Get("redirectSettings.redirectUsingFqdn"); value.Exists() && !data.SendFqdnToPeerDevicesInsteadOfIp.IsNull() {
		data.SendFqdnToPeerDevicesInsteadOfIp = types.BoolValue(value.Bool())
	} else {
		data.SendFqdnToPeerDevicesInsteadOfIp = types.BoolNull()
	}
	if value := res.Get("redirectSettings.ikev2RedirectPhase"); value.Exists() && !data.Ikev2RedirectPhase.IsNull() {
		data.Ikev2RedirectPhase = types.StringValue(value.String())
	} else {
		data.Ikev2RedirectPhase = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNRALoadBalancing) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data VPNRALoadBalancing) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	return body
}

// End of section. //template:end toBodyPutDelete
