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

func TestAccDataSourceFmcVPNS2SIPSECSettings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_s2s_ipsec_settings.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "crypto_map_type", "STATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "ikev2_mode", "TUNNEL"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "security_association_strength_enforcement", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "reverse_route_injection", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "perfect_forward_secrecy", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "perfect_forward_secrecy_modulus_group", "14"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "lifetime_duration", "28800"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "lifetime_size", "4608000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "validate_incoming_icmp_error_messages", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "do_not_fragment_policy", "NONE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "tfc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "tfc_burst_bytes", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "tfc_payload_bytes", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_ipsec_settings.test", "tfc_timeout", "0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVPNS2SIPSECSettingsPrerequisitesConfig + testAccDataSourceFmcVPNS2SIPSECSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcVPNS2SIPSECSettingsPrerequisitesConfig = `
resource "fmc_vpn_s2s" "test" {
  name             = "my_s2s_vpn_ipsec_settings"
  route_based      = false
  network_topology = "POINT_TO_POINT"
  ikev1            = true
  ikev2            = true
}

resource "fmc_ikev2_ipsec_proposals" "test" {
  items = {
    my_s2s_vpn_ipsec_settings = {
      description     = "IKEv2 IPsec Proposal 1"
      esp_encryptions = ["AES-256"]
      esp_hashes      = ["SHA-256"]
    }
  }
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVPNS2SIPSECSettingsConfig() string {
	config := `resource "fmc_vpn_s2s_ipsec_settings" "test" {` + "\n"
	config += `	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n"
	config += `	crypto_map_type = "STATIC"` + "\n"
	config += `	ikev2_mode = "TUNNEL"` + "\n"
	config += `	ikev2_ipsec_proposals = [{` + "\n"
	config += `		id = fmc_ikev2_ipsec_proposals.test.items["my_s2s_vpn_ipsec_settings"].id` + "\n"
	config += `		name = "my_s2s_vpn_ipsec_settings"` + "\n"
	config += `	}]` + "\n"
	config += `	security_association_strength_enforcement = false` + "\n"
	config += `	reverse_route_injection = true` + "\n"
	config += `	perfect_forward_secrecy = true` + "\n"
	config += `	perfect_forward_secrecy_modulus_group = "14"` + "\n"
	config += `	lifetime_duration = 28800` + "\n"
	config += `	lifetime_size = 4608000` + "\n"
	config += `	validate_incoming_icmp_error_messages = false` + "\n"
	config += `	do_not_fragment_policy = "NONE"` + "\n"
	config += `	tfc = true` + "\n"
	config += `	tfc_burst_bytes = 0` + "\n"
	config += `	tfc_payload_bytes = 0` + "\n"
	config += `	tfc_timeout = 0` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vpn_s2s_ipsec_settings" "test" {
			id = fmc_vpn_s2s_ipsec_settings.test.id
			vpn_s2s_id = fmc_vpn_s2s.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
