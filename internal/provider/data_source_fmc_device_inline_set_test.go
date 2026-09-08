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

func TestAccDataSourceFmcDeviceInlineSet(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" || os.Getenv("TF_VAR_interface_2_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name and TF_VAR_interface_2_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_inline_set.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_inline_set.test", "name", "my_inline_set"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_inline_set.test", "mtu", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_inline_set.test", "fail_safe", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_inline_set.test", "tap_mode", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_inline_set.test", "propagate_link_state", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_inline_set.test", "strict_tcp_enforcement", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_inline_set.test", "snort_fail_open_busy", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_inline_set.test", "snort_fail_open_down", "true"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceInlineSetPrerequisitesConfig + testAccDataSourceFmcDeviceInlineSetConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceInlineSetPrerequisitesConfig + testAccNamedDataSourceFmcDeviceInlineSetConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceInlineSetPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" { default = null } // tests will set $TF_VAR_interface_name
variable "interface_2_name" { default = null } // tests will set $TF_VAR_interface_2_name

resource "fmc_device_physical_interface" "test" {
  device_id    = var.device_id
  name         = var.interface_name
  logical_name = "inline_member_1"
}

resource "fmc_device_physical_interface" "test_2" {
  device_id    = var.device_id
  name         = var.interface_2_name
  logical_name = "inline_member_2"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceInlineSetConfig() string {
	config := `resource "fmc_device_inline_set" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	name = "my_inline_set"` + "\n"
	config += `	mtu = 1500` + "\n"
	config += `	fail_safe = true` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `		first_interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		first_interface_name = fmc_device_physical_interface.test.name` + "\n"
	config += `		first_interface_type = fmc_device_physical_interface.test.type` + "\n"
	config += `		second_interface_id = fmc_device_physical_interface.test_2.id` + "\n"
	config += `		second_interface_name = fmc_device_physical_interface.test_2.name` + "\n"
	config += `		second_interface_type = fmc_device_physical_interface.test_2.type` + "\n"
	config += `	}]` + "\n"
	config += `	tap_mode = false` + "\n"
	config += `	propagate_link_state = true` + "\n"
	config += `	strict_tcp_enforcement = true` + "\n"
	config += `	snort_fail_open_busy = false` + "\n"
	config += `	snort_fail_open_down = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_inline_set" "test" {
			id = fmc_device_inline_set.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceInlineSetConfig() string {
	config := `resource "fmc_device_inline_set" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	name = "my_inline_set"` + "\n"
	config += `	mtu = 1500` + "\n"
	config += `	fail_safe = true` + "\n"
	config += `	interface_pairs = [{` + "\n"
	config += `		first_interface_id = fmc_device_physical_interface.test.id` + "\n"
	config += `		first_interface_name = fmc_device_physical_interface.test.name` + "\n"
	config += `		first_interface_type = fmc_device_physical_interface.test.type` + "\n"
	config += `		second_interface_id = fmc_device_physical_interface.test_2.id` + "\n"
	config += `		second_interface_name = fmc_device_physical_interface.test_2.name` + "\n"
	config += `		second_interface_type = fmc_device_physical_interface.test_2.type` + "\n"
	config += `	}]` + "\n"
	config += `	tap_mode = false` + "\n"
	config += `	propagate_link_state = true` + "\n"
	config += `	strict_tcp_enforcement = true` + "\n"
	config += `	snort_fail_open_busy = false` + "\n"
	config += `	snort_fail_open_down = true` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_inline_set" "test" {
			device_id = var.device_id
			name = fmc_device_inline_set.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
