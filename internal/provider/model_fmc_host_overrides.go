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
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type HostOverrides struct {
	Id         types.String             `tfsdk:"id"`
	Domain     types.String             `tfsdk:"domain"`
	ParentName types.String             `tfsdk:"parent_name"`
	ParentId   types.String             `tfsdk:"parent_id"`
	Overrides  []HostOverridesOverrides `tfsdk:"overrides"`
}

type HostOverridesOverrides struct {
	TargetId    types.String `tfsdk:"target_id"`
	TargetType  types.String `tfsdk:"target_type"`
	Description types.String `tfsdk:"description"`
	Ip          types.String `tfsdk:"ip"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data HostOverrides) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/hosts"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data HostOverrides) toBody(ctx context.Context, state HostOverrides) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.ParentName.IsNull() {
		body, _ = sjson.Set(body, "name", data.ParentName.ValueString())
	}
	if !data.ParentId.IsNull() {
		body, _ = sjson.Set(body, "overrides.parent.id", data.ParentId.ValueString())
	}
	if len(data.Overrides) > 0 {
		body, _ = sjson.Set(body, "dummy_overrides", []any{})
		for _, item := range data.Overrides {
			itemBody := ""
			if !item.TargetId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overrides.target.id", item.TargetId.ValueString())
			}
			if !item.TargetType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overrides.target.type", item.TargetType.ValueString())
			}
			if !item.Description.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "description", item.Description.ValueString())
			}
			if !item.Ip.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "value", item.Ip.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dummy_overrides.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *HostOverrides) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.ParentName = types.StringValue(value.String())
	} else {
		data.ParentName = types.StringNull()
	}
	if value := res.Get("overrides.parent.id"); value.Exists() {
		data.ParentId = types.StringValue(value.String())
	} else {
		data.ParentId = types.StringNull()
	}
	if value := res.Get("dummy_overrides"); value.Exists() {
		data.Overrides = make([]HostOverridesOverrides, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := HostOverridesOverrides{}
			if value := res.Get("overrides.target.id"); value.Exists() {
				data.TargetId = types.StringValue(value.String())
			} else {
				data.TargetId = types.StringNull()
			}
			if value := res.Get("overrides.target.type"); value.Exists() {
				data.TargetType = types.StringValue(value.String())
			} else {
				data.TargetType = types.StringNull()
			}
			if value := res.Get("description"); value.Exists() {
				data.Description = types.StringValue(value.String())
			} else {
				data.Description = types.StringNull()
			}
			if value := res.Get("value"); value.Exists() {
				data.Ip = types.StringValue(value.String())
			} else {
				data.Ip = types.StringNull()
			}
			(*parent).Overrides = append((*parent).Overrides, data)
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
func (data *HostOverrides) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.ParentName.IsNull() {
		data.ParentName = types.StringValue(value.String())
	} else {
		data.ParentName = types.StringNull()
	}
	if value := res.Get("overrides.parent.id"); value.Exists() && !data.ParentId.IsNull() {
		data.ParentId = types.StringValue(value.String())
	} else {
		data.ParentId = types.StringNull()
	}
	for i := 0; i < len(data.Overrides); i++ {
		keys := [...]string{"overrides.target.id"}
		keyValues := [...]string{data.Overrides[i].TargetId.ValueString()}

		parent := &data
		data := (*parent).Overrides[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dummy_overrides").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Overrides[%d] = %+v",
				i,
				(*parent).Overrides[i],
			))
			(*parent).Overrides = slices.Delete((*parent).Overrides, i, i+1)
			i--

			continue
		}
		if value := res.Get("overrides.target.id"); value.Exists() && !data.TargetId.IsNull() {
			data.TargetId = types.StringValue(value.String())
		} else {
			data.TargetId = types.StringNull()
		}
		if value := res.Get("overrides.target.type"); value.Exists() && !data.TargetType.IsNull() {
			data.TargetType = types.StringValue(value.String())
		} else {
			data.TargetType = types.StringNull()
		}
		if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
			data.Description = types.StringValue(value.String())
		} else {
			data.Description = types.StringNull()
		}
		if value := res.Get("value"); value.Exists() && !data.Ip.IsNull() {
			data.Ip = types.StringValue(value.String())
		} else {
			data.Ip = types.StringNull()
		}
		(*parent).Overrides[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *HostOverrides) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

func (data HostOverrides) toBodyOverrides(ctx context.Context, state HostOverrides) string {
	body := data.toBody(ctx, state)

	r := gjson.Get(body, "dummy_overrides")
	if !r.Exists() {
		return body
	}

	r.ForEach(func(key, value gjson.Result) bool {
		updated := value.Raw
		updated, _ = sjson.Set(updated, "name", data.ParentName.ValueString())
		updated, _ = sjson.Set(updated, "overrides.parent.id", data.ParentId.ValueString())
		body, _ = sjson.SetRaw(body, "dummy_overrides."+key.String(), updated)
		return true
	})

	return gjson.Get(body, "dummy_overrides").String()
}

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// synthesizeOverrides transforms the API response
// (which uses real field names and contains injected parent fields) back into the dummy_* structure
func (data HostOverrides) synthesizeOverrides(ctx context.Context, res gjson.Result) gjson.Result {
	body := ""

	// Map top-level response fields to dummy attributes
	if value := res.Get("items.0.name"); value.Exists() {
		body, _ = sjson.Set(body, "name", value.Value())
	}
	if value := res.Get("items.0.overrides.parent.id"); value.Exists() {
		body, _ = sjson.Set(body, "overrides.parent.id", value.Value())
	}

	body, _ = sjson.SetRaw(body, "dummy_overrides", res.Get("items").Raw)

	return gjson.Parse(body)
}

// End of section. //template:end synthesizeOverrides
