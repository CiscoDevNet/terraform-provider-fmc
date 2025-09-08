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
	"maps"
	"slices"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type CertificateMaps struct {
	Id     types.String                    `tfsdk:"id"`
	Domain types.String                    `tfsdk:"domain"`
	Items  map[string]CertificateMapsItems `tfsdk:"items"`
}

type CertificateMapsItems struct {
	Id    types.String                `tfsdk:"id"`
	Type  types.String                `tfsdk:"type"`
	Rules []CertificateMapsItemsRules `tfsdk:"rules"`
}

type CertificateMapsItemsRules struct {
	Field     types.String `tfsdk:"field"`
	Component types.String `tfsdk:"component"`
	Operator  types.String `tfsdk:"operator"`
	Value     types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkCreateCertificateMaps = version.Must(version.NewVersion("999"))
var minFMCVersionBulkDeleteCertificateMaps = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data CertificateMaps) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/certificatemaps"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data CertificateMaps) toBody(ctx context.Context, state CertificateMaps) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "items", []any{})
		for key, item := range data.Items {
			itemBody, _ := sjson.Set("{}", "name", key)
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "CertificateMap")
			if len(item.Rules) > 0 {
				itemBody, _ = sjson.Set(itemBody, "rules", []any{})
				for _, childItem := range item.Rules {
					itemChildBody := ""
					if !childItem.Field.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "field", childItem.Field.ValueString())
					}
					if !childItem.Component.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "component", childItem.Component.ValueString())
					}
					if !childItem.Operator.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "operator", childItem.Operator.ValueString())
					}
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "rules.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *CertificateMaps) fromBody(ctx context.Context, res gjson.Result) {
	for k := range data.Items {
		parent := &data
		data := (*parent).Items[k]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("name").String() == k {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if !res.Exists() {
			tflog.Debug(ctx, fmt.Sprintf("subresource not found, removing: name=%v", k))
			delete((*parent).Items, k)
			continue
		}
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("rules"); value.Exists() {
			data.Rules = make([]CertificateMapsItemsRules, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := CertificateMapsItemsRules{}
				if value := res.Get("field"); value.Exists() {
					data.Field = types.StringValue(value.String())
				} else {
					data.Field = types.StringNull()
				}
				if value := res.Get("component"); value.Exists() {
					data.Component = types.StringValue(value.String())
				} else {
					data.Component = types.StringNull()
				}
				if value := res.Get("operator"); value.Exists() {
					data.Operator = types.StringValue(value.String())
				} else {
					data.Operator = types.StringNull()
				}
				if value := res.Get("value"); value.Exists() {
					data.Value = types.StringValue(value.String())
				} else {
					data.Value = types.StringNull()
				}
				(*parent).Rules = append((*parent).Rules, data)
				return true
			})
		}
		(*parent).Items[k] = data
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *CertificateMaps) fromBodyPartial(ctx context.Context, res gjson.Result) {
	for i := range data.Items {
		parent := &data
		data := (*parent).Items[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if v.Get("id").String() == data.Id.ValueString() && data.Id.ValueString() != "" {
					res = v
					return false // break ForEach
				}
				return true
			},
		)
		if value := res.Get("id"); value.Exists() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		for i := 0; i < len(data.Rules); i++ {
			keys := [...]string{"field", "component", "operator", "value"}
			keyValues := [...]string{data.Rules[i].Field.ValueString(), data.Rules[i].Component.ValueString(), data.Rules[i].Operator.ValueString(), data.Rules[i].Value.ValueString()}

			parent := &data
			data := (*parent).Rules[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("rules").ForEach(
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
				tflog.Debug(ctx, fmt.Sprintf("removing Rules[%d] = %+v",
					i,
					(*parent).Rules[i],
				))
				(*parent).Rules = slices.Delete((*parent).Rules, i, i+1)
				i--

				continue
			}
			if value := res.Get("field"); value.Exists() && !data.Field.IsNull() {
				data.Field = types.StringValue(value.String())
			} else {
				data.Field = types.StringNull()
			}
			if value := res.Get("component"); value.Exists() && !data.Component.IsNull() {
				data.Component = types.StringValue(value.String())
			} else {
				data.Component = types.StringNull()
			}
			if value := res.Get("operator"); value.Exists() && !data.Operator.IsNull() {
				data.Operator = types.StringValue(value.String())
			} else {
				data.Operator = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() && !data.Value.IsNull() {
				data.Value = types.StringValue(value.String())
			} else {
				data.Value = types.StringNull()
			}
			(*parent).Rules[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *CertificateMaps) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	for i, val := range data.Items {
		var r gjson.Result
		res.Get("items").ForEach(
			func(_, v gjson.Result) bool {
				if val.Id.IsUnknown() {
					if v.Get("name").String() == i {
						r = v
						return false // break ForEach
					}
				} else {
					if v.Get("id").String() == val.Id.ValueString() && val.Id.ValueString() != "" {
						r = v
						return false // break ForEach
					}
				}

				return true
			},
		)
		if v := data.Items[i]; v.Id.IsUnknown() {
			if value := r.Get("id"); value.Exists() {
				v.Id = types.StringValue(value.String())
			} else {
				v.Id = types.StringNull()
			}
			data.Items[i] = v
		}
		if v := data.Items[i]; v.Type.IsUnknown() {
			if value := r.Get("type"); value.Exists() {
				v.Type = types.StringValue(value.String())
			} else {
				v.Type = types.StringNull()
			}
			data.Items[i] = v
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin Clone

func (data *CertificateMaps) Clone() CertificateMaps {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data CertificateMaps) toBodyNonBulk(ctx context.Context, state CertificateMaps) string {
	// This is one-by-one update, so only one element to update is expected
	if len(data.Items) > 1 {
		tflog.Error(ctx, "Found more than one element to change. Only one will be changed.")
	}

	// Utilize existing toBody function
	body := data.toBody(ctx, state)

	// Get first element only
	return gjson.Get(body, "0").String()
}

// End of section. //template:end toBodyNonBulk

// Section below is generated&owned by "gen/generator.go". //template:begin findObjectsToBeReplaced

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

// End of section. //template:end clearItemIds
