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

func TestAccDataSourceFmcIntrusionRule(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_intrusion_rule.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_intrusion_rule.test", "rule_data", "alert icmp any any -> any any ( sid:10000301; gid:2000; msg:\"CUSTOM RULE1\"; classtype:icmp-event; rev:1; )"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_intrusion_rule.test", "rule_groups.0.id", "12da34567890-1234-5678-90ab-cdef12345678"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_intrusion_rule.test", "rule_groups.0.name", "Custom Rule Group 1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_intrusion_rule.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcIntrusionRuleConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcIntrusionRuleConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcIntrusionRuleConfig() string {
	config := `resource "fmc_intrusion_rule" "test" {` + "\n"
	config += `	rule_data = "alert icmp any any -> any any ( sid:10000301; gid:2000; msg:\"CUSTOM RULE1\"; classtype:icmp-event; rev:1; )"` + "\n"
	config += `	rule_groups = [{` + "\n"
	config += `		id = "12da34567890-1234-5678-90ab-cdef12345678"` + "\n"
	config += `		name = "Custom Rule Group 1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_intrusion_rule" "test" {
			id = fmc_intrusion_rule.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcIntrusionRuleConfig() string {
	config := `resource "fmc_intrusion_rule" "test" {` + "\n"
	config += `	rule_data = "alert icmp any any -> any any ( sid:10000301; gid:2000; msg:\"CUSTOM RULE1\"; classtype:icmp-event; rev:1; )"` + "\n"
	config += `	rule_groups = [{` + "\n"
	config += `		id = "12da34567890-1234-5678-90ab-cdef12345678"` + "\n"
	config += `		name = "Custom Rule Group 1"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_intrusion_rule" "test" {
			name = fmc_intrusion_rule.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
