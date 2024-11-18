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

func TestAccFmcPortGroups(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_port_groups.test", "items.port_group_1.id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_port_groups.test", "items.port_group_1.type", "PortObjectGroup"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_port_groups.test", "items.port_group_1.description", "My port group description"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_port_groups.test", "items.port_group_1.ports.0.type", "ProtocolPortObject"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcPortGroupsPrerequisitesConfig + testAccFmcPortGroupsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcPortGroupsPrerequisitesConfig + testAccFmcPortGroupsConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcPortGroupsPrerequisitesConfig = `
resource "fmc_port" "test" {
  name        = "port_1"
  description = "My PORT id"
  protocol    = "TCP"
  port        = "443"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcPortGroupsConfig_minimum() string {
	config := `resource "fmc_port_groups" "test" {` + "\n"
	config += `	items = { "port_group_1" = {` + "\n"
	config += `		type = "PortObjectGroup"` + "\n"
	config += `		ports = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `			type = "ProtocolPortObject"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcPortGroupsConfig_all() string {
	config := `resource "fmc_port_groups" "test" {` + "\n"
	config += `	items = { "port_group_1" = {` + "\n"
	config += `		type = "PortObjectGroup"` + "\n"
	config += `		description = "My port group description"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		ports = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `			type = "ProtocolPortObject"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
