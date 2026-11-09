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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc

func TestAccFmcAccessControlPolicyAdvancedLogging(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_access_control_policy_advanced_logging.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy_advanced_logging.test", "enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy_advanced_logging.test", "log_format", "CSV"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy_advanced_logging.test", "send_to_other_destination", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_access_control_policy_advanced_logging.test", "send_to_splunk", "true"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcAccessControlPolicyAdvancedLoggingPrerequisitesConfig + testAccFmcAccessControlPolicyAdvancedLoggingConfig_all(),
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

const testAccFmcAccessControlPolicyAdvancedLoggingPrerequisitesConfig = `
resource "fmc_access_control_policy" "test" {
  name           = "acp_advanced_logging"
  default_action = "BLOCK"
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcAccessControlPolicyAdvancedLoggingConfig_all() string {
	config := `resource "fmc_access_control_policy_advanced_logging" "test" {` + "\n"
	config += `	access_control_policy_id = fmc_access_control_policy.test.id` + "\n"
	config += `	enabled = true` + "\n"
	config += `	log_format = "CSV"` + "\n"
	config += `	send_to_other_destination = true` + "\n"
	config += `	send_to_splunk = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
