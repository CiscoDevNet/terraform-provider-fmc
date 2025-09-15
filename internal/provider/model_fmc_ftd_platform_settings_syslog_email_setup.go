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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type FTDPlatformSettingsSyslogEmailSetup struct {
	Id                    types.String                                      `tfsdk:"id"`
	Domain                types.String                                      `tfsdk:"domain"`
	FtdPlatformSettingsId types.String                                      `tfsdk:"ftd_platform_settings_id"`
	Type                  types.String                                      `tfsdk:"type"`
	SourceEmailAddress    types.String                                      `tfsdk:"source_email_address"`
	Destinations          []FTDPlatformSettingsSyslogEmailSetupDestinations `tfsdk:"destinations"`
}

type FTDPlatformSettingsSyslogEmailSetupDestinations struct {
	EmailAddresses types.List   `tfsdk:"email_addresses"`
	LogLevel       types.String `tfsdk:"log_level"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionCreateFTDPlatformSettingsSyslogEmailSetup = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsSyslogEmailSetup) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/syslog/loggingemailsetups", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsSyslogEmailSetup) toBody(ctx context.Context, state FTDPlatformSettingsSyslogEmailSetup) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.SourceEmailAddress.IsNull() {
		body, _ = sjson.Set(body, "sourceEmail", data.SourceEmailAddress.ValueString())
	}
	if len(data.Destinations) > 0 {
		body, _ = sjson.Set(body, "destinationEmails", []any{})
		for _, item := range data.Destinations {
			itemBody := ""
			if !item.EmailAddresses.IsNull() {
				var values []string
				item.EmailAddresses.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "emailIds", values)
			}
			if !item.LogLevel.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "logLevel", item.LogLevel.ValueString())
			}
			body, _ = sjson.SetRaw(body, "destinationEmails.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsSyslogEmailSetup) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("sourceEmail"); value.Exists() {
		data.SourceEmailAddress = types.StringValue(value.String())
	} else {
		data.SourceEmailAddress = types.StringNull()
	}
	if value := res.Get("destinationEmails"); value.Exists() {
		data.Destinations = make([]FTDPlatformSettingsSyslogEmailSetupDestinations, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsSyslogEmailSetupDestinations{}
			if value := res.Get("emailIds"); value.Exists() {
				data.EmailAddresses = helpers.GetStringList(value.Array())
			} else {
				data.EmailAddresses = types.ListNull(types.StringType)
			}
			if value := res.Get("logLevel"); value.Exists() {
				data.LogLevel = types.StringValue(value.String())
			} else {
				data.LogLevel = types.StringNull()
			}
			(*parent).Destinations = append((*parent).Destinations, data)
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
func (data *FTDPlatformSettingsSyslogEmailSetup) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("sourceEmail"); value.Exists() && !data.SourceEmailAddress.IsNull() {
		data.SourceEmailAddress = types.StringValue(value.String())
	} else {
		data.SourceEmailAddress = types.StringNull()
	}
	for i := 0; i < len(data.Destinations); i++ {
		keys := [...]string{"logLevel"}
		keyValues := [...]string{data.Destinations[i].LogLevel.ValueString()}

		parent := &data
		data := (*parent).Destinations[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("destinationEmails").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Destinations[%d] = %+v",
				i,
				(*parent).Destinations[i],
			))
			(*parent).Destinations = slices.Delete((*parent).Destinations, i, i+1)
			i--

			continue
		}
		if value := res.Get("emailIds"); value.Exists() && !data.EmailAddresses.IsNull() {
			data.EmailAddresses = helpers.GetStringList(value.Array())
		} else {
			data.EmailAddresses = types.ListNull(types.StringType)
		}
		if value := res.Get("logLevel"); value.Exists() && !data.LogLevel.IsNull() {
			data.LogLevel = types.StringValue(value.String())
		} else {
			data.LogLevel = types.StringNull()
		}
		(*parent).Destinations[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsSyslogEmailSetup) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data FTDPlatformSettingsSyslogEmailSetup) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	body, _ = sjson.Set(body, "sourceEmail", "empty@fmc.local")
	return body
}
