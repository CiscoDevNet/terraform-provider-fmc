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

type FTDPlatformSettingsDNS struct {
	Id                                     types.String                             `tfsdk:"id"`
	Domain                                 types.String                             `tfsdk:"domain"`
	FtdPlatformSettingsId                  types.String                             `tfsdk:"ftd_platform_settings_id"`
	Type                                   types.String                             `tfsdk:"type"`
	ServerGroups                           []FTDPlatformSettingsDNSServerGroups     `tfsdk:"server_groups"`
	ExpireEntryTimer                       types.Int64                              `tfsdk:"expire_entry_timer"`
	PollTimer                              types.Int64                              `tfsdk:"poll_timer"`
	InterfaceObjects                       []FTDPlatformSettingsDNSInterfaceObjects `tfsdk:"interface_objects"`
	LookupViaManagementDiagnosticInterface types.Bool                               `tfsdk:"lookup_via_management_diagnostic_interface"`
}

type FTDPlatformSettingsDNSServerGroups struct {
	ServerGroupId types.String `tfsdk:"server_group_id"`
	IsDefault     types.Bool   `tfsdk:"is_default"`
	FilterDomains types.List   `tfsdk:"filter_domains"`
}

type FTDPlatformSettingsDNSInterfaceObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionCreateFTDPlatformSettingsDNS = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsDNS) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/dnssettings", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsDNS) toBody(ctx context.Context, state FTDPlatformSettingsDNS) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if len(data.ServerGroups) > 0 {
		body, _ = sjson.Set(body, "dnsServerGroups", []any{})
		for _, item := range data.ServerGroups {
			itemBody := ""
			if !item.ServerGroupId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dnsServerGroup.id", item.ServerGroupId.ValueString())
			}
			if !item.IsDefault.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "isDefault", item.IsDefault.ValueBool())
			}
			if !item.FilterDomains.IsNull() {
				var values []string
				item.FilterDomains.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "bypassDomains", values)
			}
			body, _ = sjson.SetRaw(body, "dnsServerGroups.-1", itemBody)
		}
	}
	if !data.ExpireEntryTimer.IsNull() {
		body, _ = sjson.Set(body, "expiryTimerInMins", data.ExpireEntryTimer.ValueInt64())
	}
	if !data.PollTimer.IsNull() {
		body, _ = sjson.Set(body, "pollTimerInMins", data.PollTimer.ValueInt64())
	}
	if len(data.InterfaceObjects) > 0 {
		body, _ = sjson.Set(body, "interfaceObjects", []any{})
		for _, item := range data.InterfaceObjects {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "interfaceObjects.-1", itemBody)
		}
	}
	if !data.LookupViaManagementDiagnosticInterface.IsNull() {
		body, _ = sjson.Set(body, "enableLookupViaMgmt", data.LookupViaManagementDiagnosticInterface.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsDNS) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("dnsServerGroups"); value.Exists() {
		data.ServerGroups = make([]FTDPlatformSettingsDNSServerGroups, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsDNSServerGroups{}
			if value := res.Get("dnsServerGroup.id"); value.Exists() {
				data.ServerGroupId = types.StringValue(value.String())
			} else {
				data.ServerGroupId = types.StringNull()
			}
			if value := res.Get("isDefault"); value.Exists() {
				data.IsDefault = types.BoolValue(value.Bool())
			} else {
				data.IsDefault = types.BoolNull()
			}
			if value := res.Get("bypassDomains"); value.Exists() {
				data.FilterDomains = helpers.GetStringList(value.Array())
			} else {
				data.FilterDomains = types.ListNull(types.StringType)
			}
			(*parent).ServerGroups = append((*parent).ServerGroups, data)
			return true
		})
	}
	if value := res.Get("expiryTimerInMins"); value.Exists() {
		data.ExpireEntryTimer = types.Int64Value(value.Int())
	} else {
		data.ExpireEntryTimer = types.Int64Null()
	}
	if value := res.Get("pollTimerInMins"); value.Exists() {
		data.PollTimer = types.Int64Value(value.Int())
	} else {
		data.PollTimer = types.Int64Null()
	}
	if value := res.Get("interfaceObjects"); value.Exists() {
		data.InterfaceObjects = make([]FTDPlatformSettingsDNSInterfaceObjects, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsDNSInterfaceObjects{}
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
			(*parent).InterfaceObjects = append((*parent).InterfaceObjects, data)
			return true
		})
	}
	if value := res.Get("enableLookupViaMgmt"); value.Exists() {
		data.LookupViaManagementDiagnosticInterface = types.BoolValue(value.Bool())
	} else {
		data.LookupViaManagementDiagnosticInterface = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *FTDPlatformSettingsDNS) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	for i := 0; i < len(data.ServerGroups); i++ {
		keys := [...]string{"dnsServerGroup.id"}
		keyValues := [...]string{data.ServerGroups[i].ServerGroupId.ValueString()}

		parent := &data
		data := (*parent).ServerGroups[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dnsServerGroups").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ServerGroups[%d] = %+v",
				i,
				(*parent).ServerGroups[i],
			))
			(*parent).ServerGroups = slices.Delete((*parent).ServerGroups, i, i+1)
			i--

			continue
		}
		if value := res.Get("dnsServerGroup.id"); value.Exists() && !data.ServerGroupId.IsNull() {
			data.ServerGroupId = types.StringValue(value.String())
		} else {
			data.ServerGroupId = types.StringNull()
		}
		if value := res.Get("isDefault"); value.Exists() && !data.IsDefault.IsNull() {
			data.IsDefault = types.BoolValue(value.Bool())
		} else {
			data.IsDefault = types.BoolNull()
		}
		if value := res.Get("bypassDomains"); value.Exists() && !data.FilterDomains.IsNull() {
			data.FilterDomains = helpers.GetStringList(value.Array())
		} else {
			data.FilterDomains = types.ListNull(types.StringType)
		}
		(*parent).ServerGroups[i] = data
	}
	if value := res.Get("expiryTimerInMins"); value.Exists() && !data.ExpireEntryTimer.IsNull() {
		data.ExpireEntryTimer = types.Int64Value(value.Int())
	} else {
		data.ExpireEntryTimer = types.Int64Null()
	}
	if value := res.Get("pollTimerInMins"); value.Exists() && !data.PollTimer.IsNull() {
		data.PollTimer = types.Int64Value(value.Int())
	} else {
		data.PollTimer = types.Int64Null()
	}
	for i := 0; i < len(data.InterfaceObjects); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.InterfaceObjects[i].Id.ValueString()}

		parent := &data
		data := (*parent).InterfaceObjects[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("interfaceObjects").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing InterfaceObjects[%d] = %+v",
				i,
				(*parent).InterfaceObjects[i],
			))
			(*parent).InterfaceObjects = slices.Delete((*parent).InterfaceObjects, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
		(*parent).InterfaceObjects[i] = data
	}
	if value := res.Get("enableLookupViaMgmt"); value.Exists() && !data.LookupViaManagementDiagnosticInterface.IsNull() {
		data.LookupViaManagementDiagnosticInterface = types.BoolValue(value.Bool())
	} else {
		data.LookupViaManagementDiagnosticInterface = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsDNS) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
}

// End of section. //template:end fromBodyUnknowns

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyPutDelete

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data FTDPlatformSettingsDNS) toBodyPutDelete(ctx context.Context) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if data.Type.ValueString() != "" {
		body, _ = sjson.Set(body, "type", data.Type.ValueString())
	}
	return body
}

// End of section. //template:end toBodyPutDelete
