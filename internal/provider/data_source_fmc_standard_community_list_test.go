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

func TestAccDataSourceFmcStandardCommunityList(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_community_list.test", "name", "my_standard_community_list"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_standard_community_list.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_community_list.test", "entries.0.action", "PERMIT"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_community_list.test", "entries.0.communities", "123 456 789"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_community_list.test", "entries.0.internet", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_community_list.test", "entries.0.no_advertise", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_standard_community_list.test", "entries.0.no_export", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcStandardCommunityListConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcStandardCommunityListConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcStandardCommunityListConfig() string {
	config := `resource "fmc_standard_community_list" "test" {` + "\n"
	config += `	name = "my_standard_community_list"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "PERMIT"` + "\n"
	config += `		communities = "123 456 789"` + "\n"
	config += `		internet = true` + "\n"
	config += `		no_advertise = true` + "\n"
	config += `		no_export = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_standard_community_list" "test" {
			id = fmc_standard_community_list.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcStandardCommunityListConfig() string {
	config := `resource "fmc_standard_community_list" "test" {` + "\n"
	config += `	name = "my_standard_community_list"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "PERMIT"` + "\n"
	config += `		communities = "123 456 789"` + "\n"
	config += `		internet = true` + "\n"
	config += `		no_advertise = true` + "\n"
	config += `		no_export = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_standard_community_list" "test" {
			name = fmc_standard_community_list.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
