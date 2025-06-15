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

func TestAccDataSourceFmcRouteMap(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "name", "my_route_map"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_route_map.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "overridable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.action", "PERMIT"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.match_interface_names.0", "GigabitEthernet0/1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.match_metric_route_values.0", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.match_tag_values.0", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.match_route_type_external_1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.match_route_type_external_2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.match_route_type_internal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.match_route_type_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.match_route_type_nssa_external_1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.match_route_type_nssa_external_2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_metric_bandwidth", "1000000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_metric_type", "INTERNAL"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_as_path_prepend.0", "64512"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_as_path_prepend_last_as", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_as_path_convert_route_tag_into_as_path", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_community_specific_community", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_community_add_to_existing_communities", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_community_internet", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_community_no_advertise", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_community_no_export", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_community_route_target", "100:100,200:200"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_community_add_to_existing_extended_communities", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_automatic_tag", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_local_preference", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_weight", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_route_map.test", "entries.0.set_bgp_origin", "LOCAL_IGP"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcRouteMapPrerequisitesConfig + testAccDataSourceFmcRouteMapConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcRouteMapPrerequisitesConfig + testAccNamedDataSourceFmcRouteMapConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcRouteMapPrerequisitesConfig = `
resource "fmc_security_zone" "test" {
  name           = "my_route_map_security_zone"
  interface_type = "ROUTED"
}

resource "fmc_standard_acl" "test" {
  name        = "my_route_map_std_acl"
  entries = [
    { action = "DENY", literals = [{ value = "10.1.1.0/24" }] }
  ]
}

resource "fmc_extended_acl" "test" {
  name        = "my_route_map_extended_acl"
  entries = [
    {
      log_level                    = "WARNING"
      logging                      = "PER_ACCESS_LIST_ENTRY"
      log_interval_seconds         = 120
      action                       = "DENY"
      source_network_literals      = [{ value = "2002::0/64", type = "Network" }]
      destination_network_literals = [{ value = "2003::1", type = "Host" }]
    }
  ]
}

resource "fmc_as_path" "test" {
  name = 261
  entries = [
    { action = "PERMIT", regular_expression = "107" },
    { action = "PERMIT", regular_expression = "108" },
  ]
}

resource "fmc_standard_community_list" "test" {
  name = "my_route_map_std_community_list"
  entries = [{
    action       = "PERMIT"
    communities  = "100 200 300"
    }
  ]
}

resource "fmc_policy_list" "test" {
  name   = "my_route_map_policy_list"
  action = "PERMIT"
  metric                  = 100
  tag                     = 200
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcRouteMapConfig() string {
	config := `resource "fmc_route_map" "test" {` + "\n"
	config += `	name = "my_route_map"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "PERMIT"` + "\n"
	config += `		match_security_zones = [{` + "\n"
	config += `			id = fmc_security_zone.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		match_interface_names = ["GigabitEthernet0/1"]` + "\n"
	config += `		match_ipv4_address_access_lists = [{` + "\n"
	config += `			id = fmc_standard_acl.test.id` + "\n"
	config += `			type = "StandardAccessList"` + "\n"
	config += `		}]` + "\n"
	config += `		match_ipv4_next_hop_access_lists = [{` + "\n"
	config += `			id = fmc_standard_acl.test.id` + "\n"
	config += `			type = "StandardAccessList"` + "\n"
	config += `		}]` + "\n"
	config += `		match_ipv4_route_source_access_lists = [{` + "\n"
	config += `			id = fmc_standard_acl.test.id` + "\n"
	config += `			type = "StandardAccessList"` + "\n"
	config += `		}]` + "\n"
	config += `		match_ipv6_address_extended_access_list_id = fmc_extended_acl.test.id` + "\n"
	config += `		match_ipv6_next_hop_extended_access_list_id = fmc_extended_acl.test.id` + "\n"
	config += `		match_ipv6_route_source_extended_access_list_id = fmc_extended_acl.test.id` + "\n"
	config += `		match_bgp_as_path_lists = [{` + "\n"
	config += `			id = fmc_as_path.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		match_bgp_community_lists = [{` + "\n"
	config += `			id = fmc_standard_community_list.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		match_bgp_policy_lists = [{` + "\n"
	config += `			id = fmc_policy_list.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		match_metric_route_values = [100]` + "\n"
	config += `		match_tag_values = [110]` + "\n"
	config += `		match_route_type_external_1 = true` + "\n"
	config += `		match_route_type_external_2 = true` + "\n"
	config += `		match_route_type_internal = true` + "\n"
	config += `		match_route_type_local = true` + "\n"
	config += `		match_route_type_nssa_external_1 = true` + "\n"
	config += `		match_route_type_nssa_external_2 = true` + "\n"
	config += `		set_metric_bandwidth = 1000000` + "\n"
	config += `		set_metric_type = "INTERNAL"` + "\n"
	config += `		set_bgp_as_path_prepend = [64512]` + "\n"
	config += `		set_bgp_as_path_prepend_last_as = 2` + "\n"
	config += `		set_bgp_as_path_convert_route_tag_into_as_path = true` + "\n"
	config += `		set_bgp_community_specific_community = 100` + "\n"
	config += `		set_bgp_community_add_to_existing_communities = true` + "\n"
	config += `		set_bgp_community_internet = true` + "\n"
	config += `		set_bgp_community_no_advertise = true` + "\n"
	config += `		set_bgp_community_no_export = true` + "\n"
	config += `		set_bgp_community_route_target = "100:100,200:200"` + "\n"
	config += `		set_bgp_community_add_to_existing_extended_communities = true` + "\n"
	config += `		set_bgp_automatic_tag = true` + "\n"
	config += `		set_bgp_local_preference = 100` + "\n"
	config += `		set_bgp_weight = 100` + "\n"
	config += `		set_bgp_origin = "LOCAL_IGP"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_route_map" "test" {
			id = fmc_route_map.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcRouteMapConfig() string {
	config := `resource "fmc_route_map" "test" {` + "\n"
	config += `	name = "my_route_map"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	entries = [{` + "\n"
	config += `		action = "PERMIT"` + "\n"
	config += `		match_security_zones = [{` + "\n"
	config += `			id = fmc_security_zone.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		match_interface_names = ["GigabitEthernet0/1"]` + "\n"
	config += `		match_ipv4_address_access_lists = [{` + "\n"
	config += `			id = fmc_standard_acl.test.id` + "\n"
	config += `			type = "StandardAccessList"` + "\n"
	config += `		}]` + "\n"
	config += `		match_ipv4_next_hop_access_lists = [{` + "\n"
	config += `			id = fmc_standard_acl.test.id` + "\n"
	config += `			type = "StandardAccessList"` + "\n"
	config += `		}]` + "\n"
	config += `		match_ipv4_route_source_access_lists = [{` + "\n"
	config += `			id = fmc_standard_acl.test.id` + "\n"
	config += `			type = "StandardAccessList"` + "\n"
	config += `		}]` + "\n"
	config += `		match_ipv6_address_extended_access_list_id = fmc_extended_acl.test.id` + "\n"
	config += `		match_ipv6_next_hop_extended_access_list_id = fmc_extended_acl.test.id` + "\n"
	config += `		match_ipv6_route_source_extended_access_list_id = fmc_extended_acl.test.id` + "\n"
	config += `		match_bgp_as_path_lists = [{` + "\n"
	config += `			id = fmc_as_path.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		match_bgp_community_lists = [{` + "\n"
	config += `			id = fmc_standard_community_list.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		match_bgp_policy_lists = [{` + "\n"
	config += `			id = fmc_policy_list.test.id` + "\n"
	config += `		}]` + "\n"
	config += `		match_metric_route_values = [100]` + "\n"
	config += `		match_tag_values = [110]` + "\n"
	config += `		match_route_type_external_1 = true` + "\n"
	config += `		match_route_type_external_2 = true` + "\n"
	config += `		match_route_type_internal = true` + "\n"
	config += `		match_route_type_local = true` + "\n"
	config += `		match_route_type_nssa_external_1 = true` + "\n"
	config += `		match_route_type_nssa_external_2 = true` + "\n"
	config += `		set_metric_bandwidth = 1000000` + "\n"
	config += `		set_metric_type = "INTERNAL"` + "\n"
	config += `		set_bgp_as_path_prepend = [64512]` + "\n"
	config += `		set_bgp_as_path_prepend_last_as = 2` + "\n"
	config += `		set_bgp_as_path_convert_route_tag_into_as_path = true` + "\n"
	config += `		set_bgp_community_specific_community = 100` + "\n"
	config += `		set_bgp_community_add_to_existing_communities = true` + "\n"
	config += `		set_bgp_community_internet = true` + "\n"
	config += `		set_bgp_community_no_advertise = true` + "\n"
	config += `		set_bgp_community_no_export = true` + "\n"
	config += `		set_bgp_community_route_target = "100:100,200:200"` + "\n"
	config += `		set_bgp_community_add_to_existing_extended_communities = true` + "\n"
	config += `		set_bgp_automatic_tag = true` + "\n"
	config += `		set_bgp_local_preference = 100` + "\n"
	config += `		set_bgp_weight = 100` + "\n"
	config += `		set_bgp_origin = "LOCAL_IGP"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_route_map" "test" {
			name = fmc_route_map.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
