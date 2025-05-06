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

func TestAccFmcVPNS2S(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s.test", "name", "my_ftd_s2s_vpn"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_s2s.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s.test", "route_based", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s.test", "topology_type", "POINT_TO_POINT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s.test", "ikev1_enable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s.test", "ikev2_enable", "true"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_s2s.test", "ipsec_policy_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_s2s.test", "ike_policy_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_s2s.test", "advanced_settings_policy_id"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNS2SConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNS2SConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_vpn_s2s.test",
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

func testAccFmcVPNS2SConfig_minimum() string {
	config := `resource "fmc_vpn_s2s" "test" {` + "\n"
	config += `	name = "my_ftd_s2s_vpn"` + "\n"
	config += `	route_based = true` + "\n"
	config += `	topology_type = "POINT_TO_POINT"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNS2SConfig_all() string {
	config := `resource "fmc_vpn_s2s" "test" {` + "\n"
	config += `	name = "my_ftd_s2s_vpn"` + "\n"
	config += `	route_based = true` + "\n"
	config += `	topology_type = "POINT_TO_POINT"` + "\n"
	config += `	ikev1_enable = false` + "\n"
	config += `	ikev2_enable = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
