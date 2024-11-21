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

func TestAccDataSourceFmcPortGroups(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_port_groups.test", "items.port_group_1.id"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_port_groups.test", "items.port_group_1.description", "My port group description"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_port_groups.test", "items.port_group_1.ports.0.id"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_port_groups.test", "items.port_group_1.ports.0.type", "ProtocolPortObject"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcPortGroupsPrerequisitesConfig + testAccDataSourceFmcPortGroupsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcPortGroupsPrerequisitesConfig = `
resource "fmc_port" "test" {
  name        = "port_1"
  description = "My PORT id"
  protocol    = "TCP"
  port        = "443"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcPortGroupsConfig() string {
	config := `resource "fmc_port_groups" "test" {` + "\n"
	config += `	items = { "port_group_1" = {` + "\n"
	config += `		description = "My port group description"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		ports = [{` + "\n"
	config += `			id = fmc_port.test.id` + "\n"
	config += `			type = "ProtocolPortObject"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_port_groups" "test" {
			depends_on = [fmc_port_groups.test]
			items = {
				"port_group_1" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
