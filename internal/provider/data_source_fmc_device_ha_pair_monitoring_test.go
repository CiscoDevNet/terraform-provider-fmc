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

func TestAccDataSourceFmcDeviceHAPairMonitoring(t *testing.T) {
	if os.Getenv("TF_VAR_device_ha_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_ha_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_ha_pair_monitoring.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair_monitoring.test", "logical_name", "outside"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair_monitoring.test", "monitor_interface", "true"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_ha_pair_monitoring.test", "ipv4_active_address"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair_monitoring.test", "ipv4_standby_address", "10.1.1.2"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_ha_pair_monitoring.test", "ipv4_netmask"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair_monitoring.test", "ipv6_addresses.0.active_address", "2006::1/30"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ha_pair_monitoring.test", "ipv6_addresses.0.standby_address", "2006::2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceHAPairMonitoringPrerequisitesConfig + testAccDataSourceFmcDeviceHAPairMonitoringConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceHAPairMonitoringPrerequisitesConfig + testAccNamedDataSourceFmcDeviceHAPairMonitoringConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceHAPairMonitoringPrerequisitesConfig = `
variable "device_ha_id" { default = null } // tests will set $TF_VAR_device_ha_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceHAPairMonitoringConfig() string {
	config := `resource "fmc_device_ha_pair_monitoring" "test" {` + "\n"
	config += `	ha_pair_id = var.device_ha_id` + "\n"
	config += `	logical_name = "outside"` + "\n"
	config += `	monitor_interface = true` + "\n"
	config += `	ipv4_standby_address = "10.1.1.2"` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		active_address = "2006::1/30"` + "\n"
	config += `		standby_address = "2006::2"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ha_pair_monitoring" "test" {
			id = fmc_device_ha_pair_monitoring.test.id
			ha_pair_id = var.device_ha_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceHAPairMonitoringConfig() string {
	config := `resource "fmc_device_ha_pair_monitoring" "test" {` + "\n"
	config += `	ha_pair_id = var.device_ha_id` + "\n"
	config += `	logical_name = "outside"` + "\n"
	config += `	monitor_interface = true` + "\n"
	config += `	ipv4_standby_address = "10.1.1.2"` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		active_address = "2006::1/30"` + "\n"
	config += `		standby_address = "2006::2"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ha_pair_monitoring" "test" {
			ha_pair_id = var.device_ha_id
			logical_name = fmc_device_ha_pair_monitoring.test.logical_name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
