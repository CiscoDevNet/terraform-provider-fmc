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

type DeviceDDNS struct {
	Id                        types.String                       `tfsdk:"id"`
	Domain                    types.String                       `tfsdk:"domain"`
	DeviceId                  types.String                       `tfsdk:"device_id"`
	Type                      types.String                       `tfsdk:"type"`
	DhcpClientRequestType     types.String                       `tfsdk:"dhcp_client_request_type"`
	DhcpClientBroadcast       types.Bool                         `tfsdk:"dhcp_client_broadcast"`
	DynamicDnsUpdate          types.String                       `tfsdk:"dynamic_dns_update"`
	DhcpClientRequestOverride types.Bool                         `tfsdk:"dhcp_client_request_override"`
	DhcpClientIdInterfaces    []DeviceDDNSDhcpClientIdInterfaces `tfsdk:"dhcp_client_id_interfaces"`
	DdnsUpdateMethods         []DeviceDDNSDdnsUpdateMethods      `tfsdk:"ddns_update_methods"`
	DdnsInterfaceSettings     []DeviceDDNSDdnsInterfaceSettings  `tfsdk:"ddns_interface_settings"`
}

type DeviceDDNSDhcpClientIdInterfaces struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

type DeviceDDNSDdnsUpdateMethods struct {
	Name                 types.String `tfsdk:"name"`
	Method               types.String `tfsdk:"method"`
	UpdateIntervalDay    types.Int64  `tfsdk:"update_interval_day"`
	UpdateIntervalHour   types.Int64  `tfsdk:"update_interval_hour"`
	UpdateIntervalMinute types.Int64  `tfsdk:"update_interval_minute"`
	UpdateIntervalSecond types.Int64  `tfsdk:"update_interval_second"`
	WebUrl               types.String `tfsdk:"web_url"`
	WebUpdateType        types.String `tfsdk:"web_update_type"`
	UpdateRecords        types.String `tfsdk:"update_records"`
}

type DeviceDDNSDdnsInterfaceSettings struct {
	InterfaceId               types.String `tfsdk:"interface_id"`
	InterfaceName             types.String `tfsdk:"interface_name"`
	InterfaceType             types.String `tfsdk:"interface_type"`
	MethodName                types.String `tfsdk:"method_name"`
	Hostname                  types.String `tfsdk:"hostname"`
	DhcpClientRequestType     types.String `tfsdk:"dhcp_client_request_type"`
	DynamicDnsUpdate          types.String `tfsdk:"dynamic_dns_update"`
	DhcpClientRequestOverride types.Bool   `tfsdk:"dhcp_client_request_override"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions
var minFMCVersionDeviceDDNS = version.Must(version.NewVersion("7.4"))

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceDDNS) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/dhcp/ddnssettings", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceDDNS) toBody(ctx context.Context, state DeviceDDNS) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.DhcpClientRequestType.IsNull() {
		body, _ = sjson.Set(body, "dhcpClientRequestType", data.DhcpClientRequestType.ValueString())
	}
	if !data.DhcpClientBroadcast.IsNull() {
		body, _ = sjson.Set(body, "enableDHCPClientBroadcast", data.DhcpClientBroadcast.ValueBool())
	}
	if !data.DynamicDnsUpdate.IsNull() {
		body, _ = sjson.Set(body, "dynamicDDNSUpdateType", data.DynamicDnsUpdate.ValueString())
	}
	if !data.DhcpClientRequestOverride.IsNull() {
		body, _ = sjson.Set(body, "overrideDHCPClientRequest", data.DhcpClientRequestOverride.ValueBool())
	}
	if len(data.DhcpClientIdInterfaces) > 0 {
		body, _ = sjson.Set(body, "dhcpClientIdInterfaces", []any{})
		for _, item := range data.DhcpClientIdInterfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Type.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "type", item.Type.ValueString())
			}
			body, _ = sjson.SetRaw(body, "dhcpClientIdInterfaces.-1", itemBody)
		}
	}
	if len(data.DdnsUpdateMethods) > 0 {
		body, _ = sjson.Set(body, "ddnsUpdateMethods", []any{})
		for _, item := range data.DdnsUpdateMethods {
			itemBody := ""
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			if !item.Method.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "methodOption", item.Method.ValueString())
			}
			if !item.UpdateIntervalDay.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "updateInterval.day", item.UpdateIntervalDay.ValueInt64())
			}
			if !item.UpdateIntervalHour.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "updateInterval.hourOfTheDay", item.UpdateIntervalHour.ValueInt64())
			}
			if !item.UpdateIntervalMinute.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "updateInterval.minute", item.UpdateIntervalMinute.ValueInt64())
			}
			if !item.UpdateIntervalSecond.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "updateInterval.second", item.UpdateIntervalSecond.ValueInt64())
			}
			if !item.WebUrl.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "webURL", item.WebUrl.ValueString())
			}
			if !item.WebUpdateType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "webUpdateType", item.WebUpdateType.ValueString())
			}
			if !item.UpdateRecords.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "updateRecords", item.UpdateRecords.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ddnsUpdateMethods.-1", itemBody)
		}
	}
	if len(data.DdnsInterfaceSettings) > 0 {
		body, _ = sjson.Set(body, "ddnsInterfaceSettings", []any{})
		for _, item := range data.DdnsInterfaceSettings {
			itemBody := ""
			if !item.InterfaceId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.id", item.InterfaceId.ValueString())
			}
			if !item.InterfaceName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.name", item.InterfaceName.ValueString())
			}
			if !item.InterfaceType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "interface.type", item.InterfaceType.ValueString())
			}
			if !item.MethodName.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "updateMethodName", item.MethodName.ValueString())
			}
			if !item.Hostname.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "hostname", item.Hostname.ValueString())
			}
			if !item.DhcpClientRequestType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dhcpClientRequestType", item.DhcpClientRequestType.ValueString())
			}
			if !item.DynamicDnsUpdate.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "dynamicDDNSUpdateType", item.DynamicDnsUpdate.ValueString())
			}
			if !item.DhcpClientRequestOverride.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "overrideDHCPClientRequest", item.DhcpClientRequestOverride.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "ddnsInterfaceSettings.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceDDNS) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("dhcpClientRequestType"); value.Exists() {
		data.DhcpClientRequestType = types.StringValue(value.String())
	} else {
		data.DhcpClientRequestType = types.StringNull()
	}
	if value := res.Get("enableDHCPClientBroadcast"); value.Exists() {
		data.DhcpClientBroadcast = types.BoolValue(value.Bool())
	} else {
		data.DhcpClientBroadcast = types.BoolNull()
	}
	if value := res.Get("dynamicDDNSUpdateType"); value.Exists() {
		data.DynamicDnsUpdate = types.StringValue(value.String())
	} else {
		data.DynamicDnsUpdate = types.StringNull()
	}
	if value := res.Get("overrideDHCPClientRequest"); value.Exists() {
		data.DhcpClientRequestOverride = types.BoolValue(value.Bool())
	} else {
		data.DhcpClientRequestOverride = types.BoolNull()
	}
	if value := res.Get("dhcpClientIdInterfaces"); value.Exists() {
		data.DhcpClientIdInterfaces = make([]DeviceDDNSDhcpClientIdInterfaces, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceDDNSDhcpClientIdInterfaces{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
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
			(*parent).DhcpClientIdInterfaces = append((*parent).DhcpClientIdInterfaces, data)
			return true
		})
	}
	if value := res.Get("ddnsUpdateMethods"); value.Exists() {
		data.DdnsUpdateMethods = make([]DeviceDDNSDdnsUpdateMethods, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceDDNSDdnsUpdateMethods{}
			if value := res.Get("name"); value.Exists() {
				data.Name = types.StringValue(value.String())
			} else {
				data.Name = types.StringNull()
			}
			if value := res.Get("methodOption"); value.Exists() {
				data.Method = types.StringValue(value.String())
			} else {
				data.Method = types.StringNull()
			}
			if value := res.Get("updateInterval.day"); value.Exists() {
				data.UpdateIntervalDay = types.Int64Value(value.Int())
			} else {
				data.UpdateIntervalDay = types.Int64Null()
			}
			if value := res.Get("updateInterval.hourOfTheDay"); value.Exists() {
				data.UpdateIntervalHour = types.Int64Value(value.Int())
			} else {
				data.UpdateIntervalHour = types.Int64Null()
			}
			if value := res.Get("updateInterval.minute"); value.Exists() {
				data.UpdateIntervalMinute = types.Int64Value(value.Int())
			} else {
				data.UpdateIntervalMinute = types.Int64Null()
			}
			if value := res.Get("updateInterval.second"); value.Exists() {
				data.UpdateIntervalSecond = types.Int64Value(value.Int())
			} else {
				data.UpdateIntervalSecond = types.Int64Null()
			}
			if value := res.Get("webURL"); value.Exists() {
				data.WebUrl = types.StringValue(value.String())
			} else {
				data.WebUrl = types.StringNull()
			}
			if value := res.Get("webUpdateType"); value.Exists() {
				data.WebUpdateType = types.StringValue(value.String())
			} else {
				data.WebUpdateType = types.StringNull()
			}
			if value := res.Get("updateRecords"); value.Exists() {
				data.UpdateRecords = types.StringValue(value.String())
			} else {
				data.UpdateRecords = types.StringNull()
			}
			(*parent).DdnsUpdateMethods = append((*parent).DdnsUpdateMethods, data)
			return true
		})
	}
	if value := res.Get("ddnsInterfaceSettings"); value.Exists() {
		data.DdnsInterfaceSettings = make([]DeviceDDNSDdnsInterfaceSettings, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceDDNSDdnsInterfaceSettings{}
			if value := res.Get("interface.id"); value.Exists() {
				data.InterfaceId = types.StringValue(value.String())
			} else {
				data.InterfaceId = types.StringNull()
			}
			if value := res.Get("interface.name"); value.Exists() {
				data.InterfaceName = types.StringValue(value.String())
			} else {
				data.InterfaceName = types.StringNull()
			}
			if value := res.Get("interface.type"); value.Exists() {
				data.InterfaceType = types.StringValue(value.String())
			} else {
				data.InterfaceType = types.StringNull()
			}
			if value := res.Get("updateMethodName"); value.Exists() {
				data.MethodName = types.StringValue(value.String())
			} else {
				data.MethodName = types.StringNull()
			}
			if value := res.Get("hostname"); value.Exists() {
				data.Hostname = types.StringValue(value.String())
			} else {
				data.Hostname = types.StringNull()
			}
			if value := res.Get("dhcpClientRequestType"); value.Exists() {
				data.DhcpClientRequestType = types.StringValue(value.String())
			} else {
				data.DhcpClientRequestType = types.StringNull()
			}
			if value := res.Get("dynamicDDNSUpdateType"); value.Exists() {
				data.DynamicDnsUpdate = types.StringValue(value.String())
			} else {
				data.DynamicDnsUpdate = types.StringNull()
			}
			if value := res.Get("overrideDHCPClientRequest"); value.Exists() {
				data.DhcpClientRequestOverride = types.BoolValue(value.Bool())
			} else {
				data.DhcpClientRequestOverride = types.BoolNull()
			}
			(*parent).DdnsInterfaceSettings = append((*parent).DdnsInterfaceSettings, data)
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
func (data *DeviceDDNS) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("dhcpClientRequestType"); value.Exists() && !data.DhcpClientRequestType.IsNull() {
		data.DhcpClientRequestType = types.StringValue(value.String())
	} else {
		data.DhcpClientRequestType = types.StringNull()
	}
	if value := res.Get("enableDHCPClientBroadcast"); value.Exists() && !data.DhcpClientBroadcast.IsNull() {
		data.DhcpClientBroadcast = types.BoolValue(value.Bool())
	} else {
		data.DhcpClientBroadcast = types.BoolNull()
	}
	if value := res.Get("dynamicDDNSUpdateType"); value.Exists() && !data.DynamicDnsUpdate.IsNull() {
		data.DynamicDnsUpdate = types.StringValue(value.String())
	} else {
		data.DynamicDnsUpdate = types.StringNull()
	}
	if value := res.Get("overrideDHCPClientRequest"); value.Exists() && !data.DhcpClientRequestOverride.IsNull() {
		data.DhcpClientRequestOverride = types.BoolValue(value.Bool())
	} else {
		data.DhcpClientRequestOverride = types.BoolNull()
	}
	for i := 0; i < len(data.DhcpClientIdInterfaces); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.DhcpClientIdInterfaces[i].Id.ValueString()}

		parent := &data
		data := (*parent).DhcpClientIdInterfaces[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("dhcpClientIdInterfaces").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DhcpClientIdInterfaces[%d] = %+v",
				i,
				(*parent).DhcpClientIdInterfaces[i],
			))
			(*parent).DhcpClientIdInterfaces = slices.Delete((*parent).DhcpClientIdInterfaces, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
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
		(*parent).DhcpClientIdInterfaces[i] = data
	}
	for i := 0; i < len(data.DdnsUpdateMethods); i++ {
		keys := [...]string{"name"}
		keyValues := [...]string{data.DdnsUpdateMethods[i].Name.ValueString()}

		parent := &data
		data := (*parent).DdnsUpdateMethods[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ddnsUpdateMethods").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DdnsUpdateMethods[%d] = %+v",
				i,
				(*parent).DdnsUpdateMethods[i],
			))
			(*parent).DdnsUpdateMethods = slices.Delete((*parent).DdnsUpdateMethods, i, i+1)
			i--

			continue
		}
		if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
		}
		if value := res.Get("methodOption"); value.Exists() && !data.Method.IsNull() {
			data.Method = types.StringValue(value.String())
		} else {
			data.Method = types.StringNull()
		}
		if value := res.Get("updateInterval.day"); value.Exists() && !data.UpdateIntervalDay.IsNull() {
			data.UpdateIntervalDay = types.Int64Value(value.Int())
		} else {
			data.UpdateIntervalDay = types.Int64Null()
		}
		if value := res.Get("updateInterval.hourOfTheDay"); value.Exists() && !data.UpdateIntervalHour.IsNull() {
			data.UpdateIntervalHour = types.Int64Value(value.Int())
		} else {
			data.UpdateIntervalHour = types.Int64Null()
		}
		if value := res.Get("updateInterval.minute"); value.Exists() && !data.UpdateIntervalMinute.IsNull() {
			data.UpdateIntervalMinute = types.Int64Value(value.Int())
		} else {
			data.UpdateIntervalMinute = types.Int64Null()
		}
		if value := res.Get("updateInterval.second"); value.Exists() && !data.UpdateIntervalSecond.IsNull() {
			data.UpdateIntervalSecond = types.Int64Value(value.Int())
		} else {
			data.UpdateIntervalSecond = types.Int64Null()
		}
		if value := res.Get("webURL"); value.Exists() && !data.WebUrl.IsNull() {
			data.WebUrl = types.StringValue(value.String())
		} else {
			data.WebUrl = types.StringNull()
		}
		if value := res.Get("webUpdateType"); value.Exists() && !data.WebUpdateType.IsNull() {
			data.WebUpdateType = types.StringValue(value.String())
		} else {
			data.WebUpdateType = types.StringNull()
		}
		if value := res.Get("updateRecords"); value.Exists() && !data.UpdateRecords.IsNull() {
			data.UpdateRecords = types.StringValue(value.String())
		} else {
			data.UpdateRecords = types.StringNull()
		}
		(*parent).DdnsUpdateMethods[i] = data
	}
	for i := 0; i < len(data.DdnsInterfaceSettings); i++ {
		keys := [...]string{"interface.id"}
		keyValues := [...]string{data.DdnsInterfaceSettings[i].InterfaceId.ValueString()}

		parent := &data
		data := (*parent).DdnsInterfaceSettings[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ddnsInterfaceSettings").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing DdnsInterfaceSettings[%d] = %+v",
				i,
				(*parent).DdnsInterfaceSettings[i],
			))
			(*parent).DdnsInterfaceSettings = slices.Delete((*parent).DdnsInterfaceSettings, i, i+1)
			i--

			continue
		}
		if value := res.Get("interface.id"); value.Exists() && !data.InterfaceId.IsNull() {
			data.InterfaceId = types.StringValue(value.String())
		} else {
			data.InterfaceId = types.StringNull()
		}
		if value := res.Get("interface.name"); value.Exists() && !data.InterfaceName.IsNull() {
			data.InterfaceName = types.StringValue(value.String())
		} else {
			data.InterfaceName = types.StringNull()
		}
		if value := res.Get("interface.type"); value.Exists() && !data.InterfaceType.IsNull() {
			data.InterfaceType = types.StringValue(value.String())
		} else {
			data.InterfaceType = types.StringNull()
		}
		if value := res.Get("updateMethodName"); value.Exists() && !data.MethodName.IsNull() {
			data.MethodName = types.StringValue(value.String())
		} else {
			data.MethodName = types.StringNull()
		}
		if value := res.Get("hostname"); value.Exists() && !data.Hostname.IsNull() {
			data.Hostname = types.StringValue(value.String())
		} else {
			data.Hostname = types.StringNull()
		}
		if value := res.Get("dhcpClientRequestType"); value.Exists() && !data.DhcpClientRequestType.IsNull() {
			data.DhcpClientRequestType = types.StringValue(value.String())
		} else {
			data.DhcpClientRequestType = types.StringNull()
		}
		if value := res.Get("dynamicDDNSUpdateType"); value.Exists() && !data.DynamicDnsUpdate.IsNull() {
			data.DynamicDnsUpdate = types.StringValue(value.String())
		} else {
			data.DynamicDnsUpdate = types.StringNull()
		}
		if value := res.Get("overrideDHCPClientRequest"); value.Exists() && !data.DhcpClientRequestOverride.IsNull() {
			data.DhcpClientRequestOverride = types.BoolValue(value.Bool())
		} else {
			data.DhcpClientRequestOverride = types.BoolNull()
		}
		(*parent).DdnsInterfaceSettings[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceDDNS) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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

// toBodyPutDelete is used to create the body for PUT requests to clear the resource state
func (data DeviceDDNS) toBodyPutDelete(ctx context.Context) string {
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

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBody

// End of section. //template:end adjustBody

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk

// Section below is generated&owned by "gen/generator.go". //template:begin toBodyOverrides

// End of section. //template:end toBodyOverrides

// Section below is generated&owned by "gen/generator.go". //template:begin synthesizeOverrides

// End of section. //template:end synthesizeOverrides
