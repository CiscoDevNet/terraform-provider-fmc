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

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FTDPlatformSettingsSyslogRateLimit struct {
	Id                    types.String `tfsdk:"id"`
	Domain                types.String `tfsdk:"domain"`
	FtdPlatformSettingsId types.String `tfsdk:"ftd_platform_settings_id"`
	Type                  types.String `tfsdk:"type"`
	RateLimitType         types.String `tfsdk:"rate_limit_type"`
	RateLimitValue        types.String `tfsdk:"rate_limit_value"`
	NumberOfMessages      types.Int64  `tfsdk:"number_of_messages"`
	Interval              types.Int64  `tfsdk:"interval"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionFTDPlatformSettingsSyslogRateLimit = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsSyslogRateLimit) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/syslog/ratelimits", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsSyslogRateLimit) toBody(ctx context.Context, state FTDPlatformSettingsSyslogRateLimit) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.RateLimitType.IsNull() {
		body, _ = sjson.Set(body, "rateLimitType", data.RateLimitType.ValueString())
	}
	if !data.RateLimitValue.IsNull() {
		body, _ = sjson.Set(body, "value", data.RateLimitValue.ValueString())
	}
	if !data.NumberOfMessages.IsNull() {
		body, _ = sjson.Set(body, "numberOfMessages", data.NumberOfMessages.ValueInt64())
	}
	if !data.Interval.IsNull() {
		body, _ = sjson.Set(body, "intervalInSeconds", data.Interval.ValueInt64())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsSyslogRateLimit) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("rateLimitType"); value.Exists() {
		data.RateLimitType = types.StringValue(value.String())
	} else {
		data.RateLimitType = types.StringNull()
	}
	if value := res.Get("value"); value.Exists() {
		data.RateLimitValue = types.StringValue(value.String())
	} else {
		data.RateLimitValue = types.StringNull()
	}
	if value := res.Get("numberOfMessages"); value.Exists() {
		data.NumberOfMessages = types.Int64Value(value.Int())
	} else {
		data.NumberOfMessages = types.Int64Null()
	}
	if value := res.Get("intervalInSeconds"); value.Exists() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Null()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FTDPlatformSettingsSyslogRateLimit) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("rateLimitType"); value.Exists() && !data.RateLimitType.IsNull() {
		data.RateLimitType = types.StringValue(value.String())
	} else {
		data.RateLimitType = types.StringNull()
	}
	if value := res.Get("value"); value.Exists() && !data.RateLimitValue.IsNull() {
		data.RateLimitValue = types.StringValue(value.String())
	} else {
		data.RateLimitValue = types.StringNull()
	}
	if value := res.Get("numberOfMessages"); value.Exists() && !data.NumberOfMessages.IsNull() {
		data.NumberOfMessages = types.Int64Value(value.Int())
	} else {
		data.NumberOfMessages = types.Int64Null()
	}
	if value := res.Get("intervalInSeconds"); value.Exists() && !data.Interval.IsNull() {
		data.Interval = types.Int64Value(value.Int())
	} else {
		data.Interval = types.Int64Null()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsSyslogRateLimit) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns
