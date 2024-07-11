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
func TestAccDataSourceFmcDeviceSubinterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "index", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "vlan_id", "4094"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "logical_name", "mysubinterface-0-1.7"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "description", "my description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv4_static_address", "10.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv4_static_netmask", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv6_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv6_enforce_eui", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv6_enable_auto_config", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv6_enable_dhcp_address", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv6_enable_dhcp_nonaddress", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv6_enable_ra", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv6_addresses.0.address", "2005::"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_subinterface.test", "ipv6_addresses.0.prefix", "124"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceSubinterfacePrerequisitesConfig + testAccDataSourceFmcDeviceSubinterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceFmcDeviceSubinterfacePrerequisitesConfig = `
resource "fmc_device_physical_interface" "test" {
  device_id    = var.device_id
  name         = "GigabitEthernet0/1"
  mode         = "NONE"
  logical_name = "myif-0-1"
  mtu          = 9000
}

resource "fmc_security_zone" "test" {
  name           = "routed1"
  interface_mode = "ROUTED"
}

variable "device_id" { default = null } // tests will set $TF_VAR_device_id

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceFmcDeviceSubinterfaceConfig() string {
	config := `resource "fmc_device_subinterface" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	enabled = true` + "\n"
	config += `	index = 7` + "\n"
	config += `	vlan_id = 4094` + "\n"
	config += `	logical_name = "mysubinterface-0-1.7"` + "\n"
	config += `	description = "my description"` + "\n"
	config += `	management_only = false` + "\n"
	config += `	security_zone_id = fmc_security_zone.test.id` + "\n"
	config += `	ipv4_static_address = "10.2.2.2"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `	ipv6_enable = true` + "\n"
	config += `	ipv6_enforce_eui = true` + "\n"
	config += `	ipv6_enable_auto_config = true` + "\n"
	config += `	ipv6_enable_dhcp_address = true` + "\n"
	config += `	ipv6_enable_dhcp_nonaddress = true` + "\n"
	config += `	ipv6_enable_ra = false` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `	  address = "2005::"` + "\n"
	config += `	  prefix = "124"` + "\n"
	config += `	  enforce_eui = "true"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_subinterface" "test" {
			id = fmc_device_subinterface.test.id
			device_id = fmc_device_physical_interface.test.device_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
