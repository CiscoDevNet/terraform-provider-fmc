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

func TestAccFmcPrefilterPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "name", "POLICY1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "description", "My prefilter policy"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_prefilter_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "default_action", "BLOCK_TUNNELS"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_prefilter_policy.test", "default_action_id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "default_action_log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "default_action_log_end", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "default_action_send_events_to_fmc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.name", "rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.rule_type", "PREFILTER"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.action", "FASTPATH"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.source_network_literals.0.value", "10.1.1.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.destination_network_literals.0.value", "10.2.2.0/24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.vlan_tag_literals.0.start_tag", "11"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.vlan_tag_literals.0.end_tag", "22"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.source_port_literals.0.protocol", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.source_port_literals.0.port", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.destination_port_literals.0.protocol", "6"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.destination_port_literals.0.port", "80"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.log_begin", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.log_end", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_prefilter_policy.test", "rules.0.send_events_to_fmc", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcPrefilterPolicyPrerequisitesConfig + testAccFmcPrefilterPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcPrefilterPolicyPrerequisitesConfig + testAccFmcPrefilterPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_prefilter_policy.test",
		ImportState:  true,
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

const testAccFmcPrefilterPolicyPrerequisitesConfig = `
resource "fmc_network" "test" {
  name   = "NET_fmc_access_control_policy"
  prefix = "10.0.0.0/24"
}

resource "fmc_host" "test" {
  name = "HOST_fmc_access_control_policy"
  ip   = "10.1.1.1"
}

resource "fmc_port" "test" {
  name     = "test_fmc_access_control_policy"
  protocol = "UDP"
  port     = "53"
}

resource "fmc_vlan_tag" "test" {
  name      = "VLAN_TAG_fmc_access_control_policy"
  start_tag = "10"
  end_tag   = "11" 
}

resource "fmc_security_zone" "test" {
  name           = "security_zone_1"
  interface_mode = "ROUTED"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcPrefilterPolicyConfig_minimum() string {
	config := `resource "fmc_prefilter_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `	default_action = "BLOCK_TUNNELS"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcPrefilterPolicyConfig_all() string {
	config := `resource "fmc_prefilter_policy" "test" {` + "\n"
	config += `	name = "POLICY1"` + "\n"
	config += `	description = "My prefilter policy"` + "\n"
	config += `	default_action = "BLOCK_TUNNELS"` + "\n"
	config += `	default_action_log_begin = true` + "\n"
	config += `	default_action_log_end = false` + "\n"
	config += `	default_action_send_events_to_fmc = true` + "\n"
	config += `	rules = [{` + "\n"
	config += `		name = "rule1"` + "\n"
	config += `		rule_type = "PREFILTER"` + "\n"
	config += `		enabled = true` + "\n"
	config += `		action = "FASTPATH"` + "\n"
	config += `		bidirectional = false` + "\n"
	config += `		source_interfaces = [{` + "\n"
	config += `			id = fmc_security_zone.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_interfaces = [{` + "\n"
	config += `			id = fmc_security_zone.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		source_network_literals = [{` + "\n"
	config += `			value = "10.1.1.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `		source_network_objects = [{` + "\n"
	config += `			id = fmc_network.test.id` + "\n"
	config += `			type = fmc_network.test.type` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_literals = [{` + "\n"
	config += `			value = "10.2.2.0/24"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_network_objects = [{` + "\n"
	config += `			id = fmc_host.test.id` + "\n"
	config += `			type = fmc_host.test.type` + "\n"
	config += `		}]` + "\n"
	config += `		vlan_tag_literals = [{` + "\n"
	config += `			start_tag = "11"` + "\n"
	config += `			end_tag = "22"` + "\n"
	config += `		}]` + "\n"
	config += `		vlan_tag_objects = [{` + "\n"
	config += `			id = fmc_vlan_tag.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		source_port_literals = [{` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `			port = "80"` + "\n"
	config += `		}]` + "\n"
	config += `		source_port_objects = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		destination_port_literals = [{` + "\n"
	config += `			protocol = "6"` + "\n"
	config += `			port = "80"` + "\n"
	config += `		}]` + "\n"
	config += `		destination_port_objects = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		log_begin = true` + "\n"
	config += `		log_end = true` + "\n"
	config += `		send_events_to_fmc = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
