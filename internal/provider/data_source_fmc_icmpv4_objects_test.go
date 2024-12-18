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

func TestAccDataSourceFmcICMPv4Objects(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_icmpv4_objects.test", "items.icmpv4_1.id"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_icmpv4_objects.test", "items.icmpv4_1.description", "ICMPv4 network unreachable response, type 3, code 0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_icmpv4_objects.test", "items.icmpv4_1.overridable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_icmpv4_objects.test", "items.icmpv4_1.icmp_type", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_icmpv4_objects.test", "items.icmpv4_1.code", "0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcICMPv4ObjectsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcICMPv4ObjectsConfig() string {
	config := `resource "fmc_icmpv4_objects" "test" {` + "\n"
	config += `	items = { "icmpv4_1" = {` + "\n"
	config += `		description = "ICMPv4 network unreachable response, type 3, code 0"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		icmp_type = 3` + "\n"
	config += `		code = 0` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_icmpv4_objects" "test" {
			depends_on = [fmc_icmpv4_objects.test]
			items = {
				"icmpv4_1" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
