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

func TestAccDataSourceFmcVPNRAIPSecIKEParameters(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_ra_ipsec_ike_parameters.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "ikev2_identity_sent_to_peer", "AUTO_OR_DN"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "ikev2_notification_on_tunnel_disconnect", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "ikev2_do_not_reboot_until_all_sessions_are_terminated", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "ikev2_cookie_challenge", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "ikev2_threshold_to_challenge_incoming_cookies", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "ikev2_number_of_sas_allowed_in_negotiation", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "ikev2_maximum_number_of_sas_allowed", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "ipsec_path_maximum_transmission_unit_aging_reset_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "nat_keepalive_message_traversal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_ra_ipsec_ike_parameters.test", "nat_keepalive_message_traversal_interval", "20"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVPNRAIPSecIKEParametersConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVPNRAIPSecIKEParametersConfig() string {
	config := `resource "fmc_vpn_ra_ipsec_ike_parameters" "test" {` + "\n"
	config += `	vpn_ra_id = fmc_vpn_ra.test.id` + "\n"
	config += `	ikev2_identity_sent_to_peer = "AUTO_OR_DN"` + "\n"
	config += `	ikev2_notification_on_tunnel_disconnect = true` + "\n"
	config += `	ikev2_do_not_reboot_until_all_sessions_are_terminated = false` + "\n"
	config += `	ikev2_cookie_challenge = "true"` + "\n"
	config += `	ikev2_threshold_to_challenge_incoming_cookies = 50` + "\n"
	config += `	ikev2_number_of_sas_allowed_in_negotiation = 90` + "\n"
	config += `	ikev2_maximum_number_of_sas_allowed = 50` + "\n"
	config += `	ipsec_path_maximum_transmission_unit_aging_reset_interval = 30` + "\n"
	config += `	nat_keepalive_message_traversal = true` + "\n"
	config += `	nat_keepalive_message_traversal_interval = 20` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vpn_ra_ipsec_ike_parameters" "test" {
			id = fmc_vpn_ra_ipsec_ike_parameters.test.id
			vpn_ra_id = fmc_vpn_ra.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
