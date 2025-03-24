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

func TestAccDataSourceFmcIPv6AddressPool(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ipv6_address_pool.test", "name", "my_ipv6_address_pool"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_ipv6_address_pool.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ipv6_address_pool.test", "description", "My IPv6 Address Pool object"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ipv6_address_pool.test", "start_address", "2001:db8::1/64"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ipv6_address_pool.test", "number_of_addresses", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_ipv6_address_pool.test", "overridable", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcIPv6AddressPoolConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcIPv6AddressPoolConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcIPv6AddressPoolConfig() string {
	config := `resource "fmc_ipv6_address_pool" "test" {` + "\n"
	config += `	name = "my_ipv6_address_pool"` + "\n"
	config += `	description = "My IPv6 Address Pool object"` + "\n"
	config += `	start_address = "2001:db8::1/64"` + "\n"
	config += `	number_of_addresses = 10` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_ipv6_address_pool" "test" {
			id = fmc_ipv6_address_pool.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcIPv6AddressPoolConfig() string {
	config := `resource "fmc_ipv6_address_pool" "test" {` + "\n"
	config += `	name = "my_ipv6_address_pool"` + "\n"
	config += `	description = "My IPv6 Address Pool object"` + "\n"
	config += `	start_address = "2001:db8::1/64"` + "\n"
	config += `	number_of_addresses = 10` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_ipv6_address_pool" "test" {
			name = fmc_ipv6_address_pool.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
