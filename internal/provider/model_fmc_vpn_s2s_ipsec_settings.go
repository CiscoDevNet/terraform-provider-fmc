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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type VPNS2SIPSECSettings struct {
	Id                                     types.String                             `tfsdk:"id"`
	Domain                                 types.String                             `tfsdk:"domain"`
	VpnS2sId                               types.String                             `tfsdk:"vpn_s2s_id"`
	Type                                   types.String                             `tfsdk:"type"`
	CryptoMapType                          types.String                             `tfsdk:"crypto_map_type"`
	Ikev2Mode                              types.String                             `tfsdk:"ikev2_mode"`
	Ikev1IpsecProposals                    []VPNS2SIPSECSettingsIkev1IpsecProposals `tfsdk:"ikev1_ipsec_proposals"`
	Ikev2IpsecProposals                    []VPNS2SIPSECSettingsIkev2IpsecProposals `tfsdk:"ikev2_ipsec_proposals"`
	SecurityAssociationStrengthEnforcement types.Bool                               `tfsdk:"security_association_strength_enforcement"`
	ReverseRouteInjection                  types.Bool                               `tfsdk:"reverse_route_injection"`
	PerfectForwardSecrecy                  types.Bool                               `tfsdk:"perfect_forward_secrecy"`
	PerfectForwardSecrecyModulusGroup      types.String                             `tfsdk:"perfect_forward_secrecy_modulus_group"`
	LifetimeDuration                       types.Int64                              `tfsdk:"lifetime_duration"`
	LifetimeSize                           types.Int64                              `tfsdk:"lifetime_size"`
	ValidateIncomingIcmpErrorMessages      types.Bool                               `tfsdk:"validate_incoming_icmp_error_messages"`
	DoNotFragmentPolicy                    types.String                             `tfsdk:"do_not_fragment_policy"`
	Tfc                                    types.Bool                               `tfsdk:"tfc"`
	TfcBurstBytes                          types.Int64                              `tfsdk:"tfc_burst_bytes"`
	TfcPayloadBytes                        types.Int64                              `tfsdk:"tfc_payload_bytes"`
	TfcTimeout                             types.Int64                              `tfsdk:"tfc_timeout"`
}

type VPNS2SIPSECSettingsIkev1IpsecProposals struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type VPNS2SIPSECSettingsIkev2IpsecProposals struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNS2SIPSECSettings) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftds2svpns/%v/ipsecsettings", url.QueryEscape(data.VpnS2sId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNS2SIPSECSettings) toBody(ctx context.Context, state VPNS2SIPSECSettings) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.CryptoMapType.IsNull() {
		body, _ = sjson.Set(body, "cryptoMapType", data.CryptoMapType.ValueString())
	}
	if !data.Ikev2Mode.IsNull() {
		body, _ = sjson.Set(body, "ikeV2Mode", data.Ikev2Mode.ValueString())
	}
	if len(data.Ikev1IpsecProposals) > 0 {
		body, _ = sjson.Set(body, "ikeV1IpsecProposal", []interface{}{})
		for _, item := range data.Ikev1IpsecProposals {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ikeV1IpsecProposal.-1", itemBody)
		}
	}
	if len(data.Ikev2IpsecProposals) > 0 {
		body, _ = sjson.Set(body, "ikeV2IpsecProposal", []interface{}{})
		for _, item := range data.Ikev2IpsecProposals {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ikeV2IpsecProposal.-1", itemBody)
		}
	}
	if !data.SecurityAssociationStrengthEnforcement.IsNull() {
		body, _ = sjson.Set(body, "enableSaStrengthEnforcement", data.SecurityAssociationStrengthEnforcement.ValueBool())
	}
	if !data.ReverseRouteInjection.IsNull() {
		body, _ = sjson.Set(body, "enableRRI", data.ReverseRouteInjection.ValueBool())
	}
	if !data.PerfectForwardSecrecy.IsNull() {
		body, _ = sjson.Set(body, "perfectForwardSecrecy.enabled", data.PerfectForwardSecrecy.ValueBool())
	}
	if !data.PerfectForwardSecrecyModulusGroup.IsNull() {
		body, _ = sjson.Set(body, "perfectForwardSecrecy.modulusGroup", data.PerfectForwardSecrecyModulusGroup.ValueString())
	}
	if !data.LifetimeDuration.IsNull() {
		body, _ = sjson.Set(body, "lifetimeSeconds", data.LifetimeDuration.ValueInt64())
	}
	if !data.LifetimeSize.IsNull() {
		body, _ = sjson.Set(body, "lifetimeKilobytes", data.LifetimeSize.ValueInt64())
	}
	if !data.ValidateIncomingIcmpErrorMessages.IsNull() {
		body, _ = sjson.Set(body, "validateIncomingIcmpErrorMessage", data.ValidateIncomingIcmpErrorMessages.ValueBool())
	}
	if !data.DoNotFragmentPolicy.IsNull() {
		body, _ = sjson.Set(body, "doNotFragmentPolicy", data.DoNotFragmentPolicy.ValueString())
	}
	if !data.Tfc.IsNull() {
		body, _ = sjson.Set(body, "tfcPackets.enabled", data.Tfc.ValueBool())
	}
	if !data.TfcBurstBytes.IsNull() {
		body, _ = sjson.Set(body, "tfcPackets.burstBytes", data.TfcBurstBytes.ValueInt64())
	}
	if !data.TfcPayloadBytes.IsNull() {
		body, _ = sjson.Set(body, "tfcPackets.payloadBytes", data.TfcPayloadBytes.ValueInt64())
	}
	if !data.TfcTimeout.IsNull() {
		body, _ = sjson.Set(body, "tfcPackets.timeoutSeconds", data.TfcTimeout.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNS2SIPSECSettings) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("cryptoMapType"); value.Exists() {
		data.CryptoMapType = types.StringValue(value.String())
	} else {
		data.CryptoMapType = types.StringNull()
	}
	if value := res.Get("ikeV2Mode"); value.Exists() {
		data.Ikev2Mode = types.StringValue(value.String())
	} else {
		data.Ikev2Mode = types.StringNull()
	}
	if value := res.Get("ikeV1IpsecProposal"); value.Exists() {
		data.Ikev1IpsecProposals = make([]VPNS2SIPSECSettingsIkev1IpsecProposals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNS2SIPSECSettingsIkev1IpsecProposals{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).Ikev1IpsecProposals = append((*parent).Ikev1IpsecProposals, data)
			return true
		})
	}
	if value := res.Get("ikeV2IpsecProposal"); value.Exists() {
		data.Ikev2IpsecProposals = make([]VPNS2SIPSECSettingsIkev2IpsecProposals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNS2SIPSECSettingsIkev2IpsecProposals{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).Ikev2IpsecProposals = append((*parent).Ikev2IpsecProposals, data)
			return true
		})
	}
	if value := res.Get("enableSaStrengthEnforcement"); value.Exists() {
		data.SecurityAssociationStrengthEnforcement = types.BoolValue(value.Bool())
	} else {
		data.SecurityAssociationStrengthEnforcement = types.BoolNull()
	}
	if value := res.Get("enableRRI"); value.Exists() {
		data.ReverseRouteInjection = types.BoolValue(value.Bool())
	} else {
		data.ReverseRouteInjection = types.BoolNull()
	}
	if value := res.Get("perfectForwardSecrecy.enabled"); value.Exists() {
		data.PerfectForwardSecrecy = types.BoolValue(value.Bool())
	} else {
		data.PerfectForwardSecrecy = types.BoolNull()
	}
	if value := res.Get("perfectForwardSecrecy.modulusGroup"); value.Exists() {
		data.PerfectForwardSecrecyModulusGroup = types.StringValue(value.String())
	} else {
		data.PerfectForwardSecrecyModulusGroup = types.StringNull()
	}
	if value := res.Get("lifetimeSeconds"); value.Exists() {
		data.LifetimeDuration = types.Int64Value(value.Int())
	} else {
		data.LifetimeDuration = types.Int64Null()
	}
	if value := res.Get("lifetimeKilobytes"); value.Exists() {
		data.LifetimeSize = types.Int64Value(value.Int())
	} else {
		data.LifetimeSize = types.Int64Null()
	}
	if value := res.Get("validateIncomingIcmpErrorMessage"); value.Exists() {
		data.ValidateIncomingIcmpErrorMessages = types.BoolValue(value.Bool())
	} else {
		data.ValidateIncomingIcmpErrorMessages = types.BoolNull()
	}
	if value := res.Get("doNotFragmentPolicy"); value.Exists() {
		data.DoNotFragmentPolicy = types.StringValue(value.String())
	} else {
		data.DoNotFragmentPolicy = types.StringNull()
	}
	if value := res.Get("tfcPackets.enabled"); value.Exists() {
		data.Tfc = types.BoolValue(value.Bool())
	} else {
		data.Tfc = types.BoolNull()
	}
	if value := res.Get("tfcPackets.burstBytes"); value.Exists() {
		data.TfcBurstBytes = types.Int64Value(value.Int())
	} else {
		data.TfcBurstBytes = types.Int64Null()
	}
	if value := res.Get("tfcPackets.payloadBytes"); value.Exists() {
		data.TfcPayloadBytes = types.Int64Value(value.Int())
	} else {
		data.TfcPayloadBytes = types.Int64Null()
	}
	if value := res.Get("tfcPackets.timeoutSeconds"); value.Exists() {
		data.TfcTimeout = types.Int64Value(value.Int())
	} else {
		data.TfcTimeout = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNS2SIPSECSettings) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("cryptoMapType"); value.Exists() && !data.CryptoMapType.IsNull() {
		data.CryptoMapType = types.StringValue(value.String())
	} else {
		data.CryptoMapType = types.StringNull()
	}
	if value := res.Get("ikeV2Mode"); value.Exists() && !data.Ikev2Mode.IsNull() {
		data.Ikev2Mode = types.StringValue(value.String())
	} else {
		data.Ikev2Mode = types.StringNull()
	}
	for i := 0; i < len(data.Ikev1IpsecProposals); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Ikev1IpsecProposals[i].Id.ValueString()}

		parent := &data
		data := (*parent).Ikev1IpsecProposals[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ikeV1IpsecProposal").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ikev1IpsecProposals[%d] = %+v",
				i,
				(*parent).Ikev1IpsecProposals[i],
			))
			(*parent).Ikev1IpsecProposals = slices.Delete((*parent).Ikev1IpsecProposals, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).Ikev1IpsecProposals[i] = data
	}
	for i := 0; i < len(data.Ikev2IpsecProposals); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Ikev2IpsecProposals[i].Id.ValueString()}

		parent := &data
		data := (*parent).Ikev2IpsecProposals[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ikeV2IpsecProposal").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ikev2IpsecProposals[%d] = %+v",
				i,
				(*parent).Ikev2IpsecProposals[i],
			))
			(*parent).Ikev2IpsecProposals = slices.Delete((*parent).Ikev2IpsecProposals, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		(*parent).Ikev2IpsecProposals[i] = data
	}
	if value := res.Get("enableSaStrengthEnforcement"); value.Exists() && !data.SecurityAssociationStrengthEnforcement.IsNull() {
		data.SecurityAssociationStrengthEnforcement = types.BoolValue(value.Bool())
	} else {
		data.SecurityAssociationStrengthEnforcement = types.BoolNull()
	}
	if value := res.Get("enableRRI"); value.Exists() && !data.ReverseRouteInjection.IsNull() {
		data.ReverseRouteInjection = types.BoolValue(value.Bool())
	} else {
		data.ReverseRouteInjection = types.BoolNull()
	}
	if value := res.Get("perfectForwardSecrecy.enabled"); value.Exists() && !data.PerfectForwardSecrecy.IsNull() {
		data.PerfectForwardSecrecy = types.BoolValue(value.Bool())
	} else {
		data.PerfectForwardSecrecy = types.BoolNull()
	}
	if value := res.Get("perfectForwardSecrecy.modulusGroup"); value.Exists() && !data.PerfectForwardSecrecyModulusGroup.IsNull() {
		data.PerfectForwardSecrecyModulusGroup = types.StringValue(value.String())
	} else {
		data.PerfectForwardSecrecyModulusGroup = types.StringNull()
	}
	if value := res.Get("lifetimeSeconds"); value.Exists() && !data.LifetimeDuration.IsNull() {
		data.LifetimeDuration = types.Int64Value(value.Int())
	} else {
		data.LifetimeDuration = types.Int64Null()
	}
	if value := res.Get("lifetimeKilobytes"); value.Exists() && !data.LifetimeSize.IsNull() {
		data.LifetimeSize = types.Int64Value(value.Int())
	} else {
		data.LifetimeSize = types.Int64Null()
	}
	if value := res.Get("validateIncomingIcmpErrorMessage"); value.Exists() && !data.ValidateIncomingIcmpErrorMessages.IsNull() {
		data.ValidateIncomingIcmpErrorMessages = types.BoolValue(value.Bool())
	} else {
		data.ValidateIncomingIcmpErrorMessages = types.BoolNull()
	}
	if value := res.Get("doNotFragmentPolicy"); value.Exists() && !data.DoNotFragmentPolicy.IsNull() {
		data.DoNotFragmentPolicy = types.StringValue(value.String())
	} else {
		data.DoNotFragmentPolicy = types.StringNull()
	}
	if value := res.Get("tfcPackets.enabled"); value.Exists() && !data.Tfc.IsNull() {
		data.Tfc = types.BoolValue(value.Bool())
	} else {
		data.Tfc = types.BoolNull()
	}
	if value := res.Get("tfcPackets.burstBytes"); value.Exists() && !data.TfcBurstBytes.IsNull() {
		data.TfcBurstBytes = types.Int64Value(value.Int())
	} else {
		data.TfcBurstBytes = types.Int64Null()
	}
	if value := res.Get("tfcPackets.payloadBytes"); value.Exists() && !data.TfcPayloadBytes.IsNull() {
		data.TfcPayloadBytes = types.Int64Value(value.Int())
	} else {
		data.TfcPayloadBytes = types.Int64Null()
	}
	if value := res.Get("tfcPackets.timeoutSeconds"); value.Exists() && !data.TfcTimeout.IsNull() {
		data.TfcTimeout = types.Int64Value(value.Int())
	} else {
		data.TfcTimeout = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNS2SIPSECSettings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
func (data VPNS2SIPSECSettings) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	return body
}

// End of section. //template:end toBodyPutDelete
