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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type IKEv2Policies struct {
	Id     types.String                  `tfsdk:"id"`
	Domain types.String                  `tfsdk:"domain"`
	Items  map[string]IKEv2PoliciesItems `tfsdk:"items"`
}

type IKEv2PoliciesItems struct {
	Id                   types.String `tfsdk:"id"`
	Description          types.String `tfsdk:"description"`
	Type                 types.String `tfsdk:"type"`
	Priority             types.Int64  `tfsdk:"priority"`
	Lifetime             types.Int64  `tfsdk:"lifetime"`
	IntegrityAlgorithms  types.Set    `tfsdk:"integrity_algorithms"`
	EncryptionAlgorithms types.Set    `tfsdk:"encryption_algorithms"`
	PrfAlgorithms        types.Set    `tfsdk:"prf_algorithms"`
	DhGroups             types.Set    `tfsdk:"dh_groups"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkCreateIKEv2Policies = version.Must(version.NewVersion("999"))
var minFMCVersionBulkDeleteIKEv2Policies = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data IKEv2Policies) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/ikev2policies"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data IKEv2Policies) toBody(ctx context.Context, state IKEv2Policies) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if len(data.Items) > 0 {
		body, _ = sjson.Set(body, "items", []interface{}{})
		for key, item := range data.Items {
			itemBody, _ := sjson.Set("{}", "name", key)
			if !item.Id.IsNull() && !item.Id.IsUnknown() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			itemBody, _ = sjson.Set(itemBody, "type", "IKEv2Policy")
			if !item.Priority.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "priority", item.Priority.ValueInt64())
			}
			if !item.Lifetime.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "lifetimeInSeconds", item.Lifetime.ValueInt64())
			}
			if !item.IntegrityAlgorithms.IsNull() {
				var values []string
				item.IntegrityAlgorithms.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "integrityAlgorithms", values)
			}
			if !item.EncryptionAlgorithms.IsNull() {
				var values []string
				item.EncryptionAlgorithms.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "encryptionAlgorithms", values)
			}
			if !item.PrfAlgorithms.IsNull() {
				var values []string
				item.PrfAlgorithms.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "prfIntegrityAlgorithms", values)
			}
			if !item.DhGroups.IsNull() {
				var values []string
				item.DhGroups.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "diffieHellmanGroups", values)
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *IKEv2Policies) fromBody(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("description"); value.Exists() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("priority"); value.Exists() {
			data.Priority = types.Int64Value(value.Int())
		} else {
			data.Priority = types.Int64Null()
		}
		if value := res.Get("lifetimeInSeconds"); value.Exists() {
			data.Lifetime = types.Int64Value(value.Int())
		} else {
			data.Lifetime = types.Int64Null()
		}
		if value := res.Get("integrityAlgorithms"); value.Exists() {
			data.IntegrityAlgorithms = helpers.GetStringSet(value.Array())
		} else {
			data.IntegrityAlgorithms = types.SetNull(types.StringType)
		}
		if value := res.Get("encryptionAlgorithms"); value.Exists() {
			data.EncryptionAlgorithms = helpers.GetStringSet(value.Array())
		} else {
			data.EncryptionAlgorithms = types.SetNull(types.StringType)
		}
		if value := res.Get("prfIntegrityAlgorithms"); value.Exists() {
			data.PrfAlgorithms = helpers.GetStringSet(value.Array())
		} else {
			data.PrfAlgorithms = types.SetNull(types.StringType)
		}
		if value := res.Get("diffieHellmanGroups"); value.Exists() {
			data.DhGroups = helpers.GetStringSet(value.Array())
		} else {
			data.DhGroups = types.SetNull(types.StringType)
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
func (data *IKEv2Policies) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
			data.Priority = types.Int64Value(value.Int())
		} else {
			data.Priority = types.Int64Null()
		}
		if value := res.Get("lifetimeInSeconds"); value.Exists() && !data.Lifetime.IsNull() {
			data.Lifetime = types.Int64Value(value.Int())
		} else {
			data.Lifetime = types.Int64Null()
		}
		if value := res.Get("integrityAlgorithms"); value.Exists() && !data.IntegrityAlgorithms.IsNull() {
			data.IntegrityAlgorithms = helpers.GetStringSet(value.Array())
		} else {
			data.IntegrityAlgorithms = types.SetNull(types.StringType)
		}
		if value := res.Get("encryptionAlgorithms"); value.Exists() && !data.EncryptionAlgorithms.IsNull() {
			data.EncryptionAlgorithms = helpers.GetStringSet(value.Array())
		} else {
			data.EncryptionAlgorithms = types.SetNull(types.StringType)
		}
		if value := res.Get("prfIntegrityAlgorithms"); value.Exists() && !data.PrfAlgorithms.IsNull() {
			data.PrfAlgorithms = helpers.GetStringSet(value.Array())
		} else {
			data.PrfAlgorithms = types.SetNull(types.StringType)
		}
		if value := res.Get("diffieHellmanGroups"); value.Exists() && !data.DhGroups.IsNull() {
			data.DhGroups = helpers.GetStringSet(value.Array())
		} else {
			data.DhGroups = types.SetNull(types.StringType)
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *IKEv2Policies) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

func (data *IKEv2Policies) Clone() IKEv2Policies {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data IKEv2Policies) toBodyNonBulk(ctx context.Context, state IKEv2Policies) string {
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
