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

func TestAccDataSourceFmcKeyChains(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_key_chains.test", "items.my_key_chains.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_key_chains.test", "items.my_key_chains.type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_key_chains.test", "items.my_key_chains.description", "My Host object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_key_chains.test", "items.my_key_chains.keys.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_key_chains.test", "items.my_key_chains.keys.0.key", "my_secret_key"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_key_chains.test", "items.my_key_chains.keys.0.accept_lifetime_start", "2025-08-25T12:14:23"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_key_chains.test", "items.my_key_chains.keys.0.accept_lifetime_end_type", "DATETIME"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_key_chains.test", "items.my_key_chains.keys.0.accept_lifetime_end", "2026-08-25T12:14:23"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_key_chains.test", "items.my_key_chains.keys.0.send_lifetime_start", "2025-08-25T12:14:23"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_key_chains.test", "items.my_key_chains.keys.0.send_lifetime_end_type", "DATETIME"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_key_chains.test", "items.my_key_chains.keys.0.send_lifetime_end", "2026-08-25T12:14:23"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcKeyChainsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcKeyChainsConfig() string {
	config := `resource "fmc_key_chains" "test" {` + "\n"
	config += `	items = { "my_key_chains" = {` + "\n"
	config += `		description = "My Host object"` + "\n"
	config += `		keys = [{` + "\n"
	config += `			id = 1` + "\n"
	config += `			key = "my_secret_key"` + "\n"
	config += `			accept_lifetime_start = "2025-08-25T12:14:23"` + "\n"
	config += `			accept_lifetime_end_type = "DATETIME"` + "\n"
	config += `			accept_lifetime_end = "2026-08-25T12:14:23"` + "\n"
	config += `			send_lifetime_start = "2025-08-25T12:14:23"` + "\n"
	config += `			send_lifetime_end_type = "DATETIME"` + "\n"
	config += `			send_lifetime_end = "2026-08-25T12:14:23"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_key_chains" "test" {
			depends_on = [fmc_key_chains.test]
			items = {
				"my_key_chains" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
