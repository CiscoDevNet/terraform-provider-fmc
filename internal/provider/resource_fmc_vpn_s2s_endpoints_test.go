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

func TestAccFmcVPNS2SEndpoints(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.peer_type", "PEER"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.extranet_device", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.extranet_ip_address", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.extranet_dynamic_ip", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.interface_public_ip_address", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.connection_type", "BIDIRECTIONAL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.allow_incoming_ikev2_routes", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.send_vti_ip_to_peer", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.protected_networks_acl_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.nat_traversal", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.nat_exemption_inside_interface_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.reverse_route_injection", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.local_identity_type", "EMAIL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.local_identity_string", "me@cisco.com"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.vpn_filter_acl_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.override_remote_vpn_filter_acl_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.backup_interface_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.backup_interface_public_ip_address", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.backup_local_identity_type", "EMAIL"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_vpn_s2s_endpoints.test", "items.my_ftd_01.backup_local_identity_string", "me@cisco.com"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcVPNS2SEndpointsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcVPNS2SEndpointsConfig_all(),
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

func testAccFmcVPNS2SEndpointsConfig_minimum() string {
	config := `resource "fmc_vpn_s2s_endpoints" "test" {` + "\n"
	config += `	vpn_s2s_id = TBD` + "\n"
	config += `	items = { "my_ftd_01" = {` + "\n"
	config += `		peer_type = "PEER"` + "\n"
	config += `		extranet_device = false` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcVPNS2SEndpointsConfig_all() string {
	config := `resource "fmc_vpn_s2s_endpoints" "test" {` + "\n"
	config += `	vpn_s2s_id = TBD` + "\n"
	config += `	items = { "my_ftd_01" = {` + "\n"
	config += `		peer_type = "PEER"` + "\n"
	config += `		extranet_device = false` + "\n"
	config += `		extranet_ip_address = ""` + "\n"
	config += `		extranet_dynamic_ip = ` + "\n"
	config += `		device_id = TBD` + "\n"
	config += `		interface_id = TBD` + "\n"
	config += `		interface_ipv6_address = TBD` + "\n"
	config += `		interface_public_ip_address = ""` + "\n"
	config += `		connection_type = "BIDIRECTIONAL"` + "\n"
	config += `		allow_incoming_ikev2_routes = false` + "\n"
	config += `		send_vti_ip_to_peer = false` + "\n"
	config += `		protected_networks = [{` + "\n"
	config += `			id = TBD` + "\n"
	config += `		}]` + "\n"
	config += `		protected_networks_acl_id = ""` + "\n"
	config += `		nat_traversal = false` + "\n"
	config += `		nat_exemption = false` + "\n"
	config += `		nat_exemption_inside_interface_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `		reverse_route_injection = false` + "\n"
	config += `		local_identity_type = "EMAIL"` + "\n"
	config += `		local_identity_string = "me@cisco.com"` + "\n"
	config += `		vpn_filter_acl_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `		override_remote_vpn_filter_acl_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `		backup_interface_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `		backup_interface_public_ip_address = ""` + "\n"
	config += `		backup_local_identity_type = "EMAIL"` + "\n"
	config += `		backup_local_identity_string = "me@cisco.com"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
