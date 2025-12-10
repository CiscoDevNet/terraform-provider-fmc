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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type BFDTemplates struct {
	Id     types.String                 `tfsdk:"id"`
	Domain types.String                 `tfsdk:"domain"`
	Items  map[string]BFDTemplatesItems `tfsdk:"items"`
}

type BFDTemplatesItems struct {
	Id                               types.String `tfsdk:"id"`
	Type                             types.String `tfsdk:"type"`
	HopType                          types.String `tfsdk:"hop_type"`
	Echo                             types.String `tfsdk:"echo"`
	IntervalType                     types.String `tfsdk:"interval_type"`
	Multiplier                       types.Int64  `tfsdk:"multiplier"`
	MinimumTransmit                  types.Int64  `tfsdk:"minimum_transmit"`
	MinimumReceive                   types.Int64  `tfsdk:"minimum_receive"`
	AuthenticationType               types.String `tfsdk:"authentication_type"`
	AuthenticationPassword           types.String `tfsdk:"authentication_password"`
	AuthenticationPasswordEncryption types.String `tfsdk:"authentication_password_encryption"`
	AuthenticationKeyId              types.Int64  `tfsdk:"authentication_key_id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBFDTemplates = version.Must(version.NewVersion("7.4"))
var minFMCVersionBulkCreateBFDTemplates = version.Must(version.NewVersion("999"))
var minFMCVersionBulkDeleteBFDTemplates = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data BFDTemplates) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/bfdtemplates"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data BFDTemplates) toBody(ctx context.Context, state BFDTemplates) string {
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
			itemBody, _ = sjson.Set(itemBody, "type", "BFDTemplate")
			if !item.HopType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "hopType", item.HopType.ValueString())
			}
			if !item.Echo.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "echo", item.Echo.ValueString())
			}
			if !item.IntervalType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "txRxInterval", item.IntervalType.ValueString())
			}
			if !item.Multiplier.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "txRxMultiplier", item.Multiplier.ValueInt64())
			}
			if !item.MinimumTransmit.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "minTransmit", item.MinimumTransmit.ValueInt64())
			}
			if !item.MinimumReceive.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "minReceive", item.MinimumReceive.ValueInt64())
			}
			if !item.AuthenticationType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authentication.authType", item.AuthenticationType.ValueString())
			}
			if !item.AuthenticationPassword.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authentication.authKey", item.AuthenticationPassword.ValueString())
			}
			if !item.AuthenticationPasswordEncryption.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authentication.pwdEncryption", item.AuthenticationPasswordEncryption.ValueString())
			}
			if !item.AuthenticationKeyId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "authentication.authKeyId", item.AuthenticationKeyId.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *BFDTemplates) fromBody(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("hopType"); value.Exists() {
			data.HopType = types.StringValue(value.String())
		} else {
			data.HopType = types.StringNull()
		}
		if value := res.Get("echo"); value.Exists() {
			data.Echo = types.StringValue(value.String())
		} else {
			data.Echo = types.StringNull()
		}
		if value := res.Get("txRxInterval"); value.Exists() {
			data.IntervalType = types.StringValue(value.String())
		} else {
			data.IntervalType = types.StringNull()
		}
		if value := res.Get("txRxMultiplier"); value.Exists() {
			data.Multiplier = types.Int64Value(value.Int())
		} else {
			data.Multiplier = types.Int64Null()
		}
		if value := res.Get("minTransmit"); value.Exists() {
			data.MinimumTransmit = types.Int64Value(value.Int())
		} else {
			data.MinimumTransmit = types.Int64Null()
		}
		if value := res.Get("minReceive"); value.Exists() {
			data.MinimumReceive = types.Int64Value(value.Int())
		} else {
			data.MinimumReceive = types.Int64Null()
		}
		if value := res.Get("authentication.authType"); value.Exists() {
			data.AuthenticationType = types.StringValue(value.String())
		} else {
			data.AuthenticationType = types.StringNull()
		}
		if value := res.Get("authentication.pwdEncryption"); value.Exists() {
			data.AuthenticationPasswordEncryption = types.StringValue(value.String())
		} else {
			data.AuthenticationPasswordEncryption = types.StringNull()
		}
		if value := res.Get("authentication.authKeyId"); value.Exists() {
			data.AuthenticationKeyId = types.Int64Value(value.Int())
		} else {
			data.AuthenticationKeyId = types.Int64Null()
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
func (data *BFDTemplates) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("hopType"); value.Exists() && !data.HopType.IsNull() {
			data.HopType = types.StringValue(value.String())
		} else {
			data.HopType = types.StringNull()
		}
		if value := res.Get("echo"); value.Exists() && !data.Echo.IsNull() {
			data.Echo = types.StringValue(value.String())
		} else {
			data.Echo = types.StringNull()
		}
		if value := res.Get("txRxInterval"); value.Exists() && !data.IntervalType.IsNull() {
			data.IntervalType = types.StringValue(value.String())
		} else {
			data.IntervalType = types.StringNull()
		}
		if value := res.Get("txRxMultiplier"); value.Exists() && !data.Multiplier.IsNull() {
			data.Multiplier = types.Int64Value(value.Int())
		} else {
			data.Multiplier = types.Int64Null()
		}
		if value := res.Get("minTransmit"); value.Exists() && !data.MinimumTransmit.IsNull() {
			data.MinimumTransmit = types.Int64Value(value.Int())
		} else {
			data.MinimumTransmit = types.Int64Null()
		}
		if value := res.Get("minReceive"); value.Exists() && !data.MinimumReceive.IsNull() {
			data.MinimumReceive = types.Int64Value(value.Int())
		} else {
			data.MinimumReceive = types.Int64Null()
		}
		if value := res.Get("authentication.authType"); value.Exists() && !data.AuthenticationType.IsNull() {
			data.AuthenticationType = types.StringValue(value.String())
		} else {
			data.AuthenticationType = types.StringNull()
		}
		if value := res.Get("authentication.pwdEncryption"); value.Exists() && !data.AuthenticationPasswordEncryption.IsNull() {
			data.AuthenticationPasswordEncryption = types.StringValue(value.String())
		} else {
			data.AuthenticationPasswordEncryption = types.StringNull()
		}
		if value := res.Get("authentication.authKeyId"); value.Exists() && !data.AuthenticationKeyId.IsNull() {
			data.AuthenticationKeyId = types.Int64Value(value.Int())
		} else {
			data.AuthenticationKeyId = types.Int64Null()
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *BFDTemplates) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

func (data *BFDTemplates) Clone() BFDTemplates {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data BFDTemplates) toBodyNonBulk(ctx context.Context, state BFDTemplates) string {
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

// Check if single object within bulk requires replace due to `requires_replace`
// Since here we assume object has changed, it must be present in both state and plan (data)
func (data BFDTemplates) findObjectsToBeReplaced(ctx context.Context, state BFDTemplates) BFDTemplates {
	// Prepare empty object to be filled in with objects that require replace
	var toBeReplaced BFDTemplates
	toBeReplaced.Items = make(map[string]BFDTemplatesItems)

	// Iterate over all objects in plan
	for key, item := range data.Items {
		// Check if object is present in state
		if _, ok := state.Items[key]; !ok {
			// Object is not present in state, hence it's not a candidate for replace
			continue
		}

		// Check if any field marked as `requires_replace` has changed
		if item.HopType != state.Items[key].HopType {
			toBeReplaced.Items[key] = item
			continue
		}
	}

	return toBeReplaced
}

// End of section. //template:end findObjectsToBeReplaced

// Section below is generated&owned by "gen/generator.go". //template:begin clearItemIds

func (data *BFDTemplates) clearItemsIds(ctx context.Context) {
	for key, value := range data.Items {
		tmp := value
		tmp.Id = types.StringNull()
		data.Items[key] = tmp
	}
}

// End of section. //template:end clearItemIds

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// End of section. //template:end toBodyPutDelete

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
