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

type FTDAutoNATRule struct {
	Id                                      types.String `tfsdk:"id"`
	Domain                                  types.String `tfsdk:"domain"`
	FtdNatPolicyId                          types.String `tfsdk:"ftd_nat_policy_id"`
	Type                                    types.String `tfsdk:"type"`
	NatType                                 types.String `tfsdk:"nat_type"`
	DestinationInterfaceId                  types.String `tfsdk:"destination_interface_id"`
	FallThrough                             types.Bool   `tfsdk:"fall_through"`
	Ipv6                                    types.Bool   `tfsdk:"ipv6"`
	NetToNet                                types.Bool   `tfsdk:"net_to_net"`
	NoProxyArp                              types.Bool   `tfsdk:"no_proxy_arp"`
	OriginalNetworkId                       types.String `tfsdk:"original_network_id"`
	OriginalPort                            types.Int64  `tfsdk:"original_port"`
	Protocol                                types.String `tfsdk:"protocol"`
	PerformRouteLookup                      types.Bool   `tfsdk:"perform_route_lookup"`
	SourceInterfaceId                       types.String `tfsdk:"source_interface_id"`
	TranslateDns                            types.Bool   `tfsdk:"translate_dns"`
	TranslatedNetworkId                     types.String `tfsdk:"translated_network_id"`
	TranslatedNetworkIsDestinationInterface types.Bool   `tfsdk:"translated_network_is_destination_interface"`
	TranslatedPort                          types.Int64  `tfsdk:"translated_port"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDAutoNATRule) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdnatpolicies/%v/autonatrules", url.QueryEscape(data.FtdNatPolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDAutoNATRule) toBody(ctx context.Context, state FTDAutoNATRule) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.NatType.IsNull() {
		body, _ = sjson.Set(body, "natType", data.NatType.ValueString())
	}
	if !data.DestinationInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "destinationInterface.id", data.DestinationInterfaceId.ValueString())
	}
	if !data.FallThrough.IsNull() {
		body, _ = sjson.Set(body, "fallThrough", data.FallThrough.ValueBool())
	}
	if !data.Ipv6.IsNull() {
		body, _ = sjson.Set(body, "interfaceIpv6", data.Ipv6.ValueBool())
	}
	if !data.NetToNet.IsNull() {
		body, _ = sjson.Set(body, "netToNet", data.NetToNet.ValueBool())
	}
	if !data.NoProxyArp.IsNull() {
		body, _ = sjson.Set(body, "noProxyArp", data.NoProxyArp.ValueBool())
	}
	if !data.OriginalNetworkId.IsNull() {
		body, _ = sjson.Set(body, "originalNetwork.id", data.OriginalNetworkId.ValueString())
	}
	if !data.OriginalPort.IsNull() {
		body, _ = sjson.Set(body, "originalPort", data.OriginalPort.ValueInt64())
	}
	if !data.Protocol.IsNull() {
		body, _ = sjson.Set(body, "serviceProtocol", data.Protocol.ValueString())
	}
	if !data.PerformRouteLookup.IsNull() {
		body, _ = sjson.Set(body, "routeLookup", data.PerformRouteLookup.ValueBool())
	}
	if !data.SourceInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "sourceInterface.id", data.SourceInterfaceId.ValueString())
	}
	if !data.TranslateDns.IsNull() {
		body, _ = sjson.Set(body, "dns", data.TranslateDns.ValueBool())
	}
	if !data.TranslatedNetworkId.IsNull() {
		body, _ = sjson.Set(body, "translatedNetwork.id", data.TranslatedNetworkId.ValueString())
	}
	if !data.TranslatedNetworkIsDestinationInterface.IsNull() {
		body, _ = sjson.Set(body, "interfaceInTranslatedNetwork", data.TranslatedNetworkIsDestinationInterface.ValueBool())
	}
	if !data.TranslatedPort.IsNull() {
		body, _ = sjson.Set(body, "translatedPort", data.TranslatedPort.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDAutoNATRule) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("natType"); value.Exists() {
		data.NatType = types.StringValue(value.String())
	} else {
		data.NatType = types.StringNull()
	}
	if value := res.Get("destinationInterface.id"); value.Exists() {
		data.DestinationInterfaceId = types.StringValue(value.String())
	} else {
		data.DestinationInterfaceId = types.StringNull()
	}
	if value := res.Get("fallThrough"); value.Exists() {
		data.FallThrough = types.BoolValue(value.Bool())
	} else {
		data.FallThrough = types.BoolNull()
	}
	if value := res.Get("interfaceIpv6"); value.Exists() {
		data.Ipv6 = types.BoolValue(value.Bool())
	} else {
		data.Ipv6 = types.BoolNull()
	}
	if value := res.Get("netToNet"); value.Exists() {
		data.NetToNet = types.BoolValue(value.Bool())
	} else {
		data.NetToNet = types.BoolNull()
	}
	if value := res.Get("noProxyArp"); value.Exists() {
		data.NoProxyArp = types.BoolValue(value.Bool())
	} else {
		data.NoProxyArp = types.BoolNull()
	}
	if value := res.Get("originalNetwork.id"); value.Exists() {
		data.OriginalNetworkId = types.StringValue(value.String())
	} else {
		data.OriginalNetworkId = types.StringNull()
	}
	if value := res.Get("originalPort"); value.Exists() {
		data.OriginalPort = types.Int64Value(value.Int())
	} else {
		data.OriginalPort = types.Int64Null()
	}
	if value := res.Get("serviceProtocol"); value.Exists() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("routeLookup"); value.Exists() {
		data.PerformRouteLookup = types.BoolValue(value.Bool())
	} else {
		data.PerformRouteLookup = types.BoolNull()
	}
	if value := res.Get("sourceInterface.id"); value.Exists() {
		data.SourceInterfaceId = types.StringValue(value.String())
	} else {
		data.SourceInterfaceId = types.StringNull()
	}
	if value := res.Get("dns"); value.Exists() {
		data.TranslateDns = types.BoolValue(value.Bool())
	} else {
		data.TranslateDns = types.BoolNull()
	}
	if value := res.Get("translatedNetwork.id"); value.Exists() {
		data.TranslatedNetworkId = types.StringValue(value.String())
	} else {
		data.TranslatedNetworkId = types.StringNull()
	}
	if value := res.Get("interfaceInTranslatedNetwork"); value.Exists() {
		data.TranslatedNetworkIsDestinationInterface = types.BoolValue(value.Bool())
	} else {
		data.TranslatedNetworkIsDestinationInterface = types.BoolNull()
	}
	if value := res.Get("translatedPort"); value.Exists() {
		data.TranslatedPort = types.Int64Value(value.Int())
	} else {
		data.TranslatedPort = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FTDAutoNATRule) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("natType"); value.Exists() && !data.NatType.IsNull() {
		data.NatType = types.StringValue(value.String())
	} else {
		data.NatType = types.StringNull()
	}
	if value := res.Get("destinationInterface.id"); value.Exists() && !data.DestinationInterfaceId.IsNull() {
		data.DestinationInterfaceId = types.StringValue(value.String())
	} else {
		data.DestinationInterfaceId = types.StringNull()
	}
	if value := res.Get("fallThrough"); value.Exists() && !data.FallThrough.IsNull() {
		data.FallThrough = types.BoolValue(value.Bool())
	} else {
		data.FallThrough = types.BoolNull()
	}
	if value := res.Get("interfaceIpv6"); value.Exists() && !data.Ipv6.IsNull() {
		data.Ipv6 = types.BoolValue(value.Bool())
	} else {
		data.Ipv6 = types.BoolNull()
	}
	if value := res.Get("netToNet"); value.Exists() && !data.NetToNet.IsNull() {
		data.NetToNet = types.BoolValue(value.Bool())
	} else {
		data.NetToNet = types.BoolNull()
	}
	if value := res.Get("noProxyArp"); value.Exists() && !data.NoProxyArp.IsNull() {
		data.NoProxyArp = types.BoolValue(value.Bool())
	} else {
		data.NoProxyArp = types.BoolNull()
	}
	if value := res.Get("originalNetwork.id"); value.Exists() && !data.OriginalNetworkId.IsNull() {
		data.OriginalNetworkId = types.StringValue(value.String())
	} else {
		data.OriginalNetworkId = types.StringNull()
	}
	if value := res.Get("originalPort"); value.Exists() && !data.OriginalPort.IsNull() {
		data.OriginalPort = types.Int64Value(value.Int())
	} else {
		data.OriginalPort = types.Int64Null()
	}
	if value := res.Get("serviceProtocol"); value.Exists() && !data.Protocol.IsNull() {
		data.Protocol = types.StringValue(value.String())
	} else {
		data.Protocol = types.StringNull()
	}
	if value := res.Get("routeLookup"); value.Exists() && !data.PerformRouteLookup.IsNull() {
		data.PerformRouteLookup = types.BoolValue(value.Bool())
	} else {
		data.PerformRouteLookup = types.BoolNull()
	}
	if value := res.Get("sourceInterface.id"); value.Exists() && !data.SourceInterfaceId.IsNull() {
		data.SourceInterfaceId = types.StringValue(value.String())
	} else {
		data.SourceInterfaceId = types.StringNull()
	}
	if value := res.Get("dns"); value.Exists() && !data.TranslateDns.IsNull() {
		data.TranslateDns = types.BoolValue(value.Bool())
	} else {
		data.TranslateDns = types.BoolNull()
	}
	if value := res.Get("translatedNetwork.id"); value.Exists() && !data.TranslatedNetworkId.IsNull() {
		data.TranslatedNetworkId = types.StringValue(value.String())
	} else {
		data.TranslatedNetworkId = types.StringNull()
	}
	if value := res.Get("interfaceInTranslatedNetwork"); value.Exists() && !data.TranslatedNetworkIsDestinationInterface.IsNull() {
		data.TranslatedNetworkIsDestinationInterface = types.BoolValue(value.Bool())
	} else {
		data.TranslatedNetworkIsDestinationInterface = types.BoolNull()
	}
	if value := res.Get("translatedPort"); value.Exists() && !data.TranslatedPort.IsNull() {
		data.TranslatedPort = types.Int64Value(value.Int())
	} else {
		data.TranslatedPort = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDAutoNATRule) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
