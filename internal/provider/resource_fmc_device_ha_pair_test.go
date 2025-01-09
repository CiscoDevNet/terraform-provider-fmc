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

func TestAccFmcDeviceHAPair(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_device_2_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_device_2_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "name", "FTD_HA"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_ha_pair.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "ha_link_interface_name", "GigabitEthernet0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "ha_link_logical_name", "LAN-INTERFACE"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "ha_link_use_ipv6", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "ha_link_primary_ip", "1.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "ha_link_secondary_ip", "1.1.1.2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "ha_link_netmask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "state_link_interface_name", "GigabitEthernet0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "state_link_logical_name", "Stateful-INTERFACE"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "state_link_use_ipv6", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "state_link_primary_ip", "10.10.10.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "state_link_secondary_ip", "10.10.10.2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "state_link_netmask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "encryption_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "failed_interfaces_limit", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "peer_poll_time", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "peer_poll_time_unit", "SEC"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "peer_hold_time", "15"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "peer_hold_time_unit", "SEC"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "interface_poll_time", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "interface_poll_time_unit", "SEC"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_ha_pair.test", "interface_hold_time", "25"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceHAPairPrerequisitesConfig + testAccFmcDeviceHAPairConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceHAPairPrerequisitesConfig + testAccFmcDeviceHAPairConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_device_ha_pair.test",
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

const testAccFmcDeviceHAPairPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "device_2_id" { default = null } // tests will set $TF_VAR_device_2_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceHAPairConfig_minimum() string {
	config := `resource "fmc_device_ha_pair" "test" {` + "\n"
	config += `	name = "FTD_HA"` + "\n"
	config += `	primary_device_id = var.device_id` + "\n"
	config += `	secondary_device_id = var.device_2_id` + "\n"
	config += `	ha_link_interface_id = "96d24097-41c4-4332-a4d0-a8c07ac08482"` + "\n"
	config += `	ha_link_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	ha_link_interface_type = ""` + "\n"
	config += `	ha_link_logical_name = "LAN-INTERFACE"` + "\n"
	config += `	ha_link_primary_ip = "1.1.1.1"` + "\n"
	config += `	ha_link_secondary_ip = "1.1.1.2"` + "\n"
	config += `	ha_link_netmask = "255.255.255.0"` + "\n"
	config += `	state_link_use_same_as_ha = false` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceHAPairConfig_all() string {
	config := `resource "fmc_device_ha_pair" "test" {` + "\n"
	config += `	name = "FTD_HA"` + "\n"
	config += `	primary_device_id = var.device_id` + "\n"
	config += `	secondary_device_id = var.device_2_id` + "\n"
	config += `	ha_link_interface_id = "96d24097-41c4-4332-a4d0-a8c07ac08482"` + "\n"
	config += `	ha_link_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	ha_link_interface_type = ""` + "\n"
	config += `	ha_link_logical_name = "LAN-INTERFACE"` + "\n"
	config += `	ha_link_use_ipv6 = false` + "\n"
	config += `	ha_link_primary_ip = "1.1.1.1"` + "\n"
	config += `	ha_link_secondary_ip = "1.1.1.2"` + "\n"
	config += `	ha_link_netmask = "255.255.255.0"` + "\n"
	config += `	state_link_use_same_as_ha = false` + "\n"
	config += `	state_link_interface_id = "76d24097-hj7r-7786-a4d0-a8c07ac08470"` + "\n"
	config += `	state_link_interface_name = "GigabitEthernet0/0"` + "\n"
	config += `	state_link_interface_type = "PhysicalInterface"` + "\n"
	config += `	state_link_logical_name = "Stateful-INTERFACE"` + "\n"
	config += `	state_link_use_ipv6 = false` + "\n"
	config += `	state_link_primary_ip = "10.10.10.1"` + "\n"
	config += `	state_link_secondary_ip = "10.10.10.2"` + "\n"
	config += `	state_link_netmask = "255.255.255.0"` + "\n"
	config += `	encryption_enabled = true` + "\n"
	config += `	encryption_key_generation_scheme = "AUTO"` + "\n"
	config += `	failed_interfaces_limit = 1` + "\n"
	config += `	peer_poll_time = 1` + "\n"
	config += `	peer_poll_time_unit = "SEC"` + "\n"
	config += `	peer_hold_time = 15` + "\n"
	config += `	peer_hold_time_unit = "SEC"` + "\n"
	config += `	interface_poll_time = 5` + "\n"
	config += `	interface_poll_time_unit = "SEC"` + "\n"
	config += `	interface_hold_time = 25` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
