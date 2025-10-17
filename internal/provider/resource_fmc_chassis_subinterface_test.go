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

func TestAccFmcChassisSubinterface(t *testing.T) {
	if os.Getenv("TF_VAR_chassis_id") == "" || os.Getenv("TF_VAR_chassis_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_chassis_id and TF_VAR_chassis_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_subinterface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_chassis_subinterface.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_subinterface.test", "sub_interface_id", "7"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_subinterface.test", "vlan_id", "4094"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_chassis_subinterface.test", "port_type", "DATA"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcChassisSubinterfacePrerequisitesConfig + testAccFmcChassisSubinterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcChassisSubinterfacePrerequisitesConfig + testAccFmcChassisSubinterfaceConfig_all(),
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

const testAccFmcChassisSubinterfacePrerequisitesConfig = `
variable "chassis_id" { default = null } // tests will set $TF_VAR_chassis_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcChassisSubinterfaceConfig_minimum() string {
	config := `resource "fmc_chassis_subinterface" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	interface_name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	interface_id = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	sub_interface_id = 7` + "\n"
	config += `	vlan_id = 4094` + "\n"
	config += `	port_type = "DATA"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcChassisSubinterfaceConfig_all() string {
	config := `resource "fmc_chassis_subinterface" "test" {` + "\n"
	config += `	chassis_id = var.chassis_id` + "\n"
	config += `	interface_name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	interface_id = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	sub_interface_id = 7` + "\n"
	config += `	vlan_id = 4094` + "\n"
	config += `	port_type = "DATA"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
