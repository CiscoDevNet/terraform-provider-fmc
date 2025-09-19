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
	"slices"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FTDPlatformSettingsSyslogLoggingDestination struct {
	Id                             types.String                                                   `tfsdk:"id"`
	Domain                         types.String                                                   `tfsdk:"domain"`
	FtdPlatformSettingsId          types.String                                                   `tfsdk:"ftd_platform_settings_id"`
	Type                           types.String                                                   `tfsdk:"type"`
	LoggingDestination             types.String                                                   `tfsdk:"logging_destination"`
	GlobalEventClassFilterCriteria types.String                                                   `tfsdk:"global_event_class_filter_criteria"`
	GlobalEventClassFilterValue    types.String                                                   `tfsdk:"global_event_class_filter_value"`
	EventClassFilters              []FTDPlatformSettingsSyslogLoggingDestinationEventClassFilters `tfsdk:"event_class_filters"`
}

type FTDPlatformSettingsSyslogLoggingDestinationEventClassFilters struct {
	Class    types.String `tfsdk:"class"`
	Severity types.String `tfsdk:"severity"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionFTDPlatformSettingsSyslogLoggingDestination = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsSyslogLoggingDestination) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/syslog/loggingdestinations", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsSyslogLoggingDestination) toBody(ctx context.Context, state FTDPlatformSettingsSyslogLoggingDestination) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.LoggingDestination.IsNull() {
		body, _ = sjson.Set(body, "loggingDestination", data.LoggingDestination.ValueString())
	}
	if !data.GlobalEventClassFilterCriteria.IsNull() {
		body, _ = sjson.Set(body, "allEventConfig.filterCriteria", data.GlobalEventClassFilterCriteria.ValueString())
	}
	if !data.GlobalEventClassFilterValue.IsNull() {
		body, _ = sjson.Set(body, "allEventConfig.value", data.GlobalEventClassFilterValue.ValueString())
	}
	if len(data.EventClassFilters) > 0 {
		body, _ = sjson.Set(body, "specificEventConfig", []any{})
		for _, item := range data.EventClassFilters {
			itemBody := ""
			if !item.Class.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "class", item.Class.ValueString())
			}
			if !item.Severity.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "severity", item.Severity.ValueString())
			}
			body, _ = sjson.SetRaw(body, "specificEventConfig.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsSyslogLoggingDestination) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("loggingDestination"); value.Exists() {
		data.LoggingDestination = types.StringValue(value.String())
	} else {
		data.LoggingDestination = types.StringValue("INTERNAL_BUFFER")
	}
	if value := res.Get("allEventConfig.filterCriteria"); value.Exists() {
		data.GlobalEventClassFilterCriteria = types.StringValue(value.String())
	} else {
		data.GlobalEventClassFilterCriteria = types.StringNull()
	}
	if value := res.Get("allEventConfig.value"); value.Exists() {
		data.GlobalEventClassFilterValue = types.StringValue(value.String())
	} else {
		data.GlobalEventClassFilterValue = types.StringNull()
	}
	if value := res.Get("specificEventConfig"); value.Exists() {
		data.EventClassFilters = make([]FTDPlatformSettingsSyslogLoggingDestinationEventClassFilters, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsSyslogLoggingDestinationEventClassFilters{}
			if value := res.Get("class"); value.Exists() {
				data.Class = types.StringValue(value.String())
			} else {
				data.Class = types.StringNull()
			}
			if value := res.Get("severity"); value.Exists() {
				data.Severity = types.StringValue(value.String())
			} else {
				data.Severity = types.StringNull()
			}
			(*parent).EventClassFilters = append((*parent).EventClassFilters, data)
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
func (data *FTDPlatformSettingsSyslogLoggingDestination) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("loggingDestination"); value.Exists() && !data.LoggingDestination.IsNull() {
		data.LoggingDestination = types.StringValue(value.String())
	} else if data.LoggingDestination.ValueString() != "INTERNAL_BUFFER" {
		data.LoggingDestination = types.StringNull()
	}
	if value := res.Get("allEventConfig.filterCriteria"); value.Exists() && !data.GlobalEventClassFilterCriteria.IsNull() {
		data.GlobalEventClassFilterCriteria = types.StringValue(value.String())
	} else {
		data.GlobalEventClassFilterCriteria = types.StringNull()
	}
	if value := res.Get("allEventConfig.value"); value.Exists() && !data.GlobalEventClassFilterValue.IsNull() {
		data.GlobalEventClassFilterValue = types.StringValue(value.String())
	} else {
		data.GlobalEventClassFilterValue = types.StringNull()
	}
	for i := 0; i < len(data.EventClassFilters); i++ {
		keys := [...]string{"class", "severity"}
		keyValues := [...]string{data.EventClassFilters[i].Class.ValueString(), data.EventClassFilters[i].Severity.ValueString()}

		parent := &data
		data := (*parent).EventClassFilters[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("specificEventConfig").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing EventClassFilters[%d] = %+v",
				i,
				(*parent).EventClassFilters[i],
			))
			(*parent).EventClassFilters = slices.Delete((*parent).EventClassFilters, i, i+1)
			i--

			continue
		}
		if value := res.Get("class"); value.Exists() && !data.Class.IsNull() {
			data.Class = types.StringValue(value.String())
		} else {
			data.Class = types.StringNull()
		}
		if value := res.Get("severity"); value.Exists() && !data.Severity.IsNull() {
			data.Severity = types.StringValue(value.String())
		} else {
			data.Severity = types.StringNull()
		}
		(*parent).EventClassFilters[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsSyslogLoggingDestination) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns
