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

func TestAccDataSourceFmcDeviceEtherChannelInterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_id") == "" || os.Getenv("FMC_DEVICE_ETHERCHANNEL_INTERFACE") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_id and FMC_DEVICE_ETHERCHANNEL_INTERFACE")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_etherchannel_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_etherchannel_interface.test", "is_multi_instance"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_etherchannel_interface.test", "logical_name", "myinterface-0-1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_etherchannel_interface.test", "description", "my description"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_etherchannel_interface.test", "mode", "NONE"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_etherchannel_interface.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_etherchannel_interface.test", "mtu", "9000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_etherchannel_interface.test", "ether_channel_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_etherchannel_interface.test", "ipv4_static_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_etherchannel_interface.test", "ipv4_static_netmask", "24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceEtherChannelInterfacePrerequisitesConfig + testAccDataSourceFmcDeviceEtherChannelInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceEtherChannelInterfacePrerequisitesConfig + testAccNamedDataSourceFmcDeviceEtherChannelInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceEtherChannelInterfacePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_id" { default = null } // tests will set $TF_VAR_interface_id

data "fmc_device_physical_interface" "test" {
  device_id = var.device_id
  id = var.interface_id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceEtherChannelInterfaceConfig() string {
	config := `resource "fmc_device_etherchannel_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "myinterface-0-1"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	description = "my description"` + "\n"
	config += `	mode = "NONE"` + "\n"
	config += `	mtu = 9000` + "\n"
	config += `	ether_channel_id = "1"` + "\n"
	config += `	selected_interfaces = [{` + "\n"
	config += `		id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `		name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_address = "10.1.1.1"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_etherchannel_interface" "test" {
			id = fmc_device_etherchannel_interface.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceEtherChannelInterfaceConfig() string {
	config := `resource "fmc_device_etherchannel_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "myinterface-0-1"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	description = "my description"` + "\n"
	config += `	mode = "NONE"` + "\n"
	config += `	mtu = 9000` + "\n"
	config += `	ether_channel_id = "1"` + "\n"
	config += `	selected_interfaces = [{` + "\n"
	config += `		id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `		name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	}]` + "\n"
	config += `	ipv4_static_address = "10.1.1.1"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_etherchannel_interface" "test" {
			device_id = var.device_id
			name = fmc_device_etherchannel_interface.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
