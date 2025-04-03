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

func TestAccFmcIPv4AddressPools(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ipv4_address_pools.test", "items.my_ipv4_address_pools.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ipv4_address_pools.test", "items.my_ipv4_address_pools.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv4_address_pools.test", "items.my_ipv4_address_pools.description", "My IPv4 Address Pool object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv4_address_pools.test", "items.my_ipv4_address_pools.range", "10.0.0.10-10.0.0.20"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv4_address_pools.test", "items.my_ipv4_address_pools.netmask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv4_address_pools.test", "items.my_ipv4_address_pools.overridable", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcIPv4AddressPoolsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcIPv4AddressPoolsConfig_all(),
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcIPv4AddressPoolsConfig_minimum() string {
	config := `resource "fmc_ipv4_address_pools" "test" {` + "\n"
	config += `	items = { "my_ipv4_address_pools" = {` + "\n"
	config += `		range = "10.0.0.10-10.0.0.20"` + "\n"
	config += `		netmask = "255.255.255.0"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcIPv4AddressPoolsConfig_all() string {
	config := `resource "fmc_ipv4_address_pools" "test" {` + "\n"
	config += `	items = { "my_ipv4_address_pools" = {` + "\n"
	config += `		description = "My IPv4 Address Pool object"` + "\n"
	config += `		range = "10.0.0.10-10.0.0.20"` + "\n"
	config += `		netmask = "255.255.255.0"` + "\n"
	config += `		overridable = true` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
