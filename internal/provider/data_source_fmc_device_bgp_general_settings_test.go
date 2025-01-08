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

func TestAccDataSourceFmcDeviceBGPGeneralSettings(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_bgp_general_settings.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "as_number", "65535"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "router_id", "AUTOMATIC"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "scanning_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "as_number_in_path_attribute", "50"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "log_neighbor_changes", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "tcp_path_mtu_discovery", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "reset_session_upon_failover", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "enforce_first_peer_as", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "use_dot_notation", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "aggregate_timer", "30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "default_local_preference", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "compare_med_from_different_neighbors", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "compare_router_id_in_path", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "pick_best_med", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "missing_med_as_best", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "keepalive_interval", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "hold_time", "180"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "min_hold_time", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_bgp_general_settings.test", "next_hop_delay_interval", "5"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceBGPGeneralSettingsPrerequisitesConfig + testAccDataSourceFmcDeviceBGPGeneralSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceBGPGeneralSettingsPrerequisitesConfig + testAccNamedDataSourceFmcDeviceBGPGeneralSettingsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceBGPGeneralSettingsPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceBGPGeneralSettingsConfig() string {
	config := `resource "fmc_device_bgp_general_settings" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	as_number = "65535"` + "\n"
	config += `	router_id = "AUTOMATIC"` + "\n"
	config += `	scanning_interval = 60` + "\n"
	config += `	as_number_in_path_attribute = 50` + "\n"
	config += `	log_neighbor_changes = false` + "\n"
	config += `	tcp_path_mtu_discovery = true` + "\n"
	config += `	reset_session_upon_failover = true` + "\n"
	config += `	enforce_first_peer_as = true` + "\n"
	config += `	use_dot_notation = false` + "\n"
	config += `	aggregate_timer = 30` + "\n"
	config += `	default_local_preference = 100` + "\n"
	config += `	compare_med_from_different_neighbors = true` + "\n"
	config += `	compare_router_id_in_path = true` + "\n"
	config += `	pick_best_med = true` + "\n"
	config += `	missing_med_as_best = false` + "\n"
	config += `	keepalive_interval = 60` + "\n"
	config += `	hold_time = 180` + "\n"
	config += `	min_hold_time = 0` + "\n"
	config += `	next_hop_delay_interval = 5` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bgp_general_settings" "test" {
			id = fmc_device_bgp_general_settings.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceBGPGeneralSettingsConfig() string {
	config := `resource "fmc_device_bgp_general_settings" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	as_number = "65535"` + "\n"
	config += `	router_id = "AUTOMATIC"` + "\n"
	config += `	scanning_interval = 60` + "\n"
	config += `	as_number_in_path_attribute = 50` + "\n"
	config += `	log_neighbor_changes = false` + "\n"
	config += `	tcp_path_mtu_discovery = true` + "\n"
	config += `	reset_session_upon_failover = true` + "\n"
	config += `	enforce_first_peer_as = true` + "\n"
	config += `	use_dot_notation = false` + "\n"
	config += `	aggregate_timer = 30` + "\n"
	config += `	default_local_preference = 100` + "\n"
	config += `	compare_med_from_different_neighbors = true` + "\n"
	config += `	compare_router_id_in_path = true` + "\n"
	config += `	pick_best_med = true` + "\n"
	config += `	missing_med_as_best = false` + "\n"
	config += `	keepalive_interval = 60` + "\n"
	config += `	hold_time = 180` + "\n"
	config += `	min_hold_time = 0` + "\n"
	config += `	next_hop_delay_interval = 5` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_bgp_general_settings" "test" {
			device_id = var.device_id
			as_number = fmc_device_bgp_general_settings.test.as_number
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
