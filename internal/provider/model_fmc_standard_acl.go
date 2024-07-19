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

type StandardACL struct {
	Id          types.String         `tfsdk:"id"`
	Domain      types.String         `tfsdk:"domain"`
	Name        types.String         `tfsdk:"name"`
	Description types.String         `tfsdk:"description"`
	Entries     []StandardACLEntries `tfsdk:"entries"`
}

type StandardACLEntries struct {
	Action   types.String                 `tfsdk:"action"`
	Objects  []StandardACLEntriesObjects  `tfsdk:"objects"`
	Literals []StandardACLEntriesLiterals `tfsdk:"literals"`
}

type StandardACLEntriesObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}
type StandardACLEntriesLiterals struct {
	Value types.String `tfsdk:"value"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data StandardACL) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/standardaccesslists"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data StandardACL) toBody(ctx context.Context, state StandardACL) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []interface{}{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if len(item.Objects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "networks.objects", []interface{}{})
				for _, childItem := range item.Objects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "networks.objects.-1", itemChildBody)
				}
			}
			if len(item.Literals) > 0 {
				itemBody, _ = sjson.Set(itemBody, "networks.literals", []interface{}{})
				for _, childItem := range item.Literals {
					itemChildBody := ""
					if !childItem.Value.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "value", childItem.Value.ValueString())
					}
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "AnyNonEmptyString")
					itemBody, _ = sjson.SetRaw(itemBody, "networks.literals.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *StandardACL) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]StandardACLEntries, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := StandardACLEntries{}
			if cValue := v.Get("action"); cValue.Exists() {
				item.Action = types.StringValue(cValue.String())
			} else {
				item.Action = types.StringNull()
			}
			if cValue := v.Get("networks.objects"); cValue.Exists() {
				item.Objects = make([]StandardACLEntriesObjects, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := StandardACLEntriesObjects{}
					if ccValue := cv.Get("id"); ccValue.Exists() {
						cItem.Id = types.StringValue(ccValue.String())
					} else {
						cItem.Id = types.StringNull()
					}
					if ccValue := cv.Get("type"); ccValue.Exists() {
						cItem.Type = types.StringValue(ccValue.String())
					} else {
						cItem.Type = types.StringNull()
					}
					item.Objects = append(item.Objects, cItem)
					return true
				})
			}
			if cValue := v.Get("networks.literals"); cValue.Exists() {
				item.Literals = make([]StandardACLEntriesLiterals, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := StandardACLEntriesLiterals{}
					if ccValue := cv.Get("value"); ccValue.Exists() {
						cItem.Value = types.StringValue(ccValue.String())
					} else {
						cItem.Value = types.StringNull()
					}
					item.Literals = append(item.Literals, cItem)
					return true
				})
			}
			data.Entries = append(data.Entries, item)
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
func (data *StandardACL) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	{
		l := len(res.Get("entries").Array())
		tflog.Debug(ctx, fmt.Sprintf("entries array resizing from %d to %d", len(data.Entries), l))
		for i := len(data.Entries); i < l; i++ {
			data.Entries = append(data.Entries, StandardACLEntries{})
		}
		if len(data.Entries) > l {
			data.Entries = data.Entries[:l]
		}
	}
	for i := range data.Entries {
		r := res.Get(fmt.Sprintf("entries.%d", i))
		if value := r.Get("action"); value.Exists() && !data.Entries[i].Action.IsNull() {
			data.Entries[i].Action = types.StringValue(value.String())
		} else {
			data.Entries[i].Action = types.StringNull()
		}
		for ci := 0; ci < len(data.Entries[i].Objects); ci++ {
			keys := [...]string{"id"}
			keyValues := [...]string{data.Entries[i].Objects[ci].Id.ValueString()}

			var cr gjson.Result
			r.Get("networks.objects").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.Entries[i].Objects[%d] = %+v",
					ci,
					data.Entries[i].Objects[ci],
				))
				data.Entries[i].Objects = slices.Delete(data.Entries[i].Objects, ci, ci+1)
				ci--

				continue
			}
			if value := cr.Get("id"); value.Exists() && !data.Entries[i].Objects[ci].Id.IsNull() {
				data.Entries[i].Objects[ci].Id = types.StringValue(value.String())
			} else {
				data.Entries[i].Objects[ci].Id = types.StringNull()
			}
			if value := cr.Get("type"); value.Exists() && !data.Entries[i].Objects[ci].Type.IsNull() {
				data.Entries[i].Objects[ci].Type = types.StringValue(value.String())
			} else {
				data.Entries[i].Objects[ci].Type = types.StringNull()
			}
		}
		for ci := 0; ci < len(data.Entries[i].Literals); ci++ {
			keys := [...]string{"value"}
			keyValues := [...]string{data.Entries[i].Literals[ci].Value.ValueString()}

			var cr gjson.Result
			r.Get("networks.literals").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)
			if !cr.Exists() {
				tflog.Debug(ctx, fmt.Sprintf("removing data.Entries[i].Literals[%d] = %+v",
					ci,
					data.Entries[i].Literals[ci],
				))
				data.Entries[i].Literals = slices.Delete(data.Entries[i].Literals, ci, ci+1)
				ci--

				continue
			}
			if value := cr.Get("value"); value.Exists() && !data.Entries[i].Literals[ci].Value.IsNull() {
				data.Entries[i].Literals[ci].Value = types.StringValue(value.String())
			} else {
				data.Entries[i].Literals[ci].Value = types.StringNull()
			}
		}
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *StandardACL) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
}

// End of section. //template:end fromBodyUnknowns
