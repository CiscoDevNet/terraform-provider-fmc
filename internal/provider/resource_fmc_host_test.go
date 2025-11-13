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

func TestAccFmcHost(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("fmc_host.test", "name", "my_host"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_host.test", "type"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_host.test", "description", "My Host object"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_host.test", "ip", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_host.test", "overridable", "true"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcHostConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcHostConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName: "fmc_host.test",
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

func testAccFmcHostConfig_minimum() string {
	config := `resource "fmc_host" "test" {` + "\n"
	config += `	name = "my_host"` + "\n"
	config += `	ip = "10.1.1.1"` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcHostConfig_all() string {
	config := `resource "fmc_host" "test" {` + "\n"
	config += `	name = "my_host"` + "\n"
	config += `	description = "My Host object"` + "\n"
	config += `	ip = "10.1.1.1"` + "\n"
	config += `	overridable = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

// # FMCVERSION <= 7.2
// This test fails on FMC 7.2, as setting an empty description will set description as a single space.
// Which will trigger incorrect diff on next plan.
// func TestAccFmcHost_EmptyDescription(t *testing.T) {

// 	step_01 := `resource "fmc_host" "test" {` + "\n" +
// 		`	name = "hosts_1"` + "\n" +
// 		`	ip = "1.2.3.1"` + "\n" +
// 		`	description = "host1"` + "\n" +
// 		`}` + "\n"

// 	step_02 := `resource "fmc_host" "test" {` + "\n" +
// 		`	name = "hosts_1"` + "\n" +
// 		`	ip = "1.2.3.1"` + "\n" +
// 		`	description = ""` + "\n" +
// 		`}` + "\n"

// 	steps := []resource.TestStep{{
// 		Config: step_01,
// 	}, {
// 		Config: step_02,
// 	}}

// 	resource.Test(t, resource.TestCase{
// 		PreCheck:                 func() { testAccPreCheck(t) },
// 		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
// 		Steps:                    steps,
// 	})
// }
