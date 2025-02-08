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

func TestAccDataSourceFmcDeviceVNIInterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vni_interface.test", "vni_id", "42"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vni_interface.test", "multicast_group_address", "224.0.0.24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vni_interface.test", "segment_id", "501"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vni_interface.test", "logical_name", "vni42"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vni_interface.test", "description", "my description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vni_interface.test", "ipv4_static_address", "10.2.2.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vni_interface.test", "ipv4_static_netmask", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vni_interface.test", "ipv6_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vni_interface.test", "ipv6_enable_auto_config", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceVNIInterfacePrerequisitesConfig + testAccDataSourceFmcDeviceVNIInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceVNIInterfacePrerequisitesConfig = `
resource "fmc_device_physical_interface" "test" {
  device_id    = var.device_id
  name         = "GigabitEthernet0/1"
  mode         = "NONE"
  logical_name = "myinterface-0-1"
  mtu          = 9000
}

resource "fmc_device_vtep_policy" "test" {
  device_id = var.device_id
  vteps = [{
    nve_number               = 1
    source_interface_id      = fmc_device_physical_interface.test.id
    neighbor_discovery       = "DEFAULT_MULTICAST_GROUP"
    neighbor_address_literal = "224.0.0.1"
  }]
}

resource "fmc_security_zone" "test" {
  name           = "routed1"
  interface_type = "ROUTED"
}

variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceVNIInterfaceConfig() string {
	config := `resource "fmc_device_vni_interface" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	vni_id = 42` + "\n"
	config += `	multicast_group_address = "224.0.0.24"` + "\n"
	config += `	segment_id = 501` + "\n"
	config += `	nve_number = fmc_device_vtep_policy.test.vteps[0].nve_number` + "\n"
	config += `	enabled = true` + "\n"
	config += `	logical_name = "vni42"` + "\n"
	config += `	description = "my description"` + "\n"
	config += `	security_zone_id = fmc_security_zone.test.id` + "\n"
	config += `	ipv4_static_address = "10.2.2.2"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `	ipv6_enable = true` + "\n"
	config += `	ipv6_enforce_eui = true` + "\n"
	config += `	ipv6_enable_auto_config = true` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		address = "2005::"` + "\n"
	config += `		prefix = 56` + "\n"
	config += `		enforce_eui = false` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_vni_interface" "test" {
			id = fmc_device_vni_interface.test.id
			device_id = fmc_device_physical_interface.test.device_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
