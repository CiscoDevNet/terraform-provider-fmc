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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcAccessControlPolicyInheritance(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_access_control_policy_inheritance.test", "type"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcAccessControlPolicyInheritancePrerequisitesConfig + testAccFmcAccessControlPolicyInheritanceConfig_all(),
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

const testAccFmcAccessControlPolicyInheritancePrerequisitesConfig = `
resource "fmc_access_control_policy" "test_1" {
  name           = "acp_inherit_parent"
  default_action = "BLOCK"
  categories = [
    { name = "par_Cat1", section = "mandatory" },
    { name = "par_Cat2", section = "mandatory" },
    { name = "par_Cat3", section = "default" },
    { name = "par_Cat4", section = "default" },
  ]
  rules = [
    { name = "par_R1", category_name = "par_Cat1", action = "ALLOW", source_network_literals = [{ value = "10.0.0.1" }] },
    { name = "par_R2", category_name = "par_Cat2", action = "ALLOW", source_network_literals = [{ value = "10.0.0.2" }] },
    { name = "par_R3", section = "mandatory", action = "ALLOW", source_network_literals = [{ value = "10.0.0.3" }] },
    { name = "par_R4", category_name = "par_Cat3", action = "ALLOW", source_network_literals = [{ value = "10.0.0.4" }] },
    { name = "par_R5", category_name = "par_Cat4", action = "ALLOW", source_network_literals = [{ value = "10.0.0.5" }, { value = "10.0.0.6" }] },
    { name = "par_R6", section = "default", action = "ALLOW", source_network_literals = [{ value = "10.0.0.7" }] }
  ]
}

resource "fmc_access_control_policy" "test_2" {
  name           = "acp_inherit_child"
  default_action = "BLOCK"
  categories = [
    { name = "child_Cat1", section = "mandatory" },
    { name = "child_Cat2", section = "mandatory" },
    { name = "child_Cat3", section = "default" },
    { name = "child_Cat4", section = "default" },
  ]
  rules = [
    { name = "child_R1", category_name = "child_Cat1", action = "ALLOW", source_network_literals = [{ value = "10.0.0.1" }] },
    { name = "child_R2", category_name = "child_Cat2", action = "ALLOW", source_network_literals = [{ value = "10.0.0.2" }] },
    { name = "child_R3", section = "mandatory", action = "ALLOW", source_network_literals = [{ value = "10.0.0.3" }] },
    { name = "child_R4", category_name = "child_Cat3", action = "ALLOW", source_network_literals = [{ value = "10.0.0.4" }] },
    { name = "child_R5", category_name = "child_Cat4", action = "ALLOW", source_network_literals = [{ value = "10.0.0.5" }, { value = "10.0.0.6" }] },
    { name = "child_R6", section = "default", action = "ALLOW", source_network_literals = [{ value = "10.0.0.7" }] }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcAccessControlPolicyInheritanceConfig_all() string {
	config := `resource "fmc_access_control_policy_inheritance" "test" {` + "\n"
	config += `	access_control_policy_id = fmc_access_control_policy.test_1.id` + "\n"
	config += `	base_access_control_policy_id = fmc_access_control_policy.test_2.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
