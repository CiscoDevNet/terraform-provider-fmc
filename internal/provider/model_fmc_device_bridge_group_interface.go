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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceBridgeGroupInterface struct {
	Id                  types.String                                   `tfsdk:"id"`
	Domain              types.String                                   `tfsdk:"domain"`
	DeviceId            types.String                                   `tfsdk:"device_id"`
	Type                types.String                                   `tfsdk:"type"`
	Name                types.String                                   `tfsdk:"name"`
	LogicalName         types.String                                   `tfsdk:"logical_name"`
	Description         types.String                                   `tfsdk:"description"`
	BridgeGroupId       types.Int64                                    `tfsdk:"bridge_group_id"`
	SelectedInterfaces  []DeviceBridgeGroupInterfaceSelectedInterfaces `tfsdk:"selected_interfaces"`
	Ipv4StaticAddress   types.String                                   `tfsdk:"ipv4_static_address"`
	Ipv4StaticNetmask   types.String                                   `tfsdk:"ipv4_static_netmask"`
	Ipv4DhcpObtainRoute types.Bool                                     `tfsdk:"ipv4_dhcp_obtain_route"`
	Ipv6Addresses       []DeviceBridgeGroupInterfaceIpv6Addresses      `tfsdk:"ipv6_addresses"`
	Ipv6EnableDad       types.Bool                                     `tfsdk:"ipv6_enable_dad"`
	Ipv6DadAttempts     types.Int64                                    `tfsdk:"ipv6_dad_attempts"`
	Ipv6NsInterval      types.Int64                                    `tfsdk:"ipv6_ns_interval"`
	Ipv6ReachableTime   types.Int64                                    `tfsdk:"ipv6_reachable_time"`
	ArpTableEntries     []DeviceBridgeGroupInterfaceArpTableEntries    `tfsdk:"arp_table_entries"`
}

type DeviceBridgeGroupInterfaceSelectedInterfaces struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type DeviceBridgeGroupInterfaceIpv6Addresses struct {
	Address types.String `tfsdk:"address"`
	Prefix  types.String `tfsdk:"prefix"`
}

type DeviceBridgeGroupInterfaceArpTableEntries struct {
	MacAddress  types.String `tfsdk:"mac_address"`
	IpAddress   types.String `tfsdk:"ip_address"`
	EnableAlias types.Bool   `tfsdk:"enable_alias"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceBridgeGroupInterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/bridgegroupinterfaces", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceBridgeGroupInterface) toBody(ctx context.Context, state DeviceBridgeGroupInterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.LogicalName.IsNull() {
		body, _ = sjson.Set(body, "ifname", data.LogicalName.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.BridgeGroupId.IsNull() {
		body, _ = sjson.Set(body, "bridgeGroupId", data.BridgeGroupId.ValueInt64())
	}
	if len(data.SelectedInterfaces) > 0 {
		body, _ = sjson.Set(body, "selectedInterfaces", []interface{}{})
		for _, item := range data.SelectedInterfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			if !item.Name.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "name", item.Name.ValueString())
			}
			body, _ = sjson.SetRaw(body, "selectedInterfaces.-1", itemBody)
		}
	}
	if !data.Ipv4StaticAddress.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.address", data.Ipv4StaticAddress.ValueString())
	}
	if !data.Ipv4StaticNetmask.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.netmask", data.Ipv4StaticNetmask.ValueString())
	}
	if !data.Ipv4DhcpObtainRoute.IsNull() {
		body, _ = sjson.Set(body, "ipv4.dhcp.enableDefaultRouteDHCP", data.Ipv4DhcpObtainRoute.ValueBool())
	}
	if len(data.Ipv6Addresses) > 0 {
		body, _ = sjson.Set(body, "ipv6.addresses", []interface{}{})
		for _, item := range data.Ipv6Addresses {
			itemBody := ""
			if !item.Address.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "address", item.Address.ValueString())
			}
			if !item.Prefix.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "prefix", item.Prefix.ValueString())
			}
			body, _ = sjson.SetRaw(body, "ipv6.addresses.-1", itemBody)
		}
	}
	if !data.Ipv6EnableDad.IsNull() {
		body, _ = sjson.Set(body, "ipv6.enableDADLoopback", data.Ipv6EnableDad.ValueBool())
	}
	if !data.Ipv6DadAttempts.IsNull() {
		body, _ = sjson.Set(body, "ipv6.dadAttempts", data.Ipv6DadAttempts.ValueInt64())
	}
	if !data.Ipv6NsInterval.IsNull() {
		body, _ = sjson.Set(body, "ipv6.nsInterval", data.Ipv6NsInterval.ValueInt64())
	}
	if !data.Ipv6ReachableTime.IsNull() {
		body, _ = sjson.Set(body, "ipv6.reachableTime", data.Ipv6ReachableTime.ValueInt64())
	}
	if len(data.ArpTableEntries) > 0 {
		body, _ = sjson.Set(body, "arpConfig", []interface{}{})
		for _, item := range data.ArpTableEntries {
			itemBody := ""
			if !item.MacAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "macAddress", item.MacAddress.ValueString())
			}
			if !item.IpAddress.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "ipAddress", item.IpAddress.ValueString())
			}
			if !item.EnableAlias.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "enableAlias", item.EnableAlias.ValueBool())
			}
			body, _ = sjson.SetRaw(body, "arpConfig.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceBridgeGroupInterface) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("ifname"); value.Exists() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("bridgeGroupId"); value.Exists() {
		data.BridgeGroupId = types.Int64Value(value.Int())
	} else {
		data.BridgeGroupId = types.Int64Null()
	}
	if value := res.Get("selectedInterfaces"); value.Exists() {
		data.SelectedInterfaces = make([]DeviceBridgeGroupInterfaceSelectedInterfaces, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBridgeGroupInterfaceSelectedInterfaces{}
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
			(*parent).SelectedInterfaces = append((*parent).SelectedInterfaces, data)
			return true
		})
	}
	if value := res.Get("ipv4.static.address"); value.Exists() {
		data.Ipv4StaticAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StaticAddress = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() {
		data.Ipv4StaticNetmask = types.StringValue(value.String())
	} else {
		data.Ipv4StaticNetmask = types.StringNull()
	}
	if value := res.Get("ipv4.dhcp.enableDefaultRouteDHCP"); value.Exists() {
		data.Ipv4DhcpObtainRoute = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DhcpObtainRoute = types.BoolNull()
	}
	if value := res.Get("ipv6.addresses"); value.Exists() {
		data.Ipv6Addresses = make([]DeviceBridgeGroupInterfaceIpv6Addresses, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBridgeGroupInterfaceIpv6Addresses{}
			if value := res.Get("address"); value.Exists() {
				data.Address = types.StringValue(value.String())
			} else {
				data.Address = types.StringNull()
			}
			if value := res.Get("prefix"); value.Exists() {
				data.Prefix = types.StringValue(value.String())
			} else {
				data.Prefix = types.StringNull()
			}
			(*parent).Ipv6Addresses = append((*parent).Ipv6Addresses, data)
			return true
		})
	}
	if value := res.Get("ipv6.enableDADLoopback"); value.Exists() {
		data.Ipv6EnableDad = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDad = types.BoolNull()
	}
	if value := res.Get("ipv6.dadAttempts"); value.Exists() {
		data.Ipv6DadAttempts = types.Int64Value(value.Int())
	} else {
		data.Ipv6DadAttempts = types.Int64Null()
	}
	if value := res.Get("ipv6.nsInterval"); value.Exists() {
		data.Ipv6NsInterval = types.Int64Value(value.Int())
	} else {
		data.Ipv6NsInterval = types.Int64Null()
	}
	if value := res.Get("ipv6.reachableTime"); value.Exists() {
		data.Ipv6ReachableTime = types.Int64Value(value.Int())
	} else {
		data.Ipv6ReachableTime = types.Int64Null()
	}
	if value := res.Get("arpConfig"); value.Exists() {
		data.ArpTableEntries = make([]DeviceBridgeGroupInterfaceArpTableEntries, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := DeviceBridgeGroupInterfaceArpTableEntries{}
			if value := res.Get("macAddress"); value.Exists() {
				data.MacAddress = types.StringValue(value.String())
			} else {
				data.MacAddress = types.StringNull()
			}
			if value := res.Get("ipAddress"); value.Exists() {
				data.IpAddress = types.StringValue(value.String())
			} else {
				data.IpAddress = types.StringNull()
			}
			if value := res.Get("enableAlias"); value.Exists() {
				data.EnableAlias = types.BoolValue(value.Bool())
			} else {
				data.EnableAlias = types.BoolNull()
			}
			(*parent).ArpTableEntries = append((*parent).ArpTableEntries, data)
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
func (data *DeviceBridgeGroupInterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("ifname"); value.Exists() && !data.LogicalName.IsNull() {
		data.LogicalName = types.StringValue(value.String())
	} else {
		data.LogicalName = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("bridgeGroupId"); value.Exists() && !data.BridgeGroupId.IsNull() {
		data.BridgeGroupId = types.Int64Value(value.Int())
	} else {
		data.BridgeGroupId = types.Int64Null()
	}
	for i := 0; i < len(data.SelectedInterfaces); i++ {
		keys := [...]string{"id", "name"}
		keyValues := [...]string{data.SelectedInterfaces[i].Id.ValueString(), data.SelectedInterfaces[i].Name.ValueString()}

		parent := &data
		data := (*parent).SelectedInterfaces[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("selectedInterfaces").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing SelectedInterfaces[%d] = %+v",
				i,
				(*parent).SelectedInterfaces[i],
			))
			(*parent).SelectedInterfaces = slices.Delete((*parent).SelectedInterfaces, i, i+1)
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
		(*parent).SelectedInterfaces[i] = data
	}
	if value := res.Get("ipv4.static.address"); value.Exists() && !data.Ipv4StaticAddress.IsNull() {
		data.Ipv4StaticAddress = types.StringValue(value.String())
	} else {
		data.Ipv4StaticAddress = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() && !data.Ipv4StaticNetmask.IsNull() {
		data.Ipv4StaticNetmask = types.StringValue(value.String())
	} else {
		data.Ipv4StaticNetmask = types.StringNull()
	}
	if value := res.Get("ipv4.dhcp.enableDefaultRouteDHCP"); value.Exists() && !data.Ipv4DhcpObtainRoute.IsNull() {
		data.Ipv4DhcpObtainRoute = types.BoolValue(value.Bool())
	} else {
		data.Ipv4DhcpObtainRoute = types.BoolNull()
	}
	for i := 0; i < len(data.Ipv6Addresses); i++ {
		keys := [...]string{"address", "prefix"}
		keyValues := [...]string{data.Ipv6Addresses[i].Address.ValueString(), data.Ipv6Addresses[i].Prefix.ValueString()}

		parent := &data
		data := (*parent).Ipv6Addresses[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("ipv6.addresses").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing Ipv6Addresses[%d] = %+v",
				i,
				(*parent).Ipv6Addresses[i],
			))
			(*parent).Ipv6Addresses = slices.Delete((*parent).Ipv6Addresses, i, i+1)
			i--

			continue
		}
		if value := res.Get("address"); value.Exists() && !data.Address.IsNull() {
			data.Address = types.StringValue(value.String())
		} else {
			data.Address = types.StringNull()
		}
		if value := res.Get("prefix"); value.Exists() && !data.Prefix.IsNull() {
			data.Prefix = types.StringValue(value.String())
		} else {
			data.Prefix = types.StringNull()
		}
		(*parent).Ipv6Addresses[i] = data
	}
	if value := res.Get("ipv6.enableDADLoopback"); value.Exists() && !data.Ipv6EnableDad.IsNull() {
		data.Ipv6EnableDad = types.BoolValue(value.Bool())
	} else {
		data.Ipv6EnableDad = types.BoolNull()
	}
	if value := res.Get("ipv6.dadAttempts"); value.Exists() && !data.Ipv6DadAttempts.IsNull() {
		data.Ipv6DadAttempts = types.Int64Value(value.Int())
	} else {
		data.Ipv6DadAttempts = types.Int64Null()
	}
	if value := res.Get("ipv6.nsInterval"); value.Exists() && !data.Ipv6NsInterval.IsNull() {
		data.Ipv6NsInterval = types.Int64Value(value.Int())
	} else {
		data.Ipv6NsInterval = types.Int64Null()
	}
	if value := res.Get("ipv6.reachableTime"); value.Exists() && !data.Ipv6ReachableTime.IsNull() {
		data.Ipv6ReachableTime = types.Int64Value(value.Int())
	} else {
		data.Ipv6ReachableTime = types.Int64Null()
	}
	for i := 0; i < len(data.ArpTableEntries); i++ {
		keys := [...]string{"macAddress", "ipAddress", "enableAlias"}
		keyValues := [...]string{data.ArpTableEntries[i].MacAddress.ValueString(), data.ArpTableEntries[i].IpAddress.ValueString(), strconv.FormatBool(data.ArpTableEntries[i].EnableAlias.ValueBool())}

		parent := &data
		data := (*parent).ArpTableEntries[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("arpConfig").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing ArpTableEntries[%d] = %+v",
				i,
				(*parent).ArpTableEntries[i],
			))
			(*parent).ArpTableEntries = slices.Delete((*parent).ArpTableEntries, i, i+1)
			i--

			continue
		}
		if value := res.Get("macAddress"); value.Exists() && !data.MacAddress.IsNull() {
			data.MacAddress = types.StringValue(value.String())
		} else {
			data.MacAddress = types.StringNull()
		}
		if value := res.Get("ipAddress"); value.Exists() && !data.IpAddress.IsNull() {
			data.IpAddress = types.StringValue(value.String())
		} else {
			data.IpAddress = types.StringNull()
		}
		if value := res.Get("enableAlias"); value.Exists() && !data.EnableAlias.IsNull() {
			data.EnableAlias = types.BoolValue(value.Bool())
		} else {
			data.EnableAlias = types.BoolNull()
		}
		(*parent).ArpTableEntries[i] = data
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceBridgeGroupInterface) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.Name.IsUnknown() {
		if value := res.Get("name"); value.Exists() {
			data.Name = types.StringValue(value.String())
		} else {
			data.Name = types.StringNull()
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

func (data DeviceBridgeGroupInterface) adjustBody(ctx context.Context, req string) string {
	if len(data.SelectedInterfaces) > 0 {
		req, _ = sjson.Set(req, "priority", 0)
		req, _ = sjson.Set(req, "mode", "NONE")
	}

	if !data.Ipv4DhcpObtainRoute.IsNull() {
		req, _ = sjson.Set(req, "ipv4.dhcp.dhcpRouteMetric", 0)
	}

	return req
}

// Section below is generated&owned by "gen/generator.go". //template:begin adjustBodyBulk

// End of section. //template:end adjustBodyBulk
