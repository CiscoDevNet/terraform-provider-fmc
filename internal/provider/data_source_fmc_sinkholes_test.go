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

func TestAccDataSourceFmcSinkholes(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_sinkholes.test", "items.my_sinkholes.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_sinkholes.test", "items.my_sinkholes.type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sinkholes.test", "items.my_sinkholes.ipv4_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sinkholes.test", "items.my_sinkholes.ipv6_address", "2001:db8::1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sinkholes.test", "items.my_sinkholes.action", "MONITOR"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sinkholes.test", "items.my_sinkholes.log_connection_type", "SINKHOLE_CNC"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcSinkholesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcSinkholesConfig() string {
	config := `resource "fmc_sinkholes" "test" {` + "\n"
	config += `	items = { "my_sinkholes" = {` + "\n"
	config += `		ipv4_address = "10.1.1.1"` + "\n"
	config += `		ipv6_address = "2001:db8::1"` + "\n"
	config += `		action = "MONITOR"` + "\n"
	config += `		log_connection_type = "SINKHOLE_CNC"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_sinkholes" "test" {
			depends_on = [fmc_sinkholes.test]
			items = {
				"my_sinkholes" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
