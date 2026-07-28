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

func TestAccDataSourceFmcDistinguishedName(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_distinguished_name.test", "name", "my_distinguished_name"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_distinguished_name.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_distinguished_name.test", "distinguished_name", "C=US,CN=example.com,O=Example,OU=IT"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDistinguishedNameConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcDistinguishedNameConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDistinguishedNameConfig() string {
	config := `resource "fmc_distinguished_name" "test" {` + "\n"
	config += `	name = "my_distinguished_name"` + "\n"
	config += `	distinguished_name = "C=US,CN=example.com,O=Example,OU=IT"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_distinguished_name" "test" {
			id = fmc_distinguished_name.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDistinguishedNameConfig() string {
	config := `resource "fmc_distinguished_name" "test" {` + "\n"
	config += `	name = "my_distinguished_name"` + "\n"
	config += `	distinguished_name = "C=US,CN=example.com,O=Example,OU=IT"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_distinguished_name" "test" {
			name = fmc_distinguished_name.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
