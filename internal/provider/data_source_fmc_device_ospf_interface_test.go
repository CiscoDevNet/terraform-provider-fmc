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

func TestAccDataSourceFmcDeviceOSPFInterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_ospf_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "default_cost", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "priority", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "mtu_missmatch_ignore", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "hello_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "transmit_delay", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "retransmit_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "dead_interval", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "hello_multiplier", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "point_to_point", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ospf_interface.test", "bfd", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceOSPFInterfacePrerequisitesConfig + testAccDataSourceFmcDeviceOSPFInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceOSPFInterfacePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" {default = null} // tests will set $TF_VAR_interface_name

resource "fmc_device_physical_interface" "test" {
  device_id = var.device_id
  name      = var.interface_name
  mode      = "NONE"
  logical_name = "OSPF_Interface"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceOSPFInterfaceConfig() string {
	config := `resource "fmc_device_ospf_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	default_cost = 10` + "\n"
	config += `	priority = 1` + "\n"
	config += `	mtu_missmatch_ignore = false` + "\n"
	config += `	hello_interval = 10` + "\n"
	config += `	transmit_delay = 1` + "\n"
	config += `	retransmit_interval = 5` + "\n"
	config += `	dead_interval = 40` + "\n"
	config += `	hello_multiplier = 4` + "\n"
	config += `	point_to_point = false` + "\n"
	config += `	bfd = false` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ospf_interface" "test" {
			id = fmc_device_ospf_interface.test.id
			device_id = var.device_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
