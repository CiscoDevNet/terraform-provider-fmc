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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type GroupPolicy struct {
	Id                                  types.String                  `tfsdk:"id"`
	Domain                              types.String                  `tfsdk:"domain"`
	Name                                types.String                  `tfsdk:"name"`
	Type                                types.String                  `tfsdk:"type"`
	Description                         types.String                  `tfsdk:"description"`
	EnableSslProtocol                   types.Bool                    `tfsdk:"enable_ssl_protocol"`
	EnableIpsecIkev2Protocol            types.Bool                    `tfsdk:"enable_ipsec_ikev2_protocol"`
	Ipv4AddressPools                    []GroupPolicyIpv4AddressPools `tfsdk:"ipv4_address_pools"`
	Banner                              types.String                  `tfsdk:"banner"`
	PrimaryDnsServer                    types.String                  `tfsdk:"primary_dns_server"`
	SecondaryDnsServer                  types.String                  `tfsdk:"secondary_dns_server"`
	PrimaryWinsServer                   types.String                  `tfsdk:"primary_wins_server"`
	SecondaryWinsServer                 types.String                  `tfsdk:"secondary_wins_server"`
	DhcpScopeNetworkId                  types.String                  `tfsdk:"dhcp_scope_network_id"`
	DefaultDomain                       types.String                  `tfsdk:"default_domain"`
	Ipv4SplitTunnelPolicy               types.String                  `tfsdk:"ipv4_split_tunnel_policy"`
	Ipv6SplitTunnelPolicy               types.String                  `tfsdk:"ipv6_split_tunnel_policy"`
	SplitTunnelAclId                    types.String                  `tfsdk:"split_tunnel_acl_id"`
	SplitDNSRequestPolicy               types.String                  `tfsdk:"split_d_n_s_request_policy"`
	SplitDNSDomainList                  types.String                  `tfsdk:"split_d_n_s_domain_list"`
	SecureClientClientProfileId         types.String                  `tfsdk:"secure_client_client_profile_id"`
	SecureClientManagementProfileId     types.String                  `tfsdk:"secure_client_management_profile_id"`
	SslCompression                      types.String                  `tfsdk:"ssl_compression"`
	DtlsCompression                     types.String                  `tfsdk:"dtls_compression"`
	MtuSize                             types.Int64                   `tfsdk:"mtu_size"`
	IgnoreDfBit                         types.Bool                    `tfsdk:"ignore_df_bit"`
	EnableKeepAliveMessages             types.Bool                    `tfsdk:"enable_keep_alive_messages"`
	KeepAliveMessageInterval            types.Int64                   `tfsdk:"keep_alive_message_interval"`
	EnableGatewayDpd                    types.Bool                    `tfsdk:"enable_gateway_dpd"`
	GatewayDpdInterval                  types.Int64                   `tfsdk:"gateway_dpd_interval"`
	EnableClientDpd                     types.Bool                    `tfsdk:"enable_client_dpd"`
	ClientDpdInterval                   types.Int64                   `tfsdk:"client_dpd_interval"`
	ClientBypassProtocol                types.Bool                    `tfsdk:"client_bypass_protocol"`
	EnableSslRekey                      types.Bool                    `tfsdk:"enable_ssl_rekey"`
	RekeyMethod                         types.String                  `tfsdk:"rekey_method"`
	RekeyInterval                       types.Int64                   `tfsdk:"rekey_interval"`
	ClientFirewallPrivateNetworkRulesId types.String                  `tfsdk:"client_firewall_private_network_rules_id"`
	ClientFirewallPublicNetworkRulesId  types.String                  `tfsdk:"client_firewall_public_network_rules_id"`
	CustomAttributes                    []GroupPolicyCustomAttributes `tfsdk:"custom_attributes"`
	TrafficFilterAclId                  types.String                  `tfsdk:"traffic_filter_acl_id"`
	RestrictVpnToVlanId                 types.Int64                   `tfsdk:"restrict_vpn_to_vlan_id"`
	AccessHoursId                       types.String                  `tfsdk:"access_hours_id"`
	SimultaneousLoginPerUser            types.Int64                   `tfsdk:"simultaneous_login_per_user"`
	MaxConnectionTime                   types.Int64                   `tfsdk:"max_connection_time"`
	MaxConnectionTimeAlertInterval      types.Int64                   `tfsdk:"max_connection_time_alert_interval"`
	VpnIdleTimeout                      types.Int64                   `tfsdk:"vpn_idle_timeout"`
	VpnIdleTimeoutAlertInterval         types.Int64                   `tfsdk:"vpn_idle_timeout_alert_interval"`
}

type GroupPolicyIpv4AddressPools struct {
	Id types.String `tfsdk:"id"`
}

type GroupPolicyCustomAttributes struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data GroupPolicy) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/grouppolicies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data GroupPolicy) toBody(ctx context.Context, state GroupPolicy) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.EnableSslProtocol.IsNull() {
		body, _ = sjson.Set(body, "enableSSLProtocol", data.EnableSslProtocol.ValueBool())
	}
	if !data.EnableIpsecIkev2Protocol.IsNull() {
		body, _ = sjson.Set(body, "enableIPsecIKEv2Protocol", data.EnableIpsecIkev2Protocol.ValueBool())
	}
	if len(data.Ipv4AddressPools) > 0 {
		body, _ = sjson.Set(body, "generalSettings.addressAssignment.ipv4LocalAddressPool", []interface{}{})
		for _, item := range data.Ipv4AddressPools {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "generalSettings.addressAssignment.ipv4LocalAddressPool.-1", itemBody)
		}
	}
	if !data.Banner.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.banner", data.Banner.ValueString())
	}
	if !data.PrimaryDnsServer.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.primaryDNSServer.id", data.PrimaryDnsServer.ValueString())
	}
	if !data.SecondaryDnsServer.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.secondaryDNSServer.id", data.SecondaryDnsServer.ValueString())
	}
	if !data.PrimaryWinsServer.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.primaryWINSServer.id", data.PrimaryWinsServer.ValueString())
	}
	if !data.SecondaryWinsServer.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.secondaryWINSSserver.id", data.SecondaryWinsServer.ValueString())
	}
	if !data.DhcpScopeNetworkId.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.addressAssignment.dhcpScope.id", data.DhcpScopeNetworkId.ValueString())
	}
	if !data.DefaultDomain.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.addressAssignment.defaultDomainName", data.DefaultDomain.ValueString())
	}
	if !data.Ipv4SplitTunnelPolicy.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.splitTunnelSettings.ipv4SplitTunnelPolicy", data.Ipv4SplitTunnelPolicy.ValueString())
	}
	if !data.Ipv6SplitTunnelPolicy.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.splitTunnelSettings.ipv6SplitTunnelPolicy", data.Ipv6SplitTunnelPolicy.ValueString())
	}
	if !data.SplitTunnelAclId.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.splitTunnelSettings.splitTunnelACL", data.SplitTunnelAclId.ValueString())
	}
	if !data.SplitDNSRequestPolicy.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.splitTunnelSettings.splitDNSRequestPolicy", data.SplitDNSRequestPolicy.ValueString())
	}
	if !data.SplitDNSDomainList.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.splitTunnelSettings.splitDNSDomainList", data.SplitDNSDomainList.ValueString())
	}
	if !data.SecureClientClientProfileId.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.vpnClientProfile.id", data.SecureClientClientProfileId.ValueString())
	}
	if !data.SecureClientManagementProfileId.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.managementProfile.id", data.SecureClientManagementProfileId.ValueString())
	}
	if !data.SslCompression.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.sslSettings.sslCompression", data.SslCompression.ValueString())
	}
	if !data.DtlsCompression.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.sslSettings.dtlsCompression", data.DtlsCompression.ValueString())
	}
	if !data.MtuSize.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.sslSettings.mtuSize", data.MtuSize.ValueInt64())
	}
	if !data.IgnoreDfBit.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.sslSettings.ignoreDFBit", data.IgnoreDfBit.ValueBool())
	}
	if !data.EnableKeepAliveMessages.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.enableKeepAliveMessages", data.EnableKeepAliveMessages.ValueBool())
	}
	if !data.KeepAliveMessageInterval.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.keepAliveMessageInterval", data.KeepAliveMessageInterval.ValueInt64())
	}
	if !data.EnableGatewayDpd.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.enableGatewayDPD", data.EnableGatewayDpd.ValueBool())
	}
	if !data.GatewayDpdInterval.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.gatewayDPDInterval", data.GatewayDpdInterval.ValueInt64())
	}
	if !data.EnableClientDpd.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.enableClientDPD", data.EnableClientDpd.ValueBool())
	}
	if !data.ClientDpdInterval.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.clientDPDInterval", data.ClientDpdInterval.ValueInt64())
	}
	if !data.ClientBypassProtocol.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.bypassUnsupportProtocol", data.ClientBypassProtocol.ValueBool())
	}
	if !data.EnableSslRekey.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.enableSSLRekey", data.EnableSslRekey.ValueBool())
	}
	if !data.RekeyMethod.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.rekeyMethod", data.RekeyMethod.ValueString())
	}
	if !data.RekeyInterval.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.rekeyInterval", data.RekeyInterval.ValueInt64())
	}
	if !data.ClientFirewallPrivateNetworkRulesId.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.clientFirewallPrivateNetworkRules.id", data.ClientFirewallPrivateNetworkRulesId.ValueString())
	}
	if !data.ClientFirewallPublicNetworkRulesId.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.clientFirewallPublicNetworkRules.id", data.ClientFirewallPublicNetworkRulesId.ValueString())
	}
	if len(data.CustomAttributes) > 0 {
		body, _ = sjson.Set(body, "anyConnectSettings.customAttributes", []interface{}{})
		for _, item := range data.CustomAttributes {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "customAttributeObject.id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "customAttributeObject.anyConnectAttribute", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "anyConnectSettings.customAttributes.-1", itemBody)
		}
	}
	if !data.TrafficFilterAclId.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.vpnTrafficFilterACL.id", data.TrafficFilterAclId.ValueString())
	}
	if !data.RestrictVpnToVlanId.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.restrictVPNToVLANId", data.RestrictVpnToVlanId.ValueInt64())
	}
	if !data.AccessHoursId.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.accessHours.id", data.AccessHoursId.ValueString())
	}
	if !data.SimultaneousLoginPerUser.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.simultaneousLoginPerUser", data.SimultaneousLoginPerUser.ValueInt64())
	}
	if !data.MaxConnectionTime.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.maxConnectionTimeout", data.MaxConnectionTime.ValueInt64())
	}
	if !data.MaxConnectionTimeAlertInterval.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.maxConnectionTimeAlertInterval", data.MaxConnectionTimeAlertInterval.ValueInt64())
	}
	if !data.VpnIdleTimeout.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.vpnIdleTimeout", data.VpnIdleTimeout.ValueInt64())
	}
	if !data.VpnIdleTimeoutAlertInterval.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.vpnIdleTimeoutAlertInterval", data.VpnIdleTimeoutAlertInterval.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *GroupPolicy) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("enableSSLProtocol"); value.Exists() {
		data.EnableSslProtocol = types.BoolValue(value.Bool())
	} else {
		data.EnableSslProtocol = types.BoolNull()
	}
	if value := res.Get("enableIPsecIKEv2Protocol"); value.Exists() {
		data.EnableIpsecIkev2Protocol = types.BoolValue(value.Bool())
	} else {
		data.EnableIpsecIkev2Protocol = types.BoolNull()
	}
	if value := res.Get("generalSettings.addressAssignment.ipv4LocalAddressPool"); value.Exists() {
		data.Ipv4AddressPools = make([]GroupPolicyIpv4AddressPools, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := GroupPolicyIpv4AddressPools{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Ipv4AddressPools = append((*parent).Ipv4AddressPools, data)
			return true
		})
	}
	if value := res.Get("generalSettings.banner"); value.Exists() {
		data.Banner = types.StringValue(value.String())
	} else {
		data.Banner = types.StringNull()
	}
	if value := res.Get("generalSettings.primaryDNSServer.id"); value.Exists() {
		data.PrimaryDnsServer = types.StringValue(value.String())
	} else {
		data.PrimaryDnsServer = types.StringNull()
	}
	if value := res.Get("generalSettings.secondaryDNSServer.id"); value.Exists() {
		data.SecondaryDnsServer = types.StringValue(value.String())
	} else {
		data.SecondaryDnsServer = types.StringNull()
	}
	if value := res.Get("generalSettings.primaryWINSServer.id"); value.Exists() {
		data.PrimaryWinsServer = types.StringValue(value.String())
	} else {
		data.PrimaryWinsServer = types.StringNull()
	}
	if value := res.Get("generalSettings.secondaryWINSSserver.id"); value.Exists() {
		data.SecondaryWinsServer = types.StringValue(value.String())
	} else {
		data.SecondaryWinsServer = types.StringNull()
	}
	if value := res.Get("generalSettings.addressAssignment.dhcpScope.id"); value.Exists() {
		data.DhcpScopeNetworkId = types.StringValue(value.String())
	} else {
		data.DhcpScopeNetworkId = types.StringNull()
	}
	if value := res.Get("generalSettings.addressAssignment.defaultDomainName"); value.Exists() {
		data.DefaultDomain = types.StringValue(value.String())
	} else {
		data.DefaultDomain = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.ipv4SplitTunnelPolicy"); value.Exists() {
		data.Ipv4SplitTunnelPolicy = types.StringValue(value.String())
	} else {
		data.Ipv4SplitTunnelPolicy = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.ipv6SplitTunnelPolicy"); value.Exists() {
		data.Ipv6SplitTunnelPolicy = types.StringValue(value.String())
	} else {
		data.Ipv6SplitTunnelPolicy = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitTunnelACL"); value.Exists() {
		data.SplitTunnelAclId = types.StringValue(value.String())
	} else {
		data.SplitTunnelAclId = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitDNSRequestPolicy"); value.Exists() {
		data.SplitDNSRequestPolicy = types.StringValue(value.String())
	} else {
		data.SplitDNSRequestPolicy = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitDNSDomainList"); value.Exists() {
		data.SplitDNSDomainList = types.StringValue(value.String())
	} else {
		data.SplitDNSDomainList = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.vpnClientProfile.id"); value.Exists() {
		data.SecureClientClientProfileId = types.StringValue(value.String())
	} else {
		data.SecureClientClientProfileId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.managementProfile.id"); value.Exists() {
		data.SecureClientManagementProfileId = types.StringValue(value.String())
	} else {
		data.SecureClientManagementProfileId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.sslSettings.sslCompression"); value.Exists() {
		data.SslCompression = types.StringValue(value.String())
	} else {
		data.SslCompression = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.sslSettings.dtlsCompression"); value.Exists() {
		data.DtlsCompression = types.StringValue(value.String())
	} else {
		data.DtlsCompression = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.sslSettings.mtuSize"); value.Exists() {
		data.MtuSize = types.Int64Value(value.Int())
	} else {
		data.MtuSize = types.Int64Value(1406)
	}
	if value := res.Get("anyConnectSettings.sslSettings.ignoreDFBit"); value.Exists() {
		data.IgnoreDfBit = types.BoolValue(value.Bool())
	} else {
		data.IgnoreDfBit = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableKeepAliveMessages"); value.Exists() {
		data.EnableKeepAliveMessages = types.BoolValue(value.Bool())
	} else {
		data.EnableKeepAliveMessages = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.keepAliveMessageInterval"); value.Exists() {
		data.KeepAliveMessageInterval = types.Int64Value(value.Int())
	} else {
		data.KeepAliveMessageInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableGatewayDPD"); value.Exists() {
		data.EnableGatewayDpd = types.BoolValue(value.Bool())
	} else {
		data.EnableGatewayDpd = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.gatewayDPDInterval"); value.Exists() {
		data.GatewayDpdInterval = types.Int64Value(value.Int())
	} else {
		data.GatewayDpdInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableClientDPD"); value.Exists() {
		data.EnableClientDpd = types.BoolValue(value.Bool())
	} else {
		data.EnableClientDpd = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientDPDInterval"); value.Exists() {
		data.ClientDpdInterval = types.Int64Value(value.Int())
	} else {
		data.ClientDpdInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.bypassUnsupportProtocol"); value.Exists() {
		data.ClientBypassProtocol = types.BoolValue(value.Bool())
	} else {
		data.ClientBypassProtocol = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableSSLRekey"); value.Exists() {
		data.EnableSslRekey = types.BoolValue(value.Bool())
	} else {
		data.EnableSslRekey = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.rekeyMethod"); value.Exists() {
		data.RekeyMethod = types.StringValue(value.String())
	} else {
		data.RekeyMethod = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.rekeyInterval"); value.Exists() {
		data.RekeyInterval = types.Int64Value(value.Int())
	} else {
		data.RekeyInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientFirewallPrivateNetworkRules.id"); value.Exists() {
		data.ClientFirewallPrivateNetworkRulesId = types.StringValue(value.String())
	} else {
		data.ClientFirewallPrivateNetworkRulesId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientFirewallPublicNetworkRules.id"); value.Exists() {
		data.ClientFirewallPublicNetworkRulesId = types.StringValue(value.String())
	} else {
		data.ClientFirewallPublicNetworkRulesId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.customAttributes"); value.Exists() {
		data.CustomAttributes = make([]GroupPolicyCustomAttributes, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := GroupPolicyCustomAttributes{}
			if value := res.Get("customAttributeObject.id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("customAttributeObject.anyConnectAttribute"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			(*parent).CustomAttributes = append((*parent).CustomAttributes, data)
			return true
		})
	}
	if value := res.Get("advancedSettings.vpnTrafficFilterACL.id"); value.Exists() {
		data.TrafficFilterAclId = types.StringValue(value.String())
	} else {
		data.TrafficFilterAclId = types.StringNull()
	}
	if value := res.Get("advancedSettings.restrictVPNToVLANId"); value.Exists() {
		data.RestrictVpnToVlanId = types.Int64Value(value.Int())
	} else {
		data.RestrictVpnToVlanId = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.accessHours.id"); value.Exists() {
		data.AccessHoursId = types.StringValue(value.String())
	} else {
		data.AccessHoursId = types.StringNull()
	}
	if value := res.Get("advancedSettings.sessionSettings.simultaneousLoginPerUser"); value.Exists() {
		data.SimultaneousLoginPerUser = types.Int64Value(value.Int())
	} else {
		data.SimultaneousLoginPerUser = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.maxConnectionTimeout"); value.Exists() {
		data.MaxConnectionTime = types.Int64Value(value.Int())
	} else {
		data.MaxConnectionTime = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.maxConnectionTimeAlertInterval"); value.Exists() {
		data.MaxConnectionTimeAlertInterval = types.Int64Value(value.Int())
	} else {
		data.MaxConnectionTimeAlertInterval = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.vpnIdleTimeout"); value.Exists() {
		data.VpnIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.VpnIdleTimeout = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.vpnIdleTimeoutAlertInterval"); value.Exists() {
		data.VpnIdleTimeoutAlertInterval = types.Int64Value(value.Int())
	} else {
		data.VpnIdleTimeoutAlertInterval = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *GroupPolicy) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("enableSSLProtocol"); value.Exists() && !data.EnableSslProtocol.IsNull() {
		data.EnableSslProtocol = types.BoolValue(value.Bool())
	} else {
		data.EnableSslProtocol = types.BoolNull()
	}
	if value := res.Get("enableIPsecIKEv2Protocol"); value.Exists() && !data.EnableIpsecIkev2Protocol.IsNull() {
		data.EnableIpsecIkev2Protocol = types.BoolValue(value.Bool())
	} else {
		data.EnableIpsecIkev2Protocol = types.BoolNull()
	}
	for i := 0; i < len(data.Ipv4AddressPools); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Ipv4AddressPools[i].Id.ValueString()}

		parent := &data
		data := (*parent).Ipv4AddressPools[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("generalSettings.addressAssignment.ipv4LocalAddressPool").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv4AddressPools[%d] = %+v",
				i,
				(*parent).Ipv4AddressPools[i],
			))
			(*parent).Ipv4AddressPools = slices.Delete((*parent).Ipv4AddressPools, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).Ipv4AddressPools[i] = data
	}
	if value := res.Get("generalSettings.banner"); value.Exists() && !data.Banner.IsNull() {
		data.Banner = types.StringValue(value.String())
	} else {
		data.Banner = types.StringNull()
	}
	if value := res.Get("generalSettings.primaryDNSServer.id"); value.Exists() && !data.PrimaryDnsServer.IsNull() {
		data.PrimaryDnsServer = types.StringValue(value.String())
	} else {
		data.PrimaryDnsServer = types.StringNull()
	}
	if value := res.Get("generalSettings.secondaryDNSServer.id"); value.Exists() && !data.SecondaryDnsServer.IsNull() {
		data.SecondaryDnsServer = types.StringValue(value.String())
	} else {
		data.SecondaryDnsServer = types.StringNull()
	}
	if value := res.Get("generalSettings.primaryWINSServer.id"); value.Exists() && !data.PrimaryWinsServer.IsNull() {
		data.PrimaryWinsServer = types.StringValue(value.String())
	} else {
		data.PrimaryWinsServer = types.StringNull()
	}
	if value := res.Get("generalSettings.secondaryWINSSserver.id"); value.Exists() && !data.SecondaryWinsServer.IsNull() {
		data.SecondaryWinsServer = types.StringValue(value.String())
	} else {
		data.SecondaryWinsServer = types.StringNull()
	}
	if value := res.Get("generalSettings.addressAssignment.dhcpScope.id"); value.Exists() && !data.DhcpScopeNetworkId.IsNull() {
		data.DhcpScopeNetworkId = types.StringValue(value.String())
	} else {
		data.DhcpScopeNetworkId = types.StringNull()
	}
	if value := res.Get("generalSettings.addressAssignment.defaultDomainName"); value.Exists() && !data.DefaultDomain.IsNull() {
		data.DefaultDomain = types.StringValue(value.String())
	} else {
		data.DefaultDomain = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.ipv4SplitTunnelPolicy"); value.Exists() && !data.Ipv4SplitTunnelPolicy.IsNull() {
		data.Ipv4SplitTunnelPolicy = types.StringValue(value.String())
	} else {
		data.Ipv4SplitTunnelPolicy = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.ipv6SplitTunnelPolicy"); value.Exists() && !data.Ipv6SplitTunnelPolicy.IsNull() {
		data.Ipv6SplitTunnelPolicy = types.StringValue(value.String())
	} else {
		data.Ipv6SplitTunnelPolicy = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitTunnelACL"); value.Exists() && !data.SplitTunnelAclId.IsNull() {
		data.SplitTunnelAclId = types.StringValue(value.String())
	} else {
		data.SplitTunnelAclId = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitDNSRequestPolicy"); value.Exists() && !data.SplitDNSRequestPolicy.IsNull() {
		data.SplitDNSRequestPolicy = types.StringValue(value.String())
	} else {
		data.SplitDNSRequestPolicy = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitDNSDomainList"); value.Exists() && !data.SplitDNSDomainList.IsNull() {
		data.SplitDNSDomainList = types.StringValue(value.String())
	} else {
		data.SplitDNSDomainList = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.vpnClientProfile.id"); value.Exists() && !data.SecureClientClientProfileId.IsNull() {
		data.SecureClientClientProfileId = types.StringValue(value.String())
	} else {
		data.SecureClientClientProfileId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.managementProfile.id"); value.Exists() && !data.SecureClientManagementProfileId.IsNull() {
		data.SecureClientManagementProfileId = types.StringValue(value.String())
	} else {
		data.SecureClientManagementProfileId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.sslSettings.sslCompression"); value.Exists() && !data.SslCompression.IsNull() {
		data.SslCompression = types.StringValue(value.String())
	} else {
		data.SslCompression = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.sslSettings.dtlsCompression"); value.Exists() && !data.DtlsCompression.IsNull() {
		data.DtlsCompression = types.StringValue(value.String())
	} else {
		data.DtlsCompression = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.sslSettings.mtuSize"); value.Exists() && !data.MtuSize.IsNull() {
		data.MtuSize = types.Int64Value(value.Int())
	} else if data.MtuSize.ValueInt64() != 1406 {
		data.MtuSize = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.sslSettings.ignoreDFBit"); value.Exists() && !data.IgnoreDfBit.IsNull() {
		data.IgnoreDfBit = types.BoolValue(value.Bool())
	} else {
		data.IgnoreDfBit = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableKeepAliveMessages"); value.Exists() && !data.EnableKeepAliveMessages.IsNull() {
		data.EnableKeepAliveMessages = types.BoolValue(value.Bool())
	} else {
		data.EnableKeepAliveMessages = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.keepAliveMessageInterval"); value.Exists() && !data.KeepAliveMessageInterval.IsNull() {
		data.KeepAliveMessageInterval = types.Int64Value(value.Int())
	} else {
		data.KeepAliveMessageInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableGatewayDPD"); value.Exists() && !data.EnableGatewayDpd.IsNull() {
		data.EnableGatewayDpd = types.BoolValue(value.Bool())
	} else {
		data.EnableGatewayDpd = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.gatewayDPDInterval"); value.Exists() && !data.GatewayDpdInterval.IsNull() {
		data.GatewayDpdInterval = types.Int64Value(value.Int())
	} else {
		data.GatewayDpdInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableClientDPD"); value.Exists() && !data.EnableClientDpd.IsNull() {
		data.EnableClientDpd = types.BoolValue(value.Bool())
	} else {
		data.EnableClientDpd = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientDPDInterval"); value.Exists() && !data.ClientDpdInterval.IsNull() {
		data.ClientDpdInterval = types.Int64Value(value.Int())
	} else {
		data.ClientDpdInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.bypassUnsupportProtocol"); value.Exists() && !data.ClientBypassProtocol.IsNull() {
		data.ClientBypassProtocol = types.BoolValue(value.Bool())
	} else {
		data.ClientBypassProtocol = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableSSLRekey"); value.Exists() && !data.EnableSslRekey.IsNull() {
		data.EnableSslRekey = types.BoolValue(value.Bool())
	} else {
		data.EnableSslRekey = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.rekeyMethod"); value.Exists() && !data.RekeyMethod.IsNull() {
		data.RekeyMethod = types.StringValue(value.String())
	} else {
		data.RekeyMethod = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.rekeyInterval"); value.Exists() && !data.RekeyInterval.IsNull() {
		data.RekeyInterval = types.Int64Value(value.Int())
	} else {
		data.RekeyInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientFirewallPrivateNetworkRules.id"); value.Exists() && !data.ClientFirewallPrivateNetworkRulesId.IsNull() {
		data.ClientFirewallPrivateNetworkRulesId = types.StringValue(value.String())
	} else {
		data.ClientFirewallPrivateNetworkRulesId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientFirewallPublicNetworkRules.id"); value.Exists() && !data.ClientFirewallPublicNetworkRulesId.IsNull() {
		data.ClientFirewallPublicNetworkRulesId = types.StringValue(value.String())
	} else {
		data.ClientFirewallPublicNetworkRulesId = types.StringNull()
	}
	for i := 0; i < len(data.CustomAttributes); i++ {
		keys := [...]string{"customAttributeObject.id", "customAttributeObject.anyConnectAttribute"}
		keyValues := [...]string{data.CustomAttributes[i].Id.ValueString(), data.CustomAttributes[i].Type.ValueString()}

		parent := &data
		data := (*parent).CustomAttributes[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("anyConnectSettings.customAttributes").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					if v.Get(keys[ik]).String() != keyValues[ik] {
						found = false
						break
					}
					found = true
				}
				if found {
					res = v
					return false
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("removing CustomAttributes[%d] = %+v",
				i,
				(*parent).CustomAttributes[i],
			))
			(*parent).CustomAttributes = slices.Delete((*parent).CustomAttributes, i, i+1)
			i--

			continue
		}
		if value := res.Get("customAttributeObject.id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("customAttributeObject.anyConnectAttribute"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		(*parent).CustomAttributes[i] = data
	}
	if value := res.Get("advancedSettings.vpnTrafficFilterACL.id"); value.Exists() && !data.TrafficFilterAclId.IsNull() {
		data.TrafficFilterAclId = types.StringValue(value.String())
	} else {
		data.TrafficFilterAclId = types.StringNull()
	}
	if value := res.Get("advancedSettings.restrictVPNToVLANId"); value.Exists() && !data.RestrictVpnToVlanId.IsNull() {
		data.RestrictVpnToVlanId = types.Int64Value(value.Int())
	} else {
		data.RestrictVpnToVlanId = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.accessHours.id"); value.Exists() && !data.AccessHoursId.IsNull() {
		data.AccessHoursId = types.StringValue(value.String())
	} else {
		data.AccessHoursId = types.StringNull()
	}
	if value := res.Get("advancedSettings.sessionSettings.simultaneousLoginPerUser"); value.Exists() && !data.SimultaneousLoginPerUser.IsNull() {
		data.SimultaneousLoginPerUser = types.Int64Value(value.Int())
	} else {
		data.SimultaneousLoginPerUser = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.maxConnectionTimeout"); value.Exists() && !data.MaxConnectionTime.IsNull() {
		data.MaxConnectionTime = types.Int64Value(value.Int())
	} else {
		data.MaxConnectionTime = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.maxConnectionTimeAlertInterval"); value.Exists() && !data.MaxConnectionTimeAlertInterval.IsNull() {
		data.MaxConnectionTimeAlertInterval = types.Int64Value(value.Int())
	} else {
		data.MaxConnectionTimeAlertInterval = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.vpnIdleTimeout"); value.Exists() && !data.VpnIdleTimeout.IsNull() {
		data.VpnIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.VpnIdleTimeout = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.vpnIdleTimeoutAlertInterval"); value.Exists() && !data.VpnIdleTimeoutAlertInterval.IsNull() {
		data.VpnIdleTimeoutAlertInterval = types.Int64Value(value.Int())
	} else {
		data.VpnIdleTimeoutAlertInterval = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *GroupPolicy) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
