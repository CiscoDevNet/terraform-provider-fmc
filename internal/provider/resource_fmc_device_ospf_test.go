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

func TestAccFmcDeviceOSPF(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_ospf.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "router_id", "AUTOMATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "rfc_1583_compatible", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "log_adjacency_changes", "DEFAULT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "ignore_lsa_mospf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "administrative_distance_inter_area", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "administrative_distance_intra_area", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "administrative_distance_external", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "timer_lsa_group", "240"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "default_route_always_advertise", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "default_route_metric", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "default_route_metric_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "default_route_route_map_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.type", "normal"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.networks.0.id", "123e4567-e89b-12d3-a456-426614174000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.networks.0.name", "Office Network"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.authentication", "MESSAGE_DIGEST"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.default_cost", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.ranges.0.network_object_id", "123e4567-e89b-12d3-a456-426614174000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.ranges.0.advertise", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.virtual_links.0.peer_router_host_id", "123e4567-e89b-12d3-a456-426614174000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.virtual_links.0.hello_interval", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.virtual_links.0.transmit_delay", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.virtual_links.0.retransmit_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.virtual_links.0.dead_interval", "40"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.virtual_links.0.authentication_password", "ospfAuthKey"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.inter_area_filters.0.prefix_list_id", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.inter_area_filters.0.prefix_list_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.inter_area_filters.0.filter_direction", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.route_type", "RedistributeBGP"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.as_number", "65001"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.ospf_match_external_1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.ospf_match_external_2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.ospf_match_internal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.ospf_match_nssa_external_1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.ospf_match_nssa_external_2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.process_id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.subnets", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.metric", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.metric_type", "TYPE_2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "redistributions.0.route_map_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "filter_rules.0.access_list_id", "123e4567-e89b-12d3-a456-426614174000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "filter_rules.0.traffic_direction", "incomingroutefilter"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "filter_rules.0.routing_process", "CONNECTED"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "summary_addresses.0.networks.0.id", "123e4567-e89b-12d3-a456-426614174000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "summary_addresses.0.tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "summary_addresses.0.advertise", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceOSPFConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceOSPFConfig_all(),
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

func testAccFmcDeviceOSPFConfig_minimum() string {
	config := `resource "fmc_device_ospf" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceOSPFConfig_all() string {
	config := `resource "fmc_device_ospf" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	process_id = 1` + "\n"
	config += `	router_id = "AUTOMATIC"` + "\n"
	config += `	rfc_1583_compatible = false` + "\n"
	config += `	log_adjacency_changes = "DEFAULT"` + "\n"
	config += `	ignore_lsa_mospf = false` + "\n"
	config += `	administrative_distance_inter_area = 110` + "\n"
	config += `	administrative_distance_intra_area = 110` + "\n"
	config += `	administrative_distance_external = 110` + "\n"
	config += `	timer_lsa_group = 240` + "\n"
	config += `	default_route_always_advertise = false` + "\n"
	config += `	default_route_metric = 1` + "\n"
	config += `	default_route_metric_type = ""` + "\n"
	config += `	default_route_route_map_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	areas = [{` + "\n"
	config += `		id = "1"` + "\n"
	config += `		type = "normal"` + "\n"
	config += `		networks = [{` + "\n"
	config += `			id = "123e4567-e89b-12d3-a456-426614174000"` + "\n"
	config += `			name = "Office Network"` + "\n"
	config += `		}]` + "\n"
	config += `		authentication = "MESSAGE_DIGEST"` + "\n"
	config += `		default_cost = 1` + "\n"
	config += `		ranges = [{` + "\n"
	config += `			network_object_id = "123e4567-e89b-12d3-a456-426614174000"` + "\n"
	config += `			advertise = true` + "\n"
	config += `		}]` + "\n"
	config += `		virtual_links = [{` + "\n"
	config += `			peer_router_host_id = "123e4567-e89b-12d3-a456-426614174000"` + "\n"
	config += `			hello_interval = 10` + "\n"
	config += `			transmit_delay = 1` + "\n"
	config += `			retransmit_interval = 5` + "\n"
	config += `			dead_interval = 40` + "\n"
	config += `			authentication_password = "ospfAuthKey"` + "\n"
	config += `		}]` + "\n"
	config += `		inter_area_filters = [{` + "\n"
	config += `			prefix_list_id = ""` + "\n"
	config += `			prefix_list_name = ""` + "\n"
	config += `			filter_direction = ""` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	redistributions = [{` + "\n"
	config += `		route_type = "RedistributeBGP"` + "\n"
	config += `		as_number = 65001` + "\n"
	config += `		ospf_match_external_1 = true` + "\n"
	config += `		ospf_match_external_2 = true` + "\n"
	config += `		ospf_match_internal = true` + "\n"
	config += `		ospf_match_nssa_external_1 = true` + "\n"
	config += `		ospf_match_nssa_external_2 = true` + "\n"
	config += `		process_id = 2` + "\n"
	config += `		subnets = true` + "\n"
	config += `		metric = 1` + "\n"
	config += `		metric_type = "TYPE_2"` + "\n"
	config += `		tag = 100` + "\n"
	config += `		route_map_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `	}]` + "\n"
	config += `	filter_rules = [{` + "\n"
	config += `		access_list_id = "123e4567-e89b-12d3-a456-426614174000"` + "\n"
	config += `		traffic_direction = "incomingroutefilter"` + "\n"
	config += `		routing_process = "CONNECTED"` + "\n"
	config += `	}]` + "\n"
	config += `	summary_addresses = [{` + "\n"
	config += `		networks = [{` + "\n"
	config += `			id = "123e4567-e89b-12d3-a456-426614174000"` + "\n"
	config += `		}]` + "\n"
	config += `		tag = 100` + "\n"
	config += `		advertise = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
