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

func TestAccFmcNetworkOverrides(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_network_overrides.test", "overrides.0.target_type", "Device"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_network_overrides.test", "overrides.0.description", "My Network object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_network_overrides.test", "overrides.0.prefix", "10.1.2.0/24"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcNetworkOverridesPrerequisitesConfig + testAccFmcNetworkOverridesConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcNetworkOverridesPrerequisitesConfig + testAccFmcNetworkOverridesConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_network_overrides.test",
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

const testAccFmcNetworkOverridesPrerequisitesConfig = `
variable "device_id" { default = null } // tests will set $TF_VAR_device_id

resource "fmc_network" "test" {
  name        = "my_network_override"
  prefix      = "10.0.2.0/24"
  overridable = true
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcNetworkOverridesConfig_minimum() string {
	config := `resource "fmc_network_overrides" "test" {` + "\n"
	config += `	parent_name = fmc_network.test.name` + "\n"
	config += `	parent_id = fmc_network.test.id` + "\n"
	config += `	overrides = [{` + "\n"
	config += `		target_id = var.device_id` + "\n"
	config += `		target_type = "Device"` + "\n"
	config += `		prefix = "10.1.2.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcNetworkOverridesConfig_all() string {
	config := `resource "fmc_network_overrides" "test" {` + "\n"
	config += `	parent_name = fmc_network.test.name` + "\n"
	config += `	parent_id = fmc_network.test.id` + "\n"
	config += `	overrides = [{` + "\n"
	config += `		target_id = var.device_id` + "\n"
	config += `		target_type = "Device"` + "\n"
	config += `		description = "My Network object"` + "\n"
	config += `		prefix = "10.1.2.0/24"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
