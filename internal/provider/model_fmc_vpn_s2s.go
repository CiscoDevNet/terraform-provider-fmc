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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type VPNS2S struct {
	Id                       types.String `tfsdk:"id"`
	Domain                   types.String `tfsdk:"domain"`
	Name                     types.String `tfsdk:"name"`
	Type                     types.String `tfsdk:"type"`
	RouteBased               types.Bool   `tfsdk:"route_based"`
	NetworkTopology          types.String `tfsdk:"network_topology"`
	Ikev1                    types.Bool   `tfsdk:"ikev1"`
	Ikev2                    types.Bool   `tfsdk:"ikev2"`
	IpsecPolicyId            types.String `tfsdk:"ipsec_policy_id"`
	IkePolicyId              types.String `tfsdk:"ike_policy_id"`
	AdvancedSettingsPolicyId types.String `tfsdk:"advanced_settings_policy_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNS2S) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftds2svpns"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNS2S) toBody(ctx context.Context, state VPNS2S) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.RouteBased.IsNull() {
		body, _ = sjson.Set(body, "routeBased", data.RouteBased.ValueBool())
	}
	if !data.NetworkTopology.IsNull() {
		body, _ = sjson.Set(body, "topologyType", data.NetworkTopology.ValueString())
	}
	if !data.Ikev1.IsNull() {
		body, _ = sjson.Set(body, "ikeV1Enabled", data.Ikev1.ValueBool())
	}
	if !data.Ikev2.IsNull() {
		body, _ = sjson.Set(body, "ikeV2Enabled", data.Ikev2.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNS2S) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("routeBased"); value.Exists() {
		data.RouteBased = types.BoolValue(value.Bool())
	} else {
		data.RouteBased = types.BoolNull()
	}
	if value := res.Get("topologyType"); value.Exists() {
		data.NetworkTopology = types.StringValue(value.String())
	} else {
		data.NetworkTopology = types.StringNull()
	}
	if value := res.Get("ikeV1Enabled"); value.Exists() {
		data.Ikev1 = types.BoolValue(value.Bool())
	} else {
		data.Ikev1 = types.BoolValue(false)
	}
	if value := res.Get("ikeV2Enabled"); value.Exists() {
		data.Ikev2 = types.BoolValue(value.Bool())
	} else {
		data.Ikev2 = types.BoolValue(false)
	}
	if value := res.Get("ipsecSettings.id"); value.Exists() {
		data.IpsecPolicyId = types.StringValue(value.String())
	} else {
		data.IpsecPolicyId = types.StringNull()
	}
	if value := res.Get("ikeSettings.id"); value.Exists() {
		data.IkePolicyId = types.StringValue(value.String())
	} else {
		data.IkePolicyId = types.StringNull()
	}
	if value := res.Get("advancedSettings.id"); value.Exists() {
		data.AdvancedSettingsPolicyId = types.StringValue(value.String())
	} else {
		data.AdvancedSettingsPolicyId = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNS2S) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("routeBased"); value.Exists() && !data.RouteBased.IsNull() {
		data.RouteBased = types.BoolValue(value.Bool())
	} else {
		data.RouteBased = types.BoolNull()
	}
	if value := res.Get("topologyType"); value.Exists() && !data.NetworkTopology.IsNull() {
		data.NetworkTopology = types.StringValue(value.String())
	} else {
		data.NetworkTopology = types.StringNull()
	}
	if value := res.Get("ikeV1Enabled"); value.Exists() && !data.Ikev1.IsNull() {
		data.Ikev1 = types.BoolValue(value.Bool())
	} else if data.Ikev1.ValueBool() != false {
		data.Ikev1 = types.BoolNull()
	}
	if value := res.Get("ikeV2Enabled"); value.Exists() && !data.Ikev2.IsNull() {
		data.Ikev2 = types.BoolValue(value.Bool())
	} else if data.Ikev2.ValueBool() != false {
		data.Ikev2 = types.BoolNull()
	}
	if value := res.Get("ipsecSettings.id"); value.Exists() && !data.IpsecPolicyId.IsNull() {
		data.IpsecPolicyId = types.StringValue(value.String())
	} else {
		data.IpsecPolicyId = types.StringNull()
	}
	if value := res.Get("ikeSettings.id"); value.Exists() && !data.IkePolicyId.IsNull() {
		data.IkePolicyId = types.StringValue(value.String())
	} else {
		data.IkePolicyId = types.StringNull()
	}
	if value := res.Get("advancedSettings.id"); value.Exists() && !data.AdvancedSettingsPolicyId.IsNull() {
		data.AdvancedSettingsPolicyId = types.StringValue(value.String())
	} else {
		data.AdvancedSettingsPolicyId = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNS2S) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.IpsecPolicyId.IsUnknown() {
		if value := res.Get("ipsecSettings.id"); value.Exists() {
			data.IpsecPolicyId = types.StringValue(value.String())
		} else {
			data.IpsecPolicyId = types.StringNull()
		}
	}
	if data.IkePolicyId.IsUnknown() {
		if value := res.Get("ikeSettings.id"); value.Exists() {
			data.IkePolicyId = types.StringValue(value.String())
		} else {
			data.IkePolicyId = types.StringNull()
		}
	}
	if data.AdvancedSettingsPolicyId.IsUnknown() {
		if value := res.Get("advancedSettings.id"); value.Exists() {
			data.AdvancedSettingsPolicyId = types.StringValue(value.String())
		} else {
			data.AdvancedSettingsPolicyId = types.StringNull()
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
