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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type DeviceVTIInterface struct {
	Id                               types.String `tfsdk:"id"`
	Domain                           types.String `tfsdk:"domain"`
	DeviceId                         types.String `tfsdk:"device_id"`
	Type                             types.String `tfsdk:"type"`
	Name                             types.String `tfsdk:"name"`
	TunnelType                       types.String `tfsdk:"tunnel_type"`
	LogicalName                      types.String `tfsdk:"logical_name"`
	Enabled                          types.Bool   `tfsdk:"enabled"`
	Description                      types.String `tfsdk:"description"`
	SecurityZoneId                   types.String `tfsdk:"security_zone_id"`
	Priority                         types.Int64  `tfsdk:"priority"`
	TunnelId                         types.Int64  `tfsdk:"tunnel_id"`
	TunnelSourceInterfaceId          types.String `tfsdk:"tunnel_source_interface_id"`
	TunnelSourceInterfaceName        types.String `tfsdk:"tunnel_source_interface_name"`
	TunnelSourceInterfaceIpv6Address types.String `tfsdk:"tunnel_source_interface_ipv6_address"`
	IpsecTunnelMode                  types.String `tfsdk:"ipsec_tunnel_mode"`
	Ipv4Address                      types.String `tfsdk:"ipv4_address"`
	Ipv4Netmask                      types.String `tfsdk:"ipv4_netmask"`
	Ipv6Address                      types.String `tfsdk:"ipv6_address"`
	Ipv6Prefix                       types.String `tfsdk:"ipv6_prefix"`
	BorrowIpFromInterfaceId          types.String `tfsdk:"borrow_ip_from_interface_id"`
	BorrowIpFromInterfaceName        types.String `tfsdk:"borrow_ip_from_interface_name"`
	IpBasedMonitoring                types.Bool   `tfsdk:"ip_based_monitoring"`
	IpBasedMonitoringType            types.String `tfsdk:"ip_based_monitoring_type"`
	IpBasedMonitoringNextHop         types.String `tfsdk:"ip_based_monitoring_next_hop"`
	HttpBasedApplicationMonitoring   types.Bool   `tfsdk:"http_based_application_monitoring"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data DeviceVTIInterface) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/devices/devicerecords/%v/virtualtunnelinterfaces", url.QueryEscape(data.DeviceId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data DeviceVTIInterface) toBody(ctx context.Context, state DeviceVTIInterface) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.TunnelType.IsNull() {
		body, _ = sjson.Set(body, "tunnelType", data.TunnelType.ValueString())
	}
	if !data.LogicalName.IsNull() {
		body, _ = sjson.Set(body, "ifname", data.LogicalName.ValueString())
	}
	if !data.Enabled.IsNull() {
		body, _ = sjson.Set(body, "enabled", data.Enabled.ValueBool())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if !data.SecurityZoneId.IsNull() {
		body, _ = sjson.Set(body, "securityZone.id", data.SecurityZoneId.ValueString())
	}
	body, _ = sjson.Set(body, "securityZone.type", "SecurityZone")
	if !data.Priority.IsNull() {
		body, _ = sjson.Set(body, "priority", data.Priority.ValueInt64())
	}
	if !data.TunnelId.IsNull() {
		body, _ = sjson.Set(body, "tunnelId", data.TunnelId.ValueInt64())
	}
	if !data.TunnelSourceInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "tunnelSource.id", data.TunnelSourceInterfaceId.ValueString())
	}
	if !data.TunnelSourceInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "tunnelSource.name", data.TunnelSourceInterfaceName.ValueString())
	}
	if !data.TunnelSourceInterfaceIpv6Address.IsNull() {
		body, _ = sjson.Set(body, "tunnelSrcIPv6IntfAddr", data.TunnelSourceInterfaceIpv6Address.ValueString())
	}
	if !data.IpsecTunnelMode.IsNull() {
		body, _ = sjson.Set(body, "ipsecMode", data.IpsecTunnelMode.ValueString())
	}
	if !data.Ipv4Address.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.address", data.Ipv4Address.ValueString())
	}
	if !data.Ipv4Netmask.IsNull() {
		body, _ = sjson.Set(body, "ipv4.static.netmask", data.Ipv4Netmask.ValueString())
	}
	if !data.Ipv6Address.IsNull() {
		body, _ = sjson.Set(body, "ipv6.addresses.0.address", data.Ipv6Address.ValueString())
	}
	if !data.Ipv6Prefix.IsNull() {
		body, _ = sjson.Set(body, "ipv6.addresses.0.prefix", data.Ipv6Prefix.ValueString())
	}
	if !data.BorrowIpFromInterfaceId.IsNull() {
		body, _ = sjson.Set(body, "borrowIPfrom.id", data.BorrowIpFromInterfaceId.ValueString())
	}
	if !data.BorrowIpFromInterfaceName.IsNull() {
		body, _ = sjson.Set(body, "borrowIPfrom.name", data.BorrowIpFromInterfaceName.ValueString())
	}
	if !data.IpBasedMonitoring.IsNull() {
		body, _ = sjson.Set(body, "pathMonitoring.enable", data.IpBasedMonitoring.ValueBool())
	}
	if !data.IpBasedMonitoringType.IsNull() {
		body, _ = sjson.Set(body, "pathMonitoring.type", data.IpBasedMonitoringType.ValueString())
	}
	if !data.IpBasedMonitoringNextHop.IsNull() {
		body, _ = sjson.Set(body, "pathMonitoring.monitoredIp", data.IpBasedMonitoringNextHop.ValueString())
	}
	if !data.HttpBasedApplicationMonitoring.IsNull() {
		body, _ = sjson.Set(body, "applicationMonitoring.enable", data.HttpBasedApplicationMonitoring.ValueBool())
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *DeviceVTIInterface) fromBody(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("enabled"); value.Exists() {
		data.Enabled = types.BoolValue(value.Bool())
	} else {
		data.Enabled = types.BoolValue(true)
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("securityZone.id"); value.Exists() {
		data.SecurityZoneId = types.StringValue(value.String())
	} else {
		data.SecurityZoneId = types.StringNull()
	}
	if value := res.Get("priority"); value.Exists() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("tunnelId"); value.Exists() {
		data.TunnelId = types.Int64Value(value.Int())
	} else {
		data.TunnelId = types.Int64Null()
	}
	if value := res.Get("tunnelSource.id"); value.Exists() {
		data.TunnelSourceInterfaceId = types.StringValue(value.String())
	} else {
		data.TunnelSourceInterfaceId = types.StringNull()
	}
	if value := res.Get("tunnelSource.name"); value.Exists() {
		data.TunnelSourceInterfaceName = types.StringValue(value.String())
	} else {
		data.TunnelSourceInterfaceName = types.StringNull()
	}
	if value := res.Get("tunnelSrcIPv6IntfAddr"); value.Exists() {
		data.TunnelSourceInterfaceIpv6Address = types.StringValue(value.String())
	} else {
		data.TunnelSourceInterfaceIpv6Address = types.StringNull()
	}
	if value := res.Get("ipsecMode"); value.Exists() {
		data.IpsecTunnelMode = types.StringValue(value.String())
	} else {
		data.IpsecTunnelMode = types.StringNull()
	}
	if value := res.Get("ipv4.static.address"); value.Exists() {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() {
		data.Ipv4Netmask = types.StringValue(value.String())
	} else {
		data.Ipv4Netmask = types.StringNull()
	}
	if value := res.Get("ipv6.addresses.0.address"); value.Exists() {
		data.Ipv6Address = types.StringValue(value.String())
	} else {
		data.Ipv6Address = types.StringNull()
	}
	if value := res.Get("ipv6.addresses.0.prefix"); value.Exists() {
		data.Ipv6Prefix = types.StringValue(value.String())
	} else {
		data.Ipv6Prefix = types.StringNull()
	}
	if value := res.Get("borrowIPfrom.id"); value.Exists() {
		data.BorrowIpFromInterfaceId = types.StringValue(value.String())
	} else {
		data.BorrowIpFromInterfaceId = types.StringNull()
	}
	if value := res.Get("borrowIPfrom.name"); value.Exists() {
		data.BorrowIpFromInterfaceName = types.StringValue(value.String())
	} else {
		data.BorrowIpFromInterfaceName = types.StringNull()
	}
	if value := res.Get("pathMonitoring.enable"); value.Exists() {
		data.IpBasedMonitoring = types.BoolValue(value.Bool())
	} else {
		data.IpBasedMonitoring = types.BoolNull()
	}
	if value := res.Get("pathMonitoring.type"); value.Exists() {
		data.IpBasedMonitoringType = types.StringValue(value.String())
	} else {
		data.IpBasedMonitoringType = types.StringNull()
	}
	if value := res.Get("pathMonitoring.monitoredIp"); value.Exists() {
		data.IpBasedMonitoringNextHop = types.StringValue(value.String())
	} else {
		data.IpBasedMonitoringNextHop = types.StringNull()
	}
	if value := res.Get("applicationMonitoring.enable"); value.Exists() {
		data.HttpBasedApplicationMonitoring = types.BoolValue(value.Bool())
	} else {
		data.HttpBasedApplicationMonitoring = types.BoolNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *DeviceVTIInterface) fromBodyPartial(ctx context.Context, res gjson.Result) {
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
	if value := res.Get("enabled"); value.Exists() && !data.Enabled.IsNull() {
		data.Enabled = types.BoolValue(value.Bool())
	} else if data.Enabled.ValueBool() != true {
		data.Enabled = types.BoolNull()
	}
	if value := res.Get("description"); value.Exists() && !data.Description.IsNull() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("securityZone.id"); value.Exists() && !data.SecurityZoneId.IsNull() {
		data.SecurityZoneId = types.StringValue(value.String())
	} else {
		data.SecurityZoneId = types.StringNull()
	}
	if value := res.Get("priority"); value.Exists() && !data.Priority.IsNull() {
		data.Priority = types.Int64Value(value.Int())
	} else {
		data.Priority = types.Int64Null()
	}
	if value := res.Get("tunnelId"); value.Exists() && !data.TunnelId.IsNull() {
		data.TunnelId = types.Int64Value(value.Int())
	} else {
		data.TunnelId = types.Int64Null()
	}
	if value := res.Get("tunnelSource.id"); value.Exists() && !data.TunnelSourceInterfaceId.IsNull() {
		data.TunnelSourceInterfaceId = types.StringValue(value.String())
	} else {
		data.TunnelSourceInterfaceId = types.StringNull()
	}
	if value := res.Get("tunnelSource.name"); value.Exists() && !data.TunnelSourceInterfaceName.IsNull() {
		data.TunnelSourceInterfaceName = types.StringValue(value.String())
	} else {
		data.TunnelSourceInterfaceName = types.StringNull()
	}
	if value := res.Get("tunnelSrcIPv6IntfAddr"); value.Exists() && !data.TunnelSourceInterfaceIpv6Address.IsNull() {
		data.TunnelSourceInterfaceIpv6Address = types.StringValue(value.String())
	} else {
		data.TunnelSourceInterfaceIpv6Address = types.StringNull()
	}
	if value := res.Get("ipsecMode"); value.Exists() && !data.IpsecTunnelMode.IsNull() {
		data.IpsecTunnelMode = types.StringValue(value.String())
	} else {
		data.IpsecTunnelMode = types.StringNull()
	}
	if value := res.Get("ipv4.static.address"); value.Exists() && !data.Ipv4Address.IsNull() {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
	if value := res.Get("ipv4.static.netmask"); value.Exists() && !data.Ipv4Netmask.IsNull() {
		data.Ipv4Netmask = types.StringValue(value.String())
	} else {
		data.Ipv4Netmask = types.StringNull()
	}
	if value := res.Get("ipv6.addresses.0.address"); value.Exists() && !data.Ipv6Address.IsNull() {
		data.Ipv6Address = types.StringValue(value.String())
	} else {
		data.Ipv6Address = types.StringNull()
	}
	if value := res.Get("ipv6.addresses.0.prefix"); value.Exists() && !data.Ipv6Prefix.IsNull() {
		data.Ipv6Prefix = types.StringValue(value.String())
	} else {
		data.Ipv6Prefix = types.StringNull()
	}
	if value := res.Get("borrowIPfrom.id"); value.Exists() && !data.BorrowIpFromInterfaceId.IsNull() {
		data.BorrowIpFromInterfaceId = types.StringValue(value.String())
	} else {
		data.BorrowIpFromInterfaceId = types.StringNull()
	}
	if value := res.Get("borrowIPfrom.name"); value.Exists() && !data.BorrowIpFromInterfaceName.IsNull() {
		data.BorrowIpFromInterfaceName = types.StringValue(value.String())
	} else {
		data.BorrowIpFromInterfaceName = types.StringNull()
	}
	if value := res.Get("pathMonitoring.enable"); value.Exists() && !data.IpBasedMonitoring.IsNull() {
		data.IpBasedMonitoring = types.BoolValue(value.Bool())
	} else {
		data.IpBasedMonitoring = types.BoolNull()
	}
	if value := res.Get("pathMonitoring.type"); value.Exists() && !data.IpBasedMonitoringType.IsNull() {
		data.IpBasedMonitoringType = types.StringValue(value.String())
	} else {
		data.IpBasedMonitoringType = types.StringNull()
	}
	if value := res.Get("pathMonitoring.monitoredIp"); value.Exists() && !data.IpBasedMonitoringNextHop.IsNull() {
		data.IpBasedMonitoringNextHop = types.StringValue(value.String())
	} else {
		data.IpBasedMonitoringNextHop = types.StringNull()
	}
	if value := res.Get("applicationMonitoring.enable"); value.Exists() && !data.HttpBasedApplicationMonitoring.IsNull() {
		data.HttpBasedApplicationMonitoring = types.BoolValue(value.Bool())
	} else {
		data.HttpBasedApplicationMonitoring = types.BoolNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *DeviceVTIInterface) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
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
