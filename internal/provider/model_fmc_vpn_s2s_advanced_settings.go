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

type VPNS2SAdvancedSettings struct {
	Id                                                 types.String `tfsdk:"id"`
	Domain                                             types.String `tfsdk:"domain"`
	VpnS2sId                                           types.String `tfsdk:"vpn_s2s_id"`
	Type                                               types.String `tfsdk:"type"`
	IkeKeepalive                                       types.String `tfsdk:"ike_keepalive"`
	IkeKeepaliveThreshold                              types.Int64  `tfsdk:"ike_keepalive_threshold"`
	IkeKeepaliveRetryInterval                          types.Int64  `tfsdk:"ike_keepalive_retry_interval"`
	IkeIdentitySentToPeers                             types.String `tfsdk:"ike_identity_sent_to_peers"`
	IkePeerIdentityValidation                          types.String `tfsdk:"ike_peer_identity_validation"`
	IkeEnableAggressiveMode                            types.Bool   `tfsdk:"ike_enable_aggressive_mode"`
	IkeEnableNotificationOnTunnelDisconnect            types.Bool   `tfsdk:"ike_enable_notification_on_tunnel_disconnect"`
	Ikev2CookieChallenge                               types.String `tfsdk:"ikev2_cookie_challenge"`
	Ikev2ThresholdToChallengeIncomingCookies           types.Int64  `tfsdk:"ikev2_threshold_to_challenge_incoming_cookies"`
	Ikev2PercentageOfSasAllowedInNegotiation           types.Int64  `tfsdk:"ikev2_percentage_of_sas_allowed_in_negotiation"`
	Ikev2MaximumNumberOfSasAllowedInNegotiation        types.Int64  `tfsdk:"ikev2_maximum_number_of_sas_allowed_in_negotiation"`
	IpsecEnableFragmentationBeforeEncryption           types.Bool   `tfsdk:"ipsec_enable_fragmentation_before_encryption"`
	IpsecPathMaximumTransmissionUnitAgingResetInterval types.Int64  `tfsdk:"ipsec_path_maximum_transmission_unit_aging_reset_interval"`
	NatKeepaliveMessageTraversalInterval               types.Int64  `tfsdk:"nat_keepalive_message_traversal_interval"`
	VpnIdleTimeout                                     types.Int64  `tfsdk:"vpn_idle_timeout"`
	BypassAccessControlTrafficForDecryptedTraffic      types.Bool   `tfsdk:"bypass_access_control_traffic_for_decrypted_traffic"`
	UseCertMapConfiguredInEndpointToDetermineTunnel    types.Bool   `tfsdk:"use_cert_map_configured_in_endpoint_to_determine_tunnel"`
	UseCertificateOuToDetermineTunnel                  types.Bool   `tfsdk:"use_certificate_ou_to_determine_tunnel"`
	UseIkeIdentityOuToDetermineTunnel                  types.Bool   `tfsdk:"use_ike_identity_ou_to_determine_tunnel"`
	UsePeerIpAddressToDetermineTunnel                  types.Bool   `tfsdk:"use_peer_ip_address_to_determine_tunnel"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNS2SAdvancedSettings) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftds2svpns/%v/advancedsettings", url.QueryEscape(data.VpnS2sId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNS2SAdvancedSettings) toBody(ctx context.Context, state VPNS2SAdvancedSettings) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.IkeKeepalive.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.ikeKeepaliveSettings.ikeKeepalive", data.IkeKeepalive.ValueString())
	}
	if !data.IkeKeepaliveThreshold.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.ikeKeepaliveSettings.threshold", data.IkeKeepaliveThreshold.ValueInt64())
	}
	if !data.IkeKeepaliveRetryInterval.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.ikeKeepaliveSettings.retryInterval", data.IkeKeepaliveRetryInterval.ValueInt64())
	}
	if !data.IkeIdentitySentToPeers.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.identitySentToPeer", data.IkeIdentitySentToPeers.ValueString())
	}
	if !data.IkePeerIdentityValidation.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.peerIdentityValidation", data.IkePeerIdentityValidation.ValueString())
	}
	if !data.IkeEnableAggressiveMode.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.enableAggressiveMode", data.IkeEnableAggressiveMode.ValueBool())
	}
	if !data.IkeEnableNotificationOnTunnelDisconnect.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.enableNotificationOnTunnelDisconnect", data.IkeEnableNotificationOnTunnelDisconnect.ValueBool())
	}
	if !data.Ikev2CookieChallenge.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.cookieChallenge", data.Ikev2CookieChallenge.ValueString())
	}
	if !data.Ikev2ThresholdToChallengeIncomingCookies.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.thresholdToChallengeIncomingCookies", data.Ikev2ThresholdToChallengeIncomingCookies.ValueInt64())
	}
	if !data.Ikev2PercentageOfSasAllowedInNegotiation.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.percentageOfSAsAllowedInNegotiation", data.Ikev2PercentageOfSasAllowedInNegotiation.ValueInt64())
	}
	if !data.Ikev2MaximumNumberOfSasAllowedInNegotiation.IsNull() {
		body, _ = sjson.Set(body, "advancedIkeSetting.maximumNumberOfSAsAllowed", data.Ikev2MaximumNumberOfSasAllowedInNegotiation.ValueInt64())
	}
	if !data.IpsecEnableFragmentationBeforeEncryption.IsNull() {
		body, _ = sjson.Set(body, "advancedIpsecSetting.enableFragmentationBeforeEncryption", data.IpsecEnableFragmentationBeforeEncryption.ValueBool())
	}
	if !data.IpsecPathMaximumTransmissionUnitAgingResetInterval.IsNull() {
		body, _ = sjson.Set(body, "advancedIpsecSetting.maximumTransmissionUnitAging.resetIntervalMinutes", data.IpsecPathMaximumTransmissionUnitAgingResetInterval.ValueInt64())
	}
	if !data.NatKeepaliveMessageTraversalInterval.IsNull() {
		body, _ = sjson.Set(body, "advancedTunnelSetting.natKeepaliveMessageTraversal.intervalSeconds", data.NatKeepaliveMessageTraversalInterval.ValueInt64())
	}
	if !data.VpnIdleTimeout.IsNull() {
		body, _ = sjson.Set(body, "advancedTunnelSetting.vpnIdleTimeout.timeoutMinutes", data.VpnIdleTimeout.ValueInt64())
	}
	if !data.BypassAccessControlTrafficForDecryptedTraffic.IsNull() {
		body, _ = sjson.Set(body, "advancedTunnelSetting.bypassAccessControlTrafficForDecryptedTraffic", data.BypassAccessControlTrafficForDecryptedTraffic.ValueBool())
	}
	if !data.UseCertMapConfiguredInEndpointToDetermineTunnel.IsNull() {
		body, _ = sjson.Set(body, "advancedTunnelSetting.certificateMapSettings.useCertMapConfiguredInEndpointToDetermineTunnel", data.UseCertMapConfiguredInEndpointToDetermineTunnel.ValueBool())
	}
	if !data.UseCertificateOuToDetermineTunnel.IsNull() {
		body, _ = sjson.Set(body, "advancedTunnelSetting.certificateMapSettings.useCertificateOuToDetermineTunnel", data.UseCertificateOuToDetermineTunnel.ValueBool())
	}
	if !data.UseIkeIdentityOuToDetermineTunnel.IsNull() {
		body, _ = sjson.Set(body, "advancedTunnelSetting.certificateMapSettings.useIkeIdentityOuToDetermineTunnel", data.UseIkeIdentityOuToDetermineTunnel.ValueBool())
	}
	if !data.UsePeerIpAddressToDetermineTunnel.IsNull() {
		body, _ = sjson.Set(body, "advancedTunnelSetting.certificateMapSettings.usePeerIpAddressToDetermineTunnel", data.UsePeerIpAddressToDetermineTunnel.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNS2SAdvancedSettings) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.ikeKeepaliveSettings.ikeKeepalive"); value.Exists() {
		data.IkeKeepalive = types.StringValue(value.String())
	} else {
		data.IkeKeepalive = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.ikeKeepaliveSettings.threshold"); value.Exists() {
		data.IkeKeepaliveThreshold = types.Int64Value(value.Int())
	} else {
		data.IkeKeepaliveThreshold = types.Int64Null()
	}
	if value := res.Get("advancedIkeSetting.ikeKeepaliveSettings.retryInterval"); value.Exists() {
		data.IkeKeepaliveRetryInterval = types.Int64Value(value.Int())
	} else {
		data.IkeKeepaliveRetryInterval = types.Int64Null()
	}
	if value := res.Get("advancedIkeSetting.identitySentToPeer"); value.Exists() {
		data.IkeIdentitySentToPeers = types.StringValue(value.String())
	} else {
		data.IkeIdentitySentToPeers = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.peerIdentityValidation"); value.Exists() {
		data.IkePeerIdentityValidation = types.StringValue(value.String())
	} else {
		data.IkePeerIdentityValidation = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.enableAggressiveMode"); value.Exists() {
		data.IkeEnableAggressiveMode = types.BoolValue(value.Bool())
	} else {
		data.IkeEnableAggressiveMode = types.BoolNull()
	}
	if value := res.Get("advancedIkeSetting.enableNotificationOnTunnelDisconnect"); value.Exists() {
		data.IkeEnableNotificationOnTunnelDisconnect = types.BoolValue(value.Bool())
	} else {
		data.IkeEnableNotificationOnTunnelDisconnect = types.BoolNull()
	}
	if value := res.Get("advancedIkeSetting.cookieChallenge"); value.Exists() {
		data.Ikev2CookieChallenge = types.StringValue(value.String())
	} else {
		data.Ikev2CookieChallenge = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.thresholdToChallengeIncomingCookies"); value.Exists() {
		data.Ikev2ThresholdToChallengeIncomingCookies = types.Int64Value(value.Int())
	} else {
		data.Ikev2ThresholdToChallengeIncomingCookies = types.Int64Null()
	}
	if value := res.Get("advancedIkeSetting.percentageOfSAsAllowedInNegotiation"); value.Exists() {
		data.Ikev2PercentageOfSasAllowedInNegotiation = types.Int64Value(value.Int())
	} else {
		data.Ikev2PercentageOfSasAllowedInNegotiation = types.Int64Null()
	}
	if value := res.Get("advancedIkeSetting.maximumNumberOfSAsAllowed"); value.Exists() {
		data.Ikev2MaximumNumberOfSasAllowedInNegotiation = types.Int64Value(value.Int())
	} else {
		data.Ikev2MaximumNumberOfSasAllowedInNegotiation = types.Int64Null()
	}
	if value := res.Get("advancedIpsecSetting.enableFragmentationBeforeEncryption"); value.Exists() {
		data.IpsecEnableFragmentationBeforeEncryption = types.BoolValue(value.Bool())
	} else {
		data.IpsecEnableFragmentationBeforeEncryption = types.BoolNull()
	}
	if value := res.Get("advancedIpsecSetting.maximumTransmissionUnitAging.resetIntervalMinutes"); value.Exists() {
		data.IpsecPathMaximumTransmissionUnitAgingResetInterval = types.Int64Value(value.Int())
	} else {
		data.IpsecPathMaximumTransmissionUnitAgingResetInterval = types.Int64Null()
	}
	if value := res.Get("advancedTunnelSetting.natKeepaliveMessageTraversal.intervalSeconds"); value.Exists() {
		data.NatKeepaliveMessageTraversalInterval = types.Int64Value(value.Int())
	} else {
		data.NatKeepaliveMessageTraversalInterval = types.Int64Null()
	}
	if value := res.Get("advancedTunnelSetting.vpnIdleTimeout.timeoutMinutes"); value.Exists() {
		data.VpnIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.VpnIdleTimeout = types.Int64Null()
	}
	if value := res.Get("advancedTunnelSetting.bypassAccessControlTrafficForDecryptedTraffic"); value.Exists() {
		data.BypassAccessControlTrafficForDecryptedTraffic = types.BoolValue(value.Bool())
	} else {
		data.BypassAccessControlTrafficForDecryptedTraffic = types.BoolNull()
	}
	if value := res.Get("advancedTunnelSetting.certificateMapSettings.useCertMapConfiguredInEndpointToDetermineTunnel"); value.Exists() {
		data.UseCertMapConfiguredInEndpointToDetermineTunnel = types.BoolValue(value.Bool())
	} else {
		data.UseCertMapConfiguredInEndpointToDetermineTunnel = types.BoolNull()
	}
	if value := res.Get("advancedTunnelSetting.certificateMapSettings.useCertificateOuToDetermineTunnel"); value.Exists() {
		data.UseCertificateOuToDetermineTunnel = types.BoolValue(value.Bool())
	} else {
		data.UseCertificateOuToDetermineTunnel = types.BoolNull()
	}
	if value := res.Get("advancedTunnelSetting.certificateMapSettings.useIkeIdentityOuToDetermineTunnel"); value.Exists() {
		data.UseIkeIdentityOuToDetermineTunnel = types.BoolValue(value.Bool())
	} else {
		data.UseIkeIdentityOuToDetermineTunnel = types.BoolNull()
	}
	if value := res.Get("advancedTunnelSetting.certificateMapSettings.usePeerIpAddressToDetermineTunnel"); value.Exists() {
		data.UsePeerIpAddressToDetermineTunnel = types.BoolValue(value.Bool())
	} else {
		data.UsePeerIpAddressToDetermineTunnel = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNS2SAdvancedSettings) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.ikeKeepaliveSettings.ikeKeepalive"); value.Exists() && !data.IkeKeepalive.IsNull() {
		data.IkeKeepalive = types.StringValue(value.String())
	} else {
		data.IkeKeepalive = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.ikeKeepaliveSettings.threshold"); value.Exists() && !data.IkeKeepaliveThreshold.IsNull() {
		data.IkeKeepaliveThreshold = types.Int64Value(value.Int())
	} else {
		data.IkeKeepaliveThreshold = types.Int64Null()
	}
	if value := res.Get("advancedIkeSetting.ikeKeepaliveSettings.retryInterval"); value.Exists() && !data.IkeKeepaliveRetryInterval.IsNull() {
		data.IkeKeepaliveRetryInterval = types.Int64Value(value.Int())
	} else {
		data.IkeKeepaliveRetryInterval = types.Int64Null()
	}
	if value := res.Get("advancedIkeSetting.identitySentToPeer"); value.Exists() && !data.IkeIdentitySentToPeers.IsNull() {
		data.IkeIdentitySentToPeers = types.StringValue(value.String())
	} else {
		data.IkeIdentitySentToPeers = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.peerIdentityValidation"); value.Exists() && !data.IkePeerIdentityValidation.IsNull() {
		data.IkePeerIdentityValidation = types.StringValue(value.String())
	} else {
		data.IkePeerIdentityValidation = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.enableAggressiveMode"); value.Exists() && !data.IkeEnableAggressiveMode.IsNull() {
		data.IkeEnableAggressiveMode = types.BoolValue(value.Bool())
	} else {
		data.IkeEnableAggressiveMode = types.BoolNull()
	}
	if value := res.Get("advancedIkeSetting.enableNotificationOnTunnelDisconnect"); value.Exists() && !data.IkeEnableNotificationOnTunnelDisconnect.IsNull() {
		data.IkeEnableNotificationOnTunnelDisconnect = types.BoolValue(value.Bool())
	} else {
		data.IkeEnableNotificationOnTunnelDisconnect = types.BoolNull()
	}
	if value := res.Get("advancedIkeSetting.cookieChallenge"); value.Exists() && !data.Ikev2CookieChallenge.IsNull() {
		data.Ikev2CookieChallenge = types.StringValue(value.String())
	} else {
		data.Ikev2CookieChallenge = types.StringNull()
	}
	if value := res.Get("advancedIkeSetting.thresholdToChallengeIncomingCookies"); value.Exists() && !data.Ikev2ThresholdToChallengeIncomingCookies.IsNull() {
		data.Ikev2ThresholdToChallengeIncomingCookies = types.Int64Value(value.Int())
	} else {
		data.Ikev2ThresholdToChallengeIncomingCookies = types.Int64Null()
	}
	if value := res.Get("advancedIkeSetting.percentageOfSAsAllowedInNegotiation"); value.Exists() && !data.Ikev2PercentageOfSasAllowedInNegotiation.IsNull() {
		data.Ikev2PercentageOfSasAllowedInNegotiation = types.Int64Value(value.Int())
	} else {
		data.Ikev2PercentageOfSasAllowedInNegotiation = types.Int64Null()
	}
	if value := res.Get("advancedIkeSetting.maximumNumberOfSAsAllowed"); value.Exists() && !data.Ikev2MaximumNumberOfSasAllowedInNegotiation.IsNull() {
		data.Ikev2MaximumNumberOfSasAllowedInNegotiation = types.Int64Value(value.Int())
	} else {
		data.Ikev2MaximumNumberOfSasAllowedInNegotiation = types.Int64Null()
	}
	if value := res.Get("advancedIpsecSetting.enableFragmentationBeforeEncryption"); value.Exists() && !data.IpsecEnableFragmentationBeforeEncryption.IsNull() {
		data.IpsecEnableFragmentationBeforeEncryption = types.BoolValue(value.Bool())
	} else {
		data.IpsecEnableFragmentationBeforeEncryption = types.BoolNull()
	}
	if value := res.Get("advancedIpsecSetting.maximumTransmissionUnitAging.resetIntervalMinutes"); value.Exists() && !data.IpsecPathMaximumTransmissionUnitAgingResetInterval.IsNull() {
		data.IpsecPathMaximumTransmissionUnitAgingResetInterval = types.Int64Value(value.Int())
	} else {
		data.IpsecPathMaximumTransmissionUnitAgingResetInterval = types.Int64Null()
	}
	if value := res.Get("advancedTunnelSetting.natKeepaliveMessageTraversal.intervalSeconds"); value.Exists() && !data.NatKeepaliveMessageTraversalInterval.IsNull() {
		data.NatKeepaliveMessageTraversalInterval = types.Int64Value(value.Int())
	} else {
		data.NatKeepaliveMessageTraversalInterval = types.Int64Null()
	}
	if value := res.Get("advancedTunnelSetting.vpnIdleTimeout.timeoutMinutes"); value.Exists() && !data.VpnIdleTimeout.IsNull() {
		data.VpnIdleTimeout = types.Int64Value(value.Int())
	} else {
		data.VpnIdleTimeout = types.Int64Null()
	}
	if value := res.Get("advancedTunnelSetting.bypassAccessControlTrafficForDecryptedTraffic"); value.Exists() && !data.BypassAccessControlTrafficForDecryptedTraffic.IsNull() {
		data.BypassAccessControlTrafficForDecryptedTraffic = types.BoolValue(value.Bool())
	} else {
		data.BypassAccessControlTrafficForDecryptedTraffic = types.BoolNull()
	}
	if value := res.Get("advancedTunnelSetting.certificateMapSettings.useCertMapConfiguredInEndpointToDetermineTunnel"); value.Exists() && !data.UseCertMapConfiguredInEndpointToDetermineTunnel.IsNull() {
		data.UseCertMapConfiguredInEndpointToDetermineTunnel = types.BoolValue(value.Bool())
	} else {
		data.UseCertMapConfiguredInEndpointToDetermineTunnel = types.BoolNull()
	}
	if value := res.Get("advancedTunnelSetting.certificateMapSettings.useCertificateOuToDetermineTunnel"); value.Exists() && !data.UseCertificateOuToDetermineTunnel.IsNull() {
		data.UseCertificateOuToDetermineTunnel = types.BoolValue(value.Bool())
	} else {
		data.UseCertificateOuToDetermineTunnel = types.BoolNull()
	}
	if value := res.Get("advancedTunnelSetting.certificateMapSettings.useIkeIdentityOuToDetermineTunnel"); value.Exists() && !data.UseIkeIdentityOuToDetermineTunnel.IsNull() {
		data.UseIkeIdentityOuToDetermineTunnel = types.BoolValue(value.Bool())
	} else {
		data.UseIkeIdentityOuToDetermineTunnel = types.BoolNull()
	}
	if value := res.Get("advancedTunnelSetting.certificateMapSettings.usePeerIpAddressToDetermineTunnel"); value.Exists() && !data.UsePeerIpAddressToDetermineTunnel.IsNull() {
		data.UsePeerIpAddressToDetermineTunnel = types.BoolValue(value.Bool())
	} else {
		data.UsePeerIpAddressToDetermineTunnel = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNS2SAdvancedSettings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

func (data VPNS2SAdvancedSettings) fixFields(ctx context.Context, req string) string {
	if data.IpsecPathMaximumTransmissionUnitAgingResetInterval.IsNull() {
		req, _ = sjson.Set(req, "advancedIpsecSetting.maximumTransmissionUnitAging.enabled", false)
	} else {
		req, _ = sjson.Set(req, "advancedIpsecSetting.maximumTransmissionUnitAging.enabled", true)
	}

	if data.NatKeepaliveMessageTraversalInterval.IsNull() {
		req, _ = sjson.Set(req, "advancedTunnelSetting.natKeepaliveMessageTraversal.enabled", false)
	} else {
		req, _ = sjson.Set(req, "advancedTunnelSetting.natKeepaliveMessageTraversal.enabled", true)
	}

	if data.VpnIdleTimeout.IsNull() {
		req, _ = sjson.Set(req, "advancedTunnelSetting.vpnIdleTimeout.enabled", false)
	} else {
		req, _ = sjson.Set(req, "advancedTunnelSetting.vpnIdleTimeout.enabled", true)
	}

	return req
}
