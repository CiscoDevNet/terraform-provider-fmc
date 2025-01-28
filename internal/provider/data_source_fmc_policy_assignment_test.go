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

func TestAccDataSourceFmcPolicyAssignment(t *testing.T) {
	if os.Getenv("TF_VAR_device_id") == "" {
		t.Skip("skipping test, set environment variable TF_VAR_device_id")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_policy_assignment.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttrSet("data.fmc_policy_assignment.test", "policy_name"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_assignment.test", "policy_type", "FTDNatPolicy"))
	checks = append(checks, resource.TestCheckResourceAttr("data.fmc_policy_assignment.test", "targets.0.type", "Device"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		ErrorCheck:               func(err error) error { return testAccErrorCheck(t, err) },
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceFmcPolicyAssignmentPrerequisitesConfig + testAccDataSourceFmcPolicyAssignmentConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			{
				Config: testAccDataSourceFmcPolicyAssignmentPrerequisitesConfig + testAccNamedDataSourceFmcPolicyAssignmentConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites

const testAccDataSourceFmcPolicyAssignmentPrerequisitesConfig = `
resource "fmc_ftd_nat_policy" "example" {
  name = "pa_nat_policy"
}

variable "device_id" { default = null } // tests will set $TF_VAR_device_id
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig

func testAccDataSourceFmcPolicyAssignmentConfig() string {
	config := `resource "fmc_policy_assignment" "test" {` + "\n"
	config += `	policy_id = fmc_ftd_nat_policy.example.id` + "\n"
	config += `	policy_type = "FTDNatPolicy"` + "\n"
	config += `	targets = [{` + "\n"
	config += `		id = var.device_id` + "\n"
	config += `		type = "Device"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_policy_assignment" "test" {
			id = fmc_policy_assignment.test.id
		}
	`
	return config
}

func testAccNamedDataSourceFmcPolicyAssignmentConfig() string {
	config := `resource "fmc_policy_assignment" "test" {` + "\n"
	config += `	policy_id = fmc_ftd_nat_policy.example.id` + "\n"
	config += `	policy_type = "FTDNatPolicy"` + "\n"
	config += `	targets = [{` + "\n"
	config += `		id = var.device_id` + "\n"
	config += `		type = "Device"` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "fmc_policy_assignment" "test" {
			policy_name = fmc_policy_assignment.test.policy_name
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig
