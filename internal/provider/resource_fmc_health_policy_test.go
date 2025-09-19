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

func TestAccFmcHealthPolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_health_policy.test", "name", "my_health_policy"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_health_policy.test", "description", "My health policy"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_health_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_health_policy.test", "policy_type", "DevicePolicy"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_health_policy.test", "is_default_policy", "false"))

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccFmcHealthPolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_health_policy.test",
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal
// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcHealthPolicyConfig_all() string {
	config := `resource "fmc_health_policy" "test" {` + "\n"
	config += `	name = "my_health_policy"` + "\n"
	config += `	description = "My health policy"` + "\n"
	config += `	policy_type = "DevicePolicy"` + "\n"
	config += `	is_default_policy = false` + "\n"
	config += `	health_module_run_time_interval = 10` + "\n"
	config += `	metric_collection_interval = 15` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
