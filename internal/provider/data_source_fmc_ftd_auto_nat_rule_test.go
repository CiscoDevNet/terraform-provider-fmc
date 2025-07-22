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

func TestAccDataSourceFmcFTDAutoNATRule(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ftd_auto_nat_rule.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ftd_auto_nat_rule.test", "nat_type", "STATIC"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcFTDAutoNATRulePrerequisitesConfig + testAccDataSourceFmcFTDAutoNATRuleConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcFTDAutoNATRulePrerequisitesConfig = `
resource "fmc_ftd_nat_policy" "test" {
  name         = "ftd_auto_nat_rule"
  manage_rules = false
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcFTDAutoNATRuleConfig() string {
	config := `resource "fmc_ftd_auto_nat_rule" "test" {` + "\n"
	config += `	ftd_nat_policy_id = fmc_ftd_nat_policy.test.id` + "\n"
	config += `	nat_type = "STATIC"` + "\n"
	config += `	original_network_id = fmc_hosts.test.items.nat_host3.id` + "\n"
	config += `	translated_network_id = fmc_hosts.test.items.nat_host4.id` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_ftd_auto_nat_rule" "test" {
			id = fmc_ftd_auto_nat_rule.test.id
			ftd_nat_policy_id = fmc_ftd_nat_policy.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
