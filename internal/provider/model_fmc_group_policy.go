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
	Id                                     types.String                              `tfsdk:"id"`
	Domain                                 types.String                              `tfsdk:"domain"`
	Name                                   types.String                              `tfsdk:"name"`
	Type                                   types.String                              `tfsdk:"type"`
	Description                            types.String                              `tfsdk:"description"`
	VpnProtocolSsl                         types.Bool                                `tfsdk:"vpn_protocol_ssl"`
	VpnProtocolIpsecIkev2                  types.Bool                                `tfsdk:"vpn_protocol_ipsec_ikev2"`
	Ipv4AddressPools                       []GroupPolicyIpv4AddressPools             `tfsdk:"ipv4_address_pools"`
	Banner                                 types.String                              `tfsdk:"banner"`
	PrimaryDnsServerHostId                 types.String                              `tfsdk:"primary_dns_server_host_id"`
	SecondaryDnsServerHostId               types.String                              `tfsdk:"secondary_dns_server_host_id"`
	PrimaryWinsServerHostId                types.String                              `tfsdk:"primary_wins_server_host_id"`
	SecondaryWinsServerHostId              types.String                              `tfsdk:"secondary_wins_server_host_id"`
	DhcpNetworkScopeNetworkObjectId        types.String                              `tfsdk:"dhcp_network_scope_network_object_id"`
	DefaultDomain                          types.String                              `tfsdk:"default_domain"`
	Ipv4SplitTunnelPolicy                  types.String                              `tfsdk:"ipv4_split_tunnel_policy"`
	Ipv6SplitTunnelPolicy                  types.String                              `tfsdk:"ipv6_split_tunnel_policy"`
	SplitTunnelAclId                       types.String                              `tfsdk:"split_tunnel_acl_id"`
	SplitTunnelAclType                     types.String                              `tfsdk:"split_tunnel_acl_type"`
	DnsRequestSplitTunnelPolicy            types.String                              `tfsdk:"dns_request_split_tunnel_policy"`
	SplitDnsDomainList                     types.String                              `tfsdk:"split_dns_domain_list"`
	SecureClientProfileId                  types.String                              `tfsdk:"secure_client_profile_id"`
	SecureClientManagementProfileId        types.String                              `tfsdk:"secure_client_management_profile_id"`
	SecureClientModules                    []GroupPolicySecureClientModules          `tfsdk:"secure_client_modules"`
	SslCompression                         types.String                              `tfsdk:"ssl_compression"`
	DtlsCompression                        types.String                              `tfsdk:"dtls_compression"`
	MtuSize                                types.Int64                               `tfsdk:"mtu_size"`
	IgnoreDfBit                            types.Bool                                `tfsdk:"ignore_df_bit"`
	KeepAliveMessages                      types.Bool                                `tfsdk:"keep_alive_messages"`
	KeepAliveMessagesInterval              types.Int64                               `tfsdk:"keep_alive_messages_interval"`
	GatewayDpd                             types.Bool                                `tfsdk:"gateway_dpd"`
	GatewayDpdInterval                     types.Int64                               `tfsdk:"gateway_dpd_interval"`
	ClientDpd                              types.Bool                                `tfsdk:"client_dpd"`
	ClientDpdInterval                      types.Int64                               `tfsdk:"client_dpd_interval"`
	ClientBypassProtocol                   types.Bool                                `tfsdk:"client_bypass_protocol"`
	SslRekey                               types.Bool                                `tfsdk:"ssl_rekey"`
	SslRekeyMethod                         types.String                              `tfsdk:"ssl_rekey_method"`
	SslRekeyInterval                       types.Int64                               `tfsdk:"ssl_rekey_interval"`
	ClientFirewallPrivateNetworkRulesAclId types.String                              `tfsdk:"client_firewall_private_network_rules_acl_id"`
	ClientFirewallPublicNetworkRulesAclId  types.String                              `tfsdk:"client_firewall_public_network_rules_acl_id"`
	SecureClientCustomAttributes           []GroupPolicySecureClientCustomAttributes `tfsdk:"secure_client_custom_attributes"`
	TrafficFilterAclId                     types.String                              `tfsdk:"traffic_filter_acl_id"`
	RestrictVpnToVlan                      types.Int64                               `tfsdk:"restrict_vpn_to_vlan"`
	AccessHoursTimeRangeId                 types.String                              `tfsdk:"access_hours_time_range_id"`
	SimultaneousLoginsPerUser              types.Int64                               `tfsdk:"simultaneous_logins_per_user"`
	MaximumConnectionTime                  types.Int64                               `tfsdk:"maximum_connection_time"`
	MaximumConnectionTimeAlertInterval     types.Int64                               `tfsdk:"maximum_connection_time_alert_interval"`
	VpnIdleTimeout                         types.Int64                               `tfsdk:"vpn_idle_timeout"`
	VpnIdleTimeoutAlertInterval            types.Int64                               `tfsdk:"vpn_idle_timeout_alert_interval"`
}

type GroupPolicyIpv4AddressPools struct {
	Id types.String `tfsdk:"id"`
}

type GroupPolicySecureClientModules struct {
	Type           types.String `tfsdk:"type"`
	ProfileId      types.String `tfsdk:"profile_id"`
	DownloadModule types.Bool   `tfsdk:"download_module"`
}

type GroupPolicySecureClientCustomAttributes struct {
	Id types.String `tfsdk:"id"`
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
	body, _ = sjson.Set(body, "type", "GroupPolicy")
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.VpnProtocolSsl.IsNull() {
		body, _ = sjson.Set(body, "enableSSLProtocol", data.VpnProtocolSsl.ValueBool())
	}
	if !data.VpnProtocolIpsecIkev2.IsNull() {
		body, _ = sjson.Set(body, "enableIPsecIKEv2Protocol", data.VpnProtocolIpsecIkev2.ValueBool())
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
	if !data.PrimaryDnsServerHostId.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.primaryDNSServer.id", data.PrimaryDnsServerHostId.ValueString())
	}
	if !data.SecondaryDnsServerHostId.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.secondaryDNSServer.id", data.SecondaryDnsServerHostId.ValueString())
	}
	if !data.PrimaryWinsServerHostId.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.primaryWINSServer.id", data.PrimaryWinsServerHostId.ValueString())
	}
	if !data.SecondaryWinsServerHostId.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.secondaryWINSServer.id", data.SecondaryWinsServerHostId.ValueString())
	}
	if !data.DhcpNetworkScopeNetworkObjectId.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.addressAssignment.dhcpScope.id", data.DhcpNetworkScopeNetworkObjectId.ValueString())
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
		body, _ = sjson.Set(body, "generalSettings.splitTunnelSettings.splitTunnelACL.id", data.SplitTunnelAclId.ValueString())
	}
	if !data.SplitTunnelAclType.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.splitTunnelSettings.splitTunnelACL.type", data.SplitTunnelAclType.ValueString())
	}
	if !data.DnsRequestSplitTunnelPolicy.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.splitTunnelSettings.splitDNSRequestPolicy", data.DnsRequestSplitTunnelPolicy.ValueString())
	}
	if !data.SplitDnsDomainList.IsNull() {
		body, _ = sjson.Set(body, "generalSettings.splitTunnelSettings.splitDNSDomainList", data.SplitDnsDomainList.ValueString())
	}
	if !data.SecureClientProfileId.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.vpnClientProfile.id", data.SecureClientProfileId.ValueString())
	}
	if !data.SecureClientManagementProfileId.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.managementProfile.id", data.SecureClientManagementProfileId.ValueString())
	}
	if len(data.SecureClientModules) > 0 {
		body, _ = sjson.Set(body, "anyConnectSettings.clientModules", []interface{}{})
		for _, item := range data.SecureClientModules {
			itemBody := ""
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "moduleType", item.Type.ValueString())
			}
			if !item.ProfileId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "moduleProfile.id", item.ProfileId.ValueString())
			}
			if !item.DownloadModule.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableModuleDownload", item.DownloadModule.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "anyConnectSettings.clientModules.-1", itemBody)
		}
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
	if !data.KeepAliveMessages.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.enableKeepAliveMessages", data.KeepAliveMessages.ValueBool())
	}
	if !data.KeepAliveMessagesInterval.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.keepAliveMessageInterval", data.KeepAliveMessagesInterval.ValueInt64())
	}
	if !data.GatewayDpd.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.enableGatewayDPD", data.GatewayDpd.ValueBool())
	}
	if !data.GatewayDpdInterval.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.gatewayDPDInterval", data.GatewayDpdInterval.ValueInt64())
	}
	if !data.ClientDpd.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.enableClientDPD", data.ClientDpd.ValueBool())
	}
	if !data.ClientDpdInterval.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.clientDPDInterval", data.ClientDpdInterval.ValueInt64())
	}
	if !data.ClientBypassProtocol.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.bypassUnsupportProtocol", data.ClientBypassProtocol.ValueBool())
	}
	if !data.SslRekey.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.enableSSLRekey", data.SslRekey.ValueBool())
	}
	if !data.SslRekeyMethod.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.rekeyMethod", data.SslRekeyMethod.ValueString())
	}
	if !data.SslRekeyInterval.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.rekeyInterval", data.SslRekeyInterval.ValueInt64())
	}
	if !data.ClientFirewallPrivateNetworkRulesAclId.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.clientFirewallPrivateNetworkRules.id", data.ClientFirewallPrivateNetworkRulesAclId.ValueString())
	}
	if !data.ClientFirewallPublicNetworkRulesAclId.IsNull() {
		body, _ = sjson.Set(body, "anyConnectSettings.connectionSettings.clientFirewallPublicNetworkRules.id", data.ClientFirewallPublicNetworkRulesAclId.ValueString())
	}
	if len(data.SecureClientCustomAttributes) > 0 {
		body, _ = sjson.Set(body, "anyConnectSettings.customAttributes", []interface{}{})
		for _, item := range data.SecureClientCustomAttributes {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "customAttributeObject.id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "anyConnectSettings.customAttributes.-1", itemBody)
		}
	}
	if !data.TrafficFilterAclId.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.vpnTrafficFilterACL.id", data.TrafficFilterAclId.ValueString())
	}
	if !data.RestrictVpnToVlan.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.restrictVPNToVLANId", data.RestrictVpnToVlan.ValueInt64())
	}
	if !data.AccessHoursTimeRangeId.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.accessHours.id", data.AccessHoursTimeRangeId.ValueString())
	}
	if !data.SimultaneousLoginsPerUser.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.simultaneousLoginPerUser", data.SimultaneousLoginsPerUser.ValueInt64())
	}
	if !data.MaximumConnectionTime.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.maxConnectionTimeout", data.MaximumConnectionTime.ValueInt64())
	}
	if !data.MaximumConnectionTimeAlertInterval.IsNull() {
		body, _ = sjson.Set(body, "advancedSettings.sessionSettings.maxConnectionTimeAlertInterval", data.MaximumConnectionTimeAlertInterval.ValueInt64())
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
		data.VpnProtocolSsl = types.BoolValue(value.Bool())
	} else {
		data.VpnProtocolSsl = types.BoolValue(true)
	}
	if value := res.Get("enableIPsecIKEv2Protocol"); value.Exists() {
		data.VpnProtocolIpsecIkev2 = types.BoolValue(value.Bool())
	} else {
		data.VpnProtocolIpsecIkev2 = types.BoolValue(true)
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
		data.PrimaryDnsServerHostId = types.StringValue(value.String())
	} else {
		data.PrimaryDnsServerHostId = types.StringNull()
	}
	if value := res.Get("generalSettings.secondaryDNSServer.id"); value.Exists() {
		data.SecondaryDnsServerHostId = types.StringValue(value.String())
	} else {
		data.SecondaryDnsServerHostId = types.StringNull()
	}
	if value := res.Get("generalSettings.primaryWINSServer.id"); value.Exists() {
		data.PrimaryWinsServerHostId = types.StringValue(value.String())
	} else {
		data.PrimaryWinsServerHostId = types.StringNull()
	}
	if value := res.Get("generalSettings.secondaryWINSServer.id"); value.Exists() {
		data.SecondaryWinsServerHostId = types.StringValue(value.String())
	} else {
		data.SecondaryWinsServerHostId = types.StringNull()
	}
	if value := res.Get("generalSettings.addressAssignment.dhcpScope.id"); value.Exists() {
		data.DhcpNetworkScopeNetworkObjectId = types.StringValue(value.String())
	} else {
		data.DhcpNetworkScopeNetworkObjectId = types.StringNull()
	}
	if value := res.Get("generalSettings.addressAssignment.defaultDomainName"); value.Exists() {
		data.DefaultDomain = types.StringValue(value.String())
	} else {
		data.DefaultDomain = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.ipv4SplitTunnelPolicy"); value.Exists() {
		data.Ipv4SplitTunnelPolicy = types.StringValue(value.String())
	} else {
		data.Ipv4SplitTunnelPolicy = types.StringValue("TUNNEL_ALL")
	}
	if value := res.Get("generalSettings.splitTunnelSettings.ipv6SplitTunnelPolicy"); value.Exists() {
		data.Ipv6SplitTunnelPolicy = types.StringValue(value.String())
	} else {
		data.Ipv6SplitTunnelPolicy = types.StringValue("TUNNEL_ALL")
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitTunnelACL.id"); value.Exists() {
		data.SplitTunnelAclId = types.StringValue(value.String())
	} else {
		data.SplitTunnelAclId = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitTunnelACL.type"); value.Exists() {
		data.SplitTunnelAclType = types.StringValue(value.String())
	} else {
		data.SplitTunnelAclType = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitDNSRequestPolicy"); value.Exists() {
		data.DnsRequestSplitTunnelPolicy = types.StringValue(value.String())
	} else {
		data.DnsRequestSplitTunnelPolicy = types.StringValue("USE_SPLIT_TUNNEL_SETTING")
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitDNSDomainList"); value.Exists() {
		data.SplitDnsDomainList = types.StringValue(value.String())
	} else {
		data.SplitDnsDomainList = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.vpnClientProfile.id"); value.Exists() {
		data.SecureClientProfileId = types.StringValue(value.String())
	} else {
		data.SecureClientProfileId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.managementProfile.id"); value.Exists() {
		data.SecureClientManagementProfileId = types.StringValue(value.String())
	} else {
		data.SecureClientManagementProfileId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.clientModules"); value.Exists() {
		data.SecureClientModules = make([]GroupPolicySecureClientModules, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := GroupPolicySecureClientModules{}
			if value := res.Get("moduleType"); value.Exists() {
				data.Type = types.StringValue(value.String())
			} else {
				data.Type = types.StringNull()
			}
			if value := res.Get("moduleProfile.id"); value.Exists() {
				data.ProfileId = types.StringValue(value.String())
			} else {
				data.ProfileId = types.StringNull()
			}
			if value := res.Get("enableModuleDownload"); value.Exists() {
				data.DownloadModule = types.BoolValue(value.Bool())
			} else {
				data.DownloadModule = types.BoolValue(true)
			}
			(*parent).SecureClientModules = append((*parent).SecureClientModules, data)
			return true
		})
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
		data.KeepAliveMessages = types.BoolValue(value.Bool())
	} else {
		data.KeepAliveMessages = types.BoolValue(true)
	}
	if value := res.Get("anyConnectSettings.connectionSettings.keepAliveMessageInterval"); value.Exists() {
		data.KeepAliveMessagesInterval = types.Int64Value(value.Int())
	} else {
		data.KeepAliveMessagesInterval = types.Int64Value(20)
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableGatewayDPD"); value.Exists() {
		data.GatewayDpd = types.BoolValue(value.Bool())
	} else {
		data.GatewayDpd = types.BoolValue(true)
	}
	if value := res.Get("anyConnectSettings.connectionSettings.gatewayDPDInterval"); value.Exists() {
		data.GatewayDpdInterval = types.Int64Value(value.Int())
	} else {
		data.GatewayDpdInterval = types.Int64Value(30)
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableClientDPD"); value.Exists() {
		data.ClientDpd = types.BoolValue(value.Bool())
	} else {
		data.ClientDpd = types.BoolValue(true)
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientDPDInterval"); value.Exists() {
		data.ClientDpdInterval = types.Int64Value(value.Int())
	} else {
		data.ClientDpdInterval = types.Int64Value(30)
	}
	if value := res.Get("anyConnectSettings.connectionSettings.bypassUnsupportProtocol"); value.Exists() {
		data.ClientBypassProtocol = types.BoolValue(value.Bool())
	} else {
		data.ClientBypassProtocol = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableSSLRekey"); value.Exists() {
		data.SslRekey = types.BoolValue(value.Bool())
	} else {
		data.SslRekey = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.rekeyMethod"); value.Exists() {
		data.SslRekeyMethod = types.StringValue(value.String())
	} else {
		data.SslRekeyMethod = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.rekeyInterval"); value.Exists() {
		data.SslRekeyInterval = types.Int64Value(value.Int())
	} else {
		data.SslRekeyInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientFirewallPrivateNetworkRules.id"); value.Exists() {
		data.ClientFirewallPrivateNetworkRulesAclId = types.StringValue(value.String())
	} else {
		data.ClientFirewallPrivateNetworkRulesAclId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientFirewallPublicNetworkRules.id"); value.Exists() {
		data.ClientFirewallPublicNetworkRulesAclId = types.StringValue(value.String())
	} else {
		data.ClientFirewallPublicNetworkRulesAclId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.customAttributes"); value.Exists() {
		data.SecureClientCustomAttributes = make([]GroupPolicySecureClientCustomAttributes, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := GroupPolicySecureClientCustomAttributes{}
			if value := res.Get("customAttributeObject.id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).SecureClientCustomAttributes = append((*parent).SecureClientCustomAttributes, data)
			return true
		})
	}
	if value := res.Get("advancedSettings.vpnTrafficFilterACL.id"); value.Exists() {
		data.TrafficFilterAclId = types.StringValue(value.String())
	} else {
		data.TrafficFilterAclId = types.StringNull()
	}
	if value := res.Get("advancedSettings.restrictVPNToVLANId"); value.Exists() {
		data.RestrictVpnToVlan = types.Int64Value(value.Int())
	} else {
		data.RestrictVpnToVlan = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.accessHours.id"); value.Exists() {
		data.AccessHoursTimeRangeId = types.StringValue(value.String())
	} else {
		data.AccessHoursTimeRangeId = types.StringNull()
	}
	if value := res.Get("advancedSettings.sessionSettings.simultaneousLoginPerUser"); value.Exists() {
		data.SimultaneousLoginsPerUser = types.Int64Value(value.Int())
	} else {
		data.SimultaneousLoginsPerUser = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.maxConnectionTimeout"); value.Exists() {
		data.MaximumConnectionTime = types.Int64Value(value.Int())
	} else {
		data.MaximumConnectionTime = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.maxConnectionTimeAlertInterval"); value.Exists() {
		data.MaximumConnectionTimeAlertInterval = types.Int64Value(value.Int())
	} else {
		data.MaximumConnectionTimeAlertInterval = types.Int64Null()
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
	if value := res.Get("enableSSLProtocol"); value.Exists() && !data.VpnProtocolSsl.IsNull() {
		data.VpnProtocolSsl = types.BoolValue(value.Bool())
	} else if data.VpnProtocolSsl.ValueBool() != true {
		data.VpnProtocolSsl = types.BoolNull()
	}
	if value := res.Get("enableIPsecIKEv2Protocol"); value.Exists() && !data.VpnProtocolIpsecIkev2.IsNull() {
		data.VpnProtocolIpsecIkev2 = types.BoolValue(value.Bool())
	} else if data.VpnProtocolIpsecIkev2.ValueBool() != true {
		data.VpnProtocolIpsecIkev2 = types.BoolNull()
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
	if value := res.Get("generalSettings.primaryDNSServer.id"); value.Exists() && !data.PrimaryDnsServerHostId.IsNull() {
		data.PrimaryDnsServerHostId = types.StringValue(value.String())
	} else {
		data.PrimaryDnsServerHostId = types.StringNull()
	}
	if value := res.Get("generalSettings.secondaryDNSServer.id"); value.Exists() && !data.SecondaryDnsServerHostId.IsNull() {
		data.SecondaryDnsServerHostId = types.StringValue(value.String())
	} else {
		data.SecondaryDnsServerHostId = types.StringNull()
	}
	if value := res.Get("generalSettings.primaryWINSServer.id"); value.Exists() && !data.PrimaryWinsServerHostId.IsNull() {
		data.PrimaryWinsServerHostId = types.StringValue(value.String())
	} else {
		data.PrimaryWinsServerHostId = types.StringNull()
	}
	if value := res.Get("generalSettings.secondaryWINSServer.id"); value.Exists() && !data.SecondaryWinsServerHostId.IsNull() {
		data.SecondaryWinsServerHostId = types.StringValue(value.String())
	} else {
		data.SecondaryWinsServerHostId = types.StringNull()
	}
	if value := res.Get("generalSettings.addressAssignment.dhcpScope.id"); value.Exists() && !data.DhcpNetworkScopeNetworkObjectId.IsNull() {
		data.DhcpNetworkScopeNetworkObjectId = types.StringValue(value.String())
	} else {
		data.DhcpNetworkScopeNetworkObjectId = types.StringNull()
	}
	if value := res.Get("generalSettings.addressAssignment.defaultDomainName"); value.Exists() && !data.DefaultDomain.IsNull() {
		data.DefaultDomain = types.StringValue(value.String())
	} else {
		data.DefaultDomain = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.ipv4SplitTunnelPolicy"); value.Exists() && !data.Ipv4SplitTunnelPolicy.IsNull() {
		data.Ipv4SplitTunnelPolicy = types.StringValue(value.String())
	} else if data.Ipv4SplitTunnelPolicy.ValueString() != "TUNNEL_ALL" {
		data.Ipv4SplitTunnelPolicy = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.ipv6SplitTunnelPolicy"); value.Exists() && !data.Ipv6SplitTunnelPolicy.IsNull() {
		data.Ipv6SplitTunnelPolicy = types.StringValue(value.String())
	} else if data.Ipv6SplitTunnelPolicy.ValueString() != "TUNNEL_ALL" {
		data.Ipv6SplitTunnelPolicy = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitTunnelACL.id"); value.Exists() && !data.SplitTunnelAclId.IsNull() {
		data.SplitTunnelAclId = types.StringValue(value.String())
	} else {
		data.SplitTunnelAclId = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitTunnelACL.type"); value.Exists() && !data.SplitTunnelAclType.IsNull() {
		data.SplitTunnelAclType = types.StringValue(value.String())
	} else {
		data.SplitTunnelAclType = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitDNSRequestPolicy"); value.Exists() && !data.DnsRequestSplitTunnelPolicy.IsNull() {
		data.DnsRequestSplitTunnelPolicy = types.StringValue(value.String())
	} else if data.DnsRequestSplitTunnelPolicy.ValueString() != "USE_SPLIT_TUNNEL_SETTING" {
		data.DnsRequestSplitTunnelPolicy = types.StringNull()
	}
	if value := res.Get("generalSettings.splitTunnelSettings.splitDNSDomainList"); value.Exists() && !data.SplitDnsDomainList.IsNull() {
		data.SplitDnsDomainList = types.StringValue(value.String())
	} else {
		data.SplitDnsDomainList = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.vpnClientProfile.id"); value.Exists() && !data.SecureClientProfileId.IsNull() {
		data.SecureClientProfileId = types.StringValue(value.String())
	} else {
		data.SecureClientProfileId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.managementProfile.id"); value.Exists() && !data.SecureClientManagementProfileId.IsNull() {
		data.SecureClientManagementProfileId = types.StringValue(value.String())
	} else {
		data.SecureClientManagementProfileId = types.StringNull()
	}
	for i := 0; i < len(data.SecureClientModules); i++ {
		keys := [...]string{"moduleType"}
		keyValues := [...]string{data.SecureClientModules[i].Type.ValueString()}

		parent := &data
		data := (*parent).SecureClientModules[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("anyConnectSettings.clientModules").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing SecureClientModules[%d] = %+v",
				i,
				(*parent).SecureClientModules[i],
			))
			(*parent).SecureClientModules = slices.Delete((*parent).SecureClientModules, i, i+1)
			i--

			continue
		}
		if value := res.Get("moduleType"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("moduleProfile.id"); value.Exists() && !data.ProfileId.IsNull() {
			data.ProfileId = types.StringValue(value.String())
		} else {
			data.ProfileId = types.StringNull()
		}
		if value := res.Get("enableModuleDownload"); value.Exists() && !data.DownloadModule.IsNull() {
			data.DownloadModule = types.BoolValue(value.Bool())
		} else if data.DownloadModule.ValueBool() != true {
			data.DownloadModule = types.BoolNull()
		}
		(*parent).SecureClientModules[i] = data
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
	if value := res.Get("anyConnectSettings.connectionSettings.enableKeepAliveMessages"); value.Exists() && !data.KeepAliveMessages.IsNull() {
		data.KeepAliveMessages = types.BoolValue(value.Bool())
	} else if data.KeepAliveMessages.ValueBool() != true {
		data.KeepAliveMessages = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.keepAliveMessageInterval"); value.Exists() && !data.KeepAliveMessagesInterval.IsNull() {
		data.KeepAliveMessagesInterval = types.Int64Value(value.Int())
	} else if data.KeepAliveMessagesInterval.ValueInt64() != 20 {
		data.KeepAliveMessagesInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableGatewayDPD"); value.Exists() && !data.GatewayDpd.IsNull() {
		data.GatewayDpd = types.BoolValue(value.Bool())
	} else if data.GatewayDpd.ValueBool() != true {
		data.GatewayDpd = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.gatewayDPDInterval"); value.Exists() && !data.GatewayDpdInterval.IsNull() {
		data.GatewayDpdInterval = types.Int64Value(value.Int())
	} else if data.GatewayDpdInterval.ValueInt64() != 30 {
		data.GatewayDpdInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableClientDPD"); value.Exists() && !data.ClientDpd.IsNull() {
		data.ClientDpd = types.BoolValue(value.Bool())
	} else if data.ClientDpd.ValueBool() != true {
		data.ClientDpd = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientDPDInterval"); value.Exists() && !data.ClientDpdInterval.IsNull() {
		data.ClientDpdInterval = types.Int64Value(value.Int())
	} else if data.ClientDpdInterval.ValueInt64() != 30 {
		data.ClientDpdInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.bypassUnsupportProtocol"); value.Exists() && !data.ClientBypassProtocol.IsNull() {
		data.ClientBypassProtocol = types.BoolValue(value.Bool())
	} else {
		data.ClientBypassProtocol = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.enableSSLRekey"); value.Exists() && !data.SslRekey.IsNull() {
		data.SslRekey = types.BoolValue(value.Bool())
	} else {
		data.SslRekey = types.BoolNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.rekeyMethod"); value.Exists() && !data.SslRekeyMethod.IsNull() {
		data.SslRekeyMethod = types.StringValue(value.String())
	} else {
		data.SslRekeyMethod = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.rekeyInterval"); value.Exists() && !data.SslRekeyInterval.IsNull() {
		data.SslRekeyInterval = types.Int64Value(value.Int())
	} else {
		data.SslRekeyInterval = types.Int64Null()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientFirewallPrivateNetworkRules.id"); value.Exists() && !data.ClientFirewallPrivateNetworkRulesAclId.IsNull() {
		data.ClientFirewallPrivateNetworkRulesAclId = types.StringValue(value.String())
	} else {
		data.ClientFirewallPrivateNetworkRulesAclId = types.StringNull()
	}
	if value := res.Get("anyConnectSettings.connectionSettings.clientFirewallPublicNetworkRules.id"); value.Exists() && !data.ClientFirewallPublicNetworkRulesAclId.IsNull() {
		data.ClientFirewallPublicNetworkRulesAclId = types.StringValue(value.String())
	} else {
		data.ClientFirewallPublicNetworkRulesAclId = types.StringNull()
	}
	for i := 0; i < len(data.SecureClientCustomAttributes); i++ {
		keys := [...]string{"customAttributeObject.id"}
		keyValues := [...]string{data.SecureClientCustomAttributes[i].Id.ValueString()}

		parent := &data
		data := (*parent).SecureClientCustomAttributes[i]
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
			tflog.Debug(ctx, fmt.Sprintf("removing SecureClientCustomAttributes[%d] = %+v",
				i,
				(*parent).SecureClientCustomAttributes[i],
			))
			(*parent).SecureClientCustomAttributes = slices.Delete((*parent).SecureClientCustomAttributes, i, i+1)
			i--

			continue
		}
		if value := res.Get("customAttributeObject.id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).SecureClientCustomAttributes[i] = data
	}
	if value := res.Get("advancedSettings.vpnTrafficFilterACL.id"); value.Exists() && !data.TrafficFilterAclId.IsNull() {
		data.TrafficFilterAclId = types.StringValue(value.String())
	} else {
		data.TrafficFilterAclId = types.StringNull()
	}
	if value := res.Get("advancedSettings.restrictVPNToVLANId"); value.Exists() && !data.RestrictVpnToVlan.IsNull() {
		data.RestrictVpnToVlan = types.Int64Value(value.Int())
	} else {
		data.RestrictVpnToVlan = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.accessHours.id"); value.Exists() && !data.AccessHoursTimeRangeId.IsNull() {
		data.AccessHoursTimeRangeId = types.StringValue(value.String())
	} else {
		data.AccessHoursTimeRangeId = types.StringNull()
	}
	if value := res.Get("advancedSettings.sessionSettings.simultaneousLoginPerUser"); value.Exists() && !data.SimultaneousLoginsPerUser.IsNull() {
		data.SimultaneousLoginsPerUser = types.Int64Value(value.Int())
	} else {
		data.SimultaneousLoginsPerUser = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.maxConnectionTimeout"); value.Exists() && !data.MaximumConnectionTime.IsNull() {
		data.MaximumConnectionTime = types.Int64Value(value.Int())
	} else {
		data.MaximumConnectionTime = types.Int64Null()
	}
	if value := res.Get("advancedSettings.sessionSettings.maxConnectionTimeAlertInterval"); value.Exists() && !data.MaximumConnectionTimeAlertInterval.IsNull() {
		data.MaximumConnectionTimeAlertInterval = types.Int64Value(value.Int())
	} else {
		data.MaximumConnectionTimeAlertInterval = types.Int64Null()
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
