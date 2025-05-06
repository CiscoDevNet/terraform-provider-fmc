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

func TestAccFmcVPNS2SIPSECSettings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_s2s_ipsec_settings.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "crypto_map_type", "STATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "ikev2_mode", "TUNNEL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "ikev1_ipsec_proposals.0.name", "my_ikev1_ipsec_proposal"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "ikev2_ipsec_proposals.0.name", "my_ikev2_ipsec_proposal"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "enable_sa_strength_enforcement", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "enable_reverse_route_injection", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "enable_perfect_forward_secrecy", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "perfect_forward_secrecy_modulus_group", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "lifetime_duration", "28800"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "lifetime_size", "4608000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "validate_incoming_icmp_error_messages", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "do_not_fragment_policy", "NONE"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "enable_tfc_packets", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "tfc_burst_bytes", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "tfc_payload_bytes", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_ipsec_settings.test", "tfc_timeout", ""))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNS2SIPSECSettingsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNS2SIPSECSettingsConfig_all(),
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

func testAccFmcVPNS2SIPSECSettingsConfig_minimum() string {
	config := `resource "fmc_vpn_s2s_ipsec_settings" "test" {` + "\n"
	config += `	vpn_s2s_id = TBD` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNS2SIPSECSettingsConfig_all() string {
	config := `resource "fmc_vpn_s2s_ipsec_settings" "test" {` + "\n"
	config += `	vpn_s2s_id = TBD` + "\n"
	config += `	crypto_map_type = "STATIC"` + "\n"
	config += `	ikev2_mode = "TUNNEL"` + "\n"
	config += `	ikev1_ipsec_proposals = [{` + "\n"
	config += `		id = TBD` + "\n"
	config += `		name = "my_ikev1_ipsec_proposal"` + "\n"
	config += `	}]` + "\n"
	config += `	ikev2_ipsec_proposals = [{` + "\n"
	config += `		id = TBD` + "\n"
	config += `		name = "my_ikev2_ipsec_proposal"` + "\n"
	config += `	}]` + "\n"
	config += `	enable_sa_strength_enforcement = false` + "\n"
	config += `	enable_reverse_route_injection = false` + "\n"
	config += `	enable_perfect_forward_secrecy = true` + "\n"
	config += `	perfect_forward_secrecy_modulus_group = "2"` + "\n"
	config += `	lifetime_duration = 28800` + "\n"
	config += `	lifetime_size = 4608000` + "\n"
	config += `	validate_incoming_icmp_error_messages = false` + "\n"
	config += `	do_not_fragment_policy = "NONE"` + "\n"
	config += `	enable_tfc_packets = false` + "\n"
	config += `	tfc_burst_bytes = ` + "\n"
	config += `	tfc_payload_bytes = ` + "\n"
	config += `	tfc_timeout = ` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
