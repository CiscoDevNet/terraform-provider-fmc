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

func TestAccFmcVPNRALoadBalancing(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_ra_load_balancing.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_load_balancing.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_load_balancing.test", "ipv4_group_address", "192.168.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_load_balancing.test", "ipv6_group_address", "2001:db8::1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_load_balancing.test", "interface_id", "76d24097-41c4-4558-a4d0"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_load_balancing.test", "udp_port_number", "9023"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_load_balancing.test", "ipsec", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_load_balancing.test", "redirect_using_fqdn", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_ra_load_balancing.test", "ikev2_redirect_phase", "DURING_SA_AUTHENTICATION"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNRALoadBalancingConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNRALoadBalancingConfig_all(),
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

func testAccFmcVPNRALoadBalancingConfig_minimum() string {
	config := `resource "fmc_vpn_ra_load_balancing" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	ipv4_group_address = "192.168.1.1"` + "\n"
	config += `	interface_id = "76d24097-41c4-4558-a4d0"` + "\n"
	config += `	udp_port_number = 9023` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNRALoadBalancingConfig_all() string {
	config := `resource "fmc_vpn_ra_load_balancing" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	enabled = true` + "\n"
	config += `	ipv4_group_address = "192.168.1.1"` + "\n"
	config += `	ipv6_group_address = "2001:db8::1"` + "\n"
	config += `	interface_id = "76d24097-41c4-4558-a4d0"` + "\n"
	config += `	udp_port_number = 9023` + "\n"
	config += `	ipsec = true` + "\n"
	config += `	ipsec_encryption_key = "my-secret-key"` + "\n"
	config += `	redirect_using_fqdn = false` + "\n"
	config += `	ikev2_redirect_phase = "DURING_SA_AUTHENTICATION"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
