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

func TestAccDataSourceFmcRealmLocal(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_local.test", "name", "my_ldap_realm"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_realm_local.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_realm_local.test", "version"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_realm_local.test", "description", "My realm"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcRealmLocalConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcRealmLocalConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcRealmLocalConfig() string {
	config := `resource "fmc_realm_local" "test" {` + "\n"
	config += `	name = "my_ldap_realm"` + "\n"
	config += `	description = "My realm"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_realm_local" "test" {
			id = fmc_realm_local.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcRealmLocalConfig() string {
	config := `resource "fmc_realm_local" "test" {` + "\n"
	config += `	name = "my_ldap_realm"` + "\n"
	config += `	description = "My realm"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_realm_local" "test" {
			name = fmc_realm_local.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
