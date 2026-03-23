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
resource "fmc_access_control_policy" "test" {
  name           = "acp_inherit"
  default_action = "BLOCK"
  categories = [
    { name = "Cat1", section = "mandatory" },
    { name = "Cat2", section = "mandatory" },
    { name = "Cat3", section = "default" },
    { name = "Cat4", section = "default" },
  ]
  rules = [
    { name = "Rule1", category_name = "Cat1", action = "ALLOW", source_network_literals = [{ value = "10.0.0.1" }] },
    { name = "Rule2", category_name = "Cat2", action = "ALLOW", source_network_literals = [{ value = "10.0.0.2" }] },
    { name = "Rule3", section = "mandatory", action = "ALLOW", source_network_literals = [{ value = "10.0.0.3" }] },
    { name = "Rule4", category_name = "Cat3", action = "ALLOW", source_network_literals = [{ value = "10.0.0.4" }] },
    { name = "Rule5", category_name = "Cat4", action = "ALLOW", source_network_literals = [{ value = "10.0.0.5" }, { value = "10.0.0.6" }] },
    { name = "Rule6", section = "default", action = "ALLOW", source_network_literals = [{ value = "10.0.0.7" }] }
  ]
}

resource "fmc_access_control_policy" "test_base" {
  name           = "acp_inherit_base"
  default_action = "BLOCK"
  categories = [
    { name = "base_Cat1", section = "mandatory" },
    { name = "base_Cat2", section = "mandatory" },
    { name = "base_Cat3", section = "default" },
    { name = "base_Cat4", section = "default" },
  ]
  rules = [
    { name = "base_Rule1", category_name = "base_Cat1", action = "ALLOW", source_network_literals = [{ value = "10.1.0.1" }] },
    { name = "base_Rule2", category_name = "base_Cat2", action = "ALLOW", source_network_literals = [{ value = "10.1.0.2" }] },
    { name = "base_Rule3", section = "mandatory", action = "ALLOW", source_network_literals = [{ value = "10.1.0.3" }] },
    { name = "base_Rule4", category_name = "base_Cat3", action = "ALLOW", source_network_literals = [{ value = "10.1.0.4" }] },
    { name = "base_Rule5", category_name = "base_Cat4", action = "ALLOW", source_network_literals = [{ value = "10.1.0.5" }, { value = "10.1.0.6" }] },
    { name = "base_Rule6", section = "default", action = "ALLOW", source_network_literals = [{ value = "10.1.0.7" }] }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcAccessControlPolicyInheritanceConfig_all() string {
	config := `resource "fmc_access_control_policy_inheritance" "test" {` + "\n"
	config += `	access_control_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `	base_access_control_policy_id = fmc_access_control_policy.test_base.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
