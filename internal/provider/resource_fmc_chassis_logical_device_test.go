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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcChassisLogicalDevice(t *testing.T) {
	if os.Getenv("TF_VAR_chassis_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_chassis_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_logical_device.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_logical_device.test", "device_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_logical_device.test", "device_type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "name", "my-logical-device"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "ftd_version", "7.6.0.113"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "ipv4_address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "ipv4_netmask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "ipv4_gateway", "10.10.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "ipv6_address", "2001:db8::10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "ipv6_prefix", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "ipv6_gateway", "2001:db8::1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "search_domain", "cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "fqdn", "my_logical_device.cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "firewall_mode", "ROUTED"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "dns_servers", "10.123.10.12,10.123.10.14"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "admin_state", "ENABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "permit_expert_mode", "yes"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "resource_profile_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "resource_profile_name", "my_resource_profile"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "assigned_interfaces.0.id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "device_group_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_logical_device.test", "platform_settings_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_logical_device.test", "container_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_logical_device.test", "container_type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_logical_device.test", "container_name"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_logical_device.test", "container_role"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_logical_device.test", "container_status"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcChassisLogicalDeviceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcChassisLogicalDeviceConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcChassisLogicalDeviceConfig_minimum() string {
	config := `resource "fmc_chassis_logical_device" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	name = "my-logical-device"` + "\n"
	config += `	ftd_version = "7.6.0.113"` + "\n"
	config += `	firewall_mode = "ROUTED"` + "\n"
	config += `	device_password = "my_password"` + "\n"
	config += `	resource_profile_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	resource_profile_name = "my_resource_profile"` + "\n"
	config += `	assigned_interfaces = [{` + "\n"
	config += `	}]` + "\n"
	config += `	access_control_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcChassisLogicalDeviceConfig_all() string {
	config := `resource "fmc_chassis_logical_device" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	name = "my-logical-device"` + "\n"
	config += `	ftd_version = "7.6.0.113"` + "\n"
	config += `	ipv4_address = "10.10.10.10"` + "\n"
	config += `	ipv4_netmask = "255.255.255.0"` + "\n"
	config += `	ipv4_gateway = "10.10.10.1"` + "\n"
	config += `	ipv6_address = "2001:db8::10"` + "\n"
	config += `	ipv6_prefix = 64` + "\n"
	config += `	ipv6_gateway = "2001:db8::1"` + "\n"
	config += `	search_domain = "cisco.com"` + "\n"
	config += `	fqdn = "my_logical_device.cisco.com"` + "\n"
	config += `	firewall_mode = "ROUTED"` + "\n"
	config += `	dns_servers = "10.123.10.12,10.123.10.14"` + "\n"
	config += `	device_password = "my_password"` + "\n"
	config += `	admin_state = "ENABLED"` + "\n"
	config += `	permit_expert_mode = "yes"` + "\n"
	config += `	resource_profile_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	resource_profile_name = "my_resource_profile"` + "\n"
	config += `	assigned_interfaces = [{` + "\n"
	config += `		id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	}]` + "\n"
	config += `	device_group_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	access_control_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	platform_settings_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	licenses = ["MALWARE"]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
