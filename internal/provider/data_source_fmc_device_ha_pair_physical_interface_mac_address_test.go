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

func TestAccDataSourceFmcDeviceHAPairPhysicalInterfaceMACAddress(t *testing.T) {
	if os.Getenv("TF_VAR_device_ha_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_ha_id and TF_VAR_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_ha_pair_physical_interface_mac_address.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair_physical_interface_mac_address.test", "active_mac_address", "c460.15e4.0edd"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair_physical_interface_mac_address.test", "standby_mac_address", "c460.15e4.0ed0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceHAPairPhysicalInterfaceMACAddressPrerequisitesConfig + testAccDataSourceFmcDeviceHAPairPhysicalInterfaceMACAddressConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceHAPairPhysicalInterfaceMACAddressPrerequisitesConfig + testAccNamedDataSourceFmcDeviceHAPairPhysicalInterfaceMACAddressConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceHAPairPhysicalInterfaceMACAddressPrerequisitesConfig = `
data "fmc_device_physical_interface" "test" {
  device_id = var.device_id
  name      = var.interface_name
}

variable "device_ha_id" { default = null } // tests will set $TF_VAR_device_ha_id
variable "interface_name" {default = null} // tests will set $TF_VAR_interface_name
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceHAPairPhysicalInterfaceMACAddressConfig() string {
	config := `resource "fmc_device_ha_pair_physical_interface_mac_address" "test" {` + "\n"
	config += `	device_id = var.device_ha_id` + "\n"
	config += `	interface_name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	interface_id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `	active_mac_address = "c460.15e4.0edd"` + "\n"
	config += `	standby_mac_address = "c460.15e4.0ed0"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ha_pair_physical_interface_mac_address" "test" {
			id = fmc_device_ha_pair_physical_interface_mac_address.test.id
			device_id = var.device_ha_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceHAPairPhysicalInterfaceMACAddressConfig() string {
	config := `resource "fmc_device_ha_pair_physical_interface_mac_address" "test" {` + "\n"
	config += `	device_id = var.device_ha_id` + "\n"
	config += `	interface_name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	interface_id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `	active_mac_address = "c460.15e4.0edd"` + "\n"
	config += `	standby_mac_address = "c460.15e4.0ed0"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ha_pair_physical_interface_mac_address" "test" {
			device_id = var.device_ha_id
			interface_name = fmc_device_ha_pair_physical_interface_mac_address.test.interface_name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
