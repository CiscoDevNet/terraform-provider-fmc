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

func TestAccFmcGroupPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "name", "my_group_policy"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_group_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "description", "My Group Policy object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "enable_ssl_protocol", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "enable_ipsec_ikev2_protocol", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ipv4_address_pools.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "banner", "Welcome to the VPN Connection."))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "primary_dns_server", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secondary_dns_server", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "primary_wins_server", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secondary_wins_server", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "dhcp_scope_network_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "default_domain", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ipv4_split_tunnel_policy", "TUNNEL_ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ipv6_split_tunnel_policy", "TUNNEL_ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "split_tunnel_acl_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "split_d_n_s_request_policy", "TUNNEL_ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "split_d_n_s_domain_list", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secure_client_client_profile_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secure_client_management_profile_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ssl_compression", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "dtls_compression", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "mtu_size", "1406"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ignore_df_bit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "enable_keep_alive_messages", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "keep_alive_message_interval", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "enable_gateway_dpd", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "gateway_dpd_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "enable_client_dpd", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "client_dpd_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "client_bypass_protocol", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "enable_ssl_rekey", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "rekey_method", "NEW_TUNNEL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "rekey_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "client_firewall_private_network_rules_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "client_firewall_public_network_rules_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "custom_attributes.0.id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "custom_attributes.0.type", "myAnyConnectAttribute"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "traffic_filter_acl_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "restrict_vpn_to_vlan_id", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "access_hours_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "simultaneous_login_per_user", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "max_connection_time", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "max_connection_time_alert_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "vpn_idle_timeout", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "vpn_idle_timeout_alert_interval", "1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcGroupPolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcGroupPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_group_policy.test",
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

func testAccFmcGroupPolicyConfig_minimum() string {
	config := `resource "fmc_group_policy" "test" {` + "\n"
	config += `	name = "my_group_policy"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcGroupPolicyConfig_all() string {
	config := `resource "fmc_group_policy" "test" {` + "\n"
	config += `	name = "my_group_policy"` + "\n"
	config += `	description = "My Group Policy object"` + "\n"
	config += `	enable_ssl_protocol = true` + "\n"
	config += `	enable_ipsec_ikev2_protocol = true` + "\n"
	config += `	ipv4_address_pools = [{` + "\n"
	config += `		id = ""` + "\n"
	config += `	}]` + "\n"
	config += `	banner = "Welcome to the VPN Connection."` + "\n"
	config += `	primary_dns_server = ""` + "\n"
	config += `	secondary_dns_server = ""` + "\n"
	config += `	primary_wins_server = ""` + "\n"
	config += `	secondary_wins_server = ""` + "\n"
	config += `	dhcp_scope_network_id = ""` + "\n"
	config += `	default_domain = ""` + "\n"
	config += `	ipv4_split_tunnel_policy = "TUNNEL_ALL"` + "\n"
	config += `	ipv6_split_tunnel_policy = "TUNNEL_ALL"` + "\n"
	config += `	split_tunnel_acl_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	split_d_n_s_request_policy = "TUNNEL_ALL"` + "\n"
	config += `	split_d_n_s_domain_list = ""` + "\n"
	config += `	secure_client_client_profile_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	secure_client_management_profile_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	ssl_compression = ""` + "\n"
	config += `	dtls_compression = ""` + "\n"
	config += `	mtu_size = 1406` + "\n"
	config += `	ignore_df_bit = true` + "\n"
	config += `	enable_keep_alive_messages = true` + "\n"
	config += `	keep_alive_message_interval = 20` + "\n"
	config += `	enable_gateway_dpd = true` + "\n"
	config += `	gateway_dpd_interval = 30` + "\n"
	config += `	enable_client_dpd = true` + "\n"
	config += `	client_dpd_interval = 30` + "\n"
	config += `	client_bypass_protocol = false` + "\n"
	config += `	enable_ssl_rekey = true` + "\n"
	config += `	rekey_method = "NEW_TUNNEL"` + "\n"
	config += `	rekey_interval = 60` + "\n"
	config += `	client_firewall_private_network_rules_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	client_firewall_public_network_rules_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	custom_attributes = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		type = "myAnyConnectAttribute"` + "\n"
	config += `	}]` + "\n"
	config += `	traffic_filter_acl_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	restrict_vpn_to_vlan_id = 100` + "\n"
	config += `	access_hours_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	simultaneous_login_per_user = 3` + "\n"
	config += `	max_connection_time = 3600` + "\n"
	config += `	max_connection_time_alert_interval = 1` + "\n"
	config += `	vpn_idle_timeout = 3600` + "\n"
	config += `	vpn_idle_timeout_alert_interval = 1` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
