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

func TestAccDataSourceFmcDeviceClusterHealthMonitor(t *testing.T) {
	if os.Getenv("TF_VAR_cluster_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_cluster_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_cluster_health_monitor.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "health_check", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "hold_time", "3.0"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "debounce_time", "9000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "data_interface_auto_rejoin_attempts", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "data_interface_auto_rejoin_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "data_interface_auto_rejoin_interval_variation", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "cluster_interface_auto_rejoin_attempts", "-1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "cluster_interface_auto_rejoin_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "cluster_interface_auto_rejoin_interval_variation", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "system_auto_rejoin_attempts", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "system_auto_rejoin_interval", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "system_auto_rejoin_interval_variation", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster_health_monitor.test", "service_application_monitoring", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceClusterHealthMonitorPrerequisitesConfig + testAccDataSourceFmcDeviceClusterHealthMonitorConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceClusterHealthMonitorPrerequisitesConfig = `
variable "cluster_id" { default = null } // tests will set $TF_VAR_cluster_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceClusterHealthMonitorConfig() string {
	config := `resource "fmc_device_cluster_health_monitor" "test" {` + "\n"
	config += `	cluster_id = var.cluster_id` + "\n"
	config += `	health_check = true` + "\n"
	config += `	hold_time = 3.0` + "\n"
	config += `	debounce_time = 9000` + "\n"
	config += `	data_interface_auto_rejoin_attempts = 3` + "\n"
	config += `	data_interface_auto_rejoin_interval = 5` + "\n"
	config += `	data_interface_auto_rejoin_interval_variation = 2` + "\n"
	config += `	cluster_interface_auto_rejoin_attempts = -1` + "\n"
	config += `	cluster_interface_auto_rejoin_interval = 5` + "\n"
	config += `	cluster_interface_auto_rejoin_interval_variation = 1` + "\n"
	config += `	system_auto_rejoin_attempts = 3` + "\n"
	config += `	system_auto_rejoin_interval = 5` + "\n"
	config += `	system_auto_rejoin_interval_variation = 2` + "\n"
	config += `	unmonitored_interfaces = ["GigabitEthernet0/1"]` + "\n"
	config += `	service_application_monitoring = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_cluster_health_monitor" "test" {
			id = fmc_device_cluster_health_monitor.test.id
			cluster_id = var.cluster_id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
