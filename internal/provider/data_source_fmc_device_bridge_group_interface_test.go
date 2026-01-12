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
	"slices"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcDeviceBridgeGroupInterface(t *testing.T) {
	if v := os.Getenv("FMC_VERSION"); v != "" && slices.Contains([]string{"7.7"}, v) {
		t.Skip("skipping test for FMC version " + v)
	}
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_bridge_group_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_bridge_group_interface.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "logical_name", "my_bridge_group_interface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "description", "My Bridge Group Interface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "bridge_group_id", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "ipv4_static_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "ipv4_static_netmask", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "ipv6_addresses.0.address", "2004::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "ipv6_addresses.0.prefix", "64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "arp_table_entries.0.mac_address", "0123.4567.89ab"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "arp_table_entries.0.ip_address", "10.1.1.10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bridge_group_interface.test", "arp_table_entries.0.enable_alias", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceBridgeGroupInterfacePrerequisitesConfig + testAccDataSourceFmcDeviceBridgeGroupInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceBridgeGroupInterfacePrerequisitesConfig + testAccNamedDataSourceFmcDeviceBridgeGroupInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceBridgeGroupInterfacePrerequisitesConfig = `
data "fmc_device_physical_interface" "test" {
  device_id = var.device_id
  name      = var.interface_name
}

variable "device_id" { default = null } // tests will set $TF_VAR_device_ha_id
variable "interface_name" {default = null} // tests will set $TF_VAR_interface_name
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceBridgeGroupInterfaceConfig() string {
	config := `resource "fmc_device_bridge_group_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "my_bridge_group_interface"` + "\n"
	config += `	description = "My Bridge Group Interface"` + "\n"
	config += `	bridge_group_id = 100` + "\n"
	config += `	selected_interfaces = [{` + "\n"
	config += `		id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `		name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_address = "10.1.1.1"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		address = "2004::1"` + "\n"
	config += `		prefix = "64"` + "\n"
	config += `	}]` + "\n"
	config += `	arp_table_entries = [{` + "\n"
	config += `		mac_address = "0123.4567.89ab"` + "\n"
	config += `		ip_address = "10.1.1.10"` + "\n"
	config += `		enable_alias = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bridge_group_interface" "test" {
			id = fmc_device_bridge_group_interface.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceBridgeGroupInterfaceConfig() string {
	config := `resource "fmc_device_bridge_group_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "my_bridge_group_interface"` + "\n"
	config += `	description = "My Bridge Group Interface"` + "\n"
	config += `	bridge_group_id = 100` + "\n"
	config += `	selected_interfaces = [{` + "\n"
	config += `		id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `		name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_address = "10.1.1.1"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		address = "2004::1"` + "\n"
	config += `		prefix = "64"` + "\n"
	config += `	}]` + "\n"
	config += `	arp_table_entries = [{` + "\n"
	config += `		mac_address = "0123.4567.89ab"` + "\n"
	config += `		ip_address = "10.1.1.10"` + "\n"
	config += `		enable_alias = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bridge_group_interface" "test" {
			device_id = var.device_id
			name = fmc_device_bridge_group_interface.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
