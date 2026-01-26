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

func TestAccDataSourceFmcGroupPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "name", "my_group_policy"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_group_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "description", "My Group Policy object"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "protocol_ssl", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "protocol_ipsec_ikev2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "banner", "Welcome to the VPN Connection."))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "default_domain", "example.com"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "ipv4_split_tunnel_policy", "TUNNEL_ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "ipv6_split_tunnel_policy", "TUNNEL_ALL"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "dns_request_split_tunnel_policy", "TUNNEL_SPECIFIED_DOMAINS"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "dns_request_split_tunnel_domains", "example.com,example.org"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "ssl_compression", "DISABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "dtls_compression", "DISABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "mtu_size", "1406"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "ignore_df_bit", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "keep_alive_messages", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "keep_alive_messages_interval", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "gateway_dead_peer_detection", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "gateway_dead_peer_detection_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "client_dead_peer_detection", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "client_dead_peer_detection_interval", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "client_bypass_protocol", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "ssl_rekey", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "ssl_rekey_method", "NEW_TUNNEL"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "ssl_rekey_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "restrict_vpn_to_vlan", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "simultaneous_logins_per_user", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "maximum_connection_time", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "maximum_connection_time_alert_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "idle_timeout", "3600"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_group_policy.test", "idle_timeout_alert_interval", "1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcGroupPolicyPrerequisitesConfig + testAccDataSourceFmcGroupPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcGroupPolicyPrerequisitesConfig + testAccNamedDataSourceFmcGroupPolicyConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcGroupPolicyPrerequisitesConfig = `
resource "fmc_ipv4_address_pool" "test" {
  name    = "group_policy_ipv4_address_pool1"
  range   = "10.10.10.0-10.10.10.240"
  netmask = "255.255.255.0"
}

resource "fmc_hosts" "test" {
  items = {
    "group_policy_host_1" = { ip = "192.168.10.1" }
    "group_policy_host_2" = { ip = "192.168.10.2" }
    "group_policy_host_3" = { ip = "192.168.10.3" }
    "group_policy_host_4" = { ip = "192.168.10.4" }
  }
}

resource "fmc_networks" "test" {
  items = {
    "group_policy_network_1" = { prefix = "192.168.20.0/24" }
  }
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcGroupPolicyConfig() string {
	config := `resource "fmc_group_policy" "test" {` + "\n"
	config += `	name = "my_group_policy"` + "\n"
	config += `	description = "My Group Policy object"` + "\n"
	config += `	protocol_ssl = true` + "\n"
	config += `	protocol_ipsec_ikev2 = true` + "\n"
	config += `	ipv4_address_pools = [{` + "\n"
	config += `		id = fmc_ipv4_address_pool.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	banner = "Welcome to the VPN Connection."` + "\n"
	config += `	primary_dns_server_host_id = fmc_hosts.test.items["group_policy_host_1"].id` + "\n"
	config += `	secondary_dns_server_host_id = fmc_hosts.test.items["group_policy_host_2"].id` + "\n"
	config += `	ipv4_dhcp_network_scope_network_object_id = fmc_networks.test.items["group_policy_network_1"].id` + "\n"
	config += `	default_domain = "example.com"` + "\n"
	config += `	ipv4_split_tunnel_policy = "TUNNEL_ALL"` + "\n"
	config += `	ipv6_split_tunnel_policy = "TUNNEL_ALL"` + "\n"
	config += `	dns_request_split_tunnel_policy = "TUNNEL_SPECIFIED_DOMAINS"` + "\n"
	config += `	dns_request_split_tunnel_domains = "example.com,example.org"` + "\n"
	config += `	ssl_compression = "DISABLED"` + "\n"
	config += `	dtls_compression = "DISABLED"` + "\n"
	config += `	mtu_size = 1406` + "\n"
	config += `	ignore_df_bit = true` + "\n"
	config += `	keep_alive_messages = true` + "\n"
	config += `	keep_alive_messages_interval = 20` + "\n"
	config += `	gateway_dead_peer_detection = true` + "\n"
	config += `	gateway_dead_peer_detection_interval = 30` + "\n"
	config += `	client_dead_peer_detection = true` + "\n"
	config += `	client_dead_peer_detection_interval = 30` + "\n"
	config += `	client_bypass_protocol = false` + "\n"
	config += `	ssl_rekey = true` + "\n"
	config += `	ssl_rekey_method = "NEW_TUNNEL"` + "\n"
	config += `	ssl_rekey_interval = 60` + "\n"
	config += `	restrict_vpn_to_vlan = 100` + "\n"
	config += `	simultaneous_logins_per_user = 3` + "\n"
	config += `	maximum_connection_time = 3600` + "\n"
	config += `	maximum_connection_time_alert_interval = 1` + "\n"
	config += `	idle_timeout = 3600` + "\n"
	config += `	idle_timeout_alert_interval = 1` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_group_policy" "test" {
			id = fmc_group_policy.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcGroupPolicyConfig() string {
	config := `resource "fmc_group_policy" "test" {` + "\n"
	config += `	name = "my_group_policy"` + "\n"
	config += `	description = "My Group Policy object"` + "\n"
	config += `	protocol_ssl = true` + "\n"
	config += `	protocol_ipsec_ikev2 = true` + "\n"
	config += `	ipv4_address_pools = [{` + "\n"
	config += `		id = fmc_ipv4_address_pool.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	banner = "Welcome to the VPN Connection."` + "\n"
	config += `	primary_dns_server_host_id = fmc_hosts.test.items["group_policy_host_1"].id` + "\n"
	config += `	secondary_dns_server_host_id = fmc_hosts.test.items["group_policy_host_2"].id` + "\n"
	config += `	ipv4_dhcp_network_scope_network_object_id = fmc_networks.test.items["group_policy_network_1"].id` + "\n"
	config += `	default_domain = "example.com"` + "\n"
	config += `	ipv4_split_tunnel_policy = "TUNNEL_ALL"` + "\n"
	config += `	ipv6_split_tunnel_policy = "TUNNEL_ALL"` + "\n"
	config += `	dns_request_split_tunnel_policy = "TUNNEL_SPECIFIED_DOMAINS"` + "\n"
	config += `	dns_request_split_tunnel_domains = "example.com,example.org"` + "\n"
	config += `	ssl_compression = "DISABLED"` + "\n"
	config += `	dtls_compression = "DISABLED"` + "\n"
	config += `	mtu_size = 1406` + "\n"
	config += `	ignore_df_bit = true` + "\n"
	config += `	keep_alive_messages = true` + "\n"
	config += `	keep_alive_messages_interval = 20` + "\n"
	config += `	gateway_dead_peer_detection = true` + "\n"
	config += `	gateway_dead_peer_detection_interval = 30` + "\n"
	config += `	client_dead_peer_detection = true` + "\n"
	config += `	client_dead_peer_detection_interval = 30` + "\n"
	config += `	client_bypass_protocol = false` + "\n"
	config += `	ssl_rekey = true` + "\n"
	config += `	ssl_rekey_method = "NEW_TUNNEL"` + "\n"
	config += `	ssl_rekey_interval = 60` + "\n"
	config += `	restrict_vpn_to_vlan = 100` + "\n"
	config += `	simultaneous_logins_per_user = 3` + "\n"
	config += `	maximum_connection_time = 3600` + "\n"
	config += `	maximum_connection_time_alert_interval = 1` + "\n"
	config += `	idle_timeout = 3600` + "\n"
	config += `	idle_timeout_alert_interval = 1` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_group_policy" "test" {
			name = fmc_group_policy.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
