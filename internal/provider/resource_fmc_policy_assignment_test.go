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

func TestAccFmcPolicyAssignment(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_policy_assignment.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_policy_assignment.test", "policy_name"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_policy_assignment.test", "policy_type", "FTDNatPolicy"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcPolicyAssignmentPrerequisitesConfig + testAccFmcPolicyAssignmentConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_policy_assignment.test",
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

const testAccFmcPolicyAssignmentPrerequisitesConfig = `
resource "fmc_ftd_nat_policy" "example" {
  name = "policy_assignment_nat_policy"
}

variable "device_id" { default = null } // tests will set $TF_VAR_device_id

data "fmc_device" "test" {
  id = var.device_id
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcPolicyAssignmentConfig_all() string {
	config := `resource "fmc_policy_assignment" "test" {` + "\n"
	config += `	policy_id = fmc_ftd_nat_policy.example.id` + "\n"
	config += `	policy_type = "FTDNatPolicy"` + "\n"
	config += `	targets = [{` + "\n"
	config += `		id = data.fmc_device.test.id` + "\n"
	config += `		type = data.fmc_device.test.type` + "\n"
	config += `		name = data.fmc_device.test.name` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
