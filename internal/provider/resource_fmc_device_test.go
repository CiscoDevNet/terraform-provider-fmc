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

	"github.com/CiscoDevNet/terraform-provider-fmc/internal/ftd"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

func TestAccFmcDevice(t *testing.T) {
	ftdAddr := os.Getenv("FTD_ADDR")
	if ftdAddr == "" {
		t.Skip("skipping test, set environment variable FTD_ADDR")
	}

	if os.Getenv("TF_VAR_nat_id") == "" {
		// Prepare Terraform's var.ftd_addr
		os.Setenv("TF_VAR_ftd_addr", ftdAddr)
	}
	// Prepare Terraform's var.registration_key
	os.Setenv("TF_VAR_device_registration_key", ftd.MustRandomizeKey())

	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_device.test", "id"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			PreConfig: ftd.MustInitFromEnv,
			Config:    testAccFmcDevicePrerequisitesConfig + testAccFmcDeviceConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		PreConfig: ftd.MustInitFromEnv,
		Config:    testAccFmcDevicePrerequisitesConfig + testAccFmcDeviceConfig_all(),
		Check:     resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_device.test",
		ImportState:  true,
	})

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccFmcDevicePrerequisitesConfig = `
variable "ftd_addr" { default = null } // tests will set $TF_VAR_ftd_addr
variable "device_registration_key" {} // tests will set $TF_VAR_device_registration_key
variable "licenses" { default = "BASE" } // tests will set $TF_VAR_licenses

resource "fmc_access_control_policy" "test" {
  name = "fmc_device_access_control_policy"
  default_action = "BLOCK"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcDeviceConfig_minimum() string {
	config := `resource "fmc_device" "test" {` + "\n"
	config += `	name = "my_device"` + "\n"
	config += `	host = var.ftd_addr` + "\n"
	config += `	licenses = [var.licenses]` + "\n"
	config += `	registration_key = var.device_registration_key` + "\n"
	config += `	access_control_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcDeviceConfig_all() string {
	config := `resource "fmc_device" "test" {` + "\n"
	config += `	name = "my_device"` + "\n"
	config += `	host = var.ftd_addr` + "\n"
	config += `	licenses = [var.licenses]` + "\n"
	config += `	registration_key = var.device_registration_key` + "\n"
	config += `	performance_tier = "FTDv5"` + "\n"
	config += `	snort_engine = "SNORT3"` + "\n"
	config += `	access_control_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
