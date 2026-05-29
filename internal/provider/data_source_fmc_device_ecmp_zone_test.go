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

func TestAccDataSourceFmcDeviceECMPZone(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" || os.Getenv("TF_VAR_interface_name") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id and TF_VAR_interface_name")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_device_ecmp_zone.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_device_ecmp_zone.test", "name", "my_ecmp_zone"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcDeviceECMPZonePrerequisitesConfig + testAccDataSourceFmcDeviceECMPZoneConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcDeviceECMPZonePrerequisitesConfig + testAccNamedDataSourceFmcDeviceECMPZoneConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcDeviceECMPZonePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
variable "interface_name" { default = null } // tests will set $TF_VAR_interface_name

resource "fmc_device_physical_interface" "test" {
  device_id = var.device_id
  name = var.interface_name
  logical_name = "outside"
  mode = "NONE"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcDeviceECMPZoneConfig() string {
	config := `resource "fmc_device_ecmp_zone" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	name = "my_ecmp_zone"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		id = fmc_device_physical_interface.test.id` + "\n"
	config += `		name = fmc_device_physical_interface.test.name` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ecmp_zone" "test" {
			id = fmc_device_ecmp_zone.test.id
			device_id = var.device_id
		}
	`
	return config
}

func testAccNamedDataSourceFmcDeviceECMPZoneConfig() string {
	config := `resource "fmc_device_ecmp_zone" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	name = "my_ecmp_zone"` + "\n"
	config += `	interfaces = [{` + "\n"
	config += `		id = fmc_device_physical_interface.test.id` + "\n"
	config += `		name = fmc_device_physical_interface.test.name` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_device_ecmp_zone" "test" {
			device_id = var.device_id
			name = fmc_device_ecmp_zone.test.name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
