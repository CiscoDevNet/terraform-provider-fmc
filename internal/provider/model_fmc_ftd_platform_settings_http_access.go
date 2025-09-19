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

type FTDPlatformSettingsHTTPAccess struct {
	Id                    types.String                                      `tfsdk:"id"`
	Domain                types.String                                      `tfsdk:"domain"`
	FtdPlatformSettingsId types.String                                      `tfsdk:"ftd_platform_settings_id"`
	Type                  types.String                                      `tfsdk:"type"`
	EnableHttpServer      types.Bool                                        `tfsdk:"enable_http_server"`
	HttpServerPort        types.Int64                                       `tfsdk:"http_server_port"`
	HttpConfigurations    []FTDPlatformSettingsHTTPAccessHttpConfigurations `tfsdk:"http_configurations"`
}

type FTDPlatformSettingsHTTPAccessHttpConfigurations struct {
	SourceNetworkObjectId types.String                                                      `tfsdk:"source_network_object_id"`
	InterfaceLiterals     types.Set                                                         `tfsdk:"interface_literals"`
	InterfaceObjects      []FTDPlatformSettingsHTTPAccessHttpConfigurationsInterfaceObjects `tfsdk:"interface_objects"`
}

type FTDPlatformSettingsHTTPAccessHttpConfigurationsInterfaceObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionFTDPlatformSettingsHTTPAccess = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsHTTPAccess) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/httpaccesssettings", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsHTTPAccess) toBody(ctx context.Context, state FTDPlatformSettingsHTTPAccess) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.EnableHttpServer.IsNull() {
		body, _ = sjson.Set(body, "enableHttpServer", data.EnableHttpServer.ValueBool())
	}
	if !data.HttpServerPort.IsNull() {
		body, _ = sjson.Set(body, "port", data.HttpServerPort.ValueInt64())
	}
	if len(data.HttpConfigurations) > 0 {
		body, _ = sjson.Set(body, "httpConfiguration", []any{})
		for _, item := range data.HttpConfigurations {
			itemBody := ""
			if !item.SourceNetworkObjectId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipAddress.id", item.SourceNetworkObjectId.ValueString())
			}
			if !item.InterfaceLiterals.IsNull() {
				var values []string
				item.InterfaceLiterals.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "interfaces.literals", values)
			}
			if len(item.InterfaceObjects) > 0 {
				itemBody, _ = sjson.Set(itemBody, "interfaces.objects", []any{})
				for _, childItem := range item.InterfaceObjects {
					itemChildBody := ""
					if !childItem.Id.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "id", childItem.Id.ValueString())
					}
					if !childItem.Type.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "type", childItem.Type.ValueString())
					}
					if !childItem.Name.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "name", childItem.Name.ValueString())
					}
					itemBody, _ = sjson.SetRaw(itemBody, "interfaces.objects.-1", itemChildBody)
				}
			}
			body, _ = sjson.SetRaw(body, "httpConfiguration.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsHTTPAccess) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enableHttpServer"); value.Exists() {
		data.EnableHttpServer = types.BoolValue(value.Bool())
	} else {
		data.EnableHttpServer = types.BoolNull()
	}
	if value := res.Get("port"); value.Exists() {
		data.HttpServerPort = types.Int64Value(value.Int())
	} else {
		data.HttpServerPort = types.Int64Value(443)
	}
	if value := res.Get("httpConfiguration"); value.Exists() {
		data.HttpConfigurations = make([]FTDPlatformSettingsHTTPAccessHttpConfigurations, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsHTTPAccessHttpConfigurations{}
			if value := res.Get("ipAddress.id"); value.Exists() {
				data.SourceNetworkObjectId = types.StringValue(value.String())
			} else {
				data.SourceNetworkObjectId = types.StringNull()
			}
			if value := res.Get("interfaces.literals"); value.Exists() {
				data.InterfaceLiterals = helpers.GetStringSet(value.Array())
			} else {
				data.InterfaceLiterals = types.SetNull(types.StringType)
			}
			if value := res.Get("interfaces.objects"); value.Exists() {
				data.InterfaceObjects = make([]FTDPlatformSettingsHTTPAccessHttpConfigurationsInterfaceObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := FTDPlatformSettingsHTTPAccessHttpConfigurationsInterfaceObjects{}
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
					if value := res.Get("name"); value.Exists() {
						data.Name = types.StringValue(value.String())
					} else {
						data.Name = types.StringNull()
					}
					(*parent).InterfaceObjects = append((*parent).InterfaceObjects, data)
					return true
				})
			}
			(*parent).HttpConfigurations = append((*parent).HttpConfigurations, data)
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
func (data *FTDPlatformSettingsHTTPAccess) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("enableHttpServer"); value.Exists() && !data.EnableHttpServer.IsNull() {
		data.EnableHttpServer = types.BoolValue(value.Bool())
	} else {
		data.EnableHttpServer = types.BoolNull()
	}
	if value := res.Get("port"); value.Exists() && !data.HttpServerPort.IsNull() {
		data.HttpServerPort = types.Int64Value(value.Int())
	} else if data.HttpServerPort.ValueInt64() != 443 {
		data.HttpServerPort = types.Int64Null()
	}
	for i := 0; i < len(data.HttpConfigurations); i++ {
		keys := [...]string{"ipAddress.id"}
		keyValues := [...]string{data.HttpConfigurations[i].SourceNetworkObjectId.ValueString()}

		parent := &data
		data := (*parent).HttpConfigurations[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("httpConfiguration").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing HttpConfigurations[%d] = %+v",
				i,
				(*parent).HttpConfigurations[i],
			))
			(*parent).HttpConfigurations = slices.Delete((*parent).HttpConfigurations, i, i+1)
			i--

			continue
		}
		if value := res.Get("ipAddress.id"); value.Exists() && !data.SourceNetworkObjectId.IsNull() {
			data.SourceNetworkObjectId = types.StringValue(value.String())
		} else {
			data.SourceNetworkObjectId = types.StringNull()
		}
		if value := res.Get("interfaces.literals"); value.Exists() && !data.InterfaceLiterals.IsNull() {
			data.InterfaceLiterals = helpers.GetStringSet(value.Array())
		} else {
			data.InterfaceLiterals = types.SetNull(types.StringType)
		}
		for i := 0; i < len(data.InterfaceObjects); i++ {
			keys := [...]string{"id", "type", "name"}
			keyValues := [...]string{data.InterfaceObjects[i].Id.ValueString(), data.InterfaceObjects[i].Type.ValueString(), data.InterfaceObjects[i].Name.ValueString()}

			parent := &data
			data := (*parent).InterfaceObjects[i]
			parentRes := &res
			var res gjson.Result

			parentRes.Get("interfaces.objects").ForEach(
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
			if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			(*parent).InterfaceObjects[i] = data
		}
		(*parent).HttpConfigurations[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsHTTPAccess) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
func (data FTDPlatformSettingsHTTPAccess) toBodyPutDelete(ctx context.Context) string {
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
