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

func TestAccFmcDeviceRedundantInterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" || os.Getenv("TF_VAR_interface_2_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name and TF_VAR_interface_2_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_redundant_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_redundant_interface.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_redundant_interface.test", "logical_name", "my_redundant_interface"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_redundant_interface.test", "description", "my description"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_redundant_interface.test", "mtu", "9000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_redundant_interface.test", "redundant_interface_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_redundant_interface.test", "ipv4_static_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_redundant_interface.test", "ipv4_static_netmask", "24"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceRedundantInterfacePrerequisitesConfig + testAccFmcDeviceRedundantInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceRedundantInterfacePrerequisitesConfig + testAccFmcDeviceRedundantInterfaceConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcDeviceRedundantInterfacePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" { default = null } // tests will set $TF_VAR_interface_name
variable "interface_2_name" { default = null } // tests will set $TF_VAR_interface_2_name

data "fmc_device_physical_interface" "test" {
  device_id = var.device_id
  name      = var.interface_name
}

data "fmc_device_physical_interface" "test_2" {
  device_id = var.device_id
  name      = var.interface_2_name
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceRedundantInterfaceConfig_minimum() string {
	config := `resource "fmc_device_redundant_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	redundant_interface_id = 1` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceRedundantInterfaceConfig_all() string {
	config := `resource "fmc_device_redundant_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "my_redundant_interface"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	description = "my description"` + "\n"
	config += `	mtu = 9000` + "\n"
	config += `	redundant_interface_id = 1` + "\n"
	config += `	primary_interface_id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `	primary_interface_name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	primary_interface_type = data.fmc_device_physical_interface.test.type` + "\n"
	config += `	secondary_interface_id = data.fmc_device_physical_interface.test_2.id` + "\n"
	config += `	secondary_interface_name = data.fmc_device_physical_interface.test_2.name` + "\n"
	config += `	secondary_interface_type = data.fmc_device_physical_interface.test_2.type` + "\n"
	config += `	ipv4_static_address = "10.1.1.1"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
