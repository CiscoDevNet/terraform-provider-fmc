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

func TestAccDataSourceFmcSecureClientCustomAttribute(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_secure_client_custom_attribute.test", "name", "my_secure_client_custom_attribute"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_secure_client_custom_attribute.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_secure_client_custom_attribute.test", "description", "My Secure Client Custom Attribute"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_secure_client_custom_attribute.test", "attribute_type", "USER_DEFINED_CUSTOM_ATTR"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_secure_client_custom_attribute.test", "user_defined_attribute_name", "my_user_defined_attribute"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_secure_client_custom_attribute.test", "user_defined_attribute_value", "my_value"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcSecureClientCustomAttributeConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcSecureClientCustomAttributeConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcSecureClientCustomAttributeConfig() string {
	config := `resource "fmc_secure_client_custom_attribute" "test" {` + "\n"
	config += `	name = "my_secure_client_custom_attribute"` + "\n"
	config += `	description = "My Secure Client Custom Attribute"` + "\n"
	config += `	attribute_type = "USER_DEFINED_CUSTOM_ATTR"` + "\n"
	config += `	user_defined_attribute_name = "my_user_defined_attribute"` + "\n"
	config += `	user_defined_attribute_value = "my_value"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_secure_client_custom_attribute" "test" {
			id = fmc_secure_client_custom_attribute.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcSecureClientCustomAttributeConfig() string {
	config := `resource "fmc_secure_client_custom_attribute" "test" {` + "\n"
	config += `	name = "my_secure_client_custom_attribute"` + "\n"
	config += `	description = "My Secure Client Custom Attribute"` + "\n"
	config += `	attribute_type = "USER_DEFINED_CUSTOM_ATTR"` + "\n"
	config += `	user_defined_attribute_name = "my_user_defined_attribute"` + "\n"
	config += `	user_defined_attribute_value = "my_value"` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_secure_client_custom_attribute" "test" {
			name = fmc_secure_client_custom_attribute.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
