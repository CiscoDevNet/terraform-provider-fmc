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
	checks = append(checks, resource.TestCheckResourceAttr("fmc_intrusion_policy_group_override.test", "intrusion_policy_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_intrusion_policy_group_override.test", "intrusion_rule_group_id", "76d24097-41c4-4558-a4d0-a8c07ac08471"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_intrusion_policy_group_override.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_intrusion_policy_group_override.test", "default_security_level"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_intrusion_policy_group_override.test", "override_security_level", "LEVEL_3"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcIntrusionPolicyGroupOverrideConfig_all(),
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcIntrusionPolicyGroupOverrideConfig_all() string {
	config := `resource "fmc_intrusion_policy_group_override" "test" {` + "\n"
	config += `	intrusion_policy_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	intrusion_rule_group_id = "76d24097-41c4-4558-a4d0-a8c07ac08471"` + "\n"
	config += `	override_security_level = "LEVEL_3"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
