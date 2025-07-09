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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type SecurityIntelligenceURLFeed struct {
	Id              types.String `tfsdk:"id"`
	Domain          types.String `tfsdk:"domain"`
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	FeedUrl         types.String `tfsdk:"feed_url"`
	ChecksumUrl     types.String `tfsdk:"checksum_url"`
	UpdateFrequency types.Int64  `tfsdk:"update_frequency"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data SecurityIntelligenceURLFeed) getPath() string {
	return "/api/fmc_config/v1/domain/{DOMAIN_UUID}/object/siurlfeeds"
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data SecurityIntelligenceURLFeed) toBody(ctx context.Context, state SecurityIntelligenceURLFeed) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.FeedUrl.IsNull() {
		body, _ = sjson.Set(body, "feedURL", data.FeedUrl.ValueString())
	}
	if !data.ChecksumUrl.IsNull() {
		body, _ = sjson.Set(body, "checksumURL", data.ChecksumUrl.ValueString())
	}
	if !data.UpdateFrequency.IsNull() {
		body, _ = sjson.Set(body, "updateFrequency", data.UpdateFrequency.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *SecurityIntelligenceURLFeed) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("feedURL"); value.Exists() {
		data.FeedUrl = types.StringValue(value.String())
	} else {
		data.FeedUrl = types.StringNull()
	}
	if value := res.Get("checksumURL"); value.Exists() {
		data.ChecksumUrl = types.StringValue(value.String())
	} else {
		data.ChecksumUrl = types.StringNull()
	}
	if value := res.Get("updateFrequency"); value.Exists() {
		data.UpdateFrequency = types.Int64Value(value.Int())
	} else {
		data.UpdateFrequency = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *SecurityIntelligenceURLFeed) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("feedURL"); value.Exists() && !data.FeedUrl.IsNull() {
		data.FeedUrl = types.StringValue(value.String())
	} else {
		data.FeedUrl = types.StringNull()
	}
	if value := res.Get("checksumURL"); value.Exists() && !data.ChecksumUrl.IsNull() {
		data.ChecksumUrl = types.StringValue(value.String())
	} else {
		data.ChecksumUrl = types.StringNull()
	}
	if value := res.Get("updateFrequency"); value.Exists() && !data.UpdateFrequency.IsNull() {
		data.UpdateFrequency = types.Int64Value(value.Int())
	} else {
		data.UpdateFrequency = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *SecurityIntelligenceURLFeed) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
