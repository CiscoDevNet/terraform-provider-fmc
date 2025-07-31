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

func TestAccFmcVPNRAAddressAssignmentPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra_address_assignment_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_address_assignment_policy.test", "ipv4_use_authorization_server", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_address_assignment_policy.test", "ipv4_use_dhcp", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_address_assignment_policy.test", "ipv4_internal_address_pool", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_address_assignment_policy.test", "ipv4_internal_address_pool_reuse_interval", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_address_assignment_policy.test", "ipv6_use_authorization_server", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_address_assignment_policy.test", "ipv6_internal_address_pool", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNRAAddressAssignmentPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNRAAddressAssignmentPolicyConfig_all(),
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

func testAccFmcVPNRAAddressAssignmentPolicyConfig_minimum() string {
	config := `resource "fmc_vpn_ra_address_assignment_policy" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNRAAddressAssignmentPolicyConfig_all() string {
	config := `resource "fmc_vpn_ra_address_assignment_policy" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	ipv4_use_authorization_server = true` + "\n"
	config += `	ipv4_use_dhcp = true` + "\n"
	config += `	ipv4_internal_address_pool = true` + "\n"
	config += `	ipv4_internal_address_pool_reuse_interval = 0` + "\n"
	config += `	ipv6_use_authorization_server = true` + "\n"
	config += `	ipv6_internal_address_pool = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
