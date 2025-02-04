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

func TestAccFmcFilePolicy(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "name", "file_policy_1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_file_policy.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "description", "My file policy"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "first_time_file_analysis", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "custom_detection_list", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "clean_list", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "threat_score", "DISABLED"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "inspect_archives", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "block_encrypted_archives", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "block_uninspectable_archives", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "max_archive_depth", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "file_rules.0.application_protocol", "ANY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "file_rules.0.action", "DETECT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "file_rules.0.direction_of_transfer", "ANY"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "file_rules.0.file_categories.0.id", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "file_rules.0.file_categories.0.name", "PDF files"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "file_rules.0.file_types.0.id", "19"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_file_policy.test", "file_rules.0.file_types.0.name", "7Z"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcFilePolicyConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcFilePolicyConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_file_policy.test",
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

func testAccFmcFilePolicyConfig_minimum() string {
	config := `resource "fmc_file_policy" "test" {` + "\n"
	config += `	name = "file_policy_1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcFilePolicyConfig_all() string {
	config := `resource "fmc_file_policy" "test" {` + "\n"
	config += `	name = "file_policy_1"` + "\n"
	config += `	description = "My file policy"` + "\n"
	config += `	first_time_file_analysis = true` + "\n"
	config += `	custom_detection_list = true` + "\n"
	config += `	clean_list = true` + "\n"
	config += `	threat_score = "DISABLED"` + "\n"
	config += `	inspect_archives = false` + "\n"
	config += `	block_encrypted_archives = false` + "\n"
	config += `	block_uninspectable_archives = false` + "\n"
	config += `	max_archive_depth = 2` + "\n"
	config += `	file_rules = [{` + "\n"
	config += `		application_protocol = "ANY"` + "\n"
	config += `		action = "DETECT"` + "\n"
	config += `		direction_of_transfer = "ANY"` + "\n"
	config += `		file_categories = [{` + "\n"
	config += `			id = "5"` + "\n"
	config += `			name = "PDF files"` + "\n"
	config += `		}]` + "\n"
	config += `		file_types = [{` + "\n"
	config += `			id = "19"` + "\n"
	config += `			name = "7Z"` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
