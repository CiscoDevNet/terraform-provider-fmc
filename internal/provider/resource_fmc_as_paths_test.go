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

func TestAccFmcASPaths(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_as_paths.test", "items.240.id"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_as_paths.test", "items.240.type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_as_paths.test", "items.240.overridable", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_as_paths.test", "items.240.entries.0.action", "PERMIT"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_as_paths.test", "items.240.entries.0.regular_expression", "^(100|200)$"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcASPathsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcASPathsConfig_all(),
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
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimal

func testAccFmcASPathsConfig_minimum() string {
	config := `resource "fmc_as_paths" "test" {` + "\n"
	config += `	items = { "240" = {` + "\n"
	config += `		entries = [{` + "\n"
	config += `			action = "PERMIT"` + "\n"
	config += `			regular_expression = "^(100|200)$"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcASPathsConfig_all() string {
	config := `resource "fmc_as_paths" "test" {` + "\n"
	config += `	items = { "240" = {` + "\n"
	config += `		overridable = false` + "\n"
	config += `		entries = [{` + "\n"
	config += `			action = "PERMIT"` + "\n"
	config += `			regular_expression = "^(100|200)$"` + "\n"
	config += `		}]` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll
