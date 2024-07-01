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
func TestAccFmcDevicePhysicalInterface(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "device_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "mode", "NONE"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "name", "GigabitEthernet0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "logical_name", "myinterface-0-1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "description", "my description"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "mtu", "9000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv4_static_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv4_static_netmask", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv6_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv6_enforce_eui", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv6_enable_auto_config", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv6_enable_dhcp_address", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv6_enable_dhcp_nonaddress", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv6_enable_ra", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv6_addresses.0.address", "2004::"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_physical_interface.test", "ipv6_addresses.0.prefix", "124"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDevicePhysicalInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDevicePhysicalInterfaceConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
func testAccFmcDevicePhysicalInterfaceConfig_minimum() string {
	config := `resource "fmc_device_physical_interface" "test" {` + "\n"
	config += `	device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	mode = "NONE"` + "\n"
	config += `	name = "GigabitEthernet0/1"` + "\n"
	config += `	logical_name = "iface_minimum"` + "\n"
	config += `	management_only = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccFmcDevicePhysicalInterfaceConfig_all() string {
	config := `resource "fmc_device_physical_interface" "test" {` + "\n"
	config += `	device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	mode = "NONE"` + "\n"
	config += `	name = "GigabitEthernet0/1"` + "\n"
	config += `	logical_name = "myinterface-0-1"` + "\n"
	config += `	description = "my description"` + "\n"
	config += `	management_only = false` + "\n"
	config += `	mtu = 9000` + "\n"
	config += `	ipv4_static_address = "10.1.1.1"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `	ipv6_enable = true` + "\n"
	config += `	ipv6_enforce_eui = true` + "\n"
	config += `	ipv6_enable_auto_config = true` + "\n"
	config += `	ipv6_enable_dhcp_address = true` + "\n"
	config += `	ipv6_enable_dhcp_nonaddress = true` + "\n"
	config += `	ipv6_enable_ra = false` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `	  address = "2004::"` + "\n"
	config += `	  prefix = "124"` + "\n"
	config += `	  enforce_eui = "true"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
