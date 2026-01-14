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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type IPv6PrefixList struct {
	Id      types.String            `tfsdk:"id"`
	Domain  types.String            `tfsdk:"domain"`
	Name    types.String            `tfsdk:"name"`
	Type    types.String            `tfsdk:"type"`
	Entries []IPv6PrefixListEntries `tfsdk:"entries"`
}

type IPv6PrefixListEntries struct {
	Action          types.String `tfsdk:"action"`
	Prefix          types.String `tfsdk:"prefix"`
	MinPrefixLength types.Int64  `tfsdk:"min_prefix_length"`
	MaxPrefixLength types.Int64  `tfsdk:"max_prefix_length"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data IPv6PrefixList) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/ipv6prefixlists"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data IPv6PrefixList) toBody(ctx context.Context, state IPv6PrefixList) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if len(data.Entries) > 0 {
		body, _ = sjson.Set(body, "entries", []any{})
		for _, item := range data.Entries {
			itemBody := ""
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipAddress", item.Prefix.ValueString())
			}
			if !item.MinPrefixLength.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "minPrefixLenth", item.MinPrefixLength.ValueInt64())
			}
			if !item.MaxPrefixLength.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "maxPrefixLength", item.MaxPrefixLength.ValueInt64())
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *IPv6PrefixList) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("entries"); value.Exists() {
		data.Entries = make([]IPv6PrefixListEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := IPv6PrefixListEntries{}
			if value := res.Get("action"); value.Exists() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("ipAddress"); value.Exists() {
				data.Prefix = types.StringValue(value.String())
			} else {
				data.Prefix = types.StringNull()
			}
			if value := res.Get("minPrefixLenth"); value.Exists() {
				data.MinPrefixLength = types.Int64Value(value.Int())
			} else {
				data.MinPrefixLength = types.Int64Null()
			}
			if value := res.Get("maxPrefixLength"); value.Exists() {
				data.MaxPrefixLength = types.Int64Value(value.Int())
			} else {
				data.MaxPrefixLength = types.Int64Null()
			}
			(*parent).Entries = append((*parent).Entries, data)
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
func (data *IPv6PrefixList) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	{
		l := len(res.Get("entries").Array())
		tflog.Debug(ctx, fmt.Sprintf("entries array resizing from %d to %d", len(data.Entries), l))
		for i := len(data.Entries); i < l; i++ {
			data.Entries = append(data.Entries, IPv6PrefixListEntries{})
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
		if value := res.Get("ipAddress"); value.Exists() && !data.Prefix.IsNull() {
			data.Prefix = types.StringValue(value.String())
		} else {
			data.Prefix = types.StringNull()
		}
		if value := res.Get("minPrefixLenth"); value.Exists() && !data.MinPrefixLength.IsNull() {
			data.MinPrefixLength = types.Int64Value(value.Int())
		} else {
			data.MinPrefixLength = types.Int64Null()
		}
		if value := res.Get("maxPrefixLength"); value.Exists() && !data.MaxPrefixLength.IsNull() {
			data.MaxPrefixLength = types.Int64Value(value.Int())
		} else {
			data.MaxPrefixLength = types.Int64Null()
		}
		(*parent).Entries[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *IPv6PrefixList) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

func (data IPv6PrefixList) adjustBody(ctx context.Context, req string) string {

	// Add sequence numbers to the entities
	for i := range len(data.Entries) {
		req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.sequence", i), i+1)
	}

	return req
}

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
