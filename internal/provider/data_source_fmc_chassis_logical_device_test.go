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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcChassisLogicalDevice(t *testing.T) {
	if os.Getenv("TF_VAR_chassis_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_chassis_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_chassis_logical_device.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "name", "my-logical-device"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "ftd_version", "7.6.0.113"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "ipv4_address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "ipv4_netmask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "ipv4_gateway", "10.10.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "ipv6_address", "2001:db8::10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "ipv6_prefix_length", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "ipv6_gateway", "2001:db8::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "search_domain", "cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "fqdn", "my_logical_device.cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "firewall_mode", "ROUTED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "dns_servers", "10.123.10.12,10.123.10.14"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "permit_expert_mode", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "resource_profile_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "resource_profile_name", "my_resource_profile"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "assigned_interfaces.0.id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "device_group_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "access_policy_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_logical_device.test", "platform_settings_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcChassisLogicalDeviceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcChassisLogicalDeviceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcChassisLogicalDeviceConfig() string {
	config := `resource "fmc_chassis_logical_device" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	name = "my-logical-device"` + "\n"
	config += `	ftd_version = "7.6.0.113"` + "\n"
	config += `	ipv4_address = "10.10.10.10"` + "\n"
	config += `	ipv4_netmask = "255.255.255.0"` + "\n"
	config += `	ipv4_gateway = "10.10.10.1"` + "\n"
	config += `	ipv6_address = "2001:db8::10"` + "\n"
	config += `	ipv6_prefix_length = 64` + "\n"
	config += `	ipv6_gateway = "2001:db8::1"` + "\n"
	config += `	search_domain = "cisco.com"` + "\n"
	config += `	fqdn = "my_logical_device.cisco.com"` + "\n"
	config += `	firewall_mode = "ROUTED"` + "\n"
	config += `	dns_servers = "10.123.10.12,10.123.10.14"` + "\n"
	config += `	device_password = "my_password"` + "\n"
	config += `	permit_expert_mode = ""` + "\n"
	config += `	resource_profile_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	resource_profile_name = "my_resource_profile"` + "\n"
	config += `	assigned_interfaces = [{` + "\n"
	config += `		id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	}]` + "\n"
	config += `	device_group_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	access_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	license_capabilities = ["MALWARE"]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_chassis_logical_device" "test" {
			id = fmc_chassis_logical_device.test.id
			chassis_id = var.chassis_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcChassisLogicalDeviceConfig() string {
	config := `resource "fmc_chassis_logical_device" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	name = "my-logical-device"` + "\n"
	config += `	ftd_version = "7.6.0.113"` + "\n"
	config += `	ipv4_address = "10.10.10.10"` + "\n"
	config += `	ipv4_netmask = "255.255.255.0"` + "\n"
	config += `	ipv4_gateway = "10.10.10.1"` + "\n"
	config += `	ipv6_address = "2001:db8::10"` + "\n"
	config += `	ipv6_prefix_length = 64` + "\n"
	config += `	ipv6_gateway = "2001:db8::1"` + "\n"
	config += `	search_domain = "cisco.com"` + "\n"
	config += `	fqdn = "my_logical_device.cisco.com"` + "\n"
	config += `	firewall_mode = "ROUTED"` + "\n"
	config += `	dns_servers = "10.123.10.12,10.123.10.14"` + "\n"
	config += `	device_password = "my_password"` + "\n"
	config += `	permit_expert_mode = ""` + "\n"
	config += `	resource_profile_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	resource_profile_name = "my_resource_profile"` + "\n"
	config += `	assigned_interfaces = [{` + "\n"
	config += `		id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	}]` + "\n"
	config += `	device_group_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	access_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	license_capabilities = ["MALWARE"]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_chassis_logical_device" "test" {
			chassis_id = var.chassis_id
			name = fmc_chassis_logical_device.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
