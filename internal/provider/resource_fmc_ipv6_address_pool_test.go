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

func TestAccFmcIPv6AddressPool(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv6_address_pool.test", "name", "my_ipv6_address_pool"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ipv6_address_pool.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv6_address_pool.test", "description", "My IPv4 Address Pool object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv6_address_pool.test", "start_address", "2001:db8::1/64"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv6_address_pool.test", "number_of_addresses", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv6_address_pool.test", "overridable", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcIPv6AddressPoolConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcIPv6AddressPoolConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_ipv6_address_pool.test",
		ImportState:  true,
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcIPv6AddressPoolConfig_minimum() string {
	config := `resource "fmc_ipv6_address_pool" "test" {` + "\n"
	config += `	name = "my_ipv6_address_pool"` + "\n"
	config += `	start_address = "2001:db8::1/64"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcIPv6AddressPoolConfig_all() string {
	config := `resource "fmc_ipv6_address_pool" "test" {` + "\n"
	config += `	name = "my_ipv6_address_pool"` + "\n"
	config += `	description = "My IPv4 Address Pool object"` + "\n"
	config += `	start_address = "2001:db8::1/64"` + "\n"
	config += `	number_of_addresses = 10` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
