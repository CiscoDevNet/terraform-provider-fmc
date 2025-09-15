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

type FTDPlatformSettingsICMPAccess struct {
	Id                    types.String                               `tfsdk:"id"`
	Domain                types.String                               `tfsdk:"domain"`
	FtdPlatformSettingsId types.String                               `tfsdk:"ftd_platform_settings_id"`
	Type                  types.String                               `tfsdk:"type"`
	RateLimit             types.Int64                                `tfsdk:"rate_limit"`
	BurstSize             types.Int64                                `tfsdk:"burst_size"`
	IcmpConfigs           []FTDPlatformSettingsICMPAccessIcmpConfigs `tfsdk:"icmp_configs"`
}

type FTDPlatformSettingsICMPAccessIcmpConfigs struct {
	Action                types.String                                               `tfsdk:"action"`
	IcmpServiceId         types.String                                               `tfsdk:"icmp_service_id"`
	SourceNetworkObjectId types.String                                               `tfsdk:"source_network_object_id"`
	InterfaceLiterals     types.Set                                                  `tfsdk:"interface_literals"`
	InterfaceObjects      []FTDPlatformSettingsICMPAccessIcmpConfigsInterfaceObjects `tfsdk:"interface_objects"`
}

type FTDPlatformSettingsICMPAccessIcmpConfigsInterfaceObjects struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
	Name types.String `tfsdk:"name"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionCreateFTDPlatformSettingsICMPAccess = version.Must(version.NewVersion("7.7"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data FTDPlatformSettingsICMPAccess) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/policy/ftdplatformsettingspolicies/%v/icmpsettings", url.QueryEscape(data.FtdPlatformSettingsId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data FTDPlatformSettingsICMPAccess) toBody(ctx context.Context, state FTDPlatformSettingsICMPAccess) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.RateLimit.IsNull() {
		body, _ = sjson.Set(body, "rateLimit", data.RateLimit.ValueInt64())
	}
	if !data.BurstSize.IsNull() {
		body, _ = sjson.Set(body, "burstSize", data.BurstSize.ValueInt64())
	}
	if len(data.IcmpConfigs) > 0 {
		body, _ = sjson.Set(body, "icmpConfigs", []any{})
		for _, item := range data.IcmpConfigs {
			itemBody := ""
			if !item.Action.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "icmpAction", item.Action.ValueString())
			}
			if !item.IcmpServiceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "icmpService.id", item.IcmpServiceId.ValueString())
			}
			if !item.SourceNetworkObjectId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "network.object.id", item.SourceNetworkObjectId.ValueString())
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
			body, _ = sjson.SetRaw(body, "icmpConfigs.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *FTDPlatformSettingsICMPAccess) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("rateLimit"); value.Exists() {
		data.RateLimit = types.Int64Value(value.Int())
	} else {
		data.RateLimit = types.Int64Value(1)
	}
	if value := res.Get("burstSize"); value.Exists() {
		data.BurstSize = types.Int64Value(value.Int())
	} else {
		data.BurstSize = types.Int64Value(1)
	}
	if value := res.Get("icmpConfigs"); value.Exists() {
		data.IcmpConfigs = make([]FTDPlatformSettingsICMPAccessIcmpConfigs, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := FTDPlatformSettingsICMPAccessIcmpConfigs{}
			if value := res.Get("icmpAction"); value.Exists() {
				data.Action = types.StringValue(value.String())
			} else {
				data.Action = types.StringNull()
			}
			if value := res.Get("icmpService.id"); value.Exists() {
				data.IcmpServiceId = types.StringValue(value.String())
			} else {
				data.IcmpServiceId = types.StringNull()
			}
			if value := res.Get("network.object.id"); value.Exists() {
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
				data.InterfaceObjects = make([]FTDPlatformSettingsICMPAccessIcmpConfigsInterfaceObjects, 0)
				value.ForEach(func(k, res gjson.Result) bool {
					parent := &data
					data := FTDPlatformSettingsICMPAccessIcmpConfigsInterfaceObjects{}
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
			(*parent).IcmpConfigs = append((*parent).IcmpConfigs, data)
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
func (data *FTDPlatformSettingsICMPAccess) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("rateLimit"); value.Exists() && !data.RateLimit.IsNull() {
		data.RateLimit = types.Int64Value(value.Int())
	} else if data.RateLimit.ValueInt64() != 1 {
		data.RateLimit = types.Int64Null()
	}
	if value := res.Get("burstSize"); value.Exists() && !data.BurstSize.IsNull() {
		data.BurstSize = types.Int64Value(value.Int())
	} else if data.BurstSize.ValueInt64() != 1 {
		data.BurstSize = types.Int64Null()
	}
	for i := 0; i < len(data.IcmpConfigs); i++ {
		keys := [...]string{"icmpService.id", "network.object.id"}
		keyValues := [...]string{data.IcmpConfigs[i].IcmpServiceId.ValueString(), data.IcmpConfigs[i].SourceNetworkObjectId.ValueString()}

		parent := &data
		data := (*parent).IcmpConfigs[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("icmpConfigs").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing IcmpConfigs[%d] = %+v",
				i,
				(*parent).IcmpConfigs[i],
			))
			(*parent).IcmpConfigs = slices.Delete((*parent).IcmpConfigs, i, i+1)
			i--

			continue
		}
		if value := res.Get("icmpAction"); value.Exists() && !data.Action.IsNull() {
			data.Action = types.StringValue(value.String())
		} else {
			data.Action = types.StringNull()
		}
		if value := res.Get("icmpService.id"); value.Exists() && !data.IcmpServiceId.IsNull() {
			data.IcmpServiceId = types.StringValue(value.String())
		} else {
			data.IcmpServiceId = types.StringNull()
		}
		if value := res.Get("network.object.id"); value.Exists() && !data.SourceNetworkObjectId.IsNull() {
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
		(*parent).IcmpConfigs[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *FTDPlatformSettingsICMPAccess) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
func (data FTDPlatformSettingsICMPAccess) toBodyPutDelete(ctx context.Context) string {
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
