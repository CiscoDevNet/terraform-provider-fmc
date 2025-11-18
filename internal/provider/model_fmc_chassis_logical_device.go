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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types

type ChassisLogicalDevice struct {
	Id                    types.String                             `tfsdk:"id"`
	Domain                types.String                             `tfsdk:"domain"`
	ChassisId             types.String                             `tfsdk:"chassis_id"`
	Type                  types.String                             `tfsdk:"type"`
	DeviceId              types.String                             `tfsdk:"device_id"`
	DeviceType            types.String                             `tfsdk:"device_type"`
	Name                  types.String                             `tfsdk:"name"`
	FtdVersion            types.String                             `tfsdk:"ftd_version"`
	Ipv4Address           types.String                             `tfsdk:"ipv4_address"`
	Ipv4Netmask           types.String                             `tfsdk:"ipv4_netmask"`
	Ipv4Gateway           types.String                             `tfsdk:"ipv4_gateway"`
	Ipv6Address           types.String                             `tfsdk:"ipv6_address"`
	Ipv6Prefix            types.Int64                              `tfsdk:"ipv6_prefix"`
	Ipv6Gateway           types.String                             `tfsdk:"ipv6_gateway"`
	SearchDomain          types.String                             `tfsdk:"search_domain"`
	Fqdn                  types.String                             `tfsdk:"fqdn"`
	FirewallMode          types.String                             `tfsdk:"firewall_mode"`
	DnsServers            types.String                             `tfsdk:"dns_servers"`
	DevicePassword        types.String                             `tfsdk:"device_password"`
	AdminState            types.String                             `tfsdk:"admin_state"`
	PermitExpertMode      types.String                             `tfsdk:"permit_expert_mode"`
	ResourceProfileId     types.String                             `tfsdk:"resource_profile_id"`
	ResourceProfileName   types.String                             `tfsdk:"resource_profile_name"`
	AssignedInterfaces    []ChassisLogicalDeviceAssignedInterfaces `tfsdk:"assigned_interfaces"`
	DeviceGroupId         types.String                             `tfsdk:"device_group_id"`
	AccessControlPolicyId types.String                             `tfsdk:"access_control_policy_id"`
	PlatformSettingsId    types.String                             `tfsdk:"platform_settings_id"`
	Licenses              types.Set                                `tfsdk:"licenses"`
	ContainerId           types.String                             `tfsdk:"container_id"`
	ContainerType         types.String                             `tfsdk:"container_type"`
	ContainerName         types.String                             `tfsdk:"container_name"`
	ContainerRole         types.String                             `tfsdk:"container_role"`
	ContainerStatus       types.String                             `tfsdk:"container_status"`
}

type ChassisLogicalDeviceAssignedInterfaces struct {
	Id types.String `tfsdk:"id"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin minimumVersions

// End of section. //template:end minimumVersions

// Section below is generated&owned by "gen/generator.go". //template:begin getPath

func (data ChassisLogicalDevice) getPath() string {
	return fmt.Sprintf("/api/fmc_config/v1/domain/{DOMAIN_UUID}/chassis/fmcmanagedchassis/%v/logicaldevices", url.QueryEscape(data.ChassisId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody

func (data ChassisLogicalDevice) toBody(ctx context.Context, state ChassisLogicalDevice) string {
	body := ""
	if data.Id.ValueString() != "" {
		body, _ = sjson.Set(body, "id", data.Id.ValueString())
	}
	body, _ = sjson.Set(body, "type", "LogicalDevice")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.FtdVersion.IsNull() {
		body, _ = sjson.Set(body, "ftdApplicationVersion", data.FtdVersion.ValueString())
	}
	if !data.Ipv4Address.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.ipv4.ip", data.Ipv4Address.ValueString())
	}
	if !data.Ipv4Netmask.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.ipv4.mask", data.Ipv4Netmask.ValueString())
	}
	if !data.Ipv4Gateway.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.ipv4.gateway", data.Ipv4Gateway.ValueString())
	}
	if !data.Ipv6Address.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.ipv6.ip", data.Ipv6Address.ValueString())
	}
	if !data.Ipv6Prefix.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.ipv6.prefixLength", data.Ipv6Prefix.ValueInt64())
	}
	if !data.Ipv6Gateway.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.ipv6.gateway", data.Ipv6Gateway.ValueString())
	}
	if !data.SearchDomain.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.searchDomain", data.SearchDomain.ValueString())
	}
	if !data.Fqdn.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.fqdn", data.Fqdn.ValueString())
	}
	if !data.FirewallMode.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.firewallMode", data.FirewallMode.ValueString())
	}
	if !data.DnsServers.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.dnsServers", data.DnsServers.ValueString())
	}
	if !data.DevicePassword.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.adminPassword", data.DevicePassword.ValueString())
	}
	if !data.AdminState.IsNull() {
		body, _ = sjson.Set(body, "adminState", data.AdminState.ValueString())
	}
	if !data.PermitExpertMode.IsNull() {
		body, _ = sjson.Set(body, "managementBootstrap.permitExpertMode", data.PermitExpertMode.ValueString())
	}
	if !data.ResourceProfileId.IsNull() {
		body, _ = sjson.Set(body, "resourceProfile.id", data.ResourceProfileId.ValueString())
	}
	if !data.ResourceProfileName.IsNull() {
		body, _ = sjson.Set(body, "resourceProfile.name", data.ResourceProfileName.ValueString())
	}
	if len(data.AssignedInterfaces) > 0 {
		body, _ = sjson.Set(body, "externalPortLink", []any{})
		for _, item := range data.AssignedInterfaces {
			itemBody := ""
			if !item.Id.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "id", item.Id.ValueString())
			}
			body, _ = sjson.SetRaw(body, "externalPortLink.-1", itemBody)
		}
	}
	if !data.DeviceGroupId.IsNull() {
		body, _ = sjson.Set(body, "deviceRegistration.deviceGroup.id", data.DeviceGroupId.ValueString())
	}
	if !data.AccessControlPolicyId.IsNull() {
		body, _ = sjson.Set(body, "deviceRegistration.accessPolicy.id", data.AccessControlPolicyId.ValueString())
	}
	if !data.PlatformSettingsId.IsNull() {
		body, _ = sjson.Set(body, "deviceRegistration.platformSettings.id", data.PlatformSettingsId.ValueString())
	}
	if !data.Licenses.IsNull() {
		var values []string
		data.Licenses.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "deviceRegistration.licenseCaps", values)
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody

func (data *ChassisLogicalDevice) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("metadata.deviceRecordDetails.id"); value.Exists() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("metadata.deviceRecordDetails.type"); value.Exists() {
		data.DeviceType = types.StringValue(value.String())
	} else {
		data.DeviceType = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("ftdApplicationVersion"); value.Exists() {
		data.FtdVersion = types.StringValue(value.String())
	} else {
		data.FtdVersion = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv4.ip"); value.Exists() {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv4.mask"); value.Exists() {
		data.Ipv4Netmask = types.StringValue(value.String())
	} else {
		data.Ipv4Netmask = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv4.gateway"); value.Exists() {
		data.Ipv4Gateway = types.StringValue(value.String())
	} else {
		data.Ipv4Gateway = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv6.ip"); value.Exists() {
		data.Ipv6Address = types.StringValue(value.String())
	} else {
		data.Ipv6Address = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv6.prefixLength"); value.Exists() {
		data.Ipv6Prefix = types.Int64Value(value.Int())
	} else {
		data.Ipv6Prefix = types.Int64Null()
	}
	if value := res.Get("managementBootstrap.ipv6.gateway"); value.Exists() {
		data.Ipv6Gateway = types.StringValue(value.String())
	} else {
		data.Ipv6Gateway = types.StringNull()
	}
	if value := res.Get("managementBootstrap.searchDomain"); value.Exists() {
		data.SearchDomain = types.StringValue(value.String())
	} else {
		data.SearchDomain = types.StringNull()
	}
	if value := res.Get("managementBootstrap.fqdn"); value.Exists() {
		data.Fqdn = types.StringValue(value.String())
	} else {
		data.Fqdn = types.StringNull()
	}
	if value := res.Get("managementBootstrap.firewallMode"); value.Exists() {
		data.FirewallMode = types.StringValue(value.String())
	} else {
		data.FirewallMode = types.StringNull()
	}
	if value := res.Get("managementBootstrap.dnsServers"); value.Exists() {
		data.DnsServers = types.StringValue(value.String())
	} else {
		data.DnsServers = types.StringNull()
	}
	if value := res.Get("adminState"); value.Exists() {
		data.AdminState = types.StringValue(value.String())
	} else {
		data.AdminState = types.StringValue("ENABLED")
	}
	if value := res.Get("managementBootstrap.permitExpertMode"); value.Exists() {
		data.PermitExpertMode = types.StringValue(value.String())
	} else {
		data.PermitExpertMode = types.StringNull()
	}
	if value := res.Get("resourceProfile.id"); value.Exists() {
		data.ResourceProfileId = types.StringValue(value.String())
	} else {
		data.ResourceProfileId = types.StringNull()
	}
	if value := res.Get("resourceProfile.name"); value.Exists() {
		data.ResourceProfileName = types.StringValue(value.String())
	} else {
		data.ResourceProfileName = types.StringNull()
	}
	if value := res.Get("externalPortLink"); value.Exists() {
		data.AssignedInterfaces = make([]ChassisLogicalDeviceAssignedInterfaces, 0)
		value.ForEach(func(k, res gjson.Result) bool {
			parent := &data
			data := ChassisLogicalDeviceAssignedInterfaces{}
			if value := res.Get("id"); value.Exists() {
				data.Id = types.StringValue(value.String())
			} else {
				data.Id = types.StringNull()
			}
			(*parent).AssignedInterfaces = append((*parent).AssignedInterfaces, data)
			return true
		})
	}
	if value := res.Get("deviceRegistration.deviceGroup.id"); value.Exists() {
		data.DeviceGroupId = types.StringValue(value.String())
	} else {
		data.DeviceGroupId = types.StringNull()
	}
	if value := res.Get("deviceRegistration.platformSettings.id"); value.Exists() {
		data.PlatformSettingsId = types.StringValue(value.String())
	} else {
		data.PlatformSettingsId = types.StringNull()
	}
	if value := res.Get("deviceRegistration.licenseCaps"); value.Exists() {
		data.Licenses = helpers.GetStringSet(value.Array())
	} else {
		data.Licenses = types.SetNull(types.StringType)
	}
	if value := res.Get("metadata.containerDetails.id"); value.Exists() {
		data.ContainerId = types.StringValue(value.String())
	} else {
		data.ContainerId = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.type"); value.Exists() {
		data.ContainerType = types.StringValue(value.String())
	} else {
		data.ContainerType = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.name"); value.Exists() {
		data.ContainerName = types.StringValue(value.String())
	} else {
		data.ContainerName = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.role"); value.Exists() {
		data.ContainerRole = types.StringValue(value.String())
	} else {
		data.ContainerRole = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.status"); value.Exists() {
		data.ContainerStatus = types.StringValue(value.String())
	} else {
		data.ContainerStatus = types.StringNull()
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyPartial

// fromBodyPartial reads values from a gjson.Result into a tfstate model. It ignores null attributes in order to
// uncouple the provider from the exact values that the backend API might summon to replace nulls. (Such behavior might
// easily change across versions of the backend API.) For List/Set/Map attributes, the func only updates the
// "managed" elements, instead of all elements.
func (data *ChassisLogicalDevice) fromBodyPartial(ctx context.Context, res gjson.Result) {
	if value := res.Get("type"); value.Exists() && !data.Type.IsNull() {
		data.Type = types.StringValue(value.String())
	} else {
		data.Type = types.StringNull()
	}
	if value := res.Get("metadata.deviceRecordDetails.id"); value.Exists() && !data.DeviceId.IsNull() {
		data.DeviceId = types.StringValue(value.String())
	} else {
		data.DeviceId = types.StringNull()
	}
	if value := res.Get("metadata.deviceRecordDetails.type"); value.Exists() && !data.DeviceType.IsNull() {
		data.DeviceType = types.StringValue(value.String())
	} else {
		data.DeviceType = types.StringNull()
	}
	if value := res.Get("name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("ftdApplicationVersion"); value.Exists() && !data.FtdVersion.IsNull() {
		data.FtdVersion = types.StringValue(value.String())
	} else {
		data.FtdVersion = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv4.ip"); value.Exists() && !data.Ipv4Address.IsNull() {
		data.Ipv4Address = types.StringValue(value.String())
	} else {
		data.Ipv4Address = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv4.mask"); value.Exists() && !data.Ipv4Netmask.IsNull() {
		data.Ipv4Netmask = types.StringValue(value.String())
	} else {
		data.Ipv4Netmask = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv4.gateway"); value.Exists() && !data.Ipv4Gateway.IsNull() {
		data.Ipv4Gateway = types.StringValue(value.String())
	} else {
		data.Ipv4Gateway = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv6.ip"); value.Exists() && !data.Ipv6Address.IsNull() {
		data.Ipv6Address = types.StringValue(value.String())
	} else {
		data.Ipv6Address = types.StringNull()
	}
	if value := res.Get("managementBootstrap.ipv6.prefixLength"); value.Exists() && !data.Ipv6Prefix.IsNull() {
		data.Ipv6Prefix = types.Int64Value(value.Int())
	} else {
		data.Ipv6Prefix = types.Int64Null()
	}
	if value := res.Get("managementBootstrap.ipv6.gateway"); value.Exists() && !data.Ipv6Gateway.IsNull() {
		data.Ipv6Gateway = types.StringValue(value.String())
	} else {
		data.Ipv6Gateway = types.StringNull()
	}
	if value := res.Get("managementBootstrap.searchDomain"); value.Exists() && !data.SearchDomain.IsNull() {
		data.SearchDomain = types.StringValue(value.String())
	} else {
		data.SearchDomain = types.StringNull()
	}
	if value := res.Get("managementBootstrap.fqdn"); value.Exists() && !data.Fqdn.IsNull() {
		data.Fqdn = types.StringValue(value.String())
	} else {
		data.Fqdn = types.StringNull()
	}
	if value := res.Get("managementBootstrap.firewallMode"); value.Exists() && !data.FirewallMode.IsNull() {
		data.FirewallMode = types.StringValue(value.String())
	} else {
		data.FirewallMode = types.StringNull()
	}
	if value := res.Get("managementBootstrap.dnsServers"); value.Exists() && !data.DnsServers.IsNull() {
		data.DnsServers = types.StringValue(value.String())
	} else {
		data.DnsServers = types.StringNull()
	}
	if value := res.Get("adminState"); value.Exists() && !data.AdminState.IsNull() {
		data.AdminState = types.StringValue(value.String())
	} else if data.AdminState.ValueString() != "ENABLED" {
		data.AdminState = types.StringNull()
	}
	if value := res.Get("managementBootstrap.permitExpertMode"); value.Exists() && !data.PermitExpertMode.IsNull() {
		data.PermitExpertMode = types.StringValue(value.String())
	} else {
		data.PermitExpertMode = types.StringNull()
	}
	if value := res.Get("resourceProfile.id"); value.Exists() && !data.ResourceProfileId.IsNull() {
		data.ResourceProfileId = types.StringValue(value.String())
	} else {
		data.ResourceProfileId = types.StringNull()
	}
	if value := res.Get("resourceProfile.name"); value.Exists() && !data.ResourceProfileName.IsNull() {
		data.ResourceProfileName = types.StringValue(value.String())
	} else {
		data.ResourceProfileName = types.StringNull()
	}
	for i := 0; i < len(data.AssignedInterfaces); i++ {
		keys := [...]string{"id"}
		keyValues := [...]string{data.AssignedInterfaces[i].Id.ValueString()}

		parent := &data
		data := (*parent).AssignedInterfaces[i]
		parentRes := &res
		var res gjson.Result

		parentRes.Get("externalPortLink").ForEach(
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
			tflog.Debug(ctx, fmt.Sprintf("removing AssignedInterfaces[%d] = %+v",
				i,
				(*parent).AssignedInterfaces[i],
			))
			(*parent).AssignedInterfaces = slices.Delete((*parent).AssignedInterfaces, i, i+1)
			i--

			continue
		}
		if value := res.Get("id"); value.Exists() && !data.Id.IsNull() {
			data.Id = types.StringValue(value.String())
		} else {
			data.Id = types.StringNull()
		}
		(*parent).AssignedInterfaces[i] = data
	}
	if value := res.Get("deviceRegistration.deviceGroup.id"); value.Exists() && !data.DeviceGroupId.IsNull() {
		data.DeviceGroupId = types.StringValue(value.String())
	} else {
		data.DeviceGroupId = types.StringNull()
	}
	if value := res.Get("deviceRegistration.platformSettings.id"); value.Exists() && !data.PlatformSettingsId.IsNull() {
		data.PlatformSettingsId = types.StringValue(value.String())
	} else {
		data.PlatformSettingsId = types.StringNull()
	}
	if value := res.Get("deviceRegistration.licenseCaps"); value.Exists() && !data.Licenses.IsNull() {
		data.Licenses = helpers.GetStringSet(value.Array())
	} else {
		data.Licenses = types.SetNull(types.StringType)
	}
	if value := res.Get("metadata.containerDetails.id"); value.Exists() {
		data.ContainerId = types.StringValue(value.String())
	} else {
		data.ContainerId = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.type"); value.Exists() {
		data.ContainerType = types.StringValue(value.String())
	} else {
		data.ContainerType = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.name"); value.Exists() {
		data.ContainerName = types.StringValue(value.String())
	} else {
		data.ContainerName = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.role"); value.Exists() {
		data.ContainerRole = types.StringValue(value.String())
	} else {
		data.ContainerRole = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.status"); value.Exists() {
		data.ContainerStatus = types.StringValue(value.String())
	} else {
		data.ContainerStatus = types.StringNull()
	}
}

// End of section. //template:end fromBodyPartial

// Section below is generated&owned by "gen/generator.go". //template:begin fromBodyUnknowns

// fromBodyUnknowns updates the Unknown Computed tfstate values from a JSON.
// Known values are not changed (usual for Computed attributes with UseStateForUnknown or with Default).
func (data *ChassisLogicalDevice) fromBodyUnknowns(ctx context.Context, res gjson.Result) {
	if data.Type.IsUnknown() {
		if value := res.Get("type"); value.Exists() {
			data.Type = types.StringValue(value.String())
		} else {
			data.Type = types.StringNull()
		}
	}
	if data.DeviceId.IsUnknown() {
		if value := res.Get("metadata.deviceRecordDetails.id"); value.Exists() {
			data.DeviceId = types.StringValue(value.String())
		} else {
			data.DeviceId = types.StringNull()
		}
	}
	if data.DeviceType.IsUnknown() {
		if value := res.Get("metadata.deviceRecordDetails.type"); value.Exists() {
			data.DeviceType = types.StringValue(value.String())
		} else {
			data.DeviceType = types.StringNull()
		}
	}
	if value := res.Get("metadata.containerDetails.id"); value.Exists() {
		data.ContainerId = types.StringValue(value.String())
	} else {
		data.ContainerId = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.type"); value.Exists() {
		data.ContainerType = types.StringValue(value.String())
	} else {
		data.ContainerType = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.name"); value.Exists() {
		data.ContainerName = types.StringValue(value.String())
	} else {
		data.ContainerName = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.role"); value.Exists() {
		data.ContainerRole = types.StringValue(value.String())
	} else {
		data.ContainerRole = types.StringNull()
	}
	if value := res.Get("metadata.containerDetails.status"); value.Exists() {
		data.ContainerStatus = types.StringValue(value.String())
	} else {
		data.ContainerStatus = types.StringNull()
	}
}

// End of section. //template:end fromBodyUnknowns
