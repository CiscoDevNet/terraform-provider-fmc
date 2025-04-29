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

func TestAccFmcDeviceLoopbackInterface(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_loopback_interface.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device_loopback_interface.test", "name"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_loopback_interface.test", "logical_name", "my_loopback_1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_loopback_interface.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_loopback_interface.test", "loopback_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_loopback_interface.test", "description", "my VTI interface"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_loopback_interface.test", "ipv4_static_address", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_loopback_interface.test", "ipv4_static_netmask", "24"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_loopback_interface.test", "ipv6_addresses.0.address", "2004::10"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_device_loopback_interface.test", "ipv6_addresses.0.prefix", "64"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcDeviceLoopbackInterfacePrerequisitesConfig + testAccFmcDeviceLoopbackInterfaceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcDeviceLoopbackInterfacePrerequisitesConfig + testAccFmcDeviceLoopbackInterfaceConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
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

const testAccFmcDeviceLoopbackInterfacePrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceLoopbackInterfaceConfig_minimum() string {
	config := `resource "fmc_device_loopback_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	loopback_id = 1` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceLoopbackInterfaceConfig_all() string {
	config := `resource "fmc_device_loopback_interface" "test" {` + "\n"
	config += `	device_id = var.device_id` + "\n"
	config += `	logical_name = "my_loopback_1"` + "\n"
	config += `	enabled = true` + "\n"
	config += `	loopback_id = 1` + "\n"
	config += `	description = "my VTI interface"` + "\n"
	config += `	ipv4_static_address = "10.1.1.1"` + "\n"
	config += `	ipv4_static_netmask = "24"` + "\n"
	config += `	ipv6_addresses = [{` + "\n"
	config += `		address = "2004::10"` + "\n"
	config += `		prefix = "64"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
