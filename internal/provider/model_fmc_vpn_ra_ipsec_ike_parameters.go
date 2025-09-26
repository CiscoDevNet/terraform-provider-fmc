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

type VPNRAIPSecIKEParameters struct {
	Id                                                 types.String `tfsdk:"id"`
	Domain                                             types.String `tfsdk:"domain"`
	VpnRaId                                            types.String `tfsdk:"vpn_ra_id"`
	Type                                               types.String `tfsdk:"type"`
	Ikev2IdentitySentToPeer                            types.String `tfsdk:"ikev2_identity_sent_to_peer"`
	Ikev2NotificationOnTunnelDisconnect                types.Bool   `tfsdk:"ikev2_notification_on_tunnel_disconnect"`
	Ikev2DoNotRebootUntilAllSessionsAreTerminated      types.Bool   `tfsdk:"ikev2_do_not_reboot_until_all_sessions_are_terminated"`
	Ikev2CookieChallenge                               types.String `tfsdk:"ikev2_cookie_challenge"`
	Ikev2ThresholdToChallengeIncomingCookies           types.Int64  `tfsdk:"ikev2_threshold_to_challenge_incoming_cookies"`
	Ikev2NumberOfSasAllowedInNegotiation               types.Int64  `tfsdk:"ikev2_number_of_sas_allowed_in_negotiation"`
	Ikev2MaximumNumberOfSasAllowed                     types.Int64  `tfsdk:"ikev2_maximum_number_of_sas_allowed"`
	IpsecPathMaximumTransmissionUnitAgingResetInterval types.Int64  `tfsdk:"ipsec_path_maximum_transmission_unit_aging_reset_interval"`
	NatKeepaliveMessageTraversal                       types.Bool   `tfsdk:"nat_keepalive_message_traversal"`
	NatKeepaliveMessageTraversalInterval               types.Int64  `tfsdk:"nat_keepalive_message_traversal_interval"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNRAIPSecIKEParameters) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ravpns/%v/ipsecadvancedsettings", url.QueryEscape(data.VpnRaId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNRAIPSecIKEParameters) toBody(ctx context.Context, state VPNRAIPSecIKEParameters) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "RaVpnIPsecAdvancedSetting")
	if !data.Ikev2IdentitySentToPeer.IsNull() {
		body, _ = sjson.Set(body, "ikev2settings.identitySentToPeer", data.Ikev2IdentitySentToPeer.ValueString())
	}
	if !data.Ikev2NotificationOnTunnelDisconnect.IsNull() {
		body, _ = sjson.Set(body, "ikev2settings.enableNotificationOnTunnelDisconnect", data.Ikev2NotificationOnTunnelDisconnect.ValueBool())
	}
	if !data.Ikev2DoNotRebootUntilAllSessionsAreTerminated.IsNull() {
		body, _ = sjson.Set(body, "ikev2settings.doNotRebootUntilSessionsTerminated", data.Ikev2DoNotRebootUntilAllSessionsAreTerminated.ValueBool())
	}
	if !data.Ikev2CookieChallenge.IsNull() {
		body, _ = sjson.Set(body, "ikev2settings.cookieChallenge", data.Ikev2CookieChallenge.ValueString())
	}
	if !data.Ikev2ThresholdToChallengeIncomingCookies.IsNull() {
		body, _ = sjson.Set(body, "ikev2settings.thresholdToChallengeIncomingCookies", data.Ikev2ThresholdToChallengeIncomingCookies.ValueInt64())
	}
	if !data.Ikev2NumberOfSasAllowedInNegotiation.IsNull() {
		body, _ = sjson.Set(body, "ikev2settings.percentageOfSAsAllowedInNegotiation", data.Ikev2NumberOfSasAllowedInNegotiation.ValueInt64())
	}
	if !data.Ikev2MaximumNumberOfSasAllowed.IsNull() {
		body, _ = sjson.Set(body, "ikev2settings.maximumNumberOfSAsAllowed", data.Ikev2MaximumNumberOfSasAllowed.ValueInt64())
	}
	if !data.IpsecPathMaximumTransmissionUnitAgingResetInterval.IsNull() {
		body, _ = sjson.Set(body, "ipsecsettings.maximumTransmissionUnitAging.resetIntervalMinutes", data.IpsecPathMaximumTransmissionUnitAgingResetInterval.ValueInt64())
	}
	if !data.NatKeepaliveMessageTraversal.IsNull() {
		body, _ = sjson.Set(body, "natKeepaliveMessageTraversal.enabled", data.NatKeepaliveMessageTraversal.ValueBool())
	}
	if !data.NatKeepaliveMessageTraversalInterval.IsNull() {
		body, _ = sjson.Set(body, "natKeepaliveMessageTraversal.intervalSeconds", data.NatKeepaliveMessageTraversalInterval.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNRAIPSecIKEParameters) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("ikev2settings.identitySentToPeer"); value.Exists() {
		data.Ikev2IdentitySentToPeer = types.StringValue(value.String())
	} else {
		data.Ikev2IdentitySentToPeer = types.StringNull()
	}
	if value := res.Get("ikev2settings.enableNotificationOnTunnelDisconnect"); value.Exists() {
		data.Ikev2NotificationOnTunnelDisconnect = types.BoolValue(value.Bool())
	} else {
		data.Ikev2NotificationOnTunnelDisconnect = types.BoolNull()
	}
	if value := res.Get("ikev2settings.doNotRebootUntilSessionsTerminated"); value.Exists() {
		data.Ikev2DoNotRebootUntilAllSessionsAreTerminated = types.BoolValue(value.Bool())
	} else {
		data.Ikev2DoNotRebootUntilAllSessionsAreTerminated = types.BoolNull()
	}
	if value := res.Get("ikev2settings.cookieChallenge"); value.Exists() {
		data.Ikev2CookieChallenge = types.StringValue(value.String())
	} else {
		data.Ikev2CookieChallenge = types.StringNull()
	}
	if value := res.Get("ikev2settings.thresholdToChallengeIncomingCookies"); value.Exists() {
		data.Ikev2ThresholdToChallengeIncomingCookies = types.Int64Value(value.Int())
	} else {
		data.Ikev2ThresholdToChallengeIncomingCookies = types.Int64Null()
	}
	if value := res.Get("ikev2settings.percentageOfSAsAllowedInNegotiation"); value.Exists() {
		data.Ikev2NumberOfSasAllowedInNegotiation = types.Int64Value(value.Int())
	} else {
		data.Ikev2NumberOfSasAllowedInNegotiation = types.Int64Null()
	}
	if value := res.Get("ikev2settings.maximumNumberOfSAsAllowed"); value.Exists() {
		data.Ikev2MaximumNumberOfSasAllowed = types.Int64Value(value.Int())
	} else {
		data.Ikev2MaximumNumberOfSasAllowed = types.Int64Null()
	}
	if value := res.Get("ipsecsettings.maximumTransmissionUnitAging.resetIntervalMinutes"); value.Exists() {
		data.IpsecPathMaximumTransmissionUnitAgingResetInterval = types.Int64Value(value.Int())
	} else {
		data.IpsecPathMaximumTransmissionUnitAgingResetInterval = types.Int64Null()
	}
	if value := res.Get("natKeepaliveMessageTraversal.enabled"); value.Exists() {
		data.NatKeepaliveMessageTraversal = types.BoolValue(value.Bool())
	} else {
		data.NatKeepaliveMessageTraversal = types.BoolNull()
	}
	if value := res.Get("natKeepaliveMessageTraversal.intervalSeconds"); value.Exists() {
		data.NatKeepaliveMessageTraversalInterval = types.Int64Value(value.Int())
	} else {
		data.NatKeepaliveMessageTraversalInterval = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNRAIPSecIKEParameters) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("ikev2settings.identitySentToPeer"); value.Exists() && !data.Ikev2IdentitySentToPeer.IsNull() {
		data.Ikev2IdentitySentToPeer = types.StringValue(value.String())
	} else {
		data.Ikev2IdentitySentToPeer = types.StringNull()
	}
	if value := res.Get("ikev2settings.enableNotificationOnTunnelDisconnect"); value.Exists() && !data.Ikev2NotificationOnTunnelDisconnect.IsNull() {
		data.Ikev2NotificationOnTunnelDisconnect = types.BoolValue(value.Bool())
	} else {
		data.Ikev2NotificationOnTunnelDisconnect = types.BoolNull()
	}
	if value := res.Get("ikev2settings.doNotRebootUntilSessionsTerminated"); value.Exists() && !data.Ikev2DoNotRebootUntilAllSessionsAreTerminated.IsNull() {
		data.Ikev2DoNotRebootUntilAllSessionsAreTerminated = types.BoolValue(value.Bool())
	} else {
		data.Ikev2DoNotRebootUntilAllSessionsAreTerminated = types.BoolNull()
	}
	if value := res.Get("ikev2settings.cookieChallenge"); value.Exists() && !data.Ikev2CookieChallenge.IsNull() {
		data.Ikev2CookieChallenge = types.StringValue(value.String())
	} else {
		data.Ikev2CookieChallenge = types.StringNull()
	}
	if value := res.Get("ikev2settings.thresholdToChallengeIncomingCookies"); value.Exists() && !data.Ikev2ThresholdToChallengeIncomingCookies.IsNull() {
		data.Ikev2ThresholdToChallengeIncomingCookies = types.Int64Value(value.Int())
	} else {
		data.Ikev2ThresholdToChallengeIncomingCookies = types.Int64Null()
	}
	if value := res.Get("ikev2settings.percentageOfSAsAllowedInNegotiation"); value.Exists() && !data.Ikev2NumberOfSasAllowedInNegotiation.IsNull() {
		data.Ikev2NumberOfSasAllowedInNegotiation = types.Int64Value(value.Int())
	} else {
		data.Ikev2NumberOfSasAllowedInNegotiation = types.Int64Null()
	}
	if value := res.Get("ikev2settings.maximumNumberOfSAsAllowed"); value.Exists() && !data.Ikev2MaximumNumberOfSasAllowed.IsNull() {
		data.Ikev2MaximumNumberOfSasAllowed = types.Int64Value(value.Int())
	} else {
		data.Ikev2MaximumNumberOfSasAllowed = types.Int64Null()
	}
	if value := res.Get("ipsecsettings.maximumTransmissionUnitAging.resetIntervalMinutes"); value.Exists() && !data.IpsecPathMaximumTransmissionUnitAgingResetInterval.IsNull() {
		data.IpsecPathMaximumTransmissionUnitAgingResetInterval = types.Int64Value(value.Int())
	} else {
		data.IpsecPathMaximumTransmissionUnitAgingResetInterval = types.Int64Null()
	}
	if value := res.Get("natKeepaliveMessageTraversal.enabled"); value.Exists() && !data.NatKeepaliveMessageTraversal.IsNull() {
		data.NatKeepaliveMessageTraversal = types.BoolValue(value.Bool())
	} else {
		data.NatKeepaliveMessageTraversal = types.BoolNull()
	}
	if value := res.Get("natKeepaliveMessageTraversal.intervalSeconds"); value.Exists() && !data.NatKeepaliveMessageTraversalInterval.IsNull() {
		data.NatKeepaliveMessageTraversalInterval = types.Int64Value(value.Int())
	} else {
		data.NatKeepaliveMessageTraversalInterval = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNRAIPSecIKEParameters) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

func (data VPNRAIPSecIKEParameters) adjustBody(ctx context.Context, req string) string {
	if !data.IpsecPathMaximumTransmissionUnitAgingResetInterval.IsNull() {
		req, _ = sjson.Set(req, "ipsecsettings.maximumTransmissionUnitAging.enabled", true)
	}
	return req
}
