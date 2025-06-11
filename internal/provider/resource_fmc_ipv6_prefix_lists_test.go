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

func TestAccFmcIPv6PrefixLists(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ipv6_prefix_lists.test", "items.my_ipv6_prefix_lists.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_ipv6_prefix_lists.test", "items.my_ipv6_prefix_lists.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv6_prefix_lists.test", "items.my_ipv6_prefix_lists.entries.0.action", "PERMIT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv6_prefix_lists.test", "items.my_ipv6_prefix_lists.entries.0.ip_address", "2001:1::/64"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv6_prefix_lists.test", "items.my_ipv6_prefix_lists.entries.0.min_prefix_length", "65"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_ipv6_prefix_lists.test", "items.my_ipv6_prefix_lists.entries.0.max_prefix_length", "120"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcIPv6PrefixListsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcIPv6PrefixListsConfig_all(),
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

func testAccFmcIPv6PrefixListsConfig_minimum() string {
	config := `resource "fmc_ipv6_prefix_lists" "test" {` + "\n"
	config += `	items = { "my_ipv6_prefix_lists" = {` + "\n"
	config += `		entries = [{` + "\n"
	config += `			action = "PERMIT"` + "\n"
	config += `			ip_address = "2001:1::/64"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcIPv6PrefixListsConfig_all() string {
	config := `resource "fmc_ipv6_prefix_lists" "test" {` + "\n"
	config += `	items = { "my_ipv6_prefix_lists" = {` + "\n"
	config += `		entries = [{` + "\n"
	config += `			action = "PERMIT"` + "\n"
	config += `			ip_address = "2001:1::/64"` + "\n"
	config += `			min_prefix_length = 65` + "\n"
	config += `			max_prefix_length = 120` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
