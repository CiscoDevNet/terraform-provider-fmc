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

func TestAccFmcIntrusionPolicyGroupOverride(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_intrusion_policy_group_override.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_intrusion_policy_group_override.test", "default_security_level"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_intrusion_policy_group_override.test", "override_security_level", "LEVEL_2"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcIntrusionPolicyGroupOverridePrerequisitesConfig + testAccFmcIntrusionPolicyGroupOverrideConfig_all(),
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

const testAccFmcIntrusionPolicyGroupOverridePrerequisitesConfig = `
data "fmc_intrusion_policy" "builtin" {
  name = "Balanced Security and Connectivity"
}

resource "fmc_intrusion_policy" "test" {
  name        = "Intrusion Policy Group Override"
  base_policy_id = data.fmc_intrusion_policy.builtin.id
}

resource "fmc_intrusion_rule_group" "test" {
  name = "Intrusion Policy Group Override"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcIntrusionPolicyGroupOverrideConfig_all() string {
	config := `resource "fmc_intrusion_policy_group_override" "test" {` + "\n"
	config += `	intrusion_policy_id = fmc_intrusion_policy.test.id` + "\n"
	config += `	intrusion_rule_group_id = fmc_intrusion_rule_group.test.id` + "\n"
	config += `	override_security_level = "LEVEL_2"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
