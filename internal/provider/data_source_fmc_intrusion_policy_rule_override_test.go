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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcIntrusionPolicyRuleOverride(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_intrusion_policy_rule_override.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_intrusion_policy_rule_override.test", "default_state"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_intrusion_policy_rule_override.test", "override_state", "BLOCK"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcIntrusionPolicyRuleOverridePrerequisitesConfig + testAccDataSourceFmcIntrusionPolicyRuleOverrideConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcIntrusionPolicyRuleOverridePrerequisitesConfig = `
data "fmc_intrusion_policy" "builtin" {
  name = "Balanced Security and Connectivity"
}

resource "fmc_intrusion_policy" "test" {
  name        = "Intrusion Policy Rule Override"
  base_policy_id = data.fmc_intrusion_policy.builtin.id
}

resource "fmc_intrusion_rule_group" "test" {
  name = "Intrusion Policy Rule Override"
}

resource "fmc_intrusion_rule" "test" {
  rule_data = "alert icmp any any -> any any ( sid:10000303; gid:2000; msg:\"CUSTOM RULE3\"; classtype:icmp-event; rev:1; )"
  rule_groups = [
    { id = fmc_intrusion_rule_group.test.id }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcIntrusionPolicyRuleOverrideConfig() string {
	config := `resource "fmc_intrusion_policy_rule_override" "test" {` + "\n"
	config += `	intrusion_policy_id = fmc_intrusion_policy.test.id` + "\n"
	config += `	intrusion_rule_id = fmc_intrusion_rule.test.id` + "\n"
	config += `	override_state = "BLOCK"` + "\n"
	config += `	rule_data = fmc_intrusion_rule.test.rule_data` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_intrusion_policy_rule_override" "test" {
			id = fmc_intrusion_policy_rule_override.test.id
			intrusion_policy_id = fmc_intrusion_policy.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
