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
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "vpn_protocol_ssl", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "vpn_protocol_ipsec_ikev2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ipv4_address_pools.0.id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "banner", "Welcome to the VPN Connection."))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "primary_dns_server_host_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secondary_dns_server_host_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "primary_wins_server_host_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secondary_wins_server_host_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "dhcp_network_scope_network_object_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "default_domain", "example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ipv4_split_tunnel_policy", "TUNNEL_ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ipv6_split_tunnel_policy", "TUNNEL_ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "split_tunnel_acl_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "split_tunnel_acl_type", "ExtendedAccessList"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "dns_request_split_tunnel_policy", "USE_SPLIT_TUNNEL_SETTING"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "split_dns_domain_list", "example.com,example.org"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secure_client_profile_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secure_client_management_profile_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secure_client_modules.0.type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secure_client_modules.0.profile_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secure_client_modules.0.download_module", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ssl_compression", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "dtls_compression", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "mtu_size", "1406"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ignore_df_bit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "keep_alive_messages", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "keep_alive_messages_interval", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "gateway_dpd", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "gateway_dpd_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "client_dpd", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "client_dpd_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "client_bypass_protocol", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ssl_rekey", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ssl_rekey_method", "NEW_TUNNEL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "ssl_rekey_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "client_firewall_private_network_rules_acl_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "client_firewall_public_network_rules_acl_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "secure_client_custom_attributes.0.id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "traffic_filter_acl_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "restrict_vpn_to_vlan", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "access_hours_time_range_id", "12345678-1234-1234-1234-123456"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "simultaneous_logins_per_user", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "maximum_connection_time", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_group_policy.test", "maximum_connection_time_alert_interval", "1"))
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
	config += `	vpn_protocol_ssl = true` + "\n"
	config += `	vpn_protocol_ipsec_ikev2 = true` + "\n"
	config += `	ipv4_address_pools = [{` + "\n"
	config += `		id = ""` + "\n"
	config += `	}]` + "\n"
	config += `	banner = "Welcome to the VPN Connection."` + "\n"
	config += `	primary_dns_server_host_id = ""` + "\n"
	config += `	secondary_dns_server_host_id = ""` + "\n"
	config += `	primary_wins_server_host_id = ""` + "\n"
	config += `	secondary_wins_server_host_id = ""` + "\n"
	config += `	dhcp_network_scope_network_object_id = ""` + "\n"
	config += `	default_domain = "example.com"` + "\n"
	config += `	ipv4_split_tunnel_policy = "TUNNEL_ALL"` + "\n"
	config += `	ipv6_split_tunnel_policy = "TUNNEL_ALL"` + "\n"
	config += `	split_tunnel_acl_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	split_tunnel_acl_type = "ExtendedAccessList"` + "\n"
	config += `	dns_request_split_tunnel_policy = "USE_SPLIT_TUNNEL_SETTING"` + "\n"
	config += `	split_dns_domain_list = "example.com,example.org"` + "\n"
	config += `	secure_client_profile_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	secure_client_management_profile_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	secure_client_modules = [{` + "\n"
	config += `		type = ""` + "\n"
	config += `		profile_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `		download_module = true` + "\n"
	config += `	}]` + "\n"
	config += `	ssl_compression = ""` + "\n"
	config += `	dtls_compression = ""` + "\n"
	config += `	mtu_size = 1406` + "\n"
	config += `	ignore_df_bit = true` + "\n"
	config += `	keep_alive_messages = true` + "\n"
	config += `	keep_alive_messages_interval = 20` + "\n"
	config += `	gateway_dpd = true` + "\n"
	config += `	gateway_dpd_interval = 30` + "\n"
	config += `	client_dpd = true` + "\n"
	config += `	client_dpd_interval = 30` + "\n"
	config += `	client_bypass_protocol = false` + "\n"
	config += `	ssl_rekey = true` + "\n"
	config += `	ssl_rekey_method = "NEW_TUNNEL"` + "\n"
	config += `	ssl_rekey_interval = 60` + "\n"
	config += `	client_firewall_private_network_rules_acl_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	client_firewall_public_network_rules_acl_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	secure_client_custom_attributes = [{` + "\n"
	config += `		id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	}]` + "\n"
	config += `	traffic_filter_acl_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	restrict_vpn_to_vlan = 100` + "\n"
	config += `	access_hours_time_range_id = "12345678-1234-1234-1234-123456"` + "\n"
	config += `	simultaneous_logins_per_user = 3` + "\n"
	config += `	maximum_connection_time = 3600` + "\n"
	config += `	maximum_connection_time_alert_interval = 1` + "\n"
	config += `	vpn_idle_timeout = 3600` + "\n"
	config += `	vpn_idle_timeout_alert_interval = 1` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
