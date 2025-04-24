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

func TestAccDataSourceFmcChassisPhysicalInterface(t *testing.T) {
	if os.Getenv("TF_VAR_chassis_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_chassis_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_chassis_physical_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_physical_interface.test", "name", "Ethernet1/1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_physical_interface.test", "port_type", "DATA"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_chassis_physical_interface.test", "admin_state", "ENABLED"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcChassisPhysicalInterfacePrerequisitesConfig + testAccDataSourceFmcChassisPhysicalInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcChassisPhysicalInterfacePrerequisitesConfig + testAccNamedDataSourceFmcChassisPhysicalInterfaceConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcChassisPhysicalInterfacePrerequisitesConfig = `
variable "chassis_id" { default = null } // tests will set $TF_VAR_chassis_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcChassisPhysicalInterfaceConfig() string {
	config := `resource "fmc_chassis_physical_interface" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	name = "Ethernet1/1"` + "\n"
	config += `	port_type = "DATA"` + "\n"
	config += `	admin_state = "ENABLED"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_chassis_physical_interface" "test" {
			id = fmc_chassis_physical_interface.test.id
			chassis_id = var.chassis_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcChassisPhysicalInterfaceConfig() string {
	config := `resource "fmc_chassis_physical_interface" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	name = "Ethernet1/1"` + "\n"
	config += `	port_type = "DATA"` + "\n"
	config += `	admin_state = "ENABLED"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_chassis_physical_interface" "test" {
			chassis_id = var.chassis_id
			name = fmc_chassis_physical_interface.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
