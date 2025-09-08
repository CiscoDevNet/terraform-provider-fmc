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

type StandardCommunityList struct {
	Id      types.String                   `tfsdk:"id"`
	Domain  types.String                   `tfsdk:"domain"`
	Name    types.String                   `tfsdk:"name"`
	Type    types.String                   `tfsdk:"type"`
	Entries []StandardCommunityListEntries `tfsdk:"entries"`
}

type StandardCommunityListEntries struct {
	Action      types.String `tfsdk:"action"`
	Communities types.String `tfsdk:"communities"`
	Internet    types.Bool   `tfsdk:"internet"`
	NoAdvertise types.Bool   `tfsdk:"no_advertise"`
	NoExport    types.Bool   `tfsdk:"no_export"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data StandardCommunityList) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/standardcommunitylists"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data StandardCommunityList) toBody(ctx context.Context, state StandardCommunityList) string {
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
			itemBody, _ = sjson.Set(itemBody, "type", "Standard")
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "action", item.Action.ValueString())
			}
			if !item.Communities.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "communities", item.Communities.ValueString())
			}
			if !item.Internet.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "internet", item.Internet.ValueBool())
			}
			if !item.NoAdvertise.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "noAdvertise", item.NoAdvertise.ValueBool())
			}
			if !item.NoExport.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "noExport", item.NoExport.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "entries.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *StandardCommunityList) fromBody(ctx context.Context, res gjson.Result) {
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
		data.Entries = make([]StandardCommunityListEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := StandardCommunityListEntries{}
			if value := res.Get("action"); value.Exists() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("communities"); value.Exists() {
				data.Communities = types.StringValue(value.String())
			} else {
				data.Communities = types.StringNull()
			}
			if value := res.Get("internet"); value.Exists() {
				data.Internet = types.BoolValue(value.Bool())
			} else {
				data.Internet = types.BoolNull()
			}
			if value := res.Get("noAdvertise"); value.Exists() {
				data.NoAdvertise = types.BoolValue(value.Bool())
			} else {
				data.NoAdvertise = types.BoolNull()
			}
			if value := res.Get("noExport"); value.Exists() {
				data.NoExport = types.BoolValue(value.Bool())
			} else {
				data.NoExport = types.BoolNull()
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
func (data *StandardCommunityList) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
			data.Entries = append(data.Entries, StandardCommunityListEntries{})
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
		if value := res.Get("communities"); value.Exists() && !data.Communities.IsNull() {
			data.Communities = types.StringValue(value.String())
		} else {
			data.Communities = types.StringNull()
		}
		if value := res.Get("internet"); value.Exists() && !data.Internet.IsNull() {
			data.Internet = types.BoolValue(value.Bool())
		} else {
			data.Internet = types.BoolNull()
		}
		if value := res.Get("noAdvertise"); value.Exists() && !data.NoAdvertise.IsNull() {
			data.NoAdvertise = types.BoolValue(value.Bool())
		} else {
			data.NoAdvertise = types.BoolNull()
		}
		if value := res.Get("noExport"); value.Exists() && !data.NoExport.IsNull() {
			data.NoExport = types.BoolValue(value.Bool())
		} else {
			data.NoExport = types.BoolNull()
		}
		(*parent).Entries[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *StandardCommunityList) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

func (data StandardCommunityList) adjustBody(ctx context.Context, req string) string {

	for i := range len(data.Entries) {
		req, _ = sjson.Set(req, fmt.Sprintf("entries.%d.sequence", i), i+1)
	}
	return req
}

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
