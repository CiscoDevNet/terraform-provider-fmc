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
func TestAccDataSourceFmcAccessControlPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "name", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "description", "My access control policy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action", "BLOCK"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_log_end", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_send_syslog", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "default_action_syslog_severity", "DEBUG"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "categories.0.name", "cat1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.action", "ALLOW"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.name", "rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.source_network_literals.0.value", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.destination_network_literals.0.value", "10.2.2.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.log_end", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_control_policy.test", "rules.0.send_events_to_fmc", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcAccessControlPolicyPrerequisitesConfig + testAccDataSourceFmcAccessControlPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceFmcAccessControlPolicyPrerequisitesConfig = `
resource "fmc_network" "this" {
  name   = "mynetwork1"
  prefix = "10.0.0.0/24"
}

resource "fmc_host" "this" {
  name = "myhost1"
  ip   = "10.1.1.1"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceFmcAccessControlPolicyConfig() string {
	config := `resource "fmc_access_control_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `	description = "My access control policy"` + "\n"
	config += `	default_action = "BLOCK"` + "\n"
	config += `	default_action_log_begin = true` + "\n"
	config += `	default_action_log_end = false` + "\n"
	config += `	default_action_send_events_to_fmc = true` + "\n"
	config += `	default_action_send_syslog = true` + "\n"
	config += `	default_action_syslog_severity = "DEBUG"` + "\n"
	config += `	categories = [{` + "\n"
	config += `	  name = "cat1"` + "\n"
	config += `	}]` + "\n"
	config += `	rules = [{` + "\n"
	config += `	  action = "ALLOW"` + "\n"
	config += `	  name = "rule1"` + "\n"
	config += `	  category_name = "cat1"` + "\n"
	config += `	  enabled = true` + "\n"
	config += `	  source_network_literals = [{` + "\n"
	config += `		value = "10.1.1.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	  destination_network_literals = [{` + "\n"
	config += `		value = "10.2.2.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `	  source_network_objects = [{` + "\n"
	config += `		id = fmc_network.this.id` + "\n"
	config += `		type = fmc_network.this.type` + "\n"
	config += `	}]` + "\n"
	config += `	  destination_network_objects = [{` + "\n"
	config += `		id = fmc_host.this.id` + "\n"
	config += `		type = fmc_host.this.type` + "\n"
	config += `	}]` + "\n"
	config += `	  log_begin = true` + "\n"
	config += `	  log_end = true` + "\n"
	config += `	  log_files = false` + "\n"
	config += `	  send_events_to_fmc = true` + "\n"
	config += `	  description = ""` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_access_control_policy" "test" {
			id = fmc_access_control_policy.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
