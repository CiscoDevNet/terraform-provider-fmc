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

func TestAccDataSourceFmcVLANTagGroup(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vlan_tag_group.test", "name", "fmc_vlan_tag_group"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vlan_tag_group.test", "description", "My VLAN Tag Group"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vlan_tag_group.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vlan_tag_group.test", "overridable", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVLANTagGroupPrerequisitesConfig + testAccDataSourceFmcVLANTagGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcVLANTagGroupPrerequisitesConfig + testAccNamedDataSourceFmcVLANTagGroupConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcVLANTagGroupPrerequisitesConfig = `
resource "fmc_vlan_tag" "test" {
  name        = "fmc_vlan_tag_group_vlan_tag"
  description = "My TAG id"
  overridable = false
  start_tag   = 11
  end_tag     = 12
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVLANTagGroupConfig() string {
	config := `resource "fmc_vlan_tag_group" "test" {` + "\n"
	config += `	name = "fmc_vlan_tag_group"` + "\n"
	config += `	description = "My VLAN Tag Group"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	vlan_tags = [{` + "\n"
	config += `		id = fmc_vlan_tag.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	literals = [{` + "\n"
	config += `		start_tag = 11` + "\n"
	config += `		end_tag = 22` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vlan_tag_group" "test" {
			id = fmc_vlan_tag_group.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcVLANTagGroupConfig() string {
	config := `resource "fmc_vlan_tag_group" "test" {` + "\n"
	config += `	name = "fmc_vlan_tag_group"` + "\n"
	config += `	description = "My VLAN Tag Group"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	vlan_tags = [{` + "\n"
	config += `		id = fmc_vlan_tag.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	literals = [{` + "\n"
	config += `		start_tag = 11` + "\n"
	config += `		end_tag = 22` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vlan_tag_group" "test" {
			name = fmc_vlan_tag_group.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
