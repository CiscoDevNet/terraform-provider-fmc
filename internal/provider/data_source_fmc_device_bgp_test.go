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

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource

func TestAccDataSourceFmcDeviceBGP(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_bgp.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_bgp.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_bgp.test", "as_number"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_bgp.test", "ipv4_address_family_id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_bgp.test", "ipv4_address_family_type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_default_information_orginate", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_auto_summary", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_suppress_inactive_routes", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_synchronization", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_redistribute_ibgp_into_igp", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_external_distance", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_internal_distance", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_local_distance", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_number_of_ibgp_paths", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_number_of_ebgp_paths", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.remote_as", "65534"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.bfd_fallover", "SINGLE_HOP"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.enable_address", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.shutdown_administratively", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.description", "My BGP Peer"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.filter_maximum_prefixes", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.filter_threshold_value", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.filter_restart_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.routes_advertisement_interval", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.routes_remove_private_as", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.keepalive_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.hold_time", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.minimum_hold_time", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.send_community_attribute", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.next_hop_self", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.disable_connection_verification", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.tcp_path_mtu_discovery", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.max_hop_count", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.tcp_transport_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.weight", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp.test", "ipv4_neighbors.0.version", "0"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceBGPPrerequisitesConfig + testAccDataSourceFmcDeviceBGPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceBGPPrerequisitesConfig + testAccNamedDataSourceFmcDeviceBGPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceBGPPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

resource "fmc_device_bgp_general_settings" "test" {
  device_id                   = var.device_id
  as_number                   = "6353"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceBGPConfig() string {
	config := `resource "fmc_device_bgp" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	ipv4_default_information_orginate = false` + "\n"
	config += `	ipv4_auto_summary = false` + "\n"
	config += `	ipv4_suppress_inactive_routes = false` + "\n"
	config += `	ipv4_synchronization = false` + "\n"
	config += `	ipv4_redistribute_ibgp_into_igp = false` + "\n"
	config += `	ipv4_external_distance = 20` + "\n"
	config += `	ipv4_internal_distance = 200` + "\n"
	config += `	ipv4_local_distance = 200` + "\n"
	config += `	ipv4_number_of_ibgp_paths = 1` + "\n"
	config += `	ipv4_number_of_ebgp_paths = 1` + "\n"
	config += `	ipv4_neighbors = [{` + "\n"
	config += `		address = "10.1.1.1"` + "\n"
	config += `		remote_as = "65534"` + "\n"
	config += `		bfd_fallover = "SINGLE_HOP"` + "\n"
	config += `		enable_address = true` + "\n"
	config += `		shutdown_administratively = false` + "\n"
	config += `		description = "My BGP Peer"` + "\n"
	config += `		filter_maximum_prefixes = 1` + "\n"
	config += `		filter_threshold_value = 1` + "\n"
	config += `		filter_restart_interval = 1` + "\n"
	config += `		routes_advertisement_interval = 1` + "\n"
	config += `		routes_remove_private_as = false` + "\n"
	config += `		keepalive_interval = 60` + "\n"
	config += `		hold_time = 180` + "\n"
	config += `		minimum_hold_time = 3` + "\n"
	config += `		send_community_attribute = false` + "\n"
	config += `		next_hop_self = false` + "\n"
	config += `		disable_connection_verification = false` + "\n"
	config += `		tcp_path_mtu_discovery = false` + "\n"
	config += `		max_hop_count = 1` + "\n"
	config += `		tcp_transport_mode = false` + "\n"
	config += `		weight = 0` + "\n"
	config += `		version = "0"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bgp" "test" {
			id = fmc_device_bgp.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceBGPConfig() string {
	config := `resource "fmc_device_bgp" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	ipv4_default_information_orginate = false` + "\n"
	config += `	ipv4_auto_summary = false` + "\n"
	config += `	ipv4_suppress_inactive_routes = false` + "\n"
	config += `	ipv4_synchronization = false` + "\n"
	config += `	ipv4_redistribute_ibgp_into_igp = false` + "\n"
	config += `	ipv4_external_distance = 20` + "\n"
	config += `	ipv4_internal_distance = 200` + "\n"
	config += `	ipv4_local_distance = 200` + "\n"
	config += `	ipv4_number_of_ibgp_paths = 1` + "\n"
	config += `	ipv4_number_of_ebgp_paths = 1` + "\n"
	config += `	ipv4_neighbors = [{` + "\n"
	config += `		address = "10.1.1.1"` + "\n"
	config += `		remote_as = "65534"` + "\n"
	config += `		bfd_fallover = "SINGLE_HOP"` + "\n"
	config += `		enable_address = true` + "\n"
	config += `		shutdown_administratively = false` + "\n"
	config += `		description = "My BGP Peer"` + "\n"
	config += `		filter_maximum_prefixes = 1` + "\n"
	config += `		filter_threshold_value = 1` + "\n"
	config += `		filter_restart_interval = 1` + "\n"
	config += `		routes_advertisement_interval = 1` + "\n"
	config += `		routes_remove_private_as = false` + "\n"
	config += `		keepalive_interval = 60` + "\n"
	config += `		hold_time = 180` + "\n"
	config += `		minimum_hold_time = 3` + "\n"
	config += `		send_community_attribute = false` + "\n"
	config += `		next_hop_self = false` + "\n"
	config += `		disable_connection_verification = false` + "\n"
	config += `		tcp_path_mtu_discovery = false` + "\n"
	config += `		max_hop_count = 1` + "\n"
	config += `		tcp_transport_mode = false` + "\n"
	config += `		weight = 0` + "\n"
	config += `		version = "0"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bgp" "test" {
			device_id = var.device_id
			as_number = fmc_device_bgp.test.as_number
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
