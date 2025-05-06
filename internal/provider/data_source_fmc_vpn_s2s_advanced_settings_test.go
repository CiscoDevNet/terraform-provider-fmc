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

func TestAccDataSourceFmcVPNS2SAdvancedSettings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_vpn_s2s_advanced_settings.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ike_keepalive", "ENABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ike_keepalive_threshold", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ike_keepalive_retry_interval", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ike_identity_sent_to_peers", "AUTO_OR_DN"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ike_peer_identity_validation", "REQUIRED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ike_enable_aggressive_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ike_enable_notification_on_tunnel_disconnect", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ikev2_cookie_challenge", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ikev2_threshold_to_challenge_incoming_cookies", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ikev2_percentage_of_sas_allowed_in_negotiation", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ikev2_maximum_number_of_sas_allowed_in_negotiation", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ipsec_enable_fragmentation_before_encryption", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "ipsec_path_maximum_transmission_unit_aging_reset_interval", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "nat_keepalive_message_traversal_interval", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "vpn_idle_timeout", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "bypass_access_control_traffic_for_decrypted_traffic", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "use_cert_map_configured_in_endpoint_to_determine_tunnel", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "use_certificate_ou_to_determine_tunnel", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "use_ike_identity_ou_to_determine_tunnel", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_vpn_s2s_advanced_settings.test", "use_peer_ip_address_to_determine_tunnel", "false"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcVPNS2SAdvancedSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcVPNS2SAdvancedSettingsConfig() string {
	config := `resource "fmc_vpn_s2s_advanced_settings" "test" {` + "\n"
	config += `	vpn_s2s_id = TBD` + "\n"
	config += `	ike_keepalive = "ENABLED"` + "\n"
	config += `	ike_keepalive_threshold = 10` + "\n"
	config += `	ike_keepalive_retry_interval = 2` + "\n"
	config += `	ike_identity_sent_to_peers = "AUTO_OR_DN"` + "\n"
	config += `	ike_peer_identity_validation = "REQUIRED"` + "\n"
	config += `	ike_enable_aggressive_mode = false` + "\n"
	config += `	ike_enable_notification_on_tunnel_disconnect = false` + "\n"
	config += `	ikev2_cookie_challenge = "CUSTOM"` + "\n"
	config += `	ikev2_threshold_to_challenge_incoming_cookies = 50` + "\n"
	config += `	ikev2_percentage_of_sas_allowed_in_negotiation = 100` + "\n"
	config += `	ikev2_maximum_number_of_sas_allowed_in_negotiation = ` + "\n"
	config += `	ipsec_enable_fragmentation_before_encryption = true` + "\n"
	config += `	ipsec_path_maximum_transmission_unit_aging_reset_interval = ` + "\n"
	config += `	nat_keepalive_message_traversal_interval = 20` + "\n"
	config += `	vpn_idle_timeout = 10` + "\n"
	config += `	bypass_access_control_traffic_for_decrypted_traffic = false` + "\n"
	config += `	use_cert_map_configured_in_endpoint_to_determine_tunnel = false` + "\n"
	config += `	use_certificate_ou_to_determine_tunnel = false` + "\n"
	config += `	use_ike_identity_ou_to_determine_tunnel = false` + "\n"
	config += `	use_peer_ip_address_to_determine_tunnel = false` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_vpn_s2s_advanced_settings" "test" {
			id = fmc_vpn_s2s_advanced_settings.test.id
			vpn_s2s_id = TBD
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
