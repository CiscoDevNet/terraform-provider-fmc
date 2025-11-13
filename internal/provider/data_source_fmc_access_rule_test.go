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

func TestAccDataSourceFmcAccessRule(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "action", "ALLOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "name", "rule_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "source_network_literals.0.value", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "destination_network_literals.0.value", "10.2.2.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "vlan_tag_literals.0.start_tag", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "vlan_tag_literals.0.end_tag", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "source_port_literals.0.type", "PortLiteral"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "source_port_literals.0.port", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "source_port_literals.0.protocol", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "destination_port_literals.0.type", "PortLiteral"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "destination_port_literals.0.port", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "destination_port_literals.0.protocol", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "log_connection_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "log_connection_end", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "application_filters.0.types.0.id", "WEBAPP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "application_filters.0.risks.0.id", "VERY_LOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "application_filters.0.business_relevances.0.id", "LOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "application_filters.0.categories.0.id", "118"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_rule.test", "application_filters.0.tags.0.id", "24"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcAccessRulePrerequisitesConfig + testAccDataSourceFmcAccessRuleConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcAccessRulePrerequisitesConfig + testAccNamedDataSourceFmcAccessRuleConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcAccessRulePrerequisitesConfig = `
resource "fmc_access_control_policy" "test" {
  name              = "access_rule"
  default_action    = "BLOCK"
  manage_rules      = false
  manage_categories = false
}

resource "fmc_network" "test" {
  name   = "fmc_access_rule_network"
  prefix = "10.0.0.0/24"
}

resource "fmc_host" "test" {
  name = "fmc_access_rule_host"
  ip   = "10.1.1.1"
}

resource "fmc_port" "test" {
  name     = "fmc_access_rule_port"
  protocol = "UDP"
  port     = "53"
}

resource "fmc_vlan_tag" "test" {
  name      = "fmc_access_rule_vlan_tag"
  start_tag = "10"
  end_tag   = "11"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcAccessRuleConfig() string {
	config := `resource "fmc_access_rule" "test" {` + "\n"
	config += `	access_control_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `	action = "ALLOW"` + "\n"
	config += `	name = "rule_1"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	source_network_literals = [{` + "\n"
	config += `		value = "10.1.1.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	destination_network_literals = [{` + "\n"
	config += `		value = "10.2.2.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	vlan_tag_literals = [{` + "\n"
	config += `		start_tag = "11"` + "\n"
	config += `		end_tag = "22"` + "\n"
	config += `	}]` + "\n"
	config += `	vlan_tag_objects = [{` + "\n"
	config += `		id = fmc_vlan_tag.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	source_network_objects = [{` + "\n"
	config += `		id = fmc_network.test.id` + "\n"
	config += `		type = fmc_network.test.type` + "\n"
	config += `	}]` + "\n"
	config += `	destination_network_objects = [{` + "\n"
	config += `		id = fmc_host.test.id` + "\n"
	config += `		type = fmc_host.test.type` + "\n"
	config += `	}]` + "\n"
	config += `	source_port_literals = [{` + "\n"
	config += `		type = "PortLiteral"` + "\n"
	config += `		port = "80"` + "\n"
	config += `		protocol = "6"` + "\n"
	config += `	}]` + "\n"
	config += `	destination_port_literals = [{` + "\n"
	config += `		type = "PortLiteral"` + "\n"
	config += `		port = "80"` + "\n"
	config += `		protocol = "6"` + "\n"
	config += `	}]` + "\n"
	config += `	log_connection_begin = true` + "\n"
	config += `	log_connection_end = true` + "\n"
	config += `	log_files = false` + "\n"
	config += `	send_events_to_fmc = true` + "\n"
	config += `	description = ""` + "\n"
	config += `	application_filters = [{` + "\n"
	config += `		types = [{` + "\n"
	config += `			id = "WEBAPP"` + "\n"
	config += `		}]` + "\n"
	config += `		risks = [{` + "\n"
	config += `			id = "VERY_LOW"` + "\n"
	config += `		}]` + "\n"
	config += `		business_relevances = [{` + "\n"
	config += `			id = "LOW"` + "\n"
	config += `		}]` + "\n"
	config += `		categories = [{` + "\n"
	config += `			id = "118"` + "\n"
	config += `		}]` + "\n"
	config += `		tags = [{` + "\n"
	config += `			id = "24"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_access_rule" "test" {
			id = fmc_access_rule.test.id
			access_control_policy_id = fmc_access_control_policy.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcAccessRuleConfig() string {
	config := `resource "fmc_access_rule" "test" {` + "\n"
	config += `	access_control_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `	action = "ALLOW"` + "\n"
	config += `	name = "rule_1"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	source_network_literals = [{` + "\n"
	config += `		value = "10.1.1.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	destination_network_literals = [{` + "\n"
	config += `		value = "10.2.2.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	vlan_tag_literals = [{` + "\n"
	config += `		start_tag = "11"` + "\n"
	config += `		end_tag = "22"` + "\n"
	config += `	}]` + "\n"
	config += `	vlan_tag_objects = [{` + "\n"
	config += `		id = fmc_vlan_tag.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	source_network_objects = [{` + "\n"
	config += `		id = fmc_network.test.id` + "\n"
	config += `		type = fmc_network.test.type` + "\n"
	config += `	}]` + "\n"
	config += `	destination_network_objects = [{` + "\n"
	config += `		id = fmc_host.test.id` + "\n"
	config += `		type = fmc_host.test.type` + "\n"
	config += `	}]` + "\n"
	config += `	source_port_literals = [{` + "\n"
	config += `		type = "PortLiteral"` + "\n"
	config += `		port = "80"` + "\n"
	config += `		protocol = "6"` + "\n"
	config += `	}]` + "\n"
	config += `	destination_port_literals = [{` + "\n"
	config += `		type = "PortLiteral"` + "\n"
	config += `		port = "80"` + "\n"
	config += `		protocol = "6"` + "\n"
	config += `	}]` + "\n"
	config += `	log_connection_begin = true` + "\n"
	config += `	log_connection_end = true` + "\n"
	config += `	log_files = false` + "\n"
	config += `	send_events_to_fmc = true` + "\n"
	config += `	description = ""` + "\n"
	config += `	application_filters = [{` + "\n"
	config += `		types = [{` + "\n"
	config += `			id = "WEBAPP"` + "\n"
	config += `		}]` + "\n"
	config += `		risks = [{` + "\n"
	config += `			id = "VERY_LOW"` + "\n"
	config += `		}]` + "\n"
	config += `		business_relevances = [{` + "\n"
	config += `			id = "LOW"` + "\n"
	config += `		}]` + "\n"
	config += `		categories = [{` + "\n"
	config += `			id = "118"` + "\n"
	config += `		}]` + "\n"
	config += `		tags = [{` + "\n"
	config += `			id = "24"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_access_rule" "test" {
			access_control_policy_id = fmc_access_control_policy.test.id
			name = fmc_access_rule.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
