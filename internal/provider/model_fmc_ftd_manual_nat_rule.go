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

type FTDManualNATRule struct {
	Id                             types.String `tfsdk:"id"`
	Domain                         types.String `tfsdk:"domain"`
	FtdNatPolicyId                 types.String `tfsdk:"ftd_nat_policy_id"`
	Type                           types.String `tfsdk:"type"`
	Description                    types.String `tfsdk:"description"`
	Enabled                        types.Bool   `tfsdk:"enabled"`
	Section                        types.String `tfsdk:"section"`
	NatType                        types.String `tfsdk:"nat_type"`
	FallThrough                    types.Bool   `tfsdk:"fall_through"`
	InterfaceInOriginalDestination types.Bool   `tfsdk:"interface_in_original_destination"`
	InterfaceInTranslatedSource    types.Bool   `tfsdk:"interface_in_translated_source"`
	Ipv6                           types.Bool   `tfsdk:"ipv6"`
	NetToNet                       types.Bool   `tfsdk:"net_to_net"`
	NoProxyArp                     types.Bool   `tfsdk:"no_proxy_arp"`
	Unidirectional                 types.Bool   `tfsdk:"unidirectional"`
	SourceInterfaceId              types.String `tfsdk:"source_interface_id"`
	OriginalSourceId               types.String `tfsdk:"original_source_id"`
	OriginalSourcePortId           types.String `tfsdk:"original_source_port_id"`
	OriginalDestinationId          types.String `tfsdk:"original_destination_id"`
	OriginalDestinationPortId      types.String `tfsdk:"original_destination_port_id"`
	RouteLookup                    types.Bool   `tfsdk:"route_lookup"`
	DestinationInterfaceId         types.String `tfsdk:"destination_interface_id"`
	TranslatedSourceId             types.String `tfsdk:"translated_source_id"`
	TranslatedSourcePortId         types.String `tfsdk:"translated_source_port_id"`
	TranslateDns                   types.Bool   `tfsdk:"translate_dns"`
	TranslatedDestinationId        types.String `tfsdk:"translated_destination_id"`
	TranslatedDestinationPortId    types.String `tfsdk:"translated_destination_port_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDManualNATRule) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdnatpolicies/%v/manualnatrules", url.QueryEscape(data.FtdNatPolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDManualNATRule) toBody(ctx context.Context, state FTDManualNATRule) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.Section.IsNull() {
		body, _ = sjson.Set(body, "metadata.section", data.Section.ValueString())
	}
	if !data.NatType.IsNull() {
		body, _ = sjson.Set(body, "natType", data.NatType.ValueString())
	}
	if !data.FallThrough.IsNull() {
		body, _ = sjson.Set(body, "fallThrough", data.FallThrough.ValueBool())
	}
	if !data.InterfaceInOriginalDestination.IsNull() {
		body, _ = sjson.Set(body, "interfaceInOriginalDestination", data.InterfaceInOriginalDestination.ValueBool())
	}
	if !data.InterfaceInTranslatedSource.IsNull() {
		body, _ = sjson.Set(body, "interfaceInTranslatedSource", data.InterfaceInTranslatedSource.ValueBool())
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
	if !data.Unidirectional.IsNull() {
		body, _ = sjson.Set(body, "unidirectional", data.Unidirectional.ValueBool())
	}
	if !data.SourceInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "sourceInterface.id", data.SourceInterfaceId.ValueString())
	}
	if !data.OriginalSourceId.IsNull() {
		body, _ = sjson.Set(body, "originalSource.id", data.OriginalSourceId.ValueString())
	}
	if !data.OriginalSourcePortId.IsNull() {
		body, _ = sjson.Set(body, "originalSourcePort.id", data.OriginalSourcePortId.ValueString())
	}
	if !data.OriginalDestinationId.IsNull() {
		body, _ = sjson.Set(body, "originalDestination.id", data.OriginalDestinationId.ValueString())
	}
	if !data.OriginalDestinationPortId.IsNull() {
		body, _ = sjson.Set(body, "originalDestinationPort.id", data.OriginalDestinationPortId.ValueString())
	}
	if !data.RouteLookup.IsNull() {
		body, _ = sjson.Set(body, "routeLookup", data.RouteLookup.ValueBool())
	}
	if !data.DestinationInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "destinationInterface.id", data.DestinationInterfaceId.ValueString())
	}
	if !data.TranslatedSourceId.IsNull() {
		body, _ = sjson.Set(body, "translatedSource.id", data.TranslatedSourceId.ValueString())
	}
	if !data.TranslatedSourcePortId.IsNull() {
		body, _ = sjson.Set(body, "translatedSourcePort.id", data.TranslatedSourcePortId.ValueString())
	}
	if !data.TranslateDns.IsNull() {
		body, _ = sjson.Set(body, "dns", data.TranslateDns.ValueBool())
	}
	if !data.TranslatedDestinationId.IsNull() {
		body, _ = sjson.Set(body, "translatedDestination.id", data.TranslatedDestinationId.ValueString())
	}
	if !data.TranslatedDestinationPortId.IsNull() {
		body, _ = sjson.Set(body, "translatedDestinationPort.id", data.TranslatedDestinationPortId.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDManualNATRule) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("metadata.section"); value.Exists() {
		data.Section = types.StringValue(value.String())
	} else {
		data.Section = types.StringNull()
	}
	if value := res.Get("natType"); value.Exists() {
		data.NatType = types.StringValue(value.String())
	} else {
		data.NatType = types.StringNull()
	}
	if value := res.Get("fallThrough"); value.Exists() {
		data.FallThrough = types.BoolValue(value.Bool())
	} else {
		data.FallThrough = types.BoolNull()
	}
	if value := res.Get("interfaceInOriginalDestination"); value.Exists() {
		data.InterfaceInOriginalDestination = types.BoolValue(value.Bool())
	} else {
		data.InterfaceInOriginalDestination = types.BoolNull()
	}
	if value := res.Get("interfaceInTranslatedSource"); value.Exists() {
		data.InterfaceInTranslatedSource = types.BoolValue(value.Bool())
	} else {
		data.InterfaceInTranslatedSource = types.BoolNull()
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
	if value := res.Get("unidirectional"); value.Exists() {
		data.Unidirectional = types.BoolValue(value.Bool())
	} else {
		data.Unidirectional = types.BoolNull()
	}
	if value := res.Get("sourceInterface.id"); value.Exists() {
		data.SourceInterfaceId = types.StringValue(value.String())
	} else {
		data.SourceInterfaceId = types.StringNull()
	}
	if value := res.Get("originalSource.id"); value.Exists() {
		data.OriginalSourceId = types.StringValue(value.String())
	} else {
		data.OriginalSourceId = types.StringNull()
	}
	if value := res.Get("originalSourcePort.id"); value.Exists() {
		data.OriginalSourcePortId = types.StringValue(value.String())
	} else {
		data.OriginalSourcePortId = types.StringNull()
	}
	if value := res.Get("originalDestination.id"); value.Exists() {
		data.OriginalDestinationId = types.StringValue(value.String())
	} else {
		data.OriginalDestinationId = types.StringNull()
	}
	if value := res.Get("originalDestinationPort.id"); value.Exists() {
		data.OriginalDestinationPortId = types.StringValue(value.String())
	} else {
		data.OriginalDestinationPortId = types.StringNull()
	}
	if value := res.Get("routeLookup"); value.Exists() {
		data.RouteLookup = types.BoolValue(value.Bool())
	} else {
		data.RouteLookup = types.BoolNull()
	}
	if value := res.Get("destinationInterface.id"); value.Exists() {
		data.DestinationInterfaceId = types.StringValue(value.String())
	} else {
		data.DestinationInterfaceId = types.StringNull()
	}
	if value := res.Get("translatedSource.id"); value.Exists() {
		data.TranslatedSourceId = types.StringValue(value.String())
	} else {
		data.TranslatedSourceId = types.StringNull()
	}
	if value := res.Get("translatedSourcePort.id"); value.Exists() {
		data.TranslatedSourcePortId = types.StringValue(value.String())
	} else {
		data.TranslatedSourcePortId = types.StringNull()
	}
	if value := res.Get("dns"); value.Exists() {
		data.TranslateDns = types.BoolValue(value.Bool())
	} else {
		data.TranslateDns = types.BoolNull()
	}
	if value := res.Get("translatedDestination.id"); value.Exists() {
		data.TranslatedDestinationId = types.StringValue(value.String())
	} else {
		data.TranslatedDestinationId = types.StringNull()
	}
	if value := res.Get("translatedDestinationPort.id"); value.Exists() {
		data.TranslatedDestinationPortId = types.StringValue(value.String())
	} else {
		data.TranslatedDestinationPortId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FTDManualNATRule) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("metadata.section"); value.Exists() && !data.Section.IsNull() {
		data.Section = types.StringValue(value.String())
	} else {
		data.Section = types.StringNull()
	}
	if value := res.Get("natType"); value.Exists() && !data.NatType.IsNull() {
		data.NatType = types.StringValue(value.String())
	} else {
		data.NatType = types.StringNull()
	}
	if value := res.Get("fallThrough"); value.Exists() && !data.FallThrough.IsNull() {
		data.FallThrough = types.BoolValue(value.Bool())
	} else {
		data.FallThrough = types.BoolNull()
	}
	if value := res.Get("interfaceInOriginalDestination"); value.Exists() && !data.InterfaceInOriginalDestination.IsNull() {
		data.InterfaceInOriginalDestination = types.BoolValue(value.Bool())
	} else {
		data.InterfaceInOriginalDestination = types.BoolNull()
	}
	if value := res.Get("interfaceInTranslatedSource"); value.Exists() && !data.InterfaceInTranslatedSource.IsNull() {
		data.InterfaceInTranslatedSource = types.BoolValue(value.Bool())
	} else {
		data.InterfaceInTranslatedSource = types.BoolNull()
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
	if value := res.Get("unidirectional"); value.Exists() && !data.Unidirectional.IsNull() {
		data.Unidirectional = types.BoolValue(value.Bool())
	} else {
		data.Unidirectional = types.BoolNull()
	}
	if value := res.Get("sourceInterface.id"); value.Exists() && !data.SourceInterfaceId.IsNull() {
		data.SourceInterfaceId = types.StringValue(value.String())
	} else {
		data.SourceInterfaceId = types.StringNull()
	}
	if value := res.Get("originalSource.id"); value.Exists() && !data.OriginalSourceId.IsNull() {
		data.OriginalSourceId = types.StringValue(value.String())
	} else {
		data.OriginalSourceId = types.StringNull()
	}
	if value := res.Get("originalSourcePort.id"); value.Exists() && !data.OriginalSourcePortId.IsNull() {
		data.OriginalSourcePortId = types.StringValue(value.String())
	} else {
		data.OriginalSourcePortId = types.StringNull()
	}
	if value := res.Get("originalDestination.id"); value.Exists() && !data.OriginalDestinationId.IsNull() {
		data.OriginalDestinationId = types.StringValue(value.String())
	} else {
		data.OriginalDestinationId = types.StringNull()
	}
	if value := res.Get("originalDestinationPort.id"); value.Exists() && !data.OriginalDestinationPortId.IsNull() {
		data.OriginalDestinationPortId = types.StringValue(value.String())
	} else {
		data.OriginalDestinationPortId = types.StringNull()
	}
	if value := res.Get("routeLookup"); value.Exists() && !data.RouteLookup.IsNull() {
		data.RouteLookup = types.BoolValue(value.Bool())
	} else {
		data.RouteLookup = types.BoolNull()
	}
	if value := res.Get("destinationInterface.id"); value.Exists() && !data.DestinationInterfaceId.IsNull() {
		data.DestinationInterfaceId = types.StringValue(value.String())
	} else {
		data.DestinationInterfaceId = types.StringNull()
	}
	if value := res.Get("translatedSource.id"); value.Exists() && !data.TranslatedSourceId.IsNull() {
		data.TranslatedSourceId = types.StringValue(value.String())
	} else {
		data.TranslatedSourceId = types.StringNull()
	}
	if value := res.Get("translatedSourcePort.id"); value.Exists() && !data.TranslatedSourcePortId.IsNull() {
		data.TranslatedSourcePortId = types.StringValue(value.String())
	} else {
		data.TranslatedSourcePortId = types.StringNull()
	}
	if value := res.Get("dns"); value.Exists() && !data.TranslateDns.IsNull() {
		data.TranslateDns = types.BoolValue(value.Bool())
	} else {
		data.TranslateDns = types.BoolNull()
	}
	if value := res.Get("translatedDestination.id"); value.Exists() && !data.TranslatedDestinationId.IsNull() {
		data.TranslatedDestinationId = types.StringValue(value.String())
	} else {
		data.TranslatedDestinationId = types.StringNull()
	}
	if value := res.Get("translatedDestinationPort.id"); value.Exists() && !data.TranslatedDestinationPortId.IsNull() {
		data.TranslatedDestinationPortId = types.StringValue(value.String())
	} else {
		data.TranslatedDestinationPortId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDManualNATRule) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

func (data FTDManualNATRule) adjustBody(ctx context.Context, req string) string {

	req, _ = sjson.Delete(req, "metadata")

	return req
}
