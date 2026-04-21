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

type IntrusionPolicyRuleOverride struct {
	Id                types.String `tfsdk:"id"`
	Domain            types.String `tfsdk:"domain"`
	IntrusionPolicyId types.String `tfsdk:"intrusion_policy_id"`
	IntrusionRuleId   types.String `tfsdk:"intrusion_rule_id"`
	Type              types.String `tfsdk:"type"`
	DefaultState      types.String `tfsdk:"default_state"`
	OverrideState     types.String `tfsdk:"override_state"`
	RuleData          types.String `tfsdk:"rule_data"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data IntrusionPolicyRuleOverride) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/intrusionpolicies/%v/intrusionrules", url.QueryEscape(data.IntrusionPolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data IntrusionPolicyRuleOverride) toBody(ctx context.Context, state IntrusionPolicyRuleOverride) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.IntrusionRuleId.IsNull() {
		body, _ = sjson.Set(body, "id", data.IntrusionRuleId.ValueString())
	}
	if !data.OverrideState.IsNull() {
		body, _ = sjson.Set(body, "overrideState", data.OverrideState.ValueString())
	}
	if !data.RuleData.IsNull() {
		body, _ = sjson.Set(body, "ruleData", data.RuleData.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *IntrusionPolicyRuleOverride) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("id"); value.Exists() {
		data.IntrusionRuleId = types.StringValue(value.String())
	} else {
		data.IntrusionRuleId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("defaultState"); value.Exists() {
		data.DefaultState = types.StringValue(value.String())
	} else {
		data.DefaultState = types.StringNull()
	}
	if value := res.Get("overrideState"); value.Exists() {
		data.OverrideState = types.StringValue(value.String())
	} else {
		data.OverrideState = types.StringNull()
	}
	if value := res.Get("ruleData"); value.Exists() {
		data.RuleData = types.StringValue(value.String())
	} else {
		data.RuleData = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *IntrusionPolicyRuleOverride) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("id"); value.Exists() && !data.IntrusionRuleId.IsNull() {
		data.IntrusionRuleId = types.StringValue(value.String())
	} else {
		data.IntrusionRuleId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("defaultState"); value.Exists() && !data.DefaultState.IsNull() {
		data.DefaultState = types.StringValue(value.String())
	} else {
		data.DefaultState = types.StringNull()
	}
	if value := res.Get("overrideState"); value.Exists() && !data.OverrideState.IsNull() {
		data.OverrideState = types.StringValue(value.String())
	} else {
		data.OverrideState = types.StringNull()
	}
	if value := res.Get("ruleData"); value.Exists() && !data.RuleData.IsNull() {
		data.RuleData = types.StringValue(value.String())
	} else {
		data.RuleData = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *IntrusionPolicyRuleOverride) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.DefaultState.IsUnknown() {
		if value := res.Get("defaultState"); value.Exists() {
			data.DefaultState = types.StringValue(value.String())
		} else {
			data.DefaultState = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data IntrusionPolicyRuleOverride) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	body, _ = sjson.Set(body, "ruleData", data.RuleData.ValueString())
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody
