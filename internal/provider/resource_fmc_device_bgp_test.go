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

func TestAccFmcDeviceBGP(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp.test", "as_number"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_bgp.test", "ipv4_address_family_type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_default_information_orginate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_auto_aummary", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_bgp_supress_inactive", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_synchronization", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_bgp_redistribute_internal", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_external_distance", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_internal_distance", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_local_distance", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_forward_packets_over_multipath_ibgp", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_forward_packets_over_multipath_ebgp", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_remote_as", "65534"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_bfd", "SINGLE_HOP"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.enable_address_family", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_shutdown", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_description", "My BGP Peer"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_filter_max_prefix", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_filter_threshold_value", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_filter_restart_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_routes_advertisement_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_routes_remove_private_as", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_keepalive_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_hold_time", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_min_hold_time", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_send_community_attribute", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_nexthop_self", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_disable_connection_verification", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_tcp_mtu_path_discovery", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_max_hop_count", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_tcp_transport_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_weight", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_bgp.test", "ipv4_neighbors.0.neighbor_version", "0"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceBGPPrerequisitesConfig + testAccFmcDeviceBGPConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceBGPPrerequisitesConfig + testAccFmcDeviceBGPConfig_all(),
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

const testAccFmcDeviceBGPPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

resource "fmc_device_bgp_general_settings" "test" {
  device_id                   = var.device_id
  as_number                   = "6353"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceBGPConfig_minimum() string {
	config := `resource "fmc_device_bgp" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceBGPConfig_all() string {
	config := `resource "fmc_device_bgp" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	ipv4_default_information_orginate = false` + "\n"
	config += `	ipv4_auto_aummary = false` + "\n"
	config += `	ipv4_bgp_supress_inactive = false` + "\n"
	config += `	ipv4_synchronization = false` + "\n"
	config += `	ipv4_bgp_redistribute_internal = false` + "\n"
	config += `	ipv4_external_distance = 20` + "\n"
	config += `	ipv4_internal_distance = 200` + "\n"
	config += `	ipv4_local_distance = 200` + "\n"
	config += `	ipv4_forward_packets_over_multipath_ibgp = 1` + "\n"
	config += `	ipv4_forward_packets_over_multipath_ebgp = 1` + "\n"
	config += `	ipv4_neighbors = [{` + "\n"
	config += `		neighbor_address = "10.1.1.1"` + "\n"
	config += `		neighbor_remote_as = "65534"` + "\n"
	config += `		neighbor_bfd = "SINGLE_HOP"` + "\n"
	config += `		enable_address_family = true` + "\n"
	config += `		neighbor_shutdown = false` + "\n"
	config += `		neighbor_description = "My BGP Peer"` + "\n"
	config += `		neighbor_filter_max_prefix = 1` + "\n"
	config += `		neighbor_filter_threshold_value = 1` + "\n"
	config += `		neighbor_filter_restart_interval = 1` + "\n"
	config += `		neighbor_routes_advertisement_interval = 1` + "\n"
	config += `		neighbor_routes_remove_private_as = false` + "\n"
	config += `		neighbor_keepalive_interval = 60` + "\n"
	config += `		neighbor_hold_time = 180` + "\n"
	config += `		neighbor_min_hold_time = 3` + "\n"
	config += `		neighbor_send_community_attribute = false` + "\n"
	config += `		neighbor_nexthop_self = false` + "\n"
	config += `		neighbor_disable_connection_verification = false` + "\n"
	config += `		neighbor_tcp_mtu_path_discovery = false` + "\n"
	config += `		neighbor_max_hop_count = 1` + "\n"
	config += `		neighbor_tcp_transport_mode = false` + "\n"
	config += `		neighbor_weight = 0` + "\n"
	config += `		neighbor_version = "0"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
