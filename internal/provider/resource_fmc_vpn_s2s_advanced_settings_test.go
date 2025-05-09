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

func TestAccFmcVPNS2SAdvancedSettings(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_s2s_advanced_settings.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ike_keepalive", "ENABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ike_keepalive_threshold", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ike_keepalive_retry_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ike_identity_sent_to_peers", "AUTO_OR_DN"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ike_peer_identity_validation", "DO_NOT_CHECK"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ike_aggressive_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ike_notification_on_tunnel_disconnect", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ikev2_cookie_challenge", "CUSTOM"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ikev2_threshold_to_challenge_incoming_cookies", "55"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ikev2_number_of_sas_allowed_in_negotiation", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ipsec_fragmentation_before_encryption", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "ipsec_path_maximum_transmission_unit_aging_reset_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "nat_keepalive_message_traversal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "nat_keepalive_message_traversal_interval", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "vpn_idle_timeout", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "vpn_idle_timeout_value", "25"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "bypass_access_control_traffic_for_decrypted_traffic", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "cert_use_map_configured_in_endpoint_to_determine_tunnel", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "cert_use_ou_to_determine_tunnel", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "cert_use_ike_identity_to_determine_tunnel", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_advanced_settings.test", "cert_use_peer_ip_address_to_determine_tunnel", "false"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNS2SAdvancedSettingsPrerequisitesConfig + testAccFmcVPNS2SAdvancedSettingsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNS2SAdvancedSettingsPrerequisitesConfig + testAccFmcVPNS2SAdvancedSettingsConfig_all(),
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

const testAccFmcVPNS2SAdvancedSettingsPrerequisitesConfig = `
resource "fmc_vpn_s2s" "test" {
  name             = "my_s2s_vpn_ike_settings"
  route_based      = false
  network_topology = "POINT_TO_POINT"
  ikev1            = true
  ikev2            = true
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcVPNS2SAdvancedSettingsConfig_minimum() string {
	config := `resource "fmc_vpn_s2s_advanced_settings" "test" {` + "\n"
	config += `	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNS2SAdvancedSettingsConfig_all() string {
	config := `resource "fmc_vpn_s2s_advanced_settings" "test" {` + "\n"
	config += `	vpn_s2s_id = fmc_vpn_s2s.test.id` + "\n"
	config += `	ike_keepalive = "ENABLED"` + "\n"
	config += `	ike_keepalive_threshold = 15` + "\n"
	config += `	ike_keepalive_retry_interval = 5` + "\n"
	config += `	ike_identity_sent_to_peers = "AUTO_OR_DN"` + "\n"
	config += `	ike_peer_identity_validation = "DO_NOT_CHECK"` + "\n"
	config += `	ike_aggressive_mode = false` + "\n"
	config += `	ike_notification_on_tunnel_disconnect = true` + "\n"
	config += `	ikev2_cookie_challenge = "CUSTOM"` + "\n"
	config += `	ikev2_threshold_to_challenge_incoming_cookies = 55` + "\n"
	config += `	ikev2_number_of_sas_allowed_in_negotiation = 90` + "\n"
	config += `	ikev2_maximum_number_of_sas_allowed = 50` + "\n"
	config += `	ipsec_fragmentation_before_encryption = true` + "\n"
	config += `	ipsec_path_maximum_transmission_unit_aging_reset_interval = 30` + "\n"
	config += `	nat_keepalive_message_traversal = true` + "\n"
	config += `	nat_keepalive_message_traversal_interval = 20` + "\n"
	config += `	vpn_idle_timeout = true` + "\n"
	config += `	vpn_idle_timeout_value = 25` + "\n"
	config += `	bypass_access_control_traffic_for_decrypted_traffic = false` + "\n"
	config += `	cert_use_map_configured_in_endpoint_to_determine_tunnel = false` + "\n"
	config += `	cert_use_ou_to_determine_tunnel = false` + "\n"
	config += `	cert_use_ike_identity_to_determine_tunnel = false` + "\n"
	config += `	cert_use_peer_ip_address_to_determine_tunnel = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
