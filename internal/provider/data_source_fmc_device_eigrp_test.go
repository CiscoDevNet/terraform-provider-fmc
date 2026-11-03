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

func TestAccDataSourceFmcDeviceEIGRP(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_eigrp.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "as_number", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "auto_summary", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "filter_rules.0.filter_direction", "INBOUND"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "redistributions.0.redistribute_protocol", "STATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "redistributions.0.metric_bandwidth", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "redistributions.0.metric_delay", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "redistributions.0.metric_reliability", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "redistributions.0.metric_loading", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "redistributions.0.metric_mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "summary_addresses.0.administrative_distance", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "interfaces.0.hello_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "interfaces.0.hold_time", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "interfaces.0.split_horizon", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "interfaces.0.delay_time", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "interfaces.0.md5_authentication", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "interfaces.0.authentication_key_type", "UNENCRYPTED"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "interfaces.0.authentication_key_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "router_id", "10.10.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "administrative_distance_internal", "90"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "administrative_distance_external", "170"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "log_neighbor_changes", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "log_neighbor_warnings", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "log_neighbor_warnings_interval", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "stub_receive_only", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "stub_advertise_connected", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "stub_advertise_redistributed", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "stub_advertise_static", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "stub_advertise_summary", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "default_metric_bandwidth", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "default_metric_delay_time", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "default_metric_reliability", "255"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "default_metric_loading", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_eigrp.test", "default_metric_mtu", "1500"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceEIGRPPrerequisitesConfig + testAccDataSourceFmcDeviceEIGRPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceEIGRPPrerequisitesConfig + testAccNamedByAsNumberDataSourceFmcDeviceEIGRPConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceEIGRPPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" { default = null } // tests will set $TF_VAR_interface_name

resource "fmc_device_physical_interface" "test" {
  device_id           = var.device_id
  name                = var.interface_name
  logical_name        = "eigrp_interface"
  mode                = "NONE"
  enabled             = true
  ipv4_static_address = "10.99.1.1"
  ipv4_static_netmask = "24"
}

resource "fmc_networks" "test" {
  items = {
    "eigrp_network_1" = { prefix = "10.99.1.0/24" },
    "eigrp_network_2" = { prefix = "10.99.0.0/16" },
  }
}

resource "fmc_hosts" "test" {
  items = {
    "eigrp_neighbor" = { ip = "10.99.1.2" },
  }
}

resource "fmc_standard_access_list" "test" {
  name = "eigrp_std_acl"
  entries = [
    { action = "PERMIT", literals = [{ value = "10.1.1.0/24" }] }
  ]
}

resource "fmc_route_map" "test" {
  name    = "eigrp_route_map"
  entries = [
    { action = "PERMIT" }
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceEIGRPConfig() string {
	config := `resource "fmc_device_eigrp" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	as_number = "100"` + "\n"
	config += `	networks = [{` + "\n"
	config += `		id = fmc_networks.test.items.eigrp_network_1.id` + "\n"
	config += `	}]` + "\n"
	config += `	auto_summary = false` + "\n"
	config += `	passive_interfaces_all = false` + "\n"
	config += `	passive_interfaces = [{` + "\n"
	config += `		id = fmc_device_physical_interface.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	neighbors = [{` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		host_id = fmc_hosts.test.items.eigrp_neighbor.id` + "\n"
	config += `	}]` + "\n"
	config += `	filter_rules = [{` + "\n"
	config += `		filter_direction = "INBOUND"` + "\n"
	config += `		standard_access_list_id = fmc_standard_access_list.test.id` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	redistributions = [{` + "\n"
	config += `		redistribute_protocol = "STATIC"` + "\n"
	config += `		route_map_id = fmc_route_map.test.id` + "\n"
	config += `		metric_bandwidth = 1000` + "\n"
	config += `		metric_delay = 100` + "\n"
	config += `		metric_reliability = 255` + "\n"
	config += `		metric_loading = 1` + "\n"
	config += `		metric_mtu = 1500` + "\n"
	config += `	}]` + "\n"
	config += `	summary_addresses = [{` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		network_id = fmc_networks.test.items.eigrp_network_2.id` + "\n"
	config += `		administrative_distance = 5` + "\n"
	config += `	}]` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		interface_type = fmc_device_physical_interface.test.type` + "\n"
	config += `		hello_interval = 10` + "\n"
	config += `		hold_time = 30` + "\n"
	config += `		split_horizon = false` + "\n"
	config += `		delay_time = 20` + "\n"
	config += `		md5_authentication = true` + "\n"
	config += `		authentication_key_type = "UNENCRYPTED"` + "\n"
	config += `		authentication_key_id = 2` + "\n"
	config += `		authentication_key = "my_secret"` + "\n"
	config += `	}]` + "\n"
	config += `	router_id = "10.10.10.1"` + "\n"
	config += `	default_information_in_standard_access_list_id = fmc_standard_access_list.test.id` + "\n"
	config += `	default_information_out_standard_access_list_id = fmc_standard_access_list.test.id` + "\n"
	config += `	administrative_distance_internal = 90` + "\n"
	config += `	administrative_distance_external = 170` + "\n"
	config += `	log_neighbor_changes = true` + "\n"
	config += `	log_neighbor_warnings = true` + "\n"
	config += `	log_neighbor_warnings_interval = 15` + "\n"
	config += `	stub_receive_only = false` + "\n"
	config += `	stub_advertise_connected = true` + "\n"
	config += `	stub_advertise_redistributed = false` + "\n"
	config += `	stub_advertise_static = false` + "\n"
	config += `	stub_advertise_summary = true` + "\n"
	config += `	default_metric_bandwidth = 1000` + "\n"
	config += `	default_metric_delay_time = 100` + "\n"
	config += `	default_metric_reliability = 255` + "\n"
	config += `	default_metric_loading = 1` + "\n"
	config += `	default_metric_mtu = 1500` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_eigrp" "test" {
			id = fmc_device_eigrp.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedByAsNumberDataSourceFmcDeviceEIGRPConfig() string {
	config := `resource "fmc_device_eigrp" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	as_number = "100"` + "\n"
	config += `	networks = [{` + "\n"
	config += `		id = fmc_networks.test.items.eigrp_network_1.id` + "\n"
	config += `	}]` + "\n"
	config += `	auto_summary = false` + "\n"
	config += `	passive_interfaces_all = false` + "\n"
	config += `	passive_interfaces = [{` + "\n"
	config += `		id = fmc_device_physical_interface.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	neighbors = [{` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		host_id = fmc_hosts.test.items.eigrp_neighbor.id` + "\n"
	config += `	}]` + "\n"
	config += `	filter_rules = [{` + "\n"
	config += `		filter_direction = "INBOUND"` + "\n"
	config += `		standard_access_list_id = fmc_standard_access_list.test.id` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `	}]` + "\n"
	config += `	redistributions = [{` + "\n"
	config += `		redistribute_protocol = "STATIC"` + "\n"
	config += `		route_map_id = fmc_route_map.test.id` + "\n"
	config += `		metric_bandwidth = 1000` + "\n"
	config += `		metric_delay = 100` + "\n"
	config += `		metric_reliability = 255` + "\n"
	config += `		metric_loading = 1` + "\n"
	config += `		metric_mtu = 1500` + "\n"
	config += `	}]` + "\n"
	config += `	summary_addresses = [{` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		network_id = fmc_networks.test.items.eigrp_network_2.id` + "\n"
	config += `		administrative_distance = 5` + "\n"
	config += `	}]` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		interface_type = fmc_device_physical_interface.test.type` + "\n"
	config += `		hello_interval = 10` + "\n"
	config += `		hold_time = 30` + "\n"
	config += `		split_horizon = false` + "\n"
	config += `		delay_time = 20` + "\n"
	config += `		md5_authentication = true` + "\n"
	config += `		authentication_key_type = "UNENCRYPTED"` + "\n"
	config += `		authentication_key_id = 2` + "\n"
	config += `		authentication_key = "my_secret"` + "\n"
	config += `	}]` + "\n"
	config += `	router_id = "10.10.10.1"` + "\n"
	config += `	default_information_in_standard_access_list_id = fmc_standard_access_list.test.id` + "\n"
	config += `	default_information_out_standard_access_list_id = fmc_standard_access_list.test.id` + "\n"
	config += `	administrative_distance_internal = 90` + "\n"
	config += `	administrative_distance_external = 170` + "\n"
	config += `	log_neighbor_changes = true` + "\n"
	config += `	log_neighbor_warnings = true` + "\n"
	config += `	log_neighbor_warnings_interval = 15` + "\n"
	config += `	stub_receive_only = false` + "\n"
	config += `	stub_advertise_connected = true` + "\n"
	config += `	stub_advertise_redistributed = false` + "\n"
	config += `	stub_advertise_static = false` + "\n"
	config += `	stub_advertise_summary = true` + "\n"
	config += `	default_metric_bandwidth = 1000` + "\n"
	config += `	default_metric_delay_time = 100` + "\n"
	config += `	default_metric_reliability = 255` + "\n"
	config += `	default_metric_loading = 1` + "\n"
	config += `	default_metric_mtu = 1500` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_eigrp" "test" {
			device_id = var.device_id
			as_number = fmc_device_eigrp.test.as_number
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
