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

func TestAccDataSourceFmcVLANTag(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vlan_tag.test", "name", "vlan_tag_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vlan_tag.test", "description", "My TAG id"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vlan_tag.test", "overridable", "true"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vlan_tag.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVLANTagConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcVLANTagConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVLANTagConfig() string {
	config := `resource "fmc_vlan_tag" "test" {` + "\n"
	config += `	name = "vlan_tag_1"` + "\n"
	config += `	description = "My TAG id"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	start_tag = 12` + "\n"
	config += `	end_tag = 15` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vlan_tag" "test" {
			id = fmc_vlan_tag.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcVLANTagConfig() string {
	config := `resource "fmc_vlan_tag" "test" {` + "\n"
	config += `	name = "vlan_tag_1"` + "\n"
	config += `	description = "My TAG id"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	start_tag = 12` + "\n"
	config += `	end_tag = 15` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vlan_tag" "test" {
			name = fmc_vlan_tag.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
