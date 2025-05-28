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

func TestAccFmcRouteMap(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "name", "my_route_map"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_route_map.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "overridable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.sequence_number", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.action", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.security_zones.0.id", "0050568A-4E02-1ed3-0000-004294969198"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv4_access_list_addresses.0.id", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv4_access_list_addresses.0.type", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv4_access_list_next_hops.0.id", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv4_access_list_next_hops.0.type", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv4_access_list_route_sources.0.id", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv4_access_list_route_sources.0.type", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv6_access_list_addresses.0.id", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv6_access_list_addresses.0.type", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv6_access_list_next_hops.0.id", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv6_access_list_next_hops.0.type", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv6_access_list_route_sources.0.id", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.ipv6_access_list_route_sources.0.type", "0050568A-4E02-1ed3-0000-004294969199"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.metric_route_values.0", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.tag_values.0", ""))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.route_type_external1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.route_type_external2", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.route_type_internal", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.route_type_local", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.route_type_n_s_s_a_external1", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_route_map.test", "entries.0.route_type_n_s_s_a_external2", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcRouteMapConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcRouteMapConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_route_map.test",
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

func testAccFmcRouteMapConfig_minimum() string {
	config := `resource "fmc_route_map" "test" {` + "\n"
	config += `	name = "my_route_map"` + "\n"
	config += `	entries = [{` + "\n"
	config += `		sequence_number = 10` + "\n"
	config += `		action = ""` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcRouteMapConfig_all() string {
	config := `resource "fmc_route_map" "test" {` + "\n"
	config += `	name = "my_route_map"` + "\n"
	config += `	overridable = true` + "\n"
	config += `	entries = [{` + "\n"
	config += `		sequence_number = 10` + "\n"
	config += `		action = ""` + "\n"
	config += `		security_zones = [{` + "\n"
	config += `			id = "0050568A-4E02-1ed3-0000-004294969198"` + "\n"
	config += `		}]` + "\n"
	config += `		ipv4_access_list_addresses = [{` + "\n"
	config += `			id = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `			type = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `		}]` + "\n"
	config += `		ipv4_access_list_next_hops = [{` + "\n"
	config += `			id = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `			type = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `		}]` + "\n"
	config += `		ipv4_access_list_route_sources = [{` + "\n"
	config += `			id = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `			type = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `		}]` + "\n"
	config += `		ipv6_access_list_addresses = [{` + "\n"
	config += `			id = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `			type = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `		}]` + "\n"
	config += `		ipv6_access_list_next_hops = [{` + "\n"
	config += `			id = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `			type = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `		}]` + "\n"
	config += `		ipv6_access_list_route_sources = [{` + "\n"
	config += `			id = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `			type = "0050568A-4E02-1ed3-0000-004294969199"` + "\n"
	config += `		}]` + "\n"
	config += `		metric_route_values = []` + "\n"
	config += `		tag_values = []` + "\n"
	config += `		route_type_external1 = true` + "\n"
	config += `		route_type_external2 = true` + "\n"
	config += `		route_type_internal = true` + "\n"
	config += `		route_type_local = true` + "\n"
	config += `		route_type_n_s_s_a_external1 = true` + "\n"
	config += `		route_type_n_s_s_a_external2 = true` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
