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

func TestAccDataSourceFmcSLAMonitors(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_sla_monitors.test", "items.my_sla_monitors.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_sla_monitors.test", "items.my_sla_monitors.type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sla_monitors.test", "items.my_sla_monitors.description", "My SLA Monitor object"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sla_monitors.test", "items.my_sla_monitors.sla_monitor_id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sla_monitors.test", "items.my_sla_monitors.timeout", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sla_monitors.test", "items.my_sla_monitors.frequency", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sla_monitors.test", "items.my_sla_monitors.threshold", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sla_monitors.test", "items.my_sla_monitors.data_size", "28"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sla_monitors.test", "items.my_sla_monitors.tos", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sla_monitors.test", "items.my_sla_monitors.number_of_packets", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_sla_monitors.test", "items.my_sla_monitors.monitor_address", "10.10.10.1"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcSLAMonitorsPrerequisitesConfig + testAccDataSourceFmcSLAMonitorsConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcSLAMonitorsPrerequisitesConfig = `
resource "fmc_interface_group" "test" {
  name           = "sla_monitors_interface_group"
  interface_mode = "ROUTED"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcSLAMonitorsConfig() string {
	config := `resource "fmc_sla_monitors" "test" {` + "\n"
	config += `	items = { "my_sla_monitors" = {` + "\n"
	config += `		description = "My SLA Monitor object"` + "\n"
	config += `		sla_monitor_id = 10` + "\n"
	config += `		timeout = 5000` + "\n"
	config += `		frequency = 60` + "\n"
	config += `		threshold = 5000` + "\n"
	config += `		data_size = 28` + "\n"
	config += `		tos = 20` + "\n"
	config += `		number_of_packets = 1` + "\n"
	config += `		monitor_address = "10.10.10.1"` + "\n"
	config += `		selected_interfaces = [{` + "\n"
	config += `			id = fmc_interface_group.test.id` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_sla_monitors" "test" {
			depends_on = [fmc_sla_monitors.test]
			items = {
				"my_sla_monitors" = {
				}
			}
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
