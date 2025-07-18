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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SecureClientCustomAttribute struct {
	Id                                    types.String `tfsdk:"id"`
	Domain                                types.String `tfsdk:"domain"`
	Name                                  types.String `tfsdk:"name"`
	Type                                  types.String `tfsdk:"type"`
	Description                           types.String `tfsdk:"description"`
	AttributeType                         types.String `tfsdk:"attribute_type"`
	UserDefinedAttributeName              types.String `tfsdk:"user_defined_attribute_name"`
	UserDefinedAttributeValue             types.String `tfsdk:"user_defined_attribute_value"`
	PerAppVpnValue                        types.String `tfsdk:"per_app_vpn_value"`
	DynamicSplitTunnelIncludedDomains     types.List   `tfsdk:"dynamic_split_tunnel_included_domains"`
	DynamicSplitTunnelExcludedDomains     types.List   `tfsdk:"dynamic_split_tunnel_excluded_domains"`
	DeferUpdatePromptType                 types.String `tfsdk:"defer_update_prompt_type"`
	DeferUpdateDefaultAction              types.String `tfsdk:"defer_update_default_action"`
	DeferUpdateMinimumSecureClientVersion types.String `tfsdk:"defer_update_minimum_secure_client_version"`
	DeferUpdatePromptDismissTimeout       types.Int64  `tfsdk:"defer_update_prompt_dismiss_timeout"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionCreateSecureClientCustomAttribute = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SecureClientCustomAttribute) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/anyconnectcustomattributes"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SecureClientCustomAttribute) toBody(ctx context.Context, state SecureClientCustomAttribute) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	body, _ = sjson.Set(body, "type", "AnyConnectCustomAttribute")
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.AttributeType.IsNull() {
		body, _ = sjson.Set(body, "attributeType", data.AttributeType.ValueString())
	}
	if !data.UserDefinedAttributeName.IsNull() {
		body, _ = sjson.Set(body, "userDefinedAttribute.name", data.UserDefinedAttributeName.ValueString())
	}
	if !data.UserDefinedAttributeValue.IsNull() {
		body, _ = sjson.Set(body, "userDefinedAttribute.value", data.UserDefinedAttributeValue.ValueString())
	}
	if !data.PerAppVpnValue.IsNull() {
		body, _ = sjson.Set(body, "perAppVpnBase64EncodedValue", data.PerAppVpnValue.ValueString())
	}
	if !data.DynamicSplitTunnelIncludedDomains.IsNull() {
		var values []string
		data.DynamicSplitTunnelIncludedDomains.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dynamicSplitTunnel.includeDomains", values)
	}
	if !data.DynamicSplitTunnelExcludedDomains.IsNull() {
		var values []string
		data.DynamicSplitTunnelExcludedDomains.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "dynamicSplitTunnel.excludeDomains", values)
	}
	if !data.DeferUpdatePromptType.IsNull() {
		body, _ = sjson.Set(body, "deferUpdate.promptType", data.DeferUpdatePromptType.ValueString())
	}
	if !data.DeferUpdateDefaultAction.IsNull() {
		body, _ = sjson.Set(body, "deferUpdate.defaultAction", data.DeferUpdateDefaultAction.ValueString())
	}
	if !data.DeferUpdateMinimumSecureClientVersion.IsNull() {
		body, _ = sjson.Set(body, "deferUpdate.minimumAnyConnectVersion", data.DeferUpdateMinimumSecureClientVersion.ValueString())
	}
	if !data.DeferUpdatePromptDismissTimeout.IsNull() {
		body, _ = sjson.Set(body, "deferUpdate.promptDismissTimeout", data.DeferUpdatePromptDismissTimeout.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SecureClientCustomAttribute) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("attributeType"); value.Exists() {
		data.AttributeType = types.StringValue(value.String())
	} else {
		data.AttributeType = types.StringNull()
	}
	if value := res.Get("userDefinedAttribute.name"); value.Exists() {
		data.UserDefinedAttributeName = types.StringValue(value.String())
	} else {
		data.UserDefinedAttributeName = types.StringNull()
	}
	if value := res.Get("userDefinedAttribute.value"); value.Exists() {
		data.UserDefinedAttributeValue = types.StringValue(value.String())
	} else {
		data.UserDefinedAttributeValue = types.StringNull()
	}
	if value := res.Get("perAppVpnBase64EncodedValue"); value.Exists() {
		data.PerAppVpnValue = types.StringValue(value.String())
	} else {
		data.PerAppVpnValue = types.StringNull()
	}
	if value := res.Get("dynamicSplitTunnel.includeDomains"); value.Exists() {
		data.DynamicSplitTunnelIncludedDomains = helpers.GetStringList(value.Array())
	} else {
		data.DynamicSplitTunnelIncludedDomains = types.ListNull(types.StringType)
	}
	if value := res.Get("dynamicSplitTunnel.excludeDomains"); value.Exists() {
		data.DynamicSplitTunnelExcludedDomains = helpers.GetStringList(value.Array())
	} else {
		data.DynamicSplitTunnelExcludedDomains = types.ListNull(types.StringType)
	}
	if value := res.Get("deferUpdate.promptType"); value.Exists() {
		data.DeferUpdatePromptType = types.StringValue(value.String())
	} else {
		data.DeferUpdatePromptType = types.StringNull()
	}
	if value := res.Get("deferUpdate.defaultAction"); value.Exists() {
		data.DeferUpdateDefaultAction = types.StringValue(value.String())
	} else {
		data.DeferUpdateDefaultAction = types.StringNull()
	}
	if value := res.Get("deferUpdate.minimumAnyConnectVersion"); value.Exists() {
		data.DeferUpdateMinimumSecureClientVersion = types.StringValue(value.String())
	} else {
		data.DeferUpdateMinimumSecureClientVersion = types.StringNull()
	}
	if value := res.Get("deferUpdate.promptDismissTimeout"); value.Exists() {
		data.DeferUpdatePromptDismissTimeout = types.Int64Value(value.Int())
	} else {
		data.DeferUpdatePromptDismissTimeout = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SecureClientCustomAttribute) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("attributeType"); value.Exists() && !data.AttributeType.IsNull() {
		data.AttributeType = types.StringValue(value.String())
	} else {
		data.AttributeType = types.StringNull()
	}
	if value := res.Get("userDefinedAttribute.name"); value.Exists() && !data.UserDefinedAttributeName.IsNull() {
		data.UserDefinedAttributeName = types.StringValue(value.String())
	} else {
		data.UserDefinedAttributeName = types.StringNull()
	}
	if value := res.Get("userDefinedAttribute.value"); value.Exists() && !data.UserDefinedAttributeValue.IsNull() {
		data.UserDefinedAttributeValue = types.StringValue(value.String())
	} else {
		data.UserDefinedAttributeValue = types.StringNull()
	}
	if value := res.Get("perAppVpnBase64EncodedValue"); value.Exists() && !data.PerAppVpnValue.IsNull() {
		data.PerAppVpnValue = types.StringValue(value.String())
	} else {
		data.PerAppVpnValue = types.StringNull()
	}
	if value := res.Get("dynamicSplitTunnel.includeDomains"); value.Exists() && !data.DynamicSplitTunnelIncludedDomains.IsNull() {
		data.DynamicSplitTunnelIncludedDomains = helpers.GetStringList(value.Array())
	} else {
		data.DynamicSplitTunnelIncludedDomains = types.ListNull(types.StringType)
	}
	if value := res.Get("dynamicSplitTunnel.excludeDomains"); value.Exists() && !data.DynamicSplitTunnelExcludedDomains.IsNull() {
		data.DynamicSplitTunnelExcludedDomains = helpers.GetStringList(value.Array())
	} else {
		data.DynamicSplitTunnelExcludedDomains = types.ListNull(types.StringType)
	}
	if value := res.Get("deferUpdate.promptType"); value.Exists() && !data.DeferUpdatePromptType.IsNull() {
		data.DeferUpdatePromptType = types.StringValue(value.String())
	} else {
		data.DeferUpdatePromptType = types.StringNull()
	}
	if value := res.Get("deferUpdate.defaultAction"); value.Exists() && !data.DeferUpdateDefaultAction.IsNull() {
		data.DeferUpdateDefaultAction = types.StringValue(value.String())
	} else {
		data.DeferUpdateDefaultAction = types.StringNull()
	}
	if value := res.Get("deferUpdate.minimumAnyConnectVersion"); value.Exists() && !data.DeferUpdateMinimumSecureClientVersion.IsNull() {
		data.DeferUpdateMinimumSecureClientVersion = types.StringValue(value.String())
	} else {
		data.DeferUpdateMinimumSecureClientVersion = types.StringNull()
	}
	if value := res.Get("deferUpdate.promptDismissTimeout"); value.Exists() && !data.DeferUpdatePromptDismissTimeout.IsNull() {
		data.DeferUpdatePromptDismissTimeout = types.Int64Value(value.Int())
	} else {
		data.DeferUpdatePromptDismissTimeout = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SecureClientCustomAttribute) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
