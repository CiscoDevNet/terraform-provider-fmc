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

func TestAccDataSourceFmcVPNRAIPSecCryptoMap(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_ra_ipsec_crypto_map.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "interface_id", "76d24097-41c4-4558-a4d0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "ikev2_ipsec_proposals.0.id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "client_services", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "client_services_port", "443"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "perfect_forward_secrecy", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "perfect_forward_secrecy_modulus_group", "14"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "lifetime_duration", "28800"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "lifetime_size", "4608000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "validate_incoming_icmp_error_messages", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "do_not_fragment_policy", "NONE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "tfc", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "tfc_burst_bytes", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "tfc_payload_bytes", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_crypto_map.test", "tfc_timeout", "0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVPNRAIPSecCryptoMapConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccNamedDataSourceFmcVPNRAIPSecCryptoMapConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVPNRAIPSecCryptoMapConfig() string {
	config := `resource "fmc_vpn_ra_ipsec_crypto_map" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	interface_id = "76d24097-41c4-4558-a4d0"` + "\n"
	config += `	ikev2_ipsec_proposals = [{` + "\n"
	config += `		id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	}]` + "\n"
	config += `	client_services = true` + "\n"
	config += `	client_services_port = 443` + "\n"
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
		data "fmc_vpn_ra_ipsec_crypto_map" "test" {
			id = fmc_vpn_ra_ipsec_crypto_map.test.id
			vpn_ra_id = fmc_vpn_ra.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcVPNRAIPSecCryptoMapConfig() string {
	config := `resource "fmc_vpn_ra_ipsec_crypto_map" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	interface_id = "76d24097-41c4-4558-a4d0"` + "\n"
	config += `	ikev2_ipsec_proposals = [{` + "\n"
	config += `		id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	}]` + "\n"
	config += `	client_services = true` + "\n"
	config += `	client_services_port = 443` + "\n"
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
		data "fmc_vpn_ra_ipsec_crypto_map" "test" {
			vpn_ra_id = fmc_vpn_ra.test.id
			interface_id = fmc_vpn_ra_ipsec_crypto_map.test.interface_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
