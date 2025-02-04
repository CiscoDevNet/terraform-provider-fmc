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

func TestAccFmcVLANTagGroups(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vlan_tag_groups.test", "items.vlan_tag_group_1.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vlan_tag_groups.test", "items.vlan_tag_group_1.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vlan_tag_groups.test", "items.vlan_tag_group_1.description", "My vlan tag group name"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vlan_tag_groups.test", "items.vlan_tag_group_1.overridable", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVLANTagGroupsPrerequisitesConfig + testAccFmcVLANTagGroupsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVLANTagGroupsPrerequisitesConfig + testAccFmcVLANTagGroupsConfig_all(),
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

const testAccFmcVLANTagGroupsPrerequisitesConfig = `
resource "fmc_vlan_tag" "test" {
  name        = "vlan_tag_1111"
  description = "My TAG id"
  overridable = false
  start_tag   = 11
  end_tag     = 12
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcVLANTagGroupsConfig_minimum() string {
	config := `resource "fmc_vlan_tag_groups" "test" {` + "\n"
	config += `	items = { "vlan_tag_group_1" = {` + "\n"
	config += `		vlan_tags = [{` + "\n"
	config += `			id = fmc_vlan_tag.test.id` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVLANTagGroupsConfig_all() string {
	config := `resource "fmc_vlan_tag_groups" "test" {` + "\n"
	config += `	items = { "vlan_tag_group_1" = {` + "\n"
	config += `		description = "My vlan tag group name"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		vlan_tags = [{` + "\n"
	config += `			id = fmc_vlan_tag.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		literals = [{` + "\n"
	config += `			start_tag = 11` + "\n"
	config += `			end_tag = 22` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
