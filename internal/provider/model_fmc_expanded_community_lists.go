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

type ExpandedCommunityLists struct {
	Id     types.String                           `tfsdk:"id"`
	Domain types.String                           `tfsdk:"domain"`
	Items  map[string]ExpandedCommunityListsItems `tfsdk:"items"`
}

type ExpandedCommunityListsItems struct {
	Id      types.String                         `tfsdk:"id"`
	Type    types.String                         `tfsdk:"type"`
	Entries []ExpandedCommunityListsItemsEntries `tfsdk:"entries"`
}

type ExpandedCommunityListsItemsEntries struct {
	Action            types.String `tfsdk:"action"`
	RegularExpression types.String `tfsdk:"regular_expression"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionBulkCreateExpandedCommunityLists = version.Must(version.NewVersion("999"))
var minFMCVersionBulkDeleteExpandedCommunityLists = version.Must(version.NewVersion("999"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ExpandedCommunityLists) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/expandedcommunitylists"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ExpandedCommunityLists) toBody(ctx context.Context, state ExpandedCommunityLists) string {
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
			if len(item.Entries) > 0 {
				itemBody, _ = sjson.Set(itemBody, "entries", []interface{}{})
				for _, childItem := range item.Entries {
					itemChildBody := ""
					itemChildBody, _ = sjson.Set(itemChildBody, "type", "Expanded")
					if !childItem.Action.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "action", childItem.Action.ValueString())
					}
					if !childItem.RegularExpression.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "regularExpression", childItem.RegularExpression.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "entries.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "items.-1", itemBody)
		}
	}
	return gjson.Get(body, "items").String()
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ExpandedCommunityLists) fromBody(ctx context.Context, res gjson.Result) {
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
		if value := res.Get("entries"); value.Exists() {
			data.Entries = make([]ExpandedCommunityListsItemsEntries, 0)
			value.ForEach(func(k, res gjson.Result) bool {
				parent := &data
				data := ExpandedCommunityListsItemsEntries{}
				if value := res.Get("action"); value.Exists() {
					data.Action = types.StringValue(value.String())
				} else {
					data.Action = types.StringNull()
				}
				if value := res.Get("regularExpression"); value.Exists() {
					data.RegularExpression = types.StringValue(value.String())
				} else {
					data.RegularExpression = types.StringNull()
				}
				(*parent).Entries = append((*parent).Entries, data)
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
func (data *ExpandedCommunityLists) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
		{
			l := len(res.Get("entries").Array())
			tflog.Debug(ctx, fmt.Sprintf("entries array resizing from %d to %d", len(data.Entries), l))
			for i := len(data.Entries); i < l; i++ {
				data.Entries = append(data.Entries, ExpandedCommunityListsItemsEntries{})
			}
			if len(data.Entries) > l {
				data.Entries = data.Entries[:l]
			}
		}
		for i := range data.Entries {
			parent := &data
			data := (*parent).Entries[i]
			parentRes := &res
			res := parentRes.Get(fmt.Sprintf("entries.%d", i))
			if value := res.Get("action"); value.Exists() && !data.Action.IsNull() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("regularExpression"); value.Exists() && !data.RegularExpression.IsNull() {
				data.RegularExpression = types.StringValue(value.String())
			} else {
				data.RegularExpression = types.StringNull()
			}
			(*parent).Entries[i] = data
		}
		(*parent).Items[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ExpandedCommunityLists) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

func (data *ExpandedCommunityLists) Clone() ExpandedCommunityLists {
	ret := *data
	ret.Items = maps.Clone(data.Items)

	return ret
}

// End of section. //template:end Clone

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyNonBulk

// Updates done one-by-one require different API body
func (data ExpandedCommunityLists) toBodyNonBulk(ctx context.Context, state ExpandedCommunityLists) string {
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

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// End of section. //template:end toBodyPutDelete

func (data ExpandedCommunityLists) adjustBody(ctx context.Context, req string) string {

	for i := range len(gjson.Get(req, "entries").Array()) {
		req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.sequence", i), i+1)
	}
	return req
}

func (data ExpandedCommunityLists) adjustBodyBulk(ctx context.Context, req string) string {
	// To be updated once FMC supports bulk operations for Expanded Community Lists
	return req
}
