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

func TestAccFmcSLAMonitor(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "name", "my_sla_monitor"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_sla_monitor.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "description", "My SLA Monitor object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "sla_monitor_id", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "timeout", "5000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "frequency", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "threshold", "1000"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "data_size", "28"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "tos", "20"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "number_of_packets", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_sla_monitor.test", "monitor_address", "10.10.10.1"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcSLAMonitorPrerequisitesConfig + testAccFmcSLAMonitorConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcSLAMonitorPrerequisitesConfig + testAccFmcSLAMonitorConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_sla_monitor.test",
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

const testAccFmcSLAMonitorPrerequisitesConfig = `
resource "fmc_interface_group" "test" {
  name           = "sla_monitor_interface_group"
  interface_mode = "ROUTED"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcSLAMonitorConfig_minimum() string {
	config := `resource "fmc_sla_monitor" "test" {` + "\n"
	config += `	name = "my_sla_monitor"` + "\n"
	config += `	sla_monitor_id = 10` + "\n"
	config += `	monitor_address = "10.10.10.1"` + "\n"
	config += `	selected_interfaces = [{` + "\n"
	config += `		id = fmc_interface_group.test.id` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcSLAMonitorConfig_all() string {
	config := `resource "fmc_sla_monitor" "test" {` + "\n"
	config += `	name = "my_sla_monitor"` + "\n"
	config += `	description = "My SLA Monitor object"` + "\n"
	config += `	sla_monitor_id = 10` + "\n"
	config += `	timeout = 5000` + "\n"
	config += `	frequency = 60` + "\n"
	config += `	threshold = 1000` + "\n"
	config += `	data_size = 28` + "\n"
	config += `	tos = 20` + "\n"
	config += `	number_of_packets = 1` + "\n"
	config += `	monitor_address = "10.10.10.1"` + "\n"
	config += `	selected_interfaces = [{` + "\n"
	config += `		id = fmc_interface_group.test.id` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
