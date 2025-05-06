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

type VPNS2SIKESettings struct {
	Id                                   types.String                     `tfsdk:"id"`
	Domain                               types.String                     `tfsdk:"domain"`
	VpnS2sId                             types.String                     `tfsdk:"vpn_s2s_id"`
	Type                                 types.String                     `tfsdk:"type"`
	Ikev1AuthenticationType              types.String                     `tfsdk:"ikev1_authentication_type"`
	Ikev1AutomaticPreSharedKeyLength     types.Int64                      `tfsdk:"ikev1_automatic_pre_shared_key_length"`
	Ikev1ManualPreSharedKey              types.String                     `tfsdk:"ikev1_manual_pre_shared_key"`
	Ikev1Policies                        []VPNS2SIKESettingsIkev1Policies `tfsdk:"ikev1_policies"`
	Ikev2AuthenticationType              types.String                     `tfsdk:"ikev2_authentication_type"`
	Ikev2AutomaticPreSharedKeyLength     types.Int64                      `tfsdk:"ikev2_automatic_pre_shared_key_length"`
	Ikev2ManualPreSharedKey              types.String                     `tfsdk:"ikev2_manual_pre_shared_key"`
	Ikev2EnforceHexBasedPreSharedKeyOnly types.Bool                       `tfsdk:"ikev2_enforce_hex_based_pre_shared_key_only"`
	Ikev2Policies                        []VPNS2SIKESettingsIkev2Policies `tfsdk:"ikev2_policies"`
}

type VPNS2SIKESettingsIkev1Policies struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type VPNS2SIKESettingsIkev2Policies struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data VPNS2SIKESettings) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftds2svpns/%v/ikesettings", url.QueryEscape(data.VpnS2sId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data VPNS2SIKESettings) toBody(ctx context.Context, state VPNS2SIKESettings) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Ikev1AuthenticationType.IsNull() {
		body, _ = sjson.Set(body, "ikeV1Settings.authenticationType", data.Ikev1AuthenticationType.ValueString())
	}
	if !data.Ikev1AutomaticPreSharedKeyLength.IsNull() {
		body, _ = sjson.Set(body, "ikeV1Settings.automaticPreSharedKeyLength", data.Ikev1AutomaticPreSharedKeyLength.ValueInt64())
	}
	if !data.Ikev1ManualPreSharedKey.IsNull() {
		body, _ = sjson.Set(body, "ikeV1Settings.manualPreSharedKey", data.Ikev1ManualPreSharedKey.ValueString())
	}
	if len(data.Ikev1Policies) > 0 {
		body, _ = sjson.Set(body, "ikeV1Settings.policies", []interface{}{})
		for _, item := range data.Ikev1Policies {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ikeV1Settings.policies.-1", itemBody)
		}
	}
	if !data.Ikev2AuthenticationType.IsNull() {
		body, _ = sjson.Set(body, "ikeV2Settings.authenticationType", data.Ikev2AuthenticationType.ValueString())
	}
	if !data.Ikev2AutomaticPreSharedKeyLength.IsNull() {
		body, _ = sjson.Set(body, "ikeV2Settings.automaticPreSharedKeyLength", data.Ikev2AutomaticPreSharedKeyLength.ValueInt64())
	}
	if !data.Ikev2ManualPreSharedKey.IsNull() {
		body, _ = sjson.Set(body, "ikeV2Settings.manualPreSharedKey", data.Ikev2ManualPreSharedKey.ValueString())
	}
	if !data.Ikev2EnforceHexBasedPreSharedKeyOnly.IsNull() {
		body, _ = sjson.Set(body, "ikeV2Settings.enforceHexBasedPreSharedKeyOnly", data.Ikev2EnforceHexBasedPreSharedKeyOnly.ValueBool())
	}
	if len(data.Ikev2Policies) > 0 {
		body, _ = sjson.Set(body, "ikeV2Settings.policies", []interface{}{})
		for _, item := range data.Ikev2Policies {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ikeV2Settings.policies.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *VPNS2SIKESettings) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("ikeV1Settings.authenticationType"); value.Exists() {
		data.Ikev1AuthenticationType = types.StringValue(value.String())
	} else {
		data.Ikev1AuthenticationType = types.StringNull()
	}
	if value := res.Get("ikeV1Settings.automaticPreSharedKeyLength"); value.Exists() {
		data.Ikev1AutomaticPreSharedKeyLength = types.Int64Value(value.Int())
	} else {
		data.Ikev1AutomaticPreSharedKeyLength = types.Int64Null()
	}
	if value := res.Get("ikeV1Settings.policies"); value.Exists() {
		data.Ikev1Policies = make([]VPNS2SIKESettingsIkev1Policies, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNS2SIKESettingsIkev1Policies{}
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
			(*parent).Ikev1Policies = append((*parent).Ikev1Policies, data)
			return true
		})
	}
	if value := res.Get("ikeV2Settings.authenticationType"); value.Exists() {
		data.Ikev2AuthenticationType = types.StringValue(value.String())
	} else {
		data.Ikev2AuthenticationType = types.StringNull()
	}
	if value := res.Get("ikeV2Settings.automaticPreSharedKeyLength"); value.Exists() {
		data.Ikev2AutomaticPreSharedKeyLength = types.Int64Value(value.Int())
	} else {
		data.Ikev2AutomaticPreSharedKeyLength = types.Int64Null()
	}
	if value := res.Get("ikeV2Settings.enforceHexBasedPreSharedKeyOnly"); value.Exists() {
		data.Ikev2EnforceHexBasedPreSharedKeyOnly = types.BoolValue(value.Bool())
	} else {
		data.Ikev2EnforceHexBasedPreSharedKeyOnly = types.BoolNull()
	}
	if value := res.Get("ikeV2Settings.policies"); value.Exists() {
		data.Ikev2Policies = make([]VPNS2SIKESettingsIkev2Policies, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := VPNS2SIKESettingsIkev2Policies{}
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
			(*parent).Ikev2Policies = append((*parent).Ikev2Policies, data)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *VPNS2SIKESettings) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("ikeV1Settings.authenticationType"); value.Exists() && !data.Ikev1AuthenticationType.IsNull() {
		data.Ikev1AuthenticationType = types.StringValue(value.String())
	} else {
		data.Ikev1AuthenticationType = types.StringNull()
	}
	if value := res.Get("ikeV1Settings.automaticPreSharedKeyLength"); value.Exists() && !data.Ikev1AutomaticPreSharedKeyLength.IsNull() {
		data.Ikev1AutomaticPreSharedKeyLength = types.Int64Value(value.Int())
	} else {
		data.Ikev1AutomaticPreSharedKeyLength = types.Int64Null()
	}
	for i := 0; i < len(data.Ikev1Policies); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Ikev1Policies[i].Id.ValueString()}

		parent := &data
		data := (*parent).Ikev1Policies[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ikeV1Settings.policies").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ikev1Policies[%d] = %+v",
				i,
				(*parent).Ikev1Policies[i],
			))
			(*parent).Ikev1Policies = slices.Delete((*parent).Ikev1Policies, i, i+1)
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
		(*parent).Ikev1Policies[i] = data
	}
	if value := res.Get("ikeV2Settings.authenticationType"); value.Exists() && !data.Ikev2AuthenticationType.IsNull() {
		data.Ikev2AuthenticationType = types.StringValue(value.String())
	} else {
		data.Ikev2AuthenticationType = types.StringNull()
	}
	if value := res.Get("ikeV2Settings.automaticPreSharedKeyLength"); value.Exists() && !data.Ikev2AutomaticPreSharedKeyLength.IsNull() {
		data.Ikev2AutomaticPreSharedKeyLength = types.Int64Value(value.Int())
	} else {
		data.Ikev2AutomaticPreSharedKeyLength = types.Int64Null()
	}
	if value := res.Get("ikeV2Settings.enforceHexBasedPreSharedKeyOnly"); value.Exists() && !data.Ikev2EnforceHexBasedPreSharedKeyOnly.IsNull() {
		data.Ikev2EnforceHexBasedPreSharedKeyOnly = types.BoolValue(value.Bool())
	} else {
		data.Ikev2EnforceHexBasedPreSharedKeyOnly = types.BoolNull()
	}
	for i := 0; i < len(data.Ikev2Policies); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.Ikev2Policies[i].Id.ValueString()}

		parent := &data
		data := (*parent).Ikev2Policies[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ikeV2Settings.policies").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ikev2Policies[%d] = %+v",
				i,
				(*parent).Ikev2Policies[i],
			))
			(*parent).Ikev2Policies = slices.Delete((*parent).Ikev2Policies, i, i+1)
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
		(*parent).Ikev2Policies[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *VPNS2SIKESettings) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
