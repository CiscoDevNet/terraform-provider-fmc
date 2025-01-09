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

func TestAccFmcDeviceCluster(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_device_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_device_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_cluster.test", "name", "MyDeviceClusterName1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_cluster.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_cluster.test", "control_node_vni_network", "10.10.3.0/27"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_cluster.test", "control_node_ccl_network", "10.10.4.0/27"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_cluster.test", "control_node_ccl_ipv4_address", "10.10.4.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_cluster.test", "control_node_priority", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_cluster.test", "data_devices.0.data_node_device_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_cluster.test", "data_devices.0.data_node_ccl_ipv4_address", "10.10.4.2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_cluster.test", "data_devices.0.data_node_priority", "2"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceClusterPrerequisitesConfig + testAccFmcDeviceClusterConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceClusterPrerequisitesConfig + testAccFmcDeviceClusterConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_device_cluster.test",
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

const testAccFmcDeviceClusterPrerequisitesConfig = `
data "fmc_device_physical_interface" "test" {
  device_id   = var.device_id
  id          = var.device_interface_name
}
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "device_interface_name" { default = null } // tests will set $TF_VAR_device_interface_name
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceClusterConfig_minimum() string {
	config := `resource "fmc_device_cluster" "test" {` + "\n"
	config += `	name = "MyDeviceClusterName1"` + "\n"
	config += `	cluster_key = "cisco123"` + "\n"
	config += `	control_node_device_id = var.device_id` + "\n"
	config += `	control_node_ccl_network = "10.10.4.0/27"` + "\n"
	config += `	control_node_interface_id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `	control_node_interface_name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	control_node_interface_type = data.fmc_device_physical_interface.test.type` + "\n"
	config += `	control_node_ccl_ipv4_address = "10.10.4.1"` + "\n"
	config += `	control_node_priority = 1` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceClusterConfig_all() string {
	config := `resource "fmc_device_cluster" "test" {` + "\n"
	config += `	name = "MyDeviceClusterName1"` + "\n"
	config += `	cluster_key = "cisco123"` + "\n"
	config += `	control_node_device_id = var.device_id` + "\n"
	config += `	control_node_vni_network = "10.10.3.0/27"` + "\n"
	config += `	control_node_ccl_network = "10.10.4.0/27"` + "\n"
	config += `	control_node_interface_id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `	control_node_interface_name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	control_node_interface_type = data.fmc_device_physical_interface.test.type` + "\n"
	config += `	control_node_ccl_ipv4_address = "10.10.4.1"` + "\n"
	config += `	control_node_priority = 1` + "\n"
	config += `	data_devices = [{` + "\n"
	config += `		data_node_device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `		data_node_ccl_ipv4_address = "10.10.4.2"` + "\n"
	config += `		data_node_priority = 2` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
