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

func TestAccDataSourceFmcDeviceVTIInterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_vti_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_vti_interface.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "logical_name", "my_vti_interface"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "description", "My VTI"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "priority", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "tunnel_id", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "tunnel_mode", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "ipv4_address", "10.10.10.10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "ipv4_netmask", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "ip_based_monitoring", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "ip_based_monitoring_type", "PEER_IPV4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_vti_interface.test", "ip_based_monitoring_peer_ip", "10.10.10.100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceVTIInterfacePrerequisitesConfig + testAccDataSourceFmcDeviceVTIInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceVTIInterfacePrerequisitesConfig + testAccNamedDataSourceFmcDeviceVTIInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceVTIInterfacePrerequisitesConfig = `
resource "fmc_device_physical_interface" "test" {
  device_id           = var.device_id
  name                = "GigabitEthernet0/1"
  mode                = "NONE"
  logical_name        = "myinterface-0-1"
  ipv4_static_address = "10.12.1.1"
  ipv4_static_netmask = "255.255.255.240"
}

resource "fmc_security_zone" "test" {
  name           = "my_sec_zone_vti"
  interface_type = "ROUTED"
}

variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceVTIInterfaceConfig() string {
	config := `resource "fmc_device_vti_interface" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	tunnel_type = "STATIC"` + "\n"
	config += `	logical_name = "my_vti_interface"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	description = "My VTI"` + "\n"
	config += `	security_zone_id = fmc_security_zone.test.id` + "\n"
	config += `	priority = 100` + "\n"
	config += `	tunnel_id = 100` + "\n"
	config += `	tunnel_source_interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	tunnel_source_interface_name = fmc_device_physical_interface.test.name` + "\n"
	config += `	tunnel_mode = "ipv4"` + "\n"
	config += `	ipv4_address = "10.10.10.10"` + "\n"
	config += `	ipv4_netmask = "24"` + "\n"
	config += `	ip_based_monitoring = true` + "\n"
	config += `	ip_based_monitoring_type = "PEER_IPV4"` + "\n"
	config += `	ip_based_monitoring_peer_ip = "10.10.10.100"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_vti_interface" "test" {
			id = fmc_device_vti_interface.test.id
			device_id = fmc_device_physical_interface.test.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceVTIInterfaceConfig() string {
	config := `resource "fmc_device_vti_interface" "test" {` + "\n"
	config += `	device_id = fmc_device_physical_interface.test.device_id` + "\n"
	config += `	tunnel_type = "STATIC"` + "\n"
	config += `	logical_name = "my_vti_interface"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	description = "My VTI"` + "\n"
	config += `	security_zone_id = fmc_security_zone.test.id` + "\n"
	config += `	priority = 100` + "\n"
	config += `	tunnel_id = 100` + "\n"
	config += `	tunnel_source_interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	tunnel_source_interface_name = fmc_device_physical_interface.test.name` + "\n"
	config += `	tunnel_mode = "ipv4"` + "\n"
	config += `	ipv4_address = "10.10.10.10"` + "\n"
	config += `	ipv4_netmask = "24"` + "\n"
	config += `	ip_based_monitoring = true` + "\n"
	config += `	ip_based_monitoring_type = "PEER_IPV4"` + "\n"
	config += `	ip_based_monitoring_peer_ip = "10.10.10.100"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_vti_interface" "test" {
			device_id = fmc_device_physical_interface.test.device_id
			name = fmc_device_vti_interface.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
