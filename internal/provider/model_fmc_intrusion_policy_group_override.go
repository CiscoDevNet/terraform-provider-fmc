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

type IntrusionPolicyGroupOverride struct {
	Id                    types.String `tfsdk:"id"`
	Domain                types.String `tfsdk:"domain"`
	IntrusionPolicyId     types.String `tfsdk:"intrusion_policy_id"`
	IntrusionRuleGroupId  types.String `tfsdk:"intrusion_rule_group_id"`
	Type                  types.String `tfsdk:"type"`
	DefaultSecurityLevel  types.String `tfsdk:"default_security_level"`
	OverrideSecurityLevel types.String `tfsdk:"override_security_level"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data IntrusionPolicyGroupOverride) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/intrusionpolicies/%v/intrusionrulegroups", url.QueryEscape(data.IntrusionPolicyId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data IntrusionPolicyGroupOverride) toBody(ctx context.Context, state IntrusionPolicyGroupOverride) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.IntrusionRuleGroupId.IsNull() {
		body, _ = sjson.Set(body, "id", data.IntrusionRuleGroupId.ValueString())
	}
	if !data.OverrideSecurityLevel.IsNull() {
		body, _ = sjson.Set(body, "overrideSecurityLevel", data.OverrideSecurityLevel.ValueString())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *IntrusionPolicyGroupOverride) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("id"); value.Exists() {
		data.IntrusionRuleGroupId = types.StringValue(value.String())
	} else {
		data.IntrusionRuleGroupId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("defaultSecurityLevel"); value.Exists() {
		data.DefaultSecurityLevel = types.StringValue(value.String())
	} else {
		data.DefaultSecurityLevel = types.StringNull()
	}
	if value := res.Get("overrideSecurityLevel"); value.Exists() {
		data.OverrideSecurityLevel = types.StringValue(value.String())
	} else {
		data.OverrideSecurityLevel = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *IntrusionPolicyGroupOverride) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("id"); value.Exists() && !data.IntrusionRuleGroupId.IsNull() {
		data.IntrusionRuleGroupId = types.StringValue(value.String())
	} else {
		data.IntrusionRuleGroupId = types.StringNull()
	}
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("defaultSecurityLevel"); value.Exists() && !data.DefaultSecurityLevel.IsNull() {
		data.DefaultSecurityLevel = types.StringValue(value.String())
	} else {
		data.DefaultSecurityLevel = types.StringNull()
	}
	if value := res.Get("overrideSecurityLevel"); value.Exists() && !data.OverrideSecurityLevel.IsNull() {
		data.OverrideSecurityLevel = types.StringValue(value.String())
	} else {
		data.OverrideSecurityLevel = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *IntrusionPolicyGroupOverride) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.DefaultSecurityLevel.IsUnknown() {
		if value := res.Get("defaultSecurityLevel"); value.Exists() {
			data.DefaultSecurityLevel = types.StringValue(value.String())
		} else {
			data.DefaultSecurityLevel = types.StringNull()
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

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data IntrusionPolicyGroupOverride) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	body, _ = sjson.Set(body, "overrideSecurityLevel", "DEFAULT")
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

func (data IntrusionPolicyGroupOverride) adjustBody(ctx context.Context, req string) string {
	return req
}

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides
