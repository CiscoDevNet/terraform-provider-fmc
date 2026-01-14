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

func TestAccDataSourceFmcChassisEtherChannelInterface(t *testing.T) {
	if os.Getenv("TF_VAR_chassis_id") == "" || os.Getenv("TF_VAR_chassis_interface_id") == "" || os.Getenv("FMC_CHASSIS_ETHERCHANNEL_INTERFACE") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_chassis_id and TF_VAR_chassis_interface_id and FMC_CHASSIS_ETHERCHANNEL_INTERFACE")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_chassis_etherchannel_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_chassis_etherchannel_interface.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_etherchannel_interface.test", "ether_channel_id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_etherchannel_interface.test", "port_type", "DATA"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_etherchannel_interface.test", "admin_state", "ENABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_etherchannel_interface.test", "lacp_mode", "ACTIVE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_etherchannel_interface.test", "lacp_rate", "DEFAULT"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcChassisEtherChannelInterfacePrerequisitesConfig + testAccDataSourceFmcChassisEtherChannelInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcChassisEtherChannelInterfacePrerequisitesConfig + testAccNamedDataSourceFmcChassisEtherChannelInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcChassisEtherChannelInterfacePrerequisitesConfig = `
variable "chassis_id" { default = null } // tests will set $TF_VAR_chassis_id
variable "chassis_interface_id" { default = null } // tests will set $TF_VAR_chassis_interface_id

data "fmc_chassis_physical_interface" "test" {
  chassis_id = var.chassis_id
  id = var.chassis_interface_id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcChassisEtherChannelInterfaceConfig() string {
	config := `resource "fmc_chassis_etherchannel_interface" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	ether_channel_id = 10` + "\n"
	config += `	port_type = "DATA"` + "\n"
	config += `	admin_state = "ENABLED"` + "\n"
	config += `	selected_interfaces = [{` + "\n"
	config += `		id = data.fmc_chassis_physical_interface.test.id` + "\n"
	config += `		name = data.fmc_chassis_physical_interface.test.name` + "\n"
	config += `	}]` + "\n"
	config += `	lacp_mode = "ACTIVE"` + "\n"
	config += `	lacp_rate = "DEFAULT"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_chassis_etherchannel_interface" "test" {
			id = fmc_chassis_etherchannel_interface.test.id
			chassis_id = var.chassis_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcChassisEtherChannelInterfaceConfig() string {
	config := `resource "fmc_chassis_etherchannel_interface" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	ether_channel_id = 10` + "\n"
	config += `	port_type = "DATA"` + "\n"
	config += `	admin_state = "ENABLED"` + "\n"
	config += `	selected_interfaces = [{` + "\n"
	config += `		id = data.fmc_chassis_physical_interface.test.id` + "\n"
	config += `		name = data.fmc_chassis_physical_interface.test.name` + "\n"
	config += `	}]` + "\n"
	config += `	lacp_mode = "ACTIVE"` + "\n"
	config += `	lacp_rate = "DEFAULT"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_chassis_etherchannel_interface" "test" {
			chassis_id = var.chassis_id
			name = fmc_chassis_etherchannel_interface.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
