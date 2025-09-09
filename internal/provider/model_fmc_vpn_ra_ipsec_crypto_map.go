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

type VPNRAIPSecCryptoMap struct {
	Id                                types.String                             `tfsdk:"id"`
	Domain                            types.String                             `tfsdk:"domain"`
	VpnRaId                           types.String                             `tfsdk:"vpn_ra_id"`
	Type                              types.String                             `tfsdk:"type"`
	InterfaceId                       types.String                             `tfsdk:"interface_id"`
	Ikev2IpsecProposals               []VPNRAIPSecCryptoMapIkev2IpsecProposals `tfsdk:"ikev2_ipsec_proposals"`
	ReverseRouteInjection             types.Bool                               `tfsdk:"reverse_route_injection"`
	ClientServices                    types.Bool                               `tfsdk:"client_services"`
	ClientServicesPort                types.Int64                              `tfsdk:"client_services_port"`
	PerfectForwardSecrecy             types.Bool                               `tfsdk:"perfect_forward_secrecy"`
	PerfectForwardSecrecyModulusGroup types.String                             `tfsdk:"perfect_forward_secrecy_modulus_group"`
	LifetimeDuration                  types.Int64                              `tfsdk:"lifetime_duration"`
	LifetimeSize                      types.Int64                              `tfsdk:"lifetime_size"`
	ValidateIncomingIcmpErrorMessages types.Bool                               `tfsdk:"validate_incoming_icmp_error_messages"`
	DoNotFragmentPolicy               types.String                             `tfsdk:"do_not_fragment_policy"`
	Tfc                               types.Bool                               `tfsdk:"tfc"`
	TfcBurstBytes                     types.Int64                              `tfsdk:"tfc_burst_bytes"`
	TfcPayloadBytes                   types.Int64                              `tfsdk:"tfc_payload_bytes"`
	TfcTimeout                        types.Int64                              `tfsdk:"tfc_timeout"`
}

type VPNRAIPSecCryptoMapIkev2IpsecProposals struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNRAIPSecCryptoMap) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ravpns/%v/ipseccryptomaps", url.QueryEscape(data.VpnRaId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNRAIPSecCryptoMap) toBody(ctx context.Context, state VPNRAIPSecCryptoMap) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "RaVpnIPsecCryptoMap")
	if !data.InterfaceId.IsNull() {
		body, _ = sjson.Set(body, "interfaceObject.id", data.InterfaceId.ValueString())
	}
	if len(data.Ikev2IpsecProposals) > 0 {
		body, _ = sjson.Set(body, "ikev2IpsecProposals", []any{})
		for _, item := range data.Ikev2IpsecProposals {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ikev2IpsecProposals.-1", itemBody)
		}
	}
	if !data.ReverseRouteInjection.IsNull() {
		body, _ = sjson.Set(body, "enableRRI", data.ReverseRouteInjection.ValueBool())
	}
	if !data.ClientServices.IsNull() {
		body, _ = sjson.Set(body, "enableClientServices", data.ClientServices.ValueBool())
	}
	if !data.ClientServicesPort.IsNull() {
		body, _ = sjson.Set(body, "clientServicesPort", data.ClientServicesPort.ValueInt64())
	}
	if !data.PerfectForwardSecrecy.IsNull() {
		body, _ = sjson.Set(body, "perfectForwardSecracy.enabled", data.PerfectForwardSecrecy.ValueBool())
	}
	if !data.PerfectForwardSecrecyModulusGroup.IsNull() {
		body, _ = sjson.Set(body, "perfectForwardSecracy.modulusGroup", data.PerfectForwardSecrecyModulusGroup.ValueString())
	}
	if !data.LifetimeDuration.IsNull() {
		body, _ = sjson.Set(body, "lifeTimeSeconds", data.LifetimeDuration.ValueInt64())
	}
	if !data.LifetimeSize.IsNull() {
		body, _ = sjson.Set(body, "lifeTimeKilobytes", data.LifetimeSize.ValueInt64())
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

func (data *VPNRAIPSecCryptoMap) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("interfaceObject.id"); value.Exists() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	if value := res.Get("ikev2IpsecProposals"); value.Exists() {
		data.Ikev2IpsecProposals = make([]VPNRAIPSecCryptoMapIkev2IpsecProposals, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNRAIPSecCryptoMapIkev2IpsecProposals{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).Ikev2IpsecProposals = append((*parent).Ikev2IpsecProposals, data)
			return true
		})
	}
	if value := res.Get("enableRRI"); value.Exists() {
		data.ReverseRouteInjection = types.BoolValue(value.Bool())
	} else {
		data.ReverseRouteInjection = types.BoolNull()
	}
	if value := res.Get("enableClientServices"); value.Exists() {
		data.ClientServices = types.BoolValue(value.Bool())
	} else {
		data.ClientServices = types.BoolNull()
	}
	if value := res.Get("clientServicesPort"); value.Exists() {
		data.ClientServicesPort = types.Int64Value(value.Int())
	} else {
		data.ClientServicesPort = types.Int64Null()
	}
	if value := res.Get("perfectForwardSecracy.enabled"); value.Exists() {
		data.PerfectForwardSecrecy = types.BoolValue(value.Bool())
	} else {
		data.PerfectForwardSecrecy = types.BoolNull()
	}
	if value := res.Get("perfectForwardSecracy.modulusGroup"); value.Exists() {
		data.PerfectForwardSecrecyModulusGroup = types.StringValue(value.String())
	} else {
		data.PerfectForwardSecrecyModulusGroup = types.StringNull()
	}
	if value := res.Get("lifeTimeSeconds"); value.Exists() {
		data.LifetimeDuration = types.Int64Value(value.Int())
	} else {
		data.LifetimeDuration = types.Int64Null()
	}
	if value := res.Get("lifeTimeKilobytes"); value.Exists() {
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
func (data *VPNRAIPSecCryptoMap) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("interfaceObject.id"); value.Exists() && !data.InterfaceId.IsNull() {
		data.InterfaceId = types.StringValue(value.String())
	} else {
		data.InterfaceId = types.StringNull()
	}
	for i := 0; i < len(data.Ikev2IpsecProposals); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Ikev2IpsecProposals[i].Id.ValueString()}

		parent := &data
		data := (*parent).Ikev2IpsecProposals[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ikev2IpsecProposals").ForEach(
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
		(*parent).Ikev2IpsecProposals[i] = data
	}
	if value := res.Get("enableRRI"); value.Exists() && !data.ReverseRouteInjection.IsNull() {
		data.ReverseRouteInjection = types.BoolValue(value.Bool())
	} else {
		data.ReverseRouteInjection = types.BoolNull()
	}
	if value := res.Get("enableClientServices"); value.Exists() && !data.ClientServices.IsNull() {
		data.ClientServices = types.BoolValue(value.Bool())
	} else {
		data.ClientServices = types.BoolNull()
	}
	if value := res.Get("clientServicesPort"); value.Exists() && !data.ClientServicesPort.IsNull() {
		data.ClientServicesPort = types.Int64Value(value.Int())
	} else {
		data.ClientServicesPort = types.Int64Null()
	}
	if value := res.Get("perfectForwardSecracy.enabled"); value.Exists() && !data.PerfectForwardSecrecy.IsNull() {
		data.PerfectForwardSecrecy = types.BoolValue(value.Bool())
	} else {
		data.PerfectForwardSecrecy = types.BoolNull()
	}
	if value := res.Get("perfectForwardSecracy.modulusGroup"); value.Exists() && !data.PerfectForwardSecrecyModulusGroup.IsNull() {
		data.PerfectForwardSecrecyModulusGroup = types.StringValue(value.String())
	} else {
		data.PerfectForwardSecrecyModulusGroup = types.StringNull()
	}
	if value := res.Get("lifeTimeSeconds"); value.Exists() && !data.LifetimeDuration.IsNull() {
		data.LifetimeDuration = types.Int64Value(value.Int())
	} else {
		data.LifetimeDuration = types.Int64Null()
	}
	if value := res.Get("lifeTimeKilobytes"); value.Exists() && !data.LifetimeSize.IsNull() {
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
func (data *VPNRAIPSecCryptoMap) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
