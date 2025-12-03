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
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "process_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "router_id", "10.10.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "rfc_1583_compatible", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "log_adjacency_changes", "DEFAULT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "ignore_lsa_mospf", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "administrative_distance_inter_area", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "administrative_distance_intra_area", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "administrative_distance_external", "110"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "timer_lsa_group", "240"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "default_route_always_advertise", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "default_route_metric", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "default_route_metric_type", "TYPE_2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.type", "normal"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.default_cost", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.ranges.0.advertise", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ospf.test", "areas.0.inter_area_filters.0.filter_direction", "IN"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceOSPFPrerequisitesConfig + testAccFmcDeviceOSPFConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceOSPFPrerequisitesConfig + testAccFmcDeviceOSPFConfig_all(),
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

const testAccFmcDeviceOSPFPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

resource "fmc_hosts" "test" {
  items = {
    "ospf_host_1" = { ip = "10.20.3.4" },
    "ospf_host_2" = { ip = "10.20.3.5" },
  }
}

resource "fmc_ipv4_prefix_list" "test" {
  name = "ospf_ipv4_prefix_list"
  entries = [
    {
      action            = "PERMIT"
      ip_address        = "10.10.10.0/24"
      min_prefix_length = 25
      max_prefix_length = 30
    },
    {
      action            = "PERMIT"
      ip_address        = "15.10.10.0/24"
      min_prefix_length = 25
      max_prefix_length = 30
    },
  ]
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceOSPFConfig_minimum() string {
	config := `resource "fmc_device_ospf" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	process_id = 1` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceOSPFConfig_all() string {
	config := `resource "fmc_device_ospf" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	process_id = 1` + "\n"
	config += `	router_id = "10.10.10.1"` + "\n"
	config += `	rfc_1583_compatible = false` + "\n"
	config += `	log_adjacency_changes = "DEFAULT"` + "\n"
	config += `	ignore_lsa_mospf = false` + "\n"
	config += `	administrative_distance_inter_area = 110` + "\n"
	config += `	administrative_distance_intra_area = 110` + "\n"
	config += `	administrative_distance_external = 110` + "\n"
	config += `	timer_lsa_group = 240` + "\n"
	config += `	default_route_always_advertise = false` + "\n"
	config += `	default_route_metric = 1` + "\n"
	config += `	default_route_metric_type = "TYPE_2"` + "\n"
	config += `	areas = [{` + "\n"
	config += `		id = "1"` + "\n"
	config += `		type = "normal"` + "\n"
	config += `		networks = [{` + "\n"
	config += `			id = fmc_hosts.test.items.ospf_host_1.id` + "\n"
	config += `			name = "ospf_host_1"` + "\n"
	config += `		}]` + "\n"
	config += `		default_cost = 1` + "\n"
	config += `		ranges = [{` + "\n"
	config += `			network_object_id = fmc_hosts.test.items.ospf_host_2.id` + "\n"
	config += `			advertise = true` + "\n"
	config += `		}]` + "\n"
	config += `		inter_area_filters = [{` + "\n"
	config += `			prefix_list_id = fmc_ipv4_prefix_list.test.id` + "\n"
	config += `			prefix_list_name = fmc_ipv4_prefix_list.test.name` + "\n"
	config += `			filter_direction = "IN"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
