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

func TestAccDataSourceFmcDeviceCluster(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_device_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_device_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster.test", "name", "my_device_cluster"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_cluster.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster.test", "control_node_vni_prefix", "10.10.3.0/27"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster.test", "control_node_ccl_prefix", "10.10.4.0/27"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster.test", "control_node_ccl_ipv4_address", "10.10.4.1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster.test", "control_node_priority", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster.test", "data_nodes.0.device_id", "76d24097-41c4-4558-a4d0-a8c07ac08470"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster.test", "data_nodes.0.ccl_ipv4_address", "10.10.4.2"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_cluster.test", "data_nodes.0.priority", "2"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceClusterPrerequisitesConfig + testAccDataSourceFmcDeviceClusterConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceClusterPrerequisitesConfig + testAccNamedDataSourceFmcDeviceClusterConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceClusterPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "device_interface_name" { default = null } // tests will set $TF_VAR_device_interface_name

data "fmc_device_physical_interface" "test" {
  device_id   = var.device_id
  id          = var.device_interface_name
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceClusterConfig() string {
	config := `resource "fmc_device_cluster" "test" {` + "\n"
	config += `	name = "my_device_cluster"` + "\n"
	config += `	cluster_key = "cisco123"` + "\n"
	config += `	control_node_vni_prefix = "10.10.3.0/27"` + "\n"
	config += `	control_node_ccl_prefix = "10.10.4.0/27"` + "\n"
	config += `	control_node_interface_id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `	control_node_interface_name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	control_node_interface_type = data.fmc_device_physical_interface.test.type` + "\n"
	config += `	control_node_device_id = var.device_id` + "\n"
	config += `	control_node_ccl_ipv4_address = "10.10.4.1"` + "\n"
	config += `	control_node_priority = 1` + "\n"
	config += `	data_nodes = [{` + "\n"
	config += `		device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `		ccl_ipv4_address = "10.10.4.2"` + "\n"
	config += `		priority = 2` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_cluster" "test" {
			id = fmc_device_cluster.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceClusterConfig() string {
	config := `resource "fmc_device_cluster" "test" {` + "\n"
	config += `	name = "my_device_cluster"` + "\n"
	config += `	cluster_key = "cisco123"` + "\n"
	config += `	control_node_vni_prefix = "10.10.3.0/27"` + "\n"
	config += `	control_node_ccl_prefix = "10.10.4.0/27"` + "\n"
	config += `	control_node_interface_id = data.fmc_device_physical_interface.test.id` + "\n"
	config += `	control_node_interface_name = data.fmc_device_physical_interface.test.name` + "\n"
	config += `	control_node_interface_type = data.fmc_device_physical_interface.test.type` + "\n"
	config += `	control_node_device_id = var.device_id` + "\n"
	config += `	control_node_ccl_ipv4_address = "10.10.4.1"` + "\n"
	config += `	control_node_priority = 1` + "\n"
	config += `	data_nodes = [{` + "\n"
	config += `		device_id = "76d24097-41c4-4558-a4d0-a8c07ac08470"` + "\n"
	config += `		ccl_ipv4_address = "10.10.4.2"` + "\n"
	config += `		priority = 2` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_cluster" "test" {
			name = fmc_device_cluster.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
