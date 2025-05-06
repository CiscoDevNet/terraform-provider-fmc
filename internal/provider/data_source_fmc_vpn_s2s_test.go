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

func TestAccDataSourceFmcVPNS2S(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s.test", "name", "my_ftd_s2s_vpn"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_s2s.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s.test", "route_based", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s.test", "topology_type", "POINT_TO_POINT"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s.test", "ikev1_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s.test", "ikev2_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_s2s.test", "ipsec_policy_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_s2s.test", "ike_policy_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_s2s.test", "advanced_settings_policy_id"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVPNS2SConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcVPNS2SConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVPNS2SConfig() string {
	config := `resource "fmc_vpn_s2s" "test" {` + "\n"
	config += `	name = "my_ftd_s2s_vpn"` + "\n"
	config += `	route_based = true` + "\n"
	config += `	topology_type = "POINT_TO_POINT"` + "\n"
	config += `	ikev1_enable = false` + "\n"
	config += `	ikev2_enable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vpn_s2s" "test" {
			id = fmc_vpn_s2s.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcVPNS2SConfig() string {
	config := `resource "fmc_vpn_s2s" "test" {` + "\n"
	config += `	name = "my_ftd_s2s_vpn"` + "\n"
	config += `	route_based = true` + "\n"
	config += `	topology_type = "POINT_TO_POINT"` + "\n"
	config += `	ikev1_enable = false` + "\n"
	config += `	ikev2_enable = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vpn_s2s" "test" {
			name = fmc_vpn_s2s.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
