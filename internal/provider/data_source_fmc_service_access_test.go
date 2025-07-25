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

func TestAccDataSourceFmcServiceAccess(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_service_access.test", "name", "my_service_access"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_service_access.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_service_access.test", "default_action", "DENY"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_service_access.test", "rules.0.action", "ALLOW"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcServiceAccessPrerequisitesConfig + testAccDataSourceFmcServiceAccessConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcServiceAccessPrerequisitesConfig + testAccNamedDataSourceFmcServiceAccessConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcServiceAccessPrerequisitesConfig = `
data "fmc_countries" "test" {
  items = {
    "Poland"  = {}
  }
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcServiceAccessConfig() string {
	config := `resource "fmc_service_access" "test" {` + "\n"
	config += `	name = "my_service_access"` + "\n"
	config += `	default_action = "DENY"` + "\n"
	config += `	rules = [{` + "\n"
	config += `		action = "ALLOW"` + "\n"
	config += `		geolocation_sources = [{` + "\n"
	config += `			id = data.fmc_countries.test.items["Poland"].id` + "\n"
	config += `			type = data.fmc_countries.test.items["Poland"].type` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_service_access" "test" {
			id = fmc_service_access.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcServiceAccessConfig() string {
	config := `resource "fmc_service_access" "test" {` + "\n"
	config += `	name = "my_service_access"` + "\n"
	config += `	default_action = "DENY"` + "\n"
	config += `	rules = [{` + "\n"
	config += `		action = "ALLOW"` + "\n"
	config += `		geolocation_sources = [{` + "\n"
	config += `			id = data.fmc_countries.test.items["Poland"].id` + "\n"
	config += `			type = data.fmc_countries.test.items["Poland"].type` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_service_access" "test" {
			name = fmc_service_access.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
