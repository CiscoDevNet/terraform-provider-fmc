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

func TestAccDataSourceFmcAccessCategory(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_access_category.test", "name", "my_category"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_access_category.test", "type"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcAccessCategoryPrerequisitesConfig + testAccDataSourceFmcAccessCategoryConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcAccessCategoryPrerequisitesConfig + testAccNamedDataSourceFmcAccessCategoryConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcAccessCategoryPrerequisitesConfig = `
resource "fmc_access_control_policy" "test" {
  name              = "access_category"
  default_action    = "BLOCK"
  manage_rules      = false
  manage_categories = false
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcAccessCategoryConfig() string {
	config := `resource "fmc_access_category" "test" {` + "\n"
	config += `	access_control_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `	name = "my_category"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_access_category" "test" {
			id = fmc_access_category.test.id
			access_control_policy_id = fmc_access_control_policy.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcAccessCategoryConfig() string {
	config := `resource "fmc_access_category" "test" {` + "\n"
	config += `	access_control_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `	name = "my_category"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_access_category" "test" {
			access_control_policy_id = fmc_access_control_policy.test.id
			name = fmc_access_category.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
