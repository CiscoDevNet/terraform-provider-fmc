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

func TestAccFmcHosts(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_hosts.test", "items.my_hosts.id"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_hosts.test", "items.my_hosts.description", "My Host 1"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_hosts.test", "items.my_hosts.overridable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("fmc_hosts.test", "items.my_hosts.ip", "10.1.1.1"))
	checks = append(checks, resource.TestCheckResourceAttrSet("fmc_hosts.test", "items.my_hosts.type"))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccFmcHostsConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccFmcHostsConfig_all(),
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

func testAccFmcHostsConfig_minimum() string {
	config := `resource "fmc_hosts" "test" {` + "\n"
	config += `	items = { "my_hosts" = {` + "\n"
	config += `		ip = "10.1.1.1"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimal

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll

func testAccFmcHostsConfig_all() string {
	config := `resource "fmc_hosts" "test" {` + "\n"
	config += `	items = { "my_hosts" = {` + "\n"
	config += `		description = "My Host 1"` + "\n"
	config += `		overridable = true` + "\n"
	config += `		ip = "10.1.1.1"` + "\n"
	config += `	}}` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccFmcHosts_Sequential(t *testing.T) {

	step_01 := `resource "fmc_hosts" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"hosts_1" = {` + "\n" +
		`			ip = "1.2.3.1",` + "\n" +
		`			description = "host1"` + "\n" +
		`			overridable = true` + "\n" +
		`		},` + "\n" +
		`		"hosts_2" = {` + "\n" +
		`			ip = "1.2.3.2",` + "\n" +
		`		},` + "\n" +
		`		"hosts_3" = {` + "\n" +
		`			ip = "1.2.3.3",` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step01 []resource.TestCheckFunc
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_1.ip", "1.2.3.1"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_hosts.test", "items.hosts_1.type"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_1.description", "host1"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_2.ip", "1.2.3.2"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_hosts.test", "items.hosts_2.type"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_3.ip", "1.2.3.3"))
	checks_step01 = append(checks_step01, resource.TestCheckResourceAttrSet("fmc_hosts.test", "items.hosts_3.type"))

	step_02 := `resource "fmc_hosts" "test" {` + "\n" +
		`	items = {` + "\n" +
		`		"hosts_1" = {` + "\n" +
		`			ip = "1.2.3.1",` + "\n" +
		`			description = "host1 new description"` + "\n" +
		`		},` + "\n" +
		`		"hosts_2" = {` + "\n" +
		`			ip = "1.2.3.2",` + "\n" +
		`		},` + "\n" +
		`		"hosts_4" = {` + "\n" +
		`			ip = "1.2.3.4",` + "\n" +
		`		},` + "\n" +
		`	} ` + "\n" +
		`}` + "\n"

	var checks_step02 []resource.TestCheckFunc
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_1.ip", "1.2.3.1"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_1.description", "host1 new description"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_hosts.test", "items.hosts_1.type"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_2.ip", "1.2.3.2"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_hosts.test", "items.hosts_2.type"))
	checks_step02 = append(checks_step02, resource.TestCheckNoResourceAttr("fmc_hosts.test", "items.hosts_3"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttr("fmc_hosts.test", "items.hosts_4.ip", "1.2.3.4"))
	checks_step02 = append(checks_step02, resource.TestCheckResourceAttrSet("fmc_hosts.test", "items.hosts_4.type"))

	steps := []resource.TestStep{{
		Config: step_01,
		Check:  resource.ComposeTestCheckFunc(checks_step01...),
	}, {
		Config: step_02,
		Check:  resource.ComposeTestCheckFunc(checks_step02...),
	}}

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

// # FMCVERSION <= 7.2
// This test fails on FMC 7.2, as setting an empty description will set description as a single space.
// Which will trigger incorrect diff on next plan.
// func TestAccFmcHosts_EmptyDescription(t *testing.T) {

// 	step_01 := `resource "fmc_hosts" "test" {` + "\n" +
// 		`	items = {` + "\n" +
// 		`		"hosts_1" = {` + "\n" +
// 		`			ip = "1.2.3.1",` + "\n" +
// 		`			description = "host1"` + "\n" +
// 		`		},` + "\n" +
// 		`		"hosts_2" = {` + "\n" +
// 		`			ip = "1.2.3.2",` + "\n" +
// 		`			description = ""` + "\n" +
// 		`		},` + "\n" +
// 		`		"hosts_3" = {` + "\n" +
// 		`			ip = "1.2.3.3",` + "\n" +
// 		`		},` + "\n" +
// 		`	} ` + "\n" +
// 		`}` + "\n"

// 	step_02 := `resource "fmc_hosts" "test" {` + "\n" +
// 		`	items = {` + "\n" +
// 		`		"hosts_1" = {` + "\n" +
// 		`			ip = "1.2.3.1",` + "\n" +
// 		`			description = ""` + "\n" +
// 		`		},` + "\n" +
// 		`		"hosts_2" = {` + "\n" +
// 		`			ip = "1.2.3.2",` + "\n" +
// 		`			description = "host2"` + "\n" +
// 		`		},` + "\n" +
// 		`		"hosts_3" = {` + "\n" +
// 		`			ip = "1.2.3.3",` + "\n" +
// 		`			description = "host3"` + "\n" +
// 		`		},` + "\n" +
// 		`	} ` + "\n" +
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
